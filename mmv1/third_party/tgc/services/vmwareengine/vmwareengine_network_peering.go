package vmwareengine

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineNetworkPeeringAssetType string = "vmwareengine.googleapis.com/NetworkPeering"

func ResourceConverterVmwareengineNetworkPeering() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineNetworkPeeringAssetType,
		Convert:   GetVmwareengineNetworkPeeringCaiObject,
	}
}

func GetVmwareengineNetworkPeeringCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//vmwareengine.googleapis.com/projects/{{project}}/locations/global/networkPeerings/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetVmwareengineNetworkPeeringApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VmwareengineNetworkPeeringAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://vmwareengine.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "NetworkPeering",
			Data:                 resource,
		},
	}}, nil
}

func GetVmwareengineNetworkPeeringApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v.(string)
	}
	if v, ok := d.GetOk("peer_network"); ok {
		obj["peerNetwork"] = v.(string)
	}
	if v, ok := d.GetOk("peer_network_type"); ok {
		obj["peerNetworkType"] = v.(string)
	}
	if v, ok := d.GetOk("vmware_engine_network"); ok {
		obj["vmwareEngineNetwork"] = v.(string)
	}
	if v, ok := d.GetOk("export_custom_routes"); ok {
		obj["exportCustomRoutes"] = v.(bool)
	}
	if v, ok := d.GetOk("import_custom_routes"); ok {
		obj["importCustomRoutes"] = v.(bool)
	}
	if v, ok := d.GetOk("export_custom_routes_with_public_ip"); ok {
		obj["exportCustomRoutesWithPublicIp"] = v.(bool)
	}
	if v, ok := d.GetOk("import_custom_routes_with_public_ip"); ok {
		obj["importCustomRoutesWithPublicIp"] = v.(bool)
	}
	return obj, nil
}
