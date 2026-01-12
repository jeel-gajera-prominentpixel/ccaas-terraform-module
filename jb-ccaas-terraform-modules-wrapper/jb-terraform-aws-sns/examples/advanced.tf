module "sns_advanced" {
  source                      = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sns?ref=main"
  prefix_company              = "jb"
  lob                         = "itsd"
  prefix_region               = "usw2"
  application                 = "recordings"
  env                         = "sandbox"
  signature_version           = 2
  use_name_prefix             = true
  display_name                = "complete"
  name                        = "jb "
  kms_master_key_id           = "arn:aws:kms:us-east-1:123456789012:key/abcd1234-ab12-34cd-56ef-1234567890ab"
  tracing_config              = "Active"
  fifo_topic                  = true
  content_based_deduplication = true
  delivery_policy = jsonencode({
    "http" : {
      "defaultHealthyRetryPolicy" : {
        "minDelayTarget" : 20,
        "maxDelayTarget" : 20,
        "numRetries" : 3,
        "numMaxDelayRetries" : 0,
        "numNoDelayRetries" : 0,
        "numMinDelayRetries" : 0,
        "backoffFunction" : "linear"
      },
      "disableSubscriptionOverrides" : false,
      "defaultThrottlePolicy" : {
        "maxReceivesPerSecond" : 1
      }
    }
  })
  create_topic_policy         = true
  enable_default_topic_policy = true
  subscriptions = {
    sqs = {
      protocol = "sqs"
      endpoint = "arn:aws:sqs:us-east-1:123456789012:random-queue"
    }
  }
  tags = local.tags
}
