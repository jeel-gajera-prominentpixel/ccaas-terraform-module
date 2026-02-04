module "apigv1" {
  source                       = "../terraform-aws-apigwv1"
  name                     = var.name
  create_rest_api_gateway_method      = true
  create_rest_api_gateway_integration = true
  http_method                         = var.resource_root_path
  integration_http_method             = var.root_integration_http_method
  gateway_integration_type            = var.root_integration_type
  integration_uri                     = var.root_lambda_arn
  rest_api_stage_name                 = var.stage_name
  create_rest_api_deployment           = true
  create_rest_api_gateway_stage        = true
  create_rest_api_gateway_resource     = length(var.api_resources) > 0
  api_resources                        = var.api_resources
  rest_api_endpoint_type       = var.rest_api_endpoint_type
  authorization                = var.authorization
  tags = merge(local.tags, {
    Name = var.name
  })
}