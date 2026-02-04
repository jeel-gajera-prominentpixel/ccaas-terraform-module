output "apigatewayv1_api_id" {
  description = "The API identifier"
  value       = try(aws_api_gateway_rest_api.this[0].id, "")
}

output "apigatewayv1_api_arn" {
  description = "The ARN of the API"
  value       = try(aws_api_gateway_rest_api.this[0].arn, "")
}

output "apigatewayv1_api_execution_arn" {
  description = "The ARN prefix to be used in an aws_lambda_permission's source_arn attribute or in an aws_iam_policy to authorize access to the @connections API."
  value       = try(aws_api_gateway_rest_api.this[0].execution_arn, "")
}

output "apigatewayv1_key_source" {
  description = "Source of the API key for requests."
  value       = try(aws_api_gateway_rest_api.this[0].api_key_source, "")
}

output "apigatewayv1_binary_media_types" {
  description = "List of binary media types supported by the REST API."
  value       = try(aws_api_gateway_rest_api.this[0].binary_media_types, "")
}

output "apigatewayv1_description" {
  description = "Description of the REST API."
  value       = try(aws_api_gateway_rest_api.this[0].description, "")
}

output "apigatewayv1_endpoint_configuration" {
  description = "The endpoint configuration of this RestApi showing the endpoint types of the API."
  value       = try(aws_api_gateway_rest_api.this[0].endpoint_configuration, "")
}

output "apigatewayv1_minimum_compression_size" {
  description = "Minimum response size to compress for the REST API."
  value       = try(aws_api_gateway_rest_api.this[0].minimum_compression_size, "")
}

output "apigatewayv1_policy" {
  description = "JSON formatted policy document that controls access to the API Gateway."
  value       = try(aws_api_gateway_rest_api.this[0].policy, "")
}

output "apigatewayv1_root_resource_id" {
  description = "Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'."
  value       = try(aws_api_gateway_rest_api.this[0].root_resource_id, "")
}

output "apigatewayv1_deployment_id" {
  description = "ID of the deployment"
  value       = try(aws_api_gateway_deployment.this[0].id, "")
}

# Fixed: These attributes belong to the stage, not deployment
output "apigatewayv1_stage_invoke_url" {
  description = "URL to invoke the API pointing to the stage"
  value       = try(aws_api_gateway_stage.this[0].invoke_url, "")
}

output "apigatewayv1_stage_execution_arn" {
  description = "Execution ARN to be used in lambda_permission's source_arn when allowing API Gateway to invoke a Lambda function"
  value       = try(aws_api_gateway_stage.this[0].execution_arn, "")
}

output "apigatewayv1_stage_arn" {
  description = "Stage ARN"
  value       = try(aws_api_gateway_stage.this[0].arn, "")
}

output "apigatewayv1_stage_id" {
  description = "Stage ID"
  value       = try(aws_api_gateway_stage.this[0].id, "")
}

output "apigatewayv1_stage_web_acl_arn" {
  description = "ARN of the WebAcl associated with the Stage."
  value       = try(aws_api_gateway_stage.this[0].web_acl_arn, "")
}

output "authorizer_id" {
  description = "id of the authorizer."
  value       = try(aws_api_gateway_authorizer.this[0].id, "")
}