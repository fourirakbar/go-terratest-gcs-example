package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fourirakbar/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformGcpExample(t *testing.T) {
	t.Parallel()

	// exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/create-gcs-v2")
	exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../", "")

	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)
	println(projectID)

	location := "asia-southeast1"
	// name := "terratest-gcp-example"
	storageClass := "REGIONAL"
	forceDestroy := false

	expectedBucketName := fmt.Sprintf("tmp-terratest-gcs-example-%s", strings.ToLower(random.UniqueId()))
	// website::tag::1::Configure Terraform setting path to Terraform code, bucket name, and instance name.
	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,

		Vars: map[string]interface{}{
			"project":       projectID,
			"name":          expectedBucketName,
			"location":      location,
			"storage_class": storageClass,
			"force_destroy": forceDestroy,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketURL := terraform.Output(t, terraformOptions, "bucket_url")
	expectedURL := fmt.Sprintf("gs://%s", expectedBucketName)
	assert.Equal(t, expectedURL, bucketURL)
	gcp.AssertStorageBucketExists(t, expectedBucketName)
}
