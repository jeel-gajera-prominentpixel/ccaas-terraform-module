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

variable "enabled" {
  type        = bool
  default     = true
  description = "Set this to `false` to prevent resource creation by this terraform module."
}

variable "create_rest_api" {
  type        = bool
  default     = false
  description = "Flag to control the rest api creation."
}

variable "rest_api_description" {
  type        = string
  default     = "test"
  description = "The description of the REST API"
}

variable "rest_api_endpoint_type" {
  type        = string
  default     = null
  description = "(Required) List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE, REGIONAL or PRIVATE. If unspecified, defaults to EDGE."
}

variable "create_vpc_endpoint" {
  type        = bool
  default     = true
  description = "VPC endpoint is required to access api gateway url from outside the vpc. Set this to `false` to prevent vpc endpoint creation."
}

variable "rest_api_resource_policy" {
  type        = string
  default     = ""
  description = "(Optional) custom resource policy for private rest api."
}

variable "create_rest_api_deployment" {
  type        = bool
  default     = true
  description = "Flag to control the mapping creation."
}

variable "api_deployment_description" {
  type        = string
  default     = "test"
  description = "flag to manage description of api deployment"
}

variable "stage_description" {
  type        = string
  default     = "test"
  description = "Description to set on the stage managed by the stage_name argument."
}

variable "rest_variables" {
  type        = map(string)
  default     = {}
  description = "Map to set on the stage managed by the stage_name argument."
}

variable "create_rest_api_gateway_resource" {
  type        = bool
  default     = true
  description = "flag to control the rest api gateway resources creation"
}

variable "api_resources" {
  type        = map(map(string))
  default     = {}
  description = "flag to control of resources path"
}

variable "create_rest_api_gateway_method" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway method creation."
}

variable "authorization" {
  type        = string
  default     = "NONE"
  description = "  Required The type of authorization used for the method (NONE, CUSTOM, AWS_IAM, COGNITO_USER_POOLS)"
}

variable "create_rest_api_gateway_integration" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway integration creation."
}

variable "integration_http_method" {
  type        = string
  default     = "POST"
  description = "flag to control the gateway intergration http method."
}


variable "gateway_integration_type" {
  type        = string
  default     = "AWS_PROXY"
  description = "flag tp control the gatway integration type. "
}

variable "connection_rest_api_type" {
  type        = string
  default     = "INTERNET"
  description = "Valid values are INTERNET (default for connections through the public routable internet), and VPC_LINK (for private connections between API Gateway and a network load balancer in a VPC)."
}


variable "connection_id" {
  type        = string
  default     = ""
  description = "ID of the VpcLink used for the integration. Required if connection_type is VPC_LINK"
}

variable "credentials" {
  type        = string
  default     = ""
  description = "To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string "
}

variable "request_templates" {
  type        = map(string)
  default     = null
  description = "Map of the integration's request templates."
}

variable "request_parameters" {
  type        = map(string)
  default     = null
  description = "Map of request query string parameters and headers that should be passed to the backend responder"
}

variable "cache_key_parameters" {
  type        = list(any)
  default     = []
  description = "List of cache key parameters for the integration."
}

variable "cache_namespace" {
  type        = string
  default     = ""
  description = "Integration's cache namespace."
}

variable "content_handling" {
  type        = string
  default     = "CONVERT_TO_TEXT"
  description = "Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through."
}

variable "http_method" {
  type        = string
  default     = "ANY"
  description = "HTTP method (GET, POST, PUT, DELETE, HEAD, OPTION, ANY) when calling the associated resource."
}

variable "timeout_milliseconds" {
  type        = number
  default     = null
  description = "Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds."
}

variable "integration_uri" {
  type        = string
  default     = ""
  description = "URI of the Lambda function for a Lambda proxy integration, when integration_type is AWS_PROXY. For an HTTP integration, specify a fully-qualified URL."
}

variable "create_rest_api_gateway_stage" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway stage creation."
}

variable "description_gateway_stage" {
  type        = string
  default     = "demo-test"
  description = "(optional) describe your variable"
}

variable "rest_api_stage_name" {
  type        = string
  default     = ""
  description = "The name of the stage"
}

variable "cache_cluster_enabled" {
  type        = bool
  default     = false
  description = "Whether a cache cluster is enabled for the stage"
}

variable "cache_cluster_size" {
  type        = string
  default     = "0.5"
  description = "Size of the cache cluster for the stage, if enabled. Allowed values include 0.5, 1.6, 6.1, 13.5, 28.4, 58.2, 118 and 237."
}

variable "client_certificate_id" {
  type        = string
  default     = ""
  description = "Identifier of a client certificate for the stage."
}

variable "documentation_version" {
  type        = string
  default     = ""
  description = "Version of the associated API documentation"
}

variable "stage_variables" {
  type        = map(string)
  default     = {}
  description = "Map that defines the stage variables"
}

variable "xray_tracing_enabled" {
  type        = bool
  default     = true
  description = "A flag to indicate whether to enable X-Ray tracing."
}

variable "canary_settings" {
  type        = map(any)
  default     = {}
  description = "(optional) describe your variable"
}

variable "enable_access_logs" {
  type        = bool
  default     = true
  description = "flag to manage of cloudwatch log group creation"
}

variable "log_format" {
  description = " Formatting and values recorded in the logs. For more information on configuring the log format rules visit the AWS documentation"
  type        = string
  default     = <<EOF
  {
	"requestTime": "$context.requestTime",
	"requestId": "$context.requestId",
	"httpMethod": "$context.httpMethod",
	"path": "$context.path",
	"resourcePath": "$context.resourcePath",
	"status": $context.status,
	"responseLatency": $context.responseLatency,
  "xrayTraceId": "$context.xrayTraceId",
  "integrationRequestId": "$context.integration.requestId",
	"functionResponseStatus": "$context.integration.status",
  "integrationLatency": "$context.integration.latency",
	"integrationServiceStatus": "$context.integration.integrationStatus",
  "authorizeResultStatus": "$context.authorize.status",
	"authorizerServiceStatus": "$context.authorizer.status",
	"authorizerLatency": "$context.authorizer.latency",
	"authorizerRequestId": "$context.authorizer.requestId",
  "ip": "$context.identity.sourceIp",
	"userAgent": "$context.identity.userAgent",
	"principalId": "$context.authorizer.principalId",
	"cognitoUser": "$context.identity.cognitoIdentityId",
  "user": "$context.identity.user"
}
  EOF
}

variable "create_rest_api_gateway_method_response" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway stage creation."
}

variable "create_rest_api_gateway_integration_response" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway integration response creation."
}

variable "status_code" {
  type        = string
  default     = "200"
  description = "flag to control the status code"
}

variable "response_models" {
  type = map(string)
  default = {
    "application/json" = "Empty"
  }
  description = "A map of the API models used for the response's content type"
}

variable "response_parameters" {
  type        = map(bool)
  default     = {}
  description = "Map of response parameters that can be sent to the caller. For example: response_parameters { method.response.header.X-Some-Header = true } would define that the header X-Some-Header can be provided on the response"
}

variable "integration_response_parameters" {
  type        = map(string)
  default     = {}
  description = " Map of response parameters that can be read from the backend response. For example: response_parameters = { method.response.header.X-Some-Header = integration.response.header.X-Some-Other-Header }."
}

variable "create_rest_api_gateway_authorizer" {
  type        = bool
  default     = true
  description = "Flag to control the rest api gateway authorizer creation."
}

variable "gateway_authorizer" {
  type        = string
  default     = "demo"
  description = "flag to control the gateway authorizer name."
}

variable "authorizer_iam_role" {
  type        = string
  default     = ""
  description = " Custome IAMRole for Authorizer Credentials."
}

variable "identity_source" {
  type        = string
  default     = "method.request.header.Authorization"
  description = "Source of the identity in an incoming request. Defaults to method.request.header.Authorization. For REQUEST type, this may be a comma-separated list of values, including headers, query string parameters and stage variable"
}

variable "type" {
  type        = string
  default     = "TOKEN"
  description = "Type of the authorizer. Possible values are TOKEN for a Lambda function using a single authorization token submitted in a custom header, REQUEST for a Lambda function using incoming request parameters, or COGNITO_USER_POOLS for using an Amazon Cognito user pool. Defaults to TOKEN."
}

variable "authorizer_result_ttl_in_seconds" {
  type        = number
  default     = 300
  description = "TTL of cached authorizer results in seconds. Defaults to 300."
}

variable "provider_arns" {
  type        = set(string)
  default     = []
  description = "required for type COGNITO_USER_POOLS) List of the Amazon Cognito user pool ARNs. Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}."
}

variable "domain_name" {
  type        = string
  default     = null
  description = "The domain name to use for API gateway"
}

variable "rest_api_base_path" {
  type        = string
  default     = ""
  description = "Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain."
}

variable "vpc_endpoint_ids" {
  type        = list(string)
  default     = []
  description = "VPC endpoint IDs for PRIVATE API Gateway"
}
