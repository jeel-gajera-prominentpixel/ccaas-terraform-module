package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestConnectInstaceStorageConfig validates the instance storage config configuration created by Terraform.
// It performs comprehensive testing of the instance storage config output, including structure validation,
// resource type verification, and storage configuration checks.
//
// Parameters:
//   - t: Testing object for running tests and making assertions
//   - terraformOptions: Terraform options containing the configuration to test
//   - requiredTags: List of tags that must be present in the instance storage config (not directly used in this function)
//   - instanceARN: The ARN of the Amazon Connect instance (not directly used in this function)
//   - instanceID: The ID of the Amazon Connect instance
//   - BucketName: The name of the S3 bucket used for storage
//   - BucketArn: The ARN of the S3 bucket (not directly used in this function)
//   - KMSkeyArn: The ARN of the KMS key used for encryption
//   - KinesisStreamArn: The ARN of the Kinesis stream used for streaming
//
// The function performs the following validations:
//   - Verifies the instance storage config JSON is not empty
//   - Validates the structure and format of the JSON data
//   - Confirms the resource type matches the expected value for each storage config
//   - Checks S3, Kinesis Stream, and encryption configuration for each resource type
//   - Validates the instance ID and association ID in the storage config
//
// Returns:
//   - map[string]interface{}: A map containing the parsed instance storage configurations
func TestConnectInstaceStorageConfig(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string, BucketName string, BucketArn string, KMSkeyArn string, KinesisStreamArn string) map[string]interface{} {
	resourceName := "Instance Storage Config"
	initialFailed := t.Failed()
	InstanceStorageJson := terraform.OutputJson(t, terraformOptions, "instance_storage_configs")
	assert.NotEmpty(t, InstanceStorageJson, fmt.Sprintf("%s should not be empty", resourceName))

	// TestConnectInstaceStorageConfig validates the Amazon Connect instance storage configuration output from Terraform.
	// This test function performs the following checks for each supported storage config type:
	//   - Parses the "instance_storage_configs" Terraform output as JSON and ensures it is not empty.
	//   - Iterates over each storage config type (AGENT_EVENTS, CALL_RECORDINGS, CONTACT_TRACE_RECORDS, SCHEDULED_REPORTS)
	//     and validates that:
	//     - The resource type matches the expected value.
	//     - The storage type (S3 or KINESIS_STREAM) is correct for each config.
	//     - S3 configs have the correct bucket name, prefix, and KMS encryption settings.
	//     - Kinesis Stream configs have the correct stream ARN.
	//     - Unused config blocks (e.g., s3_config for KINESIS_STREAM) are empty.
	//     - The instance ID and association ID are correctly set and match the expected values.
	//     - The storage config ID is correctly structured and matches the instance ID, association ID, and config name.
	// The function returns a map containing the parsed and validated storage config details for further assertions.

	var InstanceStorageConf map[string]interface{}

	err := json.Unmarshal([]byte(InstanceStorageJson), &InstanceStorageConf)
	assert.NoError(t, err, fmt.Sprintf("Should be able to parse %s JSON", resourceName))

	Response := make(map[string]interface{})

	for name, detailsRaw := range InstanceStorageConf {
		assert.Contains(t, []string{"AGENT_EVENTS", "CALL_RECORDINGS", "CONTACT_TRACE_RECORDS", "SCHEDULED_REPORTS"}, name, "Instance Storage Configuyration should match one of the expected names")

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
				switch name {
				case "AGENT_EVENTS":
					assert.Equal(t, "AGENT_EVENTS", detailsMap["resource_type"], "resource_type should match expected value")

					if StorageConfigConfig, ok := detailsMap["storage_config"].(map[string]interface{}); ok {
						assert.Equal(t, "KINESIS_STREAM", StorageConfigConfig["storage_type"], fmt.Sprintf("%s channel should match expected value", resourceName))

						s3Config, _ := StorageConfigConfig["s3_config"].([]interface{})
						assert.Equal(t, 0, len(s3Config), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisVideoConfig, _ := StorageConfigConfig["kinesis_video_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisVideoConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisFirehoseConfig, _ := StorageConfigConfig["kinesis_firehose_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisFirehoseConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisStreamConfig, _ := StorageConfigConfig["kinesis_stream_config"].([]interface{})
						assert.True(t, len(KinesisStreamConfig) > 0, fmt.Sprintf("%s kinesis_stream_config should not be empty", resourceName))
						streamEntry, _ := KinesisStreamConfig[0].(map[string]interface{})
						assert.Equal(t, KinesisStreamArn, streamEntry["stream_arn"], fmt.Sprintf("%s stream_arn should match expected value", resourceName))
					}
				case "CALL_RECORDINGS":
					assert.Equal(t, "CALL_RECORDINGS", detailsMap["resource_type"], "resource_type should match expected value")

					if StorageConfigConfig, ok := detailsMap["storage_config"].(map[string]interface{}); ok {
						assert.Equal(t, "S3", StorageConfigConfig["storage_type"], fmt.Sprintf("%s channel should match expected value", resourceName))

						s3Config, _ := StorageConfigConfig["s3_config"].([]interface{})
						assert.True(t, len(s3Config) > 0, fmt.Sprintf("%s s3_config should not be empty", resourceName))
						s3Entry, _ := s3Config[0].(map[string]interface{})
						assert.Equal(t, BucketName, s3Entry["bucket_name"], fmt.Sprintf("%s bucket_name should match expected value", resourceName))
						assert.Equal(t, "call_recordings", s3Entry["bucket_prefix"], fmt.Sprintf("%s bucket_name should match expected value", resourceName))

						EncryptionConfig, _ := s3Entry["encryption_config"].([]interface{})
						assert.True(t, len(EncryptionConfig) > 0, fmt.Sprintf("%s encryption_config should not be empty", resourceName))
						encryptionEntry, _ := EncryptionConfig[0].(map[string]interface{})
						assert.Equal(t, "KMS", encryptionEntry["encryption_type"], fmt.Sprintf("%s encryption_type should match expected value", resourceName))
						assert.Equal(t, KMSkeyArn, encryptionEntry["key_id"], fmt.Sprintf("%s key_id should match expected value", resourceName))

						KinesisVideoConfig, _ := StorageConfigConfig["kinesis_video_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisVideoConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisFirehoseConfig, _ := StorageConfigConfig["kinesis_firehose_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisFirehoseConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisStreamConfig, _ := StorageConfigConfig["kinesis_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisStreamConfig), fmt.Sprintf("%s kinesis_stream_config should be empty", resourceName))

					}
				case "CONTACT_TRACE_RECORDS":
					assert.Equal(t, "CONTACT_TRACE_RECORDS", detailsMap["resource_type"], "resource_type should match expected value")

					if StorageConfigConfig, ok := detailsMap["storage_config"].(map[string]interface{}); ok {
						assert.Equal(t, "KINESIS_STREAM", StorageConfigConfig["storage_type"], fmt.Sprintf("%s channel should match expected value", resourceName))

						s3Config, _ := StorageConfigConfig["s3_config"].([]interface{})
						assert.Equal(t, 0, len(s3Config), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisVideoConfig, _ := StorageConfigConfig["kinesis_video_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisVideoConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisFirehoseConfig, _ := StorageConfigConfig["kinesis_firehose_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisFirehoseConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisStreamConfig, _ := StorageConfigConfig["kinesis_stream_config"].([]interface{})
						assert.True(t, len(KinesisStreamConfig) > 0, fmt.Sprintf("%s kinesis_stream_config should not be empty", resourceName))
						streamEntry, _ := KinesisStreamConfig[0].(map[string]interface{})
						assert.Equal(t, KinesisStreamArn, streamEntry["stream_arn"], fmt.Sprintf("%s stream_arn should match expected value", resourceName))
					}
				case "SCHEDULED_REPORTS":
					assert.Equal(t, "SCHEDULED_REPORTS", detailsMap["resource_type"], "resource_type should match expected value")

					if StorageConfigConfig, ok := detailsMap["storage_config"].(map[string]interface{}); ok {
						assert.Equal(t, "S3", StorageConfigConfig["storage_type"], fmt.Sprintf("%s channel should match expected value", resourceName))

						s3Config, _ := StorageConfigConfig["s3_config"].([]interface{})
						assert.True(t, len(s3Config) > 0, fmt.Sprintf("%s s3_config should not be empty", resourceName))
						s3Entry, _ := s3Config[0].(map[string]interface{})
						assert.Equal(t, BucketName, s3Entry["bucket_name"], fmt.Sprintf("%s bucket_name should match expected value", resourceName))
						assert.Equal(t, "scheduled_reports", s3Entry["bucket_prefix"], fmt.Sprintf("%s bucket_name should match expected value", resourceName))

						EncryptionConfig, _ := s3Entry["encryption_config"].([]interface{})
						assert.True(t, len(EncryptionConfig) > 0, fmt.Sprintf("%s encryption_config should not be empty", resourceName))
						encryptionEntry, _ := EncryptionConfig[0].(map[string]interface{})
						assert.Equal(t, "KMS", encryptionEntry["encryption_type"], fmt.Sprintf("%s encryption_type should match expected value", resourceName))
						assert.Equal(t, KMSkeyArn, encryptionEntry["key_id"], fmt.Sprintf("%s key_id should match expected value", resourceName))

						KinesisVideoConfig, _ := StorageConfigConfig["kinesis_video_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisVideoConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisFirehoseConfig, _ := StorageConfigConfig["kinesis_firehose_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisFirehoseConfig), fmt.Sprintf("%s s3_config should be empty", resourceName))

						KinesisStreamConfig, _ := StorageConfigConfig["kinesis_stream_config"].([]interface{})
						assert.Equal(t, 0, len(KinesisStreamConfig), fmt.Sprintf("%s kinesis_stream_config should be empty", resourceName))

					}
				default:
					assert.Fail(t, fmt.Sprintf("Unexpected resource_type: %s", name))
				}

				assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match expected value")

				InstanceStorageConfigId := strings.Split(detailsMap["id"].(string), ":")
				if len(InstanceStorageConfigId) >= 2 {
					InstanceIDFromStorageConfig := InstanceStorageConfigId[0]
					AssociationIDFromStorageConfig := InstanceStorageConfigId[1]
					StorageConfigNameFromStorageConfig := InstanceStorageConfigId[2]

					assert.Equal(t, instanceID, InstanceIDFromStorageConfig, fmt.Sprintf("%s instance id of id should be empty", resourceName))
					assert.Equal(t, detailsMap["association_id"], AssociationIDFromStorageConfig, fmt.Sprintf("%s association_id should be empty", resourceName))
					assert.Equal(t, name, StorageConfigNameFromStorageConfig, fmt.Sprintf("%s association_id should be empty", resourceName))

				}
			}
		}
	}
	Response["Success"] = !t.Failed() || initialFailed == t.Failed()
	return Response
}
