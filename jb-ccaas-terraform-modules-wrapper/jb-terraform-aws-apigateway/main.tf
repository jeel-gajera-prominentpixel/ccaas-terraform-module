module "api-gateway" {
  source        = "../../jb-ccaas-terraform-modules/terraform-aws-apigateway-v2"
  name          = var.name == "" ? local.apigateway_name : var.name
  protocol_type = var.protocol_type

  # vpc-link-http
  create_api_domain_name = var.create_api_domain_name
  vpc_links              = var.vpc_links
  ##############################

  # Access logs
  default_stage_access_log_destination_arn = var.default_stage_access_log_destination_arn
  default_stage_access_log_format          = var.default_stage_access_log_format
  create_routes_and_integrations           = var.create_routes_and_integrations
  domain_name                              = var.domain_name
  domain_name_certificate_arn              = var.domain_name_certificate_arn
  cors_configuration                       = var.cors_configuration
  integrations                             = var.integrations
  authorizers                              = var.authorizers
  tags = merge(local.tags, {
    Name = var.name == "" ? local.apigateway_name : var.name
  })
}
