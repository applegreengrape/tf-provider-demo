variable "stage_name" {
  default = "v1"
}

variable "api-gw-name" {
  default = "pz_test"
}

variable "aws_region" {
  default = "eu-west-1"
}

variable "account_id" {
  default = "900665556514"
}

variable "endpoints" {
  type = list(object({
    path               = string
    method             = string
    lambda             = string
    request_parameters = object({})

  }))
}