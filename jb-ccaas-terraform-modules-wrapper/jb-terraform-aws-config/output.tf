output "config_recorder_id" {
  value       = module.config.aws_config_configuration_recorder_id
  description = "The id of the AWS Config Recorder that was created"
}
