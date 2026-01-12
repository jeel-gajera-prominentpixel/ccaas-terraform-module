module "acm" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-acm?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  application    = "recordings"
  env            = "sandbox"
  prefix_region  = "usw2"
  domain_name    = "example.com"
  zone_id        = "Z2ES7B9AZ6SHAE"
  subject_alternative_names = [
    "*.my-domain.com",
    "app.sub.my-domain.com",
  ]
  validation_method       = "DNS"
  validation_record_fqdns = ["example.com", "www.example.com"]
}
