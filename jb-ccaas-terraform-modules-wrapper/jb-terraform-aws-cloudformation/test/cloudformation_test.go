// Package test provides testing utilities for Cloudformation module
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
    template_body               = string(must(os.ReadFile("cloudformation_test.yaml")))
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-cloudformation-%s-%s-%s", prefixCompany, lob, application, env),
		"template_body":               template_body,
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

func TestCloudformationCreation(t *testing.T) {
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
		"id",
		"outputs",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestCloudformationConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Cloudformation"

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
    // Test: Stack ID should be present and valid
	name := fmt.Sprintf("%s-cloudformation-%s-%s-%s", prefixCompany, lob, application, env)

	stackID, ok := detailsMap["id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s stack ID should be a string", resourceName))
	assert.Contains(t, stackID, name, fmt.Sprintf("%s stack ID should contain stack name", resourceName))

	// Test: Outputs should be a map
	outputsRaw, ok := detailsMap["outputs"].(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s outputs should be a map", resourceName))

	// Test: TemplateValidation output should be correct
	validationMsg, ok := outputsRaw["TemplateValidation"].(string)
	assert.True(t, ok, fmt.Sprintf("%s TemplateValidation output should be a string", resourceName))
	assert.Contains(t, validationMsg, "Template is valid", fmt.Sprintf("%s TemplateValidation output should indicate success", resourceName))

}
