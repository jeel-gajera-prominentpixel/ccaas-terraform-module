variable "application" {
  type        = string
  description = "The application name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name"
}
variable "prefix_region" {
  type        = string
  description = "The prefix region of the cloudtrail, will be appended with the company, lob, env and region to form a acm name."
}
variable "env" {
  type        = string
  description = "Environment name"
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "lob" {
  type        = string
  description = "The lob name of the cloudtrail, will be appended with the company, lob, env and region to form a cloudtrail name"
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "enable_logging" {
  type        = bool
  default     = true
  description = "Enable logging for the trail"
}

variable "enable_log_file_validation" {
  type        = bool
  default     = true
  description = "Specifies whether log file integrity validation is enabled. Creates signed digest for validated contents of logs"
}

variable "include_global_service_events" {
  type        = bool
  default     = false
  description = "Specifies whether the trail is publishing events from global services such as IAM to the log files"
}

variable "is_multi_region_trail" {
  type        = bool
  default     = true
  description = "Specifies whether the trail is created in the current region or in all regions"
}

variable "is_organization_trail" {
  type        = bool
  default     = false
  description = "The trail is an AWS Organizations trail"
}

variable "advanced_event_selector" {
  type = list(object({
    name = optional(string)
    field_selector = list(object({
      field           = string
      ends_with       = optional(list(string))
      not_ends_with   = optional(list(string))
      equals          = optional(list(string))
      not_equals      = optional(list(string))
      starts_with     = optional(list(string))
      not_starts_with = optional(list(string))
    }))
  }))
  description = "Specifies an advanced event selector for enabling data event logging. See: https://www.terraform.io/docs/providers/aws/r/cloudtrail.html for details on this variable"
  default     = []
}

variable "s3_bucket_name" {
  description = "The name of the custom S3 bucket for CloudTrail logs"
  type        = string
}

variable "s3_key_prefix" {
  type        = string
  description = "Prefix for S3 bucket used by Cloudtrail to store logs"
  default     = null
}
