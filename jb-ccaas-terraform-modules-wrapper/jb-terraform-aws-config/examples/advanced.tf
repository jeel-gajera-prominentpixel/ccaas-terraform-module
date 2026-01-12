module "config" {
  source           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-config?ref=main"
  prefix_company   = "jb"
  lob              = "itsd"
  prefix_region    = "usw2"
  application      = "recordings"
  env              = "sandbox"
  create_sns_topic = true
  create_iam_role  = true
  managed_rules = {
    rule1 = {
      description      = "Rule 1 description"
      enabled          = true
      identifier       = "rule1_identifier"
      input_parameters = {}
      tags             = {}
    },
    rule2 = {
      description      = "Rule 2 description"
      enabled          = true
      identifier       = "rule2_identifier"
      input_parameters = {}
      tags             = {}
    }
  }
  force_destroy = false
  s3_bucket_id  = "test-bucket-123"
  s3_bucket_arn = "arn:aws:s3:::test-bucket-123"
  tags          = local.tags
}
