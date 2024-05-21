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

type VirtualHubIPInitParameters struct {

	// The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The private IP address allocation method. Possible values are Static and Dynamic is allowed. Defaults to Dynamic.
	PrivateIPAllocationMethod *string `json:"privateIpAllocationMethod,omitempty" tf:"private_ip_allocation_method,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualHubIPObservation struct {

	// The ID of the Virtual Hub IP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The private IP address allocation method. Possible values are Static and Dynamic is allowed. Defaults to Dynamic.
	PrivateIPAllocationMethod *string `json:"privateIpAllocationMethod,omitempty" tf:"private_ip_allocation_method,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the Virtual Hub within which this IP configuration should be created. Changing this forces a new resource to be created.
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`
}

type VirtualHubIPParameters struct {

	// The private IP address of the IP configuration.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The private IP address allocation method. Possible values are Static and Dynamic is allowed. Defaults to Dynamic.
	// +kubebuilder:validation:Optional
	PrivateIPAllocationMethod *string `json:"privateIpAllocationMethod,omitempty" tf:"private_ip_allocation_method,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet that the IP will reside. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Hub within which this IP configuration should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// Reference to a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

// VirtualHubIPSpec defines the desired state of VirtualHubIP
type VirtualHubIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubIPParameters `json:"forProvider"`
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
	InitProvider VirtualHubIPInitParameters `json:"initProvider,omitempty"`
}

// VirtualHubIPStatus defines the observed state of VirtualHubIP.
type VirtualHubIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualHubIP is the Schema for the VirtualHubIPs API. Manages a Virtual Hub IP. This resource is also known as a Route Server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubIPSpec   `json:"spec"`
	Status            VirtualHubIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubIPList contains a list of VirtualHubIPs
type VirtualHubIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubIP `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubIP_Kind             = "VirtualHubIP"
	VirtualHubIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualHubIP_Kind}.String()
	VirtualHubIP_KindAPIVersion   = VirtualHubIP_Kind + "." + CRDGroupVersion.String()
	VirtualHubIP_GroupVersionKind = CRDGroupVersion.WithKind(VirtualHubIP_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubIP{}, &VirtualHubIPList{})
}
