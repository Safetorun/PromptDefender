terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.40"
    }
  }

  backend "s3" {
    bucket = "tf-state-bucket-prompt-shield-base"
    key    = "tfstate"
    region = "eu-west-1"
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "eu-west-1"

  default_tags {
    tags = {
      Repo      = "https://github.com/safetorun/PromptDefender"
      Workspace = terraform.workspace
    }
  }
}