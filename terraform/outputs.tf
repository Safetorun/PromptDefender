output "api_url" {
  value = aws_api_gateway_deployment.api.invoke_url
}

output "workspace" {
  value = terraform.workspace
}

output "api_key_value" {
  value     = aws_api_gateway_api_key.api_key.value
  sensitive = true
}