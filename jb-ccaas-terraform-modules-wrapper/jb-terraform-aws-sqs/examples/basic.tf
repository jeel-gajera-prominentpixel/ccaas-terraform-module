module "sqs_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sqs?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb_sqs"
  fifo_queue     = true
  create_dlq     = true
}
