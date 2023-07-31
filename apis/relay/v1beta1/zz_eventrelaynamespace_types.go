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

type EventRelayNamespaceInitParameters struct {

	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the SKU to use. At this time the only supported value is Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventRelayNamespaceObservation struct {

	// The Azure Relay Namespace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Identifier for Azure Insights metrics.
	MetricID *string `json:"metricId,omitempty" tf:"metric_id,omitempty"`

	// The name of the resource group in which to create the Azure Relay Namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the SKU to use. At this time the only supported value is Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventRelayNamespaceParameters struct {

	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Azure Relay Namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SKU to use. At this time the only supported value is Standard.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventRelayNamespaceSpec defines the desired state of EventRelayNamespace
type EventRelayNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventRelayNamespaceParameters `json:"forProvider"`
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
	InitProvider EventRelayNamespaceInitParameters `json:"initProvider,omitempty"`
}

// EventRelayNamespaceStatus defines the observed state of EventRelayNamespace.
type EventRelayNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventRelayNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventRelayNamespace is the Schema for the EventRelayNamespaces API. Manages an Azure Relay Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventRelayNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	Spec   EventRelayNamespaceSpec   `json:"spec"`
	Status EventRelayNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventRelayNamespaceList contains a list of EventRelayNamespaces
type EventRelayNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventRelayNamespace `json:"items"`
}

// Repository type metadata.
var (
	EventRelayNamespace_Kind             = "EventRelayNamespace"
	EventRelayNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventRelayNamespace_Kind}.String()
	EventRelayNamespace_KindAPIVersion   = EventRelayNamespace_Kind + "." + CRDGroupVersion.String()
	EventRelayNamespace_GroupVersionKind = CRDGroupVersion.WithKind(EventRelayNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&EventRelayNamespace{}, &EventRelayNamespaceList{})
}
