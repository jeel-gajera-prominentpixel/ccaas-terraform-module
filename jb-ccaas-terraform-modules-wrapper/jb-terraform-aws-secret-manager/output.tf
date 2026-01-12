output "secret_arn" {
  description = "The ID of the secret"
  value       = module.secrets-manager.secret_arn
}

output "secret_id" {
  description = "The ID of the secret"
  value       = module.secrets-manager.secret_id
}

output "secret_replica" {
  description = "The ID of the secret"
  value       = module.secrets-manager.secret_replica
}

output "secret_version_id" {
  description = "The ID of the secret"
  value       = module.secrets-manager.secret_version_id
}
