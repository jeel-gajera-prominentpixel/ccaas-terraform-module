module "eventbridge" {
  source = "../terraform-aws-eventbridge-v4.2.2"

  # Core
  bus_name   = var.bus_name
  create_bus = var.bus_name == "default" ? false : var.create_bus
  create     = true

  # Rules & Targets
  rules   = var.rules
  targets = var.targets

  # Policies
  policies              = var.policies
  policy                = var.policy
  policy_json           = var.policy_json
  policy_jsons          = var.policy_jsons
  policy_statements     = var.policy_statements
  number_of_policies    = var.number_of_policies
  number_of_policy_jsons = var.number_of_policy_jsons

  # API Destinations & Connections
  api_destinations        = var.api_destinations
  create_api_destinations = var.create_api_destinations
  connections             = var.connections
  create_connections      = var.create_connections

  # Archives
  archives        = var.archives
  create_archives = var.create_archives

  # Permissions
  permissions        = var.permissions
  create_permissions = var.create_permissions

  # Pipes
  pipes        = var.pipes
  create_pipes = var.create_pipes
  append_pipe_postfix = var.append_pipe_postfix

  # Schedules
  schedules               = var.schedules
  schedule_groups         = var.schedule_groups
  schedule_group_timeouts = var.schedule_group_timeouts
  create_schedule_groups  = var.create_schedule_groups
  create_schedules        = var.create_schedules
  append_schedule_group_postfix = var.append_schedule_group_postfix
  append_schedule_postfix       = var.append_schedule_postfix

  # Rules Creation
  create_rules   = var.create_rules
  create_targets = var.create_targets
  append_rule_postfix = var.append_rule_postfix

  # Schema Discoverer
  create_schemas_discoverer      = var.create_schemas_discoverer
  schemas_discoverer_description = var.schemas_discoverer_description

  # Event Source
  event_source_name = var.event_source_name

  # IAM Role
  create_role                 = var.create_role
  role_name                   = var.role_name
  role_description            = var.role_description
  role_path                   = var.role_path
  role_force_detach_policies  = var.role_force_detach_policies
  role_permissions_boundary   = var.role_permissions_boundary
  role_tags                   = var.role_tags
  trusted_entities            = var.trusted_entities

  # Attach Policies
  attach_api_destination_policy  = var.attach_api_destination_policy
  attach_cloudwatch_policy       = var.attach_cloudwatch_policy
  attach_ecs_policy              = var.attach_ecs_policy
  attach_kinesis_firehose_policy = var.attach_kinesis_firehose_policy
  attach_kinesis_policy          = var.attach_kinesis_policy
  attach_lambda_policy           = var.attach_lambda_policy
  attach_policies                = var.attach_policies
  attach_policy                  = var.attach_policy
  attach_policy_json             = var.attach_policy_json
  attach_policy_jsons            = var.attach_policy_jsons
  attach_policy_statements       = var.attach_policy_statements
  attach_sfn_policy              = var.attach_sfn_policy
  attach_sns_policy              = var.attach_sns_policy
  attach_sqs_policy              = var.attach_sqs_policy
  attach_tracing_policy          = var.attach_tracing_policy

  # Target ARNs
  cloudwatch_target_arns       = var.cloudwatch_target_arns
  ecs_target_arns              = var.ecs_target_arns
  kinesis_firehose_target_arns = var.kinesis_firehose_target_arns
  kinesis_target_arns          = var.kinesis_target_arns
  lambda_target_arns           = var.lambda_target_arns
  sfn_target_arns              = var.sfn_target_arns
  sns_target_arns              = var.sns_target_arns
  sqs_target_arns              = var.sqs_target_arns

  # Postfix Controls
  append_connection_postfix  = var.append_connection_postfix
  append_destination_postfix = var.append_destination_postfix

  # Logging
  create_log_delivery        = var.create_log_delivery
  create_log_delivery_source = var.create_log_delivery_source

  # Tags
  tags = merge(local.tags, var.tags)
}
