provider "aws" {
  region = var.region

  default_tags {
    tags = {
      repo        = "ipfs-metadata"
    }
  }
}

data "aws_caller_identity" "current" {}
data "aws_availability_zones" "available" {}

variable "region" {
  type    = string
  default = "us-east-1"
}
