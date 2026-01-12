output "parameter_names" {
  description = "List of key names"
  value       = module.ssm_store.names
}

output "parameter_values" {
  description = "List of values"
  value       = module.ssm_store.values
  sensitive   = true
}

output "parameter_map" {
  description = "A map of the names and values created"
  value       = module.ssm_store.map
  sensitive   = true
}

output "parameter_arn_map" {
  description = "A map of the names and ARNs created"
  value       = module.ssm_store.arn_map
}
