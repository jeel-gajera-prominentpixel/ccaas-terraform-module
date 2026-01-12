module "kinesis_stream" {
  source = "../../jb-ccaas-terraform-modules/terraform-aws-kinesis-stream"

  name                      = var.name == "" ? local.kinesis_stream_name : var.name
  shard_count               = var.shard_count
  retention_period          = var.retention_period
  shard_level_metrics       = var.shard_level_metrics
  enforce_consumer_deletion = var.enforce_consumer_deletion
  encryption_type           = var.encryption_type
  kms_key_id                = var.kms_key_id
  create_policy_read_only   = var.create_policy_read_only
  create_policy_write_only  = var.create_policy_write_only
  create_policy_admin       = var.create_policy_admin

  # # Tags
  tags = merge(local.tags, {
    Name = var.name == "" ? local.kinesis_stream_name : var.name
  })
}
