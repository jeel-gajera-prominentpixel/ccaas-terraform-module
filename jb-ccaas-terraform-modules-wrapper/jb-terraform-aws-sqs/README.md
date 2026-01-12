# AWS sqs Terraform Module

## How to use this module:

### aws sqs basic module usage with the required input variables:
```terraform
module "sqs_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sqs?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb_sqs"
  fifo_queue     = true
  create_dlq     = true
}
```

### aws sqs advanced module usage with all the optional input variables:



```terraform
module "sqs_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sqs?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
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
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_sqs"></a> [sqs](#module\_sqs) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-sqs | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the sqs, will be appended with the company, lob, env and region to form a sqs name. | `string` | n/a | yes |
| <a name="input_create_dlq"></a> [create\_dlq](#input\_create\_dlq) | Determines whether to create SQS dead letter queue | `bool` | `false` | no |
| <a name="input_create_dlq_redrive_allow_policy"></a> [create\_dlq\_redrive\_allow\_policy](#input\_create\_dlq\_redrive\_allow\_policy) | Determines whether to create a redrive allow policy for the dead letter queue. | `bool` | `true` | no |
| <a name="input_create_queue_policy"></a> [create\_queue\_policy](#input\_create\_queue\_policy) | Whether to create SQS queue policy | `bool` | `false` | no |
| <a name="input_dlq_queue_policy_statements"></a> [dlq\_queue\_policy\_statements](#input\_dlq\_queue\_policy\_statements) | A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage | `any` | `{}` | no |
| <a name="input_dlq_redrive_allow_policy"></a> [dlq\_redrive\_allow\_policy](#input\_dlq\_redrive\_allow\_policy) | The JSON policy to set up the Dead Letter Queue redrive permission, see AWS docs. | `any` | `{}` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_fifo_queue"></a> [fifo\_queue](#input\_fifo\_queue) | Boolean designating a FIFO queue | `bool` | `false` | no |
| <a name="input_kms_data_key_reuse_period_seconds"></a> [kms\_data\_key\_reuse\_period\_seconds](#input\_kms\_data\_key\_reuse\_period\_seconds) | The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours) | `number` | `null` | no |
| <a name="input_kms_master_key_id"></a> [kms\_master\_key\_id](#input\_kms\_master\_key\_id) | The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the sqs, will be appended with the company, lob, env and region to form a sqs name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the sqs, will be appended with the company, lob, env and region to form a sqs name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws sqs , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_queue_policy_statements"></a> [queue\_policy\_statements](#input\_queue\_policy\_statements) | A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage | `any` | `{}` | no |
| <a name="input_redrive_policy"></a> [redrive\_policy](#input\_redrive\_policy) | The JSON policy to set up the Dead Letter Queue, see AWS docs. Note: when specifying maxReceiveCount, you must specify it as an integer (5), and not a string ("5") | `any` | `{}` | no |
| <a name="input_sqs_managed_sse_enabled"></a> [sqs\_managed\_sse\_enabled](#input\_sqs\_managed\_sse\_enabled) | Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys | `bool` | `true` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_default_sqs_dlq_arn"></a> [default\_sqs\_dlq\_arn](#output\_default\_sqs\_dlq\_arn) | The ARN of the SQS queue |
| <a name="output_default_sqs_dlq_id"></a> [default\_sqs\_dlq\_id](#output\_default\_sqs\_dlq\_id) | The URL for the created Amazon SQS queue |
| <a name="output_default_sqs_queue_arn"></a> [default\_sqs\_queue\_arn](#output\_default\_sqs\_queue\_arn) | The ARN of the SQS queue |
| <a name="output_default_sqs_queue_id"></a> [default\_sqs\_queue\_id](#output\_default\_sqs\_queue\_id) | The URL for the created Amazon SQS queue |
| <a name="output_fifo_sqs_queue_arn"></a> [fifo\_sqs\_queue\_arn](#output\_fifo\_sqs\_queue\_arn) | The ARN of the SQS queue |
| <a name="output_fifo_sqs_queue_id"></a> [fifo\_sqs\_queue\_id](#output\_fifo\_sqs\_queue\_id) | The URL for the created Amazon SQS queue |
<!-- END_TF_DOCS -->
