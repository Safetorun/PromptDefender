resource "aws_dynamodb_table" "cache_table" {
  name           = "${terraform.workspace}-Prompt-Defender-cache_table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "S"
  }
}