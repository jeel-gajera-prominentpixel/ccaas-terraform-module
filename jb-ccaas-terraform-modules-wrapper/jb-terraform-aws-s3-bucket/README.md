# AWS S3 BUCKET Terraform Module

## How to use this module:

### aws s3 bucket basic module usage with the required input variables:
```terraform
module "aws_s3_bucket_basic" {
  source = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket?ref=<version>"

  prefix_company  = "jb"
  lob             = "itsd"
  prefix_region = "usw2"
  application     = "recordings"
  env             = "sandbox"
  tags            = local.tags
}
```

### aws s3 bucket advanced module usage with all the optional input variables:
```terraform
module "aws_s3_bucket_advanced" {
  source = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket?ref=<version>"

  prefix_company = "jb"
  lob            = "itsd"
  prefix_region = "usw2"
  application    = "recordings"
  env            = "sandbox"
  force_destroy  = true
  versioning     = true
  lifecycle_rule = [{
    id     = "retention"
    status = "Enabled"
    expiration = {
      days = 7
    }
    noncurrent_version_expiration = {
      noncurrent_days = 10
    }
  }]
  tags = local.tags
}


module "s3_bucket_notifications" {
  count = var.lambda_trigger ? 1 : 0
  source = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket/modules/notification?ref=<version>"
  bucket = "example-bucket"
  lambda_notifications = {
    lambda1 = {
      function_arn  = "arn:aws:lambda:us-east-1:123456789012:function:my-example-function"
      function_name = "jb-lambda"
      events        = ["s3:ObjectCreated:Put"]
      filter_prefix = "prefix/"
      filter_suffix = ".json"
    }
  }
}
```

### aws s3 bucket module with sse encryption, using a kms key alias:
```terraform
module "main" {
  source        = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-s3-bucket?ref=<version>"
  company       = "jb"
  lob           = "itsd"
  prefix_region = "usw2"
  application   = "recordings"
  env           = "sandbox"
  force_destroy = true
  kms_key_id    = "arn:aws:kms:eu-west-1:xxxx:alias/kms-xxx-xx-xx-xxx"
  tags = {
    "Key"       = "Value",
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
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.99.0 |
| <a name="provider_external"></a> [external](#provider\_external) | 2.3.5 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_s3_bucket"></a> [s3\_bucket](#module\_s3\_bucket) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-s3-bucket | main |
| <a name="module_s3_bucket_notifications"></a> [s3\_bucket\_notifications](#module\_s3\_bucket\_notifications) | git@github.com:jetblueairways/ccaas-terraform-modules.git//terraform-aws-s3-bucket/modules/notification | main |

## Resources

| Name | Type |
|------|------|
| [aws_canonical_user_id.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/canonical_user_id) | data source |
| [aws_cloudfront_log_delivery_canonical_user_id.cloudfront](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/cloudfront_log_delivery_canonical_user_id) | data source |
| [aws_iam_policy_document.bucket_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [external_external.env](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_acl"></a> [acl](#input\_acl) | The canned ACL to apply, defaults to `private`. | `string` | `"private"` | no |
| <a name="input_allowed_kms_key_arn"></a> [allowed\_kms\_key\_arn](#input\_allowed\_kms\_key\_arn) | The ARN of KMS key which should be allowed in PutObject | `string` | `null` | no |
| <a name="input_application"></a> [application](#input\_application) | The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name. | `string` | n/a | yes |
| <a name="input_attach_access_log_delivery_policy"></a> [attach\_access\_log\_delivery\_policy](#input\_attach\_access\_log\_delivery\_policy) | Controls if S3 bucket should have S3 access log delivery policy attached | `bool` | `false` | no |
| <a name="input_attach_analytics_destination_policy"></a> [attach\_analytics\_destination\_policy](#input\_attach\_analytics\_destination\_policy) | Controls if S3 bucket should have bucket analytics destination policy attached. | `bool` | `false` | no |
| <a name="input_attach_deny_incorrect_encryption_headers"></a> [attach\_deny\_incorrect\_encryption\_headers](#input\_attach\_deny\_incorrect\_encryption\_headers) | Controls if S3 bucket should deny incorrect encryption headers policy attached. | `bool` | `false` | no |
| <a name="input_attach_deny_incorrect_kms_key_sse"></a> [attach\_deny\_incorrect\_kms\_key\_sse](#input\_attach\_deny\_incorrect\_kms\_key\_sse) | Controls if S3 bucket policy should deny usage of incorrect KMS key SSE. | `bool` | `false` | no |
| <a name="input_attach_deny_insecure_transport_policy"></a> [attach\_deny\_insecure\_transport\_policy](#input\_attach\_deny\_insecure\_transport\_policy) | Controls if S3 bucket should have deny non-SSL transport policy attached | `bool` | `false` | no |
| <a name="input_attach_deny_unencrypted_object_uploads"></a> [attach\_deny\_unencrypted\_object\_uploads](#input\_attach\_deny\_unencrypted\_object\_uploads) | Controls if S3 bucket should deny unencrypted object uploads policy attached. | `bool` | `false` | no |
| <a name="input_attach_elb_log_delivery_policy"></a> [attach\_elb\_log\_delivery\_policy](#input\_attach\_elb\_log\_delivery\_policy) | Controls if S3 bucket should have ELB log delivery policy attached | `bool` | `false` | no |
| <a name="input_attach_inventory_destination_policy"></a> [attach\_inventory\_destination\_policy](#input\_attach\_inventory\_destination\_policy) | Controls if S3 bucket should have bucket inventory destination policy attached. | `bool` | `false` | no |
| <a name="input_attach_lb_log_delivery_policy"></a> [attach\_lb\_log\_delivery\_policy](#input\_attach\_lb\_log\_delivery\_policy) | Controls if S3 bucket should have ALB/NLB log delivery policy attached | `bool` | `false` | no |
| <a name="input_attach_public_policy"></a> [attach\_public\_policy](#input\_attach\_public\_policy) | Controls if a user defined public bucket policy will be attached (set to `false` to allow upstream to apply defaults to the bucket) | `bool` | `true` | no |
| <a name="input_attach_require_latest_tls_policy"></a> [attach\_require\_latest\_tls\_policy](#input\_attach\_require\_latest\_tls\_policy) | Controls if S3 bucket should require the latest version of TLS | `bool` | `false` | no |
| <a name="input_block_public_acls"></a> [block\_public\_acls](#input\_block\_public\_acls) | Whether Amazon S3 should block public ACLs for this bucket. | `bool` | `true` | no |
| <a name="input_block_public_policy"></a> [block\_public\_policy](#input\_block\_public\_policy) | Whether Amazon S3 should block public bucket policies for this bucket. | `bool` | `true` | no |
| <a name="input_bucket"></a> [bucket](#input\_bucket) | The name of the bucket. If omitted, Terraform will assign a random, unique name. | `string` | `null` | no |
| <a name="input_cors_rule"></a> [cors\_rule](#input\_cors\_rule) | List of maps containing rules for Cross-Origin Resource Sharing. | `any` | `[]` | no |
| <a name="input_create_bucket"></a> [create\_bucket](#input\_create\_bucket) | Controls if S3 bucket should be created | `bool` | `true` | no |
| <a name="input_create_sqs_policy"></a> [create\_sqs\_policy](#input\_create\_sqs\_policy) | Whether to create a policy for SQS permissions or not? | `bool` | `true` | no |
| <a name="input_enable_grant"></a> [enable\_grant](#input\_enable\_grant) | Set to true to enable ACL policy grant | `bool` | `false` | no |
| <a name="input_env"></a> [env](#input\_env) | Environment name | `string` | n/a | yes |
| <a name="input_force_destroy"></a> [force\_destroy](#input\_force\_destroy) | A boolean that indicates all objects should be deleted when deleting the bucket. | `bool` | `false` | no |
| <a name="input_grant"></a> [grant](#input\_grant) | An ACL policy grant. Conflicts with `acl` | `any` | `[]` | no |
| <a name="input_ignore_public_acls"></a> [ignore\_public\_acls](#input\_ignore\_public\_acls) | Whether Amazon S3 should ignore public ACLs for this bucket. | `bool` | `true` | no |
| <a name="input_kms_key_id"></a> [kms\_key\_id](#input\_kms\_key\_id) | The KMS key ID used for the bucket encryption. | `string` | `null` | no |
| <a name="input_lambda_notifications"></a> [lambda\_notifications](#input\_lambda\_notifications) | n/a | `any` | `{}` | no |
| <a name="input_lambda_trigger"></a> [lambda\_trigger](#input\_lambda\_trigger) | n/a | `bool` | `false` | no |
| <a name="input_lifecycle_rule"></a> [lifecycle\_rule](#input\_lifecycle\_rule) | List of lifecycle rules to apply, see [documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_lifecycle_configuration#rule). | `list(any)` | `[]` | no |
| <a name="input_lob"></a> [lob](#input\_lob) | The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Provide resource name if you want to override with wrapper | `string` | `""` | no |
| <a name="input_object_ownership"></a> [object\_ownership](#input\_object\_ownership) | Specifies the object ownership controls on the bucket. Valid values: BucketOwnerPreferred or ObjectWriter. | `string` | `"BucketOwnerEnforced"` | no |
| <a name="input_policy"></a> [policy](#input\_policy) | A valid bucket policy JSON document. | `string` | `null` | no |
| <a name="input_prefix_company"></a> [prefix\_company](#input\_prefix\_company) | The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name | `string` | n/a | yes |
| <a name="input_prefix_region"></a> [prefix\_region](#input\_prefix\_region) | The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name. | `string` | n/a | yes |
| <a name="input_restrict_public_buckets"></a> [restrict\_public\_buckets](#input\_restrict\_public\_buckets) | Whether Amazon S3 should restrict public bucket policies for this bucket. | `bool` | `true` | no |
| <a name="input_s3_logging_bucket_id"></a> [s3\_logging\_bucket\_id](#input\_s3\_logging\_bucket\_id) | Specifies the object ownership controls on the bucket. Valid values: BucketOwnerPreferred or ObjectWriter. | `string` | `null` | no |
| <a name="input_sns_notifications"></a> [sns\_notifications](#input\_sns\_notifications) | Map of S3 bucket notifications to SNS topic | `any` | `{}` | no |
| <a name="input_sqs_notifications"></a> [sqs\_notifications](#input\_sqs\_notifications) | Map of S3 bucket notifications to SQS queue | `any` | `{}` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level. | `map(string)` | `{}` | no |
| <a name="input_versioning"></a> [versioning](#input\_versioning) | Versioning is a means of keeping multiple variants of an object in the same bucket. | `bool` | `true` | no |
| <a name="input_website"></a> [website](#input\_website) | Map containing static web-site hosting or redirect configuration. | `any` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_s3_bucket_arn"></a> [s3\_bucket\_arn](#output\_s3\_bucket\_arn) | The ARN of the bucket. Will be of format arn:aws:s3:::bucketname. |
| <a name="output_s3_bucket_bucket_domain_name"></a> [s3\_bucket\_bucket\_domain\_name](#output\_s3\_bucket\_bucket\_domain\_name) | The bucket domain name. Will be of format bucketname.s3.amazonaws.com. |
| <a name="output_s3_bucket_bucket_prefix"></a> [s3\_bucket\_bucket\_prefix](#output\_s3\_bucket\_bucket\_prefix) | Creates a unique bucket name beginning with the specified prefix. |
| <a name="output_s3_bucket_bucket_regional_domain_name"></a> [s3\_bucket\_bucket\_regional\_domain\_name](#output\_s3\_bucket\_bucket\_regional\_domain\_name) | The bucket region-specific domain name. The bucket domain name including the region name, please refer here for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent redirect issues from CloudFront to S3 Origin URL. |
| <a name="output_s3_bucket_force_destroy"></a> [s3\_bucket\_force\_destroy](#output\_s3\_bucket\_force\_destroy) | hould be deleted from the bucket when the bucket is destroyed so that the bucket can be destroyed without error. These objects are not recoverable. This only deletes objects when the bucket is destroyed. |
| <a name="output_s3_bucket_hosted_zone_id"></a> [s3\_bucket\_hosted\_zone\_id](#output\_s3\_bucket\_hosted\_zone\_id) | The Route 53 Hosted Zone ID for this bucket's region. |
| <a name="output_s3_bucket_id"></a> [s3\_bucket\_id](#output\_s3\_bucket\_id) | The name of the bucket. |
| <a name="output_s3_bucket_lifecycle_configuration_rules"></a> [s3\_bucket\_lifecycle\_configuration\_rules](#output\_s3\_bucket\_lifecycle\_configuration\_rules) | The lifecycle rules of the bucket, if the bucket is configured with lifecycle rules. If not, this will be an empty string. |
| <a name="output_s3_bucket_policy"></a> [s3\_bucket\_policy](#output\_s3\_bucket\_policy) | The policy of the bucket, if the bucket is configured with a policy. If not, this will be an empty string. |
| <a name="output_s3_bucket_region"></a> [s3\_bucket\_region](#output\_s3\_bucket\_region) | The AWS region this bucket resides in. |
| <a name="output_s3_bucket_versioning"></a> [s3\_bucket\_versioning](#output\_s3\_bucket\_versioning) | Configuration of the S3 bucket versioning state. |
| <a name="output_s3_bucket_website_domain"></a> [s3\_bucket\_website\_domain](#output\_s3\_bucket\_website\_domain) | The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. |
| <a name="output_s3_bucket_website_endpoint"></a> [s3\_bucket\_website\_endpoint](#output\_s3\_bucket\_website\_endpoint) | The website endpoint, if the bucket is configured with a website. If not, this will be an empty string. |
<!-- END_TF_DOCS -->

# How To contribute

## Install dependencies

- [pre-commit](https://pre-commit.com)
- [pre-commit-hooks](https://github.com/pre-commit/pre-commit-hooks)
- [pre-commit-terraform](https://github.com/antonbabenko/pre-commit-terraform)
- [terraform-docs](https://github.com/terraform-docs/terraform-docs)
- [tflint](https://github.com/terraform-linters/tflint)

#### MacOS

```bash
brew install pre-commit terraform-docs tflint
```

#### Ubuntu

```bash
python3 -m pip install --upgrade pip
pip3 install --no-cache-dir pre-commit
curl -L "$(curl -s https://api.github.com/repos/terraform-docs/terraform-docs/releases/latest | grep -o -E -m 1 "https://.+?-linux-amd64.tar.gz")" > terraform-docs.tgz && tar -xzf terraform-docs.tgz && rm terraform-docs.tgz && chmod +x terraform-docs && sudo mv terraform-docs /usr/bin/
curl -L "$(curl -s https://api.github.com/repos/terraform-linters/tflint/releases/latest | grep -o -E -m 1 "https://.+?_linux_amd64.zip")" > tflint.zip && unzip tflint.zip && rm tflint.zip && sudo mv tflint /usr/bin/
```

### Install the pre-commit hook globally

> Note: not needed if you use the Docker image

```bash
DIR=~/.git-template
git config --global init.templateDir ${DIR}
pre-commit init-templatedir -t pre-commit ${DIR}
```

### Run the pre-commit

Execute this command to run `pre-commit` on all files in the repository (not only changed files):

```bash
pre-commit run -a
```
