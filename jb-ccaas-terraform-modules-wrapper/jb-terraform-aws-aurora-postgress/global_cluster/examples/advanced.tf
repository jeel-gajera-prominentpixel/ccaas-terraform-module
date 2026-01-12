module "rds_advance" {
  source                           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress/global_cluster?ref=main"
  prefix_company                   = "jb"
  identifier                       = "test"
  lob                              = "itsd"
  prefix_region                    = "usw2"
  application                      = "recordings"
  env                              = "sandbox"
  create                           = false
  create_global_cluster            = true
  global_cluster_engine            = "postgres"
  global_cluster_version           = "14"
  global_cluster_db_name           = "completePostgresql"
  global_cluster_storage_encrypted = true
}
