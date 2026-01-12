module "secrets-manager" {
  source                           = "../../jb-ccaas-terraform-modules/terraform-aws-secrets-manager"
  name                             = var.name == "" ? local.secrets_manager_name : var.name
  name_prefix                      = var.name_prefix == "" ? local.secrets_manager_name : var.name_prefix
  description                      = var.description
  recovery_window_in_days          = var.recovery_window_in_days
  replica                          = var.replica
  create_policy                    = var.create_policy
  source_policy_documents          = var.source_policy_documents
  override_policy_documents        = var.override_policy_documents
  block_public_policy              = var.block_public_policy
  policy_statements                = var.policy_statements
  create_random_password           = var.create_random_password
  random_password_length           = var.random_password_length
  random_password_override_special = var.random_password_override_special
  ignore_secret_changes            = var.ignore_secret_changes
  enable_rotation                  = var.enable_rotation
  rotation_lambda_arn              = var.rotation_lambda_arn
  rotation_rules                   = var.rotation_rules
  secret_string                    = var.secret_string
  secret_binary                    = var.secret_binary
  version_stages                   = var.version_stages
  tags = merge(local.tags, {
    Name = var.name == "" ? local.secrets_manager_name : var.name
  })
}
