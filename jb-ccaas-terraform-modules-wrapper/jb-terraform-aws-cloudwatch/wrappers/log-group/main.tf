module "wrapper" {
  source            = "../../jb-ccaas-terraform-modules/terraform-aws-cloudwatch/modules/log-group"
  create            = var.create
  kms_key_id        = var.kms_key_id
  log_group_class   = var.log_group_class
  name              = var.name
  name_prefix       = var.name_prefix
  retention_in_days = var.retention_in_days
  skip_destroy      = var.skip_destroy
  tags              = var.tags
}
