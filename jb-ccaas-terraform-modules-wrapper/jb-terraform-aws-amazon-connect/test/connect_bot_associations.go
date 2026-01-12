package test

import (
	"encoding/json"
	// "fmt"
	// "strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// File: test/connect_bot_associations_test.go
func TestConnectBotAssociations(t *testing.T, terraformOptions *terraform.Options, instanceARN string, instanceID string) map[string]interface{} {
    initialFailed := t.Failed()

    botAssociationsJSON := terraform.OutputJson(t, terraformOptions, "bot_associations")

    Response := map[string]interface{}{
        "Success": !t.Failed() || initialFailed == t.Failed(),
    }

    if botAssociationsJSON == "" || botAssociationsJSON == "{}" {
        t.Logf("No bot associations configured - this is expected for basic setups")
        return Response
    }

    var botAssociations map[string]interface{}
    err := json.Unmarshal([]byte(botAssociationsJSON), &botAssociations)
    assert.NoError(t, err, "Should be able to parse bot_associations JSON")

    for name, botRaw := range botAssociations {
        botMap := botRaw.(map[string]interface{})

        // Validate bot association structure
        assert.Equal(t, instanceID, botMap["instance_id"], "Instance ID should match")

        // Validate Lex bot configuration
        lexBot := botMap["lex_bot"].(map[string]interface{})
        assert.NotEmpty(t, lexBot["name"], "Lex bot name should not be empty")
        assert.NotEmpty(t, lexBot["lex_region"], "Lex region should not be empty")

        Response[name] = botMap
    }

    return Response
}
