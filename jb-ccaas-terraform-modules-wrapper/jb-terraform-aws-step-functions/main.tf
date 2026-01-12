module "step_function" {
  source                            = "../../jb-ccaas-terraform-modules/terraform-aws-step-functions"
  name                              = var.name == "" ? local.step_function_name : var.name
  definition                        = var.definition
  type                              = var.type
  publish                           = var.publish
  logging_configuration             = var.logging_configuration
  service_integrations              = var.service_integrations
  attach_policy_json                = var.attach_policy_json
  policy_json                       = var.policy_json
  cloudwatch_log_group_name         = var.cloudwatch_log_group_name
  use_existing_cloudwatch_log_group = var.use_existing_cloudwatch_log_group
  tags = merge(local.tags, {
    Name = var.name == "" ? local.step_function_name : var.name
  })

}
