// Package test provides testing utilities for IAM module
package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variables to store dynamic IDs
var ()

func getPolicyVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-iam-policy-%s-%s-%s", prefixCompany, lob, application, env),
        "policy":                      getPolicyJSON(),
	}
	return vars
}

func TestIamPolicyCreation(t *testing.T) {
	// t.Parallel()

	terraformPolicyOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../iam-policy",
		Vars:         getPolicyVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformPolicyOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformPolicyOptions)

	// Verify additional resources
	outputs := []string{
		"id",
		"arn",
		"description",
		"name",
		"path",
		"policy",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformPolicyOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestIamPolicyConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "IamPolicy"

	terraformPolicyOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../iam-policy",
		Vars:         getPolicyVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformPolicyOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformPolicyOptions)

	outputs := terraform.OutputAll(t, terraformPolicyOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Test Cases here
	// Test: IAM policy name is correct
	expectedPolicyName := fmt.Sprintf("%s-iam-policy-%s-%s-%s", prefixCompany, lob, application, env)
	actualPolicyName, ok := detailsMap["name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s policy name should be a string", resourceName))
	assert.Equal(t, expectedPolicyName, actualPolicyName, fmt.Sprintf("%s policy name should match expected value", resourceName))

	// Test: IAM policy ARN is correctly formatted
	policyArn, ok := detailsMap["arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s policy ARN should be a string", resourceName))
	assert.Contains(t, policyArn, "arn:aws:iam::", fmt.Sprintf("%s policy ARN should contain AWS IAM prefix", resourceName))
	assert.Contains(t, policyArn, expectedPolicyName, fmt.Sprintf("%s policy ARN should include policy name", resourceName))

	// Test: IAM policy path is "/"
	path, ok := detailsMap["path"].(string)
	assert.True(t, ok, fmt.Sprintf("%s path should be a string", resourceName))
	assert.Equal(t, "/", path, fmt.Sprintf("%s policy path should be root", resourceName))

	// Test: IAM policy JSON is valid and includes correct actions/resources
	rawPolicy, ok := detailsMap["policy"].(string)
	assert.True(t, ok, fmt.Sprintf("%s policy should be a string", resourceName))

	var policyDoc map[string]interface{}
	err = json.Unmarshal([]byte(rawPolicy), &policyDoc)
	assert.NoError(t, err, fmt.Sprintf("%s policy should be valid JSON", resourceName))

	// Test: policy contains a "Statement" array
	statements, ok := policyDoc["Statement"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s policy should contain 'Statement' array", resourceName))
	assert.NotEmpty(t, statements, fmt.Sprintf("%s policy 'Statement' should not be empty", resourceName))

	// Test: at least one statement has "Action" and "Resource"
	found := false
	for _, stmtRaw := range statements {
		stmt, ok := stmtRaw.(map[string]interface{})
		if !ok {
			continue
		}

		_, hasAction := stmt["Action"]
		_, hasResource := stmt["Resource"]

		if hasAction && hasResource {
			found = true
			break
		}
	}
	assert.True(t, found, fmt.Sprintf("%s policy should contain at least one Statement with both 'Action' and 'Resource'", resourceName))
}
