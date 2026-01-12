provider "aws" {
  region = "eu-west-1"

  # Make it faster by skipping something
  skip_metadata_api_check     = true
  skip_region_validation      = true
  skip_credentials_validation = true
}

# Assume VPC/subnets/SG already exist
# data "aws_subnets" "private" { ... }
# resource "aws_security_group" "svc" { ... }

module "fargate_service" {
  source = "../.."

  name        = "demo-api"
  region      = "us-east-1"
  cluster_arn = "arn:aws:ecs:us-east-1:123456789012:cluster/main"

  image              = "123456789012.dkr.ecr.us-east-1.amazonaws.com/demo:latest"
  cpu                = "512"
  memory             = "1024"
  desired_count      = 1
  container_port     = 8080
  platform_version   = "1.4.0"

  subnet_ids         = ["subnet-aaaa", "subnet-bbbb"]
  security_group_ids = ["sg-0123456789abcdef0"]
  assign_public_ip   = false

  environment = { ENV = "stg" }

  # load_balancers = [
  #   { target_group_arn = "arn:aws:elasticloadbalancing:...", container_port = 8080 }
  # ]
}
