module "iam_policy" {
  source        = "../../jb-ccaas-terraform-modules/terraform-aws-iam/modules/iam-policy"
  name          = var.name == "" ? local.iam_policy : var.name
  create_policy = var.create_policy
  description   = var.description
  name_prefix   = var.name_prefix
  path          = var.path
  policy        = var.policy
  tags = merge(local.tags, {
    Name = var.name == "" ? local.iam_policy : var.name
  })

}
