// Package test provides testing utilities for the ECS Fargate wrapper module
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
	hash_key                    = "quickConnectName"
	range_key                   = "lob"
	deletion_protection_enabled = false
	stream_enabled              = true
	stream_view_type            = "NEW_AND_OLD_IMAGES"
	attributes                  = []map[string]string{
		{
			"name": "quickConnectName",
			"type": "S",
		},
		{
			"name": "lob",
			"type": "S",
		},
		{
			"name": "quickConnectType",
			"type": "S",
		},
	}
	global_secondary_indexes = []map[string]string{
		{
			"name":            "lob-index",
			"hash_key":        "lob",
			"range_key":       "quickConnectType",
			"projection_type": "ALL",
		},
	}
	replica_regions = []map[string]any{
		{
			"region_name":                 "us-east-1",
			"deletion_protection_enabled": false,
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
		"name":                        fmt.Sprintf("%s-dyndb-%s-%s-%s", prefixCompany, lob, application, env),
		"hash_key":                    hash_key,
		"range_key":                   range_key,
		"deletion_protection_enabled": deletion_protection_enabled,
		"stream_enabled":              stream_enabled,
		"stream_view_type":            stream_view_type,
		"attributes":                  attributes,
		"global_secondary_indexes":    global_secondary_indexes,
		"replica_regions":             replica_regions,
		"tags": map[string]string{
			"company": prefixCompany,
			"lob":     lob,
		},
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

func nameBase() string {
return fmt.Sprintf("%s-ecs-%s-%s-%s", prefixCompany, lob, application, envName)
}
// getCommonVars maps directly to your wrapper's input variables.
// Rename keys here to match your module exactly.
func getCommonVars() map[string]interface{} {
return map[string]interface{}{
"region":                    region,
"prefix_company":            prefixCompany,
"prefix_region":             prefixRegion,
"lob":                       lob,
"application":               application,
"env":                       envName,
"name":                      nameBase(),
// Network
"vpc_id":                    vpcID,
"subnet_ids":                privateSubnets,
"security_group_ids":        securityGroups,
// Task/Service
"desired_count":             desiredCount,
"task_cpu":                  cpu,
"task_memory":               memory,
"assign_public_ip":          assignPublicIP,
"enable_execute_command":    enableExecuteCmd,
// Container
"container_name":            containerName,
"container_image":           containerImage,
"container_port":            containerPort,
"healthcheck_path":          healthCheckPath,
// Observability
"logs_retention_in_days":    logsRetentionDays,
// LB (toggle if your wrapper supports it)
"enable_load_balancer":      loadbalancerEnabled,
// Tags
"tags": map[string]string{
"company":     prefixCompany,
"lob":         lob,
"application": application,
"env":         envName,
"region":      region,
"created_by":  "terratest",
},
}
}
//
// ---------- Test Harness (apply once) ----------
//
var tfOpts *terraform.Options
func TestMain(m *testing.M) {
tfOpts = terraform.WithDefaultRetryableErrors(&testing.T{}, &terraform.Options{
TerraformDir: "../",
Vars:         getCommonVars(),
NoColor:      true,
NoStderr:     true,
Logger:       logger.Discard,
EnvVars: map[string]string{
"TF_LOG":      "ERROR",
"TF_LOG_PATH": "/dev/null",
},
})
log.SetOutput(io.Discard)
// If you want to only plan in CI sometimes:
// if getenvDefault("TEST_PLAN_ONLY", "false") == "true" {
// terraform.InitAndPlan(&testing.T{}, tfOpts)
// } else {
// terraform.InitAndApply(&testing.T{}, tfOpts)
// }
terraform.InitAndApply(&testing.T{}, tfOpts)
code := m.Run()
terraform.Destroy(&testing.T{}, tfOpts)
os.Exit(code)
}
//
// ---------- Tests ----------
//
func TestECSFargate_OutputsExist(t *testing.T) {
t.Parallel()
all := terraform.OutputAll(t, tfOpts)
asJSON, _ := json.Marshal(all)
t.Logf("All outputs: %s", string(asJSON))
// These output names should match what your wrapper exports.
// Swap/rename as needed.
assertOutputString(t, all, "ecs_cluster_id")
assertOutputString(t, all, "ecs_cluster_name")
assertOutputString(t, all, "ecs_service_name")
assertOutputString(t, all, "task_definition_arn")
assertOutputString(t, all, "task_definition_family")
// Optional/conditional outputs (won’t fail the test if absent)
assertOutputStringOptional(t, all, "service_security_group_id")
assertOutputStringOptional(t, all, "cloudwatch_log_group_name")
assertOutputStringOptional(t, all, "alb_target_group_arn")
}
func TestECSFargate_ServiceConfiguration(t *testing.T) {
t.Parallel()
all := terraform.OutputAll(t, tfOpts)
// Example validations; rename keys to your outputs.
assert.Equal(t, "FARGATE", getString(all, "launch_type"), "launch_type should be FARGATE")
if v, ok := all["desired_count"]; ok {
assert.Equal(t, float64(desiredCount), v, "desired_count should match input")
}
// Network expectations
if subnets, ok := all["service_subnet_ids"].([]interface{}); ok && len(privateSubnets) > 0 {
assert.True(t, len(subnets) >= 1, "service_subnet_ids should not be empty")
}
// Exec command toggle (if you surface it)
if v, ok := all["enable_execute_command"]; ok {
assert.Equal(t, enableExecuteCmd, v.(bool), "enable_execute_command should match input")
}
}
func TestECSFargate_NamingConventions(t *testing.T) {
t.Parallel()
all := terraform.OutputAll(t, tfOpts)
expectedPrefix := nameBase()
if svc, ok := all["ecs_service_name"].(string); ok && svc != "" {
assert.Contains(t, svc, expectedPrefix, "service name should contain base name")
}
if fam, ok := all["task_definition_family"].(string); ok && fam != "" {
assert.Contains(t, fam, expectedPrefix, "task definition family should contain base name")
}
}
//
// ---------- Helpers ----------
//
func assertOutputString(t *testing.T, outputs map[string]interface{}, key string) {
val, ok := outputs[key]
assert.True(t, ok, "output %q should exist", key)
if ok {
_, isStr := val.(string)
assert.True(t, isStr, "output %q should be a string", key)
}
}
func assertOutputStringOptional(t *testing.T, outputs map[string]interface{}, key string) {
val, ok := outputs[key]
if ok {
_, isStr := val.(string)
assert.True(t, isStr, "optional output %q should be a string when present", key)
}
}
func getString(outputs map[string]interface{}, key string) string {
if v, ok := outputs[key]; ok {
if s, ok := v.(string); ok {
return s
}
}
return ""
}
func splitCSV(s string) []string {
if s == "" {
return []string{}
}
var out []string
cur := ""
for _, r := range s {
if r == ',' {
if cur != "" {
out = append(out, cur)
cur = ""
}
} else {
cur += string(r)
}
}
if cur != "" {
out = append(out, cur)
}
return out
}
func getenvDefault(k, d string) string {
if v := os.Getenv(k); v != "" {
return v
}
return d
}
func mustAtoi(s string) int {
var n int
for _, r := range s {
if r < '0' || r > '9' {
return 0
}
n = n*10 + int(r-'0')
}
return n
}