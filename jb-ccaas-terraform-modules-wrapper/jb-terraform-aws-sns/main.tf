module "sns" {
  source                      = "../../jb-ccaas-terraform-modules/terraform-aws-sns"
  name                        = var.name == "" ? local.sns_name : var.name
  signature_version           = var.signature_version
  use_name_prefix             = var.use_name_prefix
  display_name                = var.display_name
  kms_master_key_id           = var.kms_master_key_id
  tracing_config              = var.tracing_config
  fifo_topic                  = var.fifo_topic
  content_based_deduplication = var.content_based_deduplication
  delivery_policy             = var.delivery_policy
  create_topic_policy         = var.create_topic_policy
  enable_default_topic_policy = var.enable_default_topic_policy
  topic_policy_statements     = var.topic_policy_statements
  subscriptions               = var.subscriptions
  tags = merge(local.tags, {
    Name = var.name == "" ? local.sns_name : var.name
  })
}
