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

variable "parameter_write" {
  type        = list(map(string))
  description = "List of maps with the parameter values to write to SSM Parameter Store"
  default     = []
}

variable "parameter_read" {
  type        = list(string)
  description = "List of parameters to read from SSM. These must already exist otherwise an error is returned. Can be used with `parameter_write` as long as the parameters are different."
  default     = []
}

variable "kms_arn" {
  type        = string
  default     = ""
  description = "The ARN of a KMS key used to encrypt and decrypt SecretString values"
}

variable "environment" {
  description = "ID element. Usually used for region e.g. 'uw2', 'us-west-2', OR role 'prod', 'staging', 'dev', 'UAT'"
  type        = string
  default     = null
}

variable "regex_replace_chars" {
  description = "Terraform regular expression (regex) string.Characters matching the regex will be removed from the ID elements.If not set, /[^a-zA-Z0-9-]/ is used to remove all characters other than hyphens, letters and digits."
  type        = string
  default     = null
}
