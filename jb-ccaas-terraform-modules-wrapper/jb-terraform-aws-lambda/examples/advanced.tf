module "lambda_function_advance" {
  source                     = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-lambda?ref=main"
  prefix_company             = "jb"
  lob                        = "itsd"
  prefix_region              = "usw2"
  application                = "recordings"
  env                        = "sandbox"
  description                = "jb function description"
  local_existing_package     = "../lambda_function_payload.zip"
  handler                    = "lambda_function.handler"
  runtime                    = "python3.8"
  store_on_s3                = false
  s3_bucket                  = ""
  s3_prefix                  = "lambda-builds/"
  create_role                = false
  lambda_role                = "arn:aws:iam::123456789012:role/lambda-execution-role"
  create_lambda_function_url = true
  environment_variables = {
    "ENV_VAR_1" = "value1"
    "ENV_VAR_2" = "value2"
  }
  timeout = 256
  tags    = local.tags
}
