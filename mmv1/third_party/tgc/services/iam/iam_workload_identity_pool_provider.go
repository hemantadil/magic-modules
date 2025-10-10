package iam

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IamWorkloadIdentityPoolProviderAssetType string = "iam.googleapis.com/WorkloadIdentityPoolProvider"

func ResourceConverterIamWorkloadIdentityPoolProvider() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IamWorkloadIdentityPoolProviderAssetType,
		Convert:   GetIamWorkloadIdentityPoolProviderCaiObject,
	}
}

func GetIamWorkloadIdentityPoolProviderCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//iam.googleapis.com/projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetIamWorkloadIdentityPoolProviderApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: IamWorkloadIdentityPoolProviderAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://iam.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "WorkloadIdentityPoolProvider",
			Data:                 resource,
		},
	}}, nil
}

func GetIamWorkloadIdentityPoolProviderApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
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
	if v, ok := d.GetOk("attribute_mapping"); ok {
		obj["attributeMapping"] = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("attribute_condition"); ok {
		obj["attributeCondition"] = v.(string)
	}
	if v, ok := d.GetOk("aws"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			aws := l[0].(map[string]interface{})
			a := make(map[string]interface{})
			if accountId, ok := aws["account_id"]; ok {
				a["accountId"] = accountId.(string)
			}
			obj["aws"] = a
		}
	}
	if v, ok := d.GetOk("oidc"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			oidc := l[0].(map[string]interface{})
			o := make(map[string]interface{})
			if issuerUri, ok := oidc["issuer_uri"]; ok {
				o["issuerUri"] = issuerUri.(string)
			}
			if allowedAudiences, ok := oidc["allowed_audiences"]; ok {
				o["allowedAudiences"] = allowedAudiences.([]interface{})
			}
			obj["oidc"] = o
		}
	}
	return obj, nil
}
