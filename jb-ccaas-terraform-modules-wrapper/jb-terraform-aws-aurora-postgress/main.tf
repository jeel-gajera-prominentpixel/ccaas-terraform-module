module "rds_aurora_postgress" {
  source                                = "../../jb-ccaas-terraform-modules/terraform-aws-rds-aurora"
  name                                  = var.name == "" ? local.rds_name : var.name
  engine                                = var.engine
  engine_mode                           = var.engine_mode
  storage_encrypted                     = var.storage_encrypted
  manage_master_user_password           = var.manage_master_user_password
  master_password                       = var.master_password
  master_user_secret_kms_key_id         = var.master_user_secret_kms_key_id
  engine_version                        = var.engine_version
  master_username                       = var.master_username
  storage_type                          = var.storage_type
  create_db_subnet_group                = var.create_db_subnet_group
  preferred_maintenance_window          = var.preferred_maintenance_window
  security_group_use_name_prefix        = var.security_group_use_name_prefix
  instances                             = var.instances
  instance_class                        = var.instance_class
  iam_role_name                         = var.iam_role_name
  iam_role_use_name_prefix              = var.iam_role_use_name_prefix
  iam_role_description                  = var.iam_role_description
  iam_role_path                         = var.iam_role_path
  iam_role_max_session_duration         = var.iam_role_max_session_duration
  endpoints                             = var.endpoints
  vpc_id                                = var.vpc_id
  vpc_security_group_ids                = var.vpc_security_group_ids
  db_subnet_group_name                  = var.db_subnet_group_name
  subnets                               = var.subnets
  security_group_rules                  = var.security_group_rules
  apply_immediately                     = var.apply_immediately
  skip_final_snapshot                   = var.skip_final_snapshot
  autoscaling_enabled                   = var.autoscaling_enabled
  autoscaling_min_capacity              = var.autoscaling_min_capacity
  autoscaling_max_capacity              = var.autoscaling_max_capacity
  create_db_cluster_parameter_group     = var.create_db_cluster_parameter_group
  enabled_cloudwatch_logs_exports       = var.enabled_cloudwatch_logs_exports
  create_cloudwatch_log_group           = var.create_cloudwatch_log_group
  deletion_protection                   = var.deletion_protection
  port                                  = var.port
  serverlessv2_scaling_configuration    = var.serverlessv2_scaling_configuration
  scaling_configuration                 = var.scaling_configuration
  create_monitoring_role                = var.create_monitoring_role
  monitoring_interval                   = var.monitoring_interval
  allocated_storage                     = var.allocated_storage
  create_security_group                 = var.create_security_group
  security_group_name                   = var.security_group_name
  security_group_description            = var.security_group_description
  security_group_tags                   = var.security_group_tags
  database_name                         = var.database_name
  enable_http_endpoint                  = var.enable_http_endpoint
  replication_source_identifier         = var.replication_source_identifier
  is_primary_cluster                    = var.is_primary_cluster
  global_cluster_identifier             = var.global_cluster_identifier
  source_region                         = var.source_region
  kms_key_id                            = var.kms_key_id
  performance_insights_enabled          = var.performance_insights_enabled
  performance_insights_kms_key_id       = var.performance_insights_kms_key_id
  performance_insights_retention_period = var.performance_insights_retention_period
  auto_minor_version_upgrade            = var.auto_minor_version_upgrade
  preferred_backup_window               = var.preferred_backup_window
  backup_retention_period               = var.backup_retention_period
  allow_major_version_upgrade           = var.allow_major_version_upgrade

  tags = merge(local.tags, {
    Name = var.name == "" ? local.rds_name : var.name
  })
}
