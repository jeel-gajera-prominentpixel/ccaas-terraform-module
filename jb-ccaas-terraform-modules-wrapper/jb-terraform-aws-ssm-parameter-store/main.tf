module "ssm_store" {
  source              = "../../jb-ccaas-terraform-modules/terraform-aws-ssm-parameter-store"
  name                = var.name == "" ? local.ssm_parameter_name : var.name
  parameter_write     = var.parameter_write
  parameter_read      = var.parameter_read
  kms_arn             = var.kms_arn
  environment         = var.environment
  regex_replace_chars = var.regex_replace_chars
  tags = merge(local.tags, {
    Name = var.name == "" ? local.ssm_parameter_name : var.name
  })
}
