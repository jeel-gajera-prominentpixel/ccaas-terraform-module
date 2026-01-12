output "route53_record_name" {
  description = "The name of the record"
  value       = module.route_53_records.route53_record_name
}
