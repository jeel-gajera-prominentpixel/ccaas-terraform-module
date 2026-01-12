terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "= 5.76"
    }
  }
}

provider "aws" {
  assume_role {
    role_arn = "arn:aws:iam::381492173985:role/${var.role_name}"
  }
  region = var.region
}
