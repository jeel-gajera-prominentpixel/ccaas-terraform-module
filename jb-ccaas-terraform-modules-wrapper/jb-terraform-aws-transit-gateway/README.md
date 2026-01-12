# AWS transit gateway Terraform Module

## How to use this module:

### aws transit gateway basic module usage with the required input variables:
```terraform
module "transit-gateway-basic" {
  source                                = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-transit-gateway?ref=<version>"
  prefix_company                        = "jb"
  lob                                   = "itsd"
  prefix_region = "usw2"
  application                           = "recordings"
  env                                   = "sandbox"
  description                           = "Basic example TGW"
  amazon_side_asn                       = 64532
  transit_gateway_cidr_blocks           = ["10.100.0.0/16"]
  enable_auto_accept_shared_attachments = true
  vpc_attachments = {
    vpc1 = {
      vpc_id     = "vpc-02b3883c33254bc76"
      subnet_ids = ["subnet-04c1650f79a8a577c"]
    }
  }
}
```

### aws transit gateway advanced module usage with all the optional input variables:

```terraform
module "transit-gateway-advance" {
  source                                = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-transit-gateway?ref=<version>"
  prefix_company                        = "jb"
  lob                                   = "itsd"
  prefix_region = "usw2"
  application                           = "recordings"
  env                                   = "sandbox"
  description                           = "My TGW"
  amazon_side_asn                       = 64532
  transit_gateway_cidr_blocks           = ["10.100.0.0/16"]
  enable_auto_accept_shared_attachments = true
  enable_multicast_support              = false
  vpc_attachments = {
    vpc1 = {
      vpc_id     = "vpc-02b3883333253bc76"
      subnet_ids = ["subnet-04c1650f79a84577c"]

      dns_support  = true
      ipv6_support = false

      transit_gateway_default_route_table_association = false
      transit_gateway_default_route_table_propagation = false

      tgw_routes = [
        {
          destination_cidr_block = "30.0.0.0/16"
        },
        {
          blackhole              = true
          destination_cidr_block = "0.0.0.0/0"
        }
      ]
    }
  }
  ram_allow_external_principals = true
  ram_principals                = [307990089504]
  tags                          = local.tags
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
| <a name="module_transit-gateway"></a> [transit-gateway](#module\_transit-gateway) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-transit-gateway | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_amazon_side_asn"></a> [amazon\_side\_asn](#input\_amazon\_side\_asn) | The Autonomous System Number (ASN) for the Amazon side of the gateway. By default the TGW is created with the current default Amazon ASN. | `string` | `null` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name. | `string` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | Description of the Transit Gateway | `string` | `null` | no |
| <a name="input_enable_auto_accept_shared_attachments"></a> [enable\_auto\_accept\_shared\_attachments](#input\_enable\_auto\_accept\_shared\_attachments) | Whether resource attachment requests are automatically accepted | `bool` | `false` | no |
| <a name="input_enable_multicast_support"></a> [enable\_multicast\_support](#input\_enable\_multicast\_support) | Whether multicast support is enabled | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws transit gateway, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_ram_allow_external_principals"></a> [ram\_allow\_external\_principals](#input\_ram\_allow\_external\_principals) | Indicates whether principals outside your organization can be associated with a resource share. | `bool` | `false` | no |
| <a name="input_ram_principals"></a> [ram\_principals](#input\_ram\_principals) | A list of principals to share TGW with. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_transit_gateway_cidr_blocks"></a> [transit\_gateway\_cidr\_blocks](#input\_transit\_gateway\_cidr\_blocks) | One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6 | `list(string)` | `[]` | no |
| <a name="input_vpc_attachments"></a> [vpc\_attachments](#input\_vpc\_attachments) | Maps of maps of VPC details to attach to TGW. Type 'any' to disable type validation by Terraform. | `any` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ec2_transit_gateway_arn"></a> [ec2\_transit\_gateway\_arn](#output\_ec2\_transit\_gateway\_arn) | EC2 Transit Gateway Amazon Resource Name (ARN) |
| <a name="output_ec2_transit_gateway_association_default_route_table_id"></a> [ec2\_transit\_gateway\_association\_default\_route\_table\_id](#output\_ec2\_transit\_gateway\_association\_default\_route\_table\_id) | Identifier of the default association route table |
| <a name="output_ec2_transit_gateway_id"></a> [ec2\_transit\_gateway\_id](#output\_ec2\_transit\_gateway\_id) | EC2 Transit Gateway identifier |
<!-- END_TF_DOCS -->
