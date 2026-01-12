variable "application" {
  type        = string
  description = "The application name of the step function, will be appended with the company, lob, env and region to form a step function name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the step function, will be appended with the company, lob, env and region to form a step function name."
}


variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws step function, will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the step function, will be appended with the company, lob, env and region to form a step function name."
}


variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = null
}


variable "type" {
  description = "Determines whether a Standard or Express state machine is created. The default is STANDARD. Valid Values: STANDARD | EXPRESS"
  type        = string
  default     = "STANDARD"

  validation {
    condition     = contains(["STANDARD", "EXPRESS"], upper(var.type))
    error_message = "Step Function type must be one of the following (STANDARD | EXPRESS)."
  }
}


variable "publish" {
  description = "Determines whether to set a version of the state machine when it is created."
  type        = bool
  default     = false
}

variable "logging_configuration" {
  description = "Defines what execution history events are logged and where they are logged"
  type        = map(string)
  default     = {}
}

variable "service_integrations" {
  description = "Map of AWS service integrations to allow in IAM role policy"
  type        = any
  default     = {}
}

variable "definition" {
  description = "The Amazon States Language definition of the Step Function"
  type        = string
  default     = ""
}

variable "attach_policy_json" {
  description = "Controls whether policy_json should be added to IAM role"
  type        = bool
  default     = false
}
variable "policy_json" {
  description = "An additional policy document as JSON to attach to IAM role"
  type        = string
  default     = null
}

variable "cloudwatch_log_group_name" {
  description = "Name of Cloudwatch Logs group name to use."
  type        = string
  default     = null
}

variable "use_existing_cloudwatch_log_group" {
  description = "Whether to use an existing CloudWatch log group or create new"
  type        = bool
  default     = false
}
