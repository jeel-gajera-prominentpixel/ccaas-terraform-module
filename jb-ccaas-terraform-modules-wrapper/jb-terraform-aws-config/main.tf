module "config" {
  source                           = "../../jb-ccaas-terraform-modules/terraform-aws-config"
  create_sns_topic                 = var.create_sns_topic
  create_iam_role                  = var.create_iam_role
  managed_rules                    = var.managed_rules
  force_destroy                    = var.force_destroy
  s3_bucket_id                     = var.s3_bucket_id
  s3_bucket_arn                    = var.s3_bucket_arn
  global_resource_collector_region = data.aws_region.current.name
  tags                             = local.tags
}
