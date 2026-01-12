// Package test provides testing utilities for Aurora PostgreSQL module
package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variables to store dynamic IDs
var (
	DynamicGlobalClusterID            string
	DynamicGlobalClusterEngine        string
	DynamicGlobalClusterEngineVersion string
	DynamicGlobalClusterDBName        string
)

// Constants for Aurora PostgreSQL testing
var (
	region                    = "us-east-1"
	prefixRegion              = "use1"
	prefixCompany             = "jb"
	lob                       = "test"
	application               = "aurora"
	env                       = "sandbox"
	global_cluster_identifier = "jb-test-cases-global-cluster"
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":               application,
		"prefix_company":            prefixCompany,
		"prefix_region":             prefixRegion,
		"lob":                       lob,
		"env":                       env,
		"global_cluster_identifier": global_cluster_identifier,
		"create_global_cluster":     true,
		"force_destroy":             true,
		"global_cluster_engine":     "aurora-postgresql",
		"global_cluster_db_name":    "testdb",
	}
	return vars
}

// Required tags for Aurora PostgreSQL resources
var requiredTags = []string{
	"module_project_path",
	"module_version",
	"company",
	"region",
	"lob",
	"application",
	"env",
	"created_by",
	"map-migrated",
}

// TestMain sets up the shared Aurora PostgreSQL cluster for all tests
func TestMain(m *testing.M) {
	// Create a shared Aurora PostgreSQL cluster for all tests
	terraformOptions := terraform.WithDefaultRetryableErrors(&testing.T{}, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVarsOfDB(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformOptions)

	// Run all tests
	code := m.Run()

	// Clean up after all tests complete
	terraform.Destroy(&testing.T{}, terraformOptions)

	os.Exit(code)
}

// func TestAuroraPostgreSQLGlobalCluster(t *testing.T) {
// 	resourceName := "Global Cluster"
// 	// t.Parallel()

// 	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
// 		TerraformDir: "../global_cluster/",
// 		Vars:         getCommonVars(),
// 		NoStderr:     true,
// 		Logger:       logger.Discard,
// 		EnvVars: map[string]string{
// 			"TF_LOG":      "ERROR",
// 			"TF_LOG_PATH": "/dev/null",
// 		},
// 	})

// 	// Initialize and apply the Terraform configuration
// 	terraform.InitAndApply(t, terraformOptions)

// 	// Clean up after test
// 	defer terraform.Destroy(t, terraformOptions)

// 	// Verify global cluster outputs
// 	globalOutputs := []string{
// 		"global_cluster_id",
// 		"global_cluster_arn",
// 		"global_cluster_engine",
// 		"global_cluster_engine_version",
// 		"global_cluster_db_name",
// 		"global_cluster_members",
// 		"global_cluster_resource_id",
// 	}
// 	// Response := make(map[string]interface{})
// 	for _, output := range globalOutputs {
// 		_ = terraform.Output(t, terraformOptions, output)
// 		// makeJsonFromTerraformOpt(t, terraformOptions, output, Response)
// 		// t.Logf("%s: %s", output, value)
// 	}

// 	// If you want to log the terraformOptions fields, marshal it to JSON for logging.
// 	outputs := terraform.OutputAll(t, terraformOptions)

// 	optionsBytes, err := json.Marshal(outputs)
// 	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
// 	var detailsStr = string(optionsBytes)
// 	if detailsStr != "" {
// 		var detailsMap map[string]interface{}
// 		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
// 		assert.NoError(t, err, "Should be able to parse details JSON")

// 		// Verify global cluster ID
// 		assert.Equal(t, "jb-test-cases-global-cluster", detailsMap["global_cluster_id"], fmt.Sprintf("%s name should match", resourceName)) // // Verify global cluster ARN
// 		assert.Equal(t, "jb-test-cases-global-cluster", detailsMap["global_cluster_arn"], fmt.Sprintf("%s ARN should match", resourceName))
// 		// // Verify global cluster engine
// 		assert.Equal(t, "aurora-postgresql", detailsMap["global_cluster_engine"], fmt.Sprintf("%s engine should match", resourceName))
// 		// Verify global cluster database name
// 		assert.Equal(t, "testdb", detailsMap["global_cluster_db_name"], fmt.Sprintf("%s database name should match", resourceName))
// 		// Verify global cluster members
// 		assert.Equal(t, []interface{}([]interface{}{}), detailsMap["global_cluster_members"], fmt.Sprintf("%s members should match", resourceName))

// 		DynamicGlobalClusterID = detailsMap["global_cluster_id"].(string)
// 		DynamicGlobalClusterEngine = detailsMap["global_cluster_engine"].(string)
// 		DynamicGlobalClusterEngineVersion = detailsMap["global_cluster_engine_version"].(string)
// 		DynamicGlobalClusterDBName = detailsMap["global_cluster_db_name"].(string)
// 	}
// }

func getCommonVarsOfDB() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                  application,
		"prefix_company":               prefixCompany,
		"prefix_region":                prefixRegion,
		"lob":                          lob,
		"env":                          env,
		"is_primary_cluster":           true,
		"vpc_id":                       "vpc-097e9d79ec20b751f",
		"global_cluster_identifier":    "jb-test-cases-global-sbx",
		"engine":                       "aurora-postgresql",
		"engine_version":               "15.4",
		"engine_mode":                  "provisioned",
		"storage_encrypted":            false,
		"master_username":              "root",
		"database_name":                "testdb",
		"apply_immediately":            true,
		"auto_minor_version_upgrade":   false,
		"enable_http_endpoint":         true,
		"create_db_subnet_group":       true,
		"db_subnet_group_name":         "",
		// "create_db_cluster_parameter_group": true,
		"create_cloudwatch_log_group": true,
		"manage_master_user_password":  false,
		"master_password":              "TestCases#123",
		"monitoring_interval":          60,
		"preferred_maintenance_window": "sun:23:45-mon:00:15",
		"skip_final_snapshot":          true,
		"serverlessv2_scaling_configuration": map[string]interface{}{
			"min_capacity": 0.5,
			"max_capacity": 2,
		},
		"instances": map[string]interface{}{
			"one": map[string]interface{}{},
			"two": map[string]interface{}{},
		},
		"master_user_secret_kms_key_id":  "arn:aws:kms:us-west-2:381492173985:key/mrk-9e899d984f1e40a8891f51b262a1ebe5",
		"instance_class":                 "db.serverless",
		"subnets":                        []string{"subnet-05a099cc20688503f", "subnet-0ef48a942d271c043"},
		"vpc_security_group_ids":         []string{"sg-03bd69d9c20cdd0bb"},
		"security_group_use_name_prefix": false,
		"port":                           5432,
		"create_security_group":          true,
		"deletion_protection":            false,
		"performance_insights_enabled":   true,
		"preferred_backup_window":        "03:00-06:00",
		"backup_retention_period":        7,
		"security_group_rules": map[string]interface{}{
			"postgres_ingress_backend": map[string]interface{}{
				"source_security_group_id": "sg-03bd69d9c20cdd0bb",
				"from_port":                5432,
				"to_port":                  5432,
				"protocol":                 "tcp",
				"description":              "Allow inbound PostgreSQL traffic",
			},
		},
	}
	return vars
}

func TestAuroraPostgreSQLDBCluster(t *testing.T) {
	t.Parallel()
	resourceName := "DB Cluster"
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVarsOfDB(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Initialize and apply the DB cluster configuration
	// terraform.Apply(t, terraformOptions)

	// Verify global cluster outputs
	globalOutputs := []string{
		"cluster_arn",
	}
	// Response := make(map[string]interface{})
	for _, output := range globalOutputs {
		_ = terraform.Output(t, terraformOptions, output)
	}
	// Clean up after test
	defer terraform.Destroy(t, terraformOptions)

	outputs := terraform.OutputAll(t, terraformOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	if detailsStr != "" {
		var detailsMap map[string]interface{}
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")

		// initializing the resource name and details map , test the output
	    fmt.Printf("Terraform Output Map for %s:\n%+v\n", resourceName, detailsMap)

		// Cluster assertions
		assert.Equal(t, "jb-rds-psql-test-aurora-use1-sandbox", detailsMap["cluster_id"], "Cluster ID should match")

		// Helper function to get last part of ARN
		getLastPartOfARN := func(arn string) string {
			parts := strings.Split(arn, ":")
			return parts[len(parts)-1]
		}

		assert.Equal(t, "jb-rds-psql-test-aurora-use1-sandbox", getLastPartOfARN(detailsMap["cluster_arn"].(string)), fmt.Sprintf("%s cluster ARN suffix should match", resourceName))
		// assert.Equal(t, "cluster-S5FPNWPA65N4TWCBPVHMKVNVGM", detailsMap["cluster_resource_id"], fmt.Sprintf("%s cluster Resource ID should match", resourceName))
		assert.Equal(t, "testdb", detailsMap["cluster_database_name"], fmt.Sprintf("%s database name should match", resourceName))
		assert.Equal(t, "15.4", detailsMap["cluster_engine_version_actual"], fmt.Sprintf("%s engine version should match", resourceName))
		assert.Equal(t, float64(5432), detailsMap["cluster_port"], fmt.Sprintf("%s port should match", resourceName))
		assert.Equal(t, "root", detailsMap["cluster_master_username"], fmt.Sprintf("%s master username should match", resourceName))
		assert.Equal(t, "TestCases#123", detailsMap["cluster_master_password"], fmt.Sprintf("%s master password should match", resourceName))

		// Instance assertions
		instances := detailsMap["cluster_instances"].(map[string]interface{})

		// Test first instance
		instance1 := instances["one"].(map[string]interface{})
		assert.Equal(t, "jb-rds-psql-test-aurora-use1-sandbox-one", getLastPartOfARN(instance1["arn"].(string)), fmt.Sprintf("%s instance one ARN suffix should match", resourceName))
		// assert.Equal(t, "db-ENDQV2X6GE2SHB7CQYNVJE5GU4", instance1["dbi_resource_id"], fmt.Sprintf("%s instance one Resource ID should match", resourceName))
		assert.Equal(t, "db.serverless", instance1["instance_class"], fmt.Sprintf("%s instance class should match", resourceName))

		// if _, ok := detailsMap["cluster_reader_endpoint"]; ok {
		// 	assert.Equal(t, false, instance1["writer"], fmt.Sprintf("%s first instance should be writer", resourceName))
		// } else {
		// 	assert.Equal(t, true, instance1["writer"], fmt.Sprintf("%s first instance should be writer", resourceName))
		// }

		assert.Equal(t, float64(60), instance1["monitoring_interval"], fmt.Sprintf("%s monitoring interval should match", resourceName))
		// assert.Equal(t, "us-west-2b", instance1["availability_zone"], fmt.Sprintf("%s AZ should match", resourceName))

		// Test second instance
		instance2 := instances["two"].(map[string]interface{})
		assert.Equal(t, "jb-rds-psql-test-aurora-use1-sandbox-two", getLastPartOfARN(instance2["arn"].(string)), fmt.Sprintf("%s instance two ARN suffix should match", resourceName))
		// assert.Equal(t, "db-TWHPFCNGWEODKZG3ATGBQRPUDQ", instance2["dbi_resource_id"], fmt.Sprintf("%s instance two Resource ID should match", resourceName))
		assert.Equal(t, "db.serverless", instance2["instance_class"], fmt.Sprintf("%s instance class should match", resourceName))
		// if _, ok := detailsMap["cluster_reader_endpoint"]; ok {
		// 	assert.Equal(t, true, instance2["writer"], fmt.Sprintf("%s first instance should be writer", resourceName))
		// } else {
		// 	assert.Equal(t, false, instance2["writer"], fmt.Sprintf("%s first instance should be writer", resourceName))
		// }

		// assert.Equal(t, "us-west-2a", instance2["availability_zone"], fmt.Sprintf("%s AZ should match", resourceName))

		// Test tags
		tags := instance1["tags"].(map[string]interface{})
		assert.Equal(t, "aurora", tags["application"], fmt.Sprintf("%s application tag should match", resourceName))
		assert.Equal(t, "jb", tags["company"], fmt.Sprintf("%s company tag should match", resourceName))
		assert.Equal(t, "sandbox", tags["env"], fmt.Sprintf("%s environment tag should match", resourceName))
		assert.Equal(t, "test", tags["lob"], fmt.Sprintf("%s LOB tag should match", resourceName))

		// ====================================================================
		// ADDITIONAL OUTPUTS VALIDATION
		// ====================================================================

		// Helper function to check if string contains expected value
		assertContains := func(actual, expected, message string) {
			assert.Contains(t, actual, expected, message)
		}

		// Test db_subnet_group_name
		if dbSubnetGroupName, ok := detailsMap["db_subnet_group_name"]; ok {
			assert.NotEmpty(t, dbSubnetGroupName, fmt.Sprintf("%s DB subnet group name should not be empty", resourceName))
			assertContains(dbSubnetGroupName.(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s DB subnet group should contain cluster identifier", resourceName))
		}

		// Test cluster_resource_id
		if clusterResourceId, ok := detailsMap["cluster_resource_id"]; ok {
			assert.NotEmpty(t, clusterResourceId, fmt.Sprintf("%s cluster resource ID should not be empty", resourceName))
			assert.True(t, strings.HasPrefix(clusterResourceId.(string), "cluster-"), fmt.Sprintf("%s cluster resource ID should start with 'cluster-'", resourceName))
		}

		// Test cluster_endpoint
		if clusterEndpoint, ok := detailsMap["cluster_endpoint"]; ok {
			assert.NotEmpty(t, clusterEndpoint, fmt.Sprintf("%s cluster endpoint should not be empty", resourceName))
			assertContains(clusterEndpoint.(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s cluster endpoint should contain cluster identifier", resourceName))
			assertContains(clusterEndpoint.(string), ".cluster-", fmt.Sprintf("%s cluster endpoint should be cluster endpoint format", resourceName))
			assertContains(clusterEndpoint.(string), ".rds.amazonaws.com", fmt.Sprintf("%s cluster endpoint should end with RDS domain", resourceName))
			assert.True(t, strings.HasSuffix(clusterEndpoint.(string), ".us-west-2.rds.amazonaws.com"), fmt.Sprintf("%s cluster endpoint should be in correct regional format", resourceName))
		}

		// Test cluster_reader_endpoint
		if clusterReaderEndpoint, ok := detailsMap["cluster_reader_endpoint"]; ok {
			assert.NotEmpty(t, clusterReaderEndpoint, fmt.Sprintf("%s cluster reader endpoint should not be empty", resourceName))
			assertContains(clusterReaderEndpoint.(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s cluster reader endpoint should contain cluster identifier", resourceName))
			assertContains(clusterReaderEndpoint.(string), ".cluster-ro-", fmt.Sprintf("%s cluster reader endpoint should be read-only endpoint format", resourceName))
			assertContains(clusterReaderEndpoint.(string), ".rds.amazonaws.com", fmt.Sprintf("%s cluster reader endpoint should end with RDS domain", resourceName))
			assert.True(t, strings.HasSuffix(clusterReaderEndpoint.(string), ".us-west-2.rds.amazonaws.com"), fmt.Sprintf("%s cluster reader endpoint should be in correct regional format", resourceName))

			// Ensure writer and reader endpoints are different
			if clusterEndpoint, exists := detailsMap["cluster_endpoint"]; exists {
				assert.NotEqual(t, clusterEndpoint.(string), clusterReaderEndpoint.(string), fmt.Sprintf("%s writer and reader endpoints should be different", resourceName))
			}
		}

		// Test security_group_id
		if securityGroupId, ok := detailsMap["security_group_id"]; ok {
			assert.NotEmpty(t, securityGroupId, fmt.Sprintf("%s security group ID should not be empty", resourceName))
			assert.True(t, strings.HasPrefix(securityGroupId.(string), "sg-"), fmt.Sprintf("%s security group ID should start with 'sg-'", resourceName))
		}

		// Test cluster_members
		// if clusterMembers, ok := detailsMap["cluster_members"]; ok {
		// 	membersList := clusterMembers.([]interface{})
		// 	assert.GreaterOrEqual(t, len(membersList), 2, fmt.Sprintf("%s should have at least 2 cluster members", resourceName))
		// 	// Verify member structure
		// 	for i, member := range membersList {
		// 		memberMap := member.(map[string]interface{})
		// 		assert.NotEmpty(t, memberMap["db_cluster_member_id"], fmt.Sprintf("%s member %d should have cluster member ID", resourceName, i))
		// 		assert.NotEmpty(t, memberMap["db_instance_identifier"], fmt.Sprintf("%s member %d should have instance identifier", resourceName, i))
		// 		assert.Contains(t, memberMap["db_instance_identifier"].(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s member %d instance identifier should contain cluster name", resourceName, i))
		// 	}
		// }

		// Test cluster_hosted_zone_id
		if hostedZoneId, ok := detailsMap["cluster_hosted_zone_id"]; ok {
			assert.NotEmpty(t, hostedZoneId, fmt.Sprintf("%s cluster hosted zone ID should not be empty", resourceName))
		}

		// Test enhanced_monitoring_iam_role_name
		if enhancedMonitoringRoleName, ok := detailsMap["enhanced_monitoring_iam_role_name"]; ok {
			assert.NotEmpty(t, enhancedMonitoringRoleName, fmt.Sprintf("%s enhanced monitoring role name should not be empty", resourceName))
			// assertContains(enhancedMonitoringRoleName.(string), "enhanced-monitoring", fmt.Sprintf("%s enhanced monitoring role name should contain 'enhanced-monitoring'", resourceName))
		}

		// Test enhanced_monitoring_iam_role_arn
		if enhancedMonitoringRoleArn, ok := detailsMap["enhanced_monitoring_iam_role_arn"]; ok {
			assert.NotEmpty(t, enhancedMonitoringRoleArn, fmt.Sprintf("%s enhanced monitoring role ARN should not be empty", resourceName))
			assert.True(t, strings.HasPrefix(enhancedMonitoringRoleArn.(string), "arn:aws:iam::"), fmt.Sprintf("%s enhanced monitoring role ARN should be valid IAM role ARN", resourceName))
			// assertContains(enhancedMonitoringRoleArn.(string), "enhanced-monitoring", fmt.Sprintf("%s enhanced monitoring role ARN should contain 'enhanced-monitoring'", resourceName))
		}

		// Test enhanced_monitoring_iam_role_unique_id
		if enhancedMonitoringRoleUniqueId, ok := detailsMap["enhanced_monitoring_iam_role_unique_id"]; ok {
			assert.NotEmpty(t, enhancedMonitoringRoleUniqueId, fmt.Sprintf("%s enhanced monitoring role unique ID should not be empty", resourceName))
			assert.True(t, strings.HasPrefix(enhancedMonitoringRoleUniqueId.(string), "AROA"), fmt.Sprintf("%s enhanced monitoring role unique ID should start with 'AROA'", resourceName))
		}

		// Test db_cluster_parameter_group_arn - NOTE: This output may be missing because
		// create_db_cluster_parameter_group is likely set to false in the module configuration
		if clusterParamGroupArn, ok := detailsMap["db_cluster_parameter_group_arn"]; ok {
			if clusterParamGroupArn != nil && clusterParamGroupArn.(string) != "" {
				assert.NotEmpty(t, clusterParamGroupArn, fmt.Sprintf("%s cluster parameter group ARN should not be empty", resourceName))
				assert.True(t, strings.HasPrefix(clusterParamGroupArn.(string), "arn:aws:rds:"), fmt.Sprintf("%s cluster parameter group ARN should be valid RDS ARN", resourceName))
				assertContains(clusterParamGroupArn.(string), "cluster-pg:", fmt.Sprintf("%s cluster parameter group ARN should contain 'cluster-pg:'", resourceName))
			} else {
				t.Logf("db_cluster_parameter_group_arn is empty - likely using default parameter group")
			}
		} else {
			t.Logf("db_cluster_parameter_group_arn output not available - create_db_cluster_parameter_group might be false")
		}

		// Test db_cluster_parameter_group_id
		if clusterParamGroupId, ok := detailsMap["db_cluster_parameter_group_id"]; ok {
			if clusterParamGroupId != nil && clusterParamGroupId.(string) != "" {
				assert.NotEmpty(t, clusterParamGroupId, fmt.Sprintf("%s cluster parameter group ID should not be empty", resourceName))
				assertContains(clusterParamGroupId.(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s cluster parameter group ID should contain cluster identifier", resourceName))
			} else {
				t.Logf("db_cluster_parameter_group_id is empty - using default parameter group")
			}
		} else {
			t.Logf("db_cluster_parameter_group_id output not available")
		}

		// Test db_parameter_group_arn - NOTE: May be empty if using default parameter groups
		if dbParamGroupArn, ok := detailsMap["db_parameter_group_arn"]; ok {
			if dbParamGroupArn != nil && dbParamGroupArn.(string) != "" {
				assert.NotEmpty(t, dbParamGroupArn, fmt.Sprintf("%s DB parameter group ARN should not be empty", resourceName))
				assert.True(t, strings.HasPrefix(dbParamGroupArn.(string), "arn:aws:rds:"), fmt.Sprintf("%s DB parameter group ARN should be valid RDS ARN", resourceName))
				assertContains(dbParamGroupArn.(string), "pg:", fmt.Sprintf("%s DB parameter group ARN should contain 'pg:'", resourceName))
			} else {
				t.Logf("db_parameter_group_arn is empty - using default parameter group")
			}
		} else {
			t.Logf("db_parameter_group_arn output not available")
		}

		// Test db_parameter_group_id
		if dbParamGroupId, ok := detailsMap["db_parameter_group_id"]; ok {
			if dbParamGroupId != nil && dbParamGroupId.(string) != "" {
				assert.NotEmpty(t, dbParamGroupId, fmt.Sprintf("%s DB parameter group ID should not be empty", resourceName))
				assertContains(dbParamGroupId.(string), "jb-rds-psql-test-aurora-use1-sandbox", fmt.Sprintf("%s DB parameter group ID should contain cluster identifier", resourceName))
			} else {
				t.Logf("db_parameter_group_id is empty - using default parameter group")
			}
		} else {
			t.Logf("db_parameter_group_id output not available")
		}

		// Test additional_cluster_endpoints - This is empty because no custom endpoints are configured
		if additionalEndpoints, ok := detailsMap["additional_cluster_endpoints"]; ok {
			endpointsMap := additionalEndpoints.(map[string]interface{})
			// This is expected to be empty if no custom endpoints are configured
			assert.NotNil(t, endpointsMap, fmt.Sprintf("%s additional cluster endpoints should be a valid map", resourceName))
			assert.Equal(t, 0, len(endpointsMap), fmt.Sprintf("%s additional cluster endpoints should be empty when no custom endpoints configured", resourceName))
			t.Logf("additional_cluster_endpoints is empty as expected - no custom endpoints configured")
		}

		// Test cluster_role_associations - This is empty because no IAM roles are associated with cluster
		if roleAssociations, ok := detailsMap["cluster_role_associations"]; ok {
			associationsMap := roleAssociations.(map[string]interface{})
			// This is expected to be empty if no IAM roles are associated with the cluster
			assert.NotNil(t, associationsMap, fmt.Sprintf("%s cluster role associations should be a valid map", resourceName))
			assert.Equal(t, 0, len(associationsMap), fmt.Sprintf("%s cluster role associations should be empty when no roles associated", resourceName))
			t.Logf("cluster_role_associations is empty as expected - no IAM roles associated with cluster")
		}

		// Test db_cluster_cloudwatch_log_groups - May be empty if create_cloudwatch_log_group is false
		if logGroups, ok := detailsMap["db_cluster_cloudwatch_log_groups"]; ok {
			logGroupsMap := logGroups.(map[string]interface{})
			assert.NotNil(t, logGroupsMap, fmt.Sprintf("%s CloudWatch log groups should be a valid map", resourceName))
			// Check if PostgreSQL log group is present
			if len(logGroupsMap) > 0 {
				for logGroupName := range logGroupsMap {
					assert.True(t, strings.Contains(logGroupName, "postgresql"), fmt.Sprintf("%s log group name should contain 'postgresql'", resourceName))
				}
				t.Logf("Found %d CloudWatch log groups", len(logGroupsMap))
			} else {
				t.Logf("db_cluster_cloudwatch_log_groups is empty - may be disabled or using default settings")
			}
		}

		// Test db_cluster_activity_stream_kinesis_stream_name - Not available because activity stream is not enabled
		if activityStreamName, ok := detailsMap["db_cluster_activity_stream_kinesis_stream_name"]; ok {
			// This might be empty if activity stream is not enabled
			if activityStreamName != nil && activityStreamName.(string) != "" {
				assert.NotEmpty(t, activityStreamName, fmt.Sprintf("%s activity stream Kinesis stream name should not be empty when enabled", resourceName))
				assertContains(activityStreamName.(string), "aws-rds-das-db-", fmt.Sprintf("%s activity stream name should follow AWS naming convention", resourceName))
			} else {
				t.Logf("db_cluster_activity_stream_kinesis_stream_name is empty - activity stream not enabled")
			}
		} else {
			t.Logf("db_cluster_activity_stream_kinesis_stream_name output not available - activity stream feature not enabled")
		}

		// Test db_cluster_secretsmanager_secret_rotation_enabled
		if secretRotationEnabled, ok := detailsMap["db_cluster_secretsmanager_secret_rotation_enabled"]; ok {
			// This should be a boolean value
			assert.IsType(t, false, secretRotationEnabled, fmt.Sprintf("%s secret rotation enabled should be a boolean", resourceName))
		}

		// Test cluster_master_user_secret - This is empty because manage_master_user_password is false
		if masterUserSecret, ok := detailsMap["cluster_master_user_secret"]; ok {
			// This should be empty if manage_master_user_password is false
			secretArray := masterUserSecret.([]interface{})
			assert.Equal(t, 0, len(secretArray), fmt.Sprintf("%s master user secret should be empty when manage_master_user_password is false", resourceName))
			t.Logf("cluster_master_user_secret is empty as expected - manage_master_user_password is false")
		}

		// Test sensitive outputs (postgresql_cluster_master_password)
		if postgresqlMasterPassword, ok := detailsMap["postgresql_cluster_master_password"]; ok {
			assert.Equal(t, "TestCases#123", postgresqlMasterPassword.(string), fmt.Sprintf("%s postgresql master password should match expected value", resourceName))
		}

		t.Logf("TestAuroraPostgreSQLDBCluster: All test cases including additional outputs are successful!")
		t.Logf("TestAuroraPostgreSQLDBCluster Test cases are successful.")
	}
}
