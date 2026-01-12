package test

import (
	"encoding/json"
	// "fmt"
	// "strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// File: test/connect_users_test.go
func TestConnectUsers(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
    initialFailed := t.Failed()

    usersJSON := terraform.OutputJson(t, terraformOptions, "users")

    Response := map[string]interface{}{
        "Success": !t.Failed() || initialFailed == t.Failed(),
    }

    if usersJSON == "" || usersJSON == "{}" {
        t.Logf("No users configured - this is expected for basic setups")
        return Response
    }

    var users map[string]interface{}
    err := json.Unmarshal([]byte(usersJSON), &users)
    assert.NoError(t, err, "Should be able to parse users JSON")

    for username, userRaw := range users {
        userMap := userRaw.(map[string]interface{})

        // Validate user structure
        assert.NotEmpty(t, userMap["user_id"], "User ID should not be empty")
        assert.Equal(t, instanceID, userMap["instance_id"], "Instance ID should match")
        assert.Equal(t, username, userMap["username"], "Username should match")

        // Validate ARN format
        arn := userMap["arn"].(string)
        assert.Contains(t, arn, instanceARN, "User ARN should contain instance ARN")
        assert.Contains(t, arn, "user", "ARN should contain user")

        // Validate identity info
        identityInfo := userMap["identity_info"].(map[string]interface{})
        assert.NotEmpty(t, identityInfo["first_name"], "First name should not be empty")
        assert.NotEmpty(t, identityInfo["last_name"], "Last name should not be empty")

        // Validate phone config if present
        if phoneConfig, ok := userMap["phone_config"].(map[string]interface{}); ok {
            assert.Contains(t, []string{"SOFT_PHONE", "DESK_PHONE"}, phoneConfig["phone_type"], "Phone type should be valid")
        }

        // Validate security profile IDs
        securityProfileIds := userMap["security_profile_ids"].([]interface{})
        assert.NotEmpty(t, securityProfileIds, "Security profile IDs should not be empty")

        // Validate tags if present
        if tags, ok := userMap["tags"].(map[string]interface{}); ok && len(tags) > 0 {
            assertTagsExist(t, tags, requiredTags)
        }

        Response[username] = userMap
    }

    return Response
}
