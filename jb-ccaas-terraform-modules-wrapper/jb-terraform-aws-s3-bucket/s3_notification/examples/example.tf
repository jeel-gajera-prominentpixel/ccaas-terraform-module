module "s3_notification" {
  source               = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket/s3_notification?ref=main"
  application          = "recordings"
  prefix_company       = "jb"
  prefix_region        = "usw2"
  env                  = "sandbox"
  lob                  = "itsd"
  create               = true
  create_sns_policy    = false
  create_sqs_policy    = false
  bucket               = "cla-test"
  bucket_arn           = "arn:aws:s3:::cla-test"
  eventbridge          = false
  lambda_notifications = {}
  sqs_notifications    = {}
  sns_notifications    = {}
}
