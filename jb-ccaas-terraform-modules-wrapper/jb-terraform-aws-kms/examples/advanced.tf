module "aws_kms" {
  source                = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kms?ref=main"
  prefix_company        = "jb"
  service               = "s3"
  lob                   = "itsd"
  prefix_region         = "usw2"
  application           = "recordings"
  env                   = "sandbox"
  description           = "Primary key of replica key example"
  multi_region          = true
  enable_default_policy = true
  tags                  = local.tags
}
