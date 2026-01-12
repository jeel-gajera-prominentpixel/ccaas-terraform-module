module "s3_bucket" {
  # https://github.com/terraform-aws-modules/terraform-aws-s3-bucket
  source                   = "../../jb-ccaas-terraform-modules/terraform-aws-s3-bucket"
  bucket                   = var.name == "" ? local.bucket_name : var.name
  force_destroy            = var.force_destroy
  acl                      = var.enable_grant == true ? null : var.acl
  block_public_acls        = var.block_public_acls
  block_public_policy      = var.block_public_policy
  ignore_public_acls       = var.ignore_public_acls
  restrict_public_buckets  = var.restrict_public_buckets
  object_ownership         = var.object_ownership
  control_object_ownership = true
  attach_policy            = true
  create_bucket            = var.create_bucket
  policy                   = data.aws_iam_policy_document.bucket_policy.json
  grant                    = var.enable_grant == true ? local.grant : var.grant
  server_side_encryption_configuration = {
    rule = {
      apply_server_side_encryption_by_default = {
        kms_master_key_id = var.kms_key_id
        sse_algorithm     = var.kms_key_id != null ? "aws:kms" : "AES256"
      }
    }
  }
  versioning = {
    enabled = var.versioning
  }

  logging = var.s3_logging_bucket_id == null ? {} : {
    target_bucket = var.s3_logging_bucket_id
    # target_bucket = data.aws_s3_bucket.logging.id
    target_prefix = format("S3Logs/%s/", var.name == "" ? local.bucket_name : var.name)
  }
  website                                  = var.website
  cors_rule                                = var.cors_rule
  lifecycle_rule                           = var.lifecycle_rule
  attach_elb_log_delivery_policy           = var.attach_elb_log_delivery_policy
  attach_lb_log_delivery_policy            = var.attach_lb_log_delivery_policy
  attach_access_log_delivery_policy        = var.attach_access_log_delivery_policy
  attach_deny_insecure_transport_policy    = var.attach_deny_insecure_transport_policy
  attach_require_latest_tls_policy         = var.attach_require_latest_tls_policy
  attach_public_policy                     = var.attach_public_policy
  attach_inventory_destination_policy      = var.attach_inventory_destination_policy
  attach_analytics_destination_policy      = var.attach_analytics_destination_policy
  attach_deny_incorrect_encryption_headers = var.attach_deny_incorrect_encryption_headers
  attach_deny_incorrect_kms_key_sse        = var.attach_deny_incorrect_kms_key_sse
  allowed_kms_key_arn                      = var.allowed_kms_key_arn
  attach_deny_unencrypted_object_uploads   = var.attach_deny_unencrypted_object_uploads
  replication_configuration                = var.replication_configuration
  tags = merge(local.tags, {
    Name = var.name == "" ? local.bucket_name : var.name
  })
}

module "s3_bucket_notifications" {
  count                = var.lambda_trigger ? 1 : 0
  source               = "../../jb-ccaas-terraform-modules/terraform-aws-s3-bucket/modules/notification"
  bucket               = var.name == "" ? local.bucket_name : var.name
  lambda_notifications = var.lambda_notifications
  sqs_notifications    = var.sqs_notifications
  sns_notifications    = var.sns_notifications
  create_sqs_policy    = var.create_sqs_policy
}
