module "s3_object" {
  source         = "../"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  create_object  = true
  bucket         = "cla-test"
  key            = "test_key"
}
