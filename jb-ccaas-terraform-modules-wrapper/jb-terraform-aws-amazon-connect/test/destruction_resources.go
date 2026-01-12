package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func ResourceRemoval(t *testing.T, terraformOptions *terraform.Options, varsDependsRemoval map[string]interface{}) map[string]interface{} {
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         varsDependsRemoval,
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})

	delete(varsDependsRemoval, "lambda_function_associations")
	// delete(varsDependsRemoval, "instance_storage_configs")
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         varsDependsRemoval, // This will now remove quick connects
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})
	terraform.Apply(t, terraformOptions)

	delete(varsDependsRemoval, "routing_profiles")
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         varsDependsRemoval, // This will now remove quick connects
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})
	terraform.Apply(t, terraformOptions)

	if groups, ok := varsDependsRemoval["user_hierarchy_groups"].(map[string]interface{}); ok {

		hierarchyGroups := []string{"MOCAgent", "MCCAgent", "AOGAgent", "MOCPlanningSupervisor", "MCCDutySupervisor", "AOGSupervisor", "MCCDutyManager", "maint"}
		for _, group := range hierarchyGroups {
			delete(groups, group)
			varsDependsRemoval["user_hierarchy_groups"] = groups
			terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "../",
				Vars:         varsDependsRemoval,
				NoStderr:     true,
				Logger:       logger.Discard,
				EnvVars: map[string]string{
					"TF_LOG":      "ERROR",
					"TF_LOG_PATH": "/dev/null",
				},
			})
			terraform.Apply(t, terraformOptions)
		}
	}

	delete(varsDependsRemoval, "user_hierarchy_structure")
	delete(varsDependsRemoval, "quick_connects")
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         varsDependsRemoval, // This will now remove quick connects
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})
	terraform.Apply(t, terraformOptions)

	delete(varsDependsRemoval, "queues")
	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         varsDependsRemoval, // This will now remove queues
		NoStderr:     true,
		Logger:       logger.Discard,
		EnvVars: map[string]string{
			"TF_LOG":      "ERROR",
			"TF_LOG_PATH": "/dev/null",
		},
	})
	terraform.Apply(t, terraformOptions)

	return varsDependsRemoval
}
