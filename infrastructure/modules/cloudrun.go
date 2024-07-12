package modules

import (
	"fmt"

	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudrunservice"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudrunserviceiambinding"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudRun struct {
	Name    string
	Account serviceaccount.ServiceAccount
	Traffic []*cloudrunservice.CloudRunServiceTraffic
}

func (cr CloudRun) New(stack cdktf.TerraformStack) {
	image := fmt.Sprintf("gcr.io/%s/%s/app:latest", infrastructure.ProjectName, cr.Name)
	cloudrun := cloudrunservice.NewCloudRunService(stack, jsii.String("cloudrun"), &cloudrunservice.CloudRunServiceConfig{
		Location: jsii.String(infrastructure.Region),
		Name:     jsii.String(cr.Name + "-service"),
		Template: &cloudrunservice.CloudRunServiceTemplate{
			Spec: &cloudrunservice.CloudRunServiceTemplateSpec{
				Containers: []*cloudrunservice.CloudRunServiceTemplateSpecContainers{
					{
						Image: jsii.String(image),
					},
				},
				ServiceAccountName: cr.Account.Email(),
			},
		},
		Traffic: cr.Traffic,
	})
	cloudrunserviceiambinding.NewCloudRunServiceIamBinding(stack, jsii.String("cloudrun_iam"), &cloudrunserviceiambinding.CloudRunServiceIamBindingConfig{
		Service:  cloudrun.Name(),
		Location: cloudrun.Location(),
		Role:     jsii.String("roles/run.invoker"),
		Members: &[]*string{
			jsii.String("allUsers"),
		},
	})
}
