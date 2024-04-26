resource "aws_dynamodb_table" "cache_table" { #tfsec:ignore:aws-dynamodb-table-customer-key
  name           = "${terraform.workspace}-Prompt-Defender-cache_table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "Id"

  point_in_time_recovery {
    enabled = true
  }

  server_side_encryption {
    enabled = true
  }

  attribute {
    name = "Id"
    type = "S"
  }
}