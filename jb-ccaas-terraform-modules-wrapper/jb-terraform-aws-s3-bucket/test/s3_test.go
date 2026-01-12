// Package test provides testing utilities for S3 Bucket module
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
	region               = "us-east-1"
	prefixRegion         = "use1"
	prefixCompany        = "jb"
	lob                  = "test"
	application          = "cases"
	env                  = "sandbox"
	force_destroy        = true
	acl                  = "null"
	lambda_trigger       = true
	create_bucket        = true
	object_ownership     = "BucketOwnerEnforced"
	versioning           = false
	lambda_notifications = map[string]map[string]interface{}{
		"lambda1": {
			"function_arn":  "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function",
			"function_name": "jb-ccaas-test-cases-lambda-function",
			"events":        []string{"s3:ObjectCreated:*"},
			"filter_prefix": "CTR/",
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":          application,
		"prefix_company":       prefixCompany,
		"prefix_region":        prefixRegion,
		"lob":                  lob,
		"env":                  env,
		"name":                 fmt.Sprintf("%s-s3-%s-%s-%s", prefixCompany, lob, application, env),
		"force_destroy":        force_destroy,
		"acl":                  acl,
		"lambda_trigger":       lambda_trigger,
		"create_bucket":        create_bucket,
		"object_ownership":     object_ownership,
		"versioning":           versioning,
		"lambda_notifications": lambda_notifications,
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

func TestS3Creation(t *testing.T) {
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
		"s3_bucket_arn",
		"s3_bucket_id",
		"s3_bucket_bucket_domain_name",
		"s3_bucket_bucket_regional_domain_name",
		"s3_bucket_force_destroy",
		"s3_bucket_bucket_prefix",
		// "s3_bucket_acl",
		"s3_bucket_versioning",
		"s3_bucket_hosted_zone_id",
		"s3_bucket_lifecycle_configuration_rules",
		"s3_bucket_policy",
		"s3_bucket_region",
		"s3_bucket_website_endpoint",
		"s3_bucket_website_domain",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestS3Configuration(t *testing.T) {
	t.Parallel()
	resourceName := "S3 Bucket"
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

	// Test case 1: Verify S3 Bucket exists
	assert.NotNil(t, detailsMap, fmt.Sprintf("%s details should not be nil", resourceName))

	// Test case 2: Verify S3 Bucket ARN
	assert.Equal(t, "arn:aws:s3:::jb-s3-test-cases-sandbox", detailsMap["s3_bucket_arn"], "S3 bucket ARN should match expected value")

	// Test case 3: Verify S3 Bucket Domain Name
	assert.Equal(t, "jb-s3-test-cases-sandbox.s3.amazonaws.com", detailsMap["s3_bucket_bucket_domain_name"], "S3 bucket domain name should match expected value")

	// Test case 4: Verify S3 Bucket Regional Domain Name
	assert.Equal(t, "jb-s3-test-cases-sandbox.s3.us-west-2.amazonaws.com", detailsMap["s3_bucket_bucket_regional_domain_name"], "S3 bucket regional domain name should match expected value")

	// Test case 5: Verify Force Destroy Setting
	assert.Equal(t, true, detailsMap["s3_bucket_force_destroy"], "Force destroy should be enabled")

	// Test case 6: Verify S3 Bucket ID
	assert.Equal(t, "jb-s3-test-cases-sandbox", detailsMap["s3_bucket_id"], "S3 bucket ID should match expected value")

	// Test case 7: Verify S3 Bucket Region
	assert.Equal(t, "us-west-2", detailsMap["s3_bucket_region"], "S3 bucket region should match expected value")

	// Test case 8: Verify Hosted Zone ID
	assert.Equal(t, "Z3BJ6K6RIION7M", detailsMap["s3_bucket_hosted_zone_id"], "S3 bucket hosted zone ID should match expected value")

	// Test case 9: Verify Versioning Configuration
	versioningConfig := detailsMap["s3_bucket_versioning"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, false, versioningConfig["enabled"], "Versioning should be disabled")
	assert.Equal(t, false, versioningConfig["mfa_delete"], "MFA delete should be disabled")

	// Test case 10: Verify Bucket Policy
	expectedPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []interface{}{
			map[string]interface{}{
				"Sid":       "AllowSSLRequestsOnly",
				"Effect":    "Deny",
				"Principal": "*",
				"Action":    "s3:*",
				"Resource": []interface{}{
					"arn:aws:s3:::jb-s3-test-cases-sandbox/*",
					"arn:aws:s3:::jb-s3-test-cases-sandbox",
				},
				"Condition": map[string]interface{}{
					"Bool": map[string]interface{}{
						"aws:SecureTransport": "false",
					},
				},
			},
		},
	}

	var actualPolicy map[string]interface{}
	err = json.Unmarshal([]byte(detailsMap["s3_bucket_policy"].(string)), &actualPolicy)
	assert.NoError(t, err, "Should be able to parse bucket policy JSON")
	assert.Equal(t, expectedPolicy, actualPolicy, "Bucket policy should match expected configuration")

	// Test case 11: Verify Website Configuration
	assert.Empty(t, detailsMap["s3_bucket_website_domain"], "Website domain should be empty")
	assert.Empty(t, detailsMap["s3_bucket_website_endpoint"], "Website endpoint should be empty")

	defer terraform.Destroy(t, terraformOptions)
}
