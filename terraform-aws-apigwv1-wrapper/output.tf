output "apigatewayv1_api_id" {
  value       = module.apigv1.apigatewayv1_api_id
  description = "The API identifier"
}

output "apigatewayv1_api_arn" {
  value       = module.apigv1.apigatewayv1_api_arn
  description = "The ARN of the API"
}

output "apigatewayv1_api_execution_arn" {
  value       = module.apigv1.apigatewayv1_api_execution_arn
  description = "The ARN prefix to be used in an aws_lambda_permission's source_arn attribute or in an aws_iam_policy to authorize access to the @connections API."
}

output "apigatewayv1_key_source" {
  value       = module.apigv1.apigatewayv1_key_source
  description = "Source of the API key for requests."
}

output "apigatewayv1_binary_media_types" {
  value       = module.apigv1.apigatewayv1_binary_media_types
  description = "List of binary media types supported by the REST API."
}

output "apigatewayv1_description" {
  value       = module.apigv1.apigatewayv1_description
  description = "Description of the REST API."
}

output "apigatewayv1_endpoint_configuration" {
  value       = module.apigv1.apigatewayv1_endpoint_configuration
  description = "The endpoint configuration of this RestApi showing the endpoint types of the API."
}

output "apigatewayv1_minimum_compression_size" {
  value       = module.apigv1.apigatewayv1_minimum_compression_size
  description = "Minimum response size to compress for the REST API."
}

output "apigatewayv1_policy" {
  value       = module.apigv1.apigatewayv1_policy
  description = "JSON formatted policy document that controls access to the API Gateway."
}

output "apigatewayv1_root_resource_id" {
  value       = module.apigv1.apigatewayv1_root_resource_id
  description = "Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'."
}

output "apigatewayv1_stage_arn" {
  value       = module.apigv1.apigatewayv1_stage_arn
  description = "Stage ARN"
}

output "apigatewayv1_stage_id" {
  value       = module.apigv1.apigatewayv1_stage_id
  description = "Stage ID"
}

output "apigatewayv1_stage_invoke_url" {
  value       = module.apigv1.apigatewayv1_stage_invoke_url
  description = "URL to invoke the API pointing to the stage, e.g., https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod"
}

output "apigatewayv1_stage_execution_arn" {
  value       = module.apigv1.apigatewayv1_stage_execution_arn
  description = "Execution ARN to be used in lambda_permission's source_arn when allowing API Gateway to invoke a Lambda function, e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod"
}

output "apigatewayv1_stage_web_acl_arn" {
  value       = module.apigv1.apigatewayv1_stage_web_acl_arn
  description = "ARN of the WebAcl associated with the Stage."
}
