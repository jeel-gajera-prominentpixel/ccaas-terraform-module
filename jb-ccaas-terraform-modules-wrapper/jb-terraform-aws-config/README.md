# AWS config Terraform Module

## How to use this module:

### aws config module usage with the required input variables:

```terraform
module "config" {
  source           = git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-config?ref=<version>
  prefix_company   = "jb"
  lob              = "itsd"
  prefix_region = "usw2"
  application      = "recordings"
  env              = "sandbox"
  create_sns_topic = true
  create_iam_role  = true
  managed_rules = {
    rule1 = {
      description      = "Rule 1 description"
      enabled          = true
      identifier       = "rule1_identifier"
      input_parameters = {}
      tags             = {}
    },
    rule2 = {
      description      = "Rule 2 description"
      enabled          = true
      identifier       = "rule2_identifier"
      input_parameters = {}
      tags             = {}
    }
  }
  force_destroy = false
  s3_bucket_id  = "test-bucket-123"
  s3_bucket_arn = "arn:aws:s3:::test-bucket-123"
  tags          = local.tags
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
| <a name="module_config"></a> [config](#module\_config) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-config | main |

## Resources

| Name | Type |
|------|------|
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the aws config, will be appended with the company, lob, env and region to form a aws config name. | `string` | n/a | yes |
| <a name="input_create_iam_role"></a> [create\_iam\_role](#input\_create\_iam\_role) | Flag to indicate whether an IAM Role should be created to grant the proper permissions for AWS Config | `bool` | `false` | no |
| <a name="input_create_sns_topic"></a> [create\_sns\_topic](#input\_create\_sns\_topic) | Flag to indicate whether an SNS topic should be created for notifications<br>If you want to send findings to a new SNS topic, set this to true and provide a valid configuration for subscribers | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_force_destroy"></a> [force\_destroy](#input\_force\_destroy) | A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable | `bool` | `false` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the aws config, will be appended with the company, lob, env and region to form a aws config name. | `string` | n/a | yes |
| <a name="input_managed_rules"></a> [managed\_rules](#input\_managed\_rules) | A list of AWS Managed Rules that should be enabled on the account.<br><br>See the following for a list of possible rules to enable:<br>https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html | <pre>map(object({<br>    description      = string<br>    identifier       = string<br>    input_parameters = any<br>    tags             = map(string)<br>    enabled          = bool<br>  }))</pre> | `{}` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the aws config, will be appended with the company, lob, env and region to form a aws config name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the aws config , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_s3_bucket_arn"></a> [s3\_bucket\_arn](#input\_s3\_bucket\_arn) | The ARN of the S3 bucket used to store the configuration history | `string` | n/a | yes |
| <a name="input_s3_bucket_id"></a> [s3\_bucket\_id](#input\_s3\_bucket\_id) | The id (name) of the S3 bucket used to store the configuration history | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_config_recorder_id"></a> [config\_recorder\_id](#output\_config\_recorder\_id) | The id of the AWS Config Recorder that was created |
<!-- END_TF_DOCS -->
