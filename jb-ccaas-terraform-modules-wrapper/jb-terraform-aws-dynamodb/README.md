# AWS dynamodb-table Terraform Module

## How to use this module:

### aws dynamodb-table basic module usage with the required input variables:
```terraform
module "dynamodb-table" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-dynamodb?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
}
```

### aws dynamodb-table advanced module usage with all the optional input variables:


```terraform
module "dynamodb-table" {
  source                      = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-dynamodb?ref=<version>"
  prefix_company              = "jb"
  lob                         = "itsd"
  prefix_region = "usw2"
  application                 = "recordings"
  env                         = "sandbox"
  name                        = "jb-dynamodb-table"
  hash_key                    = "id"
  range_key                   = "title"
  table_class                 = "STANDARD"
  deletion_protection_enabled = false
  attributes = [
    {
      name = "id"
      type = "N"
    },
    {
      name = "title"
      type = "S"
    },
  ]
  import_table = {
    input_format           = "DYNAMODB_JSON"
    input_compression_type = "NONE"
    bucket                 = "jb-bucket"
    key_prefix             = "/example"
  }
  stream_enabled                     = true
  stream_view_type                   = "NEW_AND_OLD_IMAGES"
  server_side_encryption_enabled     = true
  server_side_encryption_kms_key_arn = "arn:aws:kms:us-east-1:123456789012:key/abcd1234-a123-456a-a12b-a123b4cd5678"
  global_secondary_indexes = [
    {
      name               = "TitleIndex"
      hash_key           = "title"
      range_key          = "age"
      projection_type    = "INCLUDE"
      non_key_attributes = ["id"]
    }
  ]
  replica_regions = [{
    region_name            = "eu-west-2"
    kms_key_arn            = "arn:aws:kms:us-east-1:123456789012:key/abcd1234-a123-456a-a12b"
    propagate_tags         = true
    point_in_time_recovery = true
  }]
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
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.5 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_dynamodb-table"></a> [dynamodb-table](#module\_dynamodb-table) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-dynamodb-table | feature/update-dynamodb-module |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the dynamodb-table, will be appended with the company, lob, env and region to form a dynamodb-table name. | `string` | n/a | yes |
| <a name="input_attributes"></a> [attributes](#input\_attributes) | List of nested attribute definitions. Only required for hash\_key and range\_key attributes. Each attribute has two properties: name - (Required) The name of the attribute, type - (Required) Attribute type, which must be a scalar type: S, N, or B for (S)tring, (N)umber or (B)inary data | `list(map(string))` | `[]` | no |
| <a name="input_deletion_protection_enabled"></a> [deletion\_protection\_enabled](#input\_deletion\_protection\_enabled) | Enables deletion protection for table | `bool` | `null` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_global_secondary_indexes"></a> [global\_secondary\_indexes](#input\_global\_secondary\_indexes) | Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. | `any` | `[]` | no |
| <a name="input_hash_key"></a> [hash\_key](#input\_hash\_key) | The attribute to use as the hash (partition) key. Must also be defined as an attribute | `string` | `null` | no |
| <a name="input_import_table"></a> [import\_table](#input\_import\_table) | Configurations for importing s3 data into a new table. | `any` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the dynamodb-table, will be appended with the company, lob, env and region to form a dynamodb-table name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_point_in_time_recovery_enabled"></a> [point\_in\_time\_recovery\_enabled](#input\_point\_in\_time\_recovery\_enabled) | Whether to enable point-in-time recovery | `bool` | `false` | no |
| <a name="input_point_in_time_recovery_period_in_days"></a> [point\_in\_time\_recovery\_period\_in\_days](#input\_point\_in\_time\_recovery\_period\_in\_days) | Number of preceding days for which continuous backups are taken and maintained. Default 35 | `number` | `null` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the dynamodb-table, will be appended with the company, lob, env and region to form a dynamodb-table name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the dynamodb-table , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_range_key"></a> [range\_key](#input\_range\_key) | The attribute to use as the range (sort) key. Must also be defined as an attribute | `string` | `null` | no |
| <a name="input_replica_regions"></a> [replica\_regions](#input\_replica\_regions) | Region names for creating replicas for a global DynamoDB table. | `any` | `[]` | no |
| <a name="input_server_side_encryption_enabled"></a> [server\_side\_encryption\_enabled](#input\_server\_side\_encryption\_enabled) | Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK) | `bool` | `false` | no |
| <a name="input_server_side_encryption_kms_key_arn"></a> [server\_side\_encryption\_kms\_key\_arn](#input\_server\_side\_encryption\_kms\_key\_arn) | The ARN of the CMK that should be used for the AWS KMS encryption. This attribute should only be specified if the key is different from the default DynamoDB CMK, alias/aws/dynamodb. | `string` | `null` | no |
| <a name="input_stream_enabled"></a> [stream\_enabled](#input\_stream\_enabled) | Indicates whether Streams are to be enabled (true) or disabled (false). | `bool` | `false` | no |
| <a name="input_stream_view_type"></a> [stream\_view\_type](#input\_stream\_view\_type) | When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are KEYS\_ONLY, NEW\_IMAGE, OLD\_IMAGE, NEW\_AND\_OLD\_IMAGES. | `string` | `null` | no |
| <a name="input_table_class"></a> [table\_class](#input\_table\_class) | The storage class of the table. Valid values are STANDARD and STANDARD\_INFREQUENT\_ACCESS | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_ttl_attribute_name"></a> [ttl\_attribute\_name](#input\_ttl\_attribute\_name) | The name of the table attribute to store the TTL timestamp in | `string` | `""` | no |
| <a name="input_ttl_enabled"></a> [ttl\_enabled](#input\_ttl\_enabled) | Indicates whether ttl is enabled | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dynamodb_table_arn"></a> [dynamodb\_table\_arn](#output\_dynamodb\_table\_arn) | ARN of the DynamoDB table |
| <a name="output_dynamodb_table_attributes"></a> [dynamodb\_table\_attributes](#output\_dynamodb\_table\_attributes) | Dynamo DB table attributes |
| <a name="output_dynamodb_table_billing_mode"></a> [dynamodb\_table\_billing\_mode](#output\_dynamodb\_table\_billing\_mode) | Billing mode of the DynamoDB table |
| <a name="output_dynamodb_table_deletion_protection_enabled"></a> [dynamodb\_table\_deletion\_protection\_enabled](#output\_dynamodb\_table\_deletion\_protection\_enabled) | Enables deletion protection for table |
| <a name="output_dynamodb_table_global_secondary_index"></a> [dynamodb\_table\_global\_secondary\_index](#output\_dynamodb\_table\_global\_secondary\_index) | Global secondary index of the DynamoDB table |
| <a name="output_dynamodb_table_hash_key"></a> [dynamodb\_table\_hash\_key](#output\_dynamodb\_table\_hash\_key) | Attribute to use as the hash (partition) key. |
| <a name="output_dynamodb_table_id"></a> [dynamodb\_table\_id](#output\_dynamodb\_table\_id) | ID of the DynamoDB table |
| <a name="output_dynamodb_table_local_secondary_index"></a> [dynamodb\_table\_local\_secondary\_index](#output\_dynamodb\_table\_local\_secondary\_index) | Local secondary index of the DynamoDB table |
| <a name="output_dynamodb_table_name"></a> [dynamodb\_table\_name](#output\_dynamodb\_table\_name) | Name of the DynamoDB table |
| <a name="output_dynamodb_table_range_key"></a> [dynamodb\_table\_range\_key](#output\_dynamodb\_table\_range\_key) | Range key of the DynamoDB table |
| <a name="output_dynamodb_table_replica"></a> [dynamodb\_table\_replica](#output\_dynamodb\_table\_replica) | Replica of the DynamoDB table |
| <a name="output_dynamodb_table_stream_arn"></a> [dynamodb\_table\_stream\_arn](#output\_dynamodb\_table\_stream\_arn) | The ARN of the Table Stream. Only available when var.stream\_enabled is true |
| <a name="output_dynamodb_table_stream_label"></a> [dynamodb\_table\_stream\_label](#output\_dynamodb\_table\_stream\_label) | A timestamp, in ISO 8601 format of the Table Stream. Only available when var.stream\_enabled is true |
<!-- END_TF_DOCS -->
