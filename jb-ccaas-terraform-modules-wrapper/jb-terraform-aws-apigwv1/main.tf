
module "apigv1" {
  source                       = "../../jb-ccaas-terraform-modules/terraform-aws-apigwv1"
  api_name                     = var.name == "" ? local.apigv1_name : var.name
  stage_name                   = var.stage_name
  root_integration_http_method = var.root_integration_http_method
  root_integration_type        = var.root_integration_type
  root_lambda_arn              = var.root_lambda_arn
  description                  = var.description
  types                        = var.types
  authorization                = var.authorization
  resource_root_path           = var.resource_root_path
  enable_waf_association       = var.enable_waf_association
  web_acl_arn                  = var.web_acl_arn
  resource_paths               = var.resource_paths
  tags = merge(local.tags, {
    Name = var.name == "" ? local.apigv1_name : var.name
  })
}
