
##----------------------------------------------------------------------------------
## Below resource will Provides a REST API resource.
##----------------------------------------------------------------------------------
resource "aws_api_gateway_rest_api" "rest_api" {
  count = var.enabled && var.create_rest_api ? 1 : 0

  name        = var.name
  description = var.rest_api_description
  tags        = var.tags

  endpoint_configuration {
    types            = [var.rest_api_endpoint_type]
    vpc_endpoint_ids = var.rest_api_endpoint_type == "PRIVATE" ? (var.create_vpc_endpoint ? [aws_vpc_endpoint.rest_api_private[0].id] : var.vpc_endpoint_id) : null
  }
}

##--------------------------------------------------------------------------------
# Resource Policy for [aws_api_gateway_rest_api.rest_api]
##--------------------------------------------------------------------------------
resource "aws_api_gateway_rest_api_policy" "rest_api_resource_policy" {
  count = var.enabled && var.create_rest_api && var.rest_api_endpoint_type == "PRIVATE" ? 1 : 0

  rest_api_id = aws_api_gateway_rest_api.rest_api[0].id
  policy      = var.rest_api_resource_policy != "" ? var.rest_api_resource_policy : <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Deny",
            "Principal": "*",
            "Action": "execute-api:Invoke",
            "Resource": "${aws_api_gateway_rest_api.rest_api[0].execution_arn}/*",
            "Condition": {
                "StringNotEquals": {
                    "aws:sourceVpce": "${aws_vpc_endpoint.rest_api_private[0].id}"
                }
            }
        },
        {
            "Effect": "Allow",
            "Principal": "*",
            "Action": "execute-api:Invoke",
            "Resource": "${aws_api_gateway_rest_api.rest_api[0].execution_arn}/*"
        }
    ]
}  
  EOF
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Deployment.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_deployment" "rest_api_deployment" {
  count             = var.enabled && var.create_rest_api && var.create_rest_api_deployment ? 1 : 0
  rest_api_id       = aws_api_gateway_rest_api.rest_api[0].id
  description       = var.api_deployment_description
  stage_description = var.stage_description
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
  connection_type         = var.connection_rest_api_type
  connection_id           = var.connection_id
  credentials             = var.credentials
  request_templates       = var.request_templates
  request_parameters      = var.request_parameters
  cache_namespace         = var.cache_namespace
  content_handling        = var.content_handling
  cache_key_parameters    = var.cache_key_parameters
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
  connection_type         = var.connection_rest_api_type
  connection_id           = var.connection_id
  credentials             = var.credentials
  request_templates       = var.request_templates
  request_parameters      = var.request_parameters
  cache_namespace         = var.cache_namespace
  content_handling        = var.content_handling
  cache_key_parameters    = var.cache_key_parameters
  type                    = var.gateway_integration_type
  timeout_milliseconds    = var.timeout_milliseconds
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
  cache_cluster_enabled = var.cache_cluster_enabled
  cache_cluster_size    = var.cache_cluster_size
  client_certificate_id = var.client_certificate_id
  documentation_version = var.documentation_version
  variables             = var.stage_variables
  xray_tracing_enabled  = var.xray_tracing_enabled

  dynamic "canary_settings" {
    for_each = var.canary_settings
    content {
      percent_traffic          = canary_settings.percent_traffic.value
      stage_variable_overrides = canary_settings.stage_variable_overrides.value
      use_stage_cache          = canary_settings.use_stage_cache.value
    }
  }

  dynamic "access_log_settings" {
    for_each = var.enable_access_logs == true ? [1] : []

    content {
      destination_arn = aws_cloudwatch_log_group.rest_api_log[0].arn
      format          = replace(var.log_format, "\n", "")
    }
  }

  lifecycle {
    create_before_destroy = true
  }

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
  response_models     = var.response_models
  response_parameters = var.response_parameters
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
  content_handling    = var.content_handling
  response_parameters = var.integration_response_parameters
  depends_on = [
    aws_api_gateway_method.rest_api_method,
    aws_api_gateway_integration.rest_api_integration
  ]
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Gateway Authorizer.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_authorizer" "rest_api_authorizer" {
  count                            = var.enabled && var.create_rest_api_gateway_authorizer && var.create_rest_api ? 1 : 0
  name                             = var.gateway_authorizer
  rest_api_id                      = aws_api_gateway_rest_api.rest_api[0].id
  authorizer_uri                   = var.integration_uri
  authorizer_credentials           = var.authorizer_iam_role != "" ? var.authorizer_iam_role : aws_iam_role.rest_api_iam_role[0].arn
  identity_source                  = var.identity_source
  type                             = var.type
  authorizer_result_ttl_in_seconds = var.authorizer_result_ttl_in_seconds
  provider_arns                    = var.provider_arns
}

##----------------------------------------------------------------------------------
## Below resource will Manages an Amazon REST API Base Path Mapping.
##----------------------------------------------------------------------------------

resource "aws_api_gateway_base_path_mapping" "rest_api_base_path" {
  count       = var.enabled && var.create_rest_api_gateway_authorizer && var.create_rest_api ? 1 : 0
  api_id      = aws_api_gateway_rest_api.rest_api[0].id
  domain_name = var.domain_name
  base_path   = var.rest_api_base_path
  stage_name  = var.rest_api_stage_name
}