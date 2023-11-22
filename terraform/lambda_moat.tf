resource "aws_iam_role" "lambda_role_moat" {
  name = "${terraform.workspace}-lambda_role_moat"

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
  role       = aws_iam_role.lambda_role_moat.name
  policy_arn = "arn:aws:iam::aws:policy/ComprehendFullAccess"
}


resource "aws_lambda_function" "aws_lambda_moat" {
  function_name    = "${terraform.workspace}-PromptDefender-Moat"
  handler          = "main"
  role             = aws_iam_role.lambda_role_moat.arn
  filename         = data.archive_file.lambda_moat_zip.output_path
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_moat_zip.output_base64sha256

  environment {
    variables = {
      open_ai_api_key = var.openai_secret_key
    }
  }
}

data "archive_file" "lambda_moat_zip" {
  type        = "zip"
  source_file = var.lambda_moat_path
  output_path = "moat_function.zip"
}


variable "lambda_moat_path" {
  type    = string
  default = "../deployments/aws/lambda_moat/main"
}