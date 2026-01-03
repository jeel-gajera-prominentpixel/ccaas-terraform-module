# AWS connect advance Terraform Module

## How to use this module:

### aws lambda function basic module usage with the required input variables:
```terraform
module "amazon_connect_basic" {
  source                     = "git@github.com:CloverHealth/ccaas-terraform-modules-wrapper.git//terraform-aws-amazon-connect?ref=<version>"
  prefix_company             = "ch"
  lob                        = ""
  prefix_region = "usw2"
  prefix_region = var.prefix_region
  application                = "recordings"
  env                        = "sandbox"
  tags = local.tags
}

```

### aws lambda function advanced module usage with all the optional input variables:
```terraform
module "amazon_connect_advance" {
  source                     = "git@github.com:CloverHealth/ccaas-terraform-modules-wrapper.git//terraform-aws-amazon-connect?ref=<version>"
  prefix_company             = "ch"
  lob                        = ""
  prefix_region = "usw2"
  application                = "recordings"
  env                        = "sandbox"
  bucket_name = "test-bucket"
  crt_firehose_arn = "arn:aws:firehose:region:account-id:deliverystream/delivery-stream-name"
  contact_flows = {}
  quick_connects = {}
  lambda_function_associations = {}
  contact_flow_modules = {}
  routing_profiles  = {}
  security_profiles = {}
  users = {}
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
| <a name="provider_awscc"></a> [awscc](#provider\_awscc) | 1.43.0 |
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_amazon_connect"></a> [amazon\_connect](#module\_amazon\_connect) | git@github.com:CloverHealth/ccaas-terraform-modules.git//terraform-aws-amazonconnect | v1.0.2 |
| <a name="module_connect_azure_ad_policy"></a> [connect\_azure\_ad\_policy](#module\_connect\_azure\_ad\_policy) | git@github.com:CloverHealth/ccaas-terraform-modules.git//terraform-aws-iam/modules/iam-policy | v1.0.2 |
| <a name="module_connect_federation_policy"></a> [connect\_federation\_policy](#module\_connect\_federation\_policy) | git@github.com:CloverHealth/ccaas-terraform-modules.git//terraform-aws-iam/modules/iam-policy | v1.0.2 |

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.AWSSSO_connect_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.AWSConnectPolicyAttachment1](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_role_policy_attachment.AWSConnectPolicyAttachment2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_saml_provider.AWSSSO_admin_Connect](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_saml_provider) | resource |
| [aws_iam_saml_provider.AWSSSO_agent_Connect](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_saml_provider) | resource |
| [awscc_connect_approved_origin.this](https://registry.terraform.io/providers/hashicorp/awscc/latest/docs/resources/connect_approved_origin) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_iam_policy_document.AWSConnectTrustPolicy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_iam_policy_document.connect_azure_ad_policy_document](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_iam_policy_document.connect_federation_policy_document](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_account_number"></a> [account\_number](#input\_account\_number) | n/a | `string` | `null` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_approved_dov1.0.1s"></a> [approved\_dov1.0.1s](#input\_approved\_dov1.0.1s) | List of approved dov1.0.1s for the Connect instance | `list(string)` | `[]` | no |
| <a name="input_connect_admin_sso_xml_filepath"></a> [connect\_admin\_sso\_xml\_filepath](#input\_connect\_admin\_sso\_xml\_filepath) | Specifies xml filepath | `string` | `"./"` | no |
| <a name="input_connect_agent_sso_xml_filepath"></a> [connect\_agent\_sso\_xml\_filepath](#input\_connect\_agent\_sso\_xml\_filepath) | Specifies xml filepath | `string` | `"./"` | no |
| <a name="input_contact_flow_modules"></a> [contact\_flow\_modules](#input\_contact\_flow\_modules) | A map of Amazon Connect Contact Flow Modules.<br><br>The key of the map is the Contact Flow Module `name`. The value is the configuration for that Contact Flow, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_contact_flow_module) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <contact_flow_module_name> = {<br>    content      = optional(string) # one required<br>    content_hash = optional(string) # one required<br>    description  = optional(string)<br>    filename     = optional(string) # one required<br>    tags         = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_contact_flows"></a> [contact\_flows](#input\_contact\_flows) | A map of Amazon Connect Contact Flows.<br><br>The key of the map is the Contact Flow `name`. The value is the configuration for that Contact Flow, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_contact_flow) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <contact_flow_name> = {<br>    content      = optional(string) # one required<br>    content_hash = optional(string) # one required<br>    description  = optional(string)<br>    filename     = optional(string) # one required<br>    tags         = optional(map(string))<br>    type         = optional(string)<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_create_awscc_connect_approved_origin"></a> [create\_awscc\_connect\_approved\_origin](#input\_create\_awscc\_connect\_approved\_origin) | Flag to control the creation of approved origins | `bool` | `false` | no |
| <a name="input_create_instance"></a> [create\_instance](#input\_create\_instance) | Controls if the aws\_connect\_instance resource should be created. Defaults to true. | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_hours_of_operations"></a> [hours\_of\_operations](#input\_hours\_of\_operations) | A map of Amazon Connect Hours of Operations.<br><br>The key of the map is the Hours of Operation `name`. The value is the configuration for that Hours of Operation, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_hours_of_operation) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <hours_of_operation_name> = {<br>    config = [<br>      {<br>        day = string<br>        end_time = {<br>          hours   = number<br>          minutes = number<br>        }<br>        start_time = {<br>          hours   = number<br>          minutes = number<br>        }<br>      }<br>    ]<br>    description = optional(string)<br>    tags        = optional(map(string))<br>    time_zone   = string<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_hours_of_operations_tags"></a> [hours\_of\_operations\_tags](#input\_hours\_of\_operations\_tags) | Additional tags to add to all Hours of Operations resources. | `map(string)` | `{}` | no |
| <a name="input_instance_contact_flow_logs_enabled"></a> [instance\_contact\_flow\_logs\_enabled](#input\_instance\_contact\_flow\_logs\_enabled) | Specifies whether contact flow logs are enabled. Defaults to false. | `bool` | `true` | no |
| <a name="input_instance_contact_lens_enabled"></a> [instance\_contact\_lens\_enabled](#input\_instance\_contact\_lens\_enabled) | Specifies whether contact lens is enabled. Defaults to true. | `bool` | `null` | no |
| <a name="input_instance_id"></a> [instance\_id](#input\_instance\_id) | If create\_instance is set to false, you may still create other resources and pass in an instance ID that was created outside this module. Ignored if create\_instance is true. | `string` | `null` | no |
| <a name="input_instance_identity_management_type"></a> [instance\_identity\_management\_type](#input\_instance\_identity\_management\_type) | Specifies the identity management type attached to the instance. Allowed values are: SAML, CONNECT\_MANAGED, EXISTING\_DIRECTORY. | `string` | `null` | no |
| <a name="input_instance_storage_configs"></a> [instance\_storage\_configs](#input\_instance\_storage\_configs) | A map of Amazon Connect Instance Storage Configs.<br><br>The key of the map is the Instance Storage Config `resource_type`. The value is the configuration for that Instance Storage Config, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_instance_storage_config#storage_config) (except `resource_type` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <instance_storage_config_resource_type> = {<br>    kinesis_firehose_config = optional({<br>      firehose_arn = string<br>    })<br>    kinesis_stream_config = optional({<br>      stream_arn = string<br>    })<br>    kinesis_video_stream_config = optional({<br>      encryption_config = {<br>        encryption_type = string<br>        key_id          = string<br>      }<br>      prefix                 = string<br>      retention_period_hours = number<br>    })<br>    s3_config = optional({<br>      bucket_name   = string<br>      bucket_prefix = string<br>      encryption_config = optional({<br>        encryption_type = string<br>        key_id          = string<br>      })<br>    })<br>    storage_type = string<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_lambda_function_associations"></a> [lambda\_function\_associations](#input\_lambda\_function\_associations) | A map of Lambda Function ARNs to associate to the Amazon Connect Instance, the key is a static/arbitrary name and value is the Lambda ARN.<br><br>Example/available options:<pre>{<br>  <lambda_function_association_name> = string<br>}</pre> | `map(string)` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name | `string` | n/a | yes |
| <a name="input_multi_party_conference_enabled"></a> [multi\_party\_conference\_enabled](#input\_multi\_party\_conference\_enabled) | Specifies multi-party calls/conference is enabled. | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_permissions_boundary_arn"></a> [permissions\_boundary\_arn](#input\_permissions\_boundary\_arn) | Permissions boundary ARN to use for IAM role | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the acm, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_queue_tags"></a> [queue\_tags](#input\_queue\_tags) | Additional tags to add to all Queue resources. | `map(string)` | `{}` | no |
| <a name="input_queues"></a> [queues](#input\_queues) | A map of Amazon Connect Queues.<br><br>The key of the map is the Queue `name`. The value is the configuration for that Queue, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_queue) (except `name` which is the key, and `instance_id which` is created or passed in).<br><br>Example/available options:<pre>{<br>  <queue_name> = {<br>    description            = optional(string)<br>    hours_of_operation_id  = string<br>    max_contacts           = optional(number)<br>    outbound_caller_config = optional({<br>      outbound_caller_id_name      = optional(string)<br>      outbound_caller_id_number_id = optional(string)<br>      outbound_flow_id             = optional(string)<br>    })<br>    quick_connect_ids = optional(list(string))<br>    status            = optional(string)<br>    tags              = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_quick_connects"></a> [quick\_connects](#input\_quick\_connects) | A map of Amazon Connect Quick Connect.<br><br>The key of the map is the Quick Connect `name`. The value is the configuration for that Quick Connect, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_quick_connect) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <quick_connect_name> = {<br>    description          = optional(string)<br>    quick_connect_config = {<br>      quick_connect_type = string<br>      phone_config = optional({<br>        phone_number = string<br>      })<br>      queue_config = optional({<br>        contact_flow_id = string<br>        queue_id        = string<br>      })<br>      user_config  = optional({<br>        contact_flow_id = string<br>        queue_id        = string<br>      })<br>    })<br>    tags = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_role_name"></a> [role\_name](#input\_role\_name) | n/a | `string` | `null` | no |
| <a name="input_routing_profiles"></a> [routing\_profiles](#input\_routing\_profiles) | A map of Amazon Connect Routing Profile.<br><br>The key of the map is the Routing Profile `name`. The value is the configuration for that Routing Profile, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_routing_profile) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <routing_profile_name> = {<br>    default_outbound_queue_id = string<br>    description               = string<br>    media_concurrencies = [<br>      {<br>        channel     = string<br>        concurrency = number<br>      }<br>    ]<br>    queue_configs = optional([<br>      {<br>        channel  = string<br>        delay    = number<br>        priority = number<br>        queue_id = string<br>      }<br>    ])<br>    tags = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_security_profiles"></a> [security\_profiles](#input\_security\_profiles) | A map of Amazon Connect Security Profile.<br><br>The key of the map is the Security Profile `name`. The value is the configuration for that Security Profile, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_security_profile) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <security_profile_name> = {<br>    description = optional(string)<br>    permissions = optional(list(string))<br>    tags        = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_sso_integration"></a> [sso\_integration](#input\_sso\_integration) | n/a | `bool` | `false` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_user_hierarchy_groups"></a> [user\_hierarchy\_groups](#input\_user\_hierarchy\_groups) | A map of Amazon Connect User Hierarchy Groups.<br><br>The key of the map is the User Hierarchy Group `name`. The value is the configuration for that User, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_user_hierarchy_group) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <user_hierarchy_group_name> = {<br>    parent_group_id  = optional(string)<br>    tags             = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_user_hierarchy_structure"></a> [user\_hierarchy\_structure](#input\_user\_hierarchy\_structure) | A map of Amazon Connect User Hierarchy Structure, containing keys for for zero or many levels: `level_one`, `level_two`, `level_three`, `level_four`, and `level_five`. The values are the `name` for that level. See [documentation here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_user_hierarchy_structure).<br><br>Example/available options:<pre>{<br>  level_one = string<br>}</pre> | `map(string)` | `{}` | no |
| <a name="input_users"></a> [users](#input\_users) | A map of Amazon Connect Users.<br><br>The key of the map is the User `name`. The value is the configuration for that User, supporting all arguments [documented here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/connect_user) (except `name` which is the key, and `instance_id` which is created or passed in).<br><br>Example/available options:<pre>{<br>  <user_name> = {<br>    directory_user_id  = optional(string)<br>    hierarchy_group_id = optional(string)<br>    identity_info = optional({<br>      email      = optional(string)<br>      first_name = optional(string)<br>      last_name  = optional(string)<br>    })<br>    password = optional(string)<br>    phone_config = {<br>      phone_type                    = string<br>      after_contact_work_time_limit = optional(number)<br>      auto_accept                   = optional(bool)<br>      desk_phone_number             = optional(string)<br>    }<br>    routing_profile_id   = string<br>    security_profile_ids = list(string)<br>    tags                 = optional(map(string))<br>  }<br>}</pre> | `any` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_bot_associations"></a> [bot\_associations](#output\_bot\_associations) | Full output attributes of bot\_associations resource(s). |
| <a name="output_connect_idp_admin_arn"></a> [connect\_idp\_admin\_arn](#output\_connect\_idp\_admin\_arn) | The ARN of connect admin panel |
| <a name="output_connect_idp_agent_arn"></a> [connect\_idp\_agent\_arn](#output\_connect\_idp\_agent\_arn) | The ARN of connect agent panel |
| <a name="output_contact_flow_modules"></a> [contact\_flow\_modules](#output\_contact\_flow\_modules) | Full output attributes of contact\_flow\_modules resource(s). |
| <a name="output_contact_flows"></a> [contact\_flows](#output\_contact\_flows) | Full output attributes of contact\_flows resource(s). |
| <a name="output_current_account_number"></a> [current\_account\_number](#output\_current\_account\_number) | Current AWS account number where the module is being executed. |
| <a name="output_current_region"></a> [current\_region](#output\_current\_region) | Current AWS region where the module is being executed. |
| <a name="output_hours_of_operations"></a> [hours\_of\_operations](#output\_hours\_of\_operations) | Full output attributes of hours\_of\_operations resource(s). |
| <a name="output_instance"></a> [instance](#output\_instance) | Full output attributes of aws\_connect\_instance resource. |
| <a name="output_instance_arn"></a> [instance\_arn](#output\_instance\_arn) | Amazon Connect instance ARN. |
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | Amazon Connect instance ID. If create\_instance = false, var.instance\_id is returned. |
| <a name="output_instance_storage_configs"></a> [instance\_storage\_configs](#output\_instance\_storage\_configs) | Full output attributes of instance\_storage\_configs resource(s). |
| <a name="output_lambda_function_associations"></a> [lambda\_function\_associations](#output\_lambda\_function\_associations) | Full output attributes of lambda\_function\_associations resource(s). |
| <a name="output_queues"></a> [queues](#output\_queues) | Full output attributes of aws\_connect\_queue resource(s). |
| <a name="output_quick_connects"></a> [quick\_connects](#output\_quick\_connects) | Full output attributes of quick\_connects resource(s). |
| <a name="output_routing_profiles"></a> [routing\_profiles](#output\_routing\_profiles) | Full output attributes of routing\_profiles resource(s). |
| <a name="output_security_profiles"></a> [security\_profiles](#output\_security\_profiles) | Full output attributes of security\_profiles resource(s). |
| <a name="output_user_hierarchy_groups"></a> [user\_hierarchy\_groups](#output\_user\_hierarchy\_groups) | Full output attributes of user\_hierarchy\_groups resource(s). |
| <a name="output_user_hierarchy_structure"></a> [user\_hierarchy\_structure](#output\_user\_hierarchy\_structure) | Full output attributes of user\_hierarchy\_structure resource(s). |
| <a name="output_users"></a> [users](#output\_users) | Full output attributes of users resource(s). |
| <a name="output_vocabularies"></a> [vocabularies](#output\_vocabularies) | Full output attributes of vocabularies resource(s). |
<!-- END_TF_DOCS -->
