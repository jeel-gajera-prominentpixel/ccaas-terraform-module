// Package test provides testing utilities for Api Gateway V1 module
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
	region                       = "us-east-1"
	prefixRegion                 = "use1"
	prefixCompany                = "jb"
	lob                          = "test"
	application                  = "cases"
	env                          = "sandbox"
	stage_name                   = "testcases"
	root_integration_http_method = "POST"
	root_integration_type        = "AWS_PROXY"
	root_lambda_arn              = "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function"
	description                  = "description for test cases"
	types                        = "REGIONAL"
	authorization                = "NONE"
	resource_root_path           = "ANY"
	enable_waf_association       = true
	web_acl_arn                  = "arn:aws:wafv2:us-west-2:381492173985:regional/webacl/jb-waf-testcases-sbx/29312adf-d64d-412f-a5e2-dc6c2568680b"
	resource_paths               = map[string]map[string]interface{}{
		"{proxy+}": {
			"lambda_arn":              "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function",
			"http_method":             "ANY",
			"integration_http_method": "POST",
			"type":                    "AWS_PROXY",
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                  application,
		"prefix_company":               prefixCompany,
		"prefix_region":                prefixRegion,
		"lob":                          lob,
		"env":                          env,
		"name":                         fmt.Sprintf("%s-apiv1-%s-%s-%s", prefixCompany, lob, application, env),
		"stage_name":                   stage_name,
		"root_integration_http_method": root_integration_http_method,
		"root_integration_type":        root_integration_type,
		"root_lambda_arn":              root_lambda_arn,
		"description":                  description,
		"types":                        types,
		"authorization":                authorization,
		"resource_root_path":           resource_root_path,
		"enable_waf_association":       enable_waf_association,
		"web_acl_arn":                  web_acl_arn,
		"resource_paths":               resource_paths,
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

func TestApiV1Creation(t *testing.T) {
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
		"apigatewayv1_api_id",
		"apigatewayv1_api_arn",
		"apigatewayv1_api_execution_arn",
		"apigatewayv1_key_source",
		"apigatewayv1_binary_media_types",
		"apigatewayv1_description",
		"apigatewayv1_endpoint_configuration",
		"apigatewayv1_minimum_compression_size",
		"apigatewayv1_policy",
		"apigatewayv1_root_resource_id",
		"apigatewayv1_stage_arn",
		"apigatewayv1_stage_id",
		"apigatewayv1_stage_invoke_url",
		"apigatewayv1_stage_execution_arn",
		"apigatewayv1_stage_web_acl_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestApiV1Configuration(t *testing.T) {
	t.Parallel()
	resourceName := "API Gateway V1"
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

	// Test case 1: Verify API Gateway V1 exists
	assert.NotNil(t, detailsMap, fmt.Sprintf("%s details should not be nil", resourceName))

	// Test case 2: Extract API Gateway ID for dynamic validation
	apiId := detailsMap["apigatewayv1_api_id"].(string)
	assert.NotEmpty(t, apiId, "API Gateway ID should not be empty")

	// Test case 3: Verify API Gateway ARN
	expectedArn := fmt.Sprintf("arn:aws:apigateway:us-west-2::/restapis/%s", apiId)
	assert.Equal(t, expectedArn, detailsMap["apigatewayv1_api_arn"], "API Gateway ARN should match expected value")

	// Test case 4: Verify API Gateway Execution ARN
	expectedExecArn := fmt.Sprintf("arn:aws:execute-api:us-west-2:381492173985:%s", apiId)
	assert.Equal(t, expectedExecArn, detailsMap["apigatewayv1_api_execution_arn"], "API Gateway execution ARN should match expected value")

	// Test case 5: Verify API Description
	assert.Equal(t, "description for test cases", detailsMap["apigatewayv1_description"], "API description should match expected value")

	// Test case 6: Verify Endpoint Configuration
	endpointConfig := detailsMap["apigatewayv1_endpoint_configuration"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, []interface{}{"REGIONAL"}, endpointConfig["types"], "Endpoint configuration type should be REGIONAL")
	assert.Empty(t, endpointConfig["vpc_endpoint_ids"], "VPC endpoint IDs should be empty")

	// Test case 7: Verify API Key Source
	assert.Equal(t, "HEADER", detailsMap["apigatewayv1_key_source"], "API key source should be HEADER")

	// Test case 8: Verify Binary Media Types
	assert.Empty(t, detailsMap["apigatewayv1_binary_media_types"], "Binary media types should be empty")

	// Test case 9: Verify Root Resource ID
	assert.NotEmpty(t, detailsMap["apigatewayv1_root_resource_id"], "Root resource ID should not be empty")

	// Test case 10: Verify Stage Configuration
	expectedStageArn := fmt.Sprintf("arn:aws:apigateway:us-west-2::/restapis/%s/stages/testcases", apiId)
	assert.Equal(t, expectedStageArn, detailsMap["apigatewayv1_stage_arn"], "Stage ARN should match expected value")

	expectedStageExecArn := fmt.Sprintf("arn:aws:execute-api:us-west-2:381492173985:%s/testcases", apiId)
	assert.Equal(t, expectedStageExecArn, detailsMap["apigatewayv1_stage_execution_arn"], "Stage execution ARN should match expected value")

	expectedStageId := fmt.Sprintf("ags-%s-testcases", apiId)
	assert.Equal(t, expectedStageId, detailsMap["apigatewayv1_stage_id"], "Stage ID should match expected value")

	expectedInvokeUrl := fmt.Sprintf("https://%s.execute-api.us-west-2.amazonaws.com/testcases", apiId)
	assert.Equal(t, expectedInvokeUrl, detailsMap["apigatewayv1_stage_invoke_url"], "Stage invoke URL should match expected value")

	// Test case 11: Verify Optional Configurations are Empty
	assert.Empty(t, detailsMap["apigatewayv1_minimum_compression_size"], "Minimum compression size should be empty")
	assert.Empty(t, detailsMap["apigatewayv1_policy"], "API policy should be empty")
	assert.Empty(t, detailsMap["apigatewayv1_stage_web_acl_arn"], "Web ACL ARN should be empty")

	defer terraform.Destroy(t, terraformOptions)
}
