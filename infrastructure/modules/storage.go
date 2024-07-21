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

type StorageBucket struct {
	Name   string
	Policy PolicyData
}

func (sb StorageBucket) New(stack cdktf.TerraformStack) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location:     jsii.String(infrastructure.Region),
		Name:         jsii.String(sb.Name + "-app-bucket"),
		ForceDestroy: jsii.Bool(true),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket:     bucket.Name(),
		PolicyData: jsii.String(sb.Policy.ToJSON()),
	})

}

type role interface {
	Role() string
}

type Role string

func (r Role) Role() string {
	return string(r)
}

const (
	RoleStorageAdmin         Role = "roles/storage.admin"
	RoleStorageObjectAdmin   Role = "roles/storage.objectAdmin"
	RoleStorageObjectViewer  Role = "roles/storage.objectViewer"
	RoleStorageObjectCreator Role = "roles/storage.objectCreator"
)

type Binding struct {
	Role    role     `json:"role"`
	Members []string `json:"members"`
}

type PolicyData struct {
	Bindings []Binding `json:"bindings"`
}

func (p *PolicyData) ToJSON() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling policy data: %v", err)
	}
	return string(data)
}
