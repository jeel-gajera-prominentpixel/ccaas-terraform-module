# AWS KMS Terraform Module

## How to use this module:

### aws KMS basic module usage with the required input variables:
```terraform
module "aws_kms" {
  source      = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kms?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  service        = "s3"
  env            = "sandbox"
  tags           = local.tags
}
```

### aws KMS advanced module usage with all the optional input variables:



```terraform
module "aws_kms" {

  source                                 = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-kms?ref=<version>"
  prefix_company        = "jb"
  service               = "s3"
  lob                   = "itsd"
  prefix_region = "usw2"
  application           = "recordings"
  env                   = "sandbox"
  description           = "Primary key of replica key example"
  multi_region          = true
  enable_default_policy = true
  tags                  = local.tags
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
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.82.1 |
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_aws_kms"></a> [aws\_kms](#module\_aws\_kms) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-kms | main |

## Resources

| Name | Type |
|------|------|
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the Key, will be appended with the company, lob, env and region to form a Key name. | `string` | n/a | yes |
| <a name="input_create_replica"></a> [create\_replica](#input\_create\_replica) | Determines whether a replica standard CMK will be created (AWS provided material) | `bool` | `false` | no |
| <a name="input_description"></a> [description](#input\_description) | The application name of the key , will be appended with the company, lob, env and region to form a key name. | `string` | `"KMS key for encryption"` | no |
| <a name="input_enable_default_policy"></a> [enable\_default\_policy](#input\_enable\_default\_policy) | n/a | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_key_statements"></a> [key\_statements](#input\_key\_statements) | A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage | `any` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the Key, will be appended with the company, lob, env and region to form a Key name | `string` | n/a | yes |
| <a name="input_multi_region"></a> [multi\_region](#input\_multi\_region) | Multi Region used to create KMS keys in multiple regions | `bool` | `false` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the Key, will be appended with the company, lob, env and region to form a Key name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the Key , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_primary_key_arn"></a> [primary\_key\_arn](#input\_primary\_key\_arn) | The primary key arn of a multi-region replica key | `string` | `null` | no |
| <a name="input_service"></a> [service](#input\_service) | The service identity for the Key. e.g. s3,rds,firehose, etc. | `string` | `"default"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_key_arn"></a> [key\_arn](#output\_key\_arn) | AWS kms key arn |
| <a name="output_key_id"></a> [key\_id](#output\_key\_id) | The globally unique identifier for the key |
<!-- END_TF_DOCS -->
