# variable "application" {
#   description = "The name of the application or identifier."
#   type        = string
# }

variable "company" {
  description = "The name of the company."
  type        = string
}

variable "environment" {
  description = "The deployment environment."
  type        = string
}

variable "repository" {
  description = "The URL of the code repository."
  type        = string
}

# variable "created_by" {
#   description = "The identifier for the creator of the resources."
#   type        = string
#   default     = "Terraform"
# }

variable "project" {
  description = "The name of the project."
  type        = string
}

variable "region" {
  description = "Region of the resource."
  type        = string
}

variable "custom_tags" {
  description = "Additional custom tags to apply to resources, complementing the default tags."
  type        = map(string)
  default     = {}
}