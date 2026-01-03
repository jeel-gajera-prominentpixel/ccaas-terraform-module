module "amazon_connect" {
  source = "../../ccaas-terraform-modules/terraform-aws-amazonconnect"

  instance_identity_management_type = var.instance_identity_management_type
  instance_inbound_calls_enabled    = true
  instance_outbound_calls_enabled   = true
  instance_alias                    = var.name == "" ? local.amazon_connect_name : var.name

  create_instance = var.create_instance
  instance_id     = var.instance_id
  # # Contact Flows / Modules
  contact_flows                      = var.contact_flows
  contact_flow_modules               = var.contact_flow_modules
  instance_contact_flow_logs_enabled = var.instance_contact_flow_logs_enabled
  instance_contact_lens_enabled      = var.instance_contact_lens_enabled
  multi_party_conference_enabled     = var.multi_party_conference_enabled
  hours_of_operations                = var.hours_of_operations
  hours_of_operations_tags           = var.hours_of_operations_tags
  queues                             = var.queues
  queue_tags                         = var.queue_tags
  user_hierarchy_structure           = var.user_hierarchy_structure
  user_hierarchy_groups              = var.user_hierarchy_groups

  # # Quick Connects
  quick_connects = var.quick_connects

  # # Routing / Security Profiles
  routing_profiles  = var.routing_profiles
  security_profiles = var.security_profiles

  instance_storage_configs = var.instance_storage_configs

  # # Lambda Function Associations
  lambda_function_associations = var.lambda_function_associations

  # # Users / Hierarchy Group / Structure
  users = var.users

  # # New variable to control lifecycle rule
  ignore_quick_connect_ids_changes = var.ignore_quick_connect_ids_changes

  # # Tags
  tags = local.tags
}

resource "awscc_connect_approved_origin" "this" {
  count = var.create_awscc_connect_approved_origin && length(var.approved_domains) > 0 ? length(var.approved_domains) : 0

  instance_id = module.amazon_connect.instance_arn
  origin      = var.approved_domains[count.index]
}
