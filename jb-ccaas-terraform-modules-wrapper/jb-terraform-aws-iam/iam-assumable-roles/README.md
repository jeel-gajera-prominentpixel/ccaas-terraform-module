# iam-assumable-roles

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
| <a name="module_iam_role"></a> [iam\_role](#module\_iam\_role) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-iam/modules/iam-assumable-role | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_admin_role_policy_arn"></a> [admin\_role\_policy\_arn](#input\_admin\_role\_policy\_arn) | Policy ARN to use for admin role | `string` | `"arn:aws:iam::aws:policy/AdministratorAccess"` | no |
| <a name="input_allow_self_assume_role"></a> [allow\_self\_assume\_role](#input\_allow\_self\_assume\_role) | Determines whether to allow the role to be [assume itself](https://aws.amazon.com/blogs/security/announcing-an-update-to-iam-role-trust-policy-behavior/) | `bool` | `false` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_attach_admin_policy"></a> [attach\_admin\_policy](#input\_attach\_admin\_policy) | Whether to attach an admin policy to a role | `bool` | `false` | no |
| <a name="input_attach_poweruser_policy"></a> [attach\_poweruser\_policy](#input\_attach\_poweruser\_policy) | Whether to attach a poweruser policy to a role | `bool` | `false` | no |
| <a name="input_attach_readonly_policy"></a> [attach\_readonly\_policy](#input\_attach\_readonly\_policy) | Whether to attach a readonly policy to a role | `bool` | `false` | no |
| <a name="input_create_custom_role_trust_policy"></a> [create\_custom\_role\_trust\_policy](#input\_create\_custom\_role\_trust\_policy) | Whether to create a custom\_role\_trust\_policy. Prevent errors with count, when custom\_role\_trust\_policy is computed | `bool` | `false` | no |
| <a name="input_create_instance_profile"></a> [create\_instance\_profile](#input\_create\_instance\_profile) | Whether to create an instance profile | `bool` | `false` | no |
| <a name="input_create_role"></a> [create\_role](#input\_create\_role) | Whether to create a role | `bool` | `false` | no |
| <a name="input_custom_role_policy_arns"></a> [custom\_role\_policy\_arns](#input\_custom\_role\_policy\_arns) | List of ARNs of IAM policies to attach to IAM role | `list(string)` | `[]` | no |
| <a name="input_custom_role_trust_policy"></a> [custom\_role\_trust\_policy](#input\_custom\_role\_trust\_policy) | A custom role trust policy. (Only valid if create\_custom\_role\_trust\_policy = true) | `string` | `""` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_force_detach_policies"></a> [force\_detach\_policies](#input\_force\_detach\_policies) | Whether policies should be detached from this role when destroying | `bool` | `false` | no |
| <a name="input_inline_policy_statements"></a> [inline\_policy\_statements](#input\_inline\_policy\_statements) | List of inline policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) to attach to IAM role as an inline policy | `any` | `[]` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_max_session_duration"></a> [max\_session\_duration](#input\_max\_session\_duration) | Maximum CLI/API session duration in seconds between 3600 and 43200 | `number` | `3600` | no |
| <a name="input_mfa_age"></a> [mfa\_age](#input\_mfa\_age) | Max age of valid MFA (in seconds) for roles which require MFA | `number` | `86400` | no |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_number_of_custom_role_policy_arns"></a> [number\_of\_custom\_role\_policy\_arns](#input\_number\_of\_custom\_role\_policy\_arns) | Number of IAM policies to attach to IAM role | `number` | `null` | no |
| <a name="input_poweruser_role_policy_arn"></a> [poweruser\_role\_policy\_arn](#input\_poweruser\_role\_policy\_arn) | Policy ARN to use for poweruser role | `string` | `"arn:aws:iam::aws:policy/PowerUserAccess"` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_readonly_role_policy_arn"></a> [readonly\_role\_policy\_arn](#input\_readonly\_role\_policy\_arn) | Policy ARN to use for readonly role | `string` | `"arn:aws:iam::aws:policy/ReadOnlyAccess"` | no |
| <a name="input_role_description"></a> [role\_description](#input\_role\_description) | IAM Role description | `string` | `""` | no |
| <a name="input_role_name"></a> [role\_name](#input\_role\_name) | IAM role name | `string` | `null` | no |
| <a name="input_role_name_prefix"></a> [role\_name\_prefix](#input\_role\_name\_prefix) | IAM role name prefix | `string` | `null` | no |
| <a name="input_role_path"></a> [role\_path](#input\_role\_path) | Path of IAM role | `string` | `"/"` | no |
| <a name="input_role_permissions_boundary_arn"></a> [role\_permissions\_boundary\_arn](#input\_role\_permissions\_boundary\_arn) | Permissions boundary ARN to use for IAM role | `string` | `""` | no |
| <a name="input_role_requires_mfa"></a> [role\_requires\_mfa](#input\_role\_requires\_mfa) | Whether role requires MFA | `bool` | `true` | no |
| <a name="input_role_requires_session_name"></a> [role\_requires\_session\_name](#input\_role\_requires\_session\_name) | Determines if the role-session-name variable is needed when assuming a role(https://aws.amazon.com/blogs/security/easily-control-naming-individual-iam-role-sessions/) | `bool` | `false` | no |
| <a name="input_role_session_name"></a> [role\_session\_name](#input\_role\_session\_name) | role\_session\_name for roles which require this parameter when being assumed. By default, you need to set your own username as role\_session\_name | `list(string)` | <pre>[<br>  "${aws:username}"<br>]</pre> | no |
| <a name="input_role_sts_externalid"></a> [role\_sts\_externalid](#input\_role\_sts\_externalid) | STS ExternalId condition values to use with a role (when MFA is not required) | `any` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_trust_policy_conditions"></a> [trust\_policy\_conditions](#input\_trust\_policy\_conditions) | [Condition constraints](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#condition) applied to the trust policy | <pre>list(object({<br>    test     = string<br>    variable = string<br>    values   = list(string)<br>  }))</pre> | `[]` | no |
| <a name="input_trusted_role_actions"></a> [trusted\_role\_actions](#input\_trusted\_role\_actions) | Additional trusted role actions | `list(string)` | <pre>[<br>  "sts:AssumeRole",<br>  "sts:TagSession"<br>]</pre> | no |
| <a name="input_trusted_role_arns"></a> [trusted\_role\_arns](#input\_trusted\_role\_arns) | ARNs of AWS entities who can assume these roles | `list(string)` | `[]` | no |
| <a name="input_trusted_role_services"></a> [trusted\_role\_services](#input\_trusted\_role\_services) | AWS Services that can assume these roles | `list(string)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_iam_instance_profile_arn"></a> [iam\_instance\_profile\_arn](#output\_iam\_instance\_profile\_arn) | ARN of IAM instance profile |
| <a name="output_iam_instance_profile_id"></a> [iam\_instance\_profile\_id](#output\_iam\_instance\_profile\_id) | IAM Instance profile's ID. |
| <a name="output_iam_instance_profile_name"></a> [iam\_instance\_profile\_name](#output\_iam\_instance\_profile\_name) | Name of IAM instance profile |
| <a name="output_iam_instance_profile_path"></a> [iam\_instance\_profile\_path](#output\_iam\_instance\_profile\_path) | Path of IAM instance profile |
| <a name="output_iam_role_arn"></a> [iam\_role\_arn](#output\_iam\_role\_arn) | ARN of IAM role |
| <a name="output_iam_role_name"></a> [iam\_role\_name](#output\_iam\_role\_name) | Name of IAM role |
| <a name="output_iam_role_path"></a> [iam\_role\_path](#output\_iam\_role\_path) | Path of IAM role |
| <a name="output_iam_role_unique_id"></a> [iam\_role\_unique\_id](#output\_iam\_role\_unique\_id) | Unique ID of IAM role |
| <a name="output_role_requires_mfa"></a> [role\_requires\_mfa](#output\_role\_requires\_mfa) | Whether IAM role requires MFA |
| <a name="output_role_sts_externalid"></a> [role\_sts\_externalid](#output\_role\_sts\_externalid) | STS ExternalId condition value to use with a role |
<!-- END_TF_DOCS -->
