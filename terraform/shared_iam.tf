resource "aws_iam_policy" "dynamodb_read_write_policy_wall" {
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem"
        ],
        Effect   = "Allow",
        Resource = aws_dynamodb_table.cache_table.arn
      },
    ],
  })
}

resource "aws_iam_policy" "ssm_read_policy_wall" {
  policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Action = [
          "ssm:GetParameter",
          "ssm:GetParameters",
        ],
        Effect   = "Allow",
        Resource = aws_ssm_parameter.openai_api_key.arn
      },
    ],
  })
}

