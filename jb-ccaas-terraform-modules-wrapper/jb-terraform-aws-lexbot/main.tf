resource "awscc_lex_bot" "this" {
  count                       = var.create_lexbot ? 1 : 0
  name                        = var.name == "" ? local.lex_name : var.name
  auto_build_bot_locales      = var.auto_build_bot_locales
  role_arn                    = var.role_arn
  data_privacy                = var.data_privacy
  idle_session_ttl_in_seconds = var.idle_session_ttl_in_seconds
  bot_file_s3_location        = var.bot_file_s3_location
  bot_tags                    = local.bot_tags
  description                 = var.description
}

resource "aws_lexv2models_bot_version" "this" {
  count                = var.create_lexbot_version ? 1 : 0
  bot_id               = var.bot_id
  description          = var.version_description
  locale_specification = var.locale_specification
  lifecycle {
    create_before_destroy = true
  }
}

resource "awscc_lex_bot_alias" "this" {
  count                       = var.create_lexbot_alias ? 1 : 0
  bot_alias_name              = var.bot_alias_name
  bot_id                      = var.bot_id
  bot_version                 = var.bot_version
  sentiment_analysis_settings = var.sentiment_analysis_settings
  bot_alias_locale_settings   = var.bot_alias_locale_settings
  conversation_log_settings   = var.conversation_log_settings
}

resource "awscc_lex_resource_policy" "this" {
  count        = var.create_alias_policy ? 1 : 0
  resource_arn = var.bot_alias_arn
  policy       = var.lex_policy
}
