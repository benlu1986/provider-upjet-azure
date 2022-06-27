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

type NetworkInterfaceNatRuleAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkInterfaceNatRuleAssociationParameters struct {

	// +kubebuilder:validation:Required
	IPConfigurationName *string `json:"ipConfigurationName" tf:"ip_configuration_name,omitempty"`

	// +crossplane:generate:reference:type=LoadBalancerNatRule
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NATRuleID *string `json:"natRuleId,omitempty" tf:"nat_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	NATRuleIDRef *v1.Reference `json:"natRuleIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NATRuleIDSelector *v1.Selector `json:"natRuleIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceNatRuleAssociationSpec defines the desired state of NetworkInterfaceNatRuleAssociation
type NetworkInterfaceNatRuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceNatRuleAssociationParameters `json:"forProvider"`
}

// NetworkInterfaceNatRuleAssociationStatus defines the observed state of NetworkInterfaceNatRuleAssociation.
type NetworkInterfaceNatRuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceNatRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociation is the Schema for the NetworkInterfaceNatRuleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceNatRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceNatRuleAssociationSpec   `json:"spec"`
	Status            NetworkInterfaceNatRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociationList contains a list of NetworkInterfaceNatRuleAssociations
type NetworkInterfaceNatRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceNatRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceNatRuleAssociation_Kind             = "NetworkInterfaceNatRuleAssociation"
	NetworkInterfaceNatRuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceNatRuleAssociation_Kind}.String()
	NetworkInterfaceNatRuleAssociation_KindAPIVersion   = NetworkInterfaceNatRuleAssociation_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceNatRuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceNatRuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceNatRuleAssociation{}, &NetworkInterfaceNatRuleAssociationList{})
}
