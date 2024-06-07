package modules

import (
	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudrunservice"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudrunserviceiambinding"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewCloudRun(stack cdktf.TerraformStack, name string, account serviceaccount.ServiceAccount) {
	cr := cloudrunservice.NewCloudRunService(stack, jsii.String("cloudrun"), &cloudrunservice.CloudRunServiceConfig{
		Location: jsii.String(infrastructure.Region),
		Name:     jsii.String(name + "-service"),
		Template: &cloudrunservice.CloudRunServiceTemplate{
			Spec: &cloudrunservice.CloudRunServiceTemplateSpec{
				Containers: []*cloudrunservice.CloudRunServiceTemplateSpecContainers{
					{
						Image: jsii.String("gcr.io/modular-monolith-sample/" + name + "/app"),
					},
				},
				ServiceAccountName: account.Email(),
			},
		},
	})

	cloudrunserviceiambinding.NewCloudRunServiceIamBinding(stack, jsii.String("cloudrun_iam"), &cloudrunserviceiambinding.CloudRunServiceIamBindingConfig{
		Service:  cr.Name(),
		Location: cr.Location(),
		Role:     jsii.String("roles/run.invoker"),
		Members: &[]*string{
			jsii.String("allUsers"),
		},
	})
}
