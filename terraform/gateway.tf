resource "aws_api_gateway_rest_api" "api" {
  name        = "PromptProtect"
  description = "My API Service"
}

resource "aws_api_gateway_resource" "verify_resource" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  path_part   = "verify"
}

resource "aws_api_gateway_method" "verify_method" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.verify_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_lambda_permission" "apigw_lambda_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.prompt_protect.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_integration" "lambda" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_resource.verify_resource.id
  http_method             = aws_api_gateway_method.verify_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.prompt_protect.invoke_arn
}

resource "aws_api_gateway_deployment" "api" {
  depends_on = [aws_api_gateway_integration.lambda]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "prod"
}

output "api_url" {
  value = aws_api_gateway_deployment.api.invoke_url
}