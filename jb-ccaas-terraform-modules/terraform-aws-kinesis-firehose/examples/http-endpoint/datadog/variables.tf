variable "name_prefix" {
  description = "Name prefix to use in resources"
  type        = string
  default     = "firehose-to-datadog"
}

variable "datadog_api_key" {
  description = "Datadog Api Key"
  type        = string
  sensitive   = true
}
