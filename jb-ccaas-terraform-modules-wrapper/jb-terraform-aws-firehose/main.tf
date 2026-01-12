module "firehose" {
  source                                           = "../../jb-ccaas-terraform-modules/terraform-aws-kinesis-firehose"
  name                                             = var.name == "" ? local.firehose_name : var.name
  destination                                      = var.destination
  s3_bucket_arn                                    = var.s3_bucket_arn
  s3_prefix                                        = var.s3_prefix
  s3_error_output_prefix                           = var.s3_error_output_prefix
  input_source                                     = var.input_source
  kinesis_source_stream_arn                        = var.kinesis_source_stream_arn
  buffering_size                                   = var.buffering_size
  buffering_interval                               = var.buffering_interval
  enable_s3_encryption                             = var.enable_s3_encryption
  s3_kms_key_arn                                   = var.s3_kms_key_arn
  enable_s3_backup                                 = var.enable_s3_backup
  s3_backup_bucket_arn                             = var.s3_backup_bucket_arn
  append_delimiter_to_record                       = var.append_delimiter_to_record
  redshift_copy_options                            = var.redshift_copy_options
  dynamic_partition_append_delimiter_to_record     = var.dynamic_partition_append_delimiter_to_record
  dynamic_partition_enable_record_deaggregation    = var.dynamic_partition_enable_record_deaggregation
  dynamic_partition_record_deaggregation_type      = var.dynamic_partition_record_deaggregation_type
  dynamic_partition_record_deaggregation_delimiter = var.dynamic_partition_record_deaggregation_delimiter
  sse_kms_key_type                                 = var.sse_kms_key_type
  s3_backup_prefix                                 = var.s3_backup_prefix
  s3_backup_error_output_prefix                    = var.s3_backup_error_output_prefix
  s3_backup_buffering_interval                     = var.s3_backup_buffering_interval
  s3_backup_buffering_size                         = var.s3_backup_buffering_size
  s3_backup_compression                            = var.s3_backup_compression
  s3_backup_use_existing_role                      = var.s3_backup_use_existing_role
  s3_backup_role_arn                               = var.s3_backup_role_arn
  s3_backup_enable_encryption                      = var.s3_backup_enable_encryption
  s3_backup_kms_key_arn                            = var.s3_backup_kms_key_arn
  s3_backup_enable_log                             = var.s3_backup_enable_log
  create_application_role                          = var.create_application_role
  create_application_role_policy                   = var.create_application_role_policy
  application_role_service_principal               = var.application_role_service_principal
  enable_sse                                       = var.enable_sse
  sse_kms_key_arn                                  = var.sse_kms_key_arn
  kinesis_source_kms_arn                           = var.kinesis_source_kms_arn
  kinesis_source_is_encrypted                      = var.kinesis_source_is_encrypted
  tags = merge(local.tags, {
    Name = var.name == "" ? local.firehose_name : var.name
  })

}
