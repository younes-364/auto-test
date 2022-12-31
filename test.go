package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/tree/master/modules/aws"

	"github.com/gruntwork-io/terratest/modules/terraform"

	// test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	test_structure "github.com/gruntwork-io/terratest/tree/master/modules/test-structure"
)

func TestCreateEc2Instance(t *testing.T) {
	t.Parallel()

	// Load the test environment variables
	test_structure.LoadIntegrationTestOptions(t, "test-fixture")

	// Create a Terraform options struct
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../auto-test",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"instance_type": "t2.micro",
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Look up the ID of the EC2 instance created by the Terraform module
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Check that the instance is running
	instance := aws.GetEc2InstanceById(t, instanceID)
	if instance.State.Name != "running" {
		t.Fatalf("Expected instance to be running, got %s", instance.State.Name)
	}

	// Check that the instance has the expected instance type
	expectedInstanceType := terraformOptions.Vars["instance_type"].(string)
	if instance.InstanceType != expectedInstanceType {
		t.Fatalf("Expected instance to be %s, got %s", expectedInstanceType, instance.InstanceType)
	}

	fmt.Println("EC2 instance created successfully")
}
