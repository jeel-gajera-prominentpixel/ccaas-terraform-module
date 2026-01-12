# AWS rds-proxy Terraform Module

## How to use this module:

### aws rds-proxy basic module usage with the required input variables:
```terraform
module "rds-proxy_basic" {
  source                 = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-proxy?ref=main"
  prefix_company         = "jb"
  identifier             = "test"
  lob                    = "itsd"
  prefix_region          = "usw2"
  application            = "recordings"
  env                    = "sandbox"
  engine_family = "POSTGRESQL"
  debug_logging = false
}


```

### aws rds-proxy advanced module usage with all the optional input variables:



```terraform
module "rds-proxy_advance" {
  source                          = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-proxy?ref=main"
  prefix_company                  = "jb"
  identifier                      = "test"
  lob                             = "itsd"
  prefix_region                   = "usw2"
  application                     = "recordings"
  env                             = "sandbox"
  create_rds_proxy       = true
  name                   = "rds proxy example"
  vpc_subnet_ids         = []
  vpc_security_group_ids = []

  endpoints = {
    read_write = {
        name                   = "read-write-endpoint"
        vpc_subnet_ids         = []
        vpc_security_group_ids = []
        tags                   = local.tags
      },
      read_only = {
        name                   = "read-only-endpoint"
        vpc_subnet_ids         = []
        vpc_security_group_ids = []
        target_role            = "READ_ONLY"
        tags                   = local.tags
      }
  }

  auth = {
      "root" = {
        description = "Cluster generated master user password"
        secret_arn  = ""
      }
    }

  engine_family = "POSTGRESQL"
  debug_logging = false

  # Target Aurora cluster
  target_db_cluster     = false
  db_cluster_identifier = "db cluster identifier"

  # Target RDS instance
  target_db_instance     = false
  db_instance_identifier = "db instance identifier"

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
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.5 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_rds_proxy"></a> [rds\_proxy](#module\_rds\_proxy) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-rds-proxy | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name. | `string` | n/a | yes |
| <a name="input_auth"></a> [auth](#input\_auth) | Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters<br>Example/available options:<pre>{<br>  "root" = {<br>    description = "Cluster generated master user password"<br>    secret_arn  = module.rds.cluster_master_user_secret[0].secret_arn<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_create_iam_policy"></a> [create\_iam\_policy](#input\_create\_iam\_policy) | Determines whether an IAM policy is created | `bool` | `true` | no |
| <a name="input_create_iam_role"></a> [create\_iam\_role](#input\_create\_iam\_role) | Determines whether an IAM role is created | `bool` | `true` | no |
| <a name="input_create_rds_proxy"></a> [create\_rds\_proxy](#input\_create\_rds\_proxy) | Whether cluster should be created (affects nearly all resources) | `bool` | `false` | no |
| <a name="input_db_cluster_identifier"></a> [db\_cluster\_identifier](#input\_db\_cluster\_identifier) | DB cluster identifier | `string` | `""` | no |
| <a name="input_db_instance_identifier"></a> [db\_instance\_identifier](#input\_db\_instance\_identifier) | DB instance identifier | `string` | `""` | no |
| <a name="input_debug_logging"></a> [debug\_logging](#input\_debug\_logging) | Whether the proxy includes detailed information about SQL statements in its logs | `bool` | `false` | no |
| <a name="input_endpoints"></a> [endpoints](#input\_endpoints) | Map of DB proxy endpoints to create and their attributes (see `aws_db_proxy_endpoint`)<br>Example/available options:<pre>{<br>  read_write = {<br>    name                   = "read-write-endpoint"<br>    vpc_subnet_ids         = module.vpc.private_subnets<br>    vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]<br>    tags                   = local.tags<br>  },<br>  read_only = {<br>    name                   = "read-only-endpoint"<br>    vpc_subnet_ids         = module.vpc.private_subnets<br>    vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]<br>    target_role            = "READ_ONLY"<br>    tags                   = local.tags<br>  }<br>}</pre> | `any` | `{}` | no |
| <a name="input_engine_family"></a> [engine\_family](#input\_engine\_family) | The kind of database engine that the proxy will connect to. Valid values are `MYSQL` or `POSTGRESQL` | `string` | `""` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_iam_policy_name"></a> [iam\_policy\_name](#input\_iam\_policy\_name) | The name of the role policy. If omitted, Terraform will assign a random, unique name | `string` | `""` | no |
| <a name="input_iam_role_description"></a> [iam\_role\_description](#input\_iam\_role\_description) | The description of the role | `string` | `""` | no |
| <a name="input_iam_role_force_detach_policies"></a> [iam\_role\_force\_detach\_policies](#input\_iam\_role\_force\_detach\_policies) | Specifies to force detaching any policies the role has before destroying it | `bool` | `true` | no |
| <a name="input_iam_role_max_session_duration"></a> [iam\_role\_max\_session\_duration](#input\_iam\_role\_max\_session\_duration) | The maximum session duration (in seconds) that you want to set for the specified role | `number` | `43200` | no |
| <a name="input_iam_role_name"></a> [iam\_role\_name](#input\_iam\_role\_name) | The name of the role. If omitted, Terraform will assign a random, unique name | `string` | `""` | no |
| <a name="input_iam_role_path"></a> [iam\_role\_path](#input\_iam\_role\_path) | The path to the role | `string` | `null` | no |
| <a name="input_iam_role_permissions_boundary"></a> [iam\_role\_permissions\_boundary](#input\_iam\_role\_permissions\_boundary) | The ARN of the policy that is used to set the permissions boundary for the role | `string` | `null` | no |
| <a name="input_iam_role_tags"></a> [iam\_role\_tags](#input\_iam\_role\_tags) | A map of tags to apply to the IAM role | `map(string)` | `{}` | no |
| <a name="input_idle_client_timeout"></a> [idle\_client\_timeout](#input\_idle\_client\_timeout) | The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it | `number` | `1800` | no |
| <a name="input_kms_key_arns"></a> [kms\_key\_arns](#input\_kms\_key\_arns) | List of KMS Key ARNs to allow access to decrypt SecretsManager secrets | `list(string)` | `[]` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name | `string` | n/a | yes |
| <a name="input_log_group_kms_key_id"></a> [log\_group\_kms\_key\_id](#input\_log\_group\_kms\_key\_id) | The ARN of the KMS Key to use when encrypting log data | `string` | `null` | no |
| <a name="input_log_group_retention_in_days"></a> [log\_group\_retention\_in\_days](#input\_log\_group\_retention\_in\_days) | Specifies the number of days you want to retain log events in the log group | `number` | `30` | no |
| <a name="input_log_group_tags"></a> [log\_group\_tags](#input\_log\_group\_tags) | A map of tags to apply to the CloudWatch log group | `map(string)` | `{}` | no |
| <a name="input_manage_log_group"></a> [manage\_log\_group](#input\_manage\_log\_group) | Determines whether Terraform will create/manage the CloudWatch log group or not. Note - this will fail if set to true after the log group has been created as the resource will already exist | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the rds proxy , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_require_tls"></a> [require\_tls](#input\_require\_tls) | A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy | `bool` | `true` | no |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager | `string` | `""` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching rdss will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_target_db_cluster"></a> [target\_db\_cluster](#input\_target\_db\_cluster) | Determines whether DB cluster is targeted by proxy | `bool` | `false` | no |
| <a name="input_target_db_instance"></a> [target\_db\_instance](#input\_target\_db\_instance) | Determines whether DB instance is targeted by proxy | `bool` | `false` | no |
| <a name="input_use_policy_name_prefix"></a> [use\_policy\_name\_prefix](#input\_use\_policy\_name\_prefix) | Whether to use unique name beginning with the specified `iam_policy_name` | `bool` | `false` | no |
| <a name="input_use_role_name_prefix"></a> [use\_role\_name\_prefix](#input\_use\_role\_name\_prefix) | Whether to use unique name beginning with the specified `iam_role_name` | `bool` | `false` | no |
| <a name="input_vpc_security_group_ids"></a> [vpc\_security\_group\_ids](#input\_vpc\_security\_group\_ids) | One or more VPC security group IDs to associate with the new proxy | `list(string)` | `[]` | no |
| <a name="input_vpc_subnet_ids"></a> [vpc\_subnet\_ids](#input\_vpc\_subnet\_ids) | One or more VPC subnet IDs to associate with the new proxy | `list(string)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_db_proxy_endpoints"></a> [db\_proxy\_endpoints](#output\_db\_proxy\_endpoints) | Array containing the full resource object and attributes for all DB proxy endpoints created |
| <a name="output_iam_role_arn"></a> [iam\_role\_arn](#output\_iam\_role\_arn) | The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager. |
| <a name="output_iam_role_name"></a> [iam\_role\_name](#output\_iam\_role\_name) | IAM role name |
| <a name="output_iam_role_unique_id"></a> [iam\_role\_unique\_id](#output\_iam\_role\_unique\_id) | Stable and unique string identifying the IAM role |
| <a name="output_log_group_arn"></a> [log\_group\_arn](#output\_log\_group\_arn) | The Amazon Resource Name (ARN) of the CloudWatch log group |
| <a name="output_proxy_arn"></a> [proxy\_arn](#output\_proxy\_arn) | The Amazon Resource Name (ARN) for the proxy |
| <a name="output_proxy_default_target_group_arn"></a> [proxy\_default\_target\_group\_arn](#output\_proxy\_default\_target\_group\_arn) | The Amazon Resource Name (ARN) for the default target group |
| <a name="output_proxy_default_target_group_id"></a> [proxy\_default\_target\_group\_id](#output\_proxy\_default\_target\_group\_id) | The ID for the default target group |
| <a name="output_proxy_default_target_group_name"></a> [proxy\_default\_target\_group\_name](#output\_proxy\_default\_target\_group\_name) | The name of the default target group |
| <a name="output_proxy_endpoint"></a> [proxy\_endpoint](#output\_proxy\_endpoint) | The endpoint that you can use to connect to the proxy |
| <a name="output_proxy_id"></a> [proxy\_id](#output\_proxy\_id) | The ID for the proxy |
| <a name="output_proxy_target_endpoint"></a> [proxy\_target\_endpoint](#output\_proxy\_target\_endpoint) | Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type |
| <a name="output_proxy_target_id"></a> [proxy\_target\_id](#output\_proxy\_target\_id) | Identifier of `db_proxy_name`, `target_group_name`, target type (e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`), and resource identifier separated by forward slashes (/) |
| <a name="output_proxy_target_port"></a> [proxy\_target\_port](#output\_proxy\_target\_port) | Port for the target RDS DB Instance or Aurora DB Cluster |
| <a name="output_proxy_target_rds_resource_id"></a> [proxy\_target\_rds\_resource\_id](#output\_proxy\_target\_rds\_resource\_id) | Identifier representing the DB Instance or DB Cluster target |
| <a name="output_proxy_target_target_arn"></a> [proxy\_target\_target\_arn](#output\_proxy\_target\_target\_arn) | Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API |
| <a name="output_proxy_target_tracked_cluster_id"></a> [proxy\_target\_tracked\_cluster\_id](#output\_proxy\_target\_tracked\_cluster\_id) | DB Cluster identifier for the DB Instance target. Not returned unless manually importing an RDS\_INSTANCE target that is part of a DB Cluster |
| <a name="output_proxy_target_type"></a> [proxy\_target\_type](#output\_proxy\_target\_type) | Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER` |
<!-- END_TF_DOCS -->
