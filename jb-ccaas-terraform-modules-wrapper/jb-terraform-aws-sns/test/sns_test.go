// Package test provides testing utilities for sns module
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
	use_name_prefix             = false
	fifo_topic                  = false
	content_based_deduplication = false
	create_topic_policy         = true
	enable_default_topic_policy = true
	topic_policy_statements     = []string{}
	subscriptions               = map[string]interface{}{
		"lambda_sub": map[string]interface{}{
			"protocol": "lambda",
			"endpoint": "arn:aws:lambda:us-west-2:381492173985:function:my-lambda-func",
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
		"name":                        fmt.Sprintf("%s-sns-%s-%s-%s", prefixCompany, lob, application, env),
		"use_name_prefix":             use_name_prefix,
		"fifo_topic":                  fifo_topic,
		"content_based_deduplication": content_based_deduplication,
		"create_topic_policy":         create_topic_policy,
		"enable_default_topic_policy": enable_default_topic_policy,
		"topic_policy_statements":     topic_policy_statements,
		"subscriptions":               subscriptions,
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

func TestSnsCreation(t *testing.T) {
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
		"sns_topic_arn",
		"sns_topic_id",
		"sns_subscriptions",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestSnsConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Sns"

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
	// Parse outputs
	snsTopicArn, ok := detailsMap["sns_topic_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s topic ARN should be a string", resourceName))
	assert.Contains(t, snsTopicArn, "arn:aws:sns:us-west-2:", fmt.Sprintf("%s topic ARN should contain region", resourceName))
	assert.Contains(t, snsTopicArn, fmt.Sprintf("%s-sns-%s-%s-%s", prefixCompany, lob, application, env), fmt.Sprintf("%s topic ARN should contain formatted name", resourceName))

	// Test: Subscription map is present and not empty
	subscriptions, ok := detailsMap["sns_subscriptions"].(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s subscriptions should be a map", resourceName))
	assert.NotEmpty(t, subscriptions, fmt.Sprintf("%s should have at least one subscription", resourceName))

	// Test: Validate lambda subscription fields
	lambdaSubRaw, ok := subscriptions["lambda_sub"].(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s lambda_sub should be a map", resourceName))

	// Check basic fields
	subArn, ok := lambdaSubRaw["arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s subscription ARN should be a string", resourceName))
	assert.Contains(t, subArn, snsTopicArn, fmt.Sprintf("%s subscription ARN should contain topic ARN", resourceName))

	endpoint, ok := lambdaSubRaw["endpoint"].(string)
	assert.True(t, ok, fmt.Sprintf("%s subscription endpoint should be a string", resourceName))
	assert.Contains(t, endpoint, "lambda", fmt.Sprintf("%s endpoint should contain 'lambda'", resourceName))

	protocol, ok := lambdaSubRaw["protocol"].(string)
	assert.True(t, ok, fmt.Sprintf("%s subscription protocol should be a string", resourceName))
	assert.Equal(t, "lambda", protocol, fmt.Sprintf("%s subscription protocol should be 'lambda'", resourceName))
}
