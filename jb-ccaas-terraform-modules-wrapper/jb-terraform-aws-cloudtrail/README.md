# AWS Cloudtrail Terraform Module

## How to use this module:

### aws Cloudtrail basic module usage with the required input variables:

```terraform
module "cloudtrail_log_group_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-cloudtrail?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags           = local.tags
}
```

### aws Cloudtrail advanced module usage with all the optional input variables:

```terraform
module "cloudtrail_log_group_advance" {
  source = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-cloudtrail?ref=<version>"
  prefix_company    = "jb"
  lob               = "itsd"
  prefix_region = "usw2"
  application       = "recordings"
  env               = "sandbox"
  enable_logging                = true
  enable_log_file_validation    = true
  include_global_service_events = false
  is_multi_region_trail         = true
  is_organization_trail         = false
  tags              = local.tags
}
```
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.0, < 2.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.27 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_aws_cloudtrail"></a> [aws\_cloudtrail](#module\_aws\_cloudtrail) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-cloudtrail | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_advanced_event_selector"></a> [advanced\_event\_selector](#input\_advanced\_event\_selector) | Specifies an advanced event selector for enabling data event logging. See: https://www.terraform.io/docs/providers/aws/r/cloudtrail.html for details on this variable | <pre>list(object({<br>    name = optional(string)<br>    field_selector = list(object({<br>      field           = string<br>      ends_with       = optional(list(string))<br>      not_ends_with   = optional(list(string))<br>      equals          = optional(list(string))<br>      not_equals      = optional(list(string))<br>      starts_with     = optional(list(string))<br>      not_starts_with = optional(list(string))<br>    }))<br>  }))</pre> | `[]` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name. | `string` | n/a | yes |
| <a name="input_enable_log_file_validation"></a> [enable\_log\_file\_validation](#input\_enable\_log\_file\_validation) | Specifies whether log file integrity validation is enabled. Creates signed digest for validated contents of logs | `bool` | `true` | no |
| <a name="input_enable_logging"></a> [enable\_logging](#input\_enable\_logging) | Enable logging for the trail | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_include_global_service_events"></a> [include\_global\_service\_events](#input\_include\_global\_service\_events) | Specifies whether the trail is publishing events from global services such as IAM to the log files | `bool` | `false` | no |
| <a name="input_is_multi_region_trail"></a> [is\_multi\_region\_trail](#input\_is\_multi\_region\_trail) | Specifies whether the trail is created in the current region or in all regions | `bool` | `true` | no |
| <a name="input_is_organization_trail"></a> [is\_organization\_trail](#input\_is\_organization\_trail) | The trail is an AWS Organizations trail | `bool` | `false` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the cloudtrail, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_s3_bucket_name"></a> [s3\_bucket\_name](#input\_s3\_bucket\_name) | The name of the custom S3 bucket for CloudTrail logs | `string` | n/a | yes |
| <a name="input_s3_key_prefix"></a> [s3\_key\_prefix](#input\_s3\_key\_prefix) | Prefix for S3 bucket used by Cloudtrail to store logs | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_cloudtrail_arn"></a> [cloudtrail\_arn](#output\_cloudtrail\_arn) | The Amazon Resource Name of the trail |
| <a name="output_cloudtrail_home_region"></a> [cloudtrail\_home\_region](#output\_cloudtrail\_home\_region) | The region in which the trail was created |
| <a name="output_cloudtrail_id"></a> [cloudtrail\_id](#output\_cloudtrail\_id) | The name of the trail |
<!-- END_TF_DOCS -->
