variable "application" {
  type        = string
  description = "The application name of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws pinoint , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the aws pinoint, will be appended with the company, lob, env and region to form a aws pinoint name."
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "sms" {
  type = object({
    sender     = string
    short_code = string
  })
  default     = null
  description = "Provides a Pinpoint SMS Channel resource."
}

variable "email" {
  type = object({
    from     = string
    identity = string
  })
  default     = null
  description = "Provides a Pinpoint Email Channel resource."
}
