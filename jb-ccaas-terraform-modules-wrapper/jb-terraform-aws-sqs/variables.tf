variable "application" {
  type        = string
  description = "The application name of the sqs, will be appended with the company, lob, env and region to form a sqs name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the sqs, will be appended with the company, lob, env and region to form a sqs name."
}


variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws sqs , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the sqs, will be appended with the company, lob, env and region to form a sqs name."
}


variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "fifo_queue" {
  description = "Boolean designating a FIFO queue"
  type        = bool
  default     = false
}

variable "create_dlq" {
  description = "Determines whether to create SQS dead letter queue"
  type        = bool
  default     = false
}

variable "redrive_policy" {
  description = "The JSON policy to set up the Dead Letter Queue, see AWS docs. Note: when specifying maxReceiveCount, you must specify it as an integer (5), and not a string (\"5\")"
  type        = any
  default     = {}
}

variable "sqs_managed_sse_enabled" {
  description = "Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys"
  type        = bool
  default     = true
}

variable "create_queue_policy" {
  description = "Whether to create SQS queue policy"
  type        = bool
  default     = false
}

variable "create_dlq_redrive_allow_policy" {
  description = "Determines whether to create a redrive allow policy for the dead letter queue."
  type        = bool
  default     = true
}

variable "dlq_redrive_allow_policy" {
  description = "The JSON policy to set up the Dead Letter Queue redrive permission, see AWS docs."
  type        = any
  default     = {}
}

variable "dlq_queue_policy_statements" {
  description = "A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage"
  type        = any
  default     = {}
}

variable "kms_master_key_id" {
  description = "The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK"
  type        = string
  default     = null
}

variable "kms_data_key_reuse_period_seconds" {
  description = "The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours)"
  type        = number
  default     = null
}

variable "queue_policy_statements" {
  description = "A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage"
  type        = any
  default     = {}
}
