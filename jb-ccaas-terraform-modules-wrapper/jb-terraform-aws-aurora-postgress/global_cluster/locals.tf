locals {
  global_cluster_identifier_name = format("%s-rds-%s-%s-etl-global-%s", var.prefix_company, var.lob, var.application, var.env)
}
