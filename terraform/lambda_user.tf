resource "aws_iam_role" "lambda_role_user" {
  name = "${terraform.workspace}-lambda_role_user"

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


resource "aws_iam_policy" "lambda_cloudwatch_logs_policy_user" { #tfsec:ignore:aws-iam-no-policy-wildcards
  name   = "${terraform.workspace}-lambda_cloudwatch_logs_policy_user"
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect   = "Allow",
        Resource = "${aws_cloudwatch_log_group.lambda_log_group_user.arn}:*"
      },
    ],
  })
}

resource "aws_cloudwatch_log_group" "lambda_log_group_user" { #tfsec:ignore:aws-cloudwatch-log-group-customer-key
  name              = "/aws/lambda/${aws_lambda_function.aws_lambda_user.function_name}"
  retention_in_days = 14
}

resource "aws_iam_role_policy_attachment" "xray_policy_attachment_user" {
  role       = aws_iam_role.lambda_role_user.name
  policy_arn = "arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess"
}

resource "aws_iam_role_policy_attachment" "lambda_cloudwatch_logs_attach_user" {
  role       = aws_iam_role.lambda_role_user.name
  policy_arn = aws_iam_policy.lambda_cloudwatch_logs_policy_user.arn
}

resource "aws_lambda_function" "aws_lambda_user" {
  function_name    = "${terraform.workspace}-PromptDefender-User"
  handler          = "bootstrap"
  role             = aws_iam_role.lambda_role_user.arn
  filename         = data.archive_file.lambda_user_zip.output_path
  runtime          = "provided.al2"
  source_code_hash = data.archive_file.lambda_user_zip.output_base64sha256

  timeout = 60

  layers = ["arn:aws:lambda:${var.aws_region}:901920570463:layer:aws-otel-collector-amd64-ver-0-90-1:1"]

  tracing_config {
    mode = "Active"
  }

  environment {
    variables = {
      USERS_TABLE = aws_dynamodb_table.UserAndSessionDb.name
    }
  }
}

resource "aws_iam_policy" "lambda_dynamodb_access" {
  name   = "${terraform.workspace}-lambda_dynamodb_access"
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan",
          "dynamodb:Query"
        ],
        Effect   = "Allow",
        Resource = aws_dynamodb_table.UserAndSessionDb.arn
      },
    ],
  })
}

resource "aws_iam_policy" "lambda_dynamodb_access-index" {
  name   = "${terraform.workspace}-lambda_dynamodb_access--index"
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action = [
          "dynamodb:Query"
        ],
        Effect   = "Allow",
        Resource = "arn:aws:dynamodb:${var.aws_region}:${data.aws_caller_identity.current.account_id}:table/${aws_dynamodb_table.UserAndSessionDb.name}/index/ApiKeyId-index"
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_access_attach" {
  role       = aws_iam_role.lambda_role_user.name
  policy_arn = aws_iam_policy.lambda_dynamodb_access.arn
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_access_attach--index" {
  role       = aws_iam_role.lambda_role_user.name
  policy_arn = aws_iam_policy.lambda_dynamodb_access-index.arn
}

data "archive_file" "lambda_user_zip" {
  type        = "zip"
  source_file = var.lambda_user_path
  output_path = "user_function.zip"
}

variable "lambda_user_path" {
  type    = string
  default = "../cmd/lambda_user/bootstrap"
}
