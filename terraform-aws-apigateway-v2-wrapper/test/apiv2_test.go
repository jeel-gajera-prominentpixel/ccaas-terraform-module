// Package test provides testing utilities for Api Gateway V2 module
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
	create_api_domain_name = false
	integrations           = map[string]map[string]interface{}{
		"ANY /{proxy+}": {
			"lambda_arn":             "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function",
			"payload_format_version": "2.0",
			"timeout_milliseconds":   30000,
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":            application,
		"prefix_company":         prefixCompany,
		"prefix_region":          prefixRegion,
		"lob":                    lob,
		"env":                    env,
		"name":                   fmt.Sprintf("%s-apiv2-%s-%s-%s", prefixCompany, lob, application, env),
		"create_api_domain_name": create_api_domain_name,
		"integrations":           integrations,
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

func TestApiV2Creation(t *testing.T) {
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
		"apigatewayv2_api_arn",
		"default_apigatewayv2_stage_invoke_url",
		"apigatewayv2_api_api_endpoint",
		"apigatewayv2_api_execution_arn",
		"apigatewayv2_domain_name_target_domain_name",
		"default_apigatewayv2_stage_domain_name",
		"apigatewayv2_domain_name_arn",
		"apigatewayv2_domain_name_id",
		"apigatewayv2_api_id",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestApiV2Configuration(t *testing.T) {
	t.Parallel()
	resourceName := "API Gateway V2"
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

	// Test case 1: Verify API Gateway V2 exists
	assert.NotNil(t, detailsMap, fmt.Sprintf("%s details should not be nil", resourceName))

	// Test case 2: Extract API Gateway ID for dynamic validation
	apiId := detailsMap["apigatewayv2_api_id"].(string)
	assert.NotEmpty(t, apiId, "API Gateway ID should not be empty")

	// Test case 3: Verify API Gateway endpoint
	expectedEndpoint := fmt.Sprintf("https://%s.execute-api.us-west-2.amazonaws.com", apiId)
	assert.Equal(t, expectedEndpoint, detailsMap["apigatewayv2_api_api_endpoint"], "API Gateway endpoint should match expected value")

	// Test case 4: Verify API Gateway ARN
	expectedArn := fmt.Sprintf("arn:aws:apigateway:us-west-2::/apis/%s", apiId)
	assert.Equal(t, expectedArn, detailsMap["apigatewayv2_api_arn"], "API Gateway ARN should match expected value")

	// Test case 5: Verify API Gateway Execution ARN
	expectedExecArn := fmt.Sprintf("arn:aws:execute-api:us-west-2:381492173985:%s", apiId)
	assert.Equal(t, expectedExecArn, detailsMap["apigatewayv2_api_execution_arn"], "API Gateway execution ARN should match expected value")

	// Test case 6: Verify default stage domain name
	expectedDomainName := fmt.Sprintf("%s.execute-api.us-west-2.amazonaws.com", apiId)
	assert.Equal(t, expectedDomainName, detailsMap["default_apigatewayv2_stage_domain_name"], "Default stage domain name should match expected value")

	// Test case 7: Verify default stage invoke URL
	expectedInvokeUrl := fmt.Sprintf("https://%s.execute-api.us-west-2.amazonaws.com/", apiId)
	assert.Equal(t, expectedInvokeUrl, detailsMap["default_apigatewayv2_stage_invoke_url"], "Default stage invoke URL should match expected value")

	// Test case 8: Verify custom domain name settings (not configured)
	assert.Empty(t, detailsMap["apigatewayv2_domain_name_arn"], "Custom domain name ARN should be empty")
	assert.Empty(t, detailsMap["apigatewayv2_domain_name_id"], "Custom domain name ID should be empty")
	assert.Empty(t, detailsMap["apigatewayv2_domain_name_target_domain_name"], "Custom domain name target should be empty")

	defer terraform.Destroy(t, terraformOptions)
}
