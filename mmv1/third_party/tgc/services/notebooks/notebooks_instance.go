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
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		obj["name"] = v.(string)
	}
	if v, ok := d.GetOk("machine_type"); ok {
		obj["machineType"] = v.(string)
	}
	if v, ok := d.GetOk("vm_image"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			vmImage := l[0].(map[string]interface{})
			vi := make(map[string]interface{})
			if project, ok := vmImage["project"]; ok {
				vi["project"] = project.(string)
			}
			if imageFamily, ok := vmImage["image_family"]; ok {
				vi["imageFamily"] = imageFamily.(string)
			}
			obj["vmImage"] = vi
		}
	}
	return obj, nil
}
