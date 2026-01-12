module "rds" {
  source                          = "../../jb-ccaas-terraform-modules/terraform-aws-rds"
  identifier                      = var.identifier == null ? local.identifier_name : var.identifier
  engine                          = var.engine
  engine_version                  = var.engine_version
  family                          = var.family
  major_engine_version            = var.major_engine_version
  instance_class                  = var.instance_class
  allocated_storage               = var.allocated_storage
  max_allocated_storage           = var.max_allocated_storage
  db_name                         = var.db_name
  username                        = var.username
  port                            = var.port
  create_db_subnet_group          = var.create_db_subnet_group
  subnet_ids                      = var.subnet_ids
  multi_az                        = var.multi_az
  db_subnet_group_name            = var.db_subnet_group_name
  vpc_security_group_ids          = var.vpc_security_group_ids
  maintenance_window              = var.maintenance_window
  backup_window                   = var.backup_window
  enabled_cloudwatch_logs_exports = var.enabled_cloudwatch_logs_exports
  create_cloudwatch_log_group     = var.create_cloudwatch_log_group
  deletion_protection             = var.deletion_protection
  create_monitoring_role          = var.create_monitoring_role
  monitoring_interval             = var.monitoring_interval
  monitoring_role_name            = var.monitoring_role_name
  monitoring_role_use_name_prefix = var.monitoring_role_use_name_prefix
  parameters                      = var.parameters
  tags = merge(local.tags, {
    Name = local.identifier_name
  })

}
