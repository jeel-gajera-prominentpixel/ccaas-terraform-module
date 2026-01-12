module "global_rds_cluster" {
  source                    = "../../jb-ccaas-terraform-modules/terraform-aws-rds-aurora"
  create                    = false
  create_global_cluster     = var.create_global_cluster
  global_cluster_identifier = var.global_cluster_identifier == null ? local.global_cluster_identifier_name : var.global_cluster_identifier
  global_cluster_engine     = var.global_cluster_engine
  global_cluster_version    = var.global_cluster_version
  global_cluster_db_name    = var.global_cluster_db_name
  # global_cluster_storage_encrypted        = var.global_cluster_storage_encrypted
  global_source_db_cluster_identifier     = var.source_db_cluster_identifier
  global_cluster_force_destroy            = var.force_destroy
  global_cluster_deletion_protection      = var.deletion_protection
  global_cluster_engine_lifecycle_support = var.engine_lifecycle_support
}
