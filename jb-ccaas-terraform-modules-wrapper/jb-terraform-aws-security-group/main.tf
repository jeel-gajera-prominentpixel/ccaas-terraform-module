module "security_group" {
  source                                = "../../jb-ccaas-terraform-modules/terraform-aws-security-group"
  name                                  = var.name == "" ? local.sg_name : var.name
  create                                = var.create
  create_sg                             = var.create_sg
  security_group_id                     = var.security_group_id
  vpc_id                                = var.vpc_id
  description                           = var.description
  create_timeout                        = var.create_timeout
  delete_timeout                        = var.delete_timeout
  ingress_rules                         = var.ingress_rules
  ingress_cidr_blocks                   = var.ingress_cidr_blocks
  ingress_ipv6_cidr_blocks              = var.ingress_ipv6_cidr_blocks
  ingress_prefix_list_ids               = var.ingress_prefix_list_ids
  ingress_with_self                     = var.ingress_with_self
  ingress_with_cidr_blocks              = var.ingress_with_cidr_blocks
  ingress_with_ipv6_cidr_blocks         = var.ingress_with_ipv6_cidr_blocks
  ingress_with_source_security_group_id = var.ingress_with_source_security_group_id
  ingress_with_prefix_list_ids          = var.ingress_with_prefix_list_ids
  egress_rules                          = var.egress_rules
  egress_cidr_blocks                    = var.egress_cidr_blocks
  egress_ipv6_cidr_blocks               = var.egress_ipv6_cidr_blocks
  egress_prefix_list_ids                = var.egress_prefix_list_ids
  egress_with_self                      = var.egress_with_self
  egress_with_cidr_blocks               = var.egress_with_cidr_blocks
  egress_with_ipv6_cidr_blocks          = var.egress_with_ipv6_cidr_blocks
  egress_with_source_security_group_id  = var.egress_with_source_security_group_id
  egress_with_prefix_list_ids           = var.egress_with_prefix_list_ids
  use_name_prefix                       = var.use_name_prefix
  tags = merge(local.tags, {
    Name = var.name == "" ? local.sg_name : var.name
  })
}
