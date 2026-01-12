// Package test provides testing utilities for DynamoDB module
package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Constants for Amazon Connect testing
var (
	region        = "us-east-1"
	prefixRegion  = "use1"
	prefixCompany = "jb"
	lob           = "test"
	application   = "cases"
	env           = "sandbox"
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":    application,
		"prefix_company": prefixCompany,
		"prefix_region":  prefixRegion,
		"lob":            lob,
		"env":            env,
	}
	return vars
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Required tags for Amazon Connect resources
var requiredTags = []string{
	"module_project_path",
	// "module_version",
	// "project_path",
	"commit_id",
	"company",
	"region",
	"lob",
	"application",
	"env",
	"created_by",
	"map-migrated",
}

// Shared variables to store between tests
var (
	botId      string
	botVersion string
)

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)

	code := m.Run()
	os.Exit(code)
}

func TestCreateLexbot(t *testing.T) {
	// First function: Create the Lexbot
	vars := getCommonVars()
	// Only create bot, not version or alias
	vars["create_lexbot"] = true
	vars["create_lexbot_version"] = false
	vars["create_lexbot_alias"] = false
	vars["name"] = "jb-test-cases-lexbot"
	vars["auto_build_bot_locales"] = true
	vars["role_arn"] = "arn:aws:iam::381492173985:role/jb-kinesis-test-cases-sandbox"
	vars["data_privacy"] = map[string]interface{}{
		"child_directed": false,
	}
	vars["idle_session_ttl_in_seconds"] = 300
	vars["bot_file_s3_location"] = map[string]interface{}{
		"s3_bucket":     "jb-test-cases-usw2-sandbox",
		"s3_object_key": "sandbox/lexbot/jb-lex-crew-callback-2-RTPNGFEMIY-LexJson.zip",
	}
	vars["bot_tags"] = []map[string]interface{}{
		{
			"key":   "lob",
			"value": lob,
		},
	}
	vars["description"] = "Test cases for lexbot module"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         vars,
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Initialize and apply terraform for bot creation
	terraform.InitAndApply(t, terraformOptions)

	// Get the outputs
	outputs := terraform.OutputAll(t, terraformOptions)

	// Extract and verify bot ID
	var exists bool
	botId, exists = outputs["bot_id"].(string)
	assert.True(t, exists, "Bot ID should exist in outputs")
	assert.NotEmpty(t, botId, "Bot ID should not be empty")

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}
	// Don't clean up after bot creation, it's needed for version creation
	terraformOptionsRemoveVersion := CreateBotVersion(t, terraformOptions, botId)
	terraformOptionsRemoveAlias := CreateBotAlias(t, terraformOptions, botId, botVersion)

	terraform.RunTerraformCommand(t, terraformOptionsRemoveAlias, terraform.FormatArgs(terraformOptionsRemoveAlias, "destroy", "-auto-approve", "-target=awscc_lex_bot_alias.this")...)
	terraform.RunTerraformCommand(t, terraformOptionsRemoveVersion, terraform.FormatArgs(terraformOptionsRemoveVersion, "destroy", "-auto-approve", "-target=aws_lexv2models_bot_version.this")...)
	defer terraform.Destroy(t, terraformOptions)
}

func CreateBotVersion(t *testing.T, terraformOptions *terraform.Options, botId string) *terraform.Options {

	// Verify we have the bot ID from previous test
	assert.NotEmpty(t, botId, "Bot ID should be available from previous test")
	vars := getCommonVars()

	// Set up variables for version creation
	// Don't create or destroy the bot, just create version
	vars["create_lexbot_version"] = true
	vars["bot_id"] = botId
	vars["version_description"] = "Test cases for lexbot version"
	vars["locale_specification"] = map[string]interface{}{
		"en_US": map[string]interface{}{
			"source_bot_version": "DRAFT",
		},
	}

	terraformOptionsVersions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         vars,
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Initialize and apply terraform for version creation
	terraform.RunTerraformCommand(t, terraformOptionsVersions, terraform.FormatArgs(terraformOptionsVersions, "apply", "-auto-approve", "-target=aws_lexv2models_bot_version.this")...)

	// Get the outputs
	outputs := terraform.OutputAll(t, terraformOptionsVersions)

	// Extract and verify bot version
	var exists bool
	botVersion, exists = outputs["bot_version"].(string)
	assert.True(t, exists, "Bot version should exist in outputs")
	assert.NotEmpty(t, botVersion, "Bot version should not be empty")

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Testcases bot version output properties
	// Test bot version output properties
	outputBotId, exists := outputs["bot_id"].(string)
	assert.True(t, exists, "Bot ID should exist in outputs")
	assert.Equal(t, botId, outputBotId, "Bot ID should match expected value")

	outputBotName, exists := outputs["bot_name"].(string)
	assert.True(t, exists, "Bot name should exist in outputs")
	assert.Equal(t, "jb-test-cases-lexbot", outputBotName, "Bot name should match expected value")

	outputBotVersion, exists := outputs["bot_version"].(string)
	assert.True(t, exists, "Bot version should exist in outputs")
	assert.Equal(t, botVersion, outputBotVersion, "Bot version should match expected value")

	// Verify empty fields as per the output
	emptyFields := []string{"bot_alias_arn", "bot_alias_id", "bot_alias_name"}
	for _, field := range emptyFields {
		value, exists := outputs[field].(string)
		assert.True(t, exists, fmt.Sprintf("%s should exist in outputs", field))
		assert.Empty(t, value, fmt.Sprintf("%s should be empty at version creation stage", field))
	}

	fmt.Printf("BotVersion: %+v\n", botVersion)
	return terraformOptionsVersions
}

func CreateBotAlias(t *testing.T, terraformOptions *terraform.Options, botId string, botVersion string) *terraform.Options {
	// Verify we have both bot ID and version from previous tests
	assert.NotEmpty(t, botId, "Bot ID should be available from previous test")
	assert.NotEmpty(t, botVersion, "Bot version should be available from previous test")

	// Set up variables for alias creation
	vars := getCommonVars()
	// Don't create/destroy bot or version, just create alias
	vars["create_lexbot_alias"] = true
	vars["bot_id"] = botId
	vars["bot_version"] = botVersion
	vars["bot_alias_name"] = "test-alias"
	vars["sentiment_analysis_settings"] = map[string]interface{}{
		"detect_sentiment": false,
	}
	vars["bot_alias_locale_settings"] = []map[string]interface{}{
		{
			"locale_id": "en_US",
			"bot_alias_locale_setting": map[string]interface{}{
				"enabled":                 true,
				"code_hook_specification": nil,
			},
		},
	}
	vars["conversation_log_settings"] = map[string]interface{}{
		"text_log_settings": []map[string]interface{}{
			{
				"enabled": true,
				"destination": map[string]interface{}{
					"cloudwatch": map[string]interface{}{
						"cloudwatch_log_group_arn": "arn:aws:logs:us-west-2:381492173985:log-group:/aws/connect/jb-connect-test-cases-sandbox:*",
						"log_prefix":               "/aws/connect/",
					},
				},
			},
		},
		"audio_log_settings": []map[string]interface{}{
			{
				"enabled": false,
				"destination": map[string]interface{}{
					"s3_bucket": map[string]interface{}{
						"s3_bucket_arn": "arn:aws:s3:::jb-test-cases-usw2-sandbox",
						"log_prefix":    "/lex-audio-logs/",
					},
				},
			},
		},
	}

	terraformOptionsAlias := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         vars,
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Initialize and apply terraform for alias creation
	terraform.RunTerraformCommand(t, terraformOptionsAlias, terraform.FormatArgs(terraformOptionsAlias, "apply", "-auto-approve", "-target=awscc_lex_bot_alias.this")...)

	// Get the outputs and verify alias creation
	outputs := terraform.OutputAll(t, terraformOptionsAlias)
	fmt.Printf("OUTPUT: %+v\n", outputs)

	// Extract bot alias ID from the output
	rawAliasId, exists := outputs["bot_alias_id"].(string)
	assert.True(t, exists, "Bot alias ID should exist in outputs")
	assert.NotEmpty(t, rawAliasId, "Bot alias ID should not be empty")
	// Split the alias ID and take the first part
	botAliasId := strings.Split(rawAliasId, "|")[0]
	assert.NotEmpty(t, botAliasId, "Bot alias ID part should not be empty")

	botAlias, exists := outputs["bot_alias_name"].(string)
	assert.True(t, exists, "Bot alias should exist in outputs")
	assert.NotEmpty(t, botAlias, "Bot alias should not be empty")
	fmt.Printf("BotAlias: %+v\n", botAlias)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// Test bot alias output properties
	aliasArn, exists := outputs["bot_alias_arn"].(string)
	assert.True(t, exists, "Bot alias ARN should exist in outputs")
	assert.Equal(t, fmt.Sprintf("arn:aws:lex:us-west-2:381492173985:bot-alias/%s/%s", botId, botAliasId), aliasArn, "Bot alias ARN should match expected value")

	aliasId, exists := outputs["bot_alias_id"].(string)
	assert.True(t, exists, "Bot alias ID should exist in outputs")
	assert.Equal(t, fmt.Sprintf("%s|%s", botAliasId, botId), aliasId, "Bot alias ID should match expected value")

	aliasName, exists := outputs["bot_alias_name"].(string)
	assert.True(t, exists, "Bot alias name should exist in outputs")
	assert.Equal(t, "test-alias", aliasName, "Bot alias name should match expected value")

	outputBotId, exists := outputs["bot_id"].(string)
	assert.True(t, exists, "Bot ID should exist in outputs")
	assert.Equal(t, botId, outputBotId, "Bot ID should match expected value")

	outputBotName, exists := outputs["bot_name"].(string)
	assert.True(t, exists, "Bot name should exist in outputs")
	assert.Equal(t, "jb-test-cases-lexbot", outputBotName, "Bot name should match expected value")

	outputBotVersion, exists := outputs["bot_version"].(string)
	assert.True(t, exists, "Bot version should exist in outputs")
	assert.Equal(t, botVersion, outputBotVersion, "Bot version should match expected value")

	return terraformOptionsAlias
}
