package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Project: jsii.String("modular-monolith-boilerplate"),
		Region:  jsii.String("asia-northeast1"),
		Zone:    jsii.String("asia-northeast1-a"),
	})

	cdktf.NewGcsBackend(stack, &cdktf.GcsBackendConfig{
		Bucket: jsii.String("modular-monolith-boilerplate-backend"),
		Prefix: jsii.String("state"),
		// 個人開発のためstate lockは不要
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	NewMyStack(app, "modular-monolith-boilerplater")
	app.Synth()
}
