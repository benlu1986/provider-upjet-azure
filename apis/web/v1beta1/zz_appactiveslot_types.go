/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AppActiveSlotInitParameters struct {

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`
}

type AppActiveSlotObservation struct {

	// The ID of the Web App Active Slot
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the last successful swap with Production.
	// The timestamp of the last successful swap with `Production`
	LastSuccessfulSwap *string `json:"lastSuccessfulSwap,omitempty" tf:"last_successful_swap,omitempty"`

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`

	// The ID of the Slot to swap with Production.
	// The ID of the Slot to swap with `Production`.
	SlotID *string `json:"slotId,omitempty" tf:"slot_id,omitempty"`
}

type AppActiveSlotParameters struct {

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	// +kubebuilder:validation:Optional
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`

	// The ID of the Slot to swap with Production.
	// The ID of the Slot to swap with `Production`.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsWebAppSlot
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SlotID *string `json:"slotId,omitempty" tf:"slot_id,omitempty"`

	// Reference to a WindowsWebAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDRef *v1.Reference `json:"slotIdRef,omitempty" tf:"-"`

	// Selector for a WindowsWebAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDSelector *v1.Selector `json:"slotIdSelector,omitempty" tf:"-"`
}

// AppActiveSlotSpec defines the desired state of AppActiveSlot
type AppActiveSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppActiveSlotParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AppActiveSlotInitParameters `json:"initProvider,omitempty"`
}

// AppActiveSlotStatus defines the observed state of AppActiveSlot.
type AppActiveSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppActiveSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppActiveSlot is the Schema for the AppActiveSlots API. Manages a Web App Active Slot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppActiveSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppActiveSlotSpec   `json:"spec"`
	Status            AppActiveSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppActiveSlotList contains a list of AppActiveSlots
type AppActiveSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppActiveSlot `json:"items"`
}

// Repository type metadata.
var (
	AppActiveSlot_Kind             = "AppActiveSlot"
	AppActiveSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppActiveSlot_Kind}.String()
	AppActiveSlot_KindAPIVersion   = AppActiveSlot_Kind + "." + CRDGroupVersion.String()
	AppActiveSlot_GroupVersionKind = CRDGroupVersion.WithKind(AppActiveSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&AppActiveSlot{}, &AppActiveSlotList{})
}
