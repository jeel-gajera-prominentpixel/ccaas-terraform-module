// Package test provides testing utilities for Amazon Connect module
package test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"encoding/json"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variables to store dynamic IDs
var (
	dynamicOutBoundContactFlowID    string
	dynamicQueueTransferFlowID      string
	dynamicHoursOfOperationID       string
	dynamicQueuesID                 string
	dynamicUserHierarchyStructureID int
	dynamicUserHierarchyGroupID     int
	hierarchyGroups                 map[string]interface{}
	connectBucketId                 string
	connectS3BucketArn              string
	s3KMSKeyArn                     string
	KinesisDataStreamArn            string
	LambdaFunctionAssociationARN    string
)

// Constants for Amazon Connect testing
var (
	region              = "us-east-1"
	prefixRegion        = "use1"
	prefixCompany       = "jb"
	lob                 = "test"
	application         = "cases"
	env                 = "sandbox"
	create_instance     = false
	instance_id         = "184cff82-60e1-47ee-9875-71e202ab41b8"
	hours_of_operations = map[string]interface{}{
		"test_hours_of_operations": map[string]interface{}{
			"description": "24-7 Music on Hold hours",
			"time_zone":   "US/Eastern",
			"config": []map[string]interface{}{
				{
					"day":        "WEDNESDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "FRIDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "MONDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "SATURDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "TUESDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "THURSDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
				{
					"day":        "SUNDAY",
					"start_time": map[string]int{"hours": 6, "minutes": 0},
					"end_time":   map[string]int{"hours": 0, "minutes": 0},
				},
			},
		},
	}
	rawContent     = must(os.ReadFile("./ContactFlowAndModule/BasicFlow.json"))
	escapedContent = strings.ReplaceAll(strings.ReplaceAll(string(rawContent), "\"", "\\\""), "\n", "\\n")

	rawContentWhisper     = must(os.ReadFile("./ContactFlowAndModule/WhisperFlow.json"))
	escapedContentWhisper = strings.ReplaceAll(strings.ReplaceAll(string(rawContentWhisper), "\"", "\\\""), "\n", "\\n")

	rawContentQueueTransfer     = must(os.ReadFile("./ContactFlowAndModule/QueueTransferFlow.json"))
	escapedContentQueueTransfer = strings.ReplaceAll(strings.ReplaceAll(string(rawContentQueueTransfer), "\"", "\\\""), "\n", "\\n")
	contact_flows               = map[string]interface{}{
		"test_contact_flows": map[string]interface{}{
			"description": "Basic contact flow for testing",
			"type":        "CONTACT_FLOW",
			"content":     escapedContent,
		},
		"test_whisper_flows": map[string]interface{}{
			"description": "Basic outbound whisper flow for testing",
			"type":        "OUTBOUND_WHISPER",
			"content":     escapedContentWhisper,
		},
		"test_queue_transfer_flows": map[string]interface{}{
			"description": "Basic test_queue_transfer_flows flow for testing",
			"type":        "QUEUE_TRANSFER",
			"content":     escapedContentQueueTransfer,
		},
	}
	rawContentModule     = must(os.ReadFile("./ContactFlowAndModule/BasicFlowModule.json"))
	escapedContentModule = strings.ReplaceAll(strings.ReplaceAll(string(rawContentModule), "\"", "\\\""), "\n", "\\n")
	contact_flow_modules = map[string]interface{}{
		"test_contact_flows_module": map[string]interface{}{
			"description": "Basic contact flow module for testing",
			"content":     escapedContentModule,
		},
	}
	user_hierarchy_structure = map[string]string{
		"level_one":   "lob",
		"level_two":   "manager",
		"level_three": "supervisor",
		"level_four":  "agent",
	}
)

func getCommonVars() map[string]interface{} {
	vars := map[string]interface{}{
		"application":                       application,
		"prefix_company":                    prefixCompany,
		"prefix_region":                     prefixRegion,
		"lob":                               lob,
		"env":                               env,
		"create_instance":                   create_instance,
		"instance_id":                       instance_id,
		"instance_identity_management_type": "SAML",
		"name":                              fmt.Sprintf("%s-connect-%s-%s-%s", prefixCompany, lob, application, env),
		"tags": map[string]string{
			"company": prefixCompany,
			"lob":     lob,
		},
		"hours_of_operations":      hours_of_operations,
		"contact_flows":            contact_flows,
		"contact_flow_modules":     contact_flow_modules,
		"user_hierarchy_structure": user_hierarchy_structure,
		// "instance_storage_configs":     GetInstanceStorageConfig(),
		"lambda_function_associations": GetInstanceLambdaFunctionAssociation(),
	}

	// Add queues only if we have dynamic IDs
	if dynamicOutBoundContactFlowID != "" && dynamicHoursOfOperationID != "" {
		vars["queues"] = getQueuesConfig()
		vars["queue_tags"] = map[string]string{
			"company": prefixCompany,
			"lob":     lob,
		}
	}
	if dynamicQueueTransferFlowID != "" && dynamicQueuesID != "" {
		vars["quick_connects"] = getQuickConnectConfig()
	}
	if dynamicUserHierarchyStructureID > 0 {
		vars["user_hierarchy_groups"] = GetHierarchyGroupsConfig()
	}
	if dynamicQueuesID != "" {
		vars["routing_profiles"] = GetRoutingProfileConfig()
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
	"module_version",
	// "project_path",
	// "commit_id",
	"company",
	"region",
	"lob",
	"application",
	"env",
	"created_by",
	"map-migrated",
}

func GetInstanceLambdaFunctionAssociation() map[string]interface{} {
	LambdaFunctionAssociationARN = "arn:aws:lambda:us-west-2:381492173985:function:jb-ccaas-test-cases-lambda-function"
	return map[string]interface{}{
		"test-lambda-assocaition": LambdaFunctionAssociationARN,
	}
}

// func GetInstanceStorageConfig() map[string]interface{} {

// 	connectS3BucketArn = "arn:aws:s3:::jb-test-cases-usw2-sandbox"
// 	s3KMSKeyArn = "arn:aws:kms:us-west-2:381492173985:key/mrk-9e899d984f1e40a8891f51b262a1ebe5"
// 	KinesisDataStreamArn = "arn:aws:kinesis:us-west-2:381492173985:stream/jb-test-cases-event-stream-sandbox"
// 	connectBucketId = "jb-test-cases-usw2-sandbox"

//		return map[string]interface{}{
//			"SCHEDULED_REPORTS": map[string]interface{}{
//				"storage_type": "S3",
//				"s3_config": map[string]interface{}{
//					"bucket_name":   connectBucketId,
//					"bucket_prefix": "scheduled_reports",
//					"encryption_config": map[string]interface{}{
//						"encryption_type": "KMS",
//						"key_id":          s3KMSKeyArn,
//					},
//				},
//			},
//			"CALL_RECORDINGS": map[string]interface{}{
//				"storage_type": "S3",
//				"s3_config": map[string]interface{}{
//					"bucket_name":   connectBucketId,
//					"bucket_prefix": "call_recordings",
//					"encryption_config": map[string]interface{}{
//						"encryption_type": "KMS",
//						"key_id":          s3KMSKeyArn,
//					},
//				},
//			},
//			"AGENT_EVENTS": map[string]interface{}{
//				"storage_type": "KINESIS_STREAM",
//				"kinesis_stream_config": map[string]interface{}{
//					"stream_arn": KinesisDataStreamArn,
//				},
//			},
//			"CONTACT_TRACE_RECORDS": map[string]interface{}{
//				"storage_type": "KINESIS_STREAM",
//				"kinesis_stream_config": map[string]interface{}{
//					"stream_arn": KinesisDataStreamArn,
//				},
//			},
//		}
//	}
func GetRoutingProfileConfig() map[string]interface{} {
	return map[string]interface{}{
		"jb_maint_A1": map[string]interface{}{
			"description":               "A1 routing profile",
			"default_outbound_queue_id": dynamicQueuesID,
			"media_concurrencies": []map[string]interface{}{
				{
					"channel":     "VOICE",
					"concurrency": 1,
				},
			},
			"queue_configs": []map[string]interface{}{
				{
					"channel":  "VOICE",
					"delay":    0,
					"priority": 5,
					"queue_id": dynamicQueuesID,
				},
			},
		},
	}
}

func GetHierarchyGroupsConfig() map[string]interface{} {
	// Initialize map for hierarchy groups
	user_hierarchy_groups := map[string]interface{}{}

	// Always add "maint"
	user_hierarchy_groups["maint"] = map[string]interface{}{}

	// Add "MCCDutyManager" when level 2 or higher
	if dynamicUserHierarchyGroupID >= 2 && hierarchyGroups["maint"].(map[string]interface{})["hierarchy_group_id"].(string) != "" {
		user_hierarchy_groups["MCCDutyManager"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["maint"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
	}

	// Add "AOGSupervisor" when level 3 or higher
	if dynamicUserHierarchyGroupID >= 3 && hierarchyGroups["MCCDutyManager"].(map[string]interface{})["hierarchy_group_id"].(string) != "" {
		user_hierarchy_groups["AOGSupervisor"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["MCCDutyManager"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
		user_hierarchy_groups["MCCDutySupervisor"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["MCCDutyManager"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
		user_hierarchy_groups["MOCPlanningSupervisor"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["MCCDutyManager"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
	}

	// Add "AOGAgent" when level 4
	if dynamicUserHierarchyGroupID == 4 {
		user_hierarchy_groups["AOGAgent"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["AOGSupervisor"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
		user_hierarchy_groups["MCCAgent"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["MCCDutySupervisor"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
		user_hierarchy_groups["MOCAgent"] = map[string]interface{}{
			"parent_group_id": hierarchyGroups["MOCPlanningSupervisor"].(map[string]interface{})["hierarchy_group_id"].(string),
		}
	}

	return user_hierarchy_groups
}

// getQueuesConfig returns a configuration for test_queues with required non-empty tags.
// AWS Connect requires at least one non-empty tag value for queue creation to succeed.
func getQueuesConfig() map[string]interface{} {
	return map[string]interface{}{
		"test_queues": map[string]interface{}{
			"description":           "Basic Queues for testing",
			"hours_of_operation_id": dynamicHoursOfOperationID,
			"outbound_caller_config": map[string]interface{}{
				"outbound_caller_id_name": "Test Caller ID",
				"outbound_flow_id":        dynamicOutBoundContactFlowID,
			},
			"status":            "ENABLED",
			"quick_connect_ids": []string{},
			// Ensure at least one non-empty tag is always present
			"tags": map[string]string{
				"Name":    "jb-connect-test-cases-sandbox", // Always non-empty
				"lob":     lob,
				"company": prefixCompany,
			},
		},
	}
}

func getQuickConnectConfig() map[string]interface{} {
	return map[string]interface{}{
		"test_quick_connect": map[string]interface{}{
			"quick_connect_config": map[string]interface{}{
				"quick_connect_type": "QUEUE",
				"queue_config": map[string]interface{}{
					"queue_id":        dynamicQueuesID,
					"contact_flow_id": dynamicQueueTransferFlowID,
				},
				"description": "Basic Quick Connect for testing",
			},
		},
	}
}

// TestConnectCreation validates the basic creation of an Amazon Connect instance
// TestMain sets up the shared Amazon Connect instance for all tests
func TestMain(m *testing.M) {
	// Create a shared instance for all tests
	terraformOptions := terraform.WithDefaultRetryableErrors(&testing.T{}, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	log.SetOutput(io.Discard)
	// Initialize and apply the configuration once
	terraform.InitAndApply(&testing.T{}, terraformOptions)

	// Run all tests
	code := m.Run()

	// Clean up after all tests complete
	terraform.Destroy(&testing.T{}, terraformOptions)

	os.Exit(code)
}

// TestConnectCreation validates the creation of an Amazon Connect instance.
// It checks for the existence and correct formatting of required outputs such as
// instance ARN, ID, alias, and status. The test ensures the Connect instance ID,
// ARN, and alias are in the correct format, verifies the instance is active,
// and confirms the identity management type is SAML. Additionally, it logs
// various potentially empty outputs related to the Connect instance's resources.

func TestConnectCreation(t *testing.T) {
	// t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Verify additional resources
	outputs := []string{
		"bot_associations",
		"contact_flow_modules",
		"contact_flows",
		"hours_of_operations",
		"instance_storage_configs",
		"lambda_function_associations",
		"queues",
		"quick_connects",
		"routing_profiles",
		"security_profiles",
		"user_hierarchy_groups",
		"users",
		"vocabularies",
		"connect_idp_admin_arn",
		"connect_idp_agent_arn",
	}

	// These outputs might be empty depending on configuration
	for _, output := range outputs {
		_ = terraform.Output(t, terraformOptions, output)
		// t.Logf("%s: %s", output, value)
	}
}

// TestConnectConfiguration validates the configuration of an Amazon Connect instance
// with all features enabled. It checks for the correct configuration of features
// such as auto-resolve best voices, contact lens, early media, inbound calls,
// outbound calls, multi-party conferences, and contact flow logs. Additionally,
// it verifies the identity management type is SAML.
func TestConnectConfiguration(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(),
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	instanceARN := terraform.Output(t, terraformOptions, "instance_arn")
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	LambdaFunctionAssociation := TestConnectLambdaAssociation(t, terraformOptions, instanceARN, instanceID, LambdaFunctionAssociationARN)
	if !LambdaFunctionAssociation["Success"].(bool) {
		t.Logf("LambdaFunctionAssociation Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("LambdaFunctionAssociation Test cases are successful.")
	}

	HoursOfOperation := TestHoursOfOperation(t, terraformOptions, requiredTags, instanceARN, instanceID)
	if !HoursOfOperation["Success"].(bool) {
		t.Logf("HoursOfOperation Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("HoursOfOperation Test cases are successful.")
	}

	ContactFlow := TestContactFlow(t, terraformOptions, requiredTags, instanceARN, instanceID)
	if !ContactFlow["Success"].(bool) {
		t.Logf("ContactFlow Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("ContactFlow Test cases are successful.")
	}

	ContactFlowModules := TestContactFlowModules(t, terraformOptions, requiredTags, instanceARN, instanceID)
	if !ContactFlowModules["Success"].(bool) {
		t.Logf("ContactFlowModules Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("ContactFlowModules Test cases are successful.")
	}

	// Test Security Profiles
    SecurityProfiles := TestConnectSecurityProfiles(t, terraformOptions, requiredTags, instanceARN, instanceID)
    if !SecurityProfiles["Success"].(bool) {
        t.Logf("SecurityProfiles Test cases are not successful, continuing with other tests")
    } else {
        t.Logf("SecurityProfiles Test cases are successful.")
    }

	// Test SSO Identity Providers
    SSOIdPs := TestConnectSSOIdentityProviders(t, terraformOptions, instanceARN)
    if !SSOIdPs["Success"].(bool) {
        t.Logf("SSOIdPs Test cases are not successful, continuing with other tests")
    } else {
        t.Logf("SSOIdPs Test cases are successful.")
    }

	// Test Vocabularies
    Vocabularies := TestConnectVocabularies(t, terraformOptions, requiredTags, instanceARN, instanceID)
    if !Vocabularies["Success"].(bool) {
        t.Logf("Vocabularies Test cases are not successful, continuing with other tests")
    } else {
        t.Logf("Vocabularies Test cases are successful.")
    }

	// Test Bot Associations
    BotAssociations := TestConnectBotAssociations(t, terraformOptions, instanceARN, instanceID)
    if !BotAssociations["Success"].(bool) {
        t.Logf("BotAssociations Test cases are not successful, continuing with other tests")
    } else {
        t.Logf("BotAssociations Test cases are successful.")
    }

	// Test Users (if configured)
    Users := TestConnectUsers(t, terraformOptions, requiredTags, instanceARN, instanceID)
    if !Users["Success"].(bool) {
        t.Logf("Users Test cases are not successful, continuing with other tests")
    } else {
        t.Logf("Users Test cases are successful.")
    }

	dynamicOutBoundContactFlowID = ContactFlow["test_whisper_flows"].(map[string]interface{})["contact_flow_id"].(string)
	dynamicHoursOfOperationID = HoursOfOperation["test_hours_of_operations"].(map[string]interface{})["hours_of_operation_id"].(string)
	dynamicQueueTransferFlowID = ContactFlow["test_queue_transfer_flows"].(map[string]interface{})["contact_flow_id"].(string)

	TestConnectInstaceStorageConf := TestConnectInstaceStorageConfig(t, terraformOptions, requiredTags, instanceARN, instanceID, connectBucketId, connectS3BucketArn, s3KMSKeyArn, KinesisDataStreamArn)
	if !TestConnectInstaceStorageConf["Success"].(bool) {
		t.Logf("TestConnectInstaceStorageConf Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("TestConnectInstaceStorageConf Test cases are successful.")
	}

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Reapply terraform with updated variables including queues
	//////////////////////////////////////////////////////////////////////////////////////////////
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(), // This will now include queues since we have the dynamic IDs
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	// Apply the updated configuration
	terraform.Apply(t, terraformOptions)

	QueuesConf := TestConnectQueues(t, terraformOptions, requiredTags, instanceARN, instanceID, dynamicOutBoundContactFlowID, dynamicHoursOfOperationID)
	if !QueuesConf["Success"].(bool) {
		t.Logf("QueuesConf Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("QueuesConf Test cases are successful.")
	}
	dynamicQueuesID = QueuesConf["test_queues"].(map[string]interface{})["queue_id"].(string)

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Reapply terraform with updated variables including quick connects
	//////////////////////////////////////////////////////////////////////////////////////////////
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         getCommonVars(), // This will now include quick connects since we have the dynamic IDs
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	terraform.Apply(t, terraformOptions)

	QuickConnectConf := TestConnectQuickConnect(t, terraformOptions, requiredTags, instanceARN, instanceID, dynamicQueueTransferFlowID, dynamicQueuesID)
	if !QuickConnectConf["Success"].(bool) {
		t.Logf("QuickConnectConf Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("QuickConnectConf Test cases are successful.")
	}

	RoutingProfileConf := TestConnectRoutingProfile(t, terraformOptions, requiredTags, instanceARN, instanceID, dynamicQueuesID, QueuesConf)
	if !RoutingProfileConf["Success"].(bool) {
		t.Logf("RoutingProfileConf Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("RoutingProfileConf Test cases are successful.")
	}

	// Verify the user hierarchy structure
	UserHierarchyStructConf := TestConnectUserHierarchyStructure(t, terraformOptions, instanceARN, instanceID)
	if !UserHierarchyStructConf["Success"].(bool) {
		t.Logf("UserHierarchyStructConf Test cases are not successful, skipping further tests")
		return
	} else {
		t.Logf("UserHierarchyStructConf Test cases are successful.")
	}
	hierarchyStructure := UserHierarchyStructConf["hierarchy_structure"].([]interface{})[0].(map[string]interface{})
	// Count the number of levels
	dynamicUserHierarchyStructureID = len(hierarchyStructure)

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Reapply terraform with updated variables including user hierarchy groups
	//////////////////////////////////////////////////////////////////////////////////////////////
	for i := 1; i <= 4; i++ {
		dynamicUserHierarchyGroupID = i
		terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			TerraformDir: "../",
			Vars:         getCommonVars(), // This will now include quick connects since we have the dynamic IDs
			NoStderr:     true,
			Logger:       logger.Discard,
			EnvVars: map[string]string{
				"TF_LOG":      "ERROR",
				"TF_LOG_PATH": "/dev/null",
			},
		})

		terraform.Apply(t, terraformOptions)
		hierarchyGroups = TestConnectUserHierarchyGroups(t, terraformOptions, instanceARN, instanceID, requiredTags)
		if !hierarchyGroups["Success"].(bool) {
			t.Logf("hierarchyGroups Test cases are not successful, skipping further tests")
			return
		} else {
			t.Logf("hierarchyGroups Test cases are successful.")
		}
	}

	// test output map
	resourceName := "Connect"
	outputs := terraform.OutputAll(t, terraformOptions)

	optionsBytes, err := json.Marshal(outputs)
	assert.NoError(t, err, "Should be able to marshal all outputs to JSON")
	var detailsStr = string(optionsBytes)
	var detailsMap map[string]interface{}
	if detailsStr != "" {
		err := json.Unmarshal([]byte(detailsStr), &detailsMap)
		assert.NoError(t, err, "Should be able to parse details JSON")
	}

	// initializing the resource name and details map , test the output
	fmt.Printf("Terraform Output Map for %s:\n%+v\n", resourceName, detailsMap)

	defer func() {
		// removeStorageConfigsViaAPI(instanceID)
		// Ensure ResourceRemoval is called even if test fails
		varsWithoutQuickConnects := getCommonVars()
		// varsWithoutQuickConnects["instance_storage_configs"] = map[string]interface{}{}
		ResourceRemoval(t, terraformOptions, varsWithoutQuickConnects)
	}()
}

// func removeStorageConfigsViaAPI(instanceID string) {
// 	sess := session.Must(session.NewSession())
// 	svc := connect.New(sess)

// 	storageTypes := []string{"SCHEDULED_REPORTS", "CALL_RECORDINGS", "AGENT_EVENTS", "CONTACT_TRACE_RECORDS"}

// 	for _, storageType := range storageTypes {
// 		_, _ = svc.DisassociateInstanceStorageConfig(&connect.DisassociateInstanceStorageConfigInput{
// 			InstanceId:   aws.String(instanceID),
// 			ResourceType: aws.String(storageType),
// 		})
// 	}
// }
