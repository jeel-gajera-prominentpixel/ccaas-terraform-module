module "apigv1" {
  source = "../terraform-aws-apigwv1" # real module path

  enabled         = true
  create_rest_api = true

  name                 = var.name
  rest_api_description = var.description
  rest_api_endpoint_type = var.rest_api_endpoint_type

  create_rest_api_gateway_resource    = var.create_rest_api_gateway_resource
  create_rest_api_gateway_method      = var.create_rest_api_gateway_method
  create_rest_api_gateway_integration = var.create_rest_api_gateway_integration

  http_method   = var.http_method
  authorization = var.authorization

  gateway_integration_type = var.gateway_integration_type
  integration_http_method  = var.integration_http_method
  integration_uri          = var.integration_uri

  api_resources = var.api_resources

  create_rest_api_deployment    = var.create_rest_api_deployment
  create_rest_api_gateway_stage = var.create_rest_api_gateway_stage
  rest_api_stage_name           = var.rest_api_stage_name

  xray_tracing_enabled = var.xray_tracing_enabled
  tags                 = var.tags
}
