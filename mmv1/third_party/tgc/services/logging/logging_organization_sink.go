package logging

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceConverterLoggingOrganizationSink() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingProjectSinkAssetType,
		Convert:   GetLoggingOrganizationSinkCaiObject,
	}
}

func GetLoggingOrganizationSinkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//logging.googleapis.com/organizations/{{org_id}}/sinks/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetLoggingOrganizationSinkApiObject(d, config)
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

func GetLoggingOrganizationSinkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		obj["name"] = v.(string)
	}
	if v, ok := d.GetOk("destination"); ok {
		obj["destination"] = v.(string)
	}
	if v, ok := d.GetOk("filter"); ok {
		obj["filter"] = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v.(string)
	}
	if v, ok := d.GetOk("disabled"); ok {
		obj["disabled"] = v.(bool)
	}
	if v, ok := d.GetOk("exclusions"); ok {
		exclusions := v.([]interface{})
		es := make([]map[string]interface{}, 0, len(exclusions))
		for _, e := range exclusions {
			exclusion := e.(map[string]interface{})
			entry := make(map[string]interface{})
			if name, ok := exclusion["name"]; ok {
				entry["name"] = name.(string)
			}
			if description, ok := exclusion["description"]; ok {
				entry["description"] = description.(string)
			}
			if filter, ok := exclusion["filter"]; ok {
				entry["filter"] = filter.(string)
			}
			if disabled, ok := exclusion["disabled"]; ok {
				entry["disabled"] = disabled.(bool)
			}
			es = append(es, entry)
		}
		obj["exclusions"] = es
	}
	if v, ok := d.GetOk("bigquery_options"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			opts := l[0].(map[string]interface{})
			bqo := make(map[string]interface{})
			if usePartitionedTables, ok := opts["use_partitioned_tables"]; ok {
				bqo["usePartitionedTables"] = usePartitionedTables.(bool)
			}
			obj["bigqueryOptions"] = bqo
		}
	}
	return obj, nil
}
