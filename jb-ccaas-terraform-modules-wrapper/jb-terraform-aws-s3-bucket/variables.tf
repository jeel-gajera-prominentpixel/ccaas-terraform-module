variable "application" {
  type        = string
  description = "The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name"
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name."
}

variable "lob" {
  type        = string
  description = "The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name"
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "kms_key_id" {
  type        = string
  nullable    = true
  default     = null
  description = "The KMS key ID used for the bucket encryption."
  validation {
    condition     = var.kms_key_id == null || can(regex("arn:aws:kms:.*:\\d{12}:alias\\/.*s|key\\/.*", var.kms_key_id))
    error_message = "The ARN supplied is not valid, see https://docs.aws.amazon.com/service-authorization/latest/reference/list_awskeymanagementservice.html#awskeymanagementservice-key"
  }
}

variable "force_destroy" {
  type        = bool
  default     = false
  description = "A boolean that indicates all objects should be deleted when deleting the bucket."
}

variable "acl" {
  type        = string
  default     = "private"
  description = "The canned ACL to apply, defaults to `private`."
}

variable "block_public_acls" {
  type        = bool
  default     = true
  description = "Whether Amazon S3 should block public ACLs for this bucket."
}

variable "block_public_policy" {
  type        = bool
  default     = true
  description = "Whether Amazon S3 should block public bucket policies for this bucket."
}

variable "ignore_public_acls" {
  type        = bool
  default     = true
  description = "Whether Amazon S3 should ignore public ACLs for this bucket."
}

variable "restrict_public_buckets" {
  type        = bool
  default     = true
  description = "Whether Amazon S3 should restrict public bucket policies for this bucket."
}

variable "policy" {
  type        = string
  default     = null
  description = "A valid bucket policy JSON document."
}

variable "versioning" {
  type        = bool
  default     = true
  description = "Versioning is a means of keeping multiple variants of an object in the same bucket."
}

variable "lifecycle_rule" {
  description = "List of lifecycle rules to apply, see [documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_lifecycle_configuration#rule)."
  type        = list(any)
  default     = []
}
variable "data_classification" {
  description = "Data classification tag - REQUIRED for all S3 buckets. Valid values: Public, Internal Use, Confidential, Personally Identifiable Information PII, Sensitive PII SPII, Restricted, PCI"
  type        = string
  validation {
    condition = contains([
      "Public",
      "Internal Use",
      "Confidential",
      "Personally Identifiable Information PII",
      "Sensitive PII SPII",
      "Restricted",
      "PCI"
    ], var.data_classification)
    error_message = "The data_classification must be one of: Public, Internal Use, Confidential, Personally Identifiable Information PII, Sensitive PII SPII, Restricted, PCI."
  }
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
  validation {
    condition     = !contains(keys(var.tags), "data_classification")
    error_message = "The 'data_classification' tag should not be specified in the tags variable. Use the dedicated 'data_classification' variable instead."
  }
}

variable "object_ownership" {
  description = "Specifies the object ownership controls on the bucket. Valid values: BucketOwnerPreferred or ObjectWriter."
  type        = string
  default     = "BucketOwnerEnforced"
}

variable "s3_logging_bucket_id" {
  description = "Specifies the object ownership controls on the bucket. Valid values: BucketOwnerPreferred or ObjectWriter."
  type        = string
  default     = null
}


variable "bucket" {
  description = "The name of the bucket. If omitted, Terraform will assign a random, unique name."
  type        = string
  default     = null
}

variable "lambda_notifications" {
  type    = any
  default = {}
}


variable "lambda_trigger" {
  type    = bool
  default = false
}

variable "sqs_notifications" {
  description = "Map of S3 bucket notifications to SQS queue"
  type        = any
  default     = {}
}

variable "sns_notifications" {
  description = "Map of S3 bucket notifications to SNS topic"
  type        = any
  default     = {}
}

variable "create_sqs_policy" {
  description = "Whether to create a policy for SQS permissions or not?"
  type        = bool
  default     = true
}

variable "enable_grant" {
  description = "Set to true to enable ACL policy grant"
  type        = bool
  default     = false
}
variable "grant" {
  description = "An ACL policy grant. Conflicts with `acl`"
  type        = any
  default     = []
}

variable "website" {
  description = "Map containing static web-site hosting or redirect configuration."
  type        = any # map(string)
  default     = {}
}

variable "cors_rule" {
  description = "List of maps containing rules for Cross-Origin Resource Sharing."
  type        = any
  default     = []
}

variable "create_bucket" {
  description = "Controls if S3 bucket should be created"
  type        = bool
  default     = true
}

variable "attach_elb_log_delivery_policy" {
  description = "Controls if S3 bucket should have ELB log delivery policy attached"
  type        = bool
  default     = false
}

variable "attach_lb_log_delivery_policy" {
  description = "Controls if S3 bucket should have ALB/NLB log delivery policy attached"
  type        = bool
  default     = false
}

variable "attach_access_log_delivery_policy" {
  description = "Controls if S3 bucket should have S3 access log delivery policy attached"
  type        = bool
  default     = false
}

variable "attach_deny_insecure_transport_policy" {
  description = "Controls if S3 bucket should have deny non-SSL transport policy attached"
  type        = bool
  default     = false
}

variable "attach_require_latest_tls_policy" {
  description = "Controls if S3 bucket should require the latest version of TLS"
  type        = bool
  default     = false
}

variable "attach_public_policy" {
  description = "Controls if a user defined public bucket policy will be attached (set to `false` to allow upstream to apply defaults to the bucket)"
  type        = bool
  default     = true
}

variable "attach_inventory_destination_policy" {
  description = "Controls if S3 bucket should have bucket inventory destination policy attached."
  type        = bool
  default     = false
}

variable "attach_analytics_destination_policy" {
  description = "Controls if S3 bucket should have bucket analytics destination policy attached."
  type        = bool
  default     = false
}

variable "attach_deny_incorrect_encryption_headers" {
  description = "Controls if S3 bucket should deny incorrect encryption headers policy attached."
  type        = bool
  default     = false
}

variable "attach_deny_incorrect_kms_key_sse" {
  description = "Controls if S3 bucket policy should deny usage of incorrect KMS key SSE."
  type        = bool
  default     = false
}

variable "allowed_kms_key_arn" {
  description = "The ARN of KMS key which should be allowed in PutObject"
  type        = string
  default     = null
}

variable "attach_deny_unencrypted_object_uploads" {
  description = "Controls if S3 bucket should deny unencrypted object uploads policy attached."
  type        = bool
  default     = false
}

variable "replication_configuration" {
  description = "Map containing cross-region replication configuration."
  type        = any
  default     = {}
}
