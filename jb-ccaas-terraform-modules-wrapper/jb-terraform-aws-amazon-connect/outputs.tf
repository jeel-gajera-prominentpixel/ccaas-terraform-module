################################################################################
# Instance
################################################################################
output "instance" {
  description = "Full output attributes of aws_connect_instance resource."
  value       = module.amazon_connect.instance
}

output "current_region" {
  description = "Current AWS region where the module is being executed."
  value       = data.aws_region.current.name
}

output "current_account_number" {
  description = "Current AWS account number where the module is being executed."
  value       = data.aws_caller_identity.current.account_id
}

output "instance_id" {
  description = "Amazon Connect instance ID. If create_instance = false, var.instance_id is returned."
  value       = module.amazon_connect.instance_id
}

output "instance_arn" {
  description = "Amazon Connect instance ARN."
  value       = module.amazon_connect.instance_arn
}
################################################################################
# Instance Storage Config
################################################################################
output "instance_storage_configs" {
  description = "Full output attributes of instance_storage_configs resource(s)."
  value       = module.amazon_connect.instance_storage_configs
}

################################################################################
# Hours of Operation
################################################################################
output "hours_of_operations" {
  description = "Full output attributes of hours_of_operations resource(s)."
  value       = module.amazon_connect.hours_of_operations
}

################################################################################
# Contact Flow / Module
################################################################################
output "contact_flows" {
  description = "Full output attributes of contact_flows resource(s)."
  value       = module.amazon_connect.contact_flows
}

output "contact_flow_modules" {
  description = "Full output attributes of contact_flow_modules resource(s)."
  value       = module.amazon_connect.contact_flow_modules
}

################################################################################
# Queue
################################################################################
output "queues" {
  description = "Full output attributes of aws_connect_queue resource(s)."
  value       = module.amazon_connect.queues
}

################################################################################
# Quick Connect
################################################################################
output "quick_connects" {
  description = "Full output attributes of quick_connects resource(s)."
  value       = module.amazon_connect.quick_connects
}

################################################################################
# Routing / Security Profiles
################################################################################
output "routing_profiles" {
  description = "Full output attributes of routing_profiles resource(s)."
  value       = module.amazon_connect.routing_profiles
}

output "security_profiles" {
  description = "Full output attributes of security_profiles resource(s)."
  value       = module.amazon_connect.security_profiles
}

################################################################################
# Vocabulary
################################################################################
output "vocabularies" {
  description = "Full output attributes of vocabularies resource(s)."
  value       = module.amazon_connect.vocabularies
}

################################################################################
# Lex Bot / Lambda Function Associations
################################################################################
output "bot_associations" {
  description = "Full output attributes of bot_associations resource(s)."
  value       = module.amazon_connect.bot_associations
}

output "lambda_function_associations" {
  description = "Full output attributes of lambda_function_associations resource(s)."
  value       = module.amazon_connect.lambda_function_associations
}

################################################################################
# Users / Hierarchy Group / Structure
################################################################################
output "users" {
  description = "Full output attributes of users resource(s)."
  value       = module.amazon_connect.users
  sensitive   = true
}

output "user_hierarchy_groups" {
  description = "Full output attributes of user_hierarchy_groups resource(s)."
  value       = module.amazon_connect.user_hierarchy_groups
}

output "user_hierarchy_structure" {
  description = "Full output attributes of user_hierarchy_structure resource(s)."
  value       = module.amazon_connect.user_hierarchy_structure
}

################################################################################
# SSO User output
################################################################################
output "connect_idp_admin_arn" {
  description = "The ARN of connect admin panel"
  value       = try(aws_iam_saml_provider.AWSSSO_admin_Connect[0].arn, "")
}

output "connect_idp_agent_arn" {
  description = "The ARN of connect agent panel"
  value       = try(aws_iam_saml_provider.AWSSSO_agent_Connect[0].arn, "")
}
