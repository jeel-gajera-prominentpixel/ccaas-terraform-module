variable "api_name" {
  type = string
}

variable "root_methods" {
  description = "Methods on root resource"
  type = map(object({
    authorization = optional(string, "NONE")
    lambda_arn    = string

    method_responses = optional(map(object({
      response_models = optional(map(string), {})
      response_parameters = optional(map(bool), {})
    })))
  }))
}

variable "proxy_methods" {
  description = "Methods on {proxy+} resource"
  type = map(object({
    authorization = optional(string, "NONE")
    lambda_arn    = string

    method_responses = optional(map(object({
      response_models = optional(map(string), {})
      response_parameters = optional(map(bool), {})
    })))
  }))
}
