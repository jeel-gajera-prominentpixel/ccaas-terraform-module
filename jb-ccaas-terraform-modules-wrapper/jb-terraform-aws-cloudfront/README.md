# AWS CLOUD FRONT Terraform Module

## How to use this module:

### aws cloud front basic module usage with the required input variables:
```terraform
module "cloudfront_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-cloudfront?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  default_cache_behavior = {
    target_origin_id       = "something"
    viewer_protocol_policy = "allow-all"

    allowed_methods = ["GET", "HEAD", "OPTIONS"]
    cached_methods  = ["GET", "HEAD"]
    compress        = true
    query_string    = true
  }
  create_origin_access_control = true
  origin_access_control = {
    s3_oac = {
      description      = "CloudFront access to S3"
      origin_type      = "s3"
      signing_behavior = "always"
      signing_protocol = "sigv4"
    }
  }
  origin = {
    appsync = {
      domain_name = "appsync.xyz.com"
      custom_origin_config = {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "match-viewer"
        origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
      }

      custom_header = [
        {
          name  = "X-Forwarded-Scheme"
          value = "https"
        },
        {
          name  = "X-Frame-Options"
          value = "SAMEORIGIN"
        }
      ]

      origin_shield = {
        enabled              = true
        origin_shield_region = "us-east-1"
      }
    }

    s3_one = {
      domain_name = "examplebucket.s3-us-west-2.amazonaws.com"
      s3_origin_config = {
        origin_access_identity = "s3_bucket_one"

      }
    }

    s3_oac = {
      domain_name           = "examplebucket.s3-us-west-2.amazonaws.com"
      origin_access_control = "s3_oac"

    }
  }
  origin_group = {
    group_one = {
      failover_status_codes      = [403, 404, 500, 502]
      primary_member_origin_id   = "appsync"
      secondary_member_origin_id = "s3_one"
    }
  }
  tags = local.tags
}
```

### aws cloud front advanced module usage with all the optional input variables:
```terraform
module "cloudfront_advanced" {
  source                        = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-cloudfront?ref=<version>"
  prefix_company                = "jb"
  lob                           = "itsd"
  prefix_region = "usw2"
  application                   = "recordings"
  env                           = "sandbox"
  comment                       = "My awesome CloudFront"
  enabled                       = true
  staging                       = false
  http_version                  = "http2and3"
  is_ipv6_enabled               = true
  price_class                   = "PriceClass_All"
  retain_on_delete              = false
  wait_for_deployment           = false
  create_origin_access_identity = true
  origin_access_identities = {
    s3_bucket_one = "My awesome CloudFront can access"
  }
  default_cache_behavior = {
    target_origin_id       = "something"
    viewer_protocol_policy = "allow-all"

    allowed_methods = ["GET", "HEAD", "OPTIONS"]
    cached_methods  = ["GET", "HEAD"]
    compress        = true
    query_string    = true
  }

  create_origin_access_control = true
  origin_access_control = {
    s3_oac = {
      description      = "CloudFront access to S3"
      origin_type      = "s3"
      signing_behavior = "always"
      signing_protocol = "sigv4"
    }
  }
  logging_config = {
    bucket = ""
    prefix = "cloudfront"
  }
  origin = {
    appsync = {
      domain_name = "appsync.xyz.com"
      custom_origin_config = {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "match-viewer"
        origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
      }

      custom_header = [
        {
          name  = "X-Forwarded-Scheme"
          value = "https"
        },
        {
          name  = "X-Frame-Options"
          value = "SAMEORIGIN"
        }
      ]

      origin_shield = {
        enabled              = true
        origin_shield_region = "us-east-1"
      }
    }

    s3_one = {
      domain_name = "examplebucket.s3-us-west-2.amazonaws.com"
      s3_origin_config = {
        origin_access_identity = "s3_bucket_one"

      }
    }

    s3_oac = {
      domain_name           = "examplebucket.s3-us-west-2.amazonaws.com"
      origin_access_control = "s3_oac"

    }
  }
  origin_group = {
    group_one = {
      failover_status_codes      = [403, 404, 500, 502]
      primary_member_origin_id   = "appsync"
      secondary_member_origin_id = "s3_one"
    }
  }
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
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.73.0 |
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_cloudfront"></a> [cloudfront](#module\_cloudfront) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-cloudfront | main |

## Resources

| Name | Type |
|------|------|
| [aws_cloudfront_origin_request_policy.custom](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_origin_request_policy) | resource |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aliases"></a> [aliases](#input\_aliases) | Extra CNAMEs (alternate domain names), if any, for this distribution. | `list(string)` | `null` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_comment"></a> [comment](#input\_comment) | Any comments you want to include about the distribution. | `string` | `null` | no |
| <a name="input_continuous_deployment_policy_id"></a> [continuous\_deployment\_policy\_id](#input\_continuous\_deployment\_policy\_id) | Identifier of a continuous deployment policy. This argument should only be set on a production distribution. | `string` | `null` | no |
| <a name="input_create_distribution"></a> [create\_distribution](#input\_create\_distribution) | Controls if CloudFront distribution should be created | `bool` | `true` | no |
| <a name="input_create_monitoring_subscription"></a> [create\_monitoring\_subscription](#input\_create\_monitoring\_subscription) | If enabled, the resource for monitoring subscription will created. | `bool` | `false` | no |
| <a name="input_create_origin_access_control"></a> [create\_origin\_access\_control](#input\_create\_origin\_access\_control) | Controls if CloudFront origin access control should be created | `bool` | `false` | no |
| <a name="input_create_origin_access_identity"></a> [create\_origin\_access\_identity](#input\_create\_origin\_access\_identity) | Controls if CloudFront origin access identity should be created | `bool` | `false` | no |
| <a name="input_create_request_policy"></a> [create\_request\_policy](#input\_create\_request\_policy) | Flag to control whether to create a custom origin request policy | `bool` | `false` | no |
| <a name="input_custom_error_response"></a> [custom\_error\_response](#input\_custom\_error\_response) | One or more custom error response elements | `any` | `{}` | no |
| <a name="input_custom_origin_request_policy"></a> [custom\_origin\_request\_policy](#input\_custom\_origin\_request\_policy) | n/a | <pre>object({<br>    name    = string<br>    comment = string<br><br>    cookies_config = object({<br>      cookie_behavior = string<br>      cookies         = optional(list(string), [])<br>    })<br><br>    headers_config = object({<br>      header_behavior = string<br>      headers         = optional(list(string), [])<br>    })<br><br>    query_strings_config = object({<br>      query_string_behavior = string<br>      query_strings         = optional(list(string), [])<br>    })<br>  })</pre> | `null` | no |
| <a name="input_default_cache_behavior"></a> [default\_cache\_behavior](#input\_default\_cache\_behavior) | The default cache behavior for this distribution | `any` | `null` | no |
| <a name="input_default_root_object"></a> [default\_root\_object](#input\_default\_root\_object) | The object that you want CloudFront to return (for example, index.html) when an end user requests the root URL. | `string` | `null` | no |
| <a name="input_enabled"></a> [enabled](#input\_enabled) | Whether the distribution is enabled to accept end user requests for content. | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_geo_restriction"></a> [geo\_restriction](#input\_geo\_restriction) | The restriction configuration for this distribution (geo\_restrictions) | `any` | `{}` | no |
| <a name="input_http_version"></a> [http\_version](#input\_http\_version) | The maximum HTTP version to support on the distribution. Allowed values are http1.1, http2, http2and3, and http3. The default is http2. | `string` | `"http2"` | no |
| <a name="input_is_ipv6_enabled"></a> [is\_ipv6\_enabled](#input\_is\_ipv6\_enabled) | Whether the IPv6 is enabled for the distribution. | `bool` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_logging_config"></a> [logging\_config](#input\_logging\_config) | The logging configuration that controls how logs are written to your distribution (maximum one). | `any` | `{}` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_ordered_cache_behavior"></a> [ordered\_cache\_behavior](#input\_ordered\_cache\_behavior) | An ordered list of cache behaviors resource for this distribution. List from top to bottom in order of precedence. The topmost cache behavior will have precedence 0. | `any` | `[]` | no |
| <a name="input_origin"></a> [origin](#input\_origin) | One or more origins for this distribution (multiples allowed). | `any` | `null` | no |
| <a name="input_origin_access_control"></a> [origin\_access\_control](#input\_origin\_access\_control) | Map of CloudFront origin access control | <pre>map(object({<br>    description      = string<br>    origin_type      = string<br>    signing_behavior = string<br>    signing_protocol = string<br>  }))</pre> | <pre>{<br>  "s3": {<br>    "description": "",<br>    "origin_type": "s3",<br>    "signing_behavior": "always",<br>    "signing_protocol": "sigv4"<br>  }<br>}</pre> | no |
| <a name="input_origin_access_identities"></a> [origin\_access\_identities](#input\_origin\_access\_identities) | Map of CloudFront origin access identities (value as a comment) | `map(string)` | `{}` | no |
| <a name="input_origin_group"></a> [origin\_group](#input\_origin\_group) | One or more origin\_group for this distribution (multiples allowed). | `any` | `{}` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the cloudfront, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_price_class"></a> [price\_class](#input\_price\_class) | The price class for this distribution. One of PriceClass\_All, PriceClass\_200, PriceClass\_100 | `string` | `null` | no |
| <a name="input_realtime_metrics_subscription_status"></a> [realtime\_metrics\_subscription\_status](#input\_realtime\_metrics\_subscription\_status) | A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are `Enabled` and `Disabled`. | `string` | `"Enabled"` | no |
| <a name="input_retain_on_delete"></a> [retain\_on\_delete](#input\_retain\_on\_delete) | Disables the distribution instead of deleting it when destroying the resource through Terraform. If this is set, the distribution needs to be deleted manually afterwards. | `bool` | `false` | no |
| <a name="input_staging"></a> [staging](#input\_staging) | Whether the distribution is a staging distribution. | `bool` | `false` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_viewer_certificate"></a> [viewer\_certificate](#input\_viewer\_certificate) | The SSL configuration for this distribution | `any` | <pre>{<br>  "cloudfront_default_certificate": true,<br>  "minimum_protocol_version": "TLSv1"<br>}</pre> | no |
| <a name="input_wait_for_deployment"></a> [wait\_for\_deployment](#input\_wait\_for\_deployment) | If enabled, the resource will wait for the distribution status to change from InProgress to Deployed. Setting this to false will skip the process. | `bool` | `true` | no |
| <a name="input_web_acl_id"></a> [web\_acl\_id](#input\_web\_acl\_id) | If you're using AWS WAF to filter CloudFront requests, the Id of the AWS WAF web ACL that is associated with the distribution. The WAF Web ACL must exist in the WAF Global (CloudFront) region and the credentials configuring this argument must have waf:GetWebACL permissions assigned. If using WAFv2, provide the ARN of the web ACL. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_cloudfront_distribution_arn"></a> [cloudfront\_distribution\_arn](#output\_cloudfront\_distribution\_arn) | The ARN (Amazon Resource Name) for the distribution. |
| <a name="output_cloudfront_distribution_domain_name"></a> [cloudfront\_distribution\_domain\_name](#output\_cloudfront\_distribution\_domain\_name) | The domain name corresponding to the distribution. |
| <a name="output_cloudfront_distribution_id"></a> [cloudfront\_distribution\_id](#output\_cloudfront\_distribution\_id) | The identifier for the distribution. |
| <a name="output_cloudfront_origin_access_identity_iam_arns"></a> [cloudfront\_origin\_access\_identity\_iam\_arns](#output\_cloudfront\_origin\_access\_identity\_iam\_arns) | The IAM arns of the origin access identities created |
| <a name="output_cloudfront_origin_access_identity_ids"></a> [cloudfront\_origin\_access\_identity\_ids](#output\_cloudfront\_origin\_access\_identity\_ids) | The IDS of the origin access identities created |
| <a name="output_custom_origin_request_policy_id"></a> [custom\_origin\_request\_policy\_id](#output\_custom\_origin\_request\_policy\_id) | n/a |
<!-- END_TF_DOCS -->
