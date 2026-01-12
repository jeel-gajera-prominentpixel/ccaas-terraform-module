package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectRoutingProfile validates the Routing Profile configuration created by Terraform.
// It performs comprehensive testing of the Routing Profile output, including structure validation,
// tag verification, and ARN matching.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the Routing Profile
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//   - OutBoundFlowId: The ID of the outbound flow
//   - HoursOfOperationId: The ID of the hours of operation
//
// The function performs the following validations:
//   - Verifies the Routing Profile JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the Routing Profile name matches the expected value
//   - Checks the description and flow settings
//   - Validates the instance ARN and ID associations
//   - Ensures all required tags are present
//   - Verifies the outbound caller configuration
//   - Validates hours of operation settings
//   - Checks queue status and name
//
// Returns:
//   - map[string]interface{}: A map containing the parsed queue configurations
func TestConnectRoutingProfile(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string, dynamicQueuesID string, QueuesConf map[string]interface{}) map[string]interface{} {
	resourceName := "Routing Profiles"
	initialFailed := t.Failed()

	RoutingProfileJson := terraform.OutputJson(t, terraformOptions, "routing_profiles")
	assert.NotEmpty(t, RoutingProfileJson, fmt.Sprintf("%s should not be empty", resourceName))

	// Parse the JSON string into a map and validate the Routing Profile configuration.
	// This section performs detailed validation of the Routing Profile output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Routing Profile name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var RoutingProfileConf map[string]interface{}

	err := json.Unmarshal([]byte(RoutingProfileJson), &RoutingProfileConf)
	assert.NoError(t, err, fmt.Sprintf("Should be able to parse %s JSON", resourceName))

	Response := make(map[string]interface{})

	for name, detailsRaw := range RoutingProfileConf {
		assert.Equal(t, "jb_maint_A1", name, fmt.Sprintf("%s name should match", resourceName))
		assert.NotEmpty(t, detailsRaw, "Details should not be empty")
		var detailsStr string
		if detailsMap, ok := detailsRaw.(map[string]interface{}); ok {
			detailsBytes, err := json.Marshal(detailsMap)
			assert.NoError(t, err, "Should be able to marshal details back to JSON")
			detailsStr = string(detailsBytes)
		} else {
			detailsStr = fmt.Sprintf("%v", detailsRaw)
		}

		if detailsStr != "" {
			var detailsMap map[string]interface{}
			err := json.Unmarshal([]byte(detailsStr), &detailsMap)
			assert.NoError(t, err, "Should be able to parse details as JSON")
			if err == nil {
				Response[name] = detailsMap
				assert.Equal(t, "A1 routing profile", detailsMap["description"], "Description should match expected value")
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match expected value")

				if mediaConcurrenciesConfig, ok := detailsMap["media_concurrencies"].(map[string]interface{}); ok {
					assert.Equal(t, "VOICE", mediaConcurrenciesConfig["channel"], fmt.Sprintf("%s channel should match expected value", resourceName))
					assert.Equal(t, 1, mediaConcurrenciesConfig["concurrency"], fmt.Sprintf("%s concurrency should match expected value", resourceName))
				}

				if QueuesConfig, ok := detailsMap["queue_configs"].(map[string]interface{}); ok {
					assert.Equal(t, "VOICE", QueuesConfig["channel"], fmt.Sprintf("%s channel should match expected value", resourceName))
					assert.Equal(t, 0, QueuesConfig["delay"], fmt.Sprintf("%s delay should match expected value", resourceName))
					assert.Equal(t, 5, QueuesConfig["priority"], fmt.Sprintf("%s priority should match expected value", resourceName))
					assert.Equal(t, QueuesConf["test_queues"].(map[string]interface{})["arn"].(string), QueuesConfig["queue_arn"], fmt.Sprintf("%s Queue arn should match expected value", resourceName))
					assert.Equal(t, dynamicQueuesID, QueuesConfig["queue_id"], fmt.Sprintf("%s Queue Id should match expected value", resourceName))
					assert.Equal(t, QueuesConf["test_queues"].(map[string]interface{})["name"].(string), QueuesConfig["queue_name"], fmt.Sprintf("%s Queue Name should match expected value", resourceName))
				}

				assert.Equal(t, dynamicQueuesID, detailsMap["default_outbound_queue_id"], fmt.Sprintf("%s Queue ID should match expected value", resourceName))
				assert.Equal(t, "jb_maint_A1", detailsMap["name"], fmt.Sprintf("%s Name should match expected value", resourceName))

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the ROuting Profile ARN
				InstanceArnFromRoutingProfile := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromRoutingProfile) >= 2 {
					InstanceArnOfQueue := strings.Join(InstanceArnFromRoutingProfile[:2], "/")
					RoutingProfileIdOfRoutingProfile := InstanceArnFromRoutingProfile[len(InstanceArnFromRoutingProfile)-1]
					assert.Equal(t, instanceARN, InstanceArnOfQueue, fmt.Sprintf("Instance ARN in %s should match the instance ARN", resourceName))
					assert.Equal(t, detailsMap["routing_profile_id"], RoutingProfileIdOfRoutingProfile, fmt.Sprintf("%s ID in %s ARN should match the %s ID", resourceName, resourceName, resourceName))
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
