// Package test provides testing utilities for Security-Group module
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
	region                                = "us-east-1"
	prefixRegion                          = "use1"
	prefixCompany                         = "jb"
	lob                                   = "test"
	application                           = "cases"
	env                                   = "sandbox"
	description 			              = "Test security group"
	vpc_id					              = "vpc-081ccbb8554d14f3e"
	egress_rules                          = []string{"all-all"}
	egress_cidr_blocks                    = []string{"0.0.0.0/0"}
	egress_ipv6_cidr_blocks               = []string{"::/0"}
	ingress_rules                         = []string{"all-all"}
	ingress_cidr_blocks                   = []string{"0.0.0.0/0"}
	ingress_ipv6_cidr_blocks              = []string{"::/0"}
	ingress_with_source_security_group_id = []string{}
	tags = map[string]string{
		"Name": fmt.Sprintf("%s-security-group-%s-%s-%s", prefixCompany, lob, application, env),
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                           application,
		"prefix_company":                        prefixCompany,
		"prefix_region":                         prefixRegion,
		"lob":                                   lob,
		"env":                                   env,
		"name":                                  fmt.Sprintf("%s-security-group-%s-%s-%s", prefixCompany, lob, application, env),
		"description":                           description,
		"vpc_id":                                vpc_id,
		"egress_rules":                          egress_rules,
		"egress_cidr_blocks":		             egress_cidr_blocks,
		"egress_ipv6_cidr_blocks":               egress_ipv6_cidr_blocks,
		"ingress_rules":                         ingress_rules,
		"ingress_cidr_blocks":                   ingress_cidr_blocks,
		"ingress_ipv6_cidr_blocks":              ingress_ipv6_cidr_blocks,
		"ingress_with_source_security_group_id": ingress_with_source_security_group_id,
		"tags":                                  tags,
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

func TestSecurityGroupCreation(t *testing.T) {
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
		"security_group_arn",
		"security_group_id",
		"security_group_vpc_id",
		"security_group_owner_id",
		"security_group_name",
		"security_group_description",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

func TestSecurityGroupConfiguration(t *testing.T) {
	t.Parallel()
	resourceName := "Security-Group"

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
	// Expected values
	expectedSGName := fmt.Sprintf("%s-security-group-%s-%s-%s", prefixCompany, lob, application, env)

	// Test: Security Group ID is a valid string
	sgID, ok := detailsMap["security_group_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s ID should be a string", resourceName))
	assert.NotEmpty(t, sgID, fmt.Sprintf("%s ID should not be empty", resourceName))

	// Test: Security Group ARN is a valid string
	sgArn, ok := detailsMap["security_group_arn"].(string)
	assert.True(t, ok, fmt.Sprintf("%s ARN should be a string", resourceName))
	assert.Contains(t, sgArn, "arn:aws:ec2:us-west-2:", fmt.Sprintf("%s ARN should contain correct region", resourceName))
	assert.Contains(t, sgArn, sgID, fmt.Sprintf("%s ARN should contain security group ID", resourceName))

	// Test: Security Group Name matches expected
	sgName, ok := detailsMap["security_group_name"].(string)
	assert.True(t, ok, fmt.Sprintf("%s name should be a string", resourceName))
	assert.Equal(t, expectedSGName, sgName, fmt.Sprintf("%s name should match expected", resourceName))

	// Test: Security Group Description is present and a string
	sgDescription, ok := detailsMap["security_group_description"].(string)
	assert.True(t, ok, fmt.Sprintf("%s description should be a string", resourceName))
	assert.NotEmpty(t, sgDescription, fmt.Sprintf("%s description should not be empty", resourceName))

	// Test: Security Group VPC ID is present and a string
	sgVpcID, ok := detailsMap["security_group_vpc_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s VPC ID should be a string", resourceName))
	assert.NotEmpty(t, sgVpcID, fmt.Sprintf("%s VPC ID should not be empty", resourceName))

	// Test: Security Group Owner ID is present
	sgOwnerID, ok := detailsMap["security_group_owner_id"].(string)
	assert.True(t, ok, fmt.Sprintf("%s owner ID should be a string", resourceName))
	assert.NotEmpty(t, sgOwnerID, fmt.Sprintf("%s owner ID should not be empty", resourceName))
}
