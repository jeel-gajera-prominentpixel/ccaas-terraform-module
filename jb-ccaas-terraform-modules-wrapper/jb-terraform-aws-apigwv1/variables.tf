variable "application" {
  type        = string
  description = "The application name of the API v1, will be appended with the company, lob, env and region to form a API v1 name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the API v1, will be appended with the company, lob, env and region to form a API v1 name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws API v1 , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the API v1, will be appended with the company, lob, env and region to form a API v1 name."
}

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
  default     = "dev"
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
  type = map(object({
    lambda_arn              = string
    http_method             = string
    integration_http_method = string
    type                    = string
  }))
  default = {
    "/proxy" = {
      lambda_arn              = "arn:aws:lambda:us-west-2:767252029631:function:Test-lambda-deployment"
      http_method             = "POST"
      integration_http_method = "POST"
      type                    = "AWS_PROXY"
    }
    "/proxy2" = {
      lambda_arn              = "arn:aws:lambda:us-west-2:767252029631:function:Test-lambda-deployment"
      http_method             = "POST"
      integration_http_method = "POST"
      type                    = "AWS_PROXY"
    }
  }
}
