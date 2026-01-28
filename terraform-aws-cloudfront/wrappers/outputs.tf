output "wrapper" {
  description = "Map of outputs of a wrapper."
  value       = module.wrapper
  # sensitive = false # No sensitive module output found
}

output "cloudfront_distribution_id" {
  description = "The identifier for the distribution."
  value       = module.wrapper.cloudfront_distribution_id
}

output "cloudfront_distribution_arn" {
  description = "The ARN (Amazon Resource Name) for the distribution."
  value       = module.wrapper.cloudfront_distribution_arn
}