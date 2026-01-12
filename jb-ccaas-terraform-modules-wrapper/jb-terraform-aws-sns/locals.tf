data "external" "env" {
  program = ["${path.module}/scripts/env.sh"]
}

locals {
  sns_name = format("%s-sns-%s-%s-%s-%s", var.prefix_company, var.lob, var.application, var.prefix_region, var.env)
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
    },
  )
}
