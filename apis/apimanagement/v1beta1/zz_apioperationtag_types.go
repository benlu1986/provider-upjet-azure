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

type APIOperationTagInitParameters struct {

	// The display name of the API Management API Operation Tag.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

type APIOperationTagObservation struct {

	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	APIOperationID *string `json:"apiOperationId,omitempty" tf:"api_operation_id,omitempty"`

	// The display name of the API Management API Operation Tag.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the API Management API Operation Tag.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIOperationTagParameters struct {

	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.APIOperation
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIOperationID *string `json:"apiOperationId,omitempty" tf:"api_operation_id,omitempty"`

	// Reference to a APIOperation in apimanagement to populate apiOperationId.
	// +kubebuilder:validation:Optional
	APIOperationIDRef *v1.Reference `json:"apiOperationIdRef,omitempty" tf:"-"`

	// Selector for a APIOperation in apimanagement to populate apiOperationId.
	// +kubebuilder:validation:Optional
	APIOperationIDSelector *v1.Selector `json:"apiOperationIdSelector,omitempty" tf:"-"`

	// The display name of the API Management API Operation Tag.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

// APIOperationTagSpec defines the desired state of APIOperationTag
type APIOperationTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIOperationTagParameters `json:"forProvider"`
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
	InitProvider APIOperationTagInitParameters `json:"initProvider,omitempty"`
}

// APIOperationTagStatus defines the observed state of APIOperationTag.
type APIOperationTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIOperationTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// APIOperationTag is the Schema for the APIOperationTags API. Manages a API Management API Operation Tag.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIOperationTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   APIOperationTagSpec   `json:"spec"`
	Status APIOperationTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationTagList contains a list of APIOperationTags
type APIOperationTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIOperationTag `json:"items"`
}

// Repository type metadata.
var (
	APIOperationTag_Kind             = "APIOperationTag"
	APIOperationTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIOperationTag_Kind}.String()
	APIOperationTag_KindAPIVersion   = APIOperationTag_Kind + "." + CRDGroupVersion.String()
	APIOperationTag_GroupVersionKind = CRDGroupVersion.WithKind(APIOperationTag_Kind)
)

func init() {
	SchemeBuilder.Register(&APIOperationTag{}, &APIOperationTagList{})
}
