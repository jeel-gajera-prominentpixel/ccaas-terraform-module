output "services" {
  description = "Map of services created and their attributes"
  value       = module.service
}

output "cluster" {
  description = "ECS Cluster attributes"
  value       = module.cluster
}

output "name" {
  description = "The name of the Service"
  value       = var.name

}
output "arn" {
  description = "The ARN of the Service"
  value       = module.service[*].arn
}