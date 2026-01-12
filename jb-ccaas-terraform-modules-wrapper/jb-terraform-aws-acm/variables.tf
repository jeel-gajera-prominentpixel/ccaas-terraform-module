variable "application" {
  type        = string
  description = "The application name of the acm, will be appended with the company, lob, env and region to form a acm name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the acm, will be appended with the company, lob, env and region to form a acm name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the acm, will be appended with the company, lob, env and region to form a acm name."

}
variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the acm, will be appended with the company, lob, env and region to form a acm name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "zone_id" {
  description = "The ID of the hosted zone to contain this record. Required when validating via Route53"
  type        = string
  default     = ""
}

variable "domain_name" {
  description = "A domain name for which the certificate should be issued"
  type        = string
  default     = ""
}

variable "subject_alternative_names" {
  description = "A list of domains that should be SANs in the issued certificate"
  type        = list(string)
  default     = []
}

variable "validation_method" {
  description = "Which method to use for validation. DNS or EMAIL are valid. This parameter must not be set for certificates that were imported into ACM and then into Terraform."
  type        = string
  default     = null
}

variable "create_route53_records" {
  description = "When validation is set to DNS, define whether to create the DNS records internally via Route53 or externally using any DNS provider"
  type        = bool
  default     = true
}

variable "validation_record_fqdns" {
  description = "When validation is set to DNS and the DNS validation records are set externally, provide the fqdns for the validation"
  type        = list(string)
  default     = []
}

variable "wait_for_validation" {
  description = "Whether to wait for the validation to complete"
  type        = bool
  default     = true
}
