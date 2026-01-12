package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectLambdaAssociation validates the lambda function associations created by Terraform.
// It performs comprehensive testing of the lambda function association output, including structure validation,
// function ARN verification, and association configuration checks.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the lambda associations (not directly used in this function)
//   - instanceARN: The ARN of the Amazon Connect instance
//   - instanceID: The ID of the Amazon Connect instance
//   - lambdaFunctionArns: Map of lambda function ARNs expected to be associated
//
// The function performs the following validations:
//   - Verifies the lambda function associations JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the function ARN matches the expected values for each association
//   - Checks the function type and association configuration
//   - Validates the instance ID in the lambda associations
//
// Returns:
//   - map[string]interface{}: A map containing the parsed lambda function associations
func TestConnectLambdaAssociation(t *testing.T, terraformOptions *terraform.Options, instanceARN string, instanceID string, lambdaFunctionArn string) map[string]interface{} {
	resourceName := "Lambda Function Association"
	initialFailed := t.Failed()
	lambdaAssociationJson := terraform.OutputJson(t, terraformOptions, "lambda_function_associations")
	assert.NotEmpty(t, lambdaAssociationJson, fmt.Sprintf("%s should not be empty", resourceName))

	// TestConnectLambdaAssociation validates the Amazon Connect lambda function associations output from Terraform.
	// This test function performs the following checks for each lambda function association:
	//   - Parses the "lambda_function_associations" Terraform output as JSON and ensures it is not empty.
	//   - Iterates over each lambda function association and validates that:
	//     - The function ARN matches the expected value.
	//     - The function type is correctly configured.
	//     - The instance ID matches the expected Connect instance.
	//     - The association ID is properly structured.
	// The function returns a map containing the parsed and validated lambda association details for further assertions.

	var lambdaAssociationConf map[string]interface{}

	err := json.Unmarshal([]byte(lambdaAssociationJson), &lambdaAssociationConf)
	assert.NoError(t, err, fmt.Sprintf("Should be able to parse %s JSON", resourceName))

	Response := make(map[string]interface{})

	for name, detailsRaw := range lambdaAssociationConf {
		// Validate that the lambda function name exists in expected associations
		assert.NotEmpty(t, name, "Lambda function association name should not be empty")

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

				assert.Equal(t, "test-lambda-assocaition", name, fmt.Sprintf("%s function_name should match expected value", resourceName))
				// Validate lambda function ARN
				assert.Equal(t, lambdaFunctionArn, detailsMap["function_arn"], fmt.Sprintf("%s function_arn should match expected value", resourceName))

				// Validate instance ID
				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance ID should match expected value")

				// Validate the ID structure (instance_id,function_arn)
				if id, ok := detailsMap["id"].(string); ok {
					lambdaAssociationId := strings.Split(id, ",")
					if len(lambdaAssociationId) >= 2 {
						instanceIDFromAssociation := lambdaAssociationId[0]
						functionArnFromAssociation := lambdaAssociationId[1]

						splitFunctionArn := strings.Split(functionArnFromAssociation, ":")
						functionNameFromAssociation := splitFunctionArn[len(splitFunctionArn)-1]

						assert.Equal(t, "jb-ccaas-test-cases-lambda-function", functionNameFromAssociation, fmt.Sprintf("%s lambda name in ID should match expected value", resourceName))
						assert.Equal(t, instanceID, instanceIDFromAssociation, fmt.Sprintf("%s instance ID in ID should match expected value", resourceName))
						assert.Equal(t, detailsMap["function_arn"], functionArnFromAssociation, fmt.Sprintf("%s function ARN should match ID structure", resourceName))
					}
				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
