package iam

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IamWorkloadIdentityPoolAssetType string = "iam.googleapis.com/WorkloadIdentityPool"

func ResourceConverterIamWorkloadIdentityPool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IamWorkloadIdentityPoolAssetType,
		Convert:   GetIamWorkloadIdentityPoolCaiObject,
	}
}

func GetIamWorkloadIdentityPoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//iam.googleapis.com/projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetIamWorkloadIdentityPoolApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: IamWorkloadIdentityPoolAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://iam.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "WorkloadIdentityPool",
			Data:                 resource,
		},
	}}, nil
}

func GetIamWorkloadIdentityPoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("display_name"); ok {
		obj["displayName"] = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v.(string)
	}
	if v, ok := d.GetOk("disabled"); ok {
		obj["disabled"] = v.(bool)
	}
	return obj, nil
}
