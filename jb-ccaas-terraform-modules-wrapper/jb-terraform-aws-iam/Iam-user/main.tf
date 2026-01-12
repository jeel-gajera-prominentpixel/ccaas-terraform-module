module "iam_user" {
  source                        = "../../jb-ccaas-terraform-modules/terraform-aws-iam/modules/iam-user"
  name                          = var.name == "" ? local.iam_user : var.name
  create_user                   = var.create_user
  create_iam_user_login_profile = var.create_iam_user_login_profile
  create_iam_access_key         = var.create_iam_access_key
  path                          = var.path
  force_destroy                 = var.force_destroy
  pgp_key                       = var.pgp_key
  iam_access_key_status         = var.iam_access_key_status
  password_reset_required       = var.password_reset_required
  password_length               = var.password_length
  upload_iam_user_ssh_key       = var.upload_iam_user_ssh_key
  ssh_key_encoding              = var.ssh_key_encoding
  ssh_public_key                = var.ssh_public_key
  permissions_boundary          = var.permissions_boundary
  policy_arns                   = var.policy_arns
  tags = merge(local.tags, {
    Name = var.name == "" ? local.iam_user : var.name
  })

}
