module "kinesis_stream_advance" {
  source                    = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kinesis-stream?ref=main"
  prefix_company            = "jb"
  lob                       = "itsd"
  prefix_region             = "usw2"
  application               = "recordings"
  env                       = "sandbox"
  shard_count               = 1
  retention_period          = 24
  shard_level_metrics       = []
  enforce_consumer_deletion = false
  encryption_type           = "NONE"
  kms_key_id                = "test-kms-key-id"
  create_policy_read_only   = true
  create_policy_write_only  = true
  create_policy_admin       = true
  tags                      = local.tags
}
