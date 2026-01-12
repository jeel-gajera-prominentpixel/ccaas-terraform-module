
module "wafv2" {
  source                        = "../../jb-ccaas-terraform-modules/terraform-aws-wafv2"
  name                          = var.name == "" ? local.wafv2_name : var.name
  description                   = var.description
  scope                         = var.scope
  default_action                = var.default_action
  waf_rules                     = var.waf_rules
  waf_ip_sets                   = var.waf_ip_sets
  visibility_config             = var.visibility_config
  enabled_logging_configuration = var.enabled_logging_configuration
  redacted_fields               = var.redacted_fields
  logging_filter                = var.logging_filter
  tags = merge(local.tags, {
    Name = var.name == "" ? local.wafv2_name : var.name
  })
}
