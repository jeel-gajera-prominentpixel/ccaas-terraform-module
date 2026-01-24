locals {
  # Default tags that will be applied to all resources unless overridden by custom tags.
  tags = {
    company     = var.company
    environment = var.environment
    project     = var.project
    repository  = var.repository_url
    region      = var.region 
  }
}
