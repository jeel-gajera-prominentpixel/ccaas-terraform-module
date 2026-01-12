# jb-terraform-aws-lexbot

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.0, < 2.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.27 |
| <a name="requirement_awscc"></a> [awscc](#requirement\_awscc) | = 1.0.0 |
| <a name="requirement_external"></a> [external](#requirement\_external) | >= 2.3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.84.0 |
| <a name="provider_awscc"></a> [awscc](#provider\_awscc) | 1.0.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_lexv2models_bot_version.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lexv2models_bot_version) | resource |
| [awscc_lex_bot.this](https://registry.terraform.io/providers/hashicorp/awscc/1.0.0/docs/resources/lex_bot) | resource |
| [awscc_lex_bot_alias.this](https://registry.terraform.io/providers/hashicorp/awscc/1.0.0/docs/resources/lex_bot_alias) | resource |
| [awscc_lex_resource_policy.this](https://registry.terraform.io/providers/hashicorp/awscc/1.0.0/docs/resources/lex_resource_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the lambda, will be appended with the company, lob, env and region to form a lambda name. | `string` | n/a | yes |
| <a name="input_auto_build_bot_locales"></a> [auto\_build\_bot\_locales](#input\_auto\_build\_bot\_locales) | Specifies whether to automatically build the bot locales | `bool` | `true` | no |
| <a name="input_bot_alias_arn"></a> [bot\_alias\_arn](#input\_bot\_alias\_arn) | Lex bot alias ARN | `string` | `null` | no |
| <a name="input_bot_alias_locale_settings"></a> [bot\_alias\_locale\_settings](#input\_bot\_alias\_locale\_settings) | A list of bot alias locale settings to add to the bot alias. You can use this to specify different Lambda functions for different locales. | <pre>list(object({<br>    locale_id = string<br>    bot_alias_locale_setting = optional(object({<br>      enabled = optional(bool, true)<br>      code_hook_specification = optional(object({<br>        lambda_code_hook = object({<br>          code_hook_interface_version = string<br>          lambda_arn                  = string<br>        })<br>      }))<br>    }))<br>  }))</pre> | `[]` | no |
| <a name="input_bot_alias_name"></a> [bot\_alias\_name](#input\_bot\_alias\_name) | Name of the bot alias | `string` | `"LIVE"` | no |
| <a name="input_bot_file_s3_location"></a> [bot\_file\_s3\_location](#input\_bot\_file\_s3\_location) | S3 location of bot definitions zip file, if it's not defined inline in CloudFormation | <pre>object({<br>    s3_bucket         = optional(string)<br>    s3_object_key     = optional(string)<br>    s3_object_version = optional(string)<br>  })</pre> | `null` | no |
| <a name="input_bot_id"></a> [bot\_id](#input\_bot\_id) | The identifier of the bot to create the version for | `string` | `null` | no |
| <a name="input_bot_tags"></a> [bot\_tags](#input\_bot\_tags) | A list of tags to add to the bot, which can only be added at bot creation | <pre>list(object({<br>    key   = string<br>    value = string<br>  }))</pre> | `[]` | no |
| <a name="input_bot_version"></a> [bot\_version](#input\_bot\_version) | The version of the bot to use for the alias | `string` | `null` | no |
| <a name="input_conversation_log_settings"></a> [conversation\_log\_settings](#input\_conversation\_log\_settings) | Settings for conversation logging including both audio and text logs | <pre>object({<br>    # Audio Log Settings<br>    audio_log_settings = optional(list(object({<br>      enabled = optional(bool, false)<br>      destination = optional(object({<br>        s3_bucket = optional(object({<br>          kms_key_arn   = optional(string)<br>          log_prefix    = optional(string)<br>          s3_bucket_arn = optional(string)<br>        }))<br>      }))<br>    })), [])<br><br>    # Text Log Settings<br>    text_log_settings = optional(list(object({<br>      enabled = optional(bool, true)<br>      destination = optional(object({<br>        cloudwatch = optional(object({<br>          cloudwatch_log_group_arn = optional(string)<br>          log_prefix               = optional(string)<br>        }))<br>      }))<br>    })), [])<br>  })</pre> | <pre>{<br>  "audio_log_settings": [],<br>  "text_log_settings": []<br>}</pre> | no |
| <a name="input_create_alias_policy"></a> [create\_alias\_policy](#input\_create\_alias\_policy) | Create Lex bot alias policy | `bool` | `false` | no |
| <a name="input_create_lexbot"></a> [create\_lexbot](#input\_create\_lexbot) | Create Lex bot | `bool` | `false` | no |
| <a name="input_create_lexbot_alias"></a> [create\_lexbot\_alias](#input\_create\_lexbot\_alias) | Create Lex bot alias | `bool` | `false` | no |
| <a name="input_create_lexbot_version"></a> [create\_lexbot\_version](#input\_create\_lexbot\_version) | Create Lex bot version | `bool` | `false` | no |
| <a name="input_data_privacy"></a> [data\_privacy](#input\_data\_privacy) | Provides information about data privacy for the bot | <pre>object({<br>    child_directed = bool<br>  })</pre> | <pre>{<br>  "child_directed": false<br>}</pre> | no |
| <a name="input_description"></a> [description](#input\_description) | Description of the Lex bot | `string` | `null` | no |
| <a name="input_enable_monitoring"></a> [enable\_monitoring](#input\_enable\_monitoring) | Enable CloudWatch monitoring for the bot | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_idle_session_ttl_in_seconds"></a> [idle\_session\_ttl\_in\_seconds](#input\_idle\_session\_ttl\_in\_seconds) | IdleSessionTTLInSeconds of the resource | `number` | `300` | no |
| <a name="input_lex_policy"></a> [lex\_policy](#input\_lex\_policy) | IAM policy for the Lex bot | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the lambda, will be appended with the company, lob, env and region to form a lambda name | `string` | n/a | yes |
| <a name="input_locale_specification"></a> [locale\_specification](#input\_locale\_specification) | Locale specification for the bot version. Default is DRAFT version for en\_US | <pre>map(object({<br>    source_bot_version = string<br>  }))</pre> | <pre>{<br>  "en_US": {<br>    "source_bot_version": "DRAFT"<br>  }<br>}</pre> | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the Lex bot | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the lambda, will be appended with the company, lob, env and region to form a lambda name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the lambda , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | ARN of the IAM role used by the bot | `string` | `null` | no |
| <a name="input_sentiment_analysis_settings"></a> [sentiment\_analysis\_settings](#input\_sentiment\_analysis\_settings) | Settings for sentiment analysis | <pre>object({<br>    detect_sentiment = bool<br>  })</pre> | <pre>{<br>  "detect_sentiment": false<br>}</pre> | no |
| <a name="input_test_bot_alias_settings"></a> [test\_bot\_alias\_settings](#input\_test\_bot\_alias\_settings) | Settings for the test bot alias if needed | <pre>object({<br>    create_test_alias = bool<br>    description       = string<br>  })</pre> | <pre>{<br>  "create_test_alias": false,<br>  "description": "Test alias for bot"<br>}</pre> | no |
| <a name="input_version_description"></a> [version\_description](#input\_version\_description) | Description of the bot version | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_bot_alias_arn"></a> [bot\_alias\_arn](#output\_bot\_alias\_arn) | ARN of the created Lex bot alias |
| <a name="output_bot_alias_id"></a> [bot\_alias\_id](#output\_bot\_alias\_id) | ARN of the created Lex bot alias |
| <a name="output_bot_alias_name"></a> [bot\_alias\_name](#output\_bot\_alias\_name) | Name of the created Lex bot alias |
| <a name="output_bot_id"></a> [bot\_id](#output\_bot\_id) | ID of the created Lex bot |
| <a name="output_bot_name"></a> [bot\_name](#output\_bot\_name) | Name of the created Lex bot |
| <a name="output_bot_version"></a> [bot\_version](#output\_bot\_version) | Version of the created Lex bot |
<!-- END_TF_DOCS -->
