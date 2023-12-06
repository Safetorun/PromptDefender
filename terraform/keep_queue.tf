resource "aws_sqs_queue" "keep_queue" {
  name = "${terraform.workspace}-keep_callback_queue"
  sqs_managed_sse_enabled = true
}
