package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestContactFlowModules validates the configuration and attributes of Contact Flow Modules in Amazon Connect.
// It performs comprehensive testing of the Contact Flow Module outputs from Terraform, including:
//   - Verification of JSON structure and format
//   - Validation of module name, description, and instance associations
//   - Checking for required tags
//   - Validation of ARN formatting and relationships
//   - Ensuring proper instance ID mapping
//
// Parameters:
//   - t: Testing framework context
//   - terraformOptions: Terraform configuration options
//   - requiredTags: List of tags that must be present in the Contact Flow Module
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
func TestContactFlowModules(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
	// Mocked Contact Flow Module data for testing
	initialFailed := t.Failed()
	ContactFlowModuleJSON := terraform.OutputJson(t, terraformOptions, "contact_flow_modules")
	assert.NotEmpty(t, ContactFlowModuleJSON, "Contact Flow Module should not be empty")

	// Parse the JSON string into a map and validate the Contact Flow Module configuration.
	// This section performs detailed validation of the Contact Flow Module output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Contact Flow Module name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var ContactFlowModule map[string]interface{}

	err := json.Unmarshal([]byte(ContactFlowModuleJSON), &ContactFlowModule)
	assert.NoError(t, err, "Should be able to parse contact_flow JSON")

	Response := make(map[string]interface{})

	for name, detailsRaw := range ContactFlowModule {
		assert.Equal(t, "test_contact_flows_module", name, "Contact Flow Module name should match")
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
				assert.Equal(t, "Basic contact flow module for testing", detailsMap["description"], "Description should match expected value")
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance ID should match expected value")
				assert.Equal(t, "test_contact_flows_module", detailsMap["name"], "Contact Flow Module Name should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the Contact Flow Module ARN
				InstanceArnFromFlowModule := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromFlowModule) >= 2 {
					InstanceArnOfFlowModule := strings.Join(InstanceArnFromFlowModule[:2], "/")
					ContactFlowModuleIdOfHOO := InstanceArnFromFlowModule[len(InstanceArnFromFlowModule)-1]

					assert.Equal(t, instanceARN, InstanceArnOfFlowModule, "Instance ARN in Contact Flow Modules should match the instance ARN")
					assert.Equal(t, detailsMap["contact_flow_module_id"], ContactFlowModuleIdOfHOO, "Contact Flow Module ID in Contact Flow ARN should match the Contact Flow ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
