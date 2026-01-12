module "dynamodb-table" {
  source                                = "../../jb-ccaas-terraform-modules/terraform-aws-dynamodb-table"
  name                                  = var.name == "" ? local.dynamodb_table_name : var.name
  hash_key                              = var.hash_key
  range_key                             = var.range_key
  table_class                           = var.table_class
  deletion_protection_enabled           = var.deletion_protection_enabled
  attributes                            = var.attributes
  import_table                          = var.import_table
  stream_enabled                        = var.stream_enabled
  stream_view_type                      = var.stream_view_type
  server_side_encryption_enabled        = var.server_side_encryption_enabled
  server_side_encryption_kms_key_arn    = var.server_side_encryption_kms_key_arn
  global_secondary_indexes              = var.global_secondary_indexes
  replica_regions                       = var.replica_regions
  ttl_enabled                           = var.ttl_enabled
  ttl_attribute_name                    = var.ttl_attribute_name
  point_in_time_recovery_enabled        = var.point_in_time_recovery_enabled
  point_in_time_recovery_period_in_days = var.point_in_time_recovery_period_in_days
  tags = merge(local.tags, {
    Name = var.name == "" ? local.dynamodb_table_name : var.name
  })

}
