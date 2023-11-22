resource "aws_sqs_queue" "keep_queue" {
  name = "${local.sanitized_branch_name}-keep_callback_queue"
}
