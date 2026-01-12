# AWS kinesis stream advance Terraform Module

## How to use this module:

### aws kinesis stream basic module usage with the required input variables:
```terraform
module "kinesis_stream_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kinesis-stream?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags           = local.tags
}
```

### aws kinesis stream advanced module usage with all the optional input variables:
```terraform
module "kinesis_stream_advance" {
  source                       = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kinesis-stream?ref=main"
  prefix_company               = "jb"
  lob                          = "itsd"
  prefix_region = "usw2"
  application                  = "recordings"
  env                          = "sandbox"
  shard_count               = 1
  retention_period          = 24
  shard_level_metrics       = []
  enforce_consumer_deletion = false
  encryption_type           = "NONE"
  kms_key_id                = "test-kms-key-id"
  create_policy_read_only   = true
  create_policy_write_only  = true
  create_policy_admin       = true
  tags                         = local.tags
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
| <a name="module_kinesis_stream"></a> [kinesis\_stream](#module\_kinesis\_stream) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-kinesis-stream | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name. | `string` | n/a | yes |
| <a name="input_create_policy_admin"></a> [create\_policy\_admin](#input\_create\_policy\_admin) | Whether to create IAM Policy (ARN) admin of the Stream | `bool` | `true` | no |
| <a name="input_create_policy_read_only"></a> [create\_policy\_read\_only](#input\_create\_policy\_read\_only) | Whether to create IAM Policy (ARN) read only of the Stream | `bool` | `true` | no |
| <a name="input_create_policy_write_only"></a> [create\_policy\_write\_only](#input\_create\_policy\_write\_only) | Whether to create IAM Policy (ARN) write only of the Stream | `bool` | `true` | no |
| <a name="input_encryption_type"></a> [encryption\_type](#input\_encryption\_type) | The encryption type to use. The only acceptable values are NONE or KMS. | `string` | `"NONE"` | no |
| <a name="input_enforce_consumer_deletion"></a> [enforce\_consumer\_deletion](#input\_enforce\_consumer\_deletion) | A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_kms_key_id"></a> [kms\_key\_id](#input\_kms\_key\_id) | The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias alias/aws/kinesis. | `string` | `""` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the kinesis stream , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_retention_period"></a> [retention\_period](#input\_retention\_period) | Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24. | `number` | `24` | no |
| <a name="input_shard_count"></a> [shard\_count](#input\_shard\_count) | The number of shards that the stream will use | `number` | `1` | no |
| <a name="input_shard_level_metrics"></a> [shard\_level\_metrics](#input\_shard\_level\_metrics) | A list of shard-level CloudWatch metrics which can be enabled for the stream. | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_kinesis_stream_arn"></a> [kinesis\_stream\_arn](#output\_kinesis\_stream\_arn) | The Amazon Resource Name (ARN) specifying the Stream |
| <a name="output_kinesis_stream_iam_policy_admin_arn"></a> [kinesis\_stream\_iam\_policy\_admin\_arn](#output\_kinesis\_stream\_iam\_policy\_admin\_arn) | The IAM Policy (ARN) admin of the Stream |
| <a name="output_kinesis_stream_iam_policy_read_only_arn"></a> [kinesis\_stream\_iam\_policy\_read\_only\_arn](#output\_kinesis\_stream\_iam\_policy\_read\_only\_arn) | The IAM Policy (ARN) read only of the Stream |
| <a name="output_kinesis_stream_iam_policy_write_only_arn"></a> [kinesis\_stream\_iam\_policy\_write\_only\_arn](#output\_kinesis\_stream\_iam\_policy\_write\_only\_arn) | The IAM Policy (ARN) write only of the Stream |
| <a name="output_kinesis_stream_name"></a> [kinesis\_stream\_name](#output\_kinesis\_stream\_name) | The unique Stream name |
| <a name="output_kinesis_stream_shard_count"></a> [kinesis\_stream\_shard\_count](#output\_kinesis\_stream\_shard\_count) | The count of Shards for this Stream |
<!-- END_TF_DOCS -->
