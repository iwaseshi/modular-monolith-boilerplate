package main

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func SetupGcsBackend(stack cdktf.TerraformStack) {
	cdktf.NewGcsBackend(stack, &cdktf.GcsBackendConfig{
		Bucket: jsii.String("modular-monolith-boilerplate-backend"),
		Prefix: jsii.String("terraform/state"),
	})
}
