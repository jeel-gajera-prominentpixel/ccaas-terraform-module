# examples

# s3_notification

## How to use this module:

### aws s3 notification module usage with the required input variables:
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

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.0, < 2.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.27 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_s3_notification"></a> [s3\_notification](#module\_s3\_notification) | git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket/s3_notification | main |

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END_TF_DOCS -->
