# AWS API GATEWAY Terraform Module

## How to use this module:

### aws api-gateway basic module usage with the required input variables:
```terraform
module "api-gateway_basic" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-apigateway?ref=<version>"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  tags = local.tags
}
```

### aws api-gateway advanced module usage with all the optional input variables:
```terraform
module "api-gateway_advance" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-apigateway?ref=<version>"
  prefix_company = "jb"
  prefix_region = "usw2"
  protocol_type  = "HTTP"
  lob            = "itsd"
  application    = "recordings"
  env            = "sandbox"

  # vpc-link-http
  create_api_domain_name = false
  vpc_links              = {}

  default_stage_access_log_destination_arn = null
  default_stage_access_log_format          = null
  domain_name                              = "example.com"
  authorizers                              = {}
  domain_name_certificate_arn              = "arn:aws:acm:eu-south-1:123456789102:certificate/12345678-1234-1234-1234-123456789012"
  cors_configuration                       = {}
  integrations                             = {}
  tags                                     = local.tags
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
| <a name="module_api-gateway"></a> [api-gateway](#module\_api-gateway) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-apigateway-v2 | main |

## Resources

| Name | Type |
|------|------|
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the api gateway, will be appended with the company, lob, env and region to form a api gateway name. | `string` | n/a | yes |
| <a name="input_authorizers"></a> [authorizers](#input\_authorizers) | Map of API gateway routes with integrations<br><br>  Example/available options:<pre>{<br>    "azure" = {<br>      authorizer_type  = "JWT"<br>      identity_sources = "$request.header.Authorization"<br>      name             = "azure-auth"<br>      audience         = ["d6a38afd-45d6-4874-d1aa-3c5c558aqcc2"]<br>      issuer           = "https://sts.windows.net/aaee026e-8f37-410e-8869-72d9154873e4/"<br>    }<br>  }</pre> | `map(any)` | `{}` | no |
| <a name="input_cors_configuration"></a> [cors\_configuration](#input\_cors\_configuration) | Map of API gateway routes with integrations<br><br>  Example/available options:<pre>{<br>    allow_credentials = true<br>    allow_methods     = ["GET", "OPTIONS", "POST"] or ["*"]<br>    max_age           = 5<br>    allow_headers = ["content-type", "x-amz-date", "authorization", "x-api-key", "x-amz-security-token", "x-amz-user-agent"]<br>    allow_origins = ["*"]<br>  }</pre> | `any` | `{}` | no |
| <a name="input_create_api_domain_name"></a> [create\_api\_domain\_name](#input\_create\_api\_domain\_name) | Whether to create API domain name resource | `bool` | `true` | no |
| <a name="input_create_routes_and_integrations"></a> [create\_routes\_and\_integrations](#input\_create\_routes\_and\_integrations) | Whether to create routes and integrations resources | `bool` | `true` | no |
| <a name="input_default_stage_access_log_destination_arn"></a> [default\_stage\_access\_log\_destination\_arn](#input\_default\_stage\_access\_log\_destination\_arn) | Default stage's ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN. | `string` | `null` | no |
| <a name="input_default_stage_access_log_format"></a> [default\_stage\_access\_log\_format](#input\_default\_stage\_access\_log\_format) | Default stage's single line format of the access logs of data, as specified by selected $context variables. | `string` | `null` | no |
| <a name="input_domain_name"></a> [domain\_name](#input\_domain\_name) | The domain name to use for API gateway | `string` | `null` | no |
| <a name="input_domain_name_certificate_arn"></a> [domain\_name\_certificate\_arn](#input\_domain\_name\_certificate\_arn) | The ARN of an AWS-managed certificate that will be used by the endpoint for the domain name | `string` | `""` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_integrations"></a> [integrations](#input\_integrations) | Map of API gateway routes with integrations<br><br>  Example/available options:<pre>{<br>      "ANY /" = {<br>          lambda_arn             = module.lambda.arn<br>          payload_format_version = "2.0"<br>          timeout_milliseconds   = 12000<br>        }<br>        "GET /some-route-with-authorizer" = {<br>          lambda_arn             = module.lambda.arn<br>          payload_format_version = "2.0"<br>          authorizer_key         = "cognito"<br>        }<br>        "POST /start-step-function" = {<br>          lambda_arn             = module.lambda.arn<br>          payload_format_version = "2.0"<br>          authorizer_key         = "cognito"<br>        }<br><br>      "<br>    }<br>  }</pre> | `map(any)` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the api gateway, will be appended with the company, lob, env and region to form a api gateway name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the api gateway, will be appended with the company, lob, env and region to form a api gateway name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_protocol_type"></a> [protocol\_type](#input\_protocol\_type) | The API protocol. Valid values: HTTP, WEBSOCKET | `string` | `"HTTP"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_vpc_links"></a> [vpc\_links](#input\_vpc\_links) | Map of API gateway routes with integrations<br><br>  Example/available options:<pre>{<br>    my-vpc = {<br>      name               = "example"<br>      security_group_ids = [security_group_id]<br>      subnet_ids         = public_subnets_id<br>    }<br>  }</pre> | `map(any)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_apigatewayv2_api_api_endpoint"></a> [apigatewayv2\_api\_api\_endpoint](#output\_apigatewayv2\_api\_api\_endpoint) | The API identifier. |
| <a name="output_apigatewayv2_api_arn"></a> [apigatewayv2\_api\_arn](#output\_apigatewayv2\_api\_arn) | The ARN of the API |
| <a name="output_apigatewayv2_api_execution_arn"></a> [apigatewayv2\_api\_execution\_arn](#output\_apigatewayv2\_api\_execution\_arn) | The ARN of the API |
| <a name="output_apigatewayv2_api_id"></a> [apigatewayv2\_api\_id](#output\_apigatewayv2\_api\_id) | The ARN of the API |
| <a name="output_apigatewayv2_domain_name_arn"></a> [apigatewayv2\_domain\_name\_arn](#output\_apigatewayv2\_domain\_name\_arn) | The ARN of the API |
| <a name="output_apigatewayv2_domain_name_id"></a> [apigatewayv2\_domain\_name\_id](#output\_apigatewayv2\_domain\_name\_id) | The ARN of the API |
| <a name="output_apigatewayv2_domain_name_target_domain_name"></a> [apigatewayv2\_domain\_name\_target\_domain\_name](#output\_apigatewayv2\_domain\_name\_target\_domain\_name) | The ARN of the API |
| <a name="output_default_apigatewayv2_stage_domain_name"></a> [default\_apigatewayv2\_stage\_domain\_name](#output\_default\_apigatewayv2\_stage\_domain\_name) | The ARN of the API |
| <a name="output_default_apigatewayv2_stage_invoke_url"></a> [default\_apigatewayv2\_stage\_invoke\_url](#output\_default\_apigatewayv2\_stage\_invoke\_url) | The URL to invoke the API pointing to the stage |
<!-- END_TF_DOCS -->
