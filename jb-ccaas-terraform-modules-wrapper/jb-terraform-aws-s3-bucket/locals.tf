# To retrieve the client-consumer project path from GitLab's $CI_PROJECT_PATH
data "external" "env" {
  program = ["${path.module}/scripts/env.sh"]
}

locals {

  bucket_name = format("%s-s3-%s-%s-%s-%s", var.prefix_company, var.lob, var.application, var.prefix_region, var.env)
  policy      = var.policy != null ? [var.policy] : null
  tags = merge(
    var.tags,
    {
      module_project_path = local.module_project_path,
      module_version      = local.module_version,
      project_path        = data.external.env.result["project_path"]
      company             = var.prefix_company
      region              = var.prefix_region
      lob                 = var.lob
      application         = var.application
      env                 = var.env
      created_by          = "terraform"
      map-migrated        = "migVSN3WXHRBU"
      data_classification = var.data_classification
    },
  )

  grant = [
    {
      type       = "CanonicalUser"
      permission = "FULL_CONTROL"
      id         = data.aws_canonical_user_id.current.id
    },
    {
      type       = "CanonicalUser"
      permission = "FULL_CONTROL"
      id         = data.aws_cloudfront_log_delivery_canonical_user_id.cloudfront.id
    },
    {
      type       = "Group"
      permission = "WRITE"
      uri        = "http://acs.amazonaws.com/groups/s3/LogDelivery"
    },
    {
      type       = "Group"
      permission = "READ_ACP"
      uri        = "http://acs.amazonaws.com/groups/s3/LogDelivery"
    }
  ]
}
