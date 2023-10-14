resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"

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

resource "aws_iam_role_policy_attachment" "comprehend_policy_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/ComprehendFullAccess"
}


resource "aws_lambda_function" "aws_Lambda_keep" {
  function_name    = "PromptDefender-Keep"
  handler          = "main"
  role             = aws_iam_role.lambda_role.arn
  filename         = data.archive_file.lambda_keep_zip.output_path
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_keep_zip.output_base64sha256

  environment {
    variables = {
      open_ai_api_key = var.openai_secret_key
    }
  }
}

data "archive_file" "lambda_keep_zip" {
  type        = "zip"
  source_file = var.lambda_keep_path
  output_path = "keep_function.zip"
}

variable "lambda_keep_path" {
  type    = string
  default = "../deployments/aws/lambda_keep/main"
}