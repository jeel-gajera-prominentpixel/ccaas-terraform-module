# API Gateway REST API Definition
resource "aws_api_gateway_rest_api" "this" {
  name        = var.api_name
  description = var.description
  put_rest_api_mode = var.types == "PRIVATE" ? "merge" : "overwrite"
  endpoint_configuration {
    types = [var.types]
  }
  tags = var.tags
}

# Root Resource Method
resource "aws_api_gateway_method" "root_method" {
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_rest_api.this.root_resource_id
  http_method   = var.resource_root_path
  authorization = var.authorization
}

# Integration with Lambda for Root Method
resource "aws_api_gateway_integration" "root_lambda_integration" {
  rest_api_id            = aws_api_gateway_rest_api.this.id
  resource_id            = aws_api_gateway_rest_api.this.root_resource_id
  http_method            = aws_api_gateway_method.root_method.http_method
  integration_http_method = var.root_integration_http_method # Typically POST for Lambda integrations
  type                   = var.root_integration_type
  uri                    = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${var.root_lambda_arn}/invocations"
}

# Lambda Permission for API Gateway to Invoke Root Function
resource "aws_lambda_permission" "root_invoke_permission" {
  action        = "lambda:InvokeFunction"
  function_name = element(split(":", var.root_lambda_arn), 6)
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.this.execution_arn}/${aws_api_gateway_method.root_method.http_method}/"
}

# Root Method Response
resource "aws_api_gateway_method_response" "root_method_response" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_rest_api.this.root_resource_id
  http_method = aws_api_gateway_method.root_method.http_method
  status_code = var.root_status_code

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin" = true
  }
}

# Root Integration Response
resource "aws_api_gateway_integration_response" "root_integration_response" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = aws_api_gateway_rest_api.this.root_resource_id
  http_method = aws_api_gateway_method.root_method.http_method
  status_code = aws_api_gateway_method_response.root_method_response.status_code

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin" = "'*'"
  }

  depends_on = [
    aws_api_gateway_integration.root_lambda_integration
  ]
}



# API Gateway Resource
resource "aws_api_gateway_resource" "api_resources" {
  for_each    = { for path, config in var.resource_paths : path => config if path != "/" }
  rest_api_id = aws_api_gateway_rest_api.this.id
  parent_id   = aws_api_gateway_rest_api.this.root_resource_id
  path_part   = each.key
}

# API Gateway Method for Each Resource (using specified HTTP method)
resource "aws_api_gateway_method" "resource_methods" {
  for_each      = aws_api_gateway_resource.api_resources
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = each.value.id
  http_method   = var.resource_paths[each.key].http_method
  authorization = var.authorization
}

resource "aws_api_gateway_integration" "lambda_integration" {
  for_each               = aws_api_gateway_method.resource_methods
  rest_api_id            = aws_api_gateway_rest_api.this.id
  resource_id            = each.value.resource_id
  http_method            = each.value.http_method
  integration_http_method = var.resource_paths[each.key].integration_http_method  # For Lambda, this typically stays as POST
  type                   = var.resource_paths[each.key].type
  uri                    = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${var.resource_paths[each.key].lambda_arn}/invocations"
}


# Lambda Permission for API Gateway to Invoke Function
resource "aws_lambda_permission" "api_gateway_invoke_permission" {
  for_each = aws_api_gateway_integration.lambda_integration
  action        = "lambda:InvokeFunction"
  function_name = regex("function:([^/]+)", each.value.uri)[0]
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.this.execution_arn}/*/*"
}

# Method Response for Each Resource
resource "aws_api_gateway_method_response" "method_response" {
  for_each    = aws_api_gateway_resource.api_resources
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = each.value.id
  http_method = var.resource_paths[each.key].http_method
  status_code = var.resource_paths[each.key].status_code

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin" = true
  }
}

# Integration Response for Each Resource
resource "aws_api_gateway_integration_response" "integration_response" {
  for_each    = aws_api_gateway_resource.api_resources
  rest_api_id = aws_api_gateway_rest_api.this.id
  resource_id = each.value.id
  http_method = var.resource_paths[each.key].http_method
  status_code = aws_api_gateway_method_response.method_response[each.key].status_code

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin" = "'*'"
  }

  depends_on = [
    aws_api_gateway_integration.lambda_integration
  ]
}

resource "aws_api_gateway_deployment" "this" {
  rest_api_id = aws_api_gateway_rest_api.this.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.this.body))
  }

  lifecycle {
    create_before_destroy = true
  }
  depends_on = [
    aws_api_gateway_integration.root_lambda_integration,  # Ensure integration is created first
    aws_api_gateway_method.root_method,            # Ensure method is created first
    aws_api_gateway_integration_response.root_integration_response,
    aws_api_gateway_integration.lambda_integration,
    aws_api_gateway_integration_response.integration_response
  ]
}

resource "aws_api_gateway_stage" "this" {
  deployment_id = aws_api_gateway_deployment.this.id
  rest_api_id   = aws_api_gateway_rest_api.this.id
  stage_name    = var.stage_name
}


resource "aws_wafv2_web_acl_association" "resource_association" {
  count         = var.enable_waf_association ? 1 : 0
  resource_arn  = aws_api_gateway_stage.this.arn
  web_acl_arn   = var.web_acl_arn
}