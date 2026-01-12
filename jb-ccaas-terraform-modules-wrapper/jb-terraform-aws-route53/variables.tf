variable "application" {
  type        = string
  description = "The application name of the route 53, will be appended with the company, lob, env and region to form a route 53 name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the route 53, will be appended with the company, lob, env and region to form a route 53 name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the route 53 , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the route 53, will be appended with the company, lob, env and region to form a route 53 name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "records" {
  description = "List of objects of DNS records"
  type        = any
  default     = []
}

variable "zone_name" {
  description = "Name of DNS zone"
  type        = string
  default     = null
}

variable "zone_id" {
  description = "ID of DNS zone"
  type        = string
  default     = null
}

variable "private_zone" {
  description = "Whether Route53 zone is private or public"
  type        = bool
  default     = false
}
