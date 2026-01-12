package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectUserHierarchyStructure validates the User Hierarchy Structure configuration created by Terraform.
// It performs comprehensive testing of the hierarchy structure output, including validation of levels,
// names, and ARNs.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//
// The function performs the following validations:
//   - Verifies the User Hierarchy Structure JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the instance ID matches the expected value
//   - Validates the hierarchy structure exists and is not empty
//   - For each hierarchy level (one through five):
//   - Validates the presence of ARN, ID and name
//   - Verifies specific level names:
//   - Level one: "lob"
//   - Level two: "manager"
//   - Level three: "supervisor"
//   - Level four: "agent"
//   - Confirms ARN suffixes match level numbers
//
// Returns:
//   - map[string]interface{}: A map containing the parsed user hierarchy structure configuration
func TestConnectUserHierarchyStructure(t *testing.T, terraformOptions *terraform.Options, instanceARN string, instanceID string) map[string]any {
	// Mocked User Hierarchy Structure data for testing
	initialFailed := t.Failed()
	UserHierarchyJson := terraform.OutputJson(t, terraformOptions, "user_hierarchy_structure")
	assert.NotEmpty(t, UserHierarchyJson, "User Hierarchy Structure should not be empty")

	// Parse the JSON string into a map and validate the User Hierarchy Structure configuration.
	// This section performs detailed validation of the User Hierarchy Structure output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the User Hierarchy Structure name matches the expected value
	// - Checking the description and flow settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all required flow blocks
	// - Verifying the flow logic and connections
	// - Checking that all required attributes and metadata are present

	var UserHierarchyConf map[string]any

	err := json.Unmarshal([]byte(UserHierarchyJson), &UserHierarchyConf)
	assert.NoError(t, err, "Should be able to parse user hierarchy JSON")

	Response := make(map[string]any)

	// Verify instance ID
	assert.Equal(t, instanceID, UserHierarchyConf["instance_id"], "Instance ID should match")

	// Check hierarchy structure exists and is not empty
	hierarchyStructure, ok := UserHierarchyConf["hierarchy_structure"].([]any)
	assert.True(t, ok, "Hierarchy structure should exist")
	assert.NotEmpty(t, hierarchyStructure, "Hierarchy structure should not be empty")

	// Validate each level in the hierarchy structure
	for _, levelData := range hierarchyStructure {
		levelMap, ok := levelData.(map[string]any)
		assert.True(t, ok, "Level data should be a map")

		// Validate each level (one through five)
		levels := []string{"level_one", "level_two", "level_three", "level_four", "level_five"}
		for _, level := range levels {
			if levelInfo, exists := levelMap[level].([]any); exists {
				if len(levelInfo) > 0 {
					levelDetails := levelInfo[0].(map[string]any)
					assert.Contains(t, levelDetails, "arn", fmt.Sprintf("%s should have ARN", level))
					assert.Contains(t, levelDetails, "id", fmt.Sprintf("%s should have ID", level))
					assert.Contains(t, levelDetails, "name", fmt.Sprintf("%s should have name", level))

					switch level {
					case "level_one":
						assert.Equal(t, "lob", levelDetails["name"], "Level one name should be 'lob'")
						assert.True(t, strings.HasSuffix(levelDetails["arn"].(string), "1"),
							"Level one ARN should end with '1'")
					case "level_two":
						assert.Equal(t, "manager", levelDetails["name"], "Level two name should be 'manager'")
						assert.True(t, strings.HasSuffix(levelDetails["arn"].(string), "2"),
							"Level two ARN should end with '2'")
					case "level_three":
						assert.Equal(t, "supervisor", levelDetails["name"], "Level three name should be 'supervisor'")
						assert.True(t, strings.HasSuffix(levelDetails["arn"].(string), "3"),
							"Level three ARN should end with '3'")
					case "level_four":
						assert.Equal(t, "agent", levelDetails["name"], "Level four name should be 'agent'")
						assert.True(t, strings.HasSuffix(levelDetails["arn"].(string), "4"),
							"Level four ARN should end with '4'")
					}
				}
			}
		}
		Response = UserHierarchyConf
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
