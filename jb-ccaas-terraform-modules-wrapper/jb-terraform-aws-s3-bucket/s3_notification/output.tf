output "s3_bucket_notification_id" {
  description = "ID of S3 bucket"
  value       = try(module.s3_notification.s3_bucket_notification_id, "")
}
