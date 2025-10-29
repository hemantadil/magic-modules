package resourcemanager

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ServiceAccountAssetType string = "iam.googleapis.com/ServiceAccount"

func ResourceConverterServiceAccount() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ServiceAccountAssetType,
		Convert:   GetServiceAccountCaiObject,
	}
}

func GetServiceAccountCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/projects/{{project}}/serviceAccounts/{{account_id}}@{{project}}.iam.gserviceaccount.com")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetServiceAccountApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ServiceAccountAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://iam.googleapis.com/$discovery/rest",
				DiscoveryName:        "ServiceAccount",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetServiceAccountApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})

	nameProp, err := expandServiceAccountName(d.Get("account_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("account_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	emailProp, err := expandServiceAccountEmail(d.Get("email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("email"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailProp)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}

	displayNameProp, err := expandServiceAccountDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	descriptionProp, err := expandServiceAccountDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	projectProp, err := expandServiceAccountProject(d.Get("project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("project"); !tpgresource.IsEmptyValue(reflect.ValueOf(projectProp)) && (ok || !reflect.DeepEqual(v, projectProp)) {
		obj["projectId"] = projectProp
	}

	return obj, nil
}

func expandServiceAccountName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	if accountId, ok := d.GetOk("account_id"); ok {
		return fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", project, accountId, project), nil
	}
	return nil, nil
}

func expandServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if email, ok := d.GetOk("email"); ok && email != "" {
		return email, nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	if accountId, ok := d.GetOk("account_id"); ok {
		return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", accountId, project), nil
	}
	return nil, nil
}

func expandServiceAccountDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceAccountDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceAccountProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetProject(d, config)
}
