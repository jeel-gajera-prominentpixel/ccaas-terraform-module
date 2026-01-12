// Package test provides testing utilities for Cloudtrail module
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
	region                        = "us-east-1"
	prefixRegion                  = "use1"
	prefixCompany                 = "jb"
	lob                           = "test"
	application                   = "cases"
	env                           = "sandbox"
    enable_logging			      = true
	enable_log_file_validation    = true
	include_global_service_events = true
	is_multi_region_trail		  = true
	is_organization_trail 		  = false
	s3_bucket_name                = "jb-test-cases-usw2-sandbox"
	s3_key_prefix                 = "sandbox/jb"
	advancedEventSelector         = []map[string]interface{}{
		{
			"name": "S3 external Object Level Logging",
			"field_selector": []map[string]interface{}{
				{
					"field":  "eventCategory",
					"equals": []string{"Data"},
				},
				{
					"field":  "resources.type",
					"equals": []string{"AWS::S3::Object"},
				},
				{
					"field":        "resources.ARN",
					"starts_with":  []string{
						"arn:aws:s3:::" + s3_bucket_name,
					},
				},
			},
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
		"name":                        fmt.Sprintf("%s-cloudtrail-%s-%s-%s", prefixCompany, lob, application, env),
		"enable_logging":              enable_logging,
		"enable_log_file_validation":  enable_log_file_validation,
		"include_global_service_events": include_global_service_events,
		"is_multi_region_trail":       is_multi_region_trail,
		"is_organization_trail":       is_organization_trail,
		"s3_bucket_name":              s3_bucket_name,
		"s3_key_prefix":               s3_key_prefix,
		"advanced_event_selector":     advancedEventSelector,
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

func TestCloudtrailCreation(t *testing.T) {
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
		"cloudtrail_id",
		"cloudtrail_home_region",
		"cloudtrail_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestCloudtrailConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Cloudtrail"

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
    // Expected trail name
	expectedTrailName := fmt.Sprintf("%s-cloudtrail-%s-%s-%s", prefixCompany, lob, application, env)
	
	// Test: CloudTrail ID is present and formatted correctly
	cloudtrailID, ok := detailsMap["cloudtrail_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s ID should be a string", resourceName))
	assert.NotEmpty(t, cloudtrailID, fmt.Sprintf("%s ID should not be empty", resourceName))
	assert.Contains(t, cloudtrailID, expectedTrailName, fmt.Sprintf("%s ID should contain expected trail name", resourceName))
	assert.Contains(t, cloudtrailID, "trail/", fmt.Sprintf("%s ID should contain 'trail/'", resourceName))

	// Test: CloudTrail ARN is present and matches the ID
	cloudtrailArn, ok := detailsMap["cloudtrail_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s ARN should be a string", resourceName))
	assert.NotEmpty(t, cloudtrailArn, fmt.Sprintf("%s ARN should not be empty", resourceName))
	assert.Equal(t, cloudtrailID, cloudtrailArn, fmt.Sprintf("%s ARN should match ID", resourceName))
	assert.Contains(t, cloudtrailArn, "arn:aws:cloudtrail", fmt.Sprintf("%s ARN should contain 'arn:aws:cloudtrail'", resourceName))

	// Test: CloudTrail ARN contains the trail name
	assert.Contains(t, cloudtrailArn, expectedTrailName, fmt.Sprintf("%s ARN should contain trail name", resourceName))

	// Test: CloudTrail home region is present and matches expected value
	homeRegion, ok := detailsMap["cloudtrail_home_region"].(string)
	assert.True(t, ok, fmt.Sprintf("%s home region should be a string", resourceName))
	assert.NotEmpty(t, homeRegion, fmt.Sprintf("%s home region should not be empty", resourceName))
	assert.Contains(t, cloudtrailArn, homeRegion, fmt.Sprintf("%s ARN should contain home region", resourceName))
    
}
