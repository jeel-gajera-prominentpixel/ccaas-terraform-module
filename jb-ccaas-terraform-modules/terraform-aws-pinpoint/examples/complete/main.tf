provider "aws" {
  region = "us-east-1"
}

module "pinpoint_devops_notifications" {
  source = "../../"
  name   = "devops-notifications"
}
