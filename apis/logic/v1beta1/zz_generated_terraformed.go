/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this IntegrationServiceEnvironment
func (mg *IntegrationServiceEnvironment) GetTerraformResourceType() string {
	return "azurerm_integration_service_environment"
}

// GetConnectionDetailsMapping for this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this IntegrationServiceEnvironment
func (tr *IntegrationServiceEnvironment) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this IntegrationServiceEnvironment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *IntegrationServiceEnvironment) LateInitialize(attrs []byte) (bool, error) {
	params := &IntegrationServiceEnvironmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *IntegrationServiceEnvironment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppActionCustom
func (mg *AppActionCustom) GetTerraformResourceType() string {
	return "azurerm_logic_app_action_custom"
}

// GetConnectionDetailsMapping for this AppActionCustom
func (tr *AppActionCustom) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppActionCustom
func (tr *AppActionCustom) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppActionCustom
func (tr *AppActionCustom) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppActionCustom
func (tr *AppActionCustom) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppActionCustom
func (tr *AppActionCustom) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppActionCustom
func (tr *AppActionCustom) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppActionCustom
func (tr *AppActionCustom) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppActionCustom using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppActionCustom) LateInitialize(attrs []byte) (bool, error) {
	params := &AppActionCustomParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppActionCustom) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppActionHTTP
func (mg *AppActionHTTP) GetTerraformResourceType() string {
	return "azurerm_logic_app_action_http"
}

// GetConnectionDetailsMapping for this AppActionHTTP
func (tr *AppActionHTTP) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppActionHTTP
func (tr *AppActionHTTP) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppActionHTTP
func (tr *AppActionHTTP) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppActionHTTP
func (tr *AppActionHTTP) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppActionHTTP
func (tr *AppActionHTTP) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppActionHTTP
func (tr *AppActionHTTP) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppActionHTTP
func (tr *AppActionHTTP) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppActionHTTP using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppActionHTTP) LateInitialize(attrs []byte) (bool, error) {
	params := &AppActionHTTPParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppActionHTTP) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppIntegrationAccount
func (mg *AppIntegrationAccount) GetTerraformResourceType() string {
	return "azurerm_logic_app_integration_account"
}

// GetConnectionDetailsMapping for this AppIntegrationAccount
func (tr *AppIntegrationAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppIntegrationAccount
func (tr *AppIntegrationAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppIntegrationAccount
func (tr *AppIntegrationAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppIntegrationAccount
func (tr *AppIntegrationAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppIntegrationAccount
func (tr *AppIntegrationAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppIntegrationAccount
func (tr *AppIntegrationAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppIntegrationAccount
func (tr *AppIntegrationAccount) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppIntegrationAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppIntegrationAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &AppIntegrationAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppIntegrationAccount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppIntegrationAccountBatchConfiguration
func (mg *AppIntegrationAccountBatchConfiguration) GetTerraformResourceType() string {
	return "azurerm_logic_app_integration_account_batch_configuration"
}

// GetConnectionDetailsMapping for this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppIntegrationAccountBatchConfiguration
func (tr *AppIntegrationAccountBatchConfiguration) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppIntegrationAccountBatchConfiguration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppIntegrationAccountBatchConfiguration) LateInitialize(attrs []byte) (bool, error) {
	params := &AppIntegrationAccountBatchConfigurationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppIntegrationAccountBatchConfiguration) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppIntegrationAccountPartner
func (mg *AppIntegrationAccountPartner) GetTerraformResourceType() string {
	return "azurerm_logic_app_integration_account_partner"
}

// GetConnectionDetailsMapping for this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppIntegrationAccountPartner
func (tr *AppIntegrationAccountPartner) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppIntegrationAccountPartner using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppIntegrationAccountPartner) LateInitialize(attrs []byte) (bool, error) {
	params := &AppIntegrationAccountPartnerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppIntegrationAccountPartner) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppIntegrationAccountSchema
func (mg *AppIntegrationAccountSchema) GetTerraformResourceType() string {
	return "azurerm_logic_app_integration_account_schema"
}

// GetConnectionDetailsMapping for this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppIntegrationAccountSchema
func (tr *AppIntegrationAccountSchema) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppIntegrationAccountSchema using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppIntegrationAccountSchema) LateInitialize(attrs []byte) (bool, error) {
	params := &AppIntegrationAccountSchemaParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppIntegrationAccountSchema) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppIntegrationAccountSession
func (mg *AppIntegrationAccountSession) GetTerraformResourceType() string {
	return "azurerm_logic_app_integration_account_session"
}

// GetConnectionDetailsMapping for this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppIntegrationAccountSession
func (tr *AppIntegrationAccountSession) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppIntegrationAccountSession using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppIntegrationAccountSession) LateInitialize(attrs []byte) (bool, error) {
	params := &AppIntegrationAccountSessionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppIntegrationAccountSession) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppTriggerCustom
func (mg *AppTriggerCustom) GetTerraformResourceType() string {
	return "azurerm_logic_app_trigger_custom"
}

// GetConnectionDetailsMapping for this AppTriggerCustom
func (tr *AppTriggerCustom) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppTriggerCustom
func (tr *AppTriggerCustom) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppTriggerCustom
func (tr *AppTriggerCustom) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppTriggerCustom
func (tr *AppTriggerCustom) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppTriggerCustom
func (tr *AppTriggerCustom) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppTriggerCustom
func (tr *AppTriggerCustom) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppTriggerCustom
func (tr *AppTriggerCustom) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppTriggerCustom using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppTriggerCustom) LateInitialize(attrs []byte) (bool, error) {
	params := &AppTriggerCustomParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppTriggerCustom) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppTriggerHTTPRequest
func (mg *AppTriggerHTTPRequest) GetTerraformResourceType() string {
	return "azurerm_logic_app_trigger_http_request"
}

// GetConnectionDetailsMapping for this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppTriggerHTTPRequest
func (tr *AppTriggerHTTPRequest) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppTriggerHTTPRequest using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppTriggerHTTPRequest) LateInitialize(attrs []byte) (bool, error) {
	params := &AppTriggerHTTPRequestParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppTriggerHTTPRequest) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppTriggerRecurrence
func (mg *AppTriggerRecurrence) GetTerraformResourceType() string {
	return "azurerm_logic_app_trigger_recurrence"
}

// GetConnectionDetailsMapping for this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppTriggerRecurrence
func (tr *AppTriggerRecurrence) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppTriggerRecurrence using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppTriggerRecurrence) LateInitialize(attrs []byte) (bool, error) {
	params := &AppTriggerRecurrenceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppTriggerRecurrence) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AppWorkflow
func (mg *AppWorkflow) GetTerraformResourceType() string {
	return "azurerm_logic_app_workflow"
}

// GetConnectionDetailsMapping for this AppWorkflow
func (tr *AppWorkflow) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AppWorkflow
func (tr *AppWorkflow) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AppWorkflow
func (tr *AppWorkflow) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AppWorkflow
func (tr *AppWorkflow) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AppWorkflow
func (tr *AppWorkflow) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AppWorkflow
func (tr *AppWorkflow) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this AppWorkflow
func (tr *AppWorkflow) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this AppWorkflow using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AppWorkflow) LateInitialize(attrs []byte) (bool, error) {
	params := &AppWorkflowParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AppWorkflow) GetTerraformSchemaVersion() int {
	return 0
}
