# AWS rds Terraform Module

## How to use this module:

### aws rds basic module usage with the required input variables:
```terraform
module "rds_basic" {
  source                           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-aurora/global_cluster?ref=main"
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
  source                           = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-rds-aurora/global_cluster?ref=main"
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

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_rds_advance"></a> [rds\_advance](#module\_rds\_advance) | git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress/global_cluster | main |
| <a name="module_rds_basic"></a> [rds\_basic](#module\_rds\_basic) | git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-aurora-postgress/global_cluster | main |

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END_TF_DOCS -->
