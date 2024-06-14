resource "aws_iam_role" "lambda_role_wall" {
  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ],
    Version = "2012-10-17"
  })
}

resource "aws_iam_role_policy_attachment" "comprehend_policy_attachment" {
  role       = aws_iam_role.lambda_role_wall.name
  policy_arn = "arn:aws:iam::aws:policy/ComprehendFullAccess"
}

resource "aws_iam_policy" "lambda_cloudwatch_logs_policy_wall" {
  #tfsec:ignore:aws-iam-no-policy-wildcards
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect   = "Allow",
        Resource = "${aws_cloudwatch_log_group.lambda_log_group_wall.arn}:*"
      },
    ],
  })
}

resource "aws_iam_policy" "sagemaker_invoke_policy" {
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "sagemaker:InvokeEndpoint"
        ],
        Effect   = "Allow",
        Resource = data.aws_ssm_parameter.sagemaker_endpoint.value
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "sagemaker_invoke_policy_attachment" {
  role       = aws_iam_role.lambda_role_wall.name
  policy_arn = aws_iam_policy.sagemaker_invoke_policy.arn
}


resource "aws_iam_role_policy_attachment" "dynamodb_read_write_policy_attachment" {
  role       = aws_iam_role.lambda_role_wall.name
  policy_arn = aws_iam_policy.dynamodb_read_write_policy_wall.arn
}

resource "aws_cloudwatch_log_group" "lambda_log_group_wall" {
  #tfsec:ignore:aws-cloudwatch-log-group-customer-key
  name              = "/aws/lambda/${aws_lambda_function.aws_lambda_wall.function_name}"
  retention_in_days = 14
}

resource "aws_iam_role_policy_attachment" "xray_policy_attachment_wall" {
  role       = aws_iam_role.lambda_role_wall.name
  policy_arn = "arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess"
}

resource "aws_iam_role_policy_attachment" "lambda_cloudwatch_logs_attach_wall" {
  role       = aws_iam_role.lambda_role_wall.name
  policy_arn = aws_iam_policy.lambda_cloudwatch_logs_policy_wall.arn
}

resource "aws_lambda_function" "aws_lambda_wall" {
  function_name = "${terraform.workspace}-PromptDefender-Wall"

  handler          = "app.lambda_handler"
  filename         = data.archive_file.lambda_wall_zip.output_path
  role             = aws_iam_role.lambda_role_wall.arn
  runtime          = var.python_version
  source_code_hash = data.archive_file.lambda_wall_zip.output_base64sha256

  timeout = 120

  layers = [aws_lambda_layer_version.lambda_layer_wall.arn, aws_lambda_layer_version.langchain_lambda_layer.arn]

  tracing_config {
    mode = "Active"
  }

  environment {
    variables = {
      open_ai_api_key              = var.openai_secret_key
      SAGEMAKER_ENDPOINT_JAILBREAK = data.aws_ssm_parameter.sagemaker_endpoint_name.value
      CACHE_TABLE_NAME             = aws_dynamodb_table.cache_table.name
    }
  }
}

data "archive_file" "lambda_wall_zip" {
  type        = "zip"
  source_dir  = var.lambda_wall_path
  output_path = "wall_function.zip"
}

variable "lambda_wall_path" {
  type    = string
  default = "../cmd/lambda_wall_py/dist"
}

