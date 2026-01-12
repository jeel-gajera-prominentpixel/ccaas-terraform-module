# AWS acm Terraform Module

## How to use this module:

### aws acm basic module usage with the required input variables:
```terraform
module "acm" {
  source         = git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-acm?ref=<version>
  prefix_company = "jb"
  lob            = "itsd"
  application    = "recordings"
  prefix_region  = "usw2"
  env            = "sandbox"
  domain_name    = "example.com"
  zone_id        = "Z2ES7B9AZ6SHAE"
  subject_alternative_names = [
    "*.my-domain.com",
    "app.sub.my-domain.com",
  ]
  validation_method       = "DNS"
  validation_record_fqdns = ["example.com", "www.example.com"]
}
```

### aws acm advanced module usage with all the optional input variables:



```terraform
module "acm" {
  source         = git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-acm?ref=<version>
  prefix_company = "jb"
  lob            = "itsd"
  application    = "recordings"
  prefix_region  = "usw2"
  env            = "sandbox"
  domain_name    = example.com
  zone_id        = "Z2ES7B9AZ6SHAE"
  subject_alternative_names = [
    "*.my-domain.com",
    "app.sub.my-domain.com",
  ]
  validation_method       = "DNS"
  create_route53_records  = false
  validation_record_fqdns = ["example.com", "www.example.com"]
  tags                    = local.tags
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
| <a name="module_acm"></a> [acm](#module\_acm) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-acm | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_create_route53_records"></a> [create\_route53\_records](#input\_create\_route53\_records) | When validation is set to DNS, define whether to create the DNS records internally via Route53 or externally using any DNS provider | `bool` | `true` | no |
| <a name="input_domain_name"></a> [domain\_name](#input\_domain\_name) | A domain name for which the certificate should be issued | `string` | `""` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_subject_alternative_names"></a> [subject\_alternative\_names](#input\_subject\_alternative\_names) | A list of domains that should be SANs in the issued certificate | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_validation_method"></a> [validation\_method](#input\_validation\_method) | Which method to use for validation. DNS or EMAIL are valid. This parameter must not be set for certificates that were imported into ACM and then into Terraform. | `string` | `null` | no |
| <a name="input_validation_record_fqdns"></a> [validation\_record\_fqdns](#input\_validation\_record\_fqdns) | When validation is set to DNS and the DNS validation records are set externally, provide the fqdns for the validation | `list(string)` | `[]` | no |
| <a name="input_wait_for_validation"></a> [wait\_for\_validation](#input\_wait\_for\_validation) | Whether to wait for the validation to complete | `bool` | `true` | no |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | The ID of the hosted zone to contain this record. Required when validating via Route53 | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_acm_certificate_arn"></a> [acm\_certificate\_arn](#output\_acm\_certificate\_arn) | The ARN of the certificate |
| <a name="output_acm_certificate_domain_validation_options"></a> [acm\_certificate\_domain\_validation\_options](#output\_acm\_certificate\_domain\_validation\_options) | A list of attributes to feed into other resources to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if DNS-validation was used. |
| <a name="output_acm_certificate_status"></a> [acm\_certificate\_status](#output\_acm\_certificate\_status) | Status of the certificate. |
| <a name="output_validation_domains"></a> [validation\_domains](#output\_validation\_domains) | List of distinct domain validation options. This is useful if subject alternative names contain wildcards. |
<!-- END_TF_DOCS -->
