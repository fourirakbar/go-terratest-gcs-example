package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformGcpExample(t *testing.T) {
	t.Parallel()

	exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/create-gcs")

	// Get the Project Id to use
	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)

	// Create all resources in the following zone
	location := "asia-southeast1"

	name := "terratest-gcp-example"
	storageClass := "REGIONAL"
	forceDestroy := false

	// Give the example bucket a unique name so we can distinguish it from any other bucket in your GCP account
	// expectedBucketName := fmt.Sprintf("terratest-gcp-example")

	// website::tag::1::Configure Terraform setting path to Terraform code, bucket name, and instance name.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"project":       projectID,
			"name":          name,
			"location":      location,
			"storage_class": storageClass,
			"force_destroy": forceDestroy,
		},
	}

	// website::tag::5::At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2::This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of some of the output variables
	bucketURL := terraform.Output(t, terraformOptions, "bucket_url")

	// website::tag::3::Verify that the new bucket url matches the expected url
	expectedURL := fmt.Sprintf("gs://%s", name)
	assert.Equal(t, expectedURL, bucketURL)

	// Verify that the Storage Bucket exists
	gcp.AssertStorageBucketExists(t, name)
}
