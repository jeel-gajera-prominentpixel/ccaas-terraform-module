// Package test provides testing utilities for IAM module
// helpers.go
package test

import (
	"encoding/json"
	"log"
)

// Shared constants
var (
	region        = "us-east-1"
	prefixRegion  = "use1"
	prefixCompany = "jb"
	lob           = "test"
	application   = "cases"
	env           = "sandbox"
)

// Required tags
var requiredTags = []string{
	"module_project_path",
	"commit_id",
	"company",
	"region",
	"lob",
	"application",
	"env",
	"created_by",
	"map-migrated",
}

// Generic helper
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Function for policy test
func getPolicyJSON() string {
	policy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Action": []string{
					"s3:ListBucket",
					"s3:GetObject",
					"s3:PutObject",
				},
				"Resource": []string{
					"arn:aws:s3:::jb-test-cases-usw2-sandbox",
					"arn:aws:s3:::jb-test-cases-usw2-sandbox/*",
				},
			},
		},
	}

	jsonBytes, err := json.Marshal(policy)
	if err != nil {
		log.Fatalf("Failed to marshal policy JSON: %v", err)
	}

	return string(jsonBytes)
}

// Function for role trust policy test
func getTrustPolicyJSON() string {
	policy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Principal": map[string]interface{}{
					"Service": "ec2.amazonaws.com",
				},
				"Action": "sts:AssumeRole",
			},
		},
	}

	jsonBytes, err := json.Marshal(policy)
	if err != nil {
		log.Fatalf("Failed to marshal trust policy JSON: %v", err)
	}
	
	return string(jsonBytes)
}

