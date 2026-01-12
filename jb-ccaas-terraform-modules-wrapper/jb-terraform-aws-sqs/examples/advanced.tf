module "sqs_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sqs?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb_sqs"
  fifo_queue     = true
  create_dlq     = true
  redrive_policy = {
    maxReceiveCount = 10
  }
  sqs_managed_sse_enabled = false
  dlq_redrive_allow_policy = {
    sourceQueueArns = "arn:aws:sqs:us-east-1:123456789012:my-queue"

  }
  create_queue_policy = true
  queue_policy_statements = {
    account = {
      sid = "AccountReadWrite"
      actions = [
        "sqs:SendMessage",
        "sqs:ReceiveMessage",
      ]
      principals = [
        {
          type        = "AWS"
          identifiers = ["arn:aws:iam::123456789012:root"]
        }
      ]
    }
  }
  create_dlq_redrive_allow_policy = false
  dlq_queue_policy_statements = {
    account = {
      sid = "AccountReadWrite"
      actions = [
        "sqs:SendMessage",
        "sqs:ReceiveMessage",
      ]
      principals = [
        {
          type        = "AWS"
          identifiers = ["arn:aws:iam::123456789012:root"]
        }
      ]
    }
  }
  kms_master_key_id                 = "0d1ba9e8-9421-498a-9c8a-01e9772b2924"
  kms_data_key_reuse_period_seconds = 3600
  tags                              = local.tags
}
