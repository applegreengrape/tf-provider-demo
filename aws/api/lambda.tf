resource "aws_s3_bucket" "lambda" {
  for_each = local.methods

  bucket = "${var.api-gw-name}-${each.value.lambda}"
  acl    = "private"

  versioning {
    enabled = true
  }

  provisioner "local-exec" {
    command = "cd api/lambda && zip ${each.value.lambda}.zip ${each.value.lambda}.py"
  }
}

resource "aws_s3_bucket_object" "key" {
  for_each = local.methods

  bucket = "${var.api-gw-name}-${each.value.lambda}"
  key    = "${var.stage_name}/${each.value.lambda}.zip"
  source = "api/lambda/${each.value.lambda}.zip"
}

resource "aws_lambda_function" "function" {
  for_each = local.methods

  function_name = each.value.lambda
  s3_bucket     = "${var.api-gw-name}-${each.value.lambda}"
  s3_key        = "${var.stage_name}/${each.value.lambda}.zip"

  handler = "${each.value.lambda}.handler"
  runtime = "python3.7"

  role = aws_iam_role.lambda.arn

  environment {
    variables = {
      apiTok = "UhaGae289taAmjYZCEhj"
    }
  }

  depends_on = [
    aws_s3_bucket.lambda,
    aws_s3_bucket_object.key,
  ]
}

data "aws_iam_policy_document" "lambda-assume-role-policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = "terraform-serverless-lambda"
  assume_role_policy = data.aws_iam_policy_document.lambda-assume-role-policy.json
}

data "aws_iam_policy_document" "lambda-role-policy" {
  dynamic "statement" {
    for_each = local.methods
    content {
      actions = [
        "dynamodb:BatchGetItem",
        "dynamodb:GetItem",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:BatchWriteItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem"
      ]
      resources = [
        "arn:aws:dynamodb:${var.aws_region}:${var.account_id}:table/${statement.value["lambda"]}",
        "arn:aws:dynamodb:${var.aws_region}:${var.account_id}:table/${statement.value["lambda"]}/*/*"
      ]
    }
  }
}

resource "aws_iam_policy" "lambda"{
  name        = "lambda_policy"
  policy = data.aws_iam_policy_document.lambda-role-policy.json
}

resource "aws_iam_role_policy_attachment" "lambda" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.lambda.arn
}