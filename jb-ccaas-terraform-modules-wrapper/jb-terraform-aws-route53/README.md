# AWS route53 Terraform Module

## How to use this module:

### aws route53 basic module usage with the required input variables:
```terraform
module "route_53_records" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-route53/modules/records/?ref=<version>"
  prefix_company = "jb"
  prefix_region  = "ap-south-1"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
}
```

### aws route53 advanced module usage with all the optional input variables:



```terraform
module "route_53_records" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-route53/modules/records/?ref=<version>"
  prefix_company = "jb"
  prefix_region  = "ap-south-1"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  zone_name      = "example.com"
  zone_id        = "ABCDEFGHIJKLMN"
  private_zone   = false
  records = [
    {
      name           = "geo"
      type           = "CNAME"
      ttl            = 5
      records        = ["europe.test.example.com."]
      set_identifier = "europe"
      geolocation_routing_policy = {
        continent = "EU"
      }
    }
  ]
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
| <a name="module_route_53_records"></a> [route\_53\_records](#module\_route\_53\_records) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-route53/modules/records/ | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the route 53, will be appended with the company, lob, env and region to form a route 53 name. | `string` | n/a | yes |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the route 53, will be appended with the company, lob, env and region to form a route 53 name. | `string` | n/a | yes |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the route 53, will be appended with the company, lob, env and region to form a route 53 name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the route 53 , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_private_zone"></a> [private\_zone](#input\_private\_zone) | Whether Route53 zone is private or public | `bool` | `false` | no |
| <a name="input_records"></a> [records](#input\_records) | List of objects of DNS records | `any` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | ID of DNS zone | `string` | `null` | no |
| <a name="input_zone_name"></a> [zone\_name](#input\_zone\_name) | Name of DNS zone | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_route53_record_name"></a> [route53\_record\_name](#output\_route53\_record\_name) | The name of the record |
<!-- END_TF_DOCS -->
