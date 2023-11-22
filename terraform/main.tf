terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  backend "s3" {
    bucket = "tf-state-bucket-prompt-shield"
    key    = "tfstate"
    region = "eu-west-1"
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "eu-west-1"

  default_tags {
    tags = {
      Repo   = "https://github.com/safetorun/PromptDefender"
      Branch = local.sanitized_branch_name
      Workspace = terraform.workspace
    }
  }
}

variable "commit_version" {
  type = string
}