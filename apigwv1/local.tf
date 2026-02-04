locals {
  effective_rest_api_id = var.create_api ? aws_api_gateway_rest_api.this[0].id : var.rest_api_id
}
