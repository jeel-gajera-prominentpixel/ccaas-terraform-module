resource "aws_api_gateway_rest_api" "this" {
  count             = var.create_api ? 1 : 0
  name              = var.api_name
  description       = var.description
  put_rest_api_mode = var.types == "PRIVATE" ? "merge" : "overwrite"

  # Conditional binary media type support
  binary_media_types = var.enable_binary_media_types ? var.binary_media_types : null
  endpoint_configuration {
    types = [var.types]
  }
  tags = var.tags
}

resource "aws_api_gateway_method" "root_method" {
  count              = var.create_root_method ? 1 : 0
  rest_api_id        = aws_api_gateway_rest_api.this[0].id
  resource_id        = aws_api_gateway_rest_api.this[0].root_resource_id
  http_method        = var.resource_root_path
  authorization      = var.authorization
  request_parameters = var.root_resource_request_parameters
}

resource "aws_api_gateway_integration" "root_lambda_integration" {
  count                   = var.create_root_method ? 1 : 0
  rest_api_id             = aws_api_gateway_rest_api.this[0].id
  resource_id             = aws_api_gateway_rest_api.this[0].root_resource_id
  http_method             = aws_api_gateway_method.root_method[0].http_method
  integration_http_method = var.root_integration_http_method
  type                    = var.root_integration_type
  uri                     = var.root_integration_type == "MOCK" ? null : "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${var.root_lambda_arn}/invocations"
  request_parameters      = var.root_integration_request_parameters
}
resource "aws_api_gateway_resource" "api_resources" {
  for_each = var.create_api ? {
    for path, config in var.resource_paths : path => config if path != "/"
  } : tomap({})
  rest_api_id = aws_api_gateway_rest_api.this[0].id
  parent_id   = aws_api_gateway_rest_api.this[0].root_resource_id # var.resource_paths[each.key].parent == "ROOT" ? aws_api_gateway_rest_api.this.root_resource_id : var.resource_paths[each.key].parent
  path_part = (
    each.key == "{proxy+}" ? "{proxy+}" : each.key
  )
}
resource "aws_api_gateway_resource" "child_resource" {
  for_each = var.create_child_resource ? {
    for key, value in var.resource_paths :
    key => value if contains(keys(value), "parent_id")
  } : {}

  rest_api_id = var.rest_api_id
  parent_id   = each.value.parent_id
  path_part   = each.key == "{proxy+}" ? "{proxy+}" : each.key
}

resource "aws_api_gateway_authorizer" "this" {
  count                            = var.create_authorizer ? 1 : 0
  name                             = var.authorizer_name
  rest_api_id                      = var.rest_api_id
  authorizer_uri                   = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${var.authorizer_lambda_arn}/invocations"
  identity_source                  = var.identity_source
  type                             = "REQUEST"
  authorizer_result_ttl_in_seconds = var.authorizer_result_ttl_in_seconds
}


resource "aws_api_gateway_method" "resource_methods" {
  for_each = {
    for path, res in var.resource_paths :
    path => res if lookup(res, "create_method", true)
  }

  rest_api_id        = var.rest_api_id
  resource_id        = each.value.resource_id
  http_method        = each.value.http_method
  authorization      = lookup(each.value, "authorization", var.authorization)
  authorizer_id      = lookup(each.value, "authorizer_id", null)
  request_parameters = each.value.request_parameters
}

resource "aws_api_gateway_integration" "lambda_integration" {
  for_each = {
    for key, method in aws_api_gateway_method.resource_methods :
    key => method if lookup(var.resource_paths[key], "create_method", true)
  }
  rest_api_id             = var.rest_api_id
  resource_id             = var.resource_paths[each.key].resource_id
  http_method             = each.value.http_method
  integration_http_method = var.resource_paths[each.key].integration_http_method
  type                    = var.resource_paths[each.key].type
  uri                     = var.resource_paths[each.key].type == "MOCK" ? null : "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${var.resource_paths[each.key].lambda_arn}/invocations"
  request_parameters      = var.resource_paths[each.key].integration_parameters
}

resource "aws_api_gateway_deployment" "this" {
  count       = var.create_method ? 1 : 0
  rest_api_id = var.rest_api_id
  # stage_name  = var.stage_name

  triggers = {
    redeployment = sha1(jsonencode({
      rest_api_id      = var.rest_api_id
      api_name         = var.api_name
      methods          = aws_api_gateway_method.resource_methods
      # Add timestamp to force redeployment when needed
      # timestamp = timestamp() # Uncomment this line to force redeployment on every apply
    }))
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [
    aws_api_gateway_rest_api.this[0],
    aws_api_gateway_method.root_method[0],
    aws_api_gateway_method.resource_methods[0],
    aws_api_gateway_integration.root_lambda_integration,
    aws_api_gateway_integration.lambda_integration,
    aws_api_gateway_resource.api_resources,
    aws_api_gateway_resource.child_resource[0]
  ]
}

resource "aws_api_gateway_stage" "this" {
  count         = var.create_method ? 1 : 0
  deployment_id = aws_api_gateway_deployment.this[0].id
  rest_api_id   = var.rest_api_id
  stage_name    = var.stage_name
}

resource "aws_api_gateway_method_settings" "this" {
  count       = var.enable_logs ? 1 : 0
  rest_api_id = aws_api_gateway_rest_api.this[0].id
  stage_name  = aws_api_gateway_stage.this[0].stage_name
  method_path = "*/*"
  settings {
    logging_level      = "INFO"
    data_trace_enabled = true
    metrics_enabled    = true
  }
}

resource "aws_wafv2_web_acl_association" "resource_association" {
  count        = var.enable_waf_association ? 1 : 0
  resource_arn = aws_api_gateway_stage.this[0].arn
  web_acl_arn  = var.web_acl_arn
}

resource "aws_api_gateway_account" "this" {
  count               = var.enable_logs ? 1 : 0
  cloudwatch_role_arn = aws_iam_role.api_gateway_logs_role[0].arn
}

resource "aws_iam_role" "api_gateway_logs_role" {
  count              = var.enable_logs ? 1 : 0
  name               = format("api-gateway-%s-logs-role", aws_api_gateway_rest_api.this[0].id)
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy" "api_gateway_logs_policy" {
  count  = var.enable_logs ? 1 : 0
  name   = "default"
  role   = aws_iam_role.api_gateway_logs_role[0].id
  policy = data.aws_iam_policy_document.cloudwatch.json
}

resource "aws_lambda_permission" "api_gateway_permission_root" {
  count         = var.create_root_method ? 1 : 0
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = element(split(":", var.root_lambda_arn), 6)
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.this[0].execution_arn}/*"
}

resource "aws_lambda_permission" "api_gateway_permission" {
  for_each = {
    for k, v in aws_api_gateway_integration.lambda_integration :
    k => v if lookup(var.resource_paths[k], "add_invoke", false)
  }

  statement_id  = "Allow${replace(each.key, "/[\\W]+/", "")}Invoke${substr(md5(var.rest_api_execution_arn), 0, 8)}"
  action        = "lambda:InvokeFunction"
  function_name = regex("function:([^/]+)", each.value.uri)[0]
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.rest_api_execution_arn}/*/*"
}

resource "aws_lambda_permission" "authorizer" {
  count         = var.create_authorizer ? 1 : 0
  statement_id  = "AllowExecutionFromAPIGatewayAuthorizer${substr(md5(var.rest_api_execution_arn), 0, 8)}"
  action        = "lambda:InvokeFunction"
  function_name = regex("function:([^:/]+)", var.authorizer_lambda_arn)[0]
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.rest_api_execution_arn}/*"
}

# Method Response Configuration
resource "aws_api_gateway_method_response" "method_response" {
  for_each = var.method_response_params

  rest_api_id = var.rest_api_id
  resource_id = var.resource_paths[each.key].resource_id
  http_method = var.resource_paths[each.key].http_method
  status_code = each.value.status_code

  response_parameters = each.value.response_parameters
  response_models     = each.value.response_models

  depends_on = [
    aws_api_gateway_method.api_method
  ]
}