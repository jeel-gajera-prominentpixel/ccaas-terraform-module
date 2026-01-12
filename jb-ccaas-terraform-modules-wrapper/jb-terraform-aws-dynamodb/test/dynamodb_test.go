// Package test provides testing utilities for DynamoDB module
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
	hash_key                    = "quickConnectName"
	range_key                   = "lob"
	deletion_protection_enabled = false
	stream_enabled              = true
	stream_view_type            = "NEW_AND_OLD_IMAGES"
	attributes                  = []map[string]string{
		{
			"name": "quickConnectName",
			"type": "S",
		},
		{
			"name": "lob",
			"type": "S",
		},
		{
			"name": "quickConnectType",
			"type": "S",
		},
	}
	global_secondary_indexes = []map[string]string{
		{
			"name":            "lob-index",
			"hash_key":        "lob",
			"range_key":       "quickConnectType",
			"projection_type": "ALL",
		},
	}
	replica_regions = []map[string]any{
		{
			"region_name":                 "us-east-1",
			"deletion_protection_enabled": false,
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-dyndb-%s-%s-%s", prefixCompany, lob, application, env),
		"hash_key":                    hash_key,
		"range_key":                   range_key,
		"deletion_protection_enabled": deletion_protection_enabled,
		"stream_enabled":              stream_enabled,
		"stream_view_type":            stream_view_type,
		"attributes":                  attributes,
		"global_secondary_indexes":    global_secondary_indexes,
		"replica_regions":             replica_regions,
		"tags": map[string]string{
			"company": prefixCompany,
			"lob":     lob,
		},
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

func TestDynamoDBCreation(t *testing.T) {
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
		"dynamodb_table_arn",
		"dynamodb_table_id",
		"dynamodb_table_stream_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestDynamoDBConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "DynamoDB"
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

	// Verify DynamoDB table ARN
	expectedTableArn := "arn:aws:dynamodb:us-west-2:381492173985:table/jb-dyndb-test-cases-sandbox"
	assert.Equal(t, expectedTableArn, detailsMap["dynamodb_table_arn"], fmt.Sprintf("%s table ARN should match expected value", resourceName))

	// Verify DynamoDB table ID and name
	expectedTableId := "jb-dyndb-test-cases-sandbox"
	assert.Equal(t, expectedTableId, detailsMap["dynamodb_table_id"], fmt.Sprintf("%s table ID should match expected value", resourceName))
	assert.Equal(t, expectedTableId, detailsMap["dynamodb_table_name"], fmt.Sprintf("%s table name should match expected value", resourceName))

	// Verify table keys
	assert.Equal(t, "quickConnectName", detailsMap["dynamodb_table_hash_key"], fmt.Sprintf("%s hash key should match expected value", resourceName))
	assert.Equal(t, "lob", detailsMap["dynamodb_table_range_key"], fmt.Sprintf("%s range key should match expected value", resourceName))

	// Verify billing mode and deletion protection
	assert.Equal(t, "PAY_PER_REQUEST", detailsMap["dynamodb_table_billing_mode"], fmt.Sprintf("%s billing mode should be PAY_PER_REQUEST", resourceName))
	assert.Equal(t, false, detailsMap["dynamodb_table_deletion_protection_enabled"], fmt.Sprintf("%s deletion protection should be disabled", resourceName))

	// Verify table attributes
	attributes, ok := detailsMap["dynamodb_table_attributes"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s table attributes should be a slice", resourceName))
	assert.Equal(t, 3, len(attributes), fmt.Sprintf("%s should have 3 attributes", resourceName))
	expectedAttrs := map[string]string{
		"lob":              "S",
		"quickConnectName": "S",
		"quickConnectType": "S",
	}
	for _, attr := range attributes {
		attrMap, ok := attr.(map[string]interface{})
		assert.True(t, ok, fmt.Sprintf("%s attribute should be a map", resourceName))
		name := attrMap["name"].(string)
		attrType := attrMap["type"].(string)
		assert.Equal(t, expectedAttrs[name], attrType, fmt.Sprintf("%s attribute %s should be of type %s", resourceName, name, expectedAttrs[name]))
	}

	// Verify Global Secondary Index
	gsi, ok := detailsMap["dynamodb_table_global_secondary_index"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s GSI should be a slice", resourceName))
	assert.Equal(t, 1, len(gsi), fmt.Sprintf("%s should have 1 GSI", resourceName))
	gsiMap := gsi[0].(map[string]interface{})
	assert.Equal(t, "lob-index", gsiMap["name"], fmt.Sprintf("%s GSI name should match", resourceName))
	assert.Equal(t, "lob", gsiMap["hash_key"], fmt.Sprintf("%s GSI hash key should match", resourceName))
	assert.Equal(t, "quickConnectType", gsiMap["range_key"], fmt.Sprintf("%s GSI range key should match", resourceName))
	assert.Equal(t, "ALL", gsiMap["projection_type"], fmt.Sprintf("%s GSI projection type should be ALL", resourceName))

	// Verify replicas
	replicas, ok := detailsMap["dynamodb_table_replica"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s replicas should be a slice", resourceName))
	assert.Equal(t, 1, len(replicas), fmt.Sprintf("%s should have 1 replica", resourceName))
	replicaMap := replicas[0].(map[string]interface{})
	assert.Equal(t, "us-east-1", replicaMap["region_name"], fmt.Sprintf("%s replica region should match", resourceName))
	assert.Equal(t, false, replicaMap["point_in_time_recovery"], fmt.Sprintf("%s replica point in time recovery should be disabled", resourceName))
	assert.Equal(t, false, replicaMap["propagate_tags"], fmt.Sprintf("%s replica tag propagation should be disabled", resourceName))

	// Verify DynamoDB table stream ARN and label
	streamArn, ok := detailsMap["dynamodb_table_stream_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s table stream ARN should be a string", resourceName))

	// Verify the static part of the stream ARN (everything before the timestamp)
	expectedStreamArnPrefix := "arn:aws:dynamodb:us-west-2:381492173985:table/jb-dyndb-test-cases-sandbox/stream/"
	assert.True(t, len(streamArn) > len(expectedStreamArnPrefix), fmt.Sprintf("%s stream ARN should be longer than the prefix", resourceName))
	assert.Equal(t, expectedStreamArnPrefix, streamArn[:len(expectedStreamArnPrefix)], fmt.Sprintf("%s stream ARN prefix should match expected value", resourceName))

	// Verify the timestamp format in the stream ARN and stream label
	timestamp := streamArn[len(expectedStreamArnPrefix):]
	streamLabel, ok := detailsMap["dynamodb_table_stream_label"].(string)
	assert.True(t, ok, fmt.Sprintf("%s stream label should be a string", resourceName))
	assert.Regexp(t, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}$`, timestamp, fmt.Sprintf("%s stream ARN timestamp should match ISO8601 format", resourceName))
	assert.Equal(t, timestamp, streamLabel, fmt.Sprintf("%s stream label should match stream ARN timestamp", resourceName))
}
