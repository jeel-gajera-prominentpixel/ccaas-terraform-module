module "aurora_postgress_advances" {
  source          = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress?ref=main"
  prefix_company  = "jb"
  lob             = "itsd"
  prefix_region   = "usw2"
  application     = "recordings"
  env             = "sandbox"
  name            = "Jb-rds"
  engine          = "aurora-postgresql"
  engine_version  = "14.7"
  master_username = "root"
  storage_type    = "aurora-iopt1"
  instances = {
    1 = {
      instance_class          = "db.r5.2xlarge"
      publicly_accessible     = true
      db_parameter_group_name = "default.aurora-postgresql14"
    }
    2 = {
      identifier     = "static-member-1"
      instance_class = "db.r5.2xlarge"
    }
    instance_class = "db.r6g.large"
    endpoints = {
      static = {
        identifier     = "static-custom-endpt"
        type           = "ANY"
        static_members = ["static-member-1"]
        tags           = { Endpoint = "static-members" }
      }
      excluded = {
        identifier       = "excluded-custom-endpt"
        type             = "READER"
        excluded_members = ["excluded-member-1"]
        tags             = { Endpoint = "excluded-members" }
      }
    }
    vpc_id               = "vpc-0a1b2c3d4e5f6g7h8i"
    db_subnet_group_name = "db-subnet-group-123"
    security_group_rules = {
      vpc_ingress = {
        cidr_blocks = ["10.10.0.0/28"]
      }
      egress_example = {
        cidr_blocks = ["10.33.0.0/28"]
        description = "Egress to corporate printer closet"
      }
    }
    apply_immediately                 = true
    skip_final_snapshot               = true
    autoscaling_enabled               = true
    autoscaling_min_capacity          = 2
    autoscaling_max_capacity          = 5
    create_db_cluster_parameter_group = true
    enabled_cloudwatch_logs_exports   = ["postgresql"]
    create_cloudwatch_log_group       = true
    scaling_configuration = {
      auto_pause               = true
      min_capacity             = 2
      max_capacity             = 16
      seconds_until_auto_pause = 300
      timeout_action           = "ForceApplyCapacityChange"
    }
    monitoring_interval = 60
    tags                = local.tags

  }
}
