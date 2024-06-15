package modules

import (
	"encoding/json"
	"log"

	infrastructure "cdk.tf/go/stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack, name string, account serviceaccount.ServiceAccount) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location:     jsii.String(infrastructure.Region),
		Name:         jsii.String(name + "-app-bucket"),
		ForceDestroy: jsii.Bool(true),
	})

	policyDataJSON, err := NewPolicyData([]Binding{
		{
			Role:    "roles/storage.admin",
			Members: []string{"serviceAccount:" + *account.Email()},
		},
		{
			Role:    "roles/storage.legacyObjectReader",
			Members: []string{"allUsers"},
		},
	}).ToJSON()
	if err != nil {
		log.Fatalf("Error marshalling policy data: %v", err)
	}
	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket:     bucket.Name(),
		PolicyData: jsii.String(policyDataJSON),
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

func (p *PolicyData) ToJSON() (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
