variable "application" {
  type        = string
  description = "The application name of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name."
}

variable "create_rds_proxy" {
  description = "Whether cluster should be created (affects nearly all resources)"
  type        = bool
  default     = false
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name"
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the rds proxy , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "lob" {
  type        = string
  description = "The lob name of the rds proxy, will be appended with the company, lob, env and region to form a rds proxy name"
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching rdss will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}

variable "vpc_subnet_ids" {
  description = "One or more VPC subnet IDs to associate with the new proxy"
  type        = list(string)
  default     = []
}

variable "vpc_security_group_ids" {
  description = "One or more VPC security group IDs to associate with the new proxy"
  type        = list(string)
  default     = []
}

# Proxy endpoints
variable "endpoints" {
  type        = any
  default     = {}
  description = <<-EOF
    Map of DB proxy endpoints to create and their attributes (see `aws_db_proxy_endpoint`)
    Example/available options:
    ```
    {
      read_write = {
        name                   = "read-write-endpoint"
        vpc_subnet_ids         = module.vpc.private_subnets
        vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]
        tags                   = local.tags
      },
      read_only = {
        name                   = "read-only-endpoint"
        vpc_subnet_ids         = module.vpc.private_subnets
        vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]
        target_role            = "READ_ONLY"
        tags                   = local.tags
      }
    }
    ```
      EOF
}

variable "auth" {
  type        = any
  default     = {}
  description = <<-EOF
    Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters
    Example/available options:
    ```
    {
      "root" = {
        description = "Cluster generated master user password"
        secret_arn  = module.rds.cluster_master_user_secret[0].secret_arn
      }
    }
    ```
      EOF
}

variable "engine_family" {
  description = "The kind of database engine that the proxy will connect to. Valid values are `MYSQL` or `POSTGRESQL`"
  type        = string
  default     = ""
}

variable "debug_logging" {
  description = "Whether the proxy includes detailed information about SQL statements in its logs"
  type        = bool
  default     = false
}

variable "target_db_cluster" {
  description = "Determines whether DB cluster is targeted by proxy"
  type        = bool
  default     = false
}

variable "db_cluster_identifier" {
  description = "DB cluster identifier"
  type        = string
  default     = ""
}

variable "target_db_instance" {
  description = "Determines whether DB instance is targeted by proxy"
  type        = bool
  default     = false
}

variable "db_instance_identifier" {
  description = "DB instance identifier"
  type        = string
  default     = ""
}

variable "idle_client_timeout" {
  description = "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it"
  type        = number
  default     = 1800
}

variable "require_tls" {
  description = "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy"
  type        = bool
  default     = true
}

variable "role_arn" {
  description = "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager"
  type        = string
  default     = ""
}

################################################################################
# IAM Role
################################################################################

variable "create_iam_role" {
  description = "Determines whether an IAM role is created"
  type        = bool
  default     = true
}

variable "iam_role_name" {
  description = "The name of the role. If omitted, Terraform will assign a random, unique name"
  type        = string
  default     = ""
}

variable "use_role_name_prefix" {
  description = "Whether to use unique name beginning with the specified `iam_role_name`"
  type        = bool
  default     = false
}

variable "iam_role_description" {
  description = "The description of the role"
  type        = string
  default     = ""
}

variable "iam_role_path" {
  description = "The path to the role"
  type        = string
  default     = null
}

variable "iam_role_force_detach_policies" {
  description = "Specifies to force detaching any policies the role has before destroying it"
  type        = bool
  default     = true
}

variable "iam_role_max_session_duration" {
  description = "The maximum session duration (in seconds) that you want to set for the specified role"
  type        = number
  default     = 43200 # 12 hours
}

variable "iam_role_permissions_boundary" {
  description = "The ARN of the policy that is used to set the permissions boundary for the role"
  type        = string
  default     = null
}

variable "iam_role_tags" {
  description = "A map of tags to apply to the IAM role"
  type        = map(string)
  default     = {}
}

# IAM Policy
variable "create_iam_policy" {
  description = "Determines whether an IAM policy is created"
  type        = bool
  default     = true
}

variable "iam_policy_name" {
  description = "The name of the role policy. If omitted, Terraform will assign a random, unique name"
  type        = string
  default     = ""
}

variable "use_policy_name_prefix" {
  description = "Whether to use unique name beginning with the specified `iam_policy_name`"
  type        = bool
  default     = false
}

variable "kms_key_arns" {
  description = "List of KMS Key ARNs to allow access to decrypt SecretsManager secrets"
  type        = list(string)
  default     = []
}

################################################################################
# CloudWatch Logs
################################################################################

variable "manage_log_group" {
  description = "Determines whether Terraform will create/manage the CloudWatch log group or not. Note - this will fail if set to true after the log group has been created as the resource will already exist"
  type        = bool
  default     = true
}

variable "log_group_retention_in_days" {
  description = "Specifies the number of days you want to retain log events in the log group"
  type        = number
  default     = 30
}

variable "log_group_kms_key_id" {
  description = "The ARN of the KMS Key to use when encrypting log data"
  type        = string
  default     = null
}

variable "log_group_tags" {
  description = "A map of tags to apply to the CloudWatch log group"
  type        = map(string)
  default     = {}
}
