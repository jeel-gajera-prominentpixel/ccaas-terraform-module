variable "tags" {
  description = "A mapping of tags to assign to all resources."
  type        = map(string)
  default     = {}
}

variable "attributes" {
  description = "List of nested attribute definitions."
  type        = list(map(string))
  default     = []
}

variable "autoscaling_enabled" {
  description = "Whether or not to enable autoscaling."
  type        = bool
  default     = false
}

variable "autoscaling_indexes" {
  description = "A map of index autoscaling configurations."
  type        = map(map(string))
  default     = {}
}

variable "autoscaling_read" {
  description = "A map of read autoscaling settings. `max_capacity` is the only required key."
  type        = map(string)
  default     = {}
}

variable "autoscaling_write" {
  description = "A map of write autoscaling settings. `max_capacity` is the only required key."
  type        = map(string)
  default     = {}
}

variable "billing_mode" {
  description = "Controls how you are billed for read/write throughput and how you manage capacity. The valid values are `PROVISIONED` or `PAY_PER_REQUEST`."
  type        = string
  default     = "PAY_PER_REQUEST"
}

variable "create_table" {
  description = "Controls if DynamoDB table and associated resources are created."
  type        = bool
  default     = true
}

variable "deletion_protection_enabled" {
  description = "Enables deletion protection for table."
  type        = bool
  default     = null
}

variable "global_secondary_indexes" {
  description = "Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc."
  type        = any
  default     = []
}

variable "hash_key" {
  description = "The attribute to use as the hash (partition) key. Must also be defined as an attribute."
  type        = string
  default     = null
}

variable "ignore_changes_global_secondary_index" {
  description = "Whether to ignore changes lifecycle to global secondary indices, useful for provisioned tables with scaling."
  type        = bool
  default     = false
}

variable "import_table" {
  description = "Configurations for importing s3 data into a new table."
  type        = any
  default     = {}
}

variable "local_secondary_indexes" {
  description = "Describe an LSI on the table; these can only be allocated at creation so you cannot change this definition after you have created the resource."
  type        = any
  default     = []
}

variable "name" {
  description = "Name of the DynamoDB table."
  type        = string
  default     = null
}

variable "on_demand_throughput" {
  description = "Sets the maximum number of read and write units for the specified on-demand table."
  type        = any
  default     = {}
}

variable "point_in_time_recovery_enabled" {
  description = "Whether to enable point-in-time recovery."
  type        = bool
  default     = false
}

variable "point_in_time_recovery_period_in_days" {
  description = "Number of preceding days for which continuous backups are taken and maintained."
  type = number
  default = 7
}

variable "range_key" {
  description = "The attribute to use as the range (sort) key. Must also be defined as an attribute."
  type        = string
  default     = null
}

variable "read_capacity" {
  description = "The number of read units for this table. If the billing_mode is PROVISIONED, this field should be greater than 0."
  type        = number
  default     = null
}

variable "replica_regions" {
  description = "Region names for creating replicas for a global DynamoDB table."
  type        = any
  default     = []
}

variable "restore_date_time" {
  description = "Time of the point-in-time recovery point to restore."
  type        = string
  default     = null
}

variable "restore_source_name" {
  description = "Name of the table to restore. Must match the name of an existing table."
  type        = string
  default     = null
}

variable "restore_source_table_arn" {
  description = "ARN of the source table to restore. Must be supplied for cross-region restores."
  type        = string
  default     = null
}

variable "restore_to_latest_time" {
  description = "If set, restores table to the most recent point-in-time recovery point."
  type        = bool
  default     = null
}

variable "server_side_encryption_enabled" {
  description = "Whether or not to enable encryption at rest using an AWS managed KMS customer manage key (CMK)."
  type        = bool
  default     = false
}

variable "server_side_encryption_kms_key_arn" {
  description = "The ARN of the CMK that should be used for the AWS KMS encryption."
  type        = string
  default     = null
}

variable "stream_enabled" {
  description = "Indicates whether Streams are to be enabled (true) or disabled (false)."
  type        = bool
  default     = false
}

variable "stream_view_type" {
  description = "When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`."
  type        = string
  default     = null
}

variable "table_class" {
  description = "The storage class of the table. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`."
  type        = string
  default     = null
}

variable "timeouts" {
  description = "Updated Terraform resource management timeouts."
  type        = map(string)
  default = {
    create = "10m"
    update = "60m"
    delete = "10m"
  }
}

variable "ttl_enabled" {
  description = "Indicates whether ttl is enabled."
  type        = bool
  default     = false
}

variable "ttl_attribute_name" {
  description = "The name of the table attribute to store the TTL timestamp in."
  type        = string
  default     = ""
}

variable "write_capacity" {
  description = "The number of write units for this table. If the billing_mode is PROVISIONED, this field should be greater than 0."
  type        = number
  default     = null
}
