output "key_arn" {
  value       = module.aws_kms.key_arn
  description = "AWS kms key arn"
}

output "key_id" {
  description = "The globally unique identifier for the key"
  value       = module.aws_kms.key_id
}
