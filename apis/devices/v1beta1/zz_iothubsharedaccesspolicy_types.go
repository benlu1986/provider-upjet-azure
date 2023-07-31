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

type IOTHubSharedAccessPolicyInitParameters struct {

	// Adds DeviceConnect permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// Adds RegistryRead permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// Adds RegistryWrite permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// Adds ServiceConnect permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

type IOTHubSharedAccessPolicyObservation struct {

	// Adds DeviceConnect permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// The ID of the IoTHub Shared Access Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Adds RegistryRead permission to this Shared Access Account. It allows read access to the identity registry.
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// Adds RegistryWrite permission to this Shared Access Account. It allows write access to the identity registry.
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Adds ServiceConnect permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

type IOTHubSharedAccessPolicyParameters struct {

	// Adds DeviceConnect permission to this Shared Access Account. It allows sending and receiving on the device-side endpoints.
	// +kubebuilder:validation:Optional
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// The name of the IoTHub to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// Adds RegistryRead permission to this Shared Access Account. It allows read access to the identity registry.
	// +kubebuilder:validation:Optional
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// Adds RegistryWrite permission to this Shared Access Account. It allows write access to the identity registry.
	// +kubebuilder:validation:Optional
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Adds ServiceConnect permission to this Shared Access Account. It allows sending and receiving on the cloud-side endpoints.
	// +kubebuilder:validation:Optional
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

// IOTHubSharedAccessPolicySpec defines the desired state of IOTHubSharedAccessPolicy
type IOTHubSharedAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubSharedAccessPolicyParameters `json:"forProvider"`
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
	InitProvider IOTHubSharedAccessPolicyInitParameters `json:"initProvider,omitempty"`
}

// IOTHubSharedAccessPolicyStatus defines the observed state of IOTHubSharedAccessPolicy.
type IOTHubSharedAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubSharedAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubSharedAccessPolicy is the Schema for the IOTHubSharedAccessPolicys API. Manages an IotHub Shared Access Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubSharedAccessPolicySpec   `json:"spec"`
	Status            IOTHubSharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubSharedAccessPolicyList contains a list of IOTHubSharedAccessPolicys
type IOTHubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubSharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IOTHubSharedAccessPolicy_Kind             = "IOTHubSharedAccessPolicy"
	IOTHubSharedAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubSharedAccessPolicy_Kind}.String()
	IOTHubSharedAccessPolicy_KindAPIVersion   = IOTHubSharedAccessPolicy_Kind + "." + CRDGroupVersion.String()
	IOTHubSharedAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubSharedAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubSharedAccessPolicy{}, &IOTHubSharedAccessPolicyList{})
}
