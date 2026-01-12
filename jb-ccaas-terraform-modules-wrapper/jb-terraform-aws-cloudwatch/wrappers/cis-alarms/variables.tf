variable "defaults" {
  description = "Map of default values which will be used for each item."
  type        = any
  default     = {}
}

variable "items" {
  description = "Maps of items to create a wrapper from. Values are passed through to the module."
  type        = any
  default     = {}
}


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
