# Aurora PostgreSQL Module Tests

This directory contains automated tests for the Aurora PostgreSQL Terraform module using [Terratest](https://terratest.gruntwork.io/).

## Test Structure

### Test Files
- `global_rds_test.go` - Main test file containing all Aurora PostgreSQL tests

### Test Functions

#### TestAuroraPostgreSQLCreation
Validates the basic creation of an Aurora PostgreSQL cluster by:
- Verifying core outputs exist (cluster ARN, ID, endpoints)
- Checking cluster configuration parameters
- Ensuring security group and subnet group creation

#### TestAuroraPostgreSQLConfiguration  
Tests advanced configuration features including:
- Engine version validation
- Instance configuration
- Enhanced monitoring setup
- Parameter group configuration

#### TestAuroraPostgreSQLGlobalCluster
Tests global cluster functionality by:
- Validating global cluster creation
- Checking engine and version configuration
- Verifying global cluster endpoints

## Running Tests

### Prerequisites
- Go 1.19 or later
- AWS credentials configured
- Terraform installed

### Execute Tests
```bash
# Run all tests
go test -v -timeout 30m

# Run specific test
go test -v -run TestAuroraPostgreSQLCreation -timeout 30m

# Run tests in parallel
go test -v -parallel 3 -timeout 45m
```

### Test Configuration
The tests use the following default configuration:
- Region: `us-east-1`
- Environment: `sandbox`
- Application: `aurora`
- LOB: `test`

## Test Outputs Validated

### Core Aurora PostgreSQL Outputs
- `cluster_arn` - Amazon Resource Name of the cluster
- `cluster_id` - RDS Cluster Identifier
- `cluster_endpoint` - Writer endpoint for the cluster
- `cluster_reader_endpoint` - Read-only endpoint
- `cluster_database_name` - Database name
- `cluster_port` - Database port
- `security_group_id` - Security group ID
- `db_subnet_group_name` - DB subnet group name

### Configuration Outputs
- `cluster_engine_version_actual` - Running engine version
- `cluster_instances` - Map of cluster instances
- `enhanced_monitoring_iam_role_arn` - Enhanced monitoring role
- `db_cluster_parameter_group_id` - Cluster parameter group
- `db_parameter_group_id` - DB parameter group

### Global Cluster Outputs
- `global_cluster_id` - Global cluster identifier
- `global_cluster_arn` - Global cluster ARN
- `global_cluster_engine` - Database engine
- `global_cluster_engine_version` - Engine version
- `global_cluster_endpoint` - Global cluster endpoint

## Notes
- Tests create real AWS resources and incur costs
- Resources are automatically cleaned up after test completion
- Use appropriate AWS credentials with sufficient permissions
- Tests run with error-level logging to reduce noise
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | = 6.0.0 |

## Providers

No providers.

## Modules

No modules.

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END_TF_DOCS -->
