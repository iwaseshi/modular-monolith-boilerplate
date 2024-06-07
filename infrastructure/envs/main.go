package main

import (
	infrastructure "cdk.tf/go/stack"
	"cdk.tf/go/stack/services/fileservice"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	SetupGoogleProvider(stack)
	SetupGcsBackend(stack)

	fileservice.DeployResources(stack)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	// envsは各環境に置き換えること
	NewMyStack(app, infrastructure.ProjectName+"envs")
	app.Synth()
}
