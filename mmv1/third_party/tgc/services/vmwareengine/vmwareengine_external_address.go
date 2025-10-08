package vmwareengine

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineExternalAddressAssetType string = "vmwareengine.googleapis.com/ExternalAddress"

func ResourceConverterVmwareengineExternalAddress() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineExternalAddressAssetType,
		Convert:   GetVmwareengineExternalAddressCaiObject,
	}
}

func GetVmwareengineExternalAddressCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//vmwareengine.googleapis.com/{{parent}}/externalAddresses/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetVmwareengineExternalAddressApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VmwareengineExternalAddressAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://vmwareengine.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "ExternalAddress",
			Data:                 resource,
		},
	}}, nil
}

func GetVmwareengineExternalAddressApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj, err := tpgresource.ConvertStringMap(d.Get("").(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	return obj, nil
}
