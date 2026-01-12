module "ssm_store" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-ssm-parameter-store?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb-parameters"
  parameter_write = [
    {
      name        = "/cp/prod/app/database/master_password"
      value       = "password1"
      type        = "String"
      description = "Production database master password"
    }
  ]
  environment = "development"
  tags        = local.tags
}
