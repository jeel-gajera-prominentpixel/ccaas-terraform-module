package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestHoursOfOperation(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
	initialFailed := t.Failed()
	// Mocked Hours of Operation data for testing
	hoursOfOperationJSON := terraform.OutputJson(t, terraformOptions, "hours_of_operations")
	assert.NotEmpty(t, hoursOfOperationJSON, "Hours of operations should not be empty")

	// Parse the JSON string into a map and validate the Hours of Operation configuration.
	// This section performs detailed validation of the Hours of Operation output, including:
	// - Verifying the structure and format of the JSON data
	// - Validating the Hours of Operation name matches the expected value
	// - Checking the description and timezone settings
	// - Validating the instance ARN and ID associations
	// - Ensuring the configuration includes all 7 days of the week
	// - Verifying the time settings (hours and minutes) are within valid ranges
	// - Checking that all required tags are present

	var hoursOfOperation map[string]interface{}

	err := json.Unmarshal([]byte(hoursOfOperationJSON), &hoursOfOperation)
	assert.NoError(t, err, "Should be able to parse hours_of_operations JSON")

	Response := make(map[string]interface{})

	// Iterate through each Hours of Operation entry and validate its configuration
	for name, detailsRaw := range hoursOfOperation {
		assert.Equal(t, "test_hours_of_operations", name, "Hours of Operation name should match")
		assert.NotEmpty(t, detailsRaw, "Details should not be empty")
		// Convert the details to a JSON string for further processing
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
				// Validate the description field
				assert.Equal(t, "24-7 Music on Hold hours", detailsMap["description"], "Description should match expected value")

				// Validate the timezone setting
				assert.Equal(t, "US/Eastern", detailsMap["time_zone"], "Time zone should match expected value")

				// Verify all required tags are present
				assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

				// Validate the instance ARN matches the Hours of Operation ARN
				InstanceArnFromHOO := strings.Split(detailsMap["arn"].(string), "/")
				if len(InstanceArnFromHOO) >= 2 {
					InstanceArnOfHOO := strings.Join(InstanceArnFromHOO[:2], "/")
					HoursOfOpsIdOfHOO := InstanceArnFromHOO[len(InstanceArnFromHOO)-1]
					assert.Equal(t, instanceARN, InstanceArnOfHOO, "Instance ARN in Hours of Operation should match the instance ARN")
					assert.Equal(t, detailsMap["hours_of_operation_id"], HoursOfOpsIdOfHOO, "Hours of Operation ID in Hours of Operation ARN should match the Hours Of Operation ID")
				} else {
					assert.Fail(t, "Invalid ARN format")
				}

				// Validate instance ID association
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance ID in Hours of Operation should match the instance ID")

				// Verify the configuration includes all 7 days
				assert.Equal(t, 7, len(detailsMap["config"].([]interface{})), "Config should have 7 days of operations")

				// Validate the configuration for each day of the week
				for _, dayConfig := range detailsMap["config"].([]interface{}) {
					entryMap, ok := dayConfig.(map[string]interface{})
					assert.True(t, ok, "Each config entry should be a map")

					day, dayOk := entryMap["day"].(string)
					startTime, startOk := entryMap["start_time"].([]interface{})
					endTime, endOk := entryMap["end_time"].([]interface{})

					assert.True(t, dayOk, "day should be a string")
					assert.True(t, startOk, "start_time should be a map")
					assert.True(t, endOk, "end_time should be a map")

					// Validate time settings
					startHours, shOk := startTime[0].(map[string]interface{})["hours"].(float64)
					startMinutes, smOk := startTime[0].(map[string]interface{})["minutes"].(float64)
					endHours, ehOk := endTime[0].(map[string]interface{})["hours"].(float64)
					endMinutes, emOk := endTime[0].(map[string]interface{})["minutes"].(float64)

					assert.True(t, shOk && smOk, "start_time should contain hours and minutes as numbers")
					assert.True(t, ehOk && emOk, "end_time should contain hours and minutes as numbers")

					// Validate time ranges
					assert.GreaterOrEqual(t, startHours, float64(0))
					assert.LessOrEqual(t, startHours, float64(23))
					assert.GreaterOrEqual(t, startMinutes, float64(0))
					assert.LessOrEqual(t, startMinutes, float64(59))

					assert.GreaterOrEqual(t, endHours, float64(0))
					assert.LessOrEqual(t, endHours, float64(23))
					assert.GreaterOrEqual(t, endMinutes, float64(0))
					assert.LessOrEqual(t, endMinutes, float64(59))

					// Validate day names
					validDays := []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
					assert.Contains(t, validDays, day, "day should be a valid weekday")
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
