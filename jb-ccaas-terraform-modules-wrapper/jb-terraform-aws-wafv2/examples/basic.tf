module "waf2_create" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-wafv2?ref=main"
  prefix_company = "cla"
  lob            = "itsd"
  prefix_region  = "use1"
  application    = "connect"
  env            = "sandbox"
  name           = "testing-waf"
  description    = "test own wrapper module"
  scope          = "CLOUDFRONT"
  default_action = "allow"
  waf_rules = [
    {
      name                       = "ownwrapper-rule-v4"
      priority                   = 1
      sampled_requests_enabled   = true
      cloudwatch_metrics_enabled = true
      action                     = "allow"    # "allow" or "block"
      ip_set_name                = "own-ipv4" # Name of the IP set to associate with the rule
      metric_name                = "allow-specified-ipv4"
    },
    {
      name                       = "ownwrapper-rule-v6"
      priority                   = 2
      sampled_requests_enabled   = true
      cloudwatch_metrics_enabled = true
      action                     = "allow"    # "allow" or "block"
      ip_set_name                = "own-ipv6" # Name of the IP set to associate with the rule
      metric_name                = "allow-specified-ipv6"
    }
  ]
  waf_ip_sets = [
    {
      name               = "own-ipv4"
      ip_address_version = "IPV4"
      addresses_list = [
        "x.x.x.x/x",
      ]
    },
    {
      name               = "own-ipv6"
      ip_address_version = "IPV6"
      addresses_list = [
        "x:x:x:x::/x",
      ]
    }
  ]
  visibility_config = {
    cloudwatch_metrics_enabled = true
    metric_name                = "cloudfront-waf-acl"
    sampled_requests_enabled   = true
  }
  tags = {}
}
