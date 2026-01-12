data "external" "env" {
  program = ["${path.module}/scripts/env.sh"]
}

locals {
  pinpoint_name = format("%s-pinpt-%s-%s-%s-%s", var.prefix_company, var.lob, var.application, var.prefix_region, var.env)
}
