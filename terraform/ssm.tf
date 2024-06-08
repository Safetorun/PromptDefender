resource "aws_ssm_parameter" "openai_api_key" {
  name  = "${terraform.workspace}-OpenAI-API-Key"
  type  = "SecureString"
  value = var.openai_secret_key
}

data "aws_ssm_parameter" "sagemaker_endpoint" {
  name = "SagemakerEndpoint"
}


data "aws_ssm_parameter" "sagemaker_endpoint_name" {
  name = "SagemakerName"
}
