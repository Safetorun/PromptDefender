resource "aws_api_gateway_rest_api" "api" {
  name        = "PromptProtect"
  description = "My API Service"
  body        = local_file.built_open_api_spec.content
  depends_on = [aws_lambda_function.prompt_protect, local_file.built_open_api_spec]
}

resource "aws_lambda_permission" "apigw_lambda_permission_protect" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.prompt_protect.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_lambda_permission" "apigw_lambda_permission_shield" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.prompt_protect_builder.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_deployment" "api" {
  depends_on = [aws_api_gateway_rest_api.api]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "prod"
}

resource "local_file" "built_open_api_spec" {
  filename = "../openapi.yml"
  content = templatefile("../openapi.yml.tpl", {
    prompt_shield_lambda_arn = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.prompt_protect.arn}/invocations",
    prompt_shield_builder_lambda_arn = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.prompt_protect_builder.arn}/invocations"
  })
}
output "api_url" {
  value = aws_api_gateway_deployment.api.invoke_url
}
