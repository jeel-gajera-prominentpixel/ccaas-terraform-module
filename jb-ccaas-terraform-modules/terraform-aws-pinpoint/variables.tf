variable "name" {
  type        = string
  description = "(Required) Project name"
}

variable "email" {
  type = object({
    from     = string
    identity = string
  })
  default     = null
  description = "Provides a Pinpoint Email Channel resource."
}

variable "sms" {
  type = object({
    sender     = string
    short_code = string
  })
  default     = null
  description = "Provides a Pinpoint SMS Channel resource."
}

variable "baidu" {
  type = object({
    api_key    = string
    secret_key = string
  })
  default     = null
  description = "Provides a Pinpoint Baidu Channel resource."
}

variable "apns" {
  type = object({
    certificate  = string
    private_key  = string
    bundle_id    = string
    team_id      = string
    token_key    = string
    token_key_id = string
  })
  default     = null
  description = "Provides a Pinpoint APNs Channel resource."
}

variable "apns_sandbox" {
  type = object({
    certificate  = string
    private_key  = string
    bundle_id    = string
    team_id      = string
    token_key    = string
    token_key_id = string
  })
  default     = null
  description = "Provides a Pinpoint APNs Sandbox Channel resource."
}

variable "apns_voip" {
  type = object({
    certificate  = string
    private_key  = string
    bundle_id    = string
    team_id      = string
    token_key    = string
    token_key_id = string
  })
  default     = null
  description = "Provides a Pinpoint APNs VoIP Channel resource."
}

variable "apns_voip_sandbox" {
  type = object({
    certificate  = string
    private_key  = string
    bundle_id    = string
    team_id      = string
    token_key    = string
    token_key_id = string
  })
  default     = null
  description = "Provides a Pinpoint APNs VoIP Sandbox Channel resource."
}
