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
| <a name="module_apigv1"></a> [apigv1](#module\_apigv1) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-apigwv1 | main |

## Resources

| Name | Type |
|------|------|
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the API v1, will be appended with the company, lob, env and region to form a API v1 name. | `string` | n/a | yes |
| <a name="input_authorization"></a> [authorization](#input\_authorization) | Authorization type for the API Gateway methods | `string` | `"NONE"` | no |
| <a name="input_description"></a> [description](#input\_description) | Description of the API Gateway REST API | `string` | n/a | yes |
| <a name="input_enable_waf_association"></a> [enable\_waf\_association](#input\_enable\_waf\_association) | Flag to enable or disable WAF association with resources | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the API v1, will be appended with the company, lob, env and region to form a API v1 name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name of the API Gateway REST API | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the API v1, will be appended with the company, lob, env and region to form a API v1 name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws API v1 , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_resource_paths"></a> [resource\_paths](#input\_resource\_paths) | Map of paths for the API Gateway resources with Lambda ARNs and HTTP methods | <pre>map(object({<br>    lambda_arn              = string<br>    http_method             = string<br>    integration_http_method = string<br>    type                    = string<br>  }))</pre> | <pre>{<br>  "/proxy": {<br>    "http_method": "POST",<br>    "integration_http_method": "POST",<br>    "lambda_arn": "arn:aws:lambda:us-west-2:767252029631:function:Test-lambda-deployment",<br>    "type": "AWS_PROXY"<br>  },<br>  "/proxy2": {<br>    "http_method": "POST",<br>    "integration_http_method": "POST",<br>    "lambda_arn": "arn:aws:lambda:us-west-2:767252029631:function:Test-lambda-deployment",<br>    "type": "AWS_PROXY"<br>  }<br>}</pre> | no |
| <a name="input_resource_root_path"></a> [resource\_root\_path](#input\_resource\_root\_path) | List of paths for the API Gateway resources | `string` | `"ANY"` | no |
| <a name="input_root_integration_http_method"></a> [root\_integration\_http\_method](#input\_root\_integration\_http\_method) | root\_integration\_http\_method for integration of lambda | `string` | `"arn:aws:lambda:region:account-id:function:root-function"` | no |
| <a name="input_root_integration_type"></a> [root\_integration\_type](#input\_root\_integration\_type) | root\_integration\_type for integration of lambda | `string` | `"arn:aws:lambda:region:account-id:function:root-function"` | no |
| <a name="input_root_lambda_arn"></a> [root\_lambda\_arn](#input\_root\_lambda\_arn) | Lambda function ARN for the root method | `string` | `"arn:aws:lambda:region:account-id:function:root-function"` | no |
| <a name="input_stage_name"></a> [stage\_name](#input\_stage\_name) | Stage Name of the API Gateway | `string` | `"dev"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A mapping of tags to assign to API gateway resources. | `map(string)` | `{}` | no |
| <a name="input_types"></a> [types](#input\_types) | Type of the API Gateway Endpoint Available values are EDGE, REGIONAL and PRIVATE. Default is REGIONAL | `string` | `"REGIONAL"` | no |
| <a name="input_web_acl_arn"></a> [web\_acl\_arn](#input\_web\_acl\_arn) | WAF ARNs to associate with the resource | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_apigatewayv1_api_arn"></a> [apigatewayv1\_api\_arn](#output\_apigatewayv1\_api\_arn) | The ARN of the API |
| <a name="output_apigatewayv1_api_execution_arn"></a> [apigatewayv1\_api\_execution\_arn](#output\_apigatewayv1\_api\_execution\_arn) | The ARN prefix to be used in an aws\_lambda\_permission's source\_arn attribute or in an aws\_iam\_policy to authorize access to the @connections API. |
| <a name="output_apigatewayv1_api_id"></a> [apigatewayv1\_api\_id](#output\_apigatewayv1\_api\_id) | The API identifier |
| <a name="output_apigatewayv1_binary_media_types"></a> [apigatewayv1\_binary\_media\_types](#output\_apigatewayv1\_binary\_media\_types) | List of binary media types supported by the REST API. |
| <a name="output_apigatewayv1_description"></a> [apigatewayv1\_description](#output\_apigatewayv1\_description) | Description of the REST API. |
| <a name="output_apigatewayv1_endpoint_configuration"></a> [apigatewayv1\_endpoint\_configuration](#output\_apigatewayv1\_endpoint\_configuration) | The endpoint configuration of this RestApi showing the endpoint types of the API. |
| <a name="output_apigatewayv1_key_source"></a> [apigatewayv1\_key\_source](#output\_apigatewayv1\_key\_source) | Source of the API key for requests. |
| <a name="output_apigatewayv1_minimum_compression_size"></a> [apigatewayv1\_minimum\_compression\_size](#output\_apigatewayv1\_minimum\_compression\_size) | Minimum response size to compress for the REST API. |
| <a name="output_apigatewayv1_policy"></a> [apigatewayv1\_policy](#output\_apigatewayv1\_policy) | JSON formatted policy document that controls access to the API Gateway. |
| <a name="output_apigatewayv1_root_resource_id"></a> [apigatewayv1\_root\_resource\_id](#output\_apigatewayv1\_root\_resource\_id) | Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'. |
| <a name="output_apigatewayv1_stage_arn"></a> [apigatewayv1\_stage\_arn](#output\_apigatewayv1\_stage\_arn) | Stage ARN |
| <a name="output_apigatewayv1_stage_execution_arn"></a> [apigatewayv1\_stage\_execution\_arn](#output\_apigatewayv1\_stage\_execution\_arn) | Execution ARN to be used in lambda\_permission's source\_arn when allowing API Gateway to invoke a Lambda function, e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod |
| <a name="output_apigatewayv1_stage_id"></a> [apigatewayv1\_stage\_id](#output\_apigatewayv1\_stage\_id) | Stage ID |
| <a name="output_apigatewayv1_stage_invoke_url"></a> [apigatewayv1\_stage\_invoke\_url](#output\_apigatewayv1\_stage\_invoke\_url) | URL to invoke the API pointing to the stage, e.g., https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod |
| <a name="output_apigatewayv1_stage_web_acl_arn"></a> [apigatewayv1\_stage\_web\_acl\_arn](#output\_apigatewayv1\_stage\_web\_acl\_arn) | ARN of the WebAcl associated with the Stage. |
<!-- END_TF_DOCS -->
