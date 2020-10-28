provider "aws" {
  region = "eu-west-1"
}

module "api-gw" {
  source = "./api"

  api-gw-name = "metadata"
  stage_name  = "v1"
  endpoints = [
    {
      path   = "tag"
      method = "GET"
      lambda = "tag"
      request_parameters = {
        "method.request.header.Authorization" = true
        "method.request.querystring.team" = true
      }
    },
  ]
}

output "endpoint_id" {
  value = module.api-gw.endpoint_id
}