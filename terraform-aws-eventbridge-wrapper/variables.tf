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

variable "custom_tags" {
  description = "Additional custom tags to apply to resources, supplementing the default tags."
  type        = map(string)
  default     = {}
}

variable "repository_url" {
  description = "The URL of the code repository."
  type        = string
  default     = null
}

# Resource specific variables.
variable "name" {
  description = "The name of the resource."
  type        = string
  default     = ""
}

variable "create_bus" {
  description = "Specifies whether to create an EventBridge Bus resource."
  type        = bool
  default     = true
}

variable "bus_name" {
  description = "The unique name of the EventBridge Bus."
  type        = string
  default     = "default"
}

variable "rules" {
  description = "A map of objects defining EventBridge Rule configurations."
  type        = map(any)
  default     = {}
}

variable "targets" {
  description = "A map of objects defining EventBridge Target configurations."
  type        = any
  default     = {}
}

variable "create" {
  description = "Configuration settings for resource creation."
  type        = any
  default     = {}
}

variable "attach" {
  description = "Configuration settings for attaching the resources."
  type        = any
  default     = {}
}

variable "target_arns_configuration" {
  description = "Configuration for specifying the ARNs of target resources."
  type        = any
  default     = {}
}

variable "append_postfix_configuration" {
  description = "Configuration for appending a relevant postfix to resource names."
  type        = any
  default     = {}
}

variable "role_configuration" {
  description = "Configuration settings related to IAM roles."
  type        = any
  default     = {}
}

variable "policies" {
  description = "A list of policy statement ARNs to attach to an IAM role."
  type        = list(string)
  default     = []
}

variable "policy_statements" {
  description = "A map of dynamic policy statements to attach to an IAM role."
  type        = any
  default     = {}
}

variable "policy_json" {
  description = "An additional policy document in JSON format to attach to an IAM role."
  type        = string
  default     = null
}

variable "policy_jsons" {
  description = "A list of additional policy documents in JSON format to attach to an IAM role."
  type        = list(string)
  default     = []
}

variable "policy" {
  description = "An additional policy document to attach to an IAM role."
  type        = string
  default     = null
}

variable "pipes" {
  description = "A map of objects defining EventBridge Pipe configurations."
  type        = any
  default     = {}
}

variable "api_destinations" {
  description = "A map of objects with EventBridge Destination definitions."
  type        = map(any)
  default     = {}
}

variable "number_of_policy_jsons" {
  description = "The number of policy documents in JSON format to attach to an IAM role."
  type        = number
  default     = 0
}

variable "number_of_policies" {
  description = "The number of policies to attach to an IAM role."
  type        = number
  default     = 0
}

variable "connections" {
  description = "A map of objects defining EventBridge Connection configurations."
  type        = any
  default     = {}
}

variable "schedule_groups" {
  description = "A map of objects defining EventBridge Schedule Group configurations."
  type        = any
  default     = {}
}

variable "schedules" {
  description = "A map of objects defining EventBridge Schedule configurations."
  type        = map(any)
  default     = {}
}

variable "schedule_group_timeouts" {
  description = "A map of objects defining create and delete timeouts for EventBridge Schedule Groups."
  type        = map(string)
  default     = {}
}

variable "permissions" {
  description = "A map of objects defining EventBridge Permission configurations."
  type        = map(any)
  default     = {}
}

variable "tags" {
  description = "A mapping of tags to assign to all resources."
  type        = map(string)
  default     = {}
}

variable "policy_path" {
  description = "Path of IAM policy to use for EventBridge."
  type        = string
  default     = null
}

variable "archives" {
  description = "A map of objects with the EventBridge Archive definitions."
  type        = map(any)
  default     = {}
}

variable "bus_description" {
  description = "Event bus description."
  type        = string
  default     = null
}

variable "ecs_pass_role_resources" {
  description = "List of approved roles to be passed."
  type        = list(string)
  default     = []
}

variable "event_source_name" {
  description = "The partner event source that the new event bus will be matched with. Must match name."
  type        = string
  default     = null
}

variable "kms_key_identifier" {
  description = "The identifier of the AWS KMS customer managed key for EventBridge to use."
  type        = string
  default     = null
}

variable "schemas_discoverer_description" {
  description = "Default schemas discoverer description."
  type        = string
  default     = "Auto schemas discoverer event"
}

variable "sns_kms_arns" {
  description = "The Amazon Resource Name (ARN) of the AWS KMS's configured for AWS SNS."
  type        = list(string)
  default     = ["*"]
}

variable "trusted_entities" {
  description = "Additional trusted entities for assuming roles (trust relationship)."
  type        = list(string)
  default     = []
}

variable "create_log_delivery" {
  description = "Whether to create log delivery resources"
  type        = bool
  default     = true
}

variable "create_log_delivery_source" {
  description = "Whether to create log delivery source"
  type        = bool
  default     = true
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

############################
# Policies
############################


variable "policy_statements" {
  type    = any
  default = {}
}


############################
# API Destinations & Connections
############################

variable "create_api_destinations" {
  type    = bool
  default = false
}

variable "create_connections" {
  type    = bool
  default = false
}

variable "append_connection_postfix" {
  type    = bool
  default = false
}

variable "append_destination_postfix" {
  type    = bool
  default = false
}

############################
# Archives
############################

variable "create_archives" {
  type    = bool
  default = false
}

############################
# Permissions
############################

variable "create_permissions" {
  type    = bool
  default = false
}

############################
# Pipes
############################


variable "create_pipes" {
  type    = bool
  default = false
}

variable "append_pipe_postfix" {
  type    = bool
  default = false
}

############################
# Schedules
############################

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

############################
# Schema Discoverer
############################

variable "create_schemas_discoverer" {
  type    = bool
  default = false
}


############################
# IAM Role Configuration
############################

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

############################
# Attach Policy Flags
############################

variable "attach_api_destination_policy" {
  type    = bool
  default = false
}

variable "attach_cloudwatch_policy" {
  type    = bool
  default = false
}

variable "attach_ecs_policy" {
  type    = bool
  default = false
}

variable "attach_kinesis_firehose_policy" {
  type    = bool
  default = false
}

variable "attach_kinesis_policy" {
  type    = bool
  default = false
}

variable "attach_lambda_policy" {
  type    = bool
  default = false
}

variable "attach_policies" {
  type    = bool
  default = false
}

variable "attach_policy" {
  type    = bool
  default = false
}

variable "attach_policy_json" {
  type    = bool
  default = false
}

variable "attach_policy_jsons" {
  type    = bool
  default = false
}

variable "attach_sfn_policy" {
  type    = bool
  default = false
}

variable "attach_sns_policy" {
  type    = bool
  default = false
}

variable "attach_sqs_policy" {
  type    = bool
  default = false
}

variable "attach_tracing_policy" {
  type    = bool
  default = false
}

############################
# Target ARNs
############################

variable "cloudwatch_target_arns" {
  type    = list(string)
  default = []
}

variable "ecs_target_arns" {
  type    = list(string)
  default = []
}

variable "kinesis_firehose_target_arns" {
  type    = list(string)
  default = []
}

variable "kinesis_target_arns" {
  type    = list(string)
  default = []
}

variable "lambda_target_arns" {
  type    = list(string)
  default = []
}

variable "sfn_target_arns" {
  type    = list(string)
  default = []
}

variable "sns_target_arns" {
  type    = list(string)
  default = []
}

variable "sqs_target_arns" {
  type    = list(string)
  default = []
}