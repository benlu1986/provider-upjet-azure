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

type CustomHeaderInitParameters struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderObservation struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderParameters struct {

	// The name of the custom header.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type TrafficManagerAzureEndpointInitParameters struct {

	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method. Defaults to false.
	AlwaysServeEnabled *bool `json:"alwaysServeEnabled,omitempty" tf:"always_serve_enabled,omitempty"`

	// One or more custom_header blocks as defined below.
	CustomHeader []CustomHeaderInitParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerAzureEndpointSubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The ID of the Azure Resource which should be used as a target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a PublicIP in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerAzureEndpointObservation struct {

	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method. Defaults to false.
	AlwaysServeEnabled *bool `json:"alwaysServeEnabled,omitempty" tf:"always_serve_enabled,omitempty"`

	// One or more custom_header blocks as defined below.
	CustomHeader []CustomHeaderObservation `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// The ID of the Azure Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the Traffic Manager Profile that this Azure Endpoint should be created within. Changing this forces a new resource to be created.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerAzureEndpointSubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The ID of the Azure Resource which should be used as a target.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerAzureEndpointParameters struct {

	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method. Defaults to false.
	// +kubebuilder:validation:Optional
	AlwaysServeEnabled *bool `json:"alwaysServeEnabled,omitempty" tf:"always_serve_enabled,omitempty"`

	// One or more custom_header blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	// +kubebuilder:validation:Optional
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the Traffic Manager Profile that this Azure Endpoint should be created within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.TrafficManagerProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// Reference to a TrafficManagerProfile in network to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDRef *v1.Reference `json:"profileIdRef,omitempty" tf:"-"`

	// Selector for a TrafficManagerProfile in network to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDSelector *v1.Selector `json:"profileIdSelector,omitempty" tf:"-"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subnet []TrafficManagerAzureEndpointSubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The ID of the Azure Resource which should be used as a target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a PublicIP in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerAzureEndpointSubnetInitParameters struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerAzureEndpointSubnetObservation struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerAzureEndpointSubnetParameters struct {

	// The first IP Address in this subnet.
	// +kubebuilder:validation:Optional
	First *string `json:"first" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	// +kubebuilder:validation:Optional
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	// +kubebuilder:validation:Optional
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

// TrafficManagerAzureEndpointSpec defines the desired state of TrafficManagerAzureEndpoint
type TrafficManagerAzureEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficManagerAzureEndpointParameters `json:"forProvider"`
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
	InitProvider TrafficManagerAzureEndpointInitParameters `json:"initProvider,omitempty"`
}

// TrafficManagerAzureEndpointStatus defines the observed state of TrafficManagerAzureEndpoint.
type TrafficManagerAzureEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficManagerAzureEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrafficManagerAzureEndpoint is the Schema for the TrafficManagerAzureEndpoints API. Manages an Azure Endpoint within a Traffic Manager Profile..
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TrafficManagerAzureEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficManagerAzureEndpointSpec   `json:"spec"`
	Status            TrafficManagerAzureEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerAzureEndpointList contains a list of TrafficManagerAzureEndpoints
type TrafficManagerAzureEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficManagerAzureEndpoint `json:"items"`
}

// Repository type metadata.
var (
	TrafficManagerAzureEndpoint_Kind             = "TrafficManagerAzureEndpoint"
	TrafficManagerAzureEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficManagerAzureEndpoint_Kind}.String()
	TrafficManagerAzureEndpoint_KindAPIVersion   = TrafficManagerAzureEndpoint_Kind + "." + CRDGroupVersion.String()
	TrafficManagerAzureEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(TrafficManagerAzureEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficManagerAzureEndpoint{}, &TrafficManagerAzureEndpointList{})
}
