output "sns_topic_arn" {
  description = "The ARN of the SNS topic, as a more obvious property (clone of id)"
  value       = module.sns.topic_arn
}

output "sns_topic_id" {
  description = "The ARN of the SNS topic"
  value       = module.sns.topic_id
}

output "sns_subscriptions" {
  description = "Map of subscriptions created and their attributes"
  value       = module.sns.subscriptions
}
