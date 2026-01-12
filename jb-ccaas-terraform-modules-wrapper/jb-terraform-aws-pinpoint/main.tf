module "pinpoint" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-pinpoint"
  name   = var.name == null ? local.pinpoint_name : var.name
  email  = var.email
  sms    = var.sms
}
