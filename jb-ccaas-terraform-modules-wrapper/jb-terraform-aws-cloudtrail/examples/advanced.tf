module "cloudtrail_log_group_advance" {
  source                        = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-cloudtrail?ref=main"
  prefix_company                = "jb"
  lob                           = "itsd"
  prefix_region                 = "usw2"
  application                   = "recordings"
  env                           = "sandbox"
  enable_logging                = true
  enable_log_file_validation    = true
  include_global_service_events = false
  is_multi_region_trail         = true
  is_organization_trail         = false
  tags                          = local.tags
}
