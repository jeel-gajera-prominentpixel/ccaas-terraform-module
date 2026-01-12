module "transit-gateway-advance" {
  source                                = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-transit-gateway?ref=main"
  prefix_company                        = "jb"
  lob                                   = "itsd"
  prefix_region                         = "usw2"
  application                           = "recordings"
  env                                   = "sandbox"
  description                           = "My TGW"
  amazon_side_asn                       = 64532
  transit_gateway_cidr_blocks           = ["10.100.0.0/16"]
  enable_auto_accept_shared_attachments = true
  enable_multicast_support              = false
  vpc_attachments = {
    vpc1 = {
      vpc_id     = "vpc-02b3883333253bc76"
      subnet_ids = ["subnet-04c1650f79a84577c"]

      dns_support  = true
      ipv6_support = false

      transit_gateway_default_route_table_association = false
      transit_gateway_default_route_table_propagation = false

      tgw_routes = [
        {
          destination_cidr_block = "30.0.0.0/16"
        },
        {
          blackhole              = true
          destination_cidr_block = "0.0.0.0/0"
        }
      ]
    }
  }
  ram_allow_external_principals = true
  ram_principals                = [307990089504]
  tags                          = local.tags
}
