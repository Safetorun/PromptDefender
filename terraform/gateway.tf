resource "aws_api_gateway_rest_api" "api" {
  name        = "${terraform.workspace}-PromptProtect"
  description = "My API Service"
  body        = local_file.built_open_api_spec.content
  depends_on  = [aws_lambda_function.aws_Lambda_keep, local_file.built_open_api_spec]
}

resource "aws_api_gateway_deployment" "api" {
  depends_on  = [aws_api_gateway_rest_api.api]
  rest_api_id = aws_api_gateway_rest_api.api.id
  triggers    = {
    redeployment = sha256(jsonencode(aws_api_gateway_rest_api.api.body))
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_cloudwatch_log_group" "api_logs" { #tfsec:ignore:aws-cloudwatch-log-group-customer-key
  name              = "/aws/api_gateway/${aws_api_gateway_rest_api.api.name}"
  retention_in_days = 14
}

resource "aws_api_gateway_stage" "api_stage" {
  stage_name    = "prod"
  rest_api_id   = aws_api_gateway_rest_api.api.id
  deployment_id = aws_api_gateway_deployment.api.id

  xray_tracing_enabled = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_logs.arn
    format          = "{\"apiId\":\"$context.apiId\",\"requestId\":\"$context.requestId\",\"requestTime\":\"$context.requestTime\",\"protocol\":\"$context.protocol\",\"httpMethod\":\"$context.httpMethod\",\"resourcePath\":\"$context.path\",\"requestHostHeader\":\"$context.domainName\",\"requestUserAgentHeader\":\"$context.identity.userAgent\",\"ip\":\"$context.identity.sourceIp\",\"status\":\"$context.status\",\"responseLength\":\"$context.responseLength\",\"durationMs\":\"$context.responseLatency\",\"caller\":\"$context.identity.caller\",\"user\":\"$context.identity.user\",\"principalId\":\"$context.authorizer.principalId\",\"cognitoIdentityId\":\"$context.identity.cognitoIdentityId\",\"userArn\":\"$context.identity.userArn\",\"apiKeyId\":\"$context.identity.apiKeyId\",\"metadata\":{\"apiKey\":\"$context.identity.apiKey\",\"stage\":\"$context.stage\"}}"
  }
}



resource "aws_lambda_permission" "apigw_lambda_permission_protect" {
  statement_id  = "${terraform.workspace}-AllowAPIGatewayInvoke-keep"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.aws_Lambda_keep.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api.execution_arn}/prod/*"
}

resource "aws_lambda_permission" "apigw_lambda_permission_shield" {
  statement_id  = "${terraform.workspace}-AllowAPIGatewayInvoke-wall"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.aws_lambda_wall.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api.execution_arn}/prod/*"
}

resource "aws_lambda_permission" "apigw_lambda_permission_user" {
  statement_id  = "${terraform.workspace}-AllowAPIGatewayInvoke-user"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.aws_lambda_user.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api.execution_arn}/prod/*"
}

resource "local_file" "built_open_api_spec" {
  filename = "../api/openapi.yml"
  content  = templatefile("../api/openapi.yml.tpl", {
    lambda_keep_arn = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.aws_Lambda_keep.arn}/invocations",
    lambda_wall_arn = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.aws_lambda_wall.arn}/invocations",
    lambda_user_arn = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${aws_lambda_function.aws_lambda_user.arn}/invocations",
  })
}

