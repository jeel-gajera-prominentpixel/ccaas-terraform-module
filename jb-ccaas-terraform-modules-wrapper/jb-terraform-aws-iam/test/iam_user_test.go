// Package test provides testing utilities for IAM module
package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	// "os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variables to store dynamic IDs
var ()

// Constants for testing
var (
	create_user			          = true
	create_iam_user_login_profile = false
	policy_arns				      = []string{
		"arn:aws:iam::381492173985:policy/jb-kinesis-test-cases-sandbox-s3",
	}
)

func getUserVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                 application,
		"prefix_company":              prefixCompany,
		"prefix_region":               prefixRegion,
		"lob":                         lob,
		"env":                         env,
		"name":                        fmt.Sprintf("%s-iam-%s-%s-%s", prefixCompany, lob, application, env),
		"create_user":                 create_user,
		"create_iam_user_login_profile": create_iam_user_login_profile,
		"policy_arns":                 policy_arns,
	}
	return vars
}

func TestIamUserCreation(t *testing.T) {
	// t.Parallel()

	terraformUserOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../Iam-user",
		Vars:         getUserVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformUserOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformUserOptions)

	// Verify additional resources
	outputs := []string{
		"iam_user_name",
		"iam_user_arn",
		"iam_user_unique_id",
		"iam_user_login_profile_key_fingerprint",
		"iam_user_login_profile_encrypted_password",
		"iam_user_login_profile_password",
		"iam_access_key_id",
		"iam_access_key_secret",
		"iam_access_key_key_fingerprint",
		"iam_access_key_encrypted_secret",
		"iam_access_key_ses_smtp_password_v4",
		"iam_access_key_encrypted_ses_smtp_password_v4",
		"iam_access_key_status",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformUserOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestIamUserConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "IamUser"

	terraformUserOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../Iam-user",
		Vars:         getUserVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformUserOptions)

	// Clean up after all tests complete
	defer terraform.Destroy(&testing.T{}, terraformUserOptions)

	outputs := terraform.OutputAll(t, terraformUserOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Test Cases here
	// Test: IAM user name matches expected format
	expectedUserName := fmt.Sprintf("%s-iam-%s-%s-%s", prefixCompany, lob, application, env)
	actualUserName, ok := detailsMap["iam_user_name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s user name should be a string", resourceName))
	assert.Equal(t, expectedUserName, actualUserName, fmt.Sprintf("%s user name should match expected value", resourceName))

	// Test: IAM user ARN is present and formatted correctly
	userArn, ok := detailsMap["iam_user_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s user ARN should be a string", resourceName))
	assert.Contains(t, userArn, "arn:aws:iam::", fmt.Sprintf("%s user ARN should contain expected AWS prefix", resourceName))
	assert.Contains(t, userArn, actualUserName, fmt.Sprintf("%s user ARN should contain user name", resourceName))

	// Test: IAM access key ID is present
	accessKeyID, ok := detailsMap["iam_access_key_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s access key ID should be a string", resourceName))
	assert.NotEmpty(t, accessKeyID, fmt.Sprintf("%s access key ID should not be empty", resourceName))

	// Test: IAM access key secret is present
	accessKeySecret, ok := detailsMap["iam_access_key_secret"].(string)
	assert.True(t, ok, fmt.Sprintf("%s access key secret should be a string", resourceName))
	assert.NotEmpty(t, accessKeySecret, fmt.Sprintf("%s access key secret should not be empty", resourceName))

	// Test: IAM access key status is "Active"
	accessKeyStatus, ok := detailsMap["iam_access_key_status"].(string)
	assert.True(t, ok, fmt.Sprintf("%s access key status should be a string", resourceName))
	assert.Equal(t, "Active", accessKeyStatus, fmt.Sprintf("%s access key status should be Active", resourceName))

	// Test: IAM SES SMTP password is present and valid
	smtpPassword, ok := detailsMap["iam_access_key_ses_smtp_password_v4"].(string)
	assert.True(t, ok, fmt.Sprintf("%s SES SMTP password should be a string", resourceName))
	assert.NotEmpty(t, smtpPassword, fmt.Sprintf("%s SES SMTP password should not be empty", resourceName))

	// Test: IAM user unique ID is present
	userUID, ok := detailsMap["iam_user_unique_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s user unique ID should be a string", resourceName))
	assert.NotEmpty(t, userUID, fmt.Sprintf("%s user unique ID should not be empty", resourceName))

	// initializing the resource name and details map , test the output
	// fmt.Printf("Terraform Output Map for %s:\n%+v\n", resourceName, detailsMap)
}
