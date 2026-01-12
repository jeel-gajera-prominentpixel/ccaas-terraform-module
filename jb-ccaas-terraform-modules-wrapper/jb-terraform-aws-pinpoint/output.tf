output "application_id" {
  value       = module.pinpoint.application_id
  description = "The Application ID of the Pinpoint App."
}

output "application_arn" {
  value       = module.pinpoint.application_arn
  description = "Amazon Resource Name (ARN) of the PinPoint Application."
}
