resource "aws_api_gateway_api_key" "api_key" {
  name    = "${terraform.workspace}-TestAPIKey"
  enabled = true
}

resource "aws_ssm_parameter" "defender_api_gateway_usage_plan_id" {
  name  = "${terraform.workspace}-defender_api_gateway_usage_plan_id"
  type  = "SecureString"
  value = aws_api_gateway_usage_plan.usage_plan.id
}

resource "aws_api_gateway_usage_plan" "usage_plan" {
  name = "PromptDefenderUsagePlan"

  api_stages {
    api_id = aws_api_gateway_rest_api.api.id
    stage  = aws_api_gateway_stage.api_stage.stage_name
  }
  depends_on = [aws_api_gateway_deployment.api]
}


resource "aws_api_gateway_usage_plan_key" "usage_plan_key" {
  key_id        = aws_api_gateway_api_key.api_key.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.usage_plan.id
}

resource "aws_api_gateway_method_settings" "method_settings" { #tfsec:ignore:aws-api-gateway-enable-cache
  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = aws_api_gateway_stage.api_stage.stage_name
  method_path = "*/*"

  settings {
    logging_level      = "ERROR"
    metrics_enabled    = true
    data_trace_enabled = false
    caching_enabled    = false
  }
}