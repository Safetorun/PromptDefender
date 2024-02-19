resource "aws_dynamodb_table" "UserAndSessionDb" { #tfsec:ignore:aws-dynamodb-table-customer-key
  name           = "${terraform.workspace}-UserAndSessionDb"
  billing_mode = "PAY_PER_REQUEST"
  hash_key       = "UserOrSessionId"
  range_key      = "ApiKeyId"

  server_side_encryption {
    enabled = true
  }

  attribute {
    name = "UserOrSessionId"
    type = "S"
  }

  attribute {
    name = "ApiKeyId"
    type = "S"
  }

  point_in_time_recovery {
    enabled = true
  }

  global_secondary_index {
    hash_key        = "ApiKeyId"
    name            = "ApiKeyId-index"
    projection_type = "ALL"
  }
}