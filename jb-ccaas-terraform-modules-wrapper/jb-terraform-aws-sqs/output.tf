output "default_sqs_queue_id" {
  description = "The URL for the created Amazon SQS queue"
  value       = module.sqs.queue_id
}

output "default_sqs_queue_arn" {
  description = "The ARN of the SQS queue"
  value       = module.sqs.queue_arn
}

output "default_sqs_dlq_id" {
  description = "The URL for the created Amazon SQS queue"
  value       = module.sqs.dead_letter_queue_id
}

output "default_sqs_dlq_arn" {
  description = "The ARN of the SQS queue"
  value       = module.sqs.dead_letter_queue_arn
}

output "fifo_sqs_queue_id" {
  description = "The URL for the created Amazon SQS queue"
  value       = module.sqs.queue_id
}

output "fifo_sqs_queue_arn" {
  description = "The ARN of the SQS queue"
  value       = module.sqs.queue_arn
}
