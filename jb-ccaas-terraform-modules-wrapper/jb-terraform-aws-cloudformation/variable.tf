variable "application" {
  type        = string
  description = "The application name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the cloudfront, will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}


variable "lob" {
  type        = string
  description = "The lob name of the cloudfront, will be appended with the company, lob, env and region to form a cloudfront name."
}

variable "enabled" {
  description = "Flag to determine if the stack should be created"
  type        = bool
  default     = true
}

variable "name" {
  description = "(Required) Stack name"
  type        = string
}

variable "template_body" {
  description = "(Optional) Structure containing the template body (max size: 51,200 bytes)"
  type        = string
  default     = null
}

variable "template_url" {
  description = "(Optional) Location of a file containing the template body (max size: 460,800 bytes)"
  type        = string
  default     = null
}

variable "capabilities" {
  description = "(Optional) A list of capabilities. Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, or CAPABILITY_AUTO_EXPAND"
  type        = set(string) # Changed from list to set as per AWS docs
  default     = null
  validation {
    condition     = var.capabilities == null ? true : alltrue([for cap in var.capabilities : contains(["CAPABILITY_IAM", "CAPABILITY_NAMED_IAM", "CAPABILITY_AUTO_EXPAND"], cap)])
    error_message = "Valid values for capabilities are: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, or CAPABILITY_AUTO_EXPAND"
  }
}

variable "disable_rollback" {
  description = "(Optional) Set to true to disable rollback of the stack if stack creation failed. Conflicts with on_failure"
  type        = bool
  default     = null
}

variable "notification_arns" {
  description = "(Optional) A list of SNS topic ARNs to publish stack related events"
  type        = set(string) # Changed from list to set
  default     = null
}

variable "on_failure" {
  description = "(Optional) Action to be taken if stack creation fails. This must be one of: DO_NOTHING, ROLLBACK, or DELETE. Conflicts with disable_rollback"
  type        = string
  default     = null
  validation {
    condition     = var.on_failure == null ? true : contains(["DO_NOTHING", "ROLLBACK", "DELETE"], var.on_failure)
    error_message = "on_failure must be one of: DO_NOTHING, ROLLBACK, or DELETE"
  }
}

variable "parameters" {
  description = "(Optional) A map of Parameter structures that specify input parameters for the stack"
  type        = map(string)
  default     = null
}

variable "policy_body" {
  description = "(Optional) Structure containing the stack policy body. Conflicts with policy_url"
  type        = string
  default     = null
}

variable "policy_url" {
  description = "(Optional) Location of a file containing the stack policy. Conflicts with policy_body"
  type        = string
  default     = null
}

variable "tags" {
  description = "(Optional) Map of resource tags to associate with this stack"
  type        = map(string)
  default     = null
}

variable "iam_role_arn" {
  description = "(Optional) The ARN of an IAM role that AWS CloudFormation assumes to create the stack"
  type        = string
  default     = null
}

variable "timeout_in_minutes" {
  description = "(Optional) The amount of time that can pass before the stack status becomes CREATE_FAILED"
  type        = number
  default     = null
}
