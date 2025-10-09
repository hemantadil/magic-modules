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
	obj, err := tpgresource.ConvertStringMap(d.Get("").(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	// The `crypto_key` field is not part of the API object, it's part of the URL.
	delete(obj, "crypto_key")
	return obj, nil
}
