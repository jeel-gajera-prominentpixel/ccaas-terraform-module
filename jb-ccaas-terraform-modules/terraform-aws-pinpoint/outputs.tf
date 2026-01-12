output "application_id" {
  value       = aws_pinpoint_app.this.application_id
  description = "The Application ID of the Pinpoint App."
}

output "application_arn" {
  value       = aws_pinpoint_app.this.arn
  description = "Amazon Resource Name (ARN) of the PinPoint Application."
}
