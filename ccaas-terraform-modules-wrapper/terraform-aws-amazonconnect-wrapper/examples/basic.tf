module "amazon_connect_basic" {
  source         = "git@github.com:CloverHealth/ccaas-terraform-modules-wrapper.git//terraform-aws-amazonconnect?ref=v1.0.2"
  prefix_company = "ch"
  lob            = "telesales"
  prefix_region  = "use1"
  application    = "connect"
  env            = "dev"
  tags           = local.tags
}
