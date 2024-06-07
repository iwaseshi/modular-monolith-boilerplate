package modules

import (
	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack, account serviceaccount.ServiceAccount) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location:     jsii.String(infrastructure.Region),
		Name:         jsii.String(infrastructure.ProjectName + "-app-bucket"),
		ForceDestroy: jsii.Bool(true),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: bucket.Name(),
		PolicyData: jsii.String(`{
			"bindings": [
				{
					"role": "roles/storage.admin",
					"members": [
						"serviceAccount:` + *account.Email() + `"
					]
				},
				{
					"role":"roles/storage.legacyObjectReader",
					"members": [
						"allUsers"
					]
				}
			]
		}`),
	})

}
