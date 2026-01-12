output "firehose_arn" {
  value       = module.firehose.kinesis_firehose_arn
  description = "Configured firehose"
}
