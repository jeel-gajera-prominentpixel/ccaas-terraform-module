// Package test provides testing utilities for Eventbridge module
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
var (
// inputMap map[string]string
)

// Constants for Amazon Connect testing
var (
	region                        = "us-east-1"
	prefixRegion                  = "use1"
	prefixCompany                 = "jb"
	lob                           = "test"
	application                   = "cases"
	env                           = "sandbox"
	create_bus                    = false
	name						  = "default"
	create_schemas_discoverer     = false
	attach_tracing_policy         = false
	attach_kinesis_policy         = false
	kinesis_target_arns           = []string{}
	attach_sfn_policy             = false
	sfn_target_arns               = []string{}
	attach_sqs_policy             = false
	sqs_target_arns               = []string{}
	attach_cloudwatch_policy      = false
	cloudwatch_target_arns        = []string{}
	append_rule_postfix           = true
	attach_ecs_policy             = false
	ecs_target_arns               = []string{}
	create_api_destinations       = false
	attach_api_destination_policy = false
	create_connections            = false
	attach_policy_json            = false
	attach_policy_jsons           = false
	policy_jsons                  = []string{}
	number_of_policy_jsons        = 0
	attach_policies               = false
	policies                      = []string{}
	number_of_policies            = 0
	attach_policy_statements      = false
	policy_statements             = []string{}
	role_name                     = "jb_eventbridge_test_role"
	rules = map[string]interface{}{
		"testcase_cron_1": map[string]interface{}{
			"description":         "Trigger for a Lambda function at every 30 minutes",
			"schedule_expression": "cron(*/30 * * * ? *)",
		},
	}

	targets = map[string]interface{}{
		"testcase_cron_1": []map[string]interface{}{
			{
				"name": "cron-for-testcase-agent",
				"arn":  "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function",
			},
		},
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                   application,
		"prefix_company":                prefixCompany,
		"prefix_region":                 prefixRegion,
		"lob":                           lob,
		"env":                           env,
		"name":                          name,
		"create_bus":                    create_bus,
		"create_schemas_discoverer":     create_schemas_discoverer,
		"attach_tracing_policy":         attach_tracing_policy,
		"attach_kinesis_policy":         attach_kinesis_policy,
		"kinesis_target_arns":           kinesis_target_arns,
		"attach_sfn_policy":             attach_sfn_policy,
		"sfn_target_arns":               sfn_target_arns,
		"attach_sqs_policy":             attach_sqs_policy,
		"sqs_target_arns":               sqs_target_arns,
		"attach_cloudwatch_policy":      attach_cloudwatch_policy,
		"cloudwatch_target_arns":        cloudwatch_target_arns,
		"append_rule_postfix":           append_rule_postfix,
		"attach_ecs_policy":             attach_ecs_policy,
		"ecs_target_arns":               ecs_target_arns,
		"create_api_destinations":       create_api_destinations,
		"attach_api_destination_policy": attach_api_destination_policy,
		"create_connections":            create_connections,
		"rules":                         rules,
		"targets":                       targets,
		"attach_policy_json":            attach_policy_json,
		"attach_policy_jsons":           attach_policy_jsons,
		"policy_jsons":                  policy_jsons,
		"number_of_policy_jsons":        number_of_policy_jsons,
		"attach_policies":               attach_policies,
		"policies":                      policies,
		"number_of_policies":            number_of_policies,
		"attach_policy_statements":      attach_policy_statements,
		"policy_statements":             policy_statements,
		"role_name":                     role_name,
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

func TestEventbridgeCreation(t *testing.T) {
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
		"eventbridge_rule_arns",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestEventbridgeConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Eventbridge"

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
	// Test: EventBridge rule ARN is present and formatted correctly
	eventbridgeArnsRaw, ok := detailsMap["eventbridge_rule_arns"]
	assert.True(t, ok, fmt.Sprintf("%s eventbridge_rule_arns output should be present", resourceName))

	eventbridgeArns, ok := eventbridgeArnsRaw.(map[string]interface{})
	assert.True(t, ok, fmt.Sprintf("%s eventbridge_rule_arns should be a map", resourceName))

	// Extract actual ARN
	ruleArnRaw, ok := eventbridgeArns["testcase_cron_1"]
	assert.True(t, ok, fmt.Sprintf("%s testcase_cron_1 rule ARN should be present", resourceName))

	ruleArn, ok := ruleArnRaw.(string)
	assert.True(t, ok, fmt.Sprintf("%s testcase_cron_1 rule ARN should be a string", resourceName))

	// Validate ARN contents
	assert.Contains(t, ruleArn, "arn:aws:events:us-west-2:", fmt.Sprintf("%s rule ARN should contain region", resourceName))

}
