# AWS pinpoint Terraform Module

## How to use this module:

### aws pinpoint basic module usage with the required input variables:
```terraform
module "pinpoint_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-amazon-pinpoint?ref=<version>"
  name           = var.name == null ? local.pinpoint_name : var.name
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
}
```

### aws pinpoint advanced module usage with all the optional input variables:



```terraform
module "pinpoint_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-amazon-pinpoint?ref=<version>"
  name           = "jb-pinpoint"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  email = {
    from     = "example@example.com"
    identity = "arn:aws:ses:us-west-2:123456789012:identity/example.com"
  }
  sms = {
    sender     = "example_sender"
    short_code = "12345"
  }
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
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.3 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_pinpoint"></a> [pinpoint](#module\_pinpoint) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-pinpoint | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name. | `string` | n/a | yes |
| <a name="input_email"></a> [email](#input\_email) | Provides a Pinpoint Email Channel resource. | <pre>object({<br>    from     = string<br>    identity = string<br>  })</pre> | `null` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws pinoint , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_sms"></a> [sms](#input\_sms) | Provides a Pinpoint SMS Channel resource. | <pre>object({<br>    sender     = string<br>    short_code = string<br>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_application_arn"></a> [application\_arn](#output\_application\_arn) | Amazon Resource Name (ARN) of the PinPoint Application. |
| <a name="output_application_id"></a> [application\_id](#output\_application\_id) | The Application ID of the Pinpoint App. |
<!-- END_TF_DOCS -->
