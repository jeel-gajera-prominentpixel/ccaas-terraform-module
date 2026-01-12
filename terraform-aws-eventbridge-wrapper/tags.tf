module "tags" {
  source = "../ccaas-terraform-modules/terraform-aws-tags"
  # application    = var.application
  company     = var.company
  environment = var.environment
  custom_tags = var.custom_tags
  project     = var.project == "" ? var.application : var.project
  repository  = var.repository_url
  region      = var.region == "" ? var.region_suffix : var.region
}
