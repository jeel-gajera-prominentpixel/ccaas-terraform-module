module "aurora_postgres_basic" {
  source                 = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress?ref=main"
  prefix_company         = "jb"
  lob                    = "itsd"
  prefix_region          = "usw2"
  application            = "recordings"
  env                    = "sandbox"
  version                = "3.0.0"
  name                   = "example-db"
  engine                 = "aurora-postgresql"
  engine_version         = "10.7"
  instance_class         = "db.t2.small"
  allocated_storage      = 10
  master_username        = "admin"
  master_password        = "password"
  subnets                = ["subnet-12345678"]
  vpc_security_group_ids = ["sg-12345678"]
}
