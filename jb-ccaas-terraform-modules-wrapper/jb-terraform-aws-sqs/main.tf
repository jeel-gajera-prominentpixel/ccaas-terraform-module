
module "sqs" {
  source                            = "../../jb-ccaas-terraform-modules/terraform-aws-sqs"
  name                              = var.name == "" ? local.sqs_name : var.name
  fifo_queue                        = var.fifo_queue
  create_dlq                        = var.create_dlq
  redrive_policy                    = var.redrive_policy
  sqs_managed_sse_enabled           = var.sqs_managed_sse_enabled
  dlq_redrive_allow_policy          = var.dlq_redrive_allow_policy
  create_queue_policy               = var.create_queue_policy
  queue_policy_statements           = var.queue_policy_statements
  create_dlq_redrive_allow_policy   = var.create_dlq_redrive_allow_policy
  dlq_queue_policy_statements       = var.dlq_queue_policy_statements
  kms_master_key_id                 = var.kms_master_key_id
  kms_data_key_reuse_period_seconds = var.kms_data_key_reuse_period_seconds
  tags = merge(local.tags, {
    Name = var.name == "" ? local.sqs_name : var.name
  })
}
