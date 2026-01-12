// Package test provides testing utilities for Firehose module
package test

import (
	"encoding/json"
	"strings"
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
    destination                 = "s3"
	s3_prefix                   = "sandbox/"
	append_delimiter_to_record  = true
	enable_s3_encryption		= false
	s3_bucket_arn               = "arn:aws:s3:::jb-test-cases-usw2-sandbox"
	buffering_size              = 10
	buffering_interval          = 300
	input_source                = "direct-put"
	tags                        = map[string]string{
		"Modifed_date": "2025-07-25",
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-firehose-%s-%s-%s", prefixCompany, lob, application, env),
		"destination":                 destination,
		"s3_prefix":                   s3_prefix,
		"append_delimiter_to_record":  append_delimiter_to_record,
		"enable_s3_encryption":        enable_s3_encryption,
		"s3_bucket_arn":               s3_bucket_arn,
		"buffering_size":              buffering_size,
		"buffering_interval":          buffering_interval,
		"tags":                        tags,
		"input_source":                input_source,
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

func TestFirehoseCreation(t *testing.T) {
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
		"firehose_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestFirehoseConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Firehose"

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

    // Test: Firehose ARN is present and a valid string
	firehoseArn, ok := detailsMap["firehose_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s firehose ARN should be a string", resourceName))
	assert.NotEmpty(t, firehoseArn, fmt.Sprintf("%s firehose ARN should not be empty", resourceName))

	// Test: Firehose ARN should contain the correct AWS service and region
	assert.Contains(t, firehoseArn, "arn:aws:firehose:us-west-2:", fmt.Sprintf("%s firehose ARN should contain correct region", resourceName))

	// Test: Firehose ARN should contain the expected delivery stream name
	expectedFirehoseStreamName := fmt.Sprintf("%s-firehose-%s-%s-%s", prefixCompany, lob, application, env)
	assert.Contains(t, firehoseArn, expectedFirehoseStreamName, fmt.Sprintf("%s firehose ARN should contain expected stream name", resourceName))

	// Test: If ARN has correct format prefix
	assert.True(t, strings.HasPrefix(firehoseArn, "arn:aws:firehose:"), fmt.Sprintf("%s firehose ARN should start with correct ARN prefix", resourceName))

	// Test: If the Firehose ARN contains the deliverystream keyword
	assert.Contains(t, firehoseArn, "deliverystream/", fmt.Sprintf("%s firehose ARN should include 'deliverystream/' keyword", resourceName))

}
