output "state_machine_id" {
  description = "The ARN of the State Machine"
  value       = module.step_function.state_machine_id
}


output "state_machine_arn" {
  description = "The ARN of the State Machine"
  value       = module.step_function.state_machine_arn
}
