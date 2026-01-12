module "lambda_function" {
  # https://github.com/terraform-aws-modules/terraform-aws-s3-bucket
  source                                    = "../../jb-ccaas-terraform-modules/terraform-aws-lambda"
  function_name                             = var.name == "" ? local.lambda_name : var.name
  description                               = var.description
  handler                                   = var.handler
  runtime                                   = var.runtime
  create_package                            = false
  local_existing_package                    = var.local_existing_package
  store_on_s3                               = var.store_on_s3
  s3_bucket                                 = var.s3_bucket
  s3_prefix                                 = var.s3_prefix
  environment_variables                     = var.environment_variables
  ignore_source_code_hash                   = var.ignore_source_code_hash
  timeout                                   = var.timeout
  publish                                   = var.publish
  create_role                               = var.create_role
  lambda_role                               = var.lambda_role
  create_lambda_function_url                = var.create_lambda_function_url
  attach_policy_statements                  = var.attach_policy_statements
  policy_statements                         = var.policy_statements
  vpc_subnet_ids                            = var.vpc_subnet_ids
  vpc_security_group_ids                    = var.vpc_security_group_ids
  memory_size                               = var.memory_size
  ephemeral_storage_size                    = var.ephemeral_storage_size
  architectures                             = var.architectures
  allowed_triggers                          = var.allowed_triggers
  authorization_type                        = var.authorization_type
  cors                                      = var.cors
  invoke_mode                               = var.invoke_mode
  compatible_runtimes                       = var.compatible_runtimes
  compatible_architectures                  = var.compatible_architectures
  trusted_entities                          = var.trusted_entities
  create_layer                              = var.create_layer
  layer_name                                = var.layer_name
  layers                                    = var.layers
  attach_network_policy                     = var.attach_network_policy
  replace_security_groups_on_destroy        = var.replace_security_groups_on_destroy
  replacement_security_group_ids            = var.replacement_security_group_ids
  attach_policy_jsons                       = var.attach_policy_jsons
  role_name                                 = var.role_name
  policy_name                               = var.policy_name
  policy_jsons                              = var.policy_jsons
  number_of_policy_jsons                    = var.number_of_policy_jsons
  create_current_version_allowed_triggers   = var.create_current_version_allowed_triggers
  create_unqualified_alias_allowed_triggers = var.create_unqualified_alias_allowed_triggers
  tracing_mode                              = var.tracing_mode
  event_source_mapping                      = var.event_source_mapping
  tags = merge(local.tags, {
    Name = var.name == "" ? local.lambda_name : var.name
  })
}
