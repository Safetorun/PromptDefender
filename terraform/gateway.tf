resource "aws_api_gateway_rest_api" "api" {
  name        = "PromptProtect"
  description = "My API Service"
  body        = templatefile("../openapi.yml", {lambda_arn = aws_lambda_function.prompt_protect.arn})
  depends_on  = [aws_lambda_function.prompt_protect]
}

resource "aws_lambda_permission" "apigw_lambda_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.prompt_protect.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_deployment" "api" {
  depends_on = [aws_api_gateway_rest_api.api]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "prod"

}

output "api_url" {
  value = aws_api_gateway_deployment.api.invoke_url
}
