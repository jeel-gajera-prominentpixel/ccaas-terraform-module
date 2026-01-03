module "amazon_connect_advance" {
  source                       = "git@github.com:CloverHealth/ccaas-terraform-modules-wrapper.git//terraform-aws-amazonconnect?ref=v1.0.2"
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
