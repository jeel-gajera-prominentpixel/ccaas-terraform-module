package test

import (
	_ "fmt"
	"log"
	"math"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	prefix_company = "digi"
	application    = "app"
	env            = "dev"
	company        = "Digiclarity"
	region_suffix  = "use1"

	baseTerraformVars = map[string]interface{}{
		"company":        company,
		"company_prefix": prefix_company,
		"region_suffix":  region_suffix,
		"application":    application,
		"environment":    env,
	}
)

func getTerraformOptions(t *testing.T, additionalVars map[string]interface{}) *terraform.Options {
	vars := make(map[string]interface{})

	for k, v := range baseTerraformVars {
		vars[k] = v
	}

	for k, v := range additionalVars {
		vars[k] = v
	}

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         vars,
	})
}

func getAWSAccountID() (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}

	svc := sts.New(sess)

	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		return "", err
	}

	return *result.Account, nil
}

func renderUI(passPercentage int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v.", err)
	}
	defer ui.Close()

	// Create a gauge widget
	g := widgets.NewGauge()
	g.Title = "Test Execution Progress"
	g.Percent = passPercentage
	g.SetRect(0, 0, 50, 5)
	g.BarColor = ui.ColorGreen
	g.BorderStyle.Fg = ui.ColorWhite
	g.LabelStyle.Fg = ui.ColorBlue

	// Render the gauge
	ui.Render(g)

	// Wait for user to close
	uiEvents := ui.PollEvents()
	for e := range uiEvents {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func TestEventbridgeBusCreation(t *testing.T) {
	additionalVars := map[string]interface{}{
		"bus_name": "test-bus",
	}

	terraformOptions := getTerraformOptions(t, additionalVars)
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	totalTests := 1
	passedTests := 0
	failedTests := 0

	busName := terraform.Output(t, terraformOptions, "eventbridge_bus_name")
	if assert.NotEmpty(t, busName, "Bus should be created.") {
		passedTests++
	} else {
		failedTests++
	}

	passPercentage := int(math.Round(float64(passedTests) / float64(totalTests) * 100))
	renderUI(passPercentage)

	// Ensure all tests pass.
	require.Equal(t, 0, failedTests, "There are failed test cases.")
}

func TestEventbridgeRuleTargetCreation(t *testing.T) {
	accountID, err := getAWSAccountID()
	if err != nil {
		t.Fatalf("Failed to fetch AWS account ID: %v", err)
	}

	additionalVars := map[string]interface{}{
		"rules": map[string]interface{}{
			"crons": map[string]interface{}{
				"description":         "Trigger for a Lambda.",
				"schedule_expression": "rate(5 minutes)",
			},
		},
		"targets": map[string]interface{}{
			"crons": []map[string]interface{}{
				{
					"name": "something-for-cron",
					"arn":  "arn:aws:lambda:us-west-2:" + accountID + ":function:MyFunction",
				},
			},
		},
	}

	terraformOptions := getTerraformOptions(t, additionalVars)
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	totalTests := 2
	passedTests := 0
	failedTests := 0

	ruleName := terraform.Output(t, terraformOptions, "eventbridge_rules")
	if assert.NotEmpty(t, ruleName, "EventBridge rule should be created.") {
		passedTests++
	} else {
		failedTests++
	}

	targetName := terraform.Output(t, terraformOptions, "eventbridge_targets")
	if assert.NotEmpty(t, targetName, "EventBridge target should be created.") {
		passedTests++
	} else {
		failedTests++
	}

	passPercentage := int(math.Round(float64(passedTests) / float64(totalTests) * 100))
	renderUI(passPercentage)

	// Ensure all tests pass.
	require.Equal(t, 0, failedTests, "There are failed test cases.")
}
