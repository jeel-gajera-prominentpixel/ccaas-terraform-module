module "aws_s3_bucket_advanced" {
  source = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket?ref=main"

  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  force_destroy  = true
  versioning     = true
  lifecycle_rule = [{
    id     = "retention"
    status = "Enabled"
    expiration = {
      days = 7
    }
    noncurrent_version_expiration = {
      noncurrent_days = 10
    }
  }]
  tags = local.tags
}


module "s3_bucket_notifications" {
  count  = var.lambda_trigger ? 1 : 0
  source = "../../jb-ccaas-terraform-modules/terraform-aws-s3-bucket/modules/notification"
  bucket = "example-bucket"
  lambda_notifications = {
    lambda1 = {
      function_arn  = "arn:aws:lambda:us-east-1:123456789012:function:my-example-function"
      function_name = "jb-lambda"
      events        = ["s3:ObjectCreated:Put"]
      filter_prefix = "prefix/"
      filter_suffix = ".json"
    }
  }
}
