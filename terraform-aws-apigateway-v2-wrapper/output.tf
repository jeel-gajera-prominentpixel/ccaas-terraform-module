output "apigatewayv2_api_arn" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_api_arn
}

output "default_apigatewayv2_stage_invoke_url" {
  description = "The URL to invoke the API pointing to the stage"
  value       = module.api-gateway.default_apigatewayv2_stage_invoke_url
}

output "apigatewayv2_api_api_endpoint" {
  description = "The API identifier."
  value       = module.api-gateway.apigatewayv2_api_api_endpoint
}

output "apigatewayv2_api_execution_arn" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_api_execution_arn
}

output "apigatewayv2_domain_name_target_domain_name" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_domain_name_target_domain_name
}

output "default_apigatewayv2_stage_domain_name" {
  description = "The ARN of the API"
  value       = module.api-gateway.default_apigatewayv2_stage_domain_name
}

output "apigatewayv2_domain_name_arn" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_domain_name_arn
}

output "apigatewayv2_domain_name_id" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_domain_name_id
}

output "apigatewayv2_api_id" {
  description = "The ARN of the API"
  value       = module.api-gateway.apigatewayv2_api_id
}
