locals {
  default_tags = {
    # application = var.application
    company    = var.company
    env        = var.environment
    repository = var.repository
    # created_by  = var.created_by
    project = var.project
    region  = var.region
  }

  tags = merge(local.default_tags, var.custom_tags)
}