resource "aws_iam_role" "lambda_role_keep" {
  name = "lambda_role_keep"

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

resource "aws_lambda_function" "aws_Lambda_keep" {
  function_name    = "PromptDefender-Keep"
  handler          = "main"
  role             = aws_iam_role.lambda_role_keep.arn
  filename         = data.archive_file.lambda_keep_zip.output_path
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_keep_zip.output_base64sha256


  environment {
    variables = {
      open_ai_api_key = var.openai_secret_key
      keep_sqs_queue_url  = aws_sqs_queue.keep_queue.url
    }
  }
}

resource "aws_iam_policy" "lambda_sqs_policy" {
  name   = "lambda_sqs_policy"
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action   = "sqs:SendMessage",
        Effect   = "Allow",
        Resource = aws_sqs_queue.keep_queue.arn
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "lambda_sqs_attach" {
  role       = aws_iam_role.lambda_role_keep.name
  policy_arn = aws_iam_policy.lambda_sqs_policy.arn
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