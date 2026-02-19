output "dynamodb_table_arn" {
  description = "ARN of the DynamoDB table."
  value       = one(module.dynamodb_table[*].dynamodb_table_arn)
}

output "dynamodb_table_id" {
  description = "ID of the DynamoDB table."
  value       = one(module.dynamodb_table[*].dynamodb_table_id)
}

output "dynamodb_table_stream_arn" {
  description = "The ARN of the Table Stream."
  value       = one(module.dynamodb_table[*].dynamodb_table_stream_arn)
}

output "dynamodb_table_stream_label" {
  description = "A timestamp, in ISO 8601 format of the Table Stream."
  value       = one(module.dynamodb_table[*].dynamodb_table_stream_label)
}
