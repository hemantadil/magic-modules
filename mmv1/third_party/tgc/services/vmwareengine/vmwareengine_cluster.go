package vmwareengine

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.comcom/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineClusterAssetType string = "vmwareengine.googleapis.com/Cluster"

func ResourceConverterVmwareengineCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineClusterAssetType,
		Convert:   GetVmwareengineClusterCaiObject,
	}
}

func GetVmwareengineClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//vmwareengine.googleapis.com/{{parent}}/clusters/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetVmwareengineClusterApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VmwareengineClusterAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://vmwareengine.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "Cluster",
			Data:                 resource,
		},
	}}, nil
}

func GetVmwareengineClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("node_type_configs"); ok {
		configs := v.(*schema.Set).List()
		ntcs := make([]map[string]interface{}, 0, len(configs))
		for _, c := range configs {
			configMap := c.(map[string]interface{})
			ntc := make(map[string]interface{})
			if nodeTypeId, ok := configMap["node_type_id"]; ok {
				ntc["nodeTypeId"] = nodeTypeId.(string)
			}
			if nodeCount, ok := configMap["node_count"]; ok {
				ntc["nodeCount"] = nodeCount.(int)
			}
			if customCoreCount, ok := configMap["custom_core_count"]; ok {
				ntc["customCoreCount"] = customCoreCount.(int)
			}
			ntcs = append(ntcs, ntc)
		}
		obj["nodeTypeConfigs"] = ntcs
	}
	return obj, nil
}
