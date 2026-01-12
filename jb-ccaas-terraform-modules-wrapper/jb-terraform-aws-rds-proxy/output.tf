#####################################################################################
################# OUTPUT FOR RDS CLUSTER PROXY & INSTANCE PROXY #####################
#####################################################################################


# RDS Proxy
output "proxy_id" {
  description = "The ID for the proxy"
  value       = try(module.rds_proxy[0].proxy_id, null)
}

output "proxy_arn" {
  description = "The Amazon Resource Name (ARN) for the proxy"
  value       = try(module.rds_proxy[0].proxy_arn, null)
}

output "proxy_endpoint" {
  description = "The endpoint that you can use to connect to the proxy"
  value       = try(module.rds_proxy[0].proxy_endpoint, null)
}

# Proxy Default Target Group
output "proxy_default_target_group_id" {
  description = "The ID for the default target group"
  value       = try(module.rds_proxy[0].proxy_default_target_group_id, null)
}

output "proxy_default_target_group_arn" {
  description = "The Amazon Resource Name (ARN) for the default target group"
  value       = try(module.rds_proxy[0].proxy_default_target_group_arn, null)
}

output "proxy_default_target_group_name" {
  description = "The name of the default target group"
  value       = try(module.rds_proxy[0].proxy_default_target_group_name, null)
}

# Proxy Target
output "proxy_target_endpoint" {
  description = "Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type"
  value       = try(module.rds_proxy[0].proxy_target_endpoint, null)
}

output "proxy_target_id" {
  description = "Identifier of `db_proxy_name`, `target_group_name`, target type (e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`), and resource identifier separated by forward slashes (/)"
  value       = try(module.rds_proxy[0].proxy_target_id, null)
}

output "proxy_target_port" {
  description = "Port for the target RDS DB Instance or Aurora DB Cluster"
  value       = try(module.rds_proxy[0].proxy_target_port, null)
}

output "proxy_target_rds_resource_id" {
  description = "Identifier representing the DB Instance or DB Cluster target"
  value       = try(module.rds_proxy[0].proxy_target_rds_resource_id, null)
}

output "proxy_target_target_arn" {
  description = "Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API"
  value       = try(module.rds_proxy[0].proxy_target_target_arn, null)
}

output "proxy_target_tracked_cluster_id" {
  description = "DB Cluster identifier for the DB Instance target. Not returned unless manually importing an RDS_INSTANCE target that is part of a DB Cluster"
  value       = try(module.rds_proxy[0].proxy_target_tracked_cluster_id, null)
}

output "proxy_target_type" {
  description = "Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`"
  value       = try(module.rds_proxy[0].proxy_target_type, null)
}

# DB proxy endpoints
output "db_proxy_endpoints" {
  description = "Array containing the full resource object and attributes for all DB proxy endpoints created"
  value       = try(module.rds_proxy[0].db_proxy_endpoints, null)
}

# CloudWatch logs
output "log_group_arn" {
  description = "The Amazon Resource Name (ARN) of the CloudWatch log group"
  value       = try(module.rds_proxy[0].log_group_arn, null)
}

# IAM role
output "iam_role_arn" {
  description = "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager."
  value       = try(module.rds_proxy[0].iam_role_arn, null)
}

output "iam_role_name" {
  description = "IAM role name"
  value       = try(module.rds_proxy[0].iam_role_name, null)
}

output "iam_role_unique_id" {
  description = "Stable and unique string identifying the IAM role"
  value       = try(module.rds_proxy[0].iam_role_unique_id, null)
}
