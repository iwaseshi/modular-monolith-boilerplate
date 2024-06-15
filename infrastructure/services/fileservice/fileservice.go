package fileservice

import (
	"cdk.tf/go/stack/modules"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	ServiceName = "fileservice"
)

func DeployResources(stack cdktf.TerraformStack) cdktf.TerraformStack {

	account := serviceaccount.NewServiceAccount(stack, jsii.String("app_sa"), &serviceaccount.ServiceAccountConfig{
		AccountId:   jsii.String(ServiceName + "-account"),
		DisplayName: jsii.String(ServiceName + " account"),
	})

	policyData := modules.NewPolicyData([]modules.Binding{
		{
			Role:    "roles/storage.admin",
			Members: []string{"serviceAccount:" + *account.Email()},
		},
		{
			Role:    "roles/storage.legacyObjectReader",
			Members: []string{"allUsers"},
		},
	})
	modules.NewStorageBucket(stack, ServiceName, *policyData)
	modules.NewCloudRun(stack, ServiceName, account)

	return stack
}
