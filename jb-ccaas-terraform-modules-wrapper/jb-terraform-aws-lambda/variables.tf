variable "application" {
  type        = string
  description = "The application name of the lambda, will be appended with the company, lob, env and region to form a lambda name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the lambda, will be appended with the company, lob, env and region to form a lambda name"
}


variable "prefix_region" {
  type        = string
  description = "The prefix region of the lambda , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "lob" {
  type        = string
  description = "The lob name of the lambda, will be appended with the company, lob, env and region to form a lambda name"
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

# specific variables for the lambda function

variable "description" {
  description = "Description of your Lambda Function (or Layer)"
  type        = string
  default     = ""
}

variable "handler" {
  description = "Lambda Function entrypoint in your code"
  type        = string
  default     = ""
}

variable "runtime" {
  description = "Lambda Function runtime"
  type        = string
  default     = ""
}

variable "local_existing_package" {
  description = "The absolute path to an existing zip-file to use"
  type        = string
  default     = null
}

variable "environment_variables" {
  description = "A map that defines environment variables for the Lambda Function."
  type        = map(string)
  default     = {}
}

variable "timeout" {
  description = "The amount of time your Lambda Function has to run in seconds."
  type        = number
  default     = 3
}

variable "lambda_role" {
  description = " IAM role ARN attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See Lambda Permission Model for more details."
  type        = string
  default     = ""
}

variable "publish" {
  description = "Whether to publish creation/change as new Lambda Function Version."
  type        = bool
  default     = false
}

variable "create_lambda_function_url" {
  description = "Controls whether the Lambda Function URL resource should be created"
  type        = bool
  default     = false
}

variable "store_on_s3" {
  description = "Whether to store produced artifacts on S3 or locally."
  type        = bool
  default     = false
}

variable "s3_bucket" {
  description = "S3 bucket to store artifacts"
  type        = string
  default     = null
}

variable "s3_prefix" {
  description = "Directory name where artifacts should be stored in the S3 bucket. If unset, the path from `artifacts_dir` is used"
  type        = string
  default     = null
}

variable "create_role" {
  description = "Controls whether IAM role for Lambda Function should be created"
  type        = bool
  default     = true
}

variable "ignore_source_code_hash" {
  description = "Whether to ignore changes to the function's source code hash. Set to true if you manage infrastructure and code deployments separately."
  type        = bool
  default     = true
}

variable "attach_policy_statements" {
  description = "Controls whether policy_statements should be added to IAM role for Lambda Function"
  type        = bool
  default     = false
}

variable "policy_statements" {
  description = "Map of dynamic policy statements to attach to Lambda Function role"
  type        = any
  default     = {}
}

variable "vpc_subnet_ids" {
  description = "List of subnet ids when Lambda Function should run in the VPC. Usually private or intra subnets"
  type        = list(string)
  default     = null
}

variable "vpc_security_group_ids" {
  description = "List of security group ids when Lambda Function should run in the VPC. "
  type        = list(string)
  default     = null
}

variable "attach_network_policy" {
  description = "Controls whether VPC/network policy should be added to IAM role for Lambda Function"
  type        = bool
  default     = false
}

variable "attach_policy_jsons" {
  description = "Controls whether policy_jsons should be added to IAM role for Lambda Function"
  default     = false
  type        = bool
}

variable "policy_jsons" {
  description = "List of additional policy documents as JSON to attach to Lambda Function role"
  type        = list(string)
  default     = []
}

variable "number_of_policy_jsons" {
  description = "Number of policies JSON to attach to IAM role for Lambda Function"
  type        = number
  default     = 0
}

variable "replace_security_groups_on_destroy" {
  description = "(Optional) When true, all security groups defined in vpc_security_group_ids will be replaced with the default security group after the function is destroyed. Set the replacement_security_group_ids variable to use a custom list of security groups for replacement instead."
  type        = bool
  default     = null
}

variable "replacement_security_group_ids" {
  description = "(Optional) List of security group IDs to assign to orphaned Lambda function network interfaces upon destruction. replace_security_groups_on_destroy must be set to true to use this attribute."
  type        = list(string)
  default     = null
}

variable "memory_size" {
  description = "Amount of memory in MB your Lambda Function can use at runtime. Valid value between 128 MB to 10,240 MB (10 GB), in 64 MB increments."
  type        = number
  default     = 128
}

variable "create_layer" {
  description = "Controls whether Lambda Layer resource should be created"
  type        = bool
  default     = false
}

variable "layer_name" {
  description = "Name of Lambda Layer to create"
  type        = string
  default     = ""
}

variable "layers" {
  description = "List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function."
  type        = list(string)
  default     = null
}

variable "ephemeral_storage_size" {
  description = "Amount of ephemeral storage (/tmp) in MB your Lambda Function can use at runtime. Valid value between 512 MB to 10,240 MB (10 GB)."
  type        = number
  default     = 512
}

variable "compatible_runtimes" {
  description = "A list of Runtimes this layer is compatible with. Up to 5 runtimes can be specified."
  type        = list(string)
  default     = []
}

variable "compatible_architectures" {
  description = "A list of Architectures Lambda layer is compatible with. Currently x86_64 and arm64 can be specified."
  type        = list(string)
  default     = null
}

variable "architectures" {
  description = "Instruction set architecture for your Lambda function. Valid values are [\"x86_64\"] and [\"arm64\"]."
  type        = list(string)
  default     = null
}

variable "allowed_triggers" {
  description = "Map of allowed triggers to create Lambda permissions"
  type        = map(any)
  default     = {}
}

variable "authorization_type" {
  description = "The type of authentication that the Lambda Function URL uses. Set to 'AWS_IAM' to restrict access to authenticated IAM users only. Set to 'NONE' to bypass IAM authentication and create a public endpoint."
  type        = string
  default     = "AWS_IAM"
}

variable "cors" {
  description = "CORS settings to be used by the Lambda Function URL"
  type        = any
  default     = {}
}

variable "invoke_mode" {
  description = "Invoke mode of the Lambda Function URL. Valid values are BUFFERED (default) and RESPONSE_STREAM."
  type        = string
  default     = null
}

variable "trusted_entities" {
  description = "List of additional trusted entities for assuming Lambda Function role (trust relationship)"
  type        = any
  default     = []
}

variable "policy_name" {
  description = "IAM policy name. It override the default value, which is the same as role_name"
  type        = string
  default     = null
}

variable "role_name" {
  description = "Name of IAM role to use for Lambda Function"
  type        = string
  default     = null
}

variable "tracing_mode" {
  description = "Tracing mode of the Lambda Function. Valid value can be either PassThrough or Active."
  type        = string
  default     = null
}

variable "create_current_version_allowed_triggers" {
  description = "Whether to allow triggers on current version of Lambda Function (this will revoke permissions from previous version because Terraform manages only current resources)"
  type        = bool
  default     = true
}

variable "create_unqualified_alias_allowed_triggers" {
  description = "Whether to allow triggers on unqualified alias pointing to $LATEST version"
  type        = bool
  default     = true
}

variable "event_source_mapping" {
  description = "Event source mapping configuration for the Lambda function"
  type        = any
  default     = {}
}
