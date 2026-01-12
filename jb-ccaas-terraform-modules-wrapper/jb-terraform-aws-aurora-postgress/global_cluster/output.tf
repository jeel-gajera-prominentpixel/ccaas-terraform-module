################################################################################
# Global Cluster
################################################################################

output "global_cluster_id" {
  description = "The ID of global cluster"
  value       = module.global_rds_cluster.global_cluster_id
}

output "global_cluster_arn" {
  description = "The ARN of global cluster"
  value       = module.global_rds_cluster.global_cluster_id
}

output "global_cluster_engine" {
  description = "The engine of global cluster"
  value       = module.global_rds_cluster.global_cluster_engine
}

output "global_cluster_engine_version" {
  description = "The engine version of global cluster"
  value       = module.global_rds_cluster.global_cluster_engine_version
}

output "global_cluster_db_name" {
  description = "The database name of global cluster"
  value       = module.global_rds_cluster.global_cluster_db_name
}

output "global_cluster_endpoint" {
  description = "The endpoint of global cluster"
  value       = module.global_rds_cluster.global_cluster_endpoint
}

output "global_cluster_members" {
  description = "The members of global cluster"
  value       = module.global_rds_cluster.global_cluster_members
}

output "global_cluster_resource_id" {
  description = "The resource id of global cluster"
  value       = module.global_rds_cluster.global_cluster_resource_id
}

output "global_cluster_tags_all" {
  description = "The all tags of global cluster"
  value       = module.global_rds_cluster.global_cluster_tags_all
}
