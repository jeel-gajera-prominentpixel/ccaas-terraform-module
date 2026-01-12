module "pinpoint_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-amazon-pinpoint?ref=main"
  name           = "jb-pinpoint"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  email = {
    from     = "example@example.com"
    identity = "arn:aws:ses:us-west-2:123456789012:identity/example.com"
  }
  sms = {
    sender     = "example_sender"
    short_code = "12345"
  }
}
