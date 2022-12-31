package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestMyTerraformCode(t *testing.T) {
	t.Parallel()

	os.Setenv("ATLAS_TOKEN", "eM6P0JyKaZzJaA.atlasv1.U0oMslxMeZtNlTMM5owqnPg0MomT8GIHa2GhyrTNl0zgE7vzksFfzCqHF2Wjzyuaflw")
	os.Setenv("ATLAS_WORKSPACE", "my-app-prod")

	// Set up the Terraform Cloud options
	terraformCloudOptions := &terraform.CloudOptions{
		Organization: "my-organization",
		Workspace:    "my-app-prod",
		Repository:   "auto-test",
		Branch:       "main",
		// Vars: map[string]interface{}{
		// 	"instance_type": "t2.micro",
		// },
		// VarsFiles: []string{"vars.tfvars"},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformCloudOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformCloudOptions)

	// Test code goes here
	// ...
}

// package test

// import (
// 	"fmt"
// 	"os"
// 	"testing"
// 	"time"

// 	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

// 	"github.com/gruntwork-io/terratest/modules/terraform"
// )

// func TestTerraformAwsHelloWorldExample(t *testing.T) {
// 	t.Parallel()

// 	// retryable errors in terraform testing.
// 	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
// 		TerraformDir: "C:\\Users\\admin\\Documents\\aterratest\\auto-test",
// 	})

// 	os.Setenv("ATLAS_TOKEN", "eM6P0JyKaZzJaA.atlasv1.U0oMslxMeZtNlTMM5owqnPg0MomT8GIHa2GhyrTNl0zgE7vzksFfzCqHF2Wjzyuaflw")
// 	os.Setenv("ATLAS_WORKSPACE", "my-app-prod")

// 	defer terraform.Destroy(t, terraformOptions)

// 	terraform.InitAndApply(t, terraformOptions)

// 	publicIp := terraform.Output(t, terraformOptions, "public_ip")

// 	url := fmt.Sprintf("http://%s:8080", publicIp)
// 	http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, World!", 30, 5*time.Second)
// }
