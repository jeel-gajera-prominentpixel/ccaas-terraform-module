<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13.1 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.63 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 3.63 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_pinpoint_apns_channel.apns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_apns_channel) | resource |
| [aws_pinpoint_apns_sandbox_channel.apns_sandbox](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_apns_sandbox_channel) | resource |
| [aws_pinpoint_apns_voip_channel.apns_voip](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_apns_voip_channel) | resource |
| [aws_pinpoint_apns_voip_sandbox_channel.apns_voip_sandbox](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_apns_voip_sandbox_channel) | resource |
| [aws_pinpoint_app.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_app) | resource |
| [aws_pinpoint_baidu_channel.baidu](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_baidu_channel) | resource |
| [aws_pinpoint_email_channel.email](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_email_channel) | resource |
| [aws_pinpoint_sms_channel.sms](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/pinpoint_sms_channel) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_apns"></a> [apns](#input\_apns) | Provides a Pinpoint APNs Channel resource. | <pre>object({<br>    certificate  = string<br>    private_key  = string<br>    bundle_id    = string<br>    team_id      = string<br>    token_key    = string<br>    token_key_id = string<br>  })</pre> | `null` | no |
| <a name="input_apns_sandbox"></a> [apns\_sandbox](#input\_apns\_sandbox) | Provides a Pinpoint APNs Sandbox Channel resource. | <pre>object({<br>    certificate  = string<br>    private_key  = string<br>    bundle_id    = string<br>    team_id      = string<br>    token_key    = string<br>    token_key_id = string<br>  })</pre> | `null` | no |
| <a name="input_apns_voip"></a> [apns\_voip](#input\_apns\_voip) | Provides a Pinpoint APNs VoIP Channel resource. | <pre>object({<br>    certificate  = string<br>    private_key  = string<br>    bundle_id    = string<br>    team_id      = string<br>    token_key    = string<br>    token_key_id = string<br>  })</pre> | `null` | no |
| <a name="input_apns_voip_sandbox"></a> [apns\_voip\_sandbox](#input\_apns\_voip\_sandbox) | Provides a Pinpoint APNs VoIP Sandbox Channel resource. | <pre>object({<br>    certificate  = string<br>    private_key  = string<br>    bundle_id    = string<br>    team_id      = string<br>    token_key    = string<br>    token_key_id = string<br>  })</pre> | `null` | no |
| <a name="input_baidu"></a> [baidu](#input\_baidu) | Provides a Pinpoint Baidu Channel resource. | <pre>object({<br>    api_key    = string<br>    secret_key = string<br>  })</pre> | `null` | no |
| <a name="input_email"></a> [email](#input\_email) | Provides a Pinpoint Email Channel resource. | <pre>object({<br>    from     = string<br>    identity = string<br>  })</pre> | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | (Required) Project name | `string` | n/a | yes |
| <a name="input_sms"></a> [sms](#input\_sms) | Provides a Pinpoint SMS Channel resource. | <pre>object({<br>    sender     = string<br>    short_code = string<br>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_application_arn"></a> [application\_arn](#output\_application\_arn) | Amazon Resource Name (ARN) of the PinPoint Application. |
| <a name="output_application_id"></a> [application\_id](#output\_application\_id) | The Application ID of the Pinpoint App. |
<!-- END_TF_DOCS -->