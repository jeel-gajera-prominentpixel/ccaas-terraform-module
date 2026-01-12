variable "application" {
  type        = string
  description = "The application name of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name."
}


variable "prefix_region" {
  type        = string
  description = "The prefix region of the eventbridge , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "lob" {
  type        = string
  description = "The lob name of the eventbridge, will be appended with the company, lob, env and region to form a eventbridge name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "create_bus" {
  description = "Controls whether EventBridge Bus resource should be created"
  type        = bool
  default     = true
}

variable "create_schemas_discoverer" {
  description = "Controls whether default schemas discoverer should be created"
  type        = bool
  default     = false
}

variable "attach_tracing_policy" {
  description = "Controls whether X-Ray tracing policy should be added to IAM role for EventBridge"
  type        = bool
  default     = false
}

variable "attach_kinesis_policy" {
  description = "Controls whether the Kinesis policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "kinesis_target_arns" {
  description = "The Amazon Resource Name (ARN) of the Kinesis Streams you want to use as EventBridge targets"
  type        = list(string)
  default     = []
}

variable "attach_sfn_policy" {
  description = "Controls whether the StepFunction policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "sfn_target_arns" {
  description = "The Amazon Resource Name (ARN) of the StepFunctions you want to use as EventBridge targets"
  type        = list(string)
  default     = []
}

variable "attach_sqs_policy" {
  description = "Controls whether the SQS policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "sqs_target_arns" {
  description = "The Amazon Resource Name (ARN) of the AWS SQS Queues you want to use as EventBridge targets"
  type        = list(string)
  default     = []
}

variable "attach_cloudwatch_policy" {
  description = "Controls whether the Cloudwatch policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "cloudwatch_target_arns" {
  description = "The Amazon Resource Name (ARN) of the Cloudwatch Log Streams you want to use as EventBridge targets"
  type        = list(string)
  default     = []
}

variable "append_rule_postfix" {
  description = "Controls whether to append '-rule' to the name of the rule"
  type        = bool
  default     = true
}

variable "attach_ecs_policy" {
  description = "Controls whether the ECS policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "ecs_target_arns" {
  description = "The Amazon Resource Name (ARN) of the AWS ECS Tasks you want to use as EventBridge targets"
  type        = list(string)
  default     = []
}

variable "rules" {
  description = "A map of objects with EventBridge Rule definitions."
  type        = map(any)
  default     = {}
}

variable "targets" {
  description = "A map of objects with EventBridge Target definitions."
  type        = any
  default     = {}
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

variable "attach_policy_jsons" {
  description = "Controls whether policy_jsons should be added to IAM role"
  type        = bool
  default     = false
}

variable "policy_jsons" {
  description = "List of additional policy documents as JSON to attach to IAM role"
  type        = list(string)
  default     = []
}

variable "number_of_policy_jsons" {
  description = "Number of policies JSON to attach to IAM role"
  type        = number
  default     = 0
}

variable "attach_policies" {
  description = "Controls whether list of policies should be added to IAM role"
  type        = bool
  default     = false
}

variable "policies" {
  description = "List of policy statements ARN to attach to IAM role"
  type        = list(string)
  default     = []
}

variable "number_of_policies" {
  description = "Number of policies to attach to IAM role"
  type        = number
  default     = 0
}

variable "attach_policy_statements" {
  description = "Controls whether policy_statements should be added to IAM role"
  type        = bool
  default     = false
}

variable "policy_statements" {
  description = "Map of dynamic policy statements to attach to IAM role"
  type        = any
  default     = {}
}

variable "api_destinations" {
  description = "A map of objects with EventBridge Destination definitions."
  type        = map(any)
  default     = {}
}

variable "create_api_destinations" {
  description = "Controls whether EventBridge Destination resources should be created"
  type        = bool
  default     = false
}

variable "attach_api_destination_policy" {
  description = "Controls whether the API Destination policy should be added to IAM role for EventBridge Target"
  type        = bool
  default     = false
}

variable "create_connections" {
  description = "Controls whether EventBridge Connection resources should be created"
  type        = bool
  default     = false
}

variable "connections" {
  description = "A map of objects with EventBridge Connection definitions."
  type        = any
  default     = {}
}


######
# IAM
######

variable "role_name" {
  description = "Name of IAM role to use for EventBridge"
  type        = string
  default     = null
}

variable "role_description" {
  description = "Description of IAM role to use for EventBridge"
  type        = string
  default     = null
}

variable "role_path" {
  description = "Path of IAM role to use for EventBridge"
  type        = string
  default     = null
}

variable "role_force_detach_policies" {
  description = "Specifies to force detaching any policies the IAM role has before destroying it."
  type        = bool
  default     = true
}

variable "role_permissions_boundary" {
  description = "The ARN of the policy that is used to set the permissions boundary for the IAM role used by EventBridge"
  type        = string
  default     = null
}

variable "role_tags" {
  description = "A map of tags to assign to IAM role"
  type        = map(string)
  default     = {}
}

variable "create_pipes" {
  description = "Controls whether EventBridge Pipes resources should be created"
  type        = bool
  default     = true
}


variable "pipes" {
  description = "A map of objects with EventBridge Pipe definitions."
  type        = any
  default     = {}
}

variable "append_pipe_postfix" {
  description = "Controls whether to append '-pipe' to the name of the pipe"
  type        = bool
  default     = true
}
