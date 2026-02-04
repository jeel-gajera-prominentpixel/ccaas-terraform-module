##################################
# REST API
##################################
resource "aws_api_gateway_rest_api" "this" {
  name = var.api_name
}

##################################
# PROXY RESOURCE
##################################
resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  parent_id   = aws_api_gateway_rest_api.this.root_resource_id
  path_part   = "{proxy+}"
}

##################################
# ROOT METHODS
##################################
resource "aws_api_gateway_method" "root" {
  for_each = var.root_methods

  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_rest_api.this.root_resource_id
  http_method   = each.key
  authorization = each.value.authorization
}

resource "aws_api_gateway_integration" "root" {
  for_each = var.root_methods

  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_rest_api.this.root_resource_id
  http_method = aws_api_gateway_method.root[each.key].http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = each.value.lambda_arn
}

##################################
# ROOT METHOD RESPONSES (DYNAMIC)
##################################
resource "aws_api_gateway_method_response" "root" {
  for_each = {
    for method, cfg in var.root_methods :
    method => cfg.method_responses
    if contains(keys(cfg), "method_responses")
  }

  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_rest_api.this.root_resource_id
  http_method = each.key
  status_code = each.value != null ? keys(each.value)[0] : null

  response_models     = lookup(each.value[keys(each.value)[0]], "response_models", {})
  response_parameters = lookup(each.value[keys(each.value)[0]], "response_parameters", {})
}

##################################
# PROXY METHODS
##################################
resource "aws_api_gateway_method" "proxy" {
  for_each = var.proxy_methods

  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = each.key
  authorization = each.value.authorization

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_integration" "proxy" {
  for_each = var.proxy_methods

  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.proxy.id
  http_method = aws_api_gateway_method.proxy[each.key].http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = each.value.lambda_arn

  request_parameters = {
    "integration.request.path.proxy" = "method.request.path.proxy"
  }
}

##################################
# PROXY METHOD RESPONSES (DYNAMIC)
##################################
resource "aws_api_gateway_method_response" "proxy" {
  for_each = {
    for method, cfg in var.proxy_methods :
    method => cfg.method_responses
    if contains(keys(cfg), "method_responses")
  }

  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.proxy.id
  http_method = each.key
  status_code = each.value != null ? keys(each.value)[0] : null

  response_models     = lookup(each.value[keys(each.value)[0]], "response_models", {})
  response_parameters = lookup(each.value[keys(each.value)[0]], "response_parameters", {})
}

##################################
# OPTIONS (ROOT + PROXY)
##################################
resource "aws_api_gateway_method" "options_root" {
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_rest_api.this.root_resource_id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "options_root" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_rest_api.this.root_resource_id
  http_method = "OPTIONS"
  type        = "MOCK"
}

resource "aws_api_gateway_method" "options_proxy" {
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "options_proxy" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_resource.proxy.id
  http_method = "OPTIONS"
  type        = "MOCK"
}
