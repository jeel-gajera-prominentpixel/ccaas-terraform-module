
locals {
  lex_name = format("%s-lex-%s-%s-%s-%s", var.prefix_company, var.lob, var.application, var.prefix_region, var.env)
  bot_tags = concat(var.bot_tags,
    [
      {
        key   = "module_project_path"
        value = local.module_project_path
      },
      {
        key   = "module_version"
        value = local.module_version
      },
      {
        key   = "company"
        value = var.prefix_company
      },
      {
        key   = "region"
        value = var.prefix_region
      },
      {
        key   = "lob"
        value = var.lob
      },
      {
        key   = "application"
        value = var.application
      },
      {
        key   = "env"
        value = var.env
      },
      {
        key   = "created_by"
        value = "terraform"
      },
      {
        key   = "map-migrated"
        value = "migVSN3WXHRBU"
      }
  ])
}
