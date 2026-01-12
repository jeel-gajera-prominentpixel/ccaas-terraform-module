# AWS EVENT BRIDGE Terraform Module

## How to use this module:

### aws eventbridge basic module usage with the required input variables:
```terraform
module "eventbridge_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-eventbridge?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags           = local.tags
}
```

### aws eventbridge advanced module usage with all the optional input variables:
```terraform
module "eventbridge_advance" {
  source                    = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-eventbridge?ref=main"
  prefix_company            = "jb"
  lob                       = "itsd"
  prefix_region = "usw2"
  application               = "recordings"
  env                       = "sandbox"
  create_bus                = true
  create_schemas_discoverer = false
  attach_tracing_policy     = false
  attach_kinesis_policy     = false
  kinesis_target_arns       = []
  attach_sfn_policy         = false
  sfn_target_arns           = []
  attach_sqs_policy         = false
  sqs_target_arns           = []
  attach_cloudwatch_policy  = false
  cloudwatch_target_arns    = []
  append_rule_postfix       = true
  attach_ecs_policy         = false
  ecs_target_arns           = []
  rules                     = {}
  targets                   = {}
  attach_policy_json        = false
  policy_json               = null
  attach_policy_jsons       = false
  policy_jsons              = []
  number_of_policy_jsons    = 0
  attach_policies           = false
  policies                  = []
  number_of_policies        = 0
  attach_policy_statements  = false
  policy_statements         = {}
  tags                      = local.tags
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
| <a name="module_eventbridge"></a> [eventbridge](#module\_eventbridge) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-eventbridge | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_destinations"></a> [api\_destinations](#input\_api\_destinations) | A map of objects with EventBridge Destination definitions. | `map(any)` | `{}` | no |
| <a name="input_append_pipe_postfix"></a> [append\_pipe\_postfix](#input\_append\_pipe\_postfix) | Controls whether to append '-pipe' to the name of the pipe | `bool` | `true` | no |
| <a name="input_append_rule_postfix"></a> [append\_rule\_postfix](#input\_append\_rule\_postfix) | Controls whether to append '-rule' to the name of the rule | `bool` | `true` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name. | `string` | n/a | yes |
| <a name="input_attach_api_destination_policy"></a> [attach\_api\_destination\_policy](#input\_attach\_api\_destination\_policy) | Controls whether the API Destination policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_cloudwatch_policy"></a> [attach\_cloudwatch\_policy](#input\_attach\_cloudwatch\_policy) | Controls whether the Cloudwatch policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_ecs_policy"></a> [attach\_ecs\_policy](#input\_attach\_ecs\_policy) | Controls whether the ECS policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_kinesis_policy"></a> [attach\_kinesis\_policy](#input\_attach\_kinesis\_policy) | Controls whether the Kinesis policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_policies"></a> [attach\_policies](#input\_attach\_policies) | Controls whether list of policies should be added to IAM role | `bool` | `false` | no |
| <a name="input_attach_policy_json"></a> [attach\_policy\_json](#input\_attach\_policy\_json) | Controls whether policy\_json should be added to IAM role | `bool` | `false` | no |
| <a name="input_attach_policy_jsons"></a> [attach\_policy\_jsons](#input\_attach\_policy\_jsons) | Controls whether policy\_jsons should be added to IAM role | `bool` | `false` | no |
| <a name="input_attach_policy_statements"></a> [attach\_policy\_statements](#input\_attach\_policy\_statements) | Controls whether policy\_statements should be added to IAM role | `bool` | `false` | no |
| <a name="input_attach_sfn_policy"></a> [attach\_sfn\_policy](#input\_attach\_sfn\_policy) | Controls whether the StepFunction policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_sqs_policy"></a> [attach\_sqs\_policy](#input\_attach\_sqs\_policy) | Controls whether the SQS policy should be added to IAM role for EventBridge Target | `bool` | `false` | no |
| <a name="input_attach_tracing_policy"></a> [attach\_tracing\_policy](#input\_attach\_tracing\_policy) | Controls whether X-Ray tracing policy should be added to IAM role for EventBridge | `bool` | `false` | no |
| <a name="input_cloudwatch_target_arns"></a> [cloudwatch\_target\_arns](#input\_cloudwatch\_target\_arns) | The Amazon Resource Name (ARN) of the Cloudwatch Log Streams you want to use as EventBridge targets | `list(string)` | `[]` | no |
| <a name="input_connections"></a> [connections](#input\_connections) | A map of objects with EventBridge Connection definitions. | `any` | `{}` | no |
| <a name="input_create_api_destinations"></a> [create\_api\_destinations](#input\_create\_api\_destinations) | Controls whether EventBridge Destination resources should be created | `bool` | `false` | no |
| <a name="input_create_bus"></a> [create\_bus](#input\_create\_bus) | Controls whether EventBridge Bus resource should be created | `bool` | `true` | no |
| <a name="input_create_connections"></a> [create\_connections](#input\_create\_connections) | Controls whether EventBridge Connection resources should be created | `bool` | `false` | no |
| <a name="input_create_pipes"></a> [create\_pipes](#input\_create\_pipes) | Controls whether EventBridge Pipes resources should be created | `bool` | `true` | no |
| <a name="input_create_schemas_discoverer"></a> [create\_schemas\_discoverer](#input\_create\_schemas\_discoverer) | Controls whether default schemas discoverer should be created | `bool` | `false` | no |
| <a name="input_ecs_target_arns"></a> [ecs\_target\_arns](#input\_ecs\_target\_arns) | The Amazon Resource Name (ARN) of the AWS ECS Tasks you want to use as EventBridge targets | `list(string)` | `[]` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_kinesis_target_arns"></a> [kinesis\_target\_arns](#input\_kinesis\_target\_arns) | The Amazon Resource Name (ARN) of the Kinesis Streams you want to use as EventBridge targets | `list(string)` | `[]` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_number_of_policies"></a> [number\_of\_policies](#input\_number\_of\_policies) | Number of policies to attach to IAM role | `number` | `0` | no |
| <a name="input_number_of_policy_jsons"></a> [number\_of\_policy\_jsons](#input\_number\_of\_policy\_jsons) | Number of policies JSON to attach to IAM role | `number` | `0` | no |
| <a name="input_pipes"></a> [pipes](#input\_pipes) | A map of objects with EventBridge Pipe definitions. | `any` | `{}` | no |
| <a name="input_policies"></a> [policies](#input\_policies) | List of policy statements ARN to attach to IAM role | `list(string)` | `[]` | no |
| <a name="input_policy_json"></a> [policy\_json](#input\_policy\_json) | An additional policy document as JSON to attach to IAM role | `string` | `null` | no |
| <a name="input_policy_jsons"></a> [policy\_jsons](#input\_policy\_jsons) | List of additional policy documents as JSON to attach to IAM role | `list(string)` | `[]` | no |
| <a name="input_policy_statements"></a> [policy\_statements](#input\_policy\_statements) | Map of dynamic policy statements to attach to IAM role | `any` | `{}` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the eventbridge , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_role_description"></a> [role\_description](#input\_role\_description) | Description of IAM role to use for EventBridge | `string` | `null` | no |
| <a name="input_role_force_detach_policies"></a> [role\_force\_detach\_policies](#input\_role\_force\_detach\_policies) | Specifies to force detaching any policies the IAM role has before destroying it. | `bool` | `true` | no |
| <a name="input_role_name"></a> [role\_name](#input\_role\_name) | Name of IAM role to use for EventBridge | `string` | `null` | no |
| <a name="input_role_path"></a> [role\_path](#input\_role\_path) | Path of IAM role to use for EventBridge | `string` | `null` | no |
| <a name="input_role_permissions_boundary"></a> [role\_permissions\_boundary](#input\_role\_permissions\_boundary) | The ARN of the policy that is used to set the permissions boundary for the IAM role used by EventBridge | `string` | `null` | no |
| <a name="input_role_tags"></a> [role\_tags](#input\_role\_tags) | A map of tags to assign to IAM role | `map(string)` | `{}` | no |
| <a name="input_rules"></a> [rules](#input\_rules) | A map of objects with EventBridge Rule definitions. | `map(any)` | `{}` | no |
| <a name="input_sfn_target_arns"></a> [sfn\_target\_arns](#input\_sfn\_target\_arns) | The Amazon Resource Name (ARN) of the StepFunctions you want to use as EventBridge targets | `list(string)` | `[]` | no |
| <a name="input_sqs_target_arns"></a> [sqs\_target\_arns](#input\_sqs\_target\_arns) | The Amazon Resource Name (ARN) of the AWS SQS Queues you want to use as EventBridge targets | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_targets"></a> [targets](#input\_targets) | A map of objects with EventBridge Target definitions. | `any` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_eventbridge_rule_arns"></a> [eventbridge\_rule\_arns](#output\_eventbridge\_rule\_arns) | The EventBridge Rule ARNs |
<!-- END_TF_DOCS -->
