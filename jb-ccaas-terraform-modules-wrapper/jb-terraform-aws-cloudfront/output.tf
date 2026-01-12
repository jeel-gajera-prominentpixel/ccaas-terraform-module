output "cloudfront_distribution_id" {
  description = "The identifier for the distribution."
  value       = module.cloudfront.cloudfront_distribution_id
}

output "cloudfront_distribution_arn" {
  description = "The ARN (Amazon Resource Name) for the distribution."
  value       = module.cloudfront.cloudfront_distribution_arn
}

output "cloudfront_distribution_domain_name" {
  description = "The domain name corresponding to the distribution."
  value       = module.cloudfront.cloudfront_distribution_domain_name
}

output "cloudfront_origin_access_identity_ids" {
  description = "The IDS of the origin access identities created"
  value       = module.cloudfront.cloudfront_origin_access_identity_ids
}

output "cloudfront_origin_access_identity_iam_arns" {
  description = "The IAM arns of the origin access identities created"
  value       = module.cloudfront.cloudfront_origin_access_identity_iam_arns
}

output "custom_origin_request_policy_id" {
  value = var.create_request_policy && length(aws_cloudfront_origin_request_policy.custom) > 0 ? aws_cloudfront_origin_request_policy.custom[0].id : null
}
