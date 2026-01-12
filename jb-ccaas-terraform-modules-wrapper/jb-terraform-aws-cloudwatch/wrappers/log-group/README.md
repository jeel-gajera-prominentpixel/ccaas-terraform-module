# Wrapper for module: `modules/log-group`

The configuration in this directory contains an implementation of a single module wrapper pattern, which allows managing several copies of a module in places where using the native Terraform 0.13+ `for_each` feature is not feasible (e.g., with Terragrunt).

You may want to use a single Terragrunt configuration file to manage multiple resources without duplicating `terragrunt.hcl` files for each copy of the same module.

This wrapper does not implement any extra functionality.

## Usage with Terragrunt

`terragrunt.hcl`:

```hcl
terraform {
  source = "tfr:///terraform-aws-modules/cloudwatch/aws//wrappers/log-group"
  # Alternative source:
  # source = "git::git@github.com:terraform-aws-modules/terraform-aws-cloudwatch.git//wrappers/log-group?ref=master"
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
  source = "terraform-aws-modules/cloudwatch/aws//wrappers/log-group"

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
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.27 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_wrapper"></a> [wrapper](#module\_wrapper) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-cloudwatch/modules/log-group | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name. | `string` | n/a | yes |
| <a name="input_create"></a> [create](#input\_create) | Whether to create the Cloudwatch log group | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_kms_key_id"></a> [kms\_key\_id](#input\_kms\_key\_id) | The ARN of the KMS Key to use when encrypting logs | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_log_group_class"></a> [log\_group\_class](#input\_log\_group\_class) | Specified the log class of the log group. Possible values are: STANDARD or INFREQUENT\_ACCESS | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | A name for the log group | `string` | `""` | no |
| <a name="input_name_prefix"></a> [name\_prefix](#input\_name\_prefix) | A name prefix for the log group | `string` | `null` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the cloudtrail, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_retention_in_days"></a> [retention\_in\_days](#input\_retention\_in\_days) | Specifies the number of days you want to retain log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653. | `number` | `null` | no |
| <a name="input_skip_destroy"></a> [skip\_destroy](#input\_skip\_destroy) | Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the Terraform state | `bool` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to add to Cloudwatch log group | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_cloudwatch_log_group_arn"></a> [cloudwatch\_log\_group\_arn](#output\_cloudwatch\_log\_group\_arn) | ARN of Cloudwatch log group |
| <a name="output_cloudwatch_log_group_name"></a> [cloudwatch\_log\_group\_name](#output\_cloudwatch\_log\_group\_name) | Name of Cloudwatch log group |
<!-- END_TF_DOCS -->
