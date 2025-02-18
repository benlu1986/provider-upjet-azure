// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ContentKeyPolicy
func (mg *ContentKeyPolicy) GetTerraformResourceType() string {
	return "azurerm_media_content_key_policy"
}

// GetConnectionDetailsMapping for this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"policy_option[*].fairplay_configuration[*].ask": "spec.forProvider.policyOption[*].fairplayConfiguration[*].askSecretRef", "policy_option[*].fairplay_configuration[*].pfx": "spec.forProvider.policyOption[*].fairplayConfiguration[*].pfxSecretRef", "policy_option[*].fairplay_configuration[*].pfx_password": "spec.forProvider.policyOption[*].fairplayConfiguration[*].pfxPasswordSecretRef", "policy_option[*].playready_configuration_license[*].grace_period": "spec.forProvider.policyOption[*].playreadyConfigurationLicense[*].gracePeriodSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].rsa_token_key_exponent": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].rsaTokenKeyExponentSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].rsa_token_key_modulus": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].rsaTokenKeyModulusSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].symmetric_token_key": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].symmetricTokenKeySecretRef", "policy_option[*].token_restriction[*].alternate_key[*].x509_token_key_raw": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].x509TokenKeyRawSecretRef", "policy_option[*].token_restriction[*].primary_rsa_token_key_exponent": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryRsaTokenKeyExponentSecretRef", "policy_option[*].token_restriction[*].primary_rsa_token_key_modulus": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryRsaTokenKeyModulusSecretRef", "policy_option[*].token_restriction[*].primary_symmetric_token_key": "spec.forProvider.policyOption[*].tokenRestriction[*].primarySymmetricTokenKeySecretRef", "policy_option[*].token_restriction[*].primary_x509_token_key_raw": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryX509TokenKeyRawSecretRef"}
}

// GetObservation of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ContentKeyPolicy
func (tr *ContentKeyPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ContentKeyPolicy
func (tr *ContentKeyPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this ContentKeyPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ContentKeyPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &ContentKeyPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ContentKeyPolicy) GetTerraformSchemaVersion() int {
	return 1
}
