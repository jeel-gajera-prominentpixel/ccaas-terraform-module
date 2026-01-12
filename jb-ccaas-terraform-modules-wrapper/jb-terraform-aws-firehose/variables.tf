variable "application" {
  type        = string
  description = "The application name of the bucket, will be appended with the company, lob, env and region to form a bucket name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the bucket, will be appended with the company, lob, env and region to form a bucket name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the bucket , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the bucket, will be appended with the company, lob, env and region to form a bucket name."
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

variable "input_source" {
  description = "This is the kinesis firehose source"
  type        = string
  default     = "direct-put"
  validation {
    error_message = "Please use a valid source!"
    condition     = contains(["direct-put", "kinesis", "waf", "msk"], var.input_source)
  }
}

variable "destination" {
  description = "This is the destination to where the data is delivered"
  type        = string
  validation {
    error_message = "Please use a valid destination!"
    condition     = contains(["s3", "extended_s3", "redshift", "opensearch", "opensearchserverless", "elasticsearch", "splunk", "http_endpoint", "datadog", "coralogix", "newrelic", "dynatrace", "honeycomb", "logicmonitor", "mongodb", "sumologic"], var.destination)
  }
}

variable "kinesis_source_stream_arn" {
  description = "The kinesis stream used as the source of the firehose delivery stream"
  type        = string
  default     = null
}

variable "s3_bucket_arn" {
  description = "The ARN of the S3 destination bucket"
  type        = string
  default     = null
}

variable "s3_error_output_prefix" {
  description = "Prefix added to failed records before writing them to S3. This prefix appears immediately following the bucket name."
  type        = string
  default     = null
}

variable "s3_backup_bucket_arn" {
  description = "The ARN of the S3 backup bucket"
  type        = string
  default     = null
}


variable "enable_sse" {
  description = "Whether to enable encryption at rest. Only makes sense when source is Direct Put"
  type        = bool
  default     = false
}

variable "sse_kms_key_arn" {
  description = "Amazon Resource Name (ARN) of the encryption key"
  type        = string
  default     = null
}

variable "enable_s3_backup" {
  description = "The Amazon S3 backup mode"
  type        = bool
  default     = false
}
variable "sse_kms_key_type" {
  type    = string
  default = "AWS_OWNED_CMK"

}

variable "enable_s3_encryption" {
  description = "Indicates if want use encryption in S3 bucket."
  type        = bool
  default     = false
}
variable "s3_kms_key_arn" {
  description = "Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will be used"
  type        = string
  default     = null
}
variable "s3_backup_prefix" {
  description = "The YYYY/MM/DD/HH time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket"
  type        = string
  default     = null
}

variable "s3_backup_error_output_prefix" {
  description = "Prefix added to failed records before writing them to S3"
  type        = string
  default     = null
}

variable "s3_backup_buffering_interval" {
  description = "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination."
  type        = number
  default     = 300
  validation {
    error_message = "Valid Values: Minimum: 60 seconds, maximum: 900 seconds."
    condition     = var.s3_backup_buffering_interval >= 60 && var.s3_backup_buffering_interval <= 900
  }
}

variable "s3_backup_buffering_size" {
  description = "Buffer incoming data to the specified size, in MBs, before delivering it to the destination."
  type        = number
  default     = 5
  validation {
    error_message = "Valid values: minimum: 1 MiB, maximum: 128 MiB."
    condition     = var.s3_backup_buffering_size >= 1 && var.s3_backup_buffering_size <= 128
  }
}


variable "s3_backup_compression" {
  description = "The compression format"
  type        = string
  default     = "UNCOMPRESSED"
  validation {
    error_message = "Valid values are UNCOMPRESSED, GZIP, ZIP, Snappy and HADOOP_SNAPPY."
    condition     = contains(["UNCOMPRESSED", "GZIP", "ZIP", "Snappy", "HADOOP_SNAPPY"], var.s3_backup_compression)
  }
}

variable "s3_backup_use_existing_role" {
  description = "Indicates if want use the kinesis firehose role to s3 backup bucket access."
  type        = bool
  default     = true
}

variable "s3_backup_role_arn" {
  description = "The role that Kinesis Data Firehose can use to access S3 Backup."
  type        = string
  default     = null
}

variable "s3_backup_enable_encryption" {
  description = "Indicates if want enable KMS Encryption in S3 Backup Bucket."
  type        = bool
  default     = false
}

variable "s3_backup_kms_key_arn" {
  description = "Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will be used."
  type        = string
  default     = null
}

variable "s3_backup_enable_log" {
  description = "Enables or disables the logging"
  type        = bool
  default     = true
}

variable "create_application_role" {
  description = "Set it to true to create role to be used by the source"
  default     = false
  type        = bool
}

variable "create_application_role_policy" {
  description = "Set it to true to create policy to the role used by the source"
  default     = false
  type        = bool
}

variable "application_role_service_principal" {
  description = "AWS Service Principal to assume application role"
  type        = string
  default     = null
}

variable "buffering_size" {
  description = "Buffer incoming data to the specified size, in MBs, before delivering it to the destination."
  type        = number
  default     = 5
  validation {
    error_message = "Valid values: minimum: 1 MiB, maximum: 128 MiB."
    condition     = var.buffering_size >= 1 && var.buffering_size <= 128
  }
}

variable "buffering_interval" {
  description = "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination"
  type        = number
  default     = 300
  validation {
    error_message = "Valid Values: Minimum: 0 seconds, maximum: 900 seconds."
    condition     = var.buffering_interval >= 0 && var.buffering_interval <= 900
  }
}

variable "s3_prefix" {
  description = "Specifies the  s3 bucket prefix"
  type        = string
  default     = null
}

variable "kinesis_source_kms_arn" {
  description = "Kinesis Source KMS Key to add Firehose role to decrypt the records."
  type        = string
  default     = null
}

variable "kinesis_source_is_encrypted" {
  description = "Indicates if Kinesis data stream source is encrypted"
  type        = bool
  default     = false
}

variable "append_delimiter_to_record" {
  description = "To configure your delivery stream to add a new line delimiter between records in objects that are delivered to Amazon S3."
  type        = bool
  default     = false
}

variable "redshift_copy_options" {
  description = "Copy options for copying the data from the s3 intermediate bucket into redshift, for example to change the default delimiter"
  type        = string
  default     = null
}

variable "dynamic_partition_append_delimiter_to_record" { # TODO Variable Deprecated. Remove on Next Major Version
  description = "DEPRECATED!! Use var append_delimiter_to_record instead!! Use To configure your delivery stream to add a new line delimiter between records in objects that are delivered to Amazon S3."
  type        = bool
  default     = false
}

variable "dynamic_partition_enable_record_deaggregation" {
  description = "Data deaggregation is the process of parsing through the records in a delivery stream and separating the records based either on valid JSON or on the specified delimiter"
  type        = bool
  default     = false
}

variable "dynamic_partition_record_deaggregation_type" {
  description = "Data deaggregation is the process of parsing through the records in a delivery stream and separating the records based either on valid JSON or on the specified delimiter"
  type        = string
  default     = "JSON"
  validation {
    error_message = "Valid values are JSON and DELIMITED."
    condition     = contains(["JSON", "DELIMITED"], var.dynamic_partition_record_deaggregation_type)
  }
}

variable "dynamic_partition_record_deaggregation_delimiter" {
  description = "Specifies the delimiter to be used for parsing through the records in the delivery stream and deaggregating them"
  type        = string
  default     = null
}
