// Package test provides testing utilities for Kinesis module
package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variables to store dynamic IDs
var ()

// Constants for Amazon Connect testing
var (
	region                      = "us-east-1"
	prefixRegion                = "use1"
	prefixCompany               = "jb"
	lob                         = "test"
	application                 = "cases"
	env                         = "sandbox"

	shard_count			        = 1
	retention_period            = 24
	shard_level_metrics         = []string{"IncomingBytes", "OutgoingBytes"}
	enforce_consumer_deletion   = true
	encryption_type			    = "KMS"
	kms_key_id			        = "alias/aws/kinesis"
	create_policy_read_only     = true
	create_policy_write_only    = true
	create_policy_admin         = true
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-kinesis-%s-%s-%s", prefixCompany, lob, application, env),

		"shard_count":                 shard_count,
		"retention_period":            retention_period,
		"shard_level_metrics":         shard_level_metrics,
		"enforce_consumer_deletion":   enforce_consumer_deletion,
		"encryption_type":             encryption_type,
		"kms_key_id":                  kms_key_id,
		"create_policy_read_only":     create_policy_read_only,
		"create_policy_write_only":    create_policy_write_only,
		"create_policy_admin":         create_policy_admin,
	}
	return vars
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Required tags for Amazon Connect resources
var requiredTags = []string{
	"module_project_path",
	// "module_version",
	// "project_path",
	"commit_id",
	"company",
	"region",
	"lob",
	"application",
	"env",
	"created_by",
	"map-migrated",
}

func TestMain(m *testing.M) {
	terraformOptions := terraform.WithDefaultRetryableErrors(&testing.T{}, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
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

func TestKinesisCreation(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Verify additional resources
	outputs := []string{
		"kinesis_stream_name",
		"kinesis_stream_shard_count",
		"kinesis_stream_arn",
		"kinesis_stream_iam_policy_read_only_arn",
		"kinesis_stream_iam_policy_write_only_arn",
		"kinesis_stream_iam_policy_admin_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestKinesisConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Kinesis"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	outputs := terraform.OutputAll(t, terraformOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Test Cases here

	// Test: Kinesis stream name is as expected
	expectedStreamName := fmt.Sprintf("%s-kinesis-%s-%s-%s", prefixCompany, lob, application, env)
	actualStreamName, ok := detailsMap["kinesis_stream_name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s stream name should be a string", resourceName))
	assert.Equal(t, expectedStreamName, actualStreamName, fmt.Sprintf("%s stream name should match expected value", resourceName))

	// Test: Kinesis stream ARN is present and formatted correctly
	streamArn, ok := detailsMap["kinesis_stream_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s stream ARN should be a string", resourceName))
	assert.Contains(t, streamArn, "arn:aws:kinesis:us-west-2:", fmt.Sprintf("%s stream ARN should contain region", resourceName))
	assert.Contains(t, streamArn, expectedStreamName, fmt.Sprintf("%s stream ARN should contain stream name", resourceName))

	// Test: Kinesis stream shard count matches input
    shardCount, ok := detailsMap["kinesis_stream_shard_count"].(float64)
	assert.True(t, ok, fmt.Sprintf("%s shard count should be a number", resourceName))
	assert.Equal(t, float64(shard_count), shardCount, fmt.Sprintf("%s shard count should match input", resourceName))

	// Test: IAM policy ARNs are present and formatted correctly
	adminArn, ok := detailsMap["kinesis_stream_iam_policy_admin_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s admin policy ARN should be a string", resourceName))
	assert.Contains(t, adminArn, fmt.Sprintf("policy/kinesis-stream-%s-admin", expectedStreamName), fmt.Sprintf("%s admin policy ARN should be formatted correctly", resourceName))

	readOnlyArn, ok := detailsMap["kinesis_stream_iam_policy_read_only_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s read-only policy ARN should be a string", resourceName))
	assert.Contains(t, readOnlyArn, fmt.Sprintf("policy/kinesis-stream-%s-read-only", expectedStreamName), fmt.Sprintf("%s read-only policy ARN should be formatted correctly", resourceName))

	writeOnlyArn, ok := detailsMap["kinesis_stream_iam_policy_write_only_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s write-only policy ARN should be a string", resourceName))
	assert.Contains(t, writeOnlyArn, fmt.Sprintf("policy/kinesis-stream-%s-write-only", expectedStreamName), fmt.Sprintf("%s write-only policy ARN should be formatted correctly", resourceName))

}
