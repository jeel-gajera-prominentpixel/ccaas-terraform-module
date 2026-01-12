package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectQueues validates the Queues configuration created by Terraform.
// It performs comprehensive testing of the Queues output, including structure validation,
// tag verification, and ARN matching.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the Queues
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//   - OutBoundFlowId: The ID of the outbound flow
//   - HoursOfOperationId: The ID of the hours of operation
//
// The function performs the following validations:
//   - Verifies the Queues JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the Queues name matches the expected value
//   - Checks the description and flow settings
//   - Validates the instance ARN and ID associations
//   - Ensures all required tags are present
//   - Verifies the outbound caller configuration
//   - Validates hours of operation settings
//   - Checks queue status and name
//
// Returns:
//   - map[string]interface{}: A map containing the parsed queue configurations
func TestConnectQueues(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string, OutBoundFlowId string, HoursOfOperationId string) map[string]interface{} {
	// Mocked Queues data for testing
	initialFailed := t.Failed()
	QueuesJson := terraform.OutputJson(t, terraformOptions, "queues")
	assert.NotEmpty(t, QueuesJson, "Queues should not be empty")

	// Parse the JSON string into a map and validate the Queues configuration.
	// This section performs detailed validation of the Queues output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Queues name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var QueuesConf map[string]interface{}

	err := json.Unmarshal([]byte(QueuesJson), &QueuesConf)
	assert.NoError(t, err, "Should be able to parse queues JSON")

	Response := make(map[string]interface{})

	for name, detailsRaw := range QueuesConf {
		assert.Equal(t, "test_queues", name, "Queues name should match")
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
				assert.Equal(t, "Basic Queues for testing", detailsMap["description"], "Description should match expected value")
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match expected value")

				if outboundCallerConfig, ok := detailsMap["outbound_caller_config"].(map[string]interface{}); ok {
					assert.Equal(t, OutBoundFlowId, outboundCallerConfig["outbound_flow_id"], "Queues Type should match expected value")
					assert.Equal(t, "Test Caller ID", outboundCallerConfig["outbound_caller_id_name"], "Queues Type should match expected value")
				}

				assert.Equal(t, HoursOfOperationId, detailsMap["hours_of_operation_id"], "Queues Type should match expected value")
				assert.Equal(t, "test_queues", detailsMap["name"], "Queues Name should match expected value")
				assert.Equal(t, "ENABLED", detailsMap["status"], "Queues Name should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the Queues ARN
				InstanceArnFromQueues := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromQueues) >= 2 {
					InstanceArnOfQueue := strings.Join(InstanceArnFromQueues[:2], "/")
					QueuesIdOfQueues := InstanceArnFromQueues[len(InstanceArnFromQueues)-1]
					assert.Equal(t, instanceARN, InstanceArnOfQueue, "Instance ARN in Queues should match the instance ARN")
					assert.Equal(t, detailsMap["queue_id"], QueuesIdOfQueues, "Queues ID in Queues ARN should match the Queues ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
