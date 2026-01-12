// Package test provides testing utilities for Cloudfront module
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
	enabled					      = true
    staging                       = false
    http_version                  = "http2"
    retain_on_delete              = false
    wait_for_deployment           = true
    create_origin_access_identity = true
	create_origin_access_control  = true
	create_request_policy         = true

    origin_access_identities = map[string]interface{}{
		"s3_bucket_one": "jb-test-cases-usw2-sbx-oai",
		"api_gw_one": "",
	}

	origin_access_control = map[string]interface{}{
		"custom-s3": map[string]interface{}{
			"description":      "Custom S3 OAC for test",
			"origin_type":      "s3",
			"signing_behavior": "always",
			"signing_protocol": "sigv4",
		},
    }

    default_cache_behavior = map[string]interface{}{
		"target_origin_id": "jb-test-cases-usw2-sandbox",
		"allowed_methods": []string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
		"viewer_protocol_policy": "https-only",
	}

    ordered_cache_behavior = []map[string]interface{}{
		{
			"path_pattern": "/ccp",
			"target_origin_id": "jb-test-cases-usw2-sandbox",
			"allowed_methods": []string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
			"viewer_protocol_policy": "https-only",
	    },
	}

    custom_error_response = []map[string]interface{}{
		{
			"error_code": 404,
			"response_code": 404,
			"response_page_path": "/index.html",
			"error_caching_min_ttl": 10,
		},
	}

	logging_config = map[string]interface{}{
		"bucket": "jb-test-cases-usw2-sandbox.s3.amazonaws.com",
		"prefix": "sandbox/jb",
	}

	origin = map[string]interface{}{
		"jb-test-cases-usw2-sandbox": map[string]interface{}{
			"domain_name": "jb-test-cases-usw2-sandbox.s3.amazonaws.com",
			"s3_origin_config": map[string]interface{}{
				"http_port": 80,
				"https_port": 443,
				"origin_type": "s3",
				"origin_protocol_policy": "match-viewer",
				"origin_ssl_protocols": []string{"TLSv1","TLSv1.1", "TLSv1.2"},
				"origin_access_identity": "s3_bucket_one",
			},
		},
	}

	custom_origin_request_policy = map[string]interface{}{
		"name": "DemoCustomOriginRequestPolicy",
		"comment": "Custom origin request policy configuration",
		"cookies_config": map[string]interface{}{
			"cookie_behavior": "whitelist",
			"cookies": []string{"ad_session"},
		},
		"headers_config": map[string]interface{}{
			"header_behavior": "allExcept",
			"headers": []string{"Host"},
		},
		"query_strings_config": map[string]interface{}{
			"query_string_behavior": "all",
			"query_strings": []string{},
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
		"name":                          fmt.Sprintf("%s-cloudfront-%s-%s-%s", prefixCompany, lob, application, env),
		"enabled":                       enabled,
		"staging":                       staging,
		"http_version":                  http_version,
		"retain_on_delete":              retain_on_delete,
		"wait_for_deployment":           wait_for_deployment,
		"create_origin_access_identity": create_origin_access_identity,
		"origin_access_identities":      origin_access_identities,
		"default_cache_behavior":        default_cache_behavior,
		"ordered_cache_behavior":        ordered_cache_behavior,
		"custom_error_response":         custom_error_response,
		"create_origin_access_control":  create_origin_access_control,
		"origin_access_control":         origin_access_control,
		"logging_config":                logging_config,
		"origin":                        origin,
		"create_request_policy":         create_request_policy,
		"custom_origin_request_policy":  custom_origin_request_policy,
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

func TestCloudfrontCreation(t *testing.T) {
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
		"cloudfront_distribution_id",
		"cloudfront_distribution_arn",
		"cloudfront_distribution_domain_name",
		"cloudfront_origin_access_identity_ids",
		"cloudfront_origin_access_identity_iam_arns",
		"custom_origin_request_policy_id",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestCloudfrontConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Cloudfront"

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

	// Test: CloudFront distribution ID is a non-empty string
	distributionID, ok := detailsMap["cloudfront_distribution_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s distribution ID should be a string", resourceName))
	assert.NotEmpty(t, distributionID, fmt.Sprintf("%s distribution ID should not be empty", resourceName))

	// Test: CloudFront distribution ARN is present and formatted correctly
	distributionArn, ok := detailsMap["cloudfront_distribution_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s distribution ARN should be a string", resourceName))
	assert.Contains(t, distributionArn, "arn:aws:cloudfront::", fmt.Sprintf("%s distribution ARN should be valid", resourceName))
	assert.Contains(t, distributionArn, distributionID, fmt.Sprintf("%s distribution ARN should contain distribution ID", resourceName))

	// Test: CloudFront distribution domain name is correctly formatted
	domainName, ok := detailsMap["cloudfront_distribution_domain_name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s domain name should be a string", resourceName))
	assert.Contains(t, domainName, "cloudfront.net", fmt.Sprintf("%s domain name should be a CloudFront domain", resourceName))

	// Test: Origin Access Identity IAM ARNs are present and correctly formatted
	oaiArns, ok := detailsMap["cloudfront_origin_access_identity_iam_arns"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s OAI IAM ARNs should be a list", resourceName))
	for _, oaiArnRaw := range oaiArns {
		oaiArn := oaiArnRaw.(string)
		assert.Contains(t, oaiArn, "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity", fmt.Sprintf("%s OAI ARN should be valid", resourceName))
	}

	// Test: Origin Access Identity IDs are present and not empty
	oaiIDs, ok := detailsMap["cloudfront_origin_access_identity_ids"].([]interface{})
	assert.True(t, ok, fmt.Sprintf("%s OAI IDs should be a list", resourceName))
	for _, oaiIDRaw := range oaiIDs {
		oaiID := oaiIDRaw.(string)
		assert.NotEmpty(t, oaiID, fmt.Sprintf("%s OAI ID should not be empty", resourceName))
	}

	// Test: Custom origin request policy ID is a valid UUID
	requestPolicyID, ok := detailsMap["custom_origin_request_policy_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s custom origin request policy ID should be a string", resourceName))
	assert.Regexp(t, `[a-f0-9\-]{36}`, requestPolicyID, fmt.Sprintf("%s origin request policy ID should be a valid UUID", resourceName))

}
