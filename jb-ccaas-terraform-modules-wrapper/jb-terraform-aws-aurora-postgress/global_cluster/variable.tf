variable "application" {
  type        = string
  description = "The application name of the rds, will be appended with the company, lob, env and region to form a rds name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the rds, will be appended with the company, lob, env and region to form a rds name"
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the rds , will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "lob" {
  type        = string
  description = "The lob name of the rds, will be appended with the company, lob, env and region to form a rds name"
}

variable "create_global_cluster" {
  description = "Whether global cluster should be created"
  type        = bool
  default     = false
}

variable "global_cluster_identifier" {
  description = "Global cluster identifier"
  type        = string
  default     = ""
}

variable "global_cluster_engine" {
  description = "The name of the database engine to be used for this global cluster. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`"
  type        = string
  default     = null
}

variable "global_cluster_version" {
  description = "The database engine version for global cluster. Updating this argument results in an outage"
  type        = string
  default     = null
}

variable "global_cluster_db_name" {
  description = "Name for an automatically created database on global cluster creation"
  type        = string
  default     = null
}

variable "global_cluster_storage_encrypted" {
  description = "Specifies whether the DB cluster is encrypted. The default is `true`"
  type        = bool
  default     = true
}


variable "source_db_cluster_identifier" {
  description = "(Optional) Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. Terraform cannot perform drift detection of this value."
  type        = string
  default     = null
}

variable "deletion_protection" {
  description = "(Optional) If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false."
  type        = bool
  default     = false
}

variable "force_destroy" {
  type        = bool # This is a boolean since it's an enable/disable flag
  description = "(Optional) Enable to remove DB Cluster members from Global Cluster on destroy. Required with source_db_cluster_identifier."
  default     = false # Optional, typically defaults to false
}

variable "engine_lifecycle_support" {
  type        = string # This would be a string to specify the engine lifecycle support level
  description = "(Optional) The life cycle type for this DB instance. This setting applies only to Aurora PostgreSQL-based global databases. Valid values are open-source-rds-extended-support, open-source-rds-extended-support-disabled. Default value is open-source-rds-extended-support. [Using Amazon RDS Extended Support]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html"
  default     = null
}
