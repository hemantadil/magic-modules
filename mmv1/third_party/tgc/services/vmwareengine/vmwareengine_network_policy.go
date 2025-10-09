package vmwareengine

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineNetworkPolicyAssetType string = "vmwareengine.googleapis.com/NetworkPolicy"

func ResourceConverterVmwareengineNetworkPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineNetworkPolicyAssetType,
		Convert:   GetVmwareengineNetworkPolicyCaiObject,
	}
}

func GetVmwareengineNetworkPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//vmwareengine.googleapis.com/projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetVmwareengineNetworkPolicyApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VmwareengineNetworkPolicyAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://vmwareengine.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "NetworkPolicy",
			Data:                 resource,
		},
	}}, nil
}

func GetVmwareengineNetworkPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("vmware_engine_network"); ok {
		obj["vmwareEngineNetwork"] = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v.(string)
	}
	if v, ok := d.GetOk("edge_services_cidr"); ok {
		obj["edgeServicesCidr"] = v.(string)
	}
	if v, ok := d.GetOk("internet_access"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			ia := make(map[string]interface{})
			ia["enabled"] = l[0].(map[string]interface{})["enabled"].(bool)
			obj["internetAccess"] = ia
		}
	}
	if v, ok := d.GetOk("external_ip"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			ei := make(map[string]interface{})
			ei["enabled"] = l[0].(map[string]interface{})["enabled"].(bool)
			obj["externalIp"] = ei
		}
	}

	return obj, nil
}
