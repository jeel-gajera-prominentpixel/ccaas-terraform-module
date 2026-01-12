data "aws_iam_policy_document" "bucket_policy" {

  source_policy_documents = concat(coalesce(local.policy, []))
  statement {
    sid    = "AllowSSLRequestsOnly"
    effect = "Deny"
    resources = [
      "arn:aws:s3:::${var.name == "" ? local.bucket_name : var.name}",
      "arn:aws:s3:::${var.name == "" ? local.bucket_name : var.name}/*",
    ]

    actions = ["s3:*"]

    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values   = ["false"]
    }

    principals {
      type        = "*"
      identifiers = ["*"]
    }
  }
}
