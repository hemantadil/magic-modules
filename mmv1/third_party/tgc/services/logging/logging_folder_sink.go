package logging

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceConverterLoggingFolderSink() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingProjectSinkAssetType,
		Convert:   GetLoggingFolderSinkCaiObject,
	}
}

func GetLoggingFolderSinkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//logging.googleapis.com/folders/{{folder}}/sinks/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetLoggingFolderSinkApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: LoggingProjectSinkAssetType,
		Resource: &cai.AssetResource{
			Version:              "v2",
			DiscoveryDocumentURI: "https://logging.googleapis.com/$discovery/rest?version=v2",
			DiscoveryName:        "LogSink",
			Data:                 resource,
		},
	}}, nil
}

func GetLoggingFolderSinkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj, err := tpgresource.ConvertStringMap(d.Get("").(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	// The `folder` field is not part of the API object, it's part of the URL.
	delete(obj, "folder")
	return obj, nil
}
