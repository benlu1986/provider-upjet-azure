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

type VirtualHubRouteTableRouteInitParameters_2 struct {

	// A list of destination addresses for this route.
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// The type of destinations. Possible values are CIDR, ResourceId and Service.
	DestinationsType *string `json:"destinationsType,omitempty" tf:"destinations_type,omitempty"`

	// The type of next hop. Currently the only possible value is ResourceId. Defaults to ResourceId.
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type,omitempty"`
}

type VirtualHubRouteTableRouteObservation_2 struct {

	// A list of destination addresses for this route.
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// The type of destinations. Possible values are CIDR, ResourceId and Service.
	DestinationsType *string `json:"destinationsType,omitempty" tf:"destinations_type,omitempty"`

	// The ID of the Virtual Hub Route Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The next hop's resource ID.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The type of next hop. Currently the only possible value is ResourceId. Defaults to ResourceId.
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type,omitempty"`

	// The ID of the Virtual Hub Route Table to link this route to. Changing this forces a new resource to be created.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`
}

type VirtualHubRouteTableRouteParameters_2 struct {

	// A list of destination addresses for this route.
	// +kubebuilder:validation:Optional
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// The type of destinations. Possible values are CIDR, ResourceId and Service.
	// +kubebuilder:validation:Optional
	DestinationsType *string `json:"destinationsType,omitempty" tf:"destinations_type,omitempty"`

	// The next hop's resource ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHubConnection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// Reference to a VirtualHubConnection in network to populate nextHop.
	// +kubebuilder:validation:Optional
	NextHopRef *v1.Reference `json:"nextHopRef,omitempty" tf:"-"`

	// Selector for a VirtualHubConnection in network to populate nextHop.
	// +kubebuilder:validation:Optional
	NextHopSelector *v1.Selector `json:"nextHopSelector,omitempty" tf:"-"`

	// The type of next hop. Currently the only possible value is ResourceId. Defaults to ResourceId.
	// +kubebuilder:validation:Optional
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type,omitempty"`

	// The ID of the Virtual Hub Route Table to link this route to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHubRouteTable
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a VirtualHubRouteTable in network to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHubRouteTable in network to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`
}

// VirtualHubRouteTableRouteSpec defines the desired state of VirtualHubRouteTableRoute
type VirtualHubRouteTableRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubRouteTableRouteParameters_2 `json:"forProvider"`
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
	InitProvider VirtualHubRouteTableRouteInitParameters_2 `json:"initProvider,omitempty"`
}

// VirtualHubRouteTableRouteStatus defines the observed state of VirtualHubRouteTableRoute.
type VirtualHubRouteTableRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubRouteTableRouteObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTableRoute is the Schema for the VirtualHubRouteTableRoutes API. Manages a Route in a Virtual Hub Route Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubRouteTableRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinations) || has(self.initProvider.destinations)",message="destinations is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationsType) || has(self.initProvider.destinationsType)",message="destinationsType is a required parameter"
	Spec   VirtualHubRouteTableRouteSpec   `json:"spec"`
	Status VirtualHubRouteTableRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTableRouteList contains a list of VirtualHubRouteTableRoutes
type VirtualHubRouteTableRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubRouteTableRoute `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubRouteTableRoute_Kind             = "VirtualHubRouteTableRoute"
	VirtualHubRouteTableRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualHubRouteTableRoute_Kind}.String()
	VirtualHubRouteTableRoute_KindAPIVersion   = VirtualHubRouteTableRoute_Kind + "." + CRDGroupVersion.String()
	VirtualHubRouteTableRoute_GroupVersionKind = CRDGroupVersion.WithKind(VirtualHubRouteTableRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubRouteTableRoute{}, &VirtualHubRouteTableRouteList{})
}
