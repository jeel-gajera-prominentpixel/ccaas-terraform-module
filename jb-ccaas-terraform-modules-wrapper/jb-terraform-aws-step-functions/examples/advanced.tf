module "step_function_advanced" {
  source         = "git@github.com:jetblueairways/ccaas-terraform-modules-wrapper.git//jb-terraform-aws-step-functions?ref=main"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "jb-step_function"
  definition     = <<EOF
{
  "Comment": "A Hello World example of the Amazon States Language using Pass states",
  "StartAt": "Hello",
  "States": {
    "Hello": {
      "Type": "Pass",
      "Result": "Hello",
      "Next": "World"
    },
    "World": {
      "Type": "Pass",
      "Result": "World",
      "End": true
    }
  }
}
EOF

  type    = "express"
  publish = true
  logging_configuration = {
    include_execution_data = true
    level                  = "ALL"
    log_destination        = "arn:aws:logs:eu-west-1:123456789012:log-group:/aws/vendedlogs/states/jb-step_function:*"
  }
  service_integrations = {
    batch_Sync = {
      events = ["event1", "event2"]
    }
    dynamodb = {
      dynamodb = ["arn:aws:dynamodb:eu-west-1:052212379155:table/Test"]
    }
  }
  attach_policy_json                = false
  policy_json                       = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "xray:GetSamplingStatisticSummaries"
            ],
            "Resource": ["*"]
        }
    ]
}
EOF
  cloudwatch_log_group_name         = "my-log-group"
  use_existing_cloudwatch_log_group = false
  tags                              = local.tags
}
