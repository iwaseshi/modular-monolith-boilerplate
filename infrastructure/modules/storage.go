package modules

import (
	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/iamworkloadidentitypool"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/iamworkloadidentitypoolprovider"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location: jsii.String(infrastructure.Region),
		Name:     jsii.String(infrastructure.ProjectName + "-app-bucket"),
	})

	account := serviceaccount.NewServiceAccount(stack, jsii.String("app_sa"), &serviceaccount.ServiceAccountConfig{
		AccountId:   jsii.String("app-account"),
		DisplayName: jsii.String("app account"),
	})

	pool := iamworkloadidentitypool.NewIamWorkloadIdentityPool(stack, jsii.String("wi_pool"), &iamworkloadidentitypool.IamWorkloadIdentityPoolConfig{
		DisplayName:            jsii.String("Workload Identity Pool"),
		WorkloadIdentityPoolId: jsii.String("wi-pool"),
	})

	iamworkloadidentitypoolprovider.NewIamWorkloadIdentityPoolProvider(stack, jsii.String("wi_provider"), &iamworkloadidentitypoolprovider.IamWorkloadIdentityPoolProviderConfig{
		DisplayName:                    jsii.String("Workload Identity Pool Provider"),
		WorkloadIdentityPoolId:         pool.WorkloadIdentityPoolId(),
		WorkloadIdentityPoolProviderId: jsii.String("my-provider"),
		AttributeMapping: &map[string]*string{
			"google.subject": jsii.String("assertion.sub"),
		},
		Oidc: &iamworkloadidentitypoolprovider.IamWorkloadIdentityPoolProviderOidc{
			IssuerUri: jsii.String("https://accounts.google.com"),
		},
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: bucket.Name(),
		PolicyData: jsii.String(`{
			"bindings": [
				{
					"role": "roles/storage.admin",
					"members": [
						"serviceAccount:` + *account.Email() + `",
						"principalSet://iam.googleapis.com/` + *pool.Name() + `/attribute.google.subject/kubernetes.io/serviceaccount/namespace/service-account"
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
