variable "openai_secret_key" {
  type      = string
  sensitive = true
}

variable "aws_region" {
  type    = string
  default = "eu-west-1"
}

variable "branch_name" {
  description = "The name of the Git branch"
  type        = string
  default     = ""
}

locals {
  sanitized_branch_name = replace(var.branch_name, "/", "-")
}
