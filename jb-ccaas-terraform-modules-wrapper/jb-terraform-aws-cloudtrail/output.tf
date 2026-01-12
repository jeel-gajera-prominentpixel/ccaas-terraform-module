output "cloudtrail_id" {
  value       = module.aws_cloudtrail.cloudtrail_id
  description = "The name of the trail"
}

output "cloudtrail_home_region" {
  value       = module.aws_cloudtrail.cloudtrail_home_region
  description = "The region in which the trail was created"
}

output "cloudtrail_arn" {
  value       = module.aws_cloudtrail.cloudtrail_arn
  description = "The Amazon Resource Name of the trail"
}
