resource "aws_api_gateway_api_key" "api_key" {
  name = "${local.sanitized_branch_name}-TestAPIKey"
  enabled = true
}

resource "aws_api_gateway_usage_plan" "usage_plan" {
  name = "PromptDefenderUsagePlan"

  api_stages {
    api_id = aws_api_gateway_rest_api.api.id
    stage  = aws_api_gateway_deployment.api.stage_name
  }
}

resource "aws_api_gateway_usage_plan_key" "usage_plan_key" {
  key_id        = aws_api_gateway_api_key.api_key.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.usage_plan.id
}

resource "aws_api_gateway_method_settings" "method_settings" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = aws_api_gateway_deployment.api.stage_name
  method_path = "*/*"

  settings {
  }
}

output "api_key_value" {
  value = aws_api_gateway_api_key.api_key.value
  sensitive = true
}