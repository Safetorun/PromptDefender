resource "aws_iam_role" "lambda_role_builder" {
  name = "lambda_role_builder"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ],
    Version = "2012-10-17"
  })
}


resource "aws_lambda_function" "prompt_protect_builder" {
  function_name    = "PromptProtect-Builder"
  handler          = "main"
  role             = aws_iam_role.lambda_role_builder.arn
  filename         = data.archive_file.lambda_zip_builder.output_path
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_zip_builder.output_base64sha256

  environment {
    variables = {
      open_ai_api_key = var.openai_secret_key
    }
  }
}

data "archive_file" "lambda_zip_builder" {
  type        = "zip"
  source_file = "../lambda_prompt_builder/main"
  output_path = "function-2.zip"
}