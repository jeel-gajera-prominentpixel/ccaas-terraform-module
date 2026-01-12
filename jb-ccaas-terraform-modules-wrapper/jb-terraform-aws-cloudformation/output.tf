output "id" {
  description = "A unique identifier of the stack"
  value       = try(aws_cloudformation_stack.stack.id, null)
}

output "outputs" {
  description = "A map of outputs from the stack"
  value       = try(aws_cloudformation_stack.stack.outputs, null)
}
