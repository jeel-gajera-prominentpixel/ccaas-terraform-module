# Iam-user

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
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.4 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_iam_user"></a> [iam\_user](#module\_iam\_user) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-iam/modules/iam-user | main |

## Resources

| Name | Type |
|------|------|
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_create_iam_access_key"></a> [create\_iam\_access\_key](#input\_create\_iam\_access\_key) | Whether to create IAM access key | `bool` | `true` | no |
| <a name="input_create_iam_user_login_profile"></a> [create\_iam\_user\_login\_profile](#input\_create\_iam\_user\_login\_profile) | Whether to create IAM user login profile | `bool` | `true` | no |
| <a name="input_create_user"></a> [create\_user](#input\_create\_user) | Whether to create the IAM user | `bool` | `true` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name. | `string` | n/a | yes |
| <a name="input_force_destroy"></a> [force\_destroy](#input\_force\_destroy) | When destroying this user, destroy even if it has non-Terraform-managed IAM access keys, login profile or MFA devices. Without force\_destroy a user with non-Terraform-managed access keys and login profile will fail to be destroyed. | `bool` | `false` | no |
| <a name="input_iam_access_key_status"></a> [iam\_access\_key\_status](#input\_iam\_access\_key\_status) | Access key status to apply. | `string` | `null` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_password_length"></a> [password\_length](#input\_password\_length) | The length of the generated password | `number` | `20` | no |
| <a name="input_password_reset_required"></a> [password\_reset\_required](#input\_password\_reset\_required) | Whether the user should be forced to reset the generated password on first login. | `bool` | `true` | no |
| <a name="input_path"></a> [path](#input\_path) | Desired path for the IAM user | `string` | `"/"` | no |
| <a name="input_permissions_boundary"></a> [permissions\_boundary](#input\_permissions\_boundary) | The ARN of the policy that is used to set the permissions boundary for the user. | `string` | `""` | no |
| <a name="input_pgp_key"></a> [pgp\_key](#input\_pgp\_key) | Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Used to encrypt password and access key. | `string` | `""` | no |
| <a name="input_policy_arns"></a> [policy\_arns](#input\_policy\_arns) | The list of ARNs of policies directly assigned to the IAM user | `list(string)` | `[]` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_ssh_key_encoding"></a> [ssh\_key\_encoding](#input\_ssh\_key\_encoding) | Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM | `string` | `"SSH"` | no |
| <a name="input_ssh_public_key"></a> [ssh\_public\_key](#input\_ssh\_public\_key) | The SSH public key. The public key must be encoded in ssh-rsa format or PEM format | `string` | `""` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_upload_iam_user_ssh_key"></a> [upload\_iam\_user\_ssh\_key](#input\_upload\_iam\_user\_ssh\_key) | Whether to upload a public ssh key to the IAM user | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_iam_access_key_encrypted_secret"></a> [iam\_access\_key\_encrypted\_secret](#output\_iam\_access\_key\_encrypted\_secret) | The encrypted secret, base64 encoded |
| <a name="output_iam_access_key_encrypted_ses_smtp_password_v4"></a> [iam\_access\_key\_encrypted\_ses\_smtp\_password\_v4](#output\_iam\_access\_key\_encrypted\_ses\_smtp\_password\_v4) | The encrypted secret access key converted into an SES SMTP password by applying AWS's Sigv4 conversion algorithm |
| <a name="output_iam_access_key_id"></a> [iam\_access\_key\_id](#output\_iam\_access\_key\_id) | The access key ID |
| <a name="output_iam_access_key_key_fingerprint"></a> [iam\_access\_key\_key\_fingerprint](#output\_iam\_access\_key\_key\_fingerprint) | The fingerprint of the PGP key used to encrypt the secret |
| <a name="output_iam_access_key_secret"></a> [iam\_access\_key\_secret](#output\_iam\_access\_key\_secret) | The access key secret |
| <a name="output_iam_access_key_ses_smtp_password_v4"></a> [iam\_access\_key\_ses\_smtp\_password\_v4](#output\_iam\_access\_key\_ses\_smtp\_password\_v4) | The secret access key converted into an SES SMTP password by applying AWS's Sigv4 conversion algorithm |
| <a name="output_iam_access_key_status"></a> [iam\_access\_key\_status](#output\_iam\_access\_key\_status) | Active or Inactive. Keys are initially active, but can be made inactive by other means. |
| <a name="output_iam_user_arn"></a> [iam\_user\_arn](#output\_iam\_user\_arn) | The ARN assigned by AWS for this user |
| <a name="output_iam_user_login_profile_encrypted_password"></a> [iam\_user\_login\_profile\_encrypted\_password](#output\_iam\_user\_login\_profile\_encrypted\_password) | The encrypted password, base64 encoded |
| <a name="output_iam_user_login_profile_key_fingerprint"></a> [iam\_user\_login\_profile\_key\_fingerprint](#output\_iam\_user\_login\_profile\_key\_fingerprint) | The fingerprint of the PGP key used to encrypt the password |
| <a name="output_iam_user_login_profile_password"></a> [iam\_user\_login\_profile\_password](#output\_iam\_user\_login\_profile\_password) | The user password |
| <a name="output_iam_user_name"></a> [iam\_user\_name](#output\_iam\_user\_name) | The user's name |
| <a name="output_iam_user_unique_id"></a> [iam\_user\_unique\_id](#output\_iam\_user\_unique\_id) | The unique ID assigned by AWS |
<!-- END_TF_DOCS -->
