variable "application" {
  type        = string
  description = "The application name of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name."
}

variable "prefix_company" {
  type        = string
  description = "The prefix company of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name."
}

variable "prefix_region" {
  type        = string
  description = "The prefix region of the aws transit gateway, will be appended with the company, lob, env and region to form a acm name."
}

variable "env" {
  type        = string
  description = "Environment name."
}

variable "lob" {
  type        = string
  description = "The lob name of the transit gateway, will be appended with the company, lob, env and region to form a transit gateway name."
}

variable "tags" {
  description = "A map of tags to assign to the resources created by this module. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags) present, tags with matching keys will overwrite those defined at the provider-level."
  type        = map(string)
  default     = {}
}

variable "description" {
  description = "Description of the Transit Gateway"
  type        = string
  default     = null
}

variable "amazon_side_asn" {
  description = "The Autonomous System Number (ASN) for the Amazon side of the gateway. By default the TGW is created with the current default Amazon ASN."
  type        = string
  default     = null
}

variable "transit_gateway_cidr_blocks" {
  description = "One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6"
  type        = list(string)
  default     = []
}

variable "enable_auto_accept_shared_attachments" {
  description = "Whether resource attachment requests are automatically accepted"
  type        = bool
  default     = false
}

variable "enable_multicast_support" {
  description = "Whether multicast support is enabled"
  type        = bool
  default     = false
}

variable "vpc_attachments" {
  description = "Maps of maps of VPC details to attach to TGW. Type 'any' to disable type validation by Terraform."
  type        = any
  default     = {}
}

variable "ram_allow_external_principals" {
  description = "Indicates whether principals outside your organization can be associated with a resource share."
  type        = bool
  default     = false
}

variable "ram_principals" {
  description = "A list of principals to share TGW with. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN"
  type        = list(string)
  default     = []
}

variable "name" {
  type        = string
  description = "Provide resource name if you want to override with wrapper"
  default     = ""
}
