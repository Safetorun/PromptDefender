variable "openai_secret_key" {
  type      = string
  sensitive = true
}

variable "aws_region" {
  type    = string
  default = "eu-west-1"
}

variable "huggingface_api_key" {
  type      = string
  sensitive = true
}