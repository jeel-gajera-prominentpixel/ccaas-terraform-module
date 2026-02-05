variable "name" {
  description = "Name of the API Gateway REST API"
  type        = string
  default     = ""
}

variable "tags" {
  description = "A mapping of tags to assign to API gateway resources."
  type        = map(string)
  default     = {}
}

variable "stage_name" {
  description = "Stage Name of the API Gateway"
  type        = string
}

variable "create_root_method" {
  description = "true for the root methode creation for lambda integration."
  type        = bool
  default     = false
}

variable "create_method" {
  description = "true for the methode creation for lambda integration."
  type        = bool
  default     = false
}

variable "root_integration_http_method" {
  description = "root_integration_http_method for integration of lambda"
  type        = string
  default     = "arn:aws:lambda:region:account-id:function:root-function"
}

variable "root_integration_type" {
  description = "root_integration_type for integration of lambda"
  type        = string
  default     = "arn:aws:lambda:region:account-id:function:root-function"
}

variable "root_lambda_arn" {
  description = "Lambda function ARN for the root method"
  type        = string
  default     = "arn:aws:lambda:region:account-id:function:root-function"
}

variable "description" {
  description = "Description of the API Gateway REST API"
  type        = string
}

variable "types" {
  description = "Type of the API Gateway Endpoint Available values are EDGE, REGIONAL and PRIVATE. Default is REGIONAL"
  type        = string
  default     = "REGIONAL"
}

variable "authorization" {
  description = "Authorization type for the API Gateway methods"
  type        = string
  default     = "NONE"
}

variable "resource_root_path" {
  description = "List of paths for the API Gateway resources"
  type        = string
  default     = "ANY"
}

variable "enable_waf_association" {
  description = "Flag to enable or disable WAF association with resources"
  type        = bool
  default     = false
}

variable "web_acl_arn" {
  description = "WAF ARNs to associate with the resource"
  type        = string
  default     = ""
}

variable "resource_paths" {
  description = "Map of paths for the API Gateway resources with Lambda ARNs and HTTP methods"
  type        = any
  default     = {}
}

variable "root_resource_request_parameters" {
  description = "Root resource request params."
  type        = any
  default     = {}
}

variable "root_integration_request_parameters" {
  description = "Root integration request params."
  type        = any
  default     = {}
}

variable "enable_logs" {
  description = "Enable API Gateway Logs."
  type        = bool
  default     = false
}

variable "create_api" {
  description = "Flag to control whether API Gateway should be created"
  type        = bool
  default     = true
}

variable "rest_api_id" {
  description = "Existing REST API ID (if create_api = false)"
  type        = string
  default     = null
}
variable "rest_api_execution_arn" {
  type        = string
  default     = null
  description = "The execution ARN of the REST API. Used for setting up Lambda permissions."
}

variable "create_child_resource" {
  type        = bool
  default     = true
  description = "Whether to create a child resource under the specified parent resource."
}

variable "parent_id" {
  type        = string
  default     = null
  description = "The ID of the parent API Gateway resource. Required if creating a child resource."
}

variable "path_part" {
  type        = string
  default     = null
  description = "The path part (e.g., 'admin', 'user') for the API Gateway resource."
}

variable "create_authorizer" {
  description = "Flag to control whether a custom API Gateway authorizer should be created"
  type        = bool
  default     = false
}

variable "authorizer_name" {
  description = "Name of the custom API Gateway authorizer"
  type        = string
  default     = "auth_middleware"
}

variable "authorizer_lambda_arn" {
  type        = string
  description = "ARN of the authorizer Lambda"
  default     = ""
  validation {
    condition     = var.create_authorizer ? length(var.authorizer_lambda_arn) > 0 : true
    error_message = "You must set authorizer_lambda_arn if create_authorizer is true."
  }
}

variable "identity_source" {
  description = "The identity source for which authorization is requested, e.g., method.request.header.token"
  type        = string
  default     = "method.request.header.token"
}

variable "authorizer_result_ttl_in_seconds" {
  description = "TTL in seconds for which API Gateway caches authorizer results"
  type        = number
  default     = 0
}

#variable "authorizer_id" {
#  description = "Optional: ID of an existing API Gateway authorizer to use instead of creating a new one"
#  type        = string
#  default     = null
#}

variable "enable_binary_media_types" {
  type        = bool
  default     = false
  description = "Enable binary media types for this API Gateway"
}

variable "binary_media_types" {
  type        = list(string)
  default     = []
  description = "List of binary media types to configure (if enabled)"
}

variable "root_method_response_params" {
  description = "Method response configuration for root method"
  type = map(object({
    status_code         = string
    response_parameters = map(bool)
    response_models     = optional(map(string), {})
  }))
  default = {}
}

variable "root_integration_response_params" {
  description = "Integration response configuration for root method"
  type = map(object({
    status_code         = string
    response_parameters = optional(map(string), {})
  }))
  default = {}
}

variable "method_response_params" {
  description = "Method response configuration for each resource path"
  type = map(object({
    status_code         = string
    response_parameters = map(bool)
    response_models     = optional(map(string), {})
  }))
  default = {}
}