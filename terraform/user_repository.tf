resource "aws_dynamodb_table" "UserAndSessionDb" { #tfsec:ignore:aws-dynamodb-table-customer-key
  name           = "${terraform.workspace}-UserAndSessionDb"
  read_capacity  = 10
  write_capacity = 5
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
}