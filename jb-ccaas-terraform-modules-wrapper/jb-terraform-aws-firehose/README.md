# AWS Firehose Terraform Module

## How to use this module:

### aws Firehose basic module usage with the required input variables:
```terraform
module "firehose" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-firehose?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags           = local.tags
}

```

### aws Firehose advanced module usage with all the optional input variables:



```terraform
module "firehose" {
  source                             = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-firehose?ref=<version>"
  prefix_company                     = "jb"
  lob                                = "itsd"
  prefix_region = "usw2"
  application                        = "recordings"
  env                                = "sandbox"
  destination                        = "s3"
  input_source                       = "kinesis"
  s3_bucket_arn                      = "arn:aws:s3:::s3-internal-connect-storage"
  s3_prifix                          = "example/"
  kinesis_source_stream_arn          = "arn:aws:kinesis:us-east-1:123456789012:stream/my-stream"
  enable_s3_backup                   = true
  s3_backup_bucket_arn               = "arn:aws:s3:::my-backup-bucket"
  sse_kms_key_type                   = "CUSTOMER_MANAGED_CMK"
  s3_backup_prefix                   = "backup/"
  s3_backup_error_output_prefix      = "error/"
  s3_backup_buffering_interval       = 100
  s3_backup_buffering_size           = 100
  s3_backup_compression              = "GZIP"
  s3_backup_use_existing_role        = false
  s3_backup_role_arn                 = "arn:aws:iam::AWS_ACCOUNT_ID:role/ROLE_NAME"
  s3_backup_enable_encryption        = true
  s3_backup_kms_key_arn              = "arn:aws:iam::AWS_ACCOUNT_ID:role/ROLE_NAME"
  s3_backup_enable_log               = true
  create_application_role            = false
  create_application_role_policy     = false
  application_role_service_principal = "arn:aws:iam::AWS_ACCOUNT_ID:role/ROLE_NAME"
  enable_sse                         = true
  sse_kms_key_arn                    = "arn:aws:kms:us-west-2:123456789012:key/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  tags                               = local.tags

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
| <a name="module_firehose"></a> [firehose](#module\_firehose) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-kinesis-firehose | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_append_delimiter_to_record"></a> [append\_delimiter\_to\_record](#input\_append\_delimiter\_to\_record) | To configure your delivery stream to add a new line delimiter between records in objects that are delivered to Amazon S3. | `bool` | `false` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_application_role_service_principal"></a> [application\_role\_service\_principal](#input\_application\_role\_service\_principal) | AWS Service Principal to assume application role | `string` | `null` | no |
| <a name="input_buffering_interval"></a> [buffering\_interval](#input\_buffering\_interval) | Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination | `number` | `300` | no |
| <a name="input_buffering_size"></a> [buffering\_size](#input\_buffering\_size) | Buffer incoming data to the specified size, in MBs, before delivering it to the destination. | `number` | `5` | no |
| <a name="input_create_application_role"></a> [create\_application\_role](#input\_create\_application\_role) | Set it to true to create role to be used by the source | `bool` | `false` | no |
| <a name="input_create_application_role_policy"></a> [create\_application\_role\_policy](#input\_create\_application\_role\_policy) | Set it to true to create policy to the role used by the source | `bool` | `false` | no |
| <a name="input_destination"></a> [destination](#input\_destination) | This is the destination to where the data is delivered | `string` | n/a | yes |
| <a name="input_dynamic_partition_append_delimiter_to_record"></a> [dynamic\_partition\_append\_delimiter\_to\_record](#input\_dynamic\_partition\_append\_delimiter\_to\_record) | DEPRECATED!! Use var append\_delimiter\_to\_record instead!! Use To configure your delivery stream to add a new line delimiter between records in objects that are delivered to Amazon S3. | `bool` | `false` | no |
| <a name="input_dynamic_partition_enable_record_deaggregation"></a> [dynamic\_partition\_enable\_record\_deaggregation](#input\_dynamic\_partition\_enable\_record\_deaggregation) | Data deaggregation is the process of parsing through the records in a delivery stream and separating the records based either on valid JSON or on the specified delimiter | `bool` | `false` | no |
| <a name="input_dynamic_partition_record_deaggregation_delimiter"></a> [dynamic\_partition\_record\_deaggregation\_delimiter](#input\_dynamic\_partition\_record\_deaggregation\_delimiter) | Specifies the delimiter to be used for parsing through the records in the delivery stream and deaggregating them | `string` | `null` | no |
| <a name="input_dynamic_partition_record_deaggregation_type"></a> [dynamic\_partition\_record\_deaggregation\_type](#input\_dynamic\_partition\_record\_deaggregation\_type) | Data deaggregation is the process of parsing through the records in a delivery stream and separating the records based either on valid JSON or on the specified delimiter | `string` | `"JSON"` | no |
| <a name="input_enable_s3_backup"></a> [enable\_s3\_backup](#input\_enable\_s3\_backup) | The Amazon S3 backup mode | `bool` | `false` | no |
| <a name="input_enable_s3_encryption"></a> [enable\_s3\_encryption](#input\_enable\_s3\_encryption) | Indicates if want use encryption in S3 bucket. | `bool` | `false` | no |
| <a name="input_enable_sse"></a> [enable\_sse](#input\_enable\_sse) | Whether to enable encryption at rest. Only makes sense when source is Direct Put | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_input_source"></a> [input\_source](#input\_input\_source) | This is the kinesis firehose source | `string` | `"direct-put"` | no |
| <a name="input_kinesis_source_is_encrypted"></a> [kinesis\_source\_is\_encrypted](#input\_kinesis\_source\_is\_encrypted) | Indicates if Kinesis data stream source is encrypted | `bool` | `false` | no |
| <a name="input_kinesis_source_kms_arn"></a> [kinesis\_source\_kms\_arn](#input\_kinesis\_source\_kms\_arn) | Kinesis Source KMS Key to add Firehose role to decrypt the records. | `string` | `null` | no |
| <a name="input_kinesis_source_stream_arn"></a> [kinesis\_source\_stream\_arn](#input\_kinesis\_source\_stream\_arn) | The kinesis stream used as the source of the firehose delivery stream | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_redshift_copy_options"></a> [redshift\_copy\_options](#input\_redshift\_copy\_options) | Copy options for copying the data from the s3 intermediate bucket into redshift, for example to change the default delimiter | `string` | `null` | no |
| <a name="input_s3_backup_bucket_arn"></a> [s3\_backup\_bucket\_arn](#input\_s3\_backup\_bucket\_arn) | The ARN of the S3 backup bucket | `string` | `null` | no |
| <a name="input_s3_backup_buffering_interval"></a> [s3\_backup\_buffering\_interval](#input\_s3\_backup\_buffering\_interval) | Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination. | `number` | `300` | no |
| <a name="input_s3_backup_buffering_size"></a> [s3\_backup\_buffering\_size](#input\_s3\_backup\_buffering\_size) | Buffer incoming data to the specified size, in MBs, before delivering it to the destination. | `number` | `5` | no |
| <a name="input_s3_backup_compression"></a> [s3\_backup\_compression](#input\_s3\_backup\_compression) | The compression format | `string` | `"UNCOMPRESSED"` | no |
| <a name="input_s3_backup_enable_encryption"></a> [s3\_backup\_enable\_encryption](#input\_s3\_backup\_enable\_encryption) | Indicates if want enable KMS Encryption in S3 Backup Bucket. | `bool` | `false` | no |
| <a name="input_s3_backup_enable_log"></a> [s3\_backup\_enable\_log](#input\_s3\_backup\_enable\_log) | Enables or disables the logging | `bool` | `true` | no |
| <a name="input_s3_backup_error_output_prefix"></a> [s3\_backup\_error\_output\_prefix](#input\_s3\_backup\_error\_output\_prefix) | Prefix added to failed records before writing them to S3 | `string` | `null` | no |
| <a name="input_s3_backup_kms_key_arn"></a> [s3\_backup\_kms\_key\_arn](#input\_s3\_backup\_kms\_key\_arn) | Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will be used. | `string` | `null` | no |
| <a name="input_s3_backup_prefix"></a> [s3\_backup\_prefix](#input\_s3\_backup\_prefix) | The YYYY/MM/DD/HH time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket | `string` | `null` | no |
| <a name="input_s3_backup_role_arn"></a> [s3\_backup\_role\_arn](#input\_s3\_backup\_role\_arn) | The role that Kinesis Data Firehose can use to access S3 Backup. | `string` | `null` | no |
| <a name="input_s3_backup_use_existing_role"></a> [s3\_backup\_use\_existing\_role](#input\_s3\_backup\_use\_existing\_role) | Indicates if want use the kinesis firehose role to s3 backup bucket access. | `bool` | `true` | no |
| <a name="input_s3_bucket_arn"></a> [s3\_bucket\_arn](#input\_s3\_bucket\_arn) | The ARN of the S3 destination bucket | `string` | `null` | no |
| <a name="input_s3_error_output_prefix"></a> [s3\_error\_output\_prefix](#input\_s3\_error\_output\_prefix) | Prefix added to failed records before writing them to S3. This prefix appears immediately following the bucket name. | `string` | `null` | no |
| <a name="input_s3_kms_key_arn"></a> [s3\_kms\_key\_arn](#input\_s3\_kms\_key\_arn) | Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will be used | `string` | `null` | no |
| <a name="input_s3_prefix"></a> [s3\_prefix](#input\_s3\_prefix) | Specifies the  s3 bucket prefix | `string` | `null` | no |
| <a name="input_sse_kms_key_arn"></a> [sse\_kms\_key\_arn](#input\_sse\_kms\_key\_arn) | Amazon Resource Name (ARN) of the encryption key | `string` | `null` | no |
| <a name="input_sse_kms_key_type"></a> [sse\_kms\_key\_type](#input\_sse\_kms\_key\_type) | n/a | `string` | `"AWS_OWNED_CMK"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_firehose_arn"></a> [firehose\_arn](#output\_firehose\_arn) | Configured firehose |
<!-- END_TF_DOCS -->
