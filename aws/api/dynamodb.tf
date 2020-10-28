resource "aws_dynamodb_table" "table" {

  for_each = local.methods

  name           = each.value.lambda
  hash_key       = each.value.lambda
  billing_mode   = "PAY_PER_REQUEST"

  attribute {
    name = each.value.lambda
    type = "S"
  }

  tags = {
    Name = "dynamodb-api-${each.value.lambda}"
  }
}