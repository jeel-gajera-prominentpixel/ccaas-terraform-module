// Package test provides testing utilities for Lambda module
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
	region                 = "us-east-1"
	prefixRegion           = "use1"
	prefixCompany          = "jb"
	lob                    = "test"
	application            = "cases"
	env                    = "sandbox"
	local_existing_package = "./test/node.zip"
	handler                = "index.handler"
	runtime                = "nodejs18.x"
	create_role            = true
	environment_variables  = map[string]string{
		"ENV_VAR_1": "value1",
		"ENV_VAR_2": "value2",
	}
	timeout          = 900
	publish          = true
	layers           = []string{"arn:aws:lambda:us-west-2:381492173985:layer:requests:14"} // replace with actual layer ARN
	allowed_triggers = map[string]map[string]string{
		"AllowExecutionFromEventBridgeCron1": {
			"service":    "events",
			"source_arn": "arn:aws:events:us-west-2:381492173985:rule/jb-eventbg-test-cases-rule", // replace with actual ARN
		},
	}
	attach_policy_jsons    = false
	number_of_policy_jsons = 0
	vpc_subnet_ids         = []string{"subnet-0273bc1b1b5073d64", "subnet-0d4c5993bd7365595"}
	vpc_security_group_ids = []string{"sg-0d4cfa9f703de24f8"}
	attach_network_policy  = true
	memory_size            = 1024
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":            application,
		"prefix_company":         prefixCompany,
		"prefix_region":          prefixRegion,
		"lob":                    lob,
		"env":                    env,
		"name":                   fmt.Sprintf("%s-lmda-%s-%s-%s", prefixCompany, lob, application, env),
		"role_name":              fmt.Sprintf("%s-lmda-%s-%s-%s-role", prefixCompany, lob, application, env),
		"local_existing_package": local_existing_package,
		"handler":                handler,
		"runtime":                runtime,
		"create_role":            create_role,
		"environment_variables":  environment_variables,
		"timeout":                timeout,
		"publish":                publish,
		"layers":                 layers,
		"allowed_triggers":       allowed_triggers,
		"attach_policy_jsons":    attach_policy_jsons,
		"number_of_policy_jsons": number_of_policy_jsons,
		"vpc_subnet_ids":         vpc_subnet_ids,
		"vpc_security_group_ids": vpc_security_group_ids,
		"attach_network_policy":  attach_network_policy,
		"memory_size":            memory_size,
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

func TestLambdaCreation(t *testing.T) {
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
		"lambda_function_arn",
		"lambda_function_name",
		"lambda_cloudwatch_log_group_arn",
		"lambda_cloudwatch_log_group_name",
		"local_filename",
		"s3_object",
		"lambda_role_arn",
		"lambda_role_name",
		"lambda_role_unique_id",
		"lambda_event_source_mapping_function_arn",
		"lambda_event_source_mapping_state",
		"lambda_event_source_mapping_state_transition_reason",
		"lambda_event_source_mapping_uuid",
		"lambda_layer_arn",
		"lambda_layer_layer_arn",
		"lambda_layer_created_date",
		"lambda_layer_source_code_size",
		"lambda_layer_version",
		"lambda_function_arn_static",
		"lambda_function_invoke_arn",
		"lambda_function_qualified_arn",
		"lambda_function_qualified_invoke_arn",
		"lambda_function_version",
		"lambda_function_last_modified",
		"lambda_function_source_code_hash",
		"lambda_function_source_code_size",
		"lambda_function_signing_job_arn",
		"lambda_function_signing_profile_version_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestLambdaConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Lambda"
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

	// Test case 1: Verify Lambda function exists
	assert.NotNil(t, detailsMap, fmt.Sprintf("%s details should not be nil", resourceName))

	// Test case 2: Verify Lambda function name
	expectedFunctionName := "jb-lmda-test-cases-sandbox"
	assert.Equal(t, expectedFunctionName, detailsMap["lambda_function_name"], "Lambda function name should match expected value")

	// Test case 3: Verify Lambda function ARN format
	expectedArnPrefix := "arn:aws:lambda:us-west-2:381492173985:function:"
	assert.Contains(t, detailsMap["lambda_function_arn"], expectedArnPrefix, "Lambda ARN should have correct prefix")
	assert.Equal(t, detailsMap["lambda_function_arn"], fmt.Sprintf("%s%s", expectedArnPrefix, expectedFunctionName), "Lambda ARN should match expected format")

	// Test case 4: Verify CloudWatch Log Group
	expectedLogGroupName := fmt.Sprintf("/aws/lambda/%s", expectedFunctionName)
	assert.Equal(t, expectedLogGroupName, detailsMap["lambda_cloudwatch_log_group_name"], "CloudWatch Log Group name should match expected value")

	// Test case 5: Verify function version and state
	assert.NotEmpty(t, detailsMap["lambda_function_version"], "Lambda function version should not be empty")
	version := detailsMap["lambda_function_version"].(string)
	assert.Regexp(t, `^\d+$`, version, "Lambda function version should be a number")

	// Test case 6: Verify IAM Role
	assert.Contains(t, detailsMap["lambda_role_arn"], fmt.Sprintf("arn:aws:iam::381492173985:role/%s", expectedFunctionName), "IAM role ARN should have correct format")
	assert.Equal(t, expectedFunctionName+"-role", detailsMap["lambda_role_name"], "IAM role name should match function name")
	assert.NotEmpty(t, detailsMap["lambda_role_unique_id"], "IAM role should have a unique ID")

	// Test case 7: Verify function configuration
	assert.NotEmpty(t, detailsMap["lambda_function_last_modified"], "Last modified timestamp should not be empty")
	assert.NotEmpty(t, detailsMap["lambda_function_source_code_size"], "Source code size should not be empty")

	// Test case 8: Verify invoke ARN
	expectedInvokeArnPrefix := "arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/"
	assert.Contains(t, detailsMap["lambda_function_invoke_arn"], expectedInvokeArnPrefix, "Lambda invoke ARN should have correct prefix")

	// Test case 9: Verify local deployment package
	assert.Equal(t, "./test/node.zip", detailsMap["local_filename"], "Local deployment package path should match expected value")

	// Test case 10: Verify qualified ARN and invoke ARN
	assert.Contains(t, detailsMap["lambda_function_qualified_arn"],
		fmt.Sprintf("%s:%s", detailsMap["lambda_function_arn"], version),
		"Qualified ARN should include version number")

	assert.Contains(t, detailsMap["lambda_function_qualified_invoke_arn"],
		fmt.Sprintf("%s:%s/invocations", detailsMap["lambda_function_arn"], version),
		"Qualified invoke ARN should include version and invocations path")

	defer terraform.Destroy(t, terraformOptions)
}
