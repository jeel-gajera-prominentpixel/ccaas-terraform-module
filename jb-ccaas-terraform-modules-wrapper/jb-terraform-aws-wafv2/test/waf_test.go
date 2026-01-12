// Package test provides testing utilities for Wafv2 module
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
	region                        = "us-east-1"
	prefixRegion                  = "use1"
	prefixCompany                 = "jb"
	lob                           = "test"
	application                   = "cases"
	env                           = "sandbox"
	description                   = "Test WAF ACL for CloudFront to restrict IP addresses"
	scope                         = "REGIONAL"
	default_action                = "block"
	enabled_logging_configuration = true

	waf_rules = []map[string]interface{}{
		{
			"name":                       fmt.Sprintf("allow-specified-ips-%s-%s", lob, application),
			"priority":                   1,
			"sampled_requests_enabled":   true,
			"cloudwatch_metrics_enabled": true,
			"action":                     "allow",
			"ip_set_name":                fmt.Sprintf("allowed-ips-%s-%s", lob, application),
			"metric_name":                fmt.Sprintf("allow-specified-ips-%s-%s", lob, application),
		},
		{
			"name":                       fmt.Sprintf("allow-specified-ipv6-%s-%s", lob, application),
			"priority":                   2,
			"sampled_requests_enabled":   true,
			"cloudwatch_metrics_enabled": true,
			"action":                     "allow",
			"ip_set_name":                fmt.Sprintf("allowed-ips-ipv6-%s-%s", lob, application),
			"metric_name":                fmt.Sprintf("allow-specified-ipv6-%s-%s", lob, application),
        },
    }

	waf_ip_sets = []map[string]interface{}{
		{
			"name":               fmt.Sprintf("allowed-ips-%s-%s", lob, application),
			"ip_address_version": "IPV4",
			"addresses_list":     []string{"64.25.25.249/32"}, //  []string defined in Go
		},
		{
			"name":               fmt.Sprintf("allowed-ips-ipv6-%s-%s", lob, application),
			"ip_address_version": "IPV6",
			"addresses_list":     []string{"2a09:bac0:1000:24::/64"}, //  []string defined in Go
		},
    }

    visibility_config = map[string]interface{}{
		"cloudwatch_metrics_enabled": true,
		"metric_name":                fmt.Sprintf("cloudfront-waf-acl-%s-%s", lob, application),
		"sampled_requests_enabled":   true,
    }

    redacted_fields = map[string]interface{}{
		"query_string": map[string]bool{
			"query_string": true,
		},
		"method": map[string]bool{
			"method": true,
		},
		"uri_path": map[string]bool{
			"uri_path": true,
		},
    }

	tags = map[string]string{
		"Name":        "waf-web-acl-test-cases-sandbox",
		"Environment": "sandbox",
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                   application,
		"prefix_company":                prefixCompany,
		"prefix_region":                 prefixRegion,
		"lob":                           lob,
		"env":                           env,
		"name":                          fmt.Sprintf("%s-waf-%s-%s-%s", prefixCompany, lob, application, env),
		"description":                   description,
		"scope":                         scope,
		"default_action":                default_action,
		"enabled_logging_configuration": enabled_logging_configuration,
		"waf_rules":                     waf_rules,
		"waf_ip_sets":                   waf_ip_sets,
		"visibility_config":             visibility_config,
		"redacted_fields":               redacted_fields,
		"tags":                          tags,
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

func TestWafCreation(t *testing.T) {
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
		"aws_wafv2_arn",
		"aws_wafv2_capacity",
		"aws_wafv2_id",
		"aws_wafv2_tags_all",
		"aws_wafv2_web_acl_logging_configuration_id",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestWafConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Wafv2"

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
	// Test: Web ACL ARN should be a valid WAF ARN
	assert := assert.New(t)

	wafArn, ok := detailsMap["aws_wafv2_arn"].(string)
	assert.True(ok, fmt.Sprintf("%s ARN should be a string", resourceName))
	assert.Contains(wafArn, "arn:aws:wafv2:us-west-2:", fmt.Sprintf("%s ARN should contain region", resourceName))
	assert.Contains(wafArn, "/webacl/", fmt.Sprintf("%s ARN should contain /webacl/", resourceName))

	// Test: WAFv2 ID should be non-empty string
	wafID, ok := detailsMap["aws_wafv2_id"].(string)
	assert.True(ok, fmt.Sprintf("%s ID should be a string", resourceName))
	assert.NotEmpty(wafID, fmt.Sprintf("%s ID should not be empty", resourceName))

	// Test: WAFv2 Capacity should be a positive number
	capacity, ok := detailsMap["aws_wafv2_capacity"].(float64)
	assert.True(ok, fmt.Sprintf("%s capacity should be a number", resourceName))
	assert.Greater(capacity, float64(0), fmt.Sprintf("%s capacity should be greater than 0", resourceName))

	// Test: Tags should exist and include required keys
	tagsMap, ok := detailsMap["aws_wafv2_tags_all"].(map[string]interface{})
	assert.True(ok, fmt.Sprintf("%s tags should be a map", resourceName))
	assert.NotEmpty(tagsMap, fmt.Sprintf("%s tags should not be empty", resourceName))

	requiredTags := []string{"Name", "Environment", "application", "lob"}
	for _, tagKey := range requiredTags {
		_, exists := tagsMap[tagKey]
		assert.True(exists, fmt.Sprintf("%s tag should contain key: %s", resourceName, tagKey))
	}

	// Test: Logging Configuration ARN exists and is well formed
	loggingConfigIDs, ok := detailsMap["aws_wafv2_web_acl_logging_configuration_id"].([]interface{})
	assert.True(ok, fmt.Sprintf("%s logging config should be a list", resourceName))
	assert.NotEmpty(loggingConfigIDs, fmt.Sprintf("%s logging config should not be empty", resourceName))

	for _, logID := range loggingConfigIDs {
		logIDStr, ok := logID.(string)
		assert.True(ok, fmt.Sprintf("%s log ID should be a string", resourceName))
		assert.Contains(logIDStr, "/webacl/", fmt.Sprintf("%s log ID should contain /webacl/", resourceName))
	}
}
