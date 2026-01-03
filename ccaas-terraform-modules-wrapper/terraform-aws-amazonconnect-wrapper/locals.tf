# To retrieve the client-consumer project path from GitLab's $CI_PROJECT_PATH
data "external" "env" {
  program = ["${path.module}/scripts/env.sh"]
}

locals {
  amazon_connect_name     = format("%s-connect-%s-%s", var.company_prefix, var.region_prefix, var.env)
  amazon_connect_iam_name = format("%s-connect-%s-%s", var.company_prefix, var.region_prefix, var.env)
  tags = merge(
    var.tags,
    {
      project_path = data.external.env.result["project_path"]
      company      = var.company_prefix
      region       = var.region_prefix
      env          = var.env
      created_by   = "terraform"
    },
  )
}
