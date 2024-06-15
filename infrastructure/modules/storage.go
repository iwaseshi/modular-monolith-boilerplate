package modules

import (
	"encoding/json"
	"log"

	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack, name string, policyData PolicyData) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location:     jsii.String(infrastructure.Region),
		Name:         jsii.String(name + "-app-bucket"),
		ForceDestroy: jsii.Bool(true),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket:     bucket.Name(),
		PolicyData: jsii.String(policyData.ToJSON()),
	})

}

type Binding struct {
	Role    string   `json:"role"`
	Members []string `json:"members"`
}

type PolicyData struct {
	Bindings []Binding `json:"bindings"`
}

func NewPolicyData(bindings []Binding) *PolicyData {
	return &PolicyData{Bindings: bindings}
}

func (p *PolicyData) ToJSON() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling policy data: %v", err)
	}
	return string(data)
}
