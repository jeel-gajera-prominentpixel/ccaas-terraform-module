locals {
  # Generates the EventBridge name using company prefix, application name, region suffix, and environment.
  eb_name = format("%s-evtbrg-%s-%s-%s", var.company_prefix, var.application, var.region_suffix, var.environment)

  # Default tags that will be applied to all resources unless overridden by custom tags.
  tags = module.tags.tags
}
