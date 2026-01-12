variable "application" {
  type        = string
  description = "The application name of the api gateway, will be appended with the company, lob, env and region to form a api gateway name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the api gateway, will be appended with the company, lob, env and region to form a api gateway name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the acm, will be appended with the company, lob, env and region to form a acm name."
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
  description = "The lob name of the api gateway, will be appended with the company, lob, env and region to form a api gateway name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "domain_name" {
  type        = string
  default     = null
  description = "The domain name to use for API gateway"
}

variable "domain_name_certificate_arn" {
  type        = string
  default     = ""
  description = "The ARN of an AWS-managed certificate that will be used by the endpoint for the domain name"
}


variable "cors_configuration" {
  type        = any
  default     = {}
  description = <<-EOF
Map of API gateway routes with integrations

  Example/available options:
  ```
    {
    allow_credentials = true
    allow_methods     = ["GET", "OPTIONS", "POST"] or ["*"]
    max_age           = 5
    allow_headers = ["content-type", "x-amz-date", "authorization", "x-api-key", "x-amz-security-token", "x-amz-user-agent"]
    allow_origins = ["*"]
  }
  ```
    EOF
}


variable "integrations" {
  type        = map(any)
  default     = {}
  description = <<-EOF
Map of API gateway routes with integrations

  Example/available options:
  ```
  {
      "ANY /" = {
          lambda_arn             = module.lambda.arn
          payload_format_version = "2.0"
          timeout_milliseconds   = 12000
        }
        "GET /some-route-with-authorizer" = {
          lambda_arn             = module.lambda.arn
          payload_format_version = "2.0"
          authorizer_key         = "cognito"
        }
        "POST /start-step-function" = {
          lambda_arn             = module.lambda.arn
          payload_format_version = "2.0"
          authorizer_key         = "cognito"
        }

      "
    }
  }
  ```
    EOF
}

variable "protocol_type" {
  type        = string
  default     = "HTTP"
  description = "The API protocol. Valid values: HTTP, WEBSOCKET"
}

variable "default_stage_access_log_destination_arn" {
  type        = string
  default     = null
  description = "Default stage's ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN."
}

variable "default_stage_access_log_format" {
  type        = string
  default     = null
  description = "Default stage's single line format of the access logs of data, as specified by selected $context variables."
}

variable "authorizers" {
  type        = map(any)
  default     = {}
  description = <<-EOF
Map of API gateway routes with integrations

  Example/available options:
  ```
  {
    "azure" = {
      authorizer_type  = "JWT"
      identity_sources = "$request.header.Authorization"
      name             = "azure-auth"
      audience         = ["d6a38afd-45d6-4874-d1aa-3c5c558aqcc2"]
      issuer           = "https://sts.windows.net/aaee026e-8f37-410e-8869-72d9154873e4/"
    }
  }
  ```
    EOF
}

variable "create_api_domain_name" {
  description = "Whether to create API domain name resource"
  type        = bool
  default     = true
}

variable "vpc_links" {
  type        = map(any)
  default     = {}
  description = <<-EOF
Map of API gateway routes with integrations

  Example/available options:
  ```
  {
    my-vpc = {
      name               = "example"
      security_group_ids = [security_group_id]
      subnet_ids         = public_subnets_id
    }
  }
  ```
    EOF
}

variable "create_routes_and_integrations" {
  description = "Whether to create routes and integrations resources"
  type        = bool
  default     = true
}
