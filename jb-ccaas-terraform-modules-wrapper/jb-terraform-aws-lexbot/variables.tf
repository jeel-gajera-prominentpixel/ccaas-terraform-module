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
  description = "Name of the Lex bot"
  type        = string
  default     = ""
}

variable "description" {
  description = "Description of the Lex bot"
  type        = string
  default     = null
}

variable "auto_build_bot_locales" {
  description = "Specifies whether to automatically build the bot locales"
  type        = bool
  default     = true
}

variable "role_arn" {
  description = "ARN of the IAM role used by the bot"
  type        = string
  default     = null
}

variable "data_privacy" {
  description = "Provides information about data privacy for the bot"
  type = object({
    child_directed = bool
  })
  default = {
    child_directed = false
  }
}

variable "idle_session_ttl_in_seconds" {
  description = "IdleSessionTTLInSeconds of the resource"
  type        = number
  default     = 300
}

variable "bot_file_s3_location" {
  description = "S3 location of bot definitions zip file, if it's not defined inline in CloudFormation"
  type = object({
    s3_bucket         = optional(string)
    s3_object_key     = optional(string)
    s3_object_version = optional(string)
  })
  default = null
}

variable "bot_tags" {
  description = "A list of tags to add to the bot, which can only be added at bot creation"
  type = list(object({
    key   = string
    value = string
  }))
  default = []
}

# Bot version variables
variable "version_description" {
  description = "Description of the bot version"
  type        = string
  default     = null
}

variable "locale_specification" {
  description = "Locale specification for the bot version. Default is DRAFT version for en_US"
  type = map(object({
    source_bot_version = string
  }))
  default = {
    "en_US" = {
      source_bot_version = "DRAFT"
    }
  }
}

# Bot alias variables
variable "bot_alias_name" {
  description = "Name of the bot alias"
  type        = string
  default     = "LIVE" # Common practice to use LIVE for production alias
}

variable "sentiment_analysis_settings" {
  description = "Settings for sentiment analysis"
  type = object({
    detect_sentiment = bool
  })
  default = {
    detect_sentiment = false
  }
}

variable "bot_alias_locale_settings" {
  description = "A list of bot alias locale settings to add to the bot alias. You can use this to specify different Lambda functions for different locales."
  type = list(object({
    locale_id = string
    bot_alias_locale_setting = optional(object({
      enabled = optional(bool, true)
      code_hook_specification = optional(object({
        lambda_code_hook = object({
          code_hook_interface_version = string
          lambda_arn                  = string
        })
      }))
    }))
  }))
  default = []
}

variable "conversation_log_settings" {
  description = "Settings for conversation logging including both audio and text logs"
  type = object({
    # Audio Log Settings
    audio_log_settings = optional(list(object({
      enabled = optional(bool, false)
      destination = optional(object({
        s3_bucket = optional(object({
          kms_key_arn   = optional(string)
          log_prefix    = optional(string)
          s3_bucket_arn = optional(string)
        }))
      }))
    })), [])

    # Text Log Settings
    text_log_settings = optional(list(object({
      enabled = optional(bool, true)
      destination = optional(object({
        cloudwatch = optional(object({
          cloudwatch_log_group_arn = optional(string)
          log_prefix               = optional(string)
        }))
      }))
    })), [])
  })

  default = {
    audio_log_settings = []
    text_log_settings  = []
  }
}

variable "lex_policy" {
  description = "IAM policy for the Lex bot"
  type        = string
  default     = null
}

variable "bot_alias_arn" {
  description = "Lex bot alias ARN"
  type        = string
  default     = null
}

# Additional configuration variables that might be useful
variable "test_bot_alias_settings" {
  description = "Settings for the test bot alias if needed"
  type = object({
    create_test_alias = bool
    description       = string
  })
  default = {
    create_test_alias = false
    description       = "Test alias for bot"
  }
}

variable "enable_monitoring" {
  description = "Enable CloudWatch monitoring for the bot"
  type        = bool
  default     = true
}


variable "create_lexbot" {
  description = "Create Lex bot"
  type        = bool
  default     = false
}

variable "create_lexbot_version" {
  description = "Create Lex bot version"
  type        = bool
  default     = false
}

variable "create_lexbot_alias" {
  description = "Create Lex bot alias"
  type        = bool
  default     = false
}

variable "create_alias_policy" {
  description = "Create Lex bot alias policy"
  type        = bool
  default     = false
}

variable "bot_id" {
  description = "The identifier of the bot to create the version for"
  type        = string
  default     = null
}

variable "bot_version" {
  description = "The version of the bot to use for the alias"
  type        = string
  default     = null
}
