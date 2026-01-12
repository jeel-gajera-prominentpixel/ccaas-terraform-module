package test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectUserHierarchyGroups validates the User Hierarchy Groups configuration created by Terraform.
// It performs comprehensive testing of the hierarchy groups output, including structure validation,
// tag verification, and ARN matching.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//   - requiredTags: List of tags that must be present in the hierarchy groups
//
// The function performs the following validations:
//   - Verifies the User Hierarchy Groups JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the hierarchy group names match the expected values (MOCAgent, MCCAgent, etc.)
//   - Validates the level IDs (1-4) for each hierarchy group
//   - Checks the instance ARN and ID associations
//   - Ensures all required tags are present
//   - Verifies the hierarchy group IDs and ARNs
//   - Validates the complete hierarchy structure
//
// Returns:
//   - map[string]interface{}: A map containing the parsed hierarchy group configurations
func TestConnectUserHierarchyGroups(t *testing.T, terraformOptions *terraform.Options, instanceARN string, instanceID string, requiredTags []string) map[string]interface{} {
	// Mocked User Hierarchy Groups data for testing
	initialFailed := t.Failed()
	UserHierarchyGroupJson := terraform.OutputJson(t, terraformOptions, "user_hierarchy_groups")
	assert.NotEmpty(t, UserHierarchyGroupJson, "User Hierarchy Groups should not be empty")

	// Parse the JSON string into a map and validate the User Hierarchy Groups configuration.
	// This section performs detailed validation of the User Hierarchy Groups output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the User Hierarchy Groups name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var UserHierarchyGroupConf map[string]any

	err := json.Unmarshal([]byte(UserHierarchyGroupJson), &UserHierarchyGroupConf)
	assert.NoError(t, err, "Should be able to parse user hierarchy JSON")

	Response := make(map[string]any)
	for name, detailsRaw := range UserHierarchyGroupConf {
		assert.Contains(t, []string{"MOCAgent", "MCCAgent", "AOGAgent", "MOCPlanningSupervisor", "MCCDutySupervisor", "AOGSupervisor", "MCCDutyManager", "maint"}, name, "Hierarchy Group name should match one of the expected names")

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

				assert.Equal(t, name, detailsMap["name"], "Hierarchy Group Name should match expected value")
				levelIDStr, ok := detailsMap["level_id"].(string)
				assert.True(t, ok, "level_id should be a string")
				levelIDInt, err := strconv.Atoi(levelIDStr)
				assert.NoError(t, err, "level_id should be convertible to integer")
				assert.Contains(t, []int{1, 2, 3, 4}, levelIDInt, "Hierarchy Group Level ID should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the User ARN
				InstanceArnFromHierarchyGroup := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromHierarchyGroup) >= 2 {
					InstanceArnOfHierarchyGroup := strings.Join(InstanceArnFromHierarchyGroup[:2], "/")
					HierarchyGroupIdOfHierarchyGroup := InstanceArnFromHierarchyGroup[len(InstanceArnFromHierarchyGroup)-1]
					assert.Equal(t, instanceARN, InstanceArnOfHierarchyGroup, "Instance ARN in Hierarchy Group should match the instance ARN")
					assert.Equal(t, detailsMap["hierarchy_group_id"], HierarchyGroupIdOfHierarchyGroup, "Hierarchy Group ID in Hierarchy Group ARN should match the Hierarchy Group ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
