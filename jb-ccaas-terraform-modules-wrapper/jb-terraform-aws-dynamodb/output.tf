output "dynamodb_table_arn" {
  description = "ARN of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_arn
}

output "dynamodb_table_id" {
  description = "ID of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_id
}

output "dynamodb_table_stream_arn" {
  description = "The ARN of the Table Stream. Only available when var.stream_enabled is true"
  value       = module.dynamodb-table.dynamodb_table_stream_arn
}

output "dynamodb_table_stream_label" {
  description = "A timestamp, in ISO 8601 format of the Table Stream. Only available when var.stream_enabled is true"
  value       = module.dynamodb-table.dynamodb_table_stream_label
}

output "dynamodb_table_attributes" {
  description = "Dynamo DB table attributes"
  value       = module.dynamodb-table.dynamodb_table_attributes
}

output "dynamodb_table_hash_key" {
  description = "Attribute to use as the hash (partition) key."
  value       = module.dynamodb-table.dynamodb_table_hash_key
}

output "dynamodb_table_name" {
  description = "Name of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_name
}

output "dynamodb_table_billing_mode" {
  description = "Billing mode of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_billing_mode
}

output "dynamodb_table_deletion_protection_enabled" {
  description = "Enables deletion protection for table"
  value       = module.dynamodb-table.dynamodb_table_deletion_protection_enabled
}

output "dynamodb_table_global_secondary_index" {
  description = "Global secondary index of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_global_secondary_index
}

output "dynamodb_table_local_secondary_index" {
  description = "Local secondary index of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_local_secondary_index
}

output "dynamodb_table_range_key" {
  description = "Range key of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_range_key
}

output "dynamodb_table_replica" {
  description = "Replica of the DynamoDB table"
  value       = module.dynamodb-table.dynamodb_table_replica
}
