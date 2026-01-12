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

// Constants for testing
var (
	create_role			            = true
	create_custom_role_trust_policy = true
	custom_role_policy_arns			= []string{
		"arn:aws:iam::381492173985:policy/jb-kinesis-test-cases-sandbox-s3",
	}
)

func getRolesVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                     application,
		"prefix_company":                  prefixCompany,
		"prefix_region":                   prefixRegion,
		"lob":                             lob,
		"env":                             env,
		"name":                            fmt.Sprintf("%s-iam-role-%s-%s-%s", prefixCompany, lob, application, env),
		"create_role":                     create_role,
		"create_custom_role_trust_policy": create_custom_role_trust_policy,
		"custom_role_trust_policy":        getTrustPolicyJSON(),
		"custom_role_policy_arns":         custom_role_policy_arns,
	}
	return vars
}

func TestIamRolesCreation(t *testing.T) {
	// t.Parallel()

	terraformRolesOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../iam-assumable-roles",
		Vars:         getRolesVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformRolesOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformRolesOptions)

	// Verify additional resources
	outputs := []string{
		"iam_role_arn",
		"iam_role_name",
		"iam_role_path",
		"iam_role_unique_id",
		"role_requires_mfa",
		"iam_instance_profile_arn",
		"iam_instance_profile_name",
		"iam_instance_profile_id",
		"iam_instance_profile_path",
		"role_sts_externalid",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformRolesOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestIamRolesConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "IamRoles"

	terraformRolesOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../iam-assumable-roles",
		Vars:         getRolesVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformRolesOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformRolesOptions)

	outputs := terraform.OutputAll(t, terraformRolesOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Test Cases here
	// Test: IAM role name matches expected format
	expectedRoleName := fmt.Sprintf("%s-iam-role-%s-%s-%s", prefixCompany, lob, application, env)
	actualRoleName, ok := detailsMap["iam_role_name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s role name should be a string", resourceName))
	assert.Equal(t, expectedRoleName, actualRoleName, fmt.Sprintf("%s role name should match expected value", resourceName))

	// Test: IAM role ARN contains expected values
	roleArn, ok := detailsMap["iam_role_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s role ARN should be a string", resourceName))
	assert.Contains(t, roleArn, "arn:aws:iam::", fmt.Sprintf("%s role ARN should contain AWS IAM prefix", resourceName))
	assert.Contains(t, roleArn, expectedRoleName, fmt.Sprintf("%s role ARN should contain role name", resourceName))

	// Test: IAM role path is "/"
	rolePath, ok := detailsMap["iam_role_path"].(string)
	assert.True(t, ok, fmt.Sprintf("%s role path should be a string", resourceName))
	assert.Equal(t, "/", rolePath, fmt.Sprintf("%s role path should be root", resourceName))

	// Test: IAM role unique ID is non-empty
	roleUID, ok := detailsMap["iam_role_unique_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s unique ID should be a string", resourceName))
	assert.NotEmpty(t, roleUID, fmt.Sprintf("%s unique ID should not be empty", resourceName))

	// Test: role_requires_mfa is true
	requiresMFA, ok := detailsMap["role_requires_mfa"].(bool)
	assert.True(t, ok, fmt.Sprintf("%s requires MFA flag should be a boolean", resourceName))
	assert.True(t, requiresMFA, fmt.Sprintf("%s should require MFA", resourceName))

	// Test: role_sts_externalid is present (can be empty array or string)
	_, okExt := detailsMap["role_sts_externalid"]
	assert.True(t, okExt, fmt.Sprintf("%s should have external ID key", resourceName))

	// Test: instance profile name is correctly formatted (if available)
	if profileNameRaw, ok := detailsMap["iam_instance_profile_name"]; ok {
		profileName, ok := profileNameRaw.(string)
		if ok && profileName != "" {
			assert.Contains(t, profileName, expectedRoleName, fmt.Sprintf("%s instance profile name should include role name", resourceName))
		}
	}
}
