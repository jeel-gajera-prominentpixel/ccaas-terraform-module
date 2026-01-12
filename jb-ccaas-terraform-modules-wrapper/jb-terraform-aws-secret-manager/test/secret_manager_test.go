// Package test provides testing utilities for Secret-Manager module
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
	ignore_secret_changes		= true
	create_policy			    = false
	block_public_policy		    = true
	secret_string 		        = "{}"
	recovery_window_in_days     = 0
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-secret-%s-%s-%s", prefixCompany, lob, application, env),
		"ignore_secret_changes":       ignore_secret_changes,
		"create_policy":               create_policy,
		"block_public_policy":         block_public_policy,
		"secret_string":               secret_string,
		"recovery_window_in_days":     recovery_window_in_days,
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

func TestSecretManagerCreation(t *testing.T) {
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
		"secret_arn",
		"secret_id",
		"secret_replica",
		"secret_version_id",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestSecretManagerConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "SecretManager"

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
	// Test: Secret ARN is a valid ARN and contains expected name
	expectedSecretName := fmt.Sprintf("%s-secret-%s-%s-%s", prefixCompany, lob, application, env)
	secretArn, ok := detailsMap["secret_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s secret ARN should be a string", resourceName))
	assert.Contains(t, secretArn, "arn:aws:secretsmanager:us-west-2:", fmt.Sprintf("%s secret ARN should contain region", resourceName))
	assert.Contains(t, secretArn, expectedSecretName, fmt.Sprintf("%s secret ARN should contain secret name", resourceName))

	// Test: Secret ID matches ARN
	secretID, ok := detailsMap["secret_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s secret ID should be a string", resourceName))
	assert.Equal(t, secretArn, secretID, fmt.Sprintf("%s secret ID should match secret ARN", resourceName))

	// Test: Secret version ID is present and non-empty
	secretVersionID, ok := detailsMap["secret_version_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s secret version ID should be a string", resourceName))
	assert.NotEmpty(t, secretVersionID, fmt.Sprintf("%s secret version ID should not be empty", resourceName))
	assert.Contains(t, secretVersionID, "terraform-", fmt.Sprintf("%s secret version ID should be Terraform-generated", resourceName))

	// Test: Secret replica is an empty list
	secretReplica, ok := detailsMap["secret_replica"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s secret replica should be a list", resourceName))
	assert.Len(t, secretReplica, 0, fmt.Sprintf("%s secret replica list should be empty", resourceName))

}
