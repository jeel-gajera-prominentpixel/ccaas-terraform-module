module "transit-gateway" {
  source                                = "../../jb-ccaas-terraform-modules/terraform-aws-transit-gateway"
  name                                  = var.name == "" ? local.transit_gateway_name : var.name
  description                           = var.description
  amazon_side_asn                       = var.amazon_side_asn
  transit_gateway_cidr_blocks           = var.transit_gateway_cidr_blocks
  enable_auto_accept_shared_attachments = var.enable_auto_accept_shared_attachments
  enable_multicast_support              = var.enable_multicast_support
  vpc_attachments                       = var.vpc_attachments
  ram_allow_external_principals         = var.ram_allow_external_principals
  ram_principals                        = var.ram_principals
  tags = merge(local.tags, {
    Name = var.name == "" ? local.transit_gateway_name : var.name
  })
}
