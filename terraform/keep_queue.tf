resource "aws_sqs_queue" "keep_queue" {
  name = "keep_callback_queue"
}
