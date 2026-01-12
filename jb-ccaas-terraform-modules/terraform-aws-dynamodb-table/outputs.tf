output "dynamodb_table_arn" {
  description = "ARN of the DynamoDB table"
  value       = try(aws_dynamodb_table.this[0].arn, aws_dynamodb_table.autoscaled[0].arn, aws_dynamodb_table.autoscaled_gsi_ignore[0].arn, "")
}

output "dynamodb_table_id" {
  description = "ID of the DynamoDB table"
  value       = try(aws_dynamodb_table.this[0].id, aws_dynamodb_table.autoscaled[0].id, aws_dynamodb_table.autoscaled_gsi_ignore[0].id, "")
}

output "dynamodb_table_stream_arn" {
  description = "The ARN of the Table Stream. Only available when var.stream_enabled is true"
  value       = var.stream_enabled ? try(aws_dynamodb_table.this[0].stream_arn, aws_dynamodb_table.autoscaled[0].stream_arn, aws_dynamodb_table.autoscaled_gsi_ignore[0].stream_arn, "") : null
}

output "dynamodb_table_stream_label" {
  description = "A timestamp, in ISO 8601 format of the Table Stream. Only available when var.stream_enabled is true"
  value       = var.stream_enabled ? try(aws_dynamodb_table.this[0].stream_label, aws_dynamodb_table.autoscaled[0].stream_label, aws_dynamodb_table.autoscaled_gsi_ignore[0].stream_label, "") : null
}

output "dynamodb_table_attributes" {
  description = "Dynamodb table attributes"
  value       = aws_dynamodb_table.this[0].attribute
}

output "dynamodb_table_hash_key" {
  description = "Attribute to use as the hash (partition) key"
  value       = aws_dynamodb_table.this[0].hash_key
}

output "dynamodb_table_name" {
  description = "Name of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].name
}

output "dynamodb_table_billing_mode" {
  description = "Billing mode of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].billing_mode
}

output "dynamodb_table_deletion_protection_enabled" {
  description = "Enables deletion protection for table"
  value       = aws_dynamodb_table.this[0].deletion_protection_enabled
}

output "dynamodb_table_global_secondary_index" {
  description = "Global secondary index of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].global_secondary_index
}

output "dynamodb_table_local_secondary_index" {
  description = "Local secondary index of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].local_secondary_index
}

output "dynamodb_table_range_key" {
  description = "Range key of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].range_key
}

output "dynamodb_table_replica" {
  description = "Replica of the DynamoDB table"
  value       = aws_dynamodb_table.this[0].replica
}