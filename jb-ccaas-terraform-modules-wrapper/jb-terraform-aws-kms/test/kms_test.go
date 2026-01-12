// Package test provides testing utilities for KMS module
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
var ()

// Constants for Amazon Connect testing
var (
	region                = "us-east-1"
	prefixRegion          = "use1"
	prefixCompany         = "jb"
	lob                   = "test"
	application           = "cases"
	env                   = "sandbox"
	description           = "Test cases for KMS module"
	multi_region          = true
	enable_default_policy = true
	key_statements        = []map[string]interface{}{
		{
			"sid":       "Enable Amazon Connect",
			"actions":   []string{"kms:Decrypt*"},
			"resources": []string{"*"},
			"principals": []map[string]interface{}{
				{
					"type":        "Service",
					"identifiers": []string{"connect.amazonaws.com"},
				},
			},
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":           application,
		"prefix_company":        prefixCompany,
		"prefix_region":         prefixRegion,
		"lob":                   lob,
		"env":                   env,
		"name":                  fmt.Sprintf("%s-kms-%s-%s-%s", prefixCompany, lob, application, env),
		"description":           description,
		"multi_region":          multi_region,
		"enable_default_policy": enable_default_policy,
		"key_statements":        key_statements,
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

func TestKMSCreation(t *testing.T) {
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
		"key_arn",
		"key_id",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestKMSConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "KMS"
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

	// Test case 1: Verify KMS key exists
	assert.NotNil(t, detailsMap, fmt.Sprintf("%s details should not be nil", resourceName))

	// Test case 2: Verify KMS key ARN format
	keyArn, ok := detailsMap["key_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s key_arn should be a string", resourceName))
	assert.Contains(t, keyArn, "arn:aws:kms", fmt.Sprintf("%s ARN should have correct prefix", resourceName))
	assert.Contains(t, keyArn, ":key/mrk-", fmt.Sprintf("%s ARN should contain multi-region key prefix 'mrk-'", resourceName))

	// Test case 3: Verify KMS key ID format
	keyId, ok := detailsMap["key_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s key_id should be a string", resourceName))
	assert.Contains(t, keyId, "mrk-", fmt.Sprintf("%s Key ID should be a multi-region key with 'mrk-' prefix", resourceName))
	assert.Greater(t, len(keyId), 30, fmt.Sprintf("%s Key ID should be greater than 30 characters long", resourceName)) // mrk- (4) + UUID (36)

	// Test case 4: Verify key ARN contains correct region and account
	arnParts := strings.Split(keyArn, ":")
	assert.Equal(t, "us-west-2", arnParts[3], fmt.Sprintf("%s key should be in us-west-2 region", resourceName))
	assert.Equal(t, "381492173985", arnParts[4], fmt.Sprintf("%s key should be in correct AWS account", resourceName))

	// Test case 5: Verify key ID in ARN matches the key_id output
	assert.Contains(t, keyArn, keyId, fmt.Sprintf("%s Key ID in ARN should match the key_id output", resourceName))
}
