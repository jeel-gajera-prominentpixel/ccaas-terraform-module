# AWS rds Terraform Module

## How to use this module:

### aws rds basic module usage with the required input variables:
```terraform
module "rds_basic" {
  source                           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress/global_cluster?ref=main"
  prefix_company                   = "jb"
  identifier                       = "test"
  lob                              = "itsd"
  prefix_region                    = "usw2"
  application                      = "recordings"
  env                              = "sandbox"
  create                           = false
  create_global_cluster            = true
  global_cluster_engine            = "postgres"
  global_cluster_version           = "14"
  global_cluster_db_name           = "completePostgresql"
  global_cluster_storage_encrypted = true
}


```

### aws rds advanced module usage with all the optional input variables:



```terraform

module "rds_advance" {
  source                           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress/global_cluster?ref=main"
  prefix_company                   = "jb"
  identifier                       = "test"
  lob                              = "itsd"
  prefix_region                    = "usw2"
  application                      = "recordings"
  env                              = "sandbox"
  create                           = false
  create_global_cluster            = true
  global_cluster_engine            = "postgres"
  global_cluster_version           = "14"
  global_cluster_db_name           = "completePostgresql"
  global_cluster_storage_encrypted = true
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

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_global_rds_cluster"></a> [global\_rds\_cluster](#module\_global\_rds\_cluster) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-rds-aurora | main |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the rds, will be appended with the company, lob, env and region to form a rds name. | `string` | n/a | yes |
| <a name="input_create_global_cluster"></a> [create\_global\_cluster](#input\_create\_global\_cluster) | Whether global cluster should be created | `bool` | `false` | no |
| <a name="input_deletion_protection"></a> [deletion\_protection](#input\_deletion\_protection) | (Optional) If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false. | `bool` | `false` | no |
| <a name="input_engine_lifecycle_support"></a> [engine\_lifecycle\_support](#input\_engine\_lifecycle\_support) | (Optional) The life cycle type for this DB instance. This setting applies only to Aurora PostgreSQL-based global databases. Valid values are open-source-rds-extended-support, open-source-rds-extended-support-disabled. Default value is open-source-rds-extended-support. [Using Amazon RDS Extended Support]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html | `string` | `null` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_force_destroy"></a> [force\_destroy](#input\_force\_destroy) | (Optional) Enable to remove DB Cluster members from Global Cluster on destroy. Required with source\_db\_cluster\_identifier. | `bool` | `false` | no |
| <a name="input_global_cluster_db_name"></a> [global\_cluster\_db\_name](#input\_global\_cluster\_db\_name) | Name for an automatically created database on global cluster creation | `string` | `null` | no |
| <a name="input_global_cluster_engine"></a> [global\_cluster\_engine](#input\_global\_cluster\_engine) | The name of the database engine to be used for this global cluster. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql` | `string` | `null` | no |
| <a name="input_global_cluster_identifier"></a> [global\_cluster\_identifier](#input\_global\_cluster\_identifier) | Global cluster identifier | `string` | `""` | no |
| <a name="input_global_cluster_storage_encrypted"></a> [global\_cluster\_storage\_encrypted](#input\_global\_cluster\_storage\_encrypted) | Specifies whether the DB cluster is encrypted. The default is `true` | `bool` | `true` | no |
| <a name="input_global_cluster_version"></a> [global\_cluster\_version](#input\_global\_cluster\_version) | The database engine version for global cluster. Updating this argument results in an outage | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the rds, will be appended with the company, lob, env and region to form a rds name | `string` | n/a | yes |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the rds, will be appended with the company, lob, env and region to form a rds name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the rds , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_source_db_cluster_identifier"></a> [source\_db\_cluster\_identifier](#input\_source\_db\_cluster\_identifier) | (Optional) Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. Terraform cannot perform drift detection of this value. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_global_cluster_arn"></a> [global\_cluster\_arn](#output\_global\_cluster\_arn) | The ARN of global cluster |
| <a name="output_global_cluster_db_name"></a> [global\_cluster\_db\_name](#output\_global\_cluster\_db\_name) | The database name of global cluster |
| <a name="output_global_cluster_endpoint"></a> [global\_cluster\_endpoint](#output\_global\_cluster\_endpoint) | The endpoint of global cluster |
| <a name="output_global_cluster_engine"></a> [global\_cluster\_engine](#output\_global\_cluster\_engine) | The engine of global cluster |
| <a name="output_global_cluster_engine_version"></a> [global\_cluster\_engine\_version](#output\_global\_cluster\_engine\_version) | The engine version of global cluster |
| <a name="output_global_cluster_id"></a> [global\_cluster\_id](#output\_global\_cluster\_id) | The ID of global cluster |
| <a name="output_global_cluster_members"></a> [global\_cluster\_members](#output\_global\_cluster\_members) | The members of global cluster |
| <a name="output_global_cluster_resource_id"></a> [global\_cluster\_resource\_id](#output\_global\_cluster\_resource\_id) | The resource id of global cluster |
| <a name="output_global_cluster_tags_all"></a> [global\_cluster\_tags\_all](#output\_global\_cluster\_tags\_all) | The all tags of global cluster |
<!-- END_TF_DOCS -->
