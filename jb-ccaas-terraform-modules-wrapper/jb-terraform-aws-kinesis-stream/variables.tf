variable "application" {
  type        = string
  description = "The application name of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name"
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the kinesis stream , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "lob" {
  type        = string
  description = "The lob name of the kinesis stream, will be appended with the company, lob, env and region to form a kinesis stream name"
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "shard_count" {
  description = "The number of shards that the stream will use"
  type        = number
  default     = 1
}

variable "retention_period" {
  description = "Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24."
  type        = number
  default     = 24
}

variable "shard_level_metrics" {
  description = "A list of shard-level CloudWatch metrics which can be enabled for the stream."
  type        = list(string)
  default     = []
}

variable "enforce_consumer_deletion" {
  description = "A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error."
  type        = bool
  default     = false
}

variable "encryption_type" {
  description = "The encryption type to use. The only acceptable values are NONE or KMS."
  type        = string
  default     = "NONE"
}

variable "kms_key_id" {
  description = "The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias alias/aws/kinesis."
  type        = string
  default     = ""
}

variable "create_policy_read_only" {
  type        = bool
  default     = true
  description = "Whether to create IAM Policy (ARN) read only of the Stream"
}

variable "create_policy_write_only" {
  type        = bool
  default     = true
  description = "Whether to create IAM Policy (ARN) write only of the Stream"
}

variable "create_policy_admin" {
  type        = bool
  default     = true
  description = "Whether to create IAM Policy (ARN) admin of the Stream"
}
