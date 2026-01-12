package test

import (
	// "encoding/json"
	// "fmt"
	// "strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// File: test/connect_sso_idp_test.go
func TestConnectSSOIdentityProviders(t *testing.T, terraformOptions *terraform.Options, instanceARN string) map[string]interface{} {
    initialFailed := t.Failed()

    // Test Admin IDP ARN
    adminIdpArn := terraform.Output(t, terraformOptions, "connect_idp_admin_arn")
    if adminIdpArn != "" {
        assert.Contains(t, adminIdpArn, "arn:aws:iam::", "Admin IDP ARN should be valid IAM ARN")
        assert.Contains(t, adminIdpArn, "saml-provider", "Admin IDP should be SAML provider")
        assert.Contains(t, adminIdpArn, "AWSSSO_admin_Connect", "Admin IDP should contain admin identifier")
    }

    // Test Agent IDP ARN
    agentIdpArn := terraform.Output(t, terraformOptions, "connect_idp_agent_arn")
    if agentIdpArn != "" {
        assert.Contains(t, agentIdpArn, "arn:aws:iam::", "Agent IDP ARN should be valid IAM ARN")
        assert.Contains(t, agentIdpArn, "saml-provider", "Agent IDP should be SAML provider")
        assert.Contains(t, agentIdpArn, "AWSSSO_agent_Connect", "Agent IDP should contain agent identifier")
    }

    Response := map[string]interface{}{
        "admin_idp_arn": adminIdpArn,
        "agent_idp_arn": agentIdpArn,
        "Success":       !t.Failed() || initialFailed == t.Failed(),
    }
    return Response
}
