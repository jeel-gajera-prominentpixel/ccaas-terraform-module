// TestConnectQuickConnect validates the configuration of Amazon Connect Quick Connects created by Terraform.
//
// The function performs comprehensive testing of Quick Connects, including:
// - Parsing and validating the Quick Connect JSON output
// - Verifying the Quick Connect name matches "test_quick_connect"
// - Validating instance ID matches the provided value
// - Checking queue configuration details including contact flow ID and queue ID
// - Validating instance ARN and Quick Connect ID from the ARN
// - Ensuring required tags are present
//
// Parameters:
//   - t: The testing framework object
//   - terraformOptions: Terraform configuration options for the test
//   - requiredTags: List of tags that must be present in the Quick Connect
//   - instanceARN: The Amazon Connect instance ARN
//   - instanceID: The Amazon Connect instance ID
//   - dynamicQueueTransferFlowID: ID of the contact flow for queue transfer
//   - dynamicQueuesID: ID of the associated queue
//
// Returns a map containing the parsed Quick Connect configuration details.
package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectQuickConnect validates the Quick Connect configuration created by Terraform.
// It performs comprehensive testing of the Quick Connect output, including structure validation,
// tag verification, and ARN matching.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the Quick Connect
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//   - dynamicQueueTransferFlowID: ID of the contact flow for queue transfer
//   - dynamicQueuesID: ID of the associated queue
//
// The function performs the following validations:
//   - Verifies the Quick Connect JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the Quick Connect name matches "test_quick_connect"
//   - Validates the instance ID matches the provided value
//   - Checks the queue configuration including contact flow ID and queue ID
//   - Validates the instance ARN and Quick Connect ID from the ARN
//   - Ensures all required tags are present
func TestConnectQuickConnect(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string, dynamicQueueTransferFlowID string, dynamicQueuesID string) map[string]interface{} {
	// Get Quick Connect JSON output from Terraform
	initialFailed := t.Failed()
	QuickConnectJson := terraform.OutputJson(t, terraformOptions, "quick_connects")
	assert.NotEmpty(t, QuickConnectJson, "Quick Connect should not be empty")

	// Parse the JSON string into a map and validate the Quick Connect configuration.
	// This section performs detailed validation of the Quick Connect output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Quick Connect name matches "test_quick_connect"
	// - Checking the instance ID matches the provided value
	// - Validating queue configuration details
	// - Verifying the ARN format and extracting instance ARN and Quick Connect ID
	// - Ensuring all required tags are present

	var QuickConnectConf map[string]interface{}

	err := json.Unmarshal([]byte(QuickConnectJson), &QuickConnectConf)
	assert.NoError(t, err, "Should be able to parse Quick Connect JSON")

	Response := make(map[string]interface{})

	for name, detailsRaw := range QuickConnectConf {
		assert.Equal(t, "test_quick_connect", name, "Quick Connect name should match")
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
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match expected value")

				if quickConnectConfig, ok := detailsMap["quick_connect_config"].([]interface{}); ok && len(quickConnectConfig) > 0 {
					if queueConfig, ok := quickConnectConfig[0].(map[string]interface{})["queue_config"].(map[string]interface{}); ok {
						assert.Equal(t, dynamicQueueTransferFlowID, queueConfig["contact_flow_id"], "Queues Type should match expected value")
						assert.Equal(t, dynamicQueuesID, queueConfig["queue_id"], "Queues Type should match expected value")
					}
				}

				assert.Equal(t, "test_quick_connect", detailsMap["name"], "Quick Connect Name should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the Quick Connect ARN
				InstanceArnFromQuickConnect := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromQuickConnect) >= 2 {
					InstanceArnOfQuickConnect := strings.Join(InstanceArnFromQuickConnect[:2], "/")
					QuickConnectIdOfQueues := InstanceArnFromQuickConnect[len(InstanceArnFromQuickConnect)-1]

					assert.Equal(t, instanceARN, InstanceArnOfQuickConnect, "Instance ARN in QuickConnect should match the instance ARN")
					assert.Equal(t, detailsMap["quick_connect_id"], QuickConnectIdOfQueues, "QuickConnect ID in QuickConnect ARN should match the QuickConnect ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
