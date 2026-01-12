package test

import (
	"encoding/json"
	// "fmt"
	// "strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// File: test/connect_security_profiles_test.go
func TestConnectSecurityProfiles(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
    initialFailed := t.Failed()

    securityProfilesJSON := terraform.OutputJson(t, terraformOptions, "security_profiles")
    assert.NotEmpty(t, securityProfilesJSON, "Security profiles should not be empty")

    var securityProfiles map[string]interface{}
    err := json.Unmarshal([]byte(securityProfilesJSON), &securityProfiles)
    assert.NoError(t, err, "Should be able to parse security_profiles JSON")

    Response := make(map[string]interface{})

    for name, profileRaw := range securityProfiles {
        profileMap := profileRaw.(map[string]interface{})

        // Validate security profile structure
        assert.NotEmpty(t, profileMap["security_profile_id"], "Security profile ID should not be empty")
        assert.Equal(t, instanceID, profileMap["instance_id"], "Instance ID should match")

        // Validate ARN format
        arn := profileMap["arn"].(string)
        assert.Contains(t, arn, instanceARN, "Security profile ARN should contain instance ARN")
        assert.Contains(t, arn, "security-profile", "ARN should contain security-profile")

        // Validate permissions structure if present
        if permissions, ok := profileMap["permissions"].([]interface{}); ok {
            assert.NotEmpty(t, permissions, "Permissions should not be empty if present")
        }

        // Validate tags
        if tags, ok := profileMap["tags"].(map[string]interface{}); ok && len(tags) > 0 {
            assertTagsExist(t, tags, requiredTags)
        }

        Response[name] = profileMap
    }

    Response["Success"] = !t.Failed() || initialFailed == t.Failed()
    return Response
}
