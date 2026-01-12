module "amazon_connect_advance" {
  source                       = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-amazon-connect?ref=main"
  prefix_company               = "jb"
  lob                          = "itsd"
  application                  = "recordings"
  prefix_region                = "usw2"
  env                          = "sandbox"
  bucket_name                  = "test-bucket"
  contact_flows                = {}
  quick_connects               = {}
  lambda_function_associations = {}
  contact_flow_modules         = {}
  routing_profiles             = {}
  security_profiles            = {}
  users                        = {}
  tags                         = local.tags
}
