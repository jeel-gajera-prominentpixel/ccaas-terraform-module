# Common variables.
variable "project" {
  description = "The name of the project associated with this resource."
  type        = string
  default     = ""
}

variable "application" {
  description = "Application name"
  type        = string
  default     = ""
}

variable "company_prefix" {
  description = "Company prefix"
  type        = string
  default     = ""
}

variable "company" {
  description = "Company name"
  type        = string
  default     = ""
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = ""
}

variable "region_suffix" {
  description = "Region suffix"
  type        = string
  default     = ""
}

variable "region" {
  description = "Region of the resource."
  type        = string
  default     = ""
}

variable "repository_url" {
  description = "The URL of the code repository."
  type        = string
  default     = null
}

variable "custom_tags" {
  description = "Additional custom tags"
  type        = map(string)
  default     = {}
}

variable "tags" {
  description = "Resource tags"
  type        = map(string)
  default     = {}
}

variable "name" {
  description = "The name of the resource."
  type        = string
  default     = ""
}

variable "bus_name" {
  description = "EventBridge Bus name"
  type        = string
  default     = "default"
}

variable "bus_description" {
  description = "Event bus description"
  type        = string
  default     = null
}

variable "create_bus" {
  description = "Create EventBridge Bus"
  type        = bool
  default     = true
}

variable "event_source_name" {
  description = "Partner event source name"
  type        = string
  default     = null
}

variable "kms_key_identifier" {
  description = "KMS key identifier"
  type        = string
  default     = null
}

variable "schedule_groups" {
  description = "Schedule group definitions"
  type        = any
  default     = {}
}

variable "schedules" {
  description = "Schedule definitions"
  type        = map(any)
  default     = {}
}

variable "schedule_group_timeouts" {
  description = "Schedule group timeouts"
  type        = map(string)
  default     = {}
}

variable "create_schedule_groups" {
  type    = bool
  default = true
}

variable "create_schedules" {
  type    = bool
  default = true
}

variable "append_schedule_group_postfix" {
  type    = bool
  default = true
}

variable "append_schedule_postfix" {
  type    = bool
  default = true
}

variable "create_log_delivery" {
  description = "Enable log delivery"
  type        = bool
  default     = true
}

variable "create_log_delivery_source" {
  description = "Enable log delivery source"
  type        = bool
  default     = true
}


variable "create_role" {
  type    = bool
  default = false
}

variable "role_name" {
  type    = string
  default = null
}

variable "role_description" {
  type    = string
  default = null
}

variable "role_path" {
  type    = string
  default = null
}

variable "role_force_detach_policies" {
  type    = bool
  default = true
}

variable "role_permissions_boundary" {
  type    = string
  default = null
}

variable "role_tags" {
  type    = map(string)
  default = {}
}

variable "trusted_entities" {
  description = "Trusted IAM entities"
  type        = list(string)
  default     = []
}

variable "policies" {
  description = "Policy ARNs"
  type        = list(string)
  default     = []
}

variable "policy" {
  type    = string
  default = null
}

variable "policy_path" {
  type    = string
  default = null
}

variable "policy_json" {
  type    = string
  default = null
}

variable "policy_jsons" {
  type    = list(string)
  default = []
}

variable "policy_statements" {
  type    = any
  default = {}
}

variable "number_of_policies" {
  type    = number
  default = 0
}

variable "number_of_policy_jsons" {
  type    = number
  default = 0
}

variable "attach_api_destination_policy" { 
  type = bool 
  default = false 
}

variable "attach_cloudwatch_policy" { 
  type = bool 
  default = false 
}

variable "attach_ecs_policy" { 
  type = bool 
  default = false 
}

variable "attach_kinesis_firehose_policy" { 
  type = bool 
  default = false 
}

variable "attach_kinesis_policy" { 
  type = bool 
  default = false 
}

variable "attach_lambda_policy" { 
  type = bool 
  default = false 
}

variable "attach_policies" { 
  type = bool 
  default = false 
}

variable "attach_policy" { 
  type = bool 
  default = false 
}

variable "attach_policy_json" { 
  type = bool 
  default = false 
}

variable "attach_policy_jsons" { 
  type = bool 
  default = false 
}

variable "attach_policy_statements" { 
  type = bool 
  default = false 
}

variable "attach_sfn_policy" { 
  type = bool 
  default = false 
}

variable "attach_sns_policy" { 
  type = bool 
  default = false 
}

variable "attach_sqs_policy" { 
  type = bool 
  default = false 
}

variable "attach_tracing_policy" { 
  type = bool 
  default = false 
}

variable "rules" {
  type    = map(any)
  default = {}
}

variable "targets" {
  type    = any
  default = {}
}

variable "create_rules" {
  type    = bool
  default = false
}

variable "create_targets" {
  type    = bool
  default = false
}

variable "append_rule_postfix" {
  type    = bool
  default = false
}

variable "pipes" { 
  type = any 
  default = {} 
}
variable "create_pipes" { 
  type = bool 
  default = false
}
variable "append_pipe_postfix" { 
  type = bool
  default = false
}

variable "api_destinations" { 
  type = map(any)
  default = {}
}
variable "create_api_destinations" { 
  type = bool 
  default = false 
}

variable "connections" { 
  type = any 
  default = {}
}
variable "create_connections" { 
  type = bool 
  default = false 
}
variable "append_connection_postfix" { 
  type = bool 
  default = false 
}
variable "append_destination_postfix" { 
  type = bool 
  default = false 
}

variable "archives" { 
  type = map(any) 
  default = {} 
}
variable "create_archives" { 
  type = bool 
  default = false 
}

variable "permissions" { 
  type = map(any) 
  default = {} 
}
variable "create_permissions" { 
  type = bool 
  default = false 
}

variable "cloudwatch_target_arns" { 
  type = list(string) 
  default = [] 
}

variable "ecs_target_arns"        { 
  type = list(string) 
  default = [] 
}

variable "kinesis_firehose_target_arns" { 
  type = list(string) 
  default = [] 
}

variable "kinesis_target_arns"    { 
  type = list(string) 
  default = [] 
}

variable "lambda_target_arns"     { 
  type = list(string) 
  default = [] 
}

variable "sfn_target_arns"        { 
  type = list(string) 
  default = [] 
}

variable "sns_target_arns"        { 
  type = list(string) 
  default = [] 
}

variable "sqs_target_arns"        { 
  type = list(string) 
  default = [] 
}

