// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WorkspaceSecurityAlertPolicyInitParameters struct {

	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action.
	// +listType=set
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to false.
	EmailAccountAdminsEnabled *bool `json:"emailAccountAdminsEnabled,omitempty" tf:"email_account_admins_enabled,omitempty"`

	// Specifies an array of email addresses to which the alert is sent.
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Possible values are Disabled, Enabled and New.
	PolicyState *string `json:"policyState,omitempty" tf:"policy_state,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to 0.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`
}

type WorkspaceSecurityAlertPolicyObservation struct {

	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action.
	// +listType=set
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to false.
	EmailAccountAdminsEnabled *bool `json:"emailAccountAdminsEnabled,omitempty" tf:"email_account_admins_enabled,omitempty"`

	// Specifies an array of email addresses to which the alert is sent.
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// The ID of the Synapse Workspace Security Alert Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Possible values are Disabled, Enabled and New.
	PolicyState *string `json:"policyState,omitempty" tf:"policy_state,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to 0.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

type WorkspaceSecurityAlertPolicyParameters struct {

	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action.
	// +kubebuilder:validation:Optional
	// +listType=set
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to false.
	// +kubebuilder:validation:Optional
	EmailAccountAdminsEnabled *bool `json:"emailAccountAdminsEnabled,omitempty" tf:"email_account_admins_enabled,omitempty"`

	// Specifies an array of email addresses to which the alert is sent.
	// +kubebuilder:validation:Optional
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Possible values are Disabled, Enabled and New.
	// +kubebuilder:validation:Optional
	PolicyState *string `json:"policyState,omitempty" tf:"policy_state,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to 0.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Specifies the identifier key of the Threat Detection audit storage account.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`

	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceSecurityAlertPolicySpec defines the desired state of WorkspaceSecurityAlertPolicy
type WorkspaceSecurityAlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceSecurityAlertPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WorkspaceSecurityAlertPolicyInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceSecurityAlertPolicyStatus defines the observed state of WorkspaceSecurityAlertPolicy.
type WorkspaceSecurityAlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceSecurityAlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkspaceSecurityAlertPolicy is the Schema for the WorkspaceSecurityAlertPolicys API. Manages a Security Alert Policy for a Synapse Workspace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyState) || (has(self.initProvider) && has(self.initProvider.policyState))",message="spec.forProvider.policyState is a required parameter"
	Spec   WorkspaceSecurityAlertPolicySpec   `json:"spec"`
	Status WorkspaceSecurityAlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceSecurityAlertPolicyList contains a list of WorkspaceSecurityAlertPolicys
type WorkspaceSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceSecurityAlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceSecurityAlertPolicy_Kind             = "WorkspaceSecurityAlertPolicy"
	WorkspaceSecurityAlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceSecurityAlertPolicy_Kind}.String()
	WorkspaceSecurityAlertPolicy_KindAPIVersion   = WorkspaceSecurityAlertPolicy_Kind + "." + CRDGroupVersion.String()
	WorkspaceSecurityAlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceSecurityAlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceSecurityAlertPolicy{}, &WorkspaceSecurityAlertPolicyList{})
}
