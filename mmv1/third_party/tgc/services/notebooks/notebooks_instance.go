package notebooks

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NotebooksInstanceAssetType string = "notebooks.googleapis.com/Instance"

func ResourceConverterNotebooksInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NotebooksInstanceAssetType,
		Convert:   GetNotebooksInstanceCaiObject,
	}
}

func GetNotebooksInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//notebooks.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetNotebooksInstanceApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: NotebooksInstanceAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://notebooks.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "Instance",
			Data:                 resource,
		},
	}}, nil
}

func GetNotebooksInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj, err := tpgresource.ConvertStringMap(d.Get("").(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	// The `project` and `location` fields are not part of the API object, they're part of the URL.
	delete(obj, "project")
	delete(obj, "location")
	return obj, nil
}
