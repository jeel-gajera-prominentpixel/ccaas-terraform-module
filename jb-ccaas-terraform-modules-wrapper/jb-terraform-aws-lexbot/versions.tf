
terraform {
  required_version = ">= 1.3.0, < 2.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.27"
    }
    external = {
      source  = "hashicorp/external"
      version = ">= 2.3.0"
    }
    awscc = {
      source  = "hashicorp/awscc"
      version = "= 1.0.0"
    }
  }
}
