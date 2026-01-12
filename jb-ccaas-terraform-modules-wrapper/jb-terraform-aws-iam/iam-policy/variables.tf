variable "application" {
  type        = string
  description = "The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name."
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

variable "create_policy" {
  description = "Whether to create the IAM policy"
  type        = bool
  default     = true
}

variable "name_prefix" {
  description = "IAM policy name prefix"
  type        = string
  default     = null
}

variable "path" {
  description = "The path of the policy in IAM"
  type        = string
  default     = "/"
}

variable "description" {
  description = "The description of the policy"
  type        = string
  default     = "IAM Policy"
}

variable "policy" {
  description = "The path of the policy in IAM (tpl file)"
  type        = string
  default     = ""
}
