# AWS step function Terraform Module

## How to use this module:

### aws step function basic module usage with the required input variables:
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

### aws step function advanced module usage with all the optional input variables:



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



# jb-terraform-aws-step-functions

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
| <a name="module_step_function"></a> [step\_function](#module\_step\_function) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-step-functions | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the step function, will be appended with the company, lob, env and region to form a step function name. | `string` | n/a | yes |
| <a name="input_attach_policy_json"></a> [attach\_policy\_json](#input\_attach\_policy\_json) | Controls whether policy\_json should be added to IAM role | `bool` | `false` | no |
| <a name="input_cloudwatch_log_group_name"></a> [cloudwatch\_log\_group\_name](#input\_cloudwatch\_log\_group\_name) | Name of Cloudwatch Logs group name to use. | `string` | `null` | no |
| <a name="input_definition"></a> [definition](#input\_definition) | The Amazon States Language definition of the Step Function | `string` | `""` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the step function, will be appended with the company, lob, env and region to form a step function name. | `string` | n/a | yes |
| <a name="input_logging_configuration"></a> [logging\_configuration](#input\_logging\_configuration) | Defines what execution history events are logged and where they are logged | `map(string)` | `{}` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `null` | no |
| <a name="input_policy_json"></a> [policy\_json](#input\_policy\_json) | An additional policy document as JSON to attach to IAM role | `string` | `null` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the step function, will be appended with the company, lob, env and region to form a step function name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws step function, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_publish"></a> [publish](#input\_publish) | Determines whether to set a version of the state machine when it is created. | `bool` | `false` | no |
| <a name="input_service_integrations"></a> [service\_integrations](#input\_service\_integrations) | Map of AWS service integrations to allow in IAM role policy | `any` | `{}` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_type"></a> [type](#input\_type) | Determines whether a Standard or Express state machine is created. The default is STANDARD. Valid Values: STANDARD \| EXPRESS | `string` | `"STANDARD"` | no |
| <a name="input_use_existing_cloudwatch_log_group"></a> [use\_existing\_cloudwatch\_log\_group](#input\_use\_existing\_cloudwatch\_log\_group) | Whether to use an existing CloudWatch log group or create new | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_state_machine_arn"></a> [state\_machine\_arn](#output\_state\_machine\_arn) | The ARN of the State Machine |
| <a name="output_state_machine_id"></a> [state\_machine\_id](#output\_state\_machine\_id) | The ARN of the State Machine |
<!-- END_TF_DOCS -->
