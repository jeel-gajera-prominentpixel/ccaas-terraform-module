output "db_subnet_group_name" {
  description = "The db subnet group name"
  value       = module.rds_aurora_postgress.db_subnet_group_name
}


output "cluster_arn" {
  description = "Amazon Resource Name (ARN) of cluster"
  value       = module.rds_aurora_postgress.cluster_arn
}

output "cluster_id" {
  description = "The RDS Cluster Identifier"
  value       = module.rds_aurora_postgress.cluster_id
}

output "cluster_resource_id" {
  description = "The RDS Cluster Resource ID"
  value       = module.rds_aurora_postgress.cluster_resource_id
}

output "cluster_endpoint" {
  description = "Writer endpoint for the cluster"
  value       = module.rds_aurora_postgress.cluster_endpoint
}

output "cluster_reader_endpoint" {
  description = "A read-only endpoint for the cluster, automatically load-balanced across replicas"
  value       = module.rds_aurora_postgress.cluster_reader_endpoint
}

output "cluster_database_name" {
  description = "Name for an automatically created database on cluster creation"
  value       = module.rds_aurora_postgress.cluster_database_name
}

output "security_group_id" {
  description = "The security group ID of the cluster"
  value       = module.rds_aurora_postgress.security_group_id
}

output "postgresql_cluster_master_password" {
  description = "The database master password"
  value       = module.rds_aurora_postgress.cluster_master_password
  sensitive   = true
}


################################################################################
# Cluster
################################################################################
output "cluster_members" {
  description = "List of RDS Instances that are a part of this cluster"
  value       = module.rds_aurora_postgress.cluster_members
}


output "cluster_engine_version_actual" {
  description = "The running version of the cluster database"
  value       = module.rds_aurora_postgress.cluster_engine_version_actual
}


output "cluster_port" {
  description = "The database port"
  value       = module.rds_aurora_postgress.cluster_port
}

output "cluster_master_password" {
  description = "The database master password"
  value       = module.rds_aurora_postgress.cluster_master_password
  sensitive   = true
}

output "cluster_master_username" {
  description = "The database master username"
  value       = module.rds_aurora_postgress.cluster_master_username
  sensitive   = true
}

output "cluster_master_user_secret" {
  description = "The generated database master user secret when `manage_master_user_password` is set to `true`"
  value       = module.rds_aurora_postgress.cluster_master_user_secret
}

output "cluster_hosted_zone_id" {
  description = "The Route53 Hosted Zone ID of the endpoint"
  value       = module.rds_aurora_postgress.cluster_hosted_zone_id
}

################################################################################
# Cluster Instance(s)
################################################################################

output "cluster_instances" {
  description = "A map of cluster instances and their attributes"
  value       = module.rds_aurora_postgress.cluster_instances
}

################################################################################
# Cluster Endpoint(s)
################################################################################

output "additional_cluster_endpoints" {
  description = "A map of additional cluster endpoints and their attributes"
  value       = module.rds_aurora_postgress.additional_cluster_endpoints
}

################################################################################
# Cluster IAM Roles
################################################################################

output "cluster_role_associations" {
  description = "A map of IAM roles associated with the cluster and their attributes"
  value       = module.rds_aurora_postgress.cluster_role_associations
}

################################################################################
# Enhanced Monitoring
################################################################################

output "enhanced_monitoring_iam_role_name" {
  description = "The name of the enhanced monitoring role"
  value       = module.rds_aurora_postgress.enhanced_monitoring_iam_role_name
}

output "enhanced_monitoring_iam_role_arn" {
  description = "The Amazon Resource Name (ARN) specifying the enhanced monitoring role"
  value       = module.rds_aurora_postgress.enhanced_monitoring_iam_role_arn
}

output "enhanced_monitoring_iam_role_unique_id" {
  description = "Stable and unique string identifying the enhanced monitoring role"
  value       = module.rds_aurora_postgress.enhanced_monitoring_iam_role_unique_id
}


################################################################################
# Cluster Parameter Group
################################################################################

output "db_cluster_parameter_group_arn" {
  description = "The ARN of the DB cluster parameter group created"
  value       = module.rds_aurora_postgress.db_cluster_parameter_group_arn
}

output "db_cluster_parameter_group_id" {
  description = "The ID of the DB cluster parameter group created"
  value       = module.rds_aurora_postgress.db_cluster_parameter_group_id
}

################################################################################
# DB Parameter Group
################################################################################

output "db_parameter_group_arn" {
  description = "The ARN of the DB parameter group created"
  value       = module.rds_aurora_postgress.db_parameter_group_arn
}

output "db_parameter_group_id" {
  description = "The ID of the DB parameter group created"
  value       = module.rds_aurora_postgress.db_parameter_group_id
}

################################################################################
# CloudWatch Log Group
################################################################################

output "db_cluster_cloudwatch_log_groups" {
  description = "Map of CloudWatch log groups created and their attributes"
  value       = module.rds_aurora_postgress.db_cluster_cloudwatch_log_groups
}

################################################################################
# Cluster Activity Stream
################################################################################

output "db_cluster_activity_stream_kinesis_stream_name" {
  description = "The name of the Amazon Kinesis data stream to be used for the database activity stream"
  value       = module.rds_aurora_postgress.db_cluster_activity_stream_kinesis_stream_name
}

################################################################################
# Managed Secret Rotation
################################################################################

output "db_cluster_secretsmanager_secret_rotation_enabled" {
  description = "Specifies whether automatic rotation is enabled for the secret"
  value       = module.rds_aurora_postgress.db_cluster_secretsmanager_secret_rotation_enabled
}
