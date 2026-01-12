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
| <a name="module_wafv2"></a> [wafv2](#module\_wafv2) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-wafv2 | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the wafv2, will be appended with the company, lob, env and region to form a wafv2 name. | `string` | n/a | yes |
| <a name="input_default_action"></a> [default\_action](#input\_default\_action) | (Required) Action to perform if none of the rules contained in the WebACL match. | `string` | n/a | yes |
| <a name="input_description"></a> [description](#input\_description) | (Optional) Friendly description of the WebACL. | `string` | `null` | no |
| <a name="input_enabled_logging_configuration"></a> [enabled\_logging\_configuration](#input\_enabled\_logging\_configuration) | (Optional) Whether to create logging configuration. | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the wafv2, will be appended with the company, lob, env and region to form a wafv2 name. | `string` | n/a | yes |
| <a name="input_logging_filter"></a> [logging\_filter](#input\_logging\_filter) | (Optional) A configuration block that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. | `any` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | (Required) Friendly name of the WebACL. | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the wafv2, will be appended with the company, lob, env and region to form a wafv2 name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws wafv2 , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_redacted_fields"></a> [redacted\_fields](#input\_redacted\_fields) | The parts of the request that you want to keep out of the logs.<br>You can only specify one of the following: `method`, `query_string`, `single_header`, or `uri_path`<br><br>method:<br>  Whether to enable redaction of the HTTP method.<br>  The method indicates the type of operation that the request is asking the origin to perform.<br>uri\_path:<br>  Whether to enable redaction of the URI path.<br>  This is the part of a web request that identifies a resource.<br>query\_string:<br>  Whether to enable redaction of the query string.<br>  This is the part of a URL that appears after a `?` character, if any.<br>single\_header:<br>  The list of names of the query headers to redact. | <pre>map(object({<br>    method        = optional(bool, false)<br>    uri_path      = optional(bool, false)<br>    query_string  = optional(bool, false)<br>    single_header = optional(list(string), null)<br>  }))</pre> | `{}` | no |
| <a name="input_scope"></a> [scope](#input\_scope) | (Required) Specifies whether this is for an AWS CloudFront distribution or for a regional application | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | (Optional) Map of key-value pairs to associate with the resource. | `map(string)` | `null` | no |
| <a name="input_visibility_config"></a> [visibility\_config](#input\_visibility\_config) | Configuration for visibility settings of WAF | <pre>object({<br>    cloudwatch_metrics_enabled = bool<br>    metric_name                = string<br>    sampled_requests_enabled   = bool<br>  })</pre> | n/a | yes |
| <a name="input_waf_ip_sets"></a> [waf\_ip\_sets](#input\_waf\_ip\_sets) | List of IP sets to be used in WAF rules | <pre>list(object({<br>    name               = string<br>    ip_address_version = string       # "IPV4" or "IPV6"<br>    addresses_list     = list(string) # List of IP addresses or CIDRs<br>  }))</pre> | n/a | yes |
| <a name="input_waf_rules"></a> [waf\_rules](#input\_waf\_rules) | List of WAF rules | <pre>list(object({<br>    name                       = string<br>    priority                   = number<br>    sampled_requests_enabled   = bool<br>    cloudwatch_metrics_enabled = bool<br>    action                     = string # "allow" or "block"<br>    ip_set_name                = string # Name of the IP set to associate with the rule<br>    metric_name                = string<br>  }))</pre> | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_wafv2_arn"></a> [aws\_wafv2\_arn](#output\_aws\_wafv2\_arn) | The ARN of the WAF WebACL. |
| <a name="output_aws_wafv2_capacity"></a> [aws\_wafv2\_capacity](#output\_aws\_wafv2\_capacity) | Web ACL capacity units (WCUs) currently being used by this web ACL. |
| <a name="output_aws_wafv2_id"></a> [aws\_wafv2\_id](#output\_aws\_wafv2\_id) | The ID of the WAF WebACL. |
| <a name="output_aws_wafv2_tags_all"></a> [aws\_wafv2\_tags\_all](#output\_aws\_wafv2\_tags\_all) | Map of tags assigned to the resource, including those inherited from the provider default\_tags configuration block. |
| <a name="output_aws_wafv2_web_acl_logging_configuration_id"></a> [aws\_wafv2\_web\_acl\_logging\_configuration\_id](#output\_aws\_wafv2\_web\_acl\_logging\_configuration\_id) | The Amazon Resource Name (ARN) of the WAFv2 Web ACL. |
<!-- END_TF_DOCS -->
