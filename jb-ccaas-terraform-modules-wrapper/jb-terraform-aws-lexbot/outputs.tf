output "bot_id" {
  description = "ID of the created Lex bot"
  value       = try(awscc_lex_bot.this[0].id, "")
}

output "bot_version" {
  description = "Version of the created Lex bot"
  value       = try(aws_lexv2models_bot_version.this[0].bot_version, "")
}

output "bot_alias_id" {
  description = "ARN of the created Lex bot alias"
  value       = try(awscc_lex_bot_alias.this[0].id, "")
}

output "bot_alias_name" {
  description = "Name of the created Lex bot alias"
  value       = try(awscc_lex_bot_alias.this[0].bot_alias_name, "")
}

output "bot_alias_arn" {
  description = "ARN of the created Lex bot alias"
  value       = try(awscc_lex_bot_alias.this[0].arn, "")
}

output "bot_name" {
  description = "Name of the created Lex bot"
  value       = try(awscc_lex_bot.this[0].name, "")
}
