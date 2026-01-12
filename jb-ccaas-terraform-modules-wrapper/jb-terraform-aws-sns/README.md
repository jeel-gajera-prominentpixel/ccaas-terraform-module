# AWS sns Terraform Module

## How to use this module:

### aws sns basic module usage with the required input variables:
```terraform
module "sns_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sns?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
}
```

### aws sns advanced module usage with all the optional input variables:
```terraform
module "sns_advanced" {
  source                      = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-sns?ref=<version>"
  prefix_company              = "jb"
  lob                         = "itsd"
  prefix_region = "usw2"
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
| <a name="module_sns"></a> [sns](#module\_sns) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-sns | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the sns, will be appended with the company, lob, env and region to form a sns name. | `string` | n/a | yes |
| <a name="input_content_based_deduplication"></a> [content\_based\_deduplication](#input\_content\_based\_deduplication) | Boolean indicating whether or not to enable content-based deduplication for FIFO topics. | `bool` | `false` | no |
| <a name="input_create_topic_policy"></a> [create\_topic\_policy](#input\_create\_topic\_policy) | Determines whether an SNS topic policy is created | `bool` | `true` | no |
| <a name="input_delivery_policy"></a> [delivery\_policy](#input\_delivery\_policy) | The SNS delivery policy | `string` | `null` | no |
| <a name="input_display_name"></a> [display\_name](#input\_display\_name) | The display name for the SNS topic | `string` | `null` | no |
| <a name="input_enable_default_topic_policy"></a> [enable\_default\_topic\_policy](#input\_enable\_default\_topic\_policy) | Specifies whether to enable the default topic policy. Defaults to `true` | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_fifo_topic"></a> [fifo\_topic](#input\_fifo\_topic) | Boolean indicating whether or not to create a FIFO (first-in-first-out) topic | `bool` | `false` | no |
| <a name="input_kms_master_key_id"></a> [kms\_master\_key\_id](#input\_kms\_master\_key\_id) | The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the sns, will be appended with the company, lob, env and region to form a sns name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the sns, will be appended with the company, lob, env and region to form a sns name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws sns , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_signature_version"></a> [signature\_version](#input\_signature\_version) | If SignatureVersion should be 1 (SHA1) or 2 (SHA256). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS. | `number` | `null` | no |
| <a name="input_subscriptions"></a> [subscriptions](#input\_subscriptions) | A map of subscription definitions to create | `any` | `{}` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_topic_policy_statements"></a> [topic\_policy\_statements](#input\_topic\_policy\_statements) | A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage | `any` | `{}` | no |
| <a name="input_tracing_config"></a> [tracing\_config](#input\_tracing\_config) | Tracing mode of an Amazon SNS topic. Valid values: PassThrough, Active. | `string` | `null` | no |
| <a name="input_use_name_prefix"></a> [use\_name\_prefix](#input\_use\_name\_prefix) | Determines whether `name` is used as a prefix | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_sns_subscriptions"></a> [sns\_subscriptions](#output\_sns\_subscriptions) | Map of subscriptions created and their attributes |
| <a name="output_sns_topic_arn"></a> [sns\_topic\_arn](#output\_sns\_topic\_arn) | The ARN of the SNS topic, as a more obvious property (clone of id) |
| <a name="output_sns_topic_id"></a> [sns\_topic\_id](#output\_sns\_topic\_id) | The ARN of the SNS topic |
<!-- END_TF_DOCS -->
