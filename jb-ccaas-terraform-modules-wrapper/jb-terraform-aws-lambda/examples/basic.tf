module "lambda_function_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-lambda?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags           = local.tags
}
