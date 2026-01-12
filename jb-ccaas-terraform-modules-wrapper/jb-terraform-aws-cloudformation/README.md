# jb-terraform-aws-cloudformation

## How to use this module:

### aws cloudformation module usage with the required input variables:

```terraform
module "s3_bucket_stack" {
  source         = "../"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "simple-s3-bucket-stack"
  template_body  = file("${path.module}/templates/example.yaml")
  parameters = {
    BucketName  = "my-unique-bucket-name-2024"
    Environment = "dev"
  }
  tags = {
    Environment = "Development"
    Project     = "Demo"
  }
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
| <a name="provider_aws"></a> [aws](#provider\_aws) | 6.8.0 |
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.5 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudformation_stack.stack](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack) | resource |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_capabilities"></a> [capabilities](#input\_capabilities) | (Optional) A list of capabilities. Valid values: CAPABILITY\_IAM, CAPABILITY\_NAMED\_IAM, or CAPABILITY\_AUTO\_EXPAND | `set(string)` | `null` | no |
| <a name="input_disable_rollback"></a> [disable\_rollback](#input\_disable\_rollback) | (Optional) Set to true to disable rollback of the stack if stack creation failed. Conflicts with on\_failure | `bool` | `null` | no |
| <a name="input_enabled"></a> [enabled](#input\_enabled) | Flag to determine if the stack should be created | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_iam_role_arn"></a> [iam\_role\_arn](#input\_iam\_role\_arn) | (Optional) The ARN of an IAM role that AWS CloudFormation assumes to create the stack | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | (Required) Stack name | `string` | n/a | yes |
| <a name="input_notification_arns"></a> [notification\_arns](#input\_notification\_arns) | (Optional) A list of SNS topic ARNs to publish stack related events | `set(string)` | `null` | no |
| <a name="input_on_failure"></a> [on\_failure](#input\_on\_failure) | (Optional) Action to be taken if stack creation fails. This must be one of: DO\_NOTHING, ROLLBACK, or DELETE. Conflicts with disable\_rollback | `string` | `null` | no |
| <a name="input_parameters"></a> [parameters](#input\_parameters) | (Optional) A map of Parameter structures that specify input parameters for the stack | `map(string)` | `null` | no |
| <a name="input_policy_body"></a> [policy\_body](#input\_policy\_body) | (Optional) Structure containing the stack policy body. Conflicts with policy\_url | `string` | `null` | no |
| <a name="input_policy_url"></a> [policy\_url](#input\_policy\_url) | (Optional) Location of a file containing the stack policy. Conflicts with policy\_body | `string` | `null` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the cloudfront, will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | (Optional) Map of resource tags to associate with this stack | `map(string)` | `null` | no |
| <a name="input_template_body"></a> [template\_body](#input\_template\_body) | (Optional) Structure containing the template body (max size: 51,200 bytes) | `string` | `null` | no |
| <a name="input_template_url"></a> [template\_url](#input\_template\_url) | (Optional) Location of a file containing the template body (max size: 460,800 bytes) | `string` | `null` | no |
| <a name="input_timeout_in_minutes"></a> [timeout\_in\_minutes](#input\_timeout\_in\_minutes) | (Optional) The amount of time that can pass before the stack status becomes CREATE\_FAILED | `number` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | A unique identifier of the stack |
| <a name="output_outputs"></a> [outputs](#output\_outputs) | A map of outputs from the stack |
<!-- END_TF_DOCS -->
