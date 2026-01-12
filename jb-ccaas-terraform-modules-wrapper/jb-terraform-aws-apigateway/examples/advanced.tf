module "api-gateway_advance" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-apigateway?ref=main"
  prefix_company = "jb"
  protocol_type  = "HTTP"
  prefix_region  = "usw2"
  lob            = "itsd"
  application    = "recordings"
  env            = "sandbox"

  # vpc-link-http
  create_api_domain_name = false
  vpc_links              = {}

  default_stage_access_log_destination_arn = null
  default_stage_access_log_format          = null
  domain_name                              = "example.com"
  authorizers                              = {}
  domain_name_certificate_arn              = "arn:aws:acm:eu-south-1:123456789102:certificate/12345678-1234-1234-1234-123456789012"
  cors_configuration                       = {}
  integrations                             = {}
  tags                                     = local.tags
}
