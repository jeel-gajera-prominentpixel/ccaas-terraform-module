module "rest_api_gateway" {
  source = "../apigwv1" 

  api_name = var.api_name

  root_methods = {
    ANY = {
      authorization = "NONE"
      lambda_arn    = var.lambda_arn

      method_responses = {
        "200" = {
          response_models = {
            "application/json" = "Empty"
          }

          response_parameters = {
            "method.response.header.Access-Control-Allow-Origin" = true
          }
        }
      }
    }
  }

  proxy_methods = {
    ANY = {
      authorization = "NONE"
      lambda_arn    = var.lambda_arn
    }
  }
}
