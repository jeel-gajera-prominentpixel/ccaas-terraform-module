
module "route_53_records" {
  source       = "../../jb-ccaas-terraform-modules/terraform-aws-route53/modules/records/"
  zone_name    = var.zone_name
  zone_id      = var.zone_id
  private_zone = var.private_zone
  records      = var.records
}
