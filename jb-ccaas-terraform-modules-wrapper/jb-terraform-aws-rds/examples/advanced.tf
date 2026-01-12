
module "rds_advance" {
  source                          = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds?ref=main"
  prefix_company                  = "jb"
  identifier                      = "test"
  lob                             = "itsd"
  prefix_region                   = "usw2"
  application                     = "recordings"
  env                             = "sandbox"
  engine                          = "postgres"
  engine_version                  = "14"
  family                          = "postgres14"
  major_engine_version            = "14"
  instance_class                  = "db.t4g.large"
  allocated_storage               = 20
  max_allocated_storage           = 100
  db_name                         = "completePostgresql"
  username                        = "complete_postgresql"
  port                            = 5432
  multi_az                        = true
  db_subnet_group_name            = "example-db-subnet-group"
  vpc_security_group_ids          = ["sg-12345678"]
  maintenance_window              = "Mon:00:00-Mon:03:00"
  backup_window                   = "03:00-06:00"
  enabled_cloudwatch_logs_exports = ["postgresql", "upgrade"]
  create_cloudwatch_log_group     = true
  deletion_protection             = false
  create_monitoring_role          = true
  monitoring_interval             = 60
  monitoring_role_name            = "example-monitoring-role-name"
  monitoring_role_use_name_prefix = true
  parameters = [
    {
      name  = "autovacuum"
      value = 1
    }
  ]
  tags = local.tags
}
