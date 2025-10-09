package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeImageAssetType string = "compute.googleapis.com/Image"

func ResourceConverterComputeImage() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeImageAssetType,
		Convert:   GetComputeImageCaiObject,
	}
}

func GetComputeImageCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//compute.googleapis.com/projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetComputeImageApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ComputeImageAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
			DiscoveryName:        "Image",
			Data:                 resource,
		},
	}}, nil
}

func GetComputeImageApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj, err := tpgresource.ConvertStringMap(d.Get("").(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	return obj, nil
}
