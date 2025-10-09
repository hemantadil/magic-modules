package kms

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CloudkmsCryptoKeyVersionAssetType string = "cloudkms.googleapis.com/CryptoKeyVersion"

func ResourceConverterCloudkmsCryptoKeyVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudkmsCryptoKeyVersionAssetType,
		Convert:   GetCloudkmsCryptoKeyVersionCaiObject,
	}
}

func GetCloudkmsCryptoKeyVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.ReplaceVars(d, config, "//cloudkms.googleapis.com/{{crypto_key}}/cryptoKeyVersions/{{name}}")
	if err != nil {
		return nil, err
	}

	resource, err := GetCloudkmsCryptoKeyVersionApiObject(d, config)
	if err != nil {
		return nil, err
	}

	return []cai.Asset{{
		Name: name,
		Type: CloudkmsCryptoKeyVersionAssetType,
		Resource: &cai.AssetResource{
			Version:              "v1",
			DiscoveryDocumentURI: "https://cloudkms.googleapis.com/$discovery/rest?version=v1",
			DiscoveryName:        "CryptoKeyVersion",
			Data:                 resource,
		},
	}}, nil
}

func GetCloudkmsCryptoKeyVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if v, ok := d.GetOk("state"); ok {
		obj["state"] = v.(string)
	}
	if v, ok := d.GetOk("external_protection_level_options"); ok {
		if l := v.([]interface{}); len(l) > 0 && l[0] != nil {
			opts := l[0].(map[string]interface{})
			epleo := make(map[string]interface{})
			if uri, ok := opts["external_key_uri"]; ok && uri != "" {
				epleo["externalKeyUri"] = uri.(string)
			}
			if path, ok := opts["ekm_connection_key_path"]; ok && path != "" {
				epleo["ekmConnectionKeyPath"] = path.(string)
			}
			obj["externalProtectionLevelOptions"] = epleo
		}
	}
	return obj, nil
}
