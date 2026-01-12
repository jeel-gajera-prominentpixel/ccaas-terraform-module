module "transit-gateway-basic" {
  source                                = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-transit-gateway?ref=main"
  prefix_company                        = "jb"
  lob                                   = "itsd"
  prefix_region                         = "usw2"
  application                           = "recordings"
  env                                   = "sandbox"
  description                           = "Basic example TGW"
  amazon_side_asn                       = 64532
  transit_gateway_cidr_blocks           = ["10.100.0.0/16"]
  enable_auto_accept_shared_attachments = true
  vpc_attachments = {
    vpc1 = {
      vpc_id     = "vpc-02b3883c33254bc76"
      subnet_ids = ["subnet-04c1650f79a8a577c"]
    }
  }
}
