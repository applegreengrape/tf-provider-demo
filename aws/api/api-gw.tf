locals {
  response_codes = toset([{
    status_code         = 200
    response_templates  = {} # TODO: Fill this in
    response_parameters = {} # TODO: Fill this in
  }])
  endpoints = toset(var.endpoints.*.path)
  methods = {
    for e in var.endpoints : "${e.method} ${e.path}" => e
  }
  responses = {
    for pair in setproduct(var.endpoints, local.response_codes) :
    "${pair[0].method} ${pair[0].path} ${pair[1].status_code}" => {
      method              = pair[0].method
      path                = pair[0].path
      method_key          = "${pair[0].method} ${pair[0].path}" # key for local.methods
      status_code         = pair[1].status_code
      response_templates  = pair[1].response_templates
      response_parameters = pair[1].response_parameters
    }
  }
}

resource "aws_api_gateway_rest_api" "api-gw-module" {
  name        = var.api-gw-name
  description = "api-gw"
}
resource "aws_api_gateway_resource" "api-gw-module" {
  for_each = local.endpoints

  rest_api_id = aws_api_gateway_rest_api.api-gw-module.id
  parent_id   = aws_api_gateway_rest_api.api-gw-module.root_resource_id
  path_part   = each.value

}
resource "aws_api_gateway_method" "api-gw-module" {
  for_each = local.methods

  rest_api_id   = aws_api_gateway_rest_api.api-gw-module.id
  resource_id   = aws_api_gateway_resource.api-gw-module[each.value.path].id
  http_method   = each.value.method
  authorization = "NONE"

  request_parameters = each.value.request_parameters
}
resource "aws_api_gateway_integration" "api-gw-module" {
  for_each = local.methods

  rest_api_id             = aws_api_gateway_rest_api.api-gw-module.id
  resource_id             = aws_api_gateway_resource.api-gw-module[each.value.path].id
  http_method             = aws_api_gateway_method.api-gw-module["${each.value.method} ${each.value.path}"].http_method
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.function["${each.value.method} ${each.value.path}"].invoke_arn
  integration_http_method = "POST"
}
resource "aws_api_gateway_integration_response" "api-gw-module" {
  for_each = local.responses

  rest_api_id         = aws_api_gateway_integration.api-gw-module[each.value.method_key].rest_api_id
  resource_id         = aws_api_gateway_integration.api-gw-module[each.value.method_key].resource_id
  http_method         = each.value.method
  status_code         = each.value.status_code
  response_parameters = each.value.response_parameters
  response_templates  = each.value.response_templates
}
resource "aws_api_gateway_deployment" "api-gw-module" {
  rest_api_id = aws_api_gateway_rest_api.api-gw-module.id
  stage_name  = var.stage_name
  depends_on = [

    aws_dynamodb_table.table,
    aws_s3_bucket.lambda,
    aws_s3_bucket_object.key,
    aws_iam_role.lambda,
    aws_lambda_function.function,
    aws_api_gateway_resource.api-gw-module,
    aws_api_gateway_method.api-gw-module,
    aws_api_gateway_integration.api-gw-module,
    aws_api_gateway_integration_response.api-gw-module,
  ]
}

resource "aws_lambda_permission" "apigw_lambda" {
  for_each = local.methods

  action        = "lambda:InvokeFunction"
  function_name = "arn:aws:lambda:${var.aws_region}:${var.account_id}:function:${each.value.lambda}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api-gw-module.execution_arn}/*/*"

  depends_on = [
    aws_api_gateway_deployment.api-gw-module,
  ]
}

output "execution_arn" {
  value = aws_api_gateway_rest_api.api-gw-module.execution_arn
  depends_on = [
    aws_api_gateway_deployment.api-gw-module,
  ]
}

output "endpoint_id" {
  value = "https://${aws_api_gateway_rest_api.api-gw-module.id}.execute-api.${var.aws_region}.amazonaws.com/${var.stage_name}/"
  depends_on = [
    aws_api_gateway_deployment.api-gw-module,
    aws_lambda_permission.apigw_lambda
  ]
}