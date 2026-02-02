module "apigv1_ad_create" {
  source                       = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-apigwv1?ref=main"
  prefix_company               = "cla"
  lob                          = "itsd"
  prefix_region                = "use1"
  application                  = "connect"
  env                          = "sandbox"
  name                         = "example-api-module"
  description                  = "test desc"
  authorization                = "NONE"
  types                        = "REGIONAL"
  stage_name                   = "dev"
  resource_root_path           = "ANY"
  root_lambda_arn              = "arn:aws:lambda:us-west-2:xxxxxxxxxxxx:function:xxxxxxxxxxxx"
  root_integration_http_method = "POST"
  root_integration_type        = "AWS_PROXY"
  resource_paths = {
    "{proxy+}" = {
      lambda_arn              = "arn:aws:lambda:us-west-2:xxxxxxxxxxxx:function:xxxxxxxxxxxx"
      http_method             = "ANY"
      integration_http_method = "POST"
      type                    = "AWS_PROXY"
    }
    "test" = {
      lambda_arn              = "arn:aws:lambda:us-west-2:xxxxxxxxxxxx:function:xxxxxxxxxxxx"
      http_method             = "GET"
      integration_http_method = "POST"
      type                    = "AWS_PROXY"
    }
  }
}
