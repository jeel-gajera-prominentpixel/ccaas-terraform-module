module "dynamodb-table" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-dynamodb?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
}
