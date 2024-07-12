package fileservice

import (
	"cdk.tf/go/stack/modules"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	Name = "fileservice"
)

func DeployResources(stack cdktf.TerraformStack) cdktf.TerraformStack {
	account := serviceaccount.NewServiceAccount(stack, jsii.String("app_sa"), &serviceaccount.ServiceAccountConfig{
		AccountId:   jsii.String(Name + "-account"),
		DisplayName: jsii.String(Name + " account"),
	})
	modules.StorageBucket{
		Name: Name,
		Policy: modules.PolicyData{
			Bindings: []modules.Binding{
				{
					Role:    "roles/storage.admin",
					Members: []string{"serviceAccount:" + *account.Email()},
				},
				{
					Role:    "roles/storage.legacyObjectReader",
					Members: []string{"allUsers"},
				},
			},
		},
	}.New(stack)
	modules.CloudRun{
		Name:    Name,
		Account: account,
	}.New(stack)
	return stack
}
