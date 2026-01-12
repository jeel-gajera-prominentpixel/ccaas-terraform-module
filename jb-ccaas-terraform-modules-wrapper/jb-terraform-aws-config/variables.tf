variable "application" {
  type        = string
  description = "The application name of the aws config, will be appended with the company, lob, env and region to form a aws config name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the aws config, will be appended with the company, lob, env and region to form a aws config name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws config , will be appended with the company, lob, env and region to form a acm name."
}

variable "lob" {
  type        = string
  description = "The lob name of the aws config, will be appended with the company, lob, env and region to form a aws config name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "create_sns_topic" {
  description = <<-DOC
    Flag to indicate whether an SNS topic should be created for notifications
    If you want to send findings to a new SNS topic, set this to true and provide a valid configuration for subscribers
  DOC

  type    = bool
  default = false
}

variable "create_iam_role" {
  description = "Flag to indicate whether an IAM Role should be created to grant the proper permissions for AWS Config"
  type        = bool
  default     = false
}

variable "managed_rules" {
  description = <<-DOC
    A list of AWS Managed Rules that should be enabled on the account.

    See the following for a list of possible rules to enable:
    https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
  DOC
  type = map(object({
    description      = string
    identifier       = string
    input_parameters = any
    tags             = map(string)
    enabled          = bool
  }))
  default = {}
}

variable "force_destroy" {
  type        = bool
  description = "A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable"
  default     = false
}

variable "s3_bucket_arn" {
  description = "The ARN of the S3 bucket used to store the configuration history"
  type        = string
}

variable "s3_bucket_id" {
  description = "The id (name) of the S3 bucket used to store the configuration history"
  type        = string
}
