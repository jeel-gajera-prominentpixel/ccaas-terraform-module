
module "rds_advance" {
  source                 = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-proxy?ref=main"
  prefix_company         = "jb"
  lob                    = "itsd"
  prefix_region          = "usw2"
  application            = "recordings"
  env                    = "sandbox"
  create_rds_proxy       = true
  name                   = "rds-proxy-example-addvance"
  vpc_subnet_ids         = []
  vpc_security_group_ids = []

  endpoints = {
    read_write = {
      name                   = "read-write-endpoint"
      vpc_subnet_ids         = []
      vpc_security_group_ids = []
      tags                   = local.tags
    },
    read_only = {
      name                   = "read-only-endpoint"
      vpc_subnet_ids         = []
      vpc_security_group_ids = []
      target_role            = "READ_ONLY"
      tags                   = local.tags
    }
  }

  auth = {
    "root" = {
      description = "Cluster generated master user password"
      secret_arn  = ""
    }
  }

  engine_family = "POSTGRESQL"
  debug_logging = false

  # Target Aurora cluster
  target_db_cluster     = false
  db_cluster_identifier = "db cluster identifier"

  # Target RDS instance
  target_db_instance     = false
  db_instance_identifier = "db instance identifier"

  tags = local.tags
}
