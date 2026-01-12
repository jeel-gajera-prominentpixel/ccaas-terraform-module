// Package test provides testing utilities for ssm parameter store module
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
    parameter_write = []map[string]string{
		{
			"name":      "OUTPUT_BUCKET_NAME",
			"value":    "jb-test-cases-usw2-sandbox",
			"type":      "String",
			"overwrite": "true",
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
		"name":                        fmt.Sprintf("%s-parameter-store-%s-%s-%s", prefixCompany, lob, application, env),
		"parameter_write":             parameter_write,
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

func TestParameterCreation(t *testing.T) {
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
		"parameter_names",
		"parameter_values",
		"parameter_map",
		"parameter_arn_map",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestParameterConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Parameter-Store"

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
	// Test: parameter_names should be a non-empty list of strings
	paramNames, ok := detailsMap["parameter_names"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s parameter_names should be a list", resourceName))
	assert.NotEmpty(t, paramNames, fmt.Sprintf("%s parameter_names should not be empty", resourceName))

	// Test: parameter_values should be a list of same length as parameter_names
	paramValues, ok := detailsMap["parameter_values"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s parameter_values should be a list", resourceName))
	assert.Equal(t, len(paramNames), len(paramValues), fmt.Sprintf("%s parameter_values should match parameter_names in length", resourceName))

	// Test: parameter_map should contain keys from parameter_names with matching values
	paramMap, ok := detailsMap["parameter_map"].(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s parameter_map should be a map", resourceName))
	for i, name := range paramNames {
		nameStr := name.(string)
		expectedValue := paramValues[i]
		actualValue, exists := paramMap[nameStr]
		assert.True(t, exists, fmt.Sprintf("%s parameter_map should contain key %s", resourceName, nameStr))
		assert.Equal(t, expectedValue, actualValue, fmt.Sprintf("%s parameter_map value for %s should match parameter_values", resourceName, nameStr))
	}

	// Test: parameter_arn_map should contain valid ARN for each parameter
	paramArnMap, ok := detailsMap["parameter_arn_map"].(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s parameter_arn_map should be a map", resourceName))
	for _, name := range paramNames {
		nameStr := name.(string)
		arnVal, exists := paramArnMap[nameStr]
		assert.True(t, exists, fmt.Sprintf("%s parameter_arn_map should contain key %s", resourceName, nameStr))
		arnStr, isString := arnVal.(string)
		assert.True(t, isString, fmt.Sprintf("%s ARN for %s should be a string", resourceName, nameStr))
		assert.Contains(t, arnStr, ":ssm:", fmt.Sprintf("%s ARN for %s should contain ':ssm:'", resourceName, nameStr))
		assert.Contains(t, arnStr, nameStr, fmt.Sprintf("%s ARN for %s should contain the parameter name", resourceName, nameStr))
	}
}
