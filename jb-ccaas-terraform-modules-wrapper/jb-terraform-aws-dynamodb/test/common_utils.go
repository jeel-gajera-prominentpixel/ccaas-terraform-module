package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// assertOutputsNonEmpty checks that the specified Terraform outputs are non-empty
func assertOutputsNonEmpty(t *testing.T, options *terraform.Options, outputs []string) {
	for _, output := range outputs {
		value := terraform.Output(t, options, output)
		assert.NotEmpty(t, value, "%s should not be empty", output)
	}
}

// assertNameFormat ensures the resource name follows the expected format
func assertNameFormat(t *testing.T, name, expectedPrefix string) {
	assert.True(t,
		strings.HasPrefix(name, expectedPrefix),
		"Name must start with '%s', got '%s'", expectedPrefix, name)
}

// // assertTagsExist checks if required tags exist in a given resource state
func assertTagsExist(t *testing.T, state map[string]interface{}, requiredTags []string) {
	var missingTags []string
	for _, tag := range requiredTags {
		if _, ok := state[tag]; !ok {
			missingTags = append(missingTags, tag)
		}
	}

	if len(missingTags) > 0 {
		assert.Fail(t, fmt.Sprintf("Missing required tags: %v", missingTags))
	}
}

// assertContains verifies that specific key-value pairs exist in the plan output using regular expressions
func assertContains(t *testing.T, plan string, checks map[string]string) {
	for key, expectedValue := range checks {
		// Create a flexible regex to match the key-value pair, allowing spaces around '=' and ignoring formatting issues
		regexPattern := key + `\s*=\s*` + regexp.QuoteMeta(expectedValue)

		// Compile the regex
		matcher, err := regexp.Compile(regexPattern)
		assert.NoError(t, err, "Failed to compile regex for key '%s' with value '%s'", key, expectedValue)

		// Check if the regex matches the plan output
		assert.True(t,
			matcher.MatchString(plan),
			"Expected '%s' to have value '%s' in the plan output, but it was not found", key, expectedValue)
	}
}

// mergeMaps combines two maps into one
func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}
