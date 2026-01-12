variable "application" {
  type        = string
  description = "The application name of the Key, will be appended with the company, lob, env and region to form a Key name."
}


variable "prefix_company" {
  type        = string
  description = "The prefix company of the Key, will be appended with the company, lob, env and region to form a Key name"
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the Key , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "service" {
  type        = string
  description = "The service identity for the Key. e.g. s3,rds,firehose, etc."
  default     = "default"
}

variable "lob" {
  type        = string
  description = "The lob name of the Key, will be appended with the company, lob, env and region to form a Key name"
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

variable "multi_region" {
  type        = bool
  description = "Multi Region used to create KMS keys in multiple regions"
  default     = false
}

variable "enable_default_policy" {
  type    = bool
  default = true
}

variable "description" {
  type        = string
  default     = "KMS key for encryption"
  description = "The application name of the key , will be appended with the company, lob, env and region to form a key name."
}

variable "primary_key_arn" {
  description = "The primary key arn of a multi-region replica key"
  type        = string
  default     = null
}

variable "create_replica" {
  description = "Determines whether a replica standard CMK will be created (AWS provided material)"
  type        = bool
  default     = false
}

variable "key_statements" {
  description = "A map of IAM policy [statements](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document#statement) for custom permission usage"
  type        = any
  default     = {}
}
