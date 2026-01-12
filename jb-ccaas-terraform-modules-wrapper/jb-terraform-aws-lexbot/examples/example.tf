module "lex_bot" {
  source         = "../"
  prefix_company = "jb"
  lob            = "itsd"
  prefix_region  = "usw2"
  application    = "recordings"
  env            = "sandbox"
  name           = "customer-service-bot"
  description    = "Basic customer service bot"
  role_arn       = "arn:aws:iam::123456789012:role/lex-bot-role"
  bot_file_s3_location = {
    s3_bucket     = "my-lex-bot-bucket"
    s3_object_key = "bot-definition.json"
  }
  locale_specification = {
    "en_US" = {
      source_bot_version = "DRAFT"
    }
  }
  bot_alias_locale_settings = [
    {
      locale_id = "en_US"
      bot_alias_locale_setting = {
        enabled = true
      }
    }
  ]
  conversation_log_settings = {
    text_log_settings = [
      {
        enabled = true
        destination = {
          cloudwatch = {
            cloudwatch_log_group_arn = "arn:aws:logs:us-east-1:123456789012:log-group:/aws/lex/customer-service-bot"
            log_prefix               = "/aws/lex/"
          }
        }
      }
    ]
  }
  lex_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowLexAccess"
        Effect = "Allow"
        Principal = {
          Service = "lexv2.amazonaws.com"
        }
        Action = [
          "lex:*"
        ]
        Resource = "*"
      }
    ]
  })
}
