module "eventbridge" {
  source                        = "../../jb-ccaas-terraform-modules/terraform-aws-eventbridge"
  bus_name                      = var.name == "" ? local.eventbridge_name : var.name
  create_bus                    = var.create_bus
  create_schemas_discoverer     = var.create_schemas_discoverer
  attach_tracing_policy         = var.attach_tracing_policy
  attach_kinesis_policy         = var.attach_kinesis_policy
  kinesis_target_arns           = var.kinesis_target_arns
  attach_sfn_policy             = var.attach_sfn_policy
  sfn_target_arns               = var.sfn_target_arns
  attach_sqs_policy             = var.attach_sqs_policy
  sqs_target_arns               = var.sqs_target_arns
  attach_cloudwatch_policy      = var.attach_cloudwatch_policy
  cloudwatch_target_arns        = var.cloudwatch_target_arns
  append_rule_postfix           = var.append_rule_postfix
  attach_ecs_policy             = var.attach_ecs_policy
  ecs_target_arns               = var.ecs_target_arns
  rules                         = var.rules
  targets                       = var.targets
  api_destinations              = var.api_destinations
  create_api_destinations       = var.create_api_destinations
  attach_policy_json            = var.attach_policy_json
  policy_json                   = var.policy_json
  attach_policy_jsons           = var.attach_policy_jsons
  policy_jsons                  = var.policy_jsons
  number_of_policy_jsons        = var.number_of_policy_jsons
  attach_policies               = var.attach_policies
  policies                      = var.policies
  number_of_policies            = var.number_of_policies
  attach_policy_statements      = var.attach_policy_statements
  attach_api_destination_policy = var.attach_api_destination_policy
  create_connections            = var.create_connections
  connections                   = var.connections
  policy_statements             = var.policy_statements
  role_name                     = var.role_name
  role_description              = var.role_description
  role_path                     = var.role_path
  role_force_detach_policies    = var.role_force_detach_policies
  role_permissions_boundary     = var.role_permissions_boundary
  role_tags                     = var.role_tags
  create_pipes                  = var.create_pipes
  pipes                         = var.pipes
  append_pipe_postfix           = var.append_pipe_postfix

  tags = merge(local.tags, {
    Name = var.name == "" ? local.eventbridge_name : var.name
  })
}
