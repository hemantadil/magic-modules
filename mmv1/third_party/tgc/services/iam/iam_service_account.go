package iam

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IamServiceAccountAssetType string = "iam.googleapis.com/ServiceAccount"

// ResourceConverterIamServiceAccount returns the ResourceConverter object for iam service account.
func ResourceConverterIamServiceAccount() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IamServiceAccountAssetType,
		Convert:   GetIamServiceAccountCaiObject,
	}
}

// GetIamServiceAccountCaiObject converts the Terraform plan data into a CAI object.
func GetIamServiceAccountCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/projects/{{project}}/serviceAccounts/{{account_id}}@{{project}}.iam.gserviceaccount.com")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIamServiceAccountApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IamServiceAccountAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://iam.googleapis.com/$discovery/rest?version=v1",
				DiscoveryName:        "ServiceAccount",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

// GetIamServiceAccountApiObject returns the CAI resource object.
func GetIamServiceAccountApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})

	nameProp, err := expandIamServiceAccountName(d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("account_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	emailProp, err := expandIamServiceAccountEmail(d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("email"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailProp)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}

	displayNameProp, err := expandIamServiceAccountDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	descriptionProp, err := expandIamServiceAccountDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	projectProp, err := expandIamServiceAccountProject(d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("project"); !tpgresource.IsEmptyValue(reflect.ValueOf(projectProp)) && (ok || !reflect.DeepEqual(v, projectProp)) {
		obj["projectId"] = projectProp
	}

	return obj, nil
}

func expandIamServiceAccountName(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	if accountId, ok := d.GetOk("account_id"); ok {
		return fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", project, accountId, project), nil
	}
	return nil, nil
}

func expandIamServiceAccountEmail(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandIamServiceAccountDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIamServiceAccountDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIamServiceAccountProject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetProject(d, config)
}
