package vmwareengine

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareenginePrivateCloudAssetType string = "vmwareengine.googleapis.com/PrivateCloud"

func ResourceConverterVmwareenginePrivateCloud() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareenginePrivateCloudAssetType,
		Convert:   GetVmwareenginePrivateCloudCaiObject,
	}
}

func GetVmwareenginePrivateCloudCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//vmwareengine.googleapis.com/projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetVmwareenginePrivateCloudApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VmwareenginePrivateCloudAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://vmwareengine.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "PrivateCloud",
			Data:                 resource,
		},
	}}, nil
}

func GetVmwareenginePrivateCloudApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v.(string)
	}
	if v, ok := d.GetOk("network_config"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			nc := make(map[string]interface{})
			nc["managementCidr"] = l[0].(map[string]interface{})["management_cidr"].(string)
			nc["vmwareEngineNetwork"] = l[0].(map[string]interface{})["vmware_engine_network"].(string)
			obj["networkConfig"] = nc
		}
	}
	if v, ok := d.GetOk("management_cluster"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			mc := make(map[string]interface{})
			mc["clusterId"] = l[0].(map[string]interface{})["cluster_id"].(string)
			obj["managementCluster"] = mc
		}
	}
	return obj, nil
}
