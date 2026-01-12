# AWS ssm parameter store Terraform Module

## How to use this module:

### aws ssm parameter store basic module usage with the required input variables:
```terraform
module "ssm_store" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-ssm-parameter-store?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb-parameters"
}
```

### aws ssm parameter store advanced module usage with all the optional input variables:


```terraform
module "ssm_store" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-ssm-parameter-store?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb-parameters"
  parameter_write = [
    {
      name        = "/cp/prod/app/database/master_password"
      value       = "password1"
      type        = "String"
      description = "Production database master password"
    }
  ]
  environment = "development"
  tags        = local.tags
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
| <a name="module_ssm_store"></a> [ssm\_store](#module\_ssm\_store) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-ssm-parameter-store | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_environment"></a> [environment](#input\_environment) | ID element. Usually used for region e.g. 'uw2', 'us-west-2', OR role 'prod', 'staging', 'dev', 'UAT' | `string` | `null` | no |
| <a name="input_kms_arn"></a> [kms\_arn](#input\_kms\_arn) | The ARN of a KMS key used to encrypt and decrypt SecretString values | `string` | `""` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_parameter_read"></a> [parameter\_read](#input\_parameter\_read) | List of parameters to read from SSM. These must already exist otherwise an error is returned. Can be used with `parameter_write` as long as the parameters are different. | `list(string)` | `[]` | no |
| <a name="input_parameter_write"></a> [parameter\_write](#input\_parameter\_write) | List of maps with the parameter values to write to SSM Parameter Store | `list(map(string))` | `[]` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_regex_replace_chars"></a> [regex\_replace\_chars](#input\_regex\_replace\_chars) | Terraform regular expression (regex) string.Characters matching the regex will be removed from the ID elements.If not set, /[^a-zA-Z0-9-]/ is used to remove all characters other than hyphens, letters and digits. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_parameter_arn_map"></a> [parameter\_arn\_map](#output\_parameter\_arn\_map) | A map of the names and ARNs created |
| <a name="output_parameter_map"></a> [parameter\_map](#output\_parameter\_map) | A map of the names and values created |
| <a name="output_parameter_names"></a> [parameter\_names](#output\_parameter\_names) | List of key names |
| <a name="output_parameter_values"></a> [parameter\_values](#output\_parameter\_values) | List of values |
<!-- END_TF_DOCS -->
