package test

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRgSuccessful(t *testing.T) {

	rgName := "rg-module-test-01"

	os.Rename("./provider.tf", "../provider.tf")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         map[string]interface{}{"name": rgName},
	})

	defer postTest(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output_rg_name := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, rgName, output_rg_name)
}

func TestMandatoryTags(t *testing.T) {

	rgName := "rg-module-test-03"

	mandatoryTags := []string{"Environment", "Cost", "Project"}

	os.Rename("./provider.tf", "../provider.tf")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars:         map[string]interface{}{"name": rgName},
	})

	defer postTest(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	logger.Log(t, "Testing Terratest Logger")
	if err != nil {
		logger.Log(t, "There was an error authenticating azure")
		return
	}

	rgClient, err := armresources.NewResourceGroupsClient("72a70c6f-90c3-4aca-a78d-6e538352c901", cred, nil)
	rg, err := rgClient.Get(context.Background(), rgName, nil)

	for _, tag := range mandatoryTags {
		assert.Contains(t, rg.Tags, tag)
	}

}

func postTest(t *testing.T, tfOptions *terraform.Options) {

	terraform.Destroy(t, tfOptions)
	os.Rename("../provider.tf", "./provider.tf")
	os.Remove("../terraform.tfstate")
}
