locals {
  # Generates the EventBridge name using company prefix, application name, region suffix, and environment.
  eb_name = var.name

  # Default tags that will be applied to all resources unless overridden by custom tags.
  tags = module.tags.tags
}
