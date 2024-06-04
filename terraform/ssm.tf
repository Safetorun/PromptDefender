resource "aws_ssm_parameter" "openai_api_key" {
  name  = "${terraform.workspace}-OpenAI-API-Key"
  type  = "SecureString"
  value = var.openai_secret_key
}
