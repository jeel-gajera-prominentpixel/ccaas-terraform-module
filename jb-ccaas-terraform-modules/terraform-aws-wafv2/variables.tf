variable "name" {
  description = "(Required) Friendly name of the WebACL."
  type        = string
}

variable "description" {
  description = "(Optional) Friendly description of the WebACL."
  type        = string
  default     = null
}

variable "scope" {
  description = "(Required) Specifies whether this is for an AWS CloudFront distribution or for a regional application"
  type        = string
}

variable "default_action" {
  description = "(Required) Action to perform if none of the rules contained in the WebACL match."
  type        = string
}

variable "tags" {
  description = "(Optional) Map of key-value pairs to associate with the resource."
  type        = map(string)
  default     = null
}

variable "waf_rules" {
  description = "List of WAF rules"
  type = list(object({
    name                       = string
    priority                   = number
    sampled_requests_enabled   = bool
    cloudwatch_metrics_enabled = bool
    action                     = string # "allow" or "block"
    ip_set_name                = string # Name of the IP set to associate with the rule
    metric_name                = string
  }))
}

variable "waf_ip_sets" {
  description = "List of IP sets to be used in WAF rules"
  type = list(object({
    name               = string
    ip_address_version = string       # "IPV4" or "IPV6"
    addresses_list     = list(string) # List of IP addresses or CIDRs
  }))
}

variable "visibility_config" {
  description = "Configuration for visibility settings of WAF"
  type = object({
    cloudwatch_metrics_enabled = bool
    metric_name                = string
    sampled_requests_enabled   = bool
  })
}

variable "enabled_logging_configuration" {
  description = "(Optional) Whether to create logging configuration."
  type        = bool
  default     = false
}

variable "redacted_fields" {
  type = map(object({
    method        = optional(bool, false)
    uri_path      = optional(bool, false)
    query_string  = optional(bool, false)
    single_header = optional(list(string), null)
  }))
  default     = {}
  description = <<-DOC
    The parts of the request that you want to keep out of the logs.
    You can only specify one of the following: `method`, `query_string`, `single_header`, or `uri_path`

    method:
      Whether to enable redaction of the HTTP method.
      The method indicates the type of operation that the request is asking the origin to perform.
    uri_path:
      Whether to enable redaction of the URI path.
      This is the part of a web request that identifies a resource.
    query_string:
      Whether to enable redaction of the query string.
      This is the part of a URL that appears after a `?` character, if any.
    single_header:
      The list of names of the query headers to redact.
  DOC
}

variable "logging_filter" {
  type        = any
  description = "(Optional) A configuration block that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation."
  default     = null
}
