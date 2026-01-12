module "aws_kms" {
  # https://github.com/terraform-aws-modules/terraform-aws-kms
  source                                 = "../../jb-ccaas-terraform-modules/terraform-aws-kms"
  description                            = var.description
  multi_region                           = var.multi_region
  enable_default_policy                  = var.enable_default_policy
  key_owners                             = [local.current_identity]
  key_administrators                     = [local.current_identity]
  key_users                              = [local.current_identity]
  key_service_users                      = [local.current_identity]
  key_symmetric_encryption_users         = [local.current_identity]
  key_hmac_users                         = [local.current_identity]
  key_asymmetric_public_encryption_users = [local.current_identity]
  key_asymmetric_sign_verify_users       = [local.current_identity]
  primary_key_arn                        = var.primary_key_arn
  create_replica                         = var.create_replica
  aliases                                = [var.name == "" ? local.Key_name : var.name]
  aliases_use_name_prefix                = false
  key_statements                         = var.key_statements

  tags = merge(local.tags, {
    Name = var.name == "" ? local.Key_name : var.name
  })
}
