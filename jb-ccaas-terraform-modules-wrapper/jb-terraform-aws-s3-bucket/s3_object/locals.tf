data "external" "env" {
  program = ["${path.module}/scripts/env.sh"]
}

locals {
  tags = merge(
    var.tags,
    {
      module_project_path = local.module_project_path,
      module_version      = local.module_version,
      project_path        = data.external.env.result["project_path"]
      company             = var.prefix_company
      lob                 = var.lob
      application         = var.application
      created_by          = "terraform"
      map-migrated        = "migVSN3WXHRBU"
    },
  )
}
