module "rds_proxy" {
  count  = var.create_rds_proxy ? 1 : 0
  source = "../../jb-ccaas-terraform-modules/terraform-aws-rds-proxy"

  name                   = var.name == null ? local.rds_proxy_name : var.name
  iam_role_name          = var.name == null ? local.rds_proxy_name : var.name
  vpc_subnet_ids         = var.vpc_subnet_ids
  vpc_security_group_ids = var.vpc_security_group_ids

  endpoints = var.endpoints

  auth = var.auth

  engine_family = var.engine_family
  debug_logging = var.debug_logging

  # Target Aurora cluster
  target_db_cluster     = var.target_db_cluster
  db_cluster_identifier = var.db_cluster_identifier

  # Target RDS instance
  target_db_instance     = var.target_db_instance
  db_instance_identifier = var.db_instance_identifier

  idle_client_timeout            = var.idle_client_timeout
  require_tls                    = var.require_tls
  role_arn                       = var.role_arn
  log_group_retention_in_days    = var.log_group_retention_in_days
  log_group_kms_key_id           = var.log_group_kms_key_id
  log_group_tags                 = var.log_group_tags
  manage_log_group               = var.manage_log_group
  create_iam_role                = var.create_iam_role
  use_role_name_prefix           = var.use_role_name_prefix
  iam_role_description           = var.iam_role_description
  iam_role_path                  = var.iam_role_path
  iam_role_force_detach_policies = var.iam_role_force_detach_policies
  iam_role_max_session_duration  = var.iam_role_max_session_duration
  iam_role_permissions_boundary  = var.iam_role_permissions_boundary
  create_iam_policy              = var.create_iam_policy
  use_policy_name_prefix         = var.use_policy_name_prefix

  tags = merge(local.tags, {
    Name = local.rds_proxy_name
  })
}
