module "aws_cloudtrail" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-cloudtrail"

  name                          = var.name == "" ? local.cloudtrail : var.name
  enable_logging                = var.enable_logging
  enable_log_file_validation    = var.enable_log_file_validation
  include_global_service_events = var.include_global_service_events
  is_multi_region_trail         = var.is_multi_region_trail
  is_organization_trail         = var.is_organization_trail
  s3_bucket_name                = var.s3_bucket_name
  s3_key_prefix                 = var.s3_key_prefix
  advanced_event_selector       = var.advanced_event_selector


  tags = merge(local.tags, {
    Name = local.cloudtrail
  })
}

# module "cloudtrail_s3_bucket" {
#   source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket?ref=feature/add-new-cloudtrail-module"
#   force_destroy  = true
#   prefix_company = var.prefix_company
#   prefix_region  = var.prefix_region
#   lob            = var.lob
#   application    = var.application
#   env            = var.env
#   tags = merge(local.tags, {
#     Name = var.name == "" ? local.cloudtrail : var.name
#   })
# }
