# Wrapper for module: `modules/cis-alarms`

The configuration in this directory contains an implementation of a single module wrapper pattern, which allows managing several copies of a module in places where using the native Terraform 0.13+ `for_each` feature is not feasible (e.g., with Terragrunt).

You may want to use a single Terragrunt configuration file to manage multiple resources without duplicating `terragrunt.hcl` files for each copy of the same module.

This wrapper does not implement any extra functionality.

## Usage with Terragrunt

`terragrunt.hcl`:

```hcl
terraform {
  source = "tfr:///terraform-aws-modules/cloudwatch/aws//wrappers/cis-alarms"
  # Alternative source:
  # source = "git::git@github.com:terraform-aws-modules/terraform-aws-cloudwatch.git//wrappers/cis-alarms?ref=master"
}

inputs = {
  defaults = { # Default values
    create = true
    tags = {
      Terraform   = "true"
      Environment = "dev"
    }
  }

  items = {
    my-item = {
      # omitted... can be any argument supported by the module
    }
    my-second-item = {
      # omitted... can be any argument supported by the module
    }
    # omitted...
  }
}
```

## Usage with Terraform

```hcl
module "wrapper" {
  source = "terraform-aws-modules/cloudwatch/aws//wrappers/cis-alarms"

  defaults = { # Default values
    create = true
    tags = {
      Terraform   = "true"
      Environment = "dev"
    }
  }

  items = {
    my-item = {
      # omitted... can be any argument supported by the module
    }
    my-second-item = {
      # omitted... can be any argument supported by the module
    }
    # omitted...
  }
}
```

## Example: Manage multiple S3 buckets in one Terragrunt layer

`eu-west-1/s3-buckets/terragrunt.hcl`:

```hcl
terraform {
  source = "tfr:///terraform-aws-modules/s3-bucket/aws//wrappers"
  # Alternative source:
  # source = "git::git@github.com:terraform-aws-modules/terraform-aws-s3-bucket.git//wrappers?ref=master"
}

inputs = {
  defaults = {
    force_destroy = true

    attach_elb_log_delivery_policy        = true
    attach_lb_log_delivery_policy         = true
    attach_deny_insecure_transport_policy = true
    attach_require_latest_tls_policy      = true
  }

  items = {
    bucket1 = {
      bucket = "my-random-bucket-1"
    }
    bucket2 = {
      bucket = "my-random-bucket-2"
      tags = {
        Secure = "probably"
      }
    }
  }
}
```
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.81 |
| <a name="requirement_random"></a> [random](#requirement\_random) | >= 2.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_wrapper"></a> [wrapper](#module\_wrapper) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-cloudwatch/modules/cis-alarms | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name. | `string` | n/a | yes |
| <a name="input_defaults"></a> [defaults](#input\_defaults) | Map of default values which will be used for each item. | `any` | `{}` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_items"></a> [items](#input\_items) | Maps of items to create a wrapper from. Values are passed through to the module. | `any` | `{}` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the cloudtrail, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_wrapper"></a> [wrapper](#output\_wrapper) | Map of outputs of a wrapper. |
<!-- END_TF_DOCS -->
