
##----------------------------------------------------------------------------------
## Below resource will Provides a REST API resource.
##----------------------------------------------------------------------------------
resource "aws_api_gateway_rest_api" "rest_api" {
  count = var.enabled && var.create_rest_api ? 1 : 0

  name        = var.name
  description = var.rest_api_description
  tags        = var.tags

  endpoint_configuration {
  types = [var.rest_api_endpoint_type]

}

}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Deployment.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_deployment" "rest_api_deployment" {
  count             = var.enabled && var.create_rest_api && var.create_rest_api_deployment ? 1 : 0
  rest_api_id       = aws_api_gateway_rest_api.rest_api[0].id
  description       = var.api_deployment_description
  variables         = var.rest_variables
  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_rest_api.rest_api[0].body,
      aws_api_gateway_rest_api.rest_api[0].root_resource_id,
      aws_api_gateway_method.rest_api_method[0].id,
      aws_api_gateway_integration.rest_api_integration[0].id,
      aws_api_gateway_integration.rest_api_integration[0].id,
    ]))
  }
  lifecycle {
    create_before_destroy = true
  }
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Resource.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_resource" "api_resources" {
  for_each    = var.enabled && var.create_rest_api && var.create_rest_api_gateway_resource ? var.api_resources : {}
  rest_api_id = aws_api_gateway_rest_api.rest_api[0].id
  parent_id   = aws_api_gateway_rest_api.rest_api[0].root_resource_id
  path_part   = each.value.path_part
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Method.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_method" "api_methods" {
  for_each      = var.enabled && var.create_rest_api && var.create_rest_api_gateway_method ? var.api_resources : {}
  rest_api_id   = aws_api_gateway_rest_api.rest_api[0].id
  resource_id   = aws_api_gateway_resource.api_resources[each.key].id
  http_method   = each.value.http_method
  authorization = var.authorization
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Integration.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_integration" "api_integrations" {
  for_each                = var.enabled && var.create_rest_api && var.create_rest_api_gateway_integration ? var.api_resources : {}
  rest_api_id             = aws_api_gateway_rest_api.rest_api[0].id
  resource_id             = aws_api_gateway_resource.api_resources[each.key].id
  http_method             = aws_api_gateway_method.api_methods[each.key].http_method
  integration_http_method = var.integration_http_method
  type                    = var.gateway_integration_type
  uri                     = each.value.uri
}


##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Method.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_method" "rest_api_method" {
  count         = var.enabled && var.create_rest_api && var.create_rest_api_gateway_method ? 1 : 0
  authorization = var.authorization
  http_method   = var.http_method
  resource_id   = aws_api_gateway_rest_api.rest_api[0].root_resource_id
  rest_api_id   = aws_api_gateway_rest_api.rest_api[0].id
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway integration.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_integration" "rest_api_integration" {
  count                   = var.enabled && var.create_rest_api && var.create_rest_api_gateway_integration ? 1 : 0
  rest_api_id             = aws_api_gateway_rest_api.rest_api[0].id
  resource_id             = aws_api_gateway_method.rest_api_method[0].resource_id
  http_method             = aws_api_gateway_method.rest_api_method[0].http_method
  integration_http_method = var.integration_http_method
  type                    = var.gateway_integration_type
  uri                     = var.integration_uri
}


##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway stage.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_stage" "rest_api_stage" {
  count                 = var.enabled && var.create_rest_api && var.create_rest_api_gateway_stage ? 1 : 0
  description           = var.description_gateway_stage
  deployment_id         = aws_api_gateway_deployment.rest_api_deployment[0].id
  rest_api_id           = aws_api_gateway_rest_api.rest_api[0].id
  stage_name            = var.rest_api_stage_name
  tags = var.tags
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Method Response.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_method_response" "rest_api_method_response" {
  count               = var.enabled && var.create_rest_api_gateway_method_response && var.create_rest_api ? 1 : 0
  rest_api_id         = aws_api_gateway_rest_api.rest_api[0].id
  resource_id         = aws_api_gateway_rest_api.rest_api[0].root_resource_id
  http_method         = aws_api_gateway_method.rest_api_method[0].http_method
  status_code         = var.status_code
}


##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Integration Response.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_integration_response" "rest_api_integration_response" {
  count               = var.enabled && var.create_rest_api_gateway_integration_response && var.create_rest_api ? 1 : 0
  rest_api_id         = aws_api_gateway_rest_api.rest_api[0].id
  resource_id         = aws_api_gateway_method.rest_api_method[0].resource_id
  http_method         = aws_api_gateway_method.rest_api_method[0].http_method
  status_code         = aws_api_gateway_method_response.rest_api_method_response[0].status_code
  depends_on = [
    aws_api_gateway_method.rest_api_method,
    aws_api_gateway_integration.rest_api_integration
  ]
}
