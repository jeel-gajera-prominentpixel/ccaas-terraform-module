# AWS Connect Test Suite Documentation

## Table of Contents
1. [Overview of AWS Connect Test Suite](#overview-of-aws-connect-test-suite)
2. [Directory Structure](#directory-structure)
3. [Main Test Components](#main-test-components)
4. [Execution Flow and Dynamic Output Handling](#execution-flow-and-dynamic-output-handling)
5. [Individual Test Function Patterns](#individual-test-function-patterns)
6. [Cleanup Strategy](#cleanup-strategy)
7. [Technical Architecture](#technical-architecture)

---

## Overview of AWS Connect Test Suite

This is a **comprehensive integration test suite** for Amazon Connect using the **Terratest framework**. The test suite validates your entire AWS Connect Terraform module with real AWS resources, ensuring all components integrate properly and dependencies are handled correctly.

### Key Features
- **Production-grade test suite** for AWS Connect Terraform modules
- **Sequential testing** with dependency management
- **Dynamic ID propagation** between test phases
- **Comprehensive validation** of all resource types
- **Automatic cleanup** after test completion

---

## Directory Structure

```
test/
├── ContactFlowAndModule/                # JSON files for contact flows
│   ├── BasicFlow.json
│   ├── BasicFlowModule.json
│   ├── QueueTransferFlow.json
│   └── WhisperFlow.json
├── common_utils.go                      # Utility functions for testing
├── connect_bot_associations.go          # Bot Association tests
├── connect_instance_storage_config.go  # Storage configuration tests
├── connect_queues.go                    # Queue-specific tests
├── connect_quick_connect.go             # Quick Connect tests
├── connect_routing_profile.go           # Routing Profile tests
├── connect_security_profiles.go         # Security Profile tests
├── connect_sso_idp.go                   # SSO Identity Provider tests
├── connect_test.go                      # Main test orchestrator
├── connect_user_hierarchy_group.go      # User Hierarchy Group tests
├── connect_user_hierarchy_structure.go # User Hierarchy Structure tests
├── connect_user.go                      # User management tests
├── connect_vocabularies.go              # Vocabulary tests
├── conntect_lambda_association.go       # Lambda Association tests
├── contact_flows_modules.go             # Contact Flow Module tests
├── contact_flows.go                     # Contact Flow tests
├── destruction_resources.go             # Cleanup and resource removal
├── hours_of_operation.go               # Hours of Operation tests
├── go.mod                              # Go module dependencies
├── go.sum                              # Go module checksums
├── provider.tf                         # Terraform provider configuration
└── README.md                           # Documentation
```

---

## Main Test Components

### 1. Test Architecture
- Uses **TestMain** as the orchestrator that sets up a shared Connect instance
- Implements **sequential testing** with dependency management
- Uses **dynamic ID propagation** between test phases
- Performs **cleanup** after all tests complete

### 2. Key Components Being Tested

#### Phase 1 - Core Resources
- ✅ **Lambda Function Associations** - Tests Connect-Lambda integrations
- ✅ **Hours of Operation** - Tests business hours configuration
- ✅ **Contact Flows** - Tests call routing logic (Basic, Whisper, Queue Transfer)
- ✅ **Contact Flow Modules** - Tests reusable flow components
- ✅ **Instance Storage Configuration** - Tests S3/Kinesis storage setup

#### Phase 2 - Dependent Resources
- ✅ **Queues** - Tests call queues (requires Hours of Operation + Contact Flows)
- ✅ **Quick Connects** - Tests quick connection options (requires Queues)
- ✅ **Routing Profiles** - Tests agent routing configurations
- ✅ **User Hierarchy Structure** - Tests organizational structure
- ✅ **User Hierarchy Groups** - Tests user group management (4-level hierarchy)

### 3. Test Configuration

#### Environment Setup
```go
region = "us-east-1"
instance_id = "184cff82-60e1-47ee-9875-71e202ab41b8" // Uses existing instance
create_instance = false  // Doesn't create new instance
```

#### Resource Naming Convention
```go
prefix: "jb-connect-test-cases-sandbox"
Pattern: "{company}-{lob}-{application}-{env}"
```

### 4. Testing Strategy

#### Sequential Dependency Testing
1. Creates base resources first
2. Captures dynamic IDs from outputs
3. Uses those IDs to configure dependent resources
4. Performs multiple terraform applies as dependencies are resolved

#### Dynamic ID Management
```go
dynamicOutBoundContactFlowID    // From contact flows
dynamicHoursOfOperationID       // From hours of operation  
dynamicQueuesID                 // From queues
dynamicUserHierarchyStructureID // From hierarchy structure
```

#### Test Data Sources
- JSON files in `ContactFlowAndModule/` for contact flow definitions
- Hardcoded configuration maps for resource properties
- Dynamic configuration building based on available IDs

### 5. Validation Approach

#### Resource Validation
- Tests resource creation and proper configuration
- Validates required tags are present
- Checks resource ARNs, IDs, and statuses
- Verifies cross-resource references work correctly

#### Cleanup Strategy
- Uses `defer` for cleanup guarantee
- Removes storage configurations via API
- Destroys all created resources

---

## Execution Flow and Dynamic Output Handling

### Test Execution Architecture

```
TestMain Starts
    ↓
Create Initial Terraform Options
    ↓
terraform.InitAndApply
    ↓
Run All Tests: m.Run()
    ↓
TestConnectCreation
    ↓
TestConnectConfiguration
    ↓
terraform.Destroy
    ↓
os.Exit
```

### Step-by-Step Execution Flow

#### Phase 1: Setup (TestMain)
```go
func TestMain(m *testing.M) {
    // 1. Create shared terraform options
    terraformOptions := terraform.WithDefaultRetryableErrors(&testing.T{}, &terraform.Options{
        TerraformDir: "../",
        Vars:         getCommonVars(),  // Initial vars without dynamic IDs
        NoStderr:     true,
        Logger:       logger.Discard,
    })

    // 2. Initialize and apply baseline configuration
    terraform.InitAndApply(&testing.T{}, terraformOptions)

    // 3. Run all tests
    code := m.Run()  // This calls TestConnectCreation and TestConnectConfiguration

    // 4. Cleanup
    terraform.Destroy(&testing.T{}, terraformOptions)
    os.Exit(code)
}
```

#### Phase 2: Basic Validation (TestConnectCreation)
```go
func TestConnectCreation(t *testing.T) {
    // Simple validation of basic outputs
    outputs := []string{
        "bot_associations", "contact_flow_modules", "contact_flows",
        "hours_of_operations", "instance_storage_configs",
        "lambda_function_associations", "queues", "quick_connects",
        "routing_profiles", "security_profiles", "user_hierarchy_groups",
        "users", "vocabularies"
    }

    // Just verify outputs exist (might be empty)
    for _, output := range outputs {
        _ = terraform.Output(t, terraformOptions, output)
    }
}
```

#### Phase 3: Complex Integration Testing (TestConnectConfiguration)

This is where the **magic happens** with dynamic output handling.

### Dynamic Output Handling Flow

#### Step 1: Test Independent Resources
```go
// Test Lambda Association (no dependencies)
LambdaFunctionAssociation := TestConnectLambdaAssociation(t, terraformOptions, instanceARN, instanceID, LambdaFunctionAssociationARN)
if !LambdaFunctionAssociation["Success"].(bool) {
    return  // Fail fast
}

// Test Hours of Operation (no dependencies)
HoursOfOperation := TestHoursOfOperation(t, terraformOptions, requiredTags, instanceARN, instanceID)
if !HoursOfOperation["Success"].(bool) {
    return
}

// Test Contact Flows (no dependencies)
ContactFlow := TestContactFlow(t, terraformOptions, requiredTags, instanceARN, instanceID)
if !ContactFlow["Success"].(bool) {
    return
}
```

#### Step 2: Extract Dynamic IDs
```go
// CRITICAL: Extract dynamic IDs from test results
dynamicOutBoundContactFlowID = ContactFlow["test_whisper_flows"].(map[string]interface{})["contact_flow_id"].(string)
dynamicHoursOfOperationID = HoursOfOperation["test_hours_of_operations"].(map[string]interface{})["hours_of_operation_id"].(string)  
dynamicQueueTransferFlowID = ContactFlow["test_queue_transfer_flows"].(map[string]interface{})["contact_flow_id"].(string)
```

#### Step 3: First Terraform Re-apply with New Dependencies
```go
// Re-apply terraform with updated variables (now includes queues)
terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../",
    Vars:         getCommonVars(), // Now includes queues because dynamic IDs are set
    // ... other options
})

terraform.Apply(t, terraformOptions)  // Apply updated config
```

#### Step 4: Test Resources with Dependencies
```go
// Test Queues (requires: outbound flow ID + hours of operation ID)  
QueuesConf := TestConnectQueues(t, terraformOptions, requiredTags, instanceARN, instanceID,
                                dynamicOutBoundContactFlowID, dynamicHoursOfOperationID)
if !QueuesConf["Success"].(bool) {
    return
}

// Extract queue ID for next phase
dynamicQueuesID = QueuesConf["test_queues"].(map[string]interface{})["queue_id"].(string)
```

#### Step 5: Second Re-apply for Next Level Dependencies
```go
// Another re-apply for quick connects
terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../",
    Vars:         getCommonVars(), // Now includes quick_connects
})

terraform.Apply(t, terraformOptions)

// Test Quick Connects (requires: queue ID + queue transfer flow ID)
QuickConnectConf := TestConnectQuickConnect(t, terraformOptions, requiredTags, instanceARN, instanceID,
                                          dynamicQueueTransferFlowID, dynamicQueuesID)
```

### How Dynamic Variables Work

#### `getCommonVars()` Function Logic:
```go
func getCommonVars() map[string]interface{} {
    vars := map[string]interface{}{
        // Base configuration always present
        "application": application,
        "instance_id": instance_id,
        // ... other base vars
    }

    // Conditional inclusion based on dynamic IDs
    if dynamicOutBoundContactFlowID != "" && dynamicHoursOfOperationID != "" {
        vars["queues"] = getQueuesConfig()  // Now queues can be created
    }

    if dynamicQueueTransferFlowID != "" && dynamicQueuesID != "" {
        vars["quick_connects"] = getQuickConnectConfig()  // Now quick connects can be created
    }

    if dynamicUserHierarchyStructureID > 0 {
        vars["user_hierarchy_groups"] = GetHierarchyGroupsConfig()
    }

    return vars
}
```

---

## Individual Test Function Patterns

Each test function follows this pattern:

### Input Parameters
```go
func TestConnectQueues(t *testing.T, terraformOptions *terraform.Options,
                      requiredTags []string, instanceARN string, instanceID string,
                      OutBoundFlowId string, HoursOfOperationId string) map[string]interface{}
```

### Processing Flow
```go
func TestConnectQueues(...) map[string]interface{} {
    initialFailed := t.Failed()  // Track test state

    // 1. Get Terraform output
    QueuesJson := terraform.OutputJson(t, terraformOptions, "queues")
    assert.NotEmpty(t, QueuesJson, "Queues should not be empty")

    // 2. Parse JSON output
    var QueuesConf map[string]interface{}
    err := json.Unmarshal([]byte(QueuesJson), &QueuesConf)
    assert.NoError(t, err, "Should be able to parse queues JSON")

    Response := make(map[string]interface{})

    // 3. Validate each resource
    for name, detailsRaw := range QueuesConf {
        // Parse resource details
        var detailsMap map[string]interface{}

        // 4. Perform validations
        assert.Equal(t, "test_queues", name, "Queues name should match")
        assert.Equal(t, instanceID, detailsMap["instance_id"], "Instance should match")
        assert.Equal(t, OutBoundFlowId, outboundCallerConfig["outbound_flow_id"], "Flow ID should match")

        // 5. Validate tags, ARNs, etc.
        assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

        // 6. Store results for next phase
        Response[name] = detailsMap
    }

    // 7. Return success status + data
    Response["Success"] = !t.Failed() || initialFailed == t.Failed()
    return Response
}
```

### Example: Hours of Operation Test Function

The `TestHoursOfOperation` function demonstrates comprehensive validation:

```go
func TestHoursOfOperation(t *testing.T, terraformOptions *terraform.Options, requiredTags []string, instanceARN string, instanceID string) map[string]interface{} {
    // Get and validate JSON output
    hoursOfOperationJSON := terraform.OutputJson(t, terraformOptions, "hours_of_operations")
    assert.NotEmpty(t, hoursOfOperationJSON, "Hours of operations should not be empty")

    // Parse JSON into map
    var hoursOfOperation map[string]interface{}
    err := json.Unmarshal([]byte(hoursOfOperationJSON), &hoursOfOperation)
    assert.NoError(t, err, "Should be able to parse hours_of_operations JSON")

    Response := make(map[string]interface{})

    // Validate each hours of operation configuration
    for name, detailsRaw := range hoursOfOperation {
        assert.Equal(t, "test_hours_of_operations", name, "Hours of Operation name should match")

        // Validate description and timezone
        assert.Equal(t, "24-7 Music on Hold hours", detailsMap["description"], "Description should match expected value")
        assert.Equal(t, "US/Eastern", detailsMap["time_zone"], "Time zone should match expected value")

        // Verify required tags
        assertTagsExist(t, detailsMap["tags"].(map[string]interface{}), requiredTags)

        // Validate ARN format and instance association
        InstanceArnFromHOO := strings.Split(detailsMap["arn"].(string), "/")
        if len(InstanceArnFromHOO) >= 2 {
            InstanceArnOfHOO := strings.Join(InstanceArnFromHOO[:2], "/")
            assert.Equal(t, instanceARN, InstanceArnOfHOO, "Instance ARN should match")
        }

        // Validate 7-day configuration
        assert.Equal(t, 7, len(detailsMap["config"].([]interface{})), "Config should have 7 days of operations")

        // Validate time settings for each day
        for _, dayConfig := range detailsMap["config"].([]interface{}) {
            entryMap, ok := dayConfig.(map[string]interface{})
            assert.True(t, ok, "Each config entry should be a map")

            // Validate time ranges (0-23 hours, 0-59 minutes)
            startHours := startTime[0].(map[string]interface{})["hours"].(float64)
            assert.GreaterOrEqual(t, startHours, float64(0))
            assert.LessOrEqual(t, startHours, float64(23))

            // Validate day names
            validDays := []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
            assert.Contains(t, validDays, day, "day should be a valid weekday")
        }

        Response[name] = detailsMap
    }

    Response["Success"] = !t.Failed() || initialFailed == t.Failed()
    return Response
}
```

---

## Cleanup Strategy

The cleanup follows **reverse dependency order** to avoid dependency conflicts:

### ResourceRemoval Function
```go
func ResourceRemoval(t *testing.T, ...) {
    // 1. Remove Lambda Function Associations (no dependencies on them)
    delete(varsDependsRemoval, "lambda_function_associations")
    terraform.Apply(t, terraformOptions)

    // 2. Remove Routing Profiles (depend on queues)
    delete(varsDependsRemoval, "routing_profiles")
    terraform.Apply(t, terraformOptions)

    // 3. Remove User Hierarchy Groups (in reverse order: children first)
    hierarchyGroups := []string{"MOCAgent", "MCCAgent", "AOGAgent",
                               "MOCPlanningSupervisor", "MCCDutySupervisor",
                               "AOGSupervisor", "MCCDutyManager", "maint"}
    for _, group := range hierarchyGroups {
        delete(groups, group)
        terraform.Apply(t, terraformOptions)
    }

    // 4. Remove Quick Connects (depend on queues)
    delete(varsDependsRemoval, "quick_connects")
    terraform.Apply(t, terraformOptions)

    // 5. Finally remove Queues (many things depend on them)
    delete(varsDependsRemoval, "queues")
    terraform.Apply(t, terraformOptions)
}
```

### Cleanup Order Strategy
1. **Lambda Function Associations** - No other resources depend on these
2. **Routing Profiles** - Depend on queues, so remove before queues
3. **User Hierarchy Groups** - Remove children before parents (reverse hierarchy order)
4. **Quick Connects** - Depend on queues and contact flows
5. **Queues** - Many resources depend on queues, so remove last

---

## Technical Architecture

### Key Benefits of This Architecture

1. **Dependency Management**: Handles complex AWS Connect resource dependencies automatically
2. **Fail-Fast**: Stops testing when foundational resources fail
3. **Dynamic Configuration**: Builds terraform vars based on available resources
4. **Proper Cleanup**: Removes resources in correct order to avoid dependency conflicts
5. **Comprehensive Validation**: Tests all aspects of each resource (ARNs, tags, configurations)

### Testing Framework Details
- **Framework**: Terratest (industry-standard for Terraform testing)
- **Test Type**: Integration tests (real AWS resources)
- **Resource Scope**: Full Connect instance with all components
- **Dependency Management**: Sequential with dynamic ID passing
- **Cleanup**: Automatic resource destruction

### Required Tags Validation
```go
var requiredTags = []string{
    "module_project_path",
    "module_version",
    "project_path",
    "commit_id",
    "company",
    "region",
    "lob",
    "application",
    "env",
    "created_by",
    "map-migrated",
}
```

### Utility Functions

#### Common Utilities (`common_utils.go`)
```go
// assertOutputsNonEmpty checks that specified Terraform outputs are non-empty
func assertOutputsNonEmpty(t *testing.T, options *terraform.Options, outputs []string)

// assertNameFormat ensures resource name follows expected format
func assertNameFormat(t *testing.T, name, expectedPrefix string)

// assertTagsExist checks if required tags exist in a given resource state
func assertTagsExist(t *testing.T, state map[string]interface{}, requiredTags []string)

// assertContains verifies that specific key-value pairs exist using regex
func assertContains(t *testing.T, plan string, checks map[string]string)

// mergeMaps combines two maps into one
func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{}
```

---

## Running the Tests

### Prerequisites
- Go 1.21 or higher
- Terraform installed
- AWS credentials configured
- Access to existing Connect instance (`184cff82-60e1-47ee-9875-71e202ab41b8`)

### Execution Commands
```bash
# Navigate to test directory
cd /ccaas-terraform-modules-wrapper/jb-terraform-aws-amazon-connect/test

# Run all tests
go test -v

# Run specific test
go test -v -run TestConnectCreation

# Run with verbose output
go test -v -timeout 60m
```

### Test Output Example
```
=== RUN   TestConnectCreation
--- PASS: TestConnectCreation (2.34s)
=== RUN   TestConnectConfiguration
    LambdaFunctionAssociation Test cases are successful.
    HoursOfOperation Test cases are successful.
    ContactFlow Test cases are successful.
    ContactFlowModules Test cases are successful.
    TestConnectInstaceStorageConf Test cases are successful.
    QueuesConf Test cases are successful.
    QuickConnectConf Test cases are successful.
    RoutingProfileConf Test cases are successful.
    UserHierarchyStructConf Test cases are successful.
    hierarchyGroups Test cases are successful.
--- PASS: TestConnectConfiguration (45.67s)
PASS
```

---

This documentation provides a comprehensive overview of your AWS Connect test suite, showing how it handles complex resource dependencies through dynamic output management and ensures thorough validation of all AWS Connect components. The test suite is designed to be production-ready and handles real-world scenarios with proper cleanup and error handling.
