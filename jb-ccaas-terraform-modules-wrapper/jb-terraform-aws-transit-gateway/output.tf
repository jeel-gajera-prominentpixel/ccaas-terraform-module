output "ec2_transit_gateway_arn" {
  description = "EC2 Transit Gateway Amazon Resource Name (ARN)"
  value       = module.transit-gateway.ec2_transit_gateway_arn
}

output "ec2_transit_gateway_id" {
  description = "EC2 Transit Gateway identifier"
  value       = module.transit-gateway.ec2_transit_gateway_id
}

output "ec2_transit_gateway_association_default_route_table_id" {
  description = "Identifier of the default association route table"
  value       = module.transit-gateway.ec2_transit_gateway_association_default_route_table_id
}
