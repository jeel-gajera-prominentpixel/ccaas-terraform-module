module "acm" {
  source                    = "../../jb-ccaas-terraform-modules/terraform-aws-acm"
  domain_name               = var.domain_name
  zone_id                   = var.zone_id
  subject_alternative_names = var.subject_alternative_names
  validation_method         = var.validation_method
  create_route53_records    = var.create_route53_records
  validation_record_fqdns   = var.validation_record_fqdns
  wait_for_validation       = var.wait_for_validation
  tags                      = local.tags
}
