module "rds_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-proxy?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  engine_family  = "POSTGRESQL"
  name           = "rds-proxy-example-addvance"
  debug_logging  = false
}
