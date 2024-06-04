package main

import (
	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func SetupGoogleProvider(stack cdktf.TerraformStack) {
	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Project: jsii.String(infrastructure.ProjectName),
		Region:  jsii.String(infrastructure.Region),
		Zone:    jsii.String(infrastructure.Region + "-a"),
	})
}
