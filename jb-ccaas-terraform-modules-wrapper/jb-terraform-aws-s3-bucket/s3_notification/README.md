# s3_notification

## How to use this module:

### aws s3 notification module usage with the required input variables:
```terraform
module "s3_notification" {
  source                = "git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-s3-bucket/modules/notification?ref=main"
  application           = "recordings"
  prefix_company        = "jb"
  prefix_region         = "usw2"
  env                   = "sandbox"
  lob                   = "itsd"
  create                = true
  create_sns_policy     = false
  create_sqs_policy     = false
  bucket                = "cla-test"
  bucket_arn            = "arn:aws:s3:::cla-test"
  eventbridge           = false
  lambda_notifications  = {}
  sqs_notifications     = {}
  sns_notifications     = {}
}
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.0, < 2.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.27 |
| <a name="requirement_external"></a> [external](#requirement\_external) | >= 2.3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.5 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_s3_notification"></a> [s3\_notification](#module\_s3\_notification) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-s3-bucket/modules/notification | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the rds, will be appended with the company, lob, env and region to form a rds name. | `string` | n/a | yes |
| <a name="input_bucket"></a> [bucket](#input\_bucket) | Name of S3 bucket to use | `string` | `""` | no |
| <a name="input_bucket_arn"></a> [bucket\_arn](#input\_bucket\_arn) | ARN of S3 bucket to use in policies | `string` | `null` | no |
| <a name="input_create"></a> [create](#input\_create) | Whether to create this resource or not? | `bool` | `true` | no |
| <a name="input_create_sns_policy"></a> [create\_sns\_policy](#input\_create\_sns\_policy) | Whether to create a policy for SNS permissions or not? | `bool` | `false` | no |
| <a name="input_create_sqs_policy"></a> [create\_sqs\_policy](#input\_create\_sqs\_policy) | Whether to create a policy for SQS permissions or not? | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_eventbridge"></a> [eventbridge](#input\_eventbridge) | Whether to enable Amazon EventBridge notifications | `bool` | `null` | no |
| <a name="input_lambda_notifications"></a> [lambda\_notifications](#input\_lambda\_notifications) | Map of S3 bucket notifications to Lambda function | `any` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the rds, will be appended with the company, lob, env and region to form a rds name | `string` | n/a | yes |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the rds, will be appended with the company, lob, env and region to form a rds name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the rds , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_sns_notifications"></a> [sns\_notifications](#input\_sns\_notifications) | Map of S3 bucket notifications to SNS topic | `any` | `{}` | no |
| <a name="input_sqs_notifications"></a> [sqs\_notifications](#input\_sqs\_notifications) | Map of S3 bucket notifications to SQS queue | `any` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_s3_bucket_notification_id"></a> [s3\_bucket\_notification\_id](#output\_s3\_bucket\_notification\_id) | ID of S3 bucket |
<!-- END_TF_DOCS -->
