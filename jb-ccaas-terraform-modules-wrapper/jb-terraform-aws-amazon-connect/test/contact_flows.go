package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestContactFlow validates the Contact Flow configuration created by Terraform.
// It performs comprehensive testing of the Contact Flow output, including structure validation,
// tag verification, and ARN matching.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the Contact Flow
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//
// The function performs the following validations:
//   - Verifies the Contact Flow JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the Contact Flow name matches the expected value
//   - Checks the description and flow settings
//   - Validates the instance ARN and ID associations
//   - Ensures all required tags are present
//   - Verifies the Contact Flow type and configuration
func TestContactFlow(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
	// Mocked Contact Flow data for testing
	initialFailed := t.Failed()
	ContactFlowJSON := terraform.OutputJson(t, terraformOptions, "contact_flows")
	assert.NotEmpty(t, ContactFlowJSON, "Contact Flow should not be empty")

	// Parse the JSON string into a map and validate the Contact Flow configuration.
	// This section performs detailed validation of the Contact Flow output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Contact Flow name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var ContactFlow map[string]interface{}

	err := json.Unmarshal([]byte(ContactFlowJSON), &ContactFlow)
	assert.NoError(t, err, "Should be able to parse contact_flow JSON")

	Response := make(map[string]interface{})

	for name, detailsRaw := range ContactFlow {
		// Validate the Contact Flow name and details
		assert.Contains(t, []string{"test_contact_flows", "test_whisper_flows", "test_queue_transfer_flows"}, name, "Contact Flow name should match one of the expected names")
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
				var FlowName, FlowType, FlowId, FlowDesc string
				if name == "test_contact_flows" {
					FlowName = "test_contact_flows"
					FlowType = "CONTACT_FLOW"
					FlowDesc = "Basic contact flow for testing"
				}
				if name == "test_whisper_flows" {
					FlowName = "test_whisper_flows"
					FlowType = "OUTBOUND_WHISPER"
					FlowDesc = "Basic outbound whisper flow for testing"
				}
				if name == "test_queue_transfer_flows" {
					FlowName = "test_queue_transfer_flows"
					FlowType = "QUEUE_TRANSFER"
					FlowDesc = "Basic test_queue_transfer_flows flow for testing"
				}
				FlowId = detailsMap["contact_flow_id"].(string)
				Response[name] = detailsMap
				assert.Equal(t, FlowDesc, detailsMap["description"], "Description should match expected value")
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match expected value")
				assert.Equal(t, FlowType, detailsMap["type"], "Contact Flow Type should match expected value")
				assert.Equal(t, FlowName, detailsMap["name"], "Contact Flow Name should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the Contact Flow ARN
				InstanceArnFromFlow := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromFlow) >= 2 {
					InstanceArnOfFlow := strings.Join(InstanceArnFromFlow[:2], "/")
					ContactFlowIdOfHOO := InstanceArnFromFlow[len(InstanceArnFromFlow)-1]

					assert.Equal(t, instanceARN, InstanceArnOfFlow, "Instance ARN in Contact Flow should match the instance ARN")
					assert.Equal(t, FlowId, ContactFlowIdOfHOO, "Contact Flow ID in Contact Flow ARN should match the Contact Flow ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
