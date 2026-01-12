package test

import (
	"encoding/json"
	// "fmt"
	// "strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// File: test/connect_vocabularies_test.go
func TestConnectVocabularies(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
    initialFailed := t.Failed()

    vocabulariesJSON := terraform.OutputJson(t, terraformOptions, "vocabularies")

    // Vocabularies might be empty, so we handle both cases
    Response := map[string]interface{}{
        "Success": !t.Failed() || initialFailed == t.Failed(),
    }

    if vocabulariesJSON == "" || vocabulariesJSON == "{}" {
        t.Logf("No vocabularies configured - this is expected for basic setups")
        return Response
    }

    var vocabularies map[string]interface{}
    err := json.Unmarshal([]byte(vocabulariesJSON), &vocabularies)
    assert.NoError(t, err, "Should be able to parse vocabularies JSON")

    for name, vocabRaw := range vocabularies {
        vocabMap := vocabRaw.(map[string]interface{})

        // Validate vocabulary structure
        assert.NotEmpty(t, vocabMap["vocabulary_id"], "Vocabulary ID should not be empty")
        assert.Equal(t, instanceID, vocabMap["instance_id"], "Instance ID should match")

        // Validate ARN format
        arn := vocabMap["arn"].(string)
        assert.Contains(t, arn, instanceARN, "Vocabulary ARN should contain instance ARN")
        assert.Contains(t, arn, "vocabulary", "ARN should contain vocabulary")

        // Validate language code
        assert.NotEmpty(t, vocabMap["language_code"], "Language code should not be empty")

        // Validate tags if present
        if tags, ok := vocabMap["tags"].(map[string]interface{}); ok && len(tags) > 0 {
            assertTagsExist(t, tags, requiredTags)
        }

        Response[name] = vocabMap
    }

    return Response
}
