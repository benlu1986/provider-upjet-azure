// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomHeaderInitParameters struct {

	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderObservation struct {

	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderParameters struct {

	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DNSConfigInitParameters struct {

	// The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
	RelativeName *string `json:"relativeName,omitempty" tf:"relative_name,omitempty"`

	// The TTL value of the Profile used by Local DNS resolvers and clients.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type DNSConfigObservation struct {

	// The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
	RelativeName *string `json:"relativeName,omitempty" tf:"relative_name,omitempty"`

	// The TTL value of the Profile used by Local DNS resolvers and clients.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type DNSConfigParameters struct {

	// The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RelativeName *string `json:"relativeName" tf:"relative_name,omitempty"`

	// The TTL value of the Profile used by Local DNS resolvers and clients.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`
}

type MonitorConfigInitParameters struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []CustomHeaderInitParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// A list of status code ranges in the format of 100-101.
	ExpectedStatusCodeRanges []*string `json:"expectedStatusCodeRanges,omitempty" tf:"expected_status_code_ranges,omitempty"`

	// The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: 30 (normal probing) and 10 (fast probing). The default value is 30.
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// The path used by the monitoring checks. Required when protocol is set to HTTP or HTTPS - cannot be set when protocol is set to TCP.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port number used by the monitoring checks.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol used by the monitoring checks, supported values are HTTP, HTTPS and TCP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If interval_in_seconds is set to 30, then timeout_in_seconds can be between 5 and 10. The default value is 10. If interval_in_seconds is set to 10, then valid values are between 5 and 9 and timeout_in_seconds is required.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between 0 and 9. The default value is 3
	ToleratedNumberOfFailures *float64 `json:"toleratedNumberOfFailures,omitempty" tf:"tolerated_number_of_failures,omitempty"`
}

type MonitorConfigObservation struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []CustomHeaderObservation `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// A list of status code ranges in the format of 100-101.
	ExpectedStatusCodeRanges []*string `json:"expectedStatusCodeRanges,omitempty" tf:"expected_status_code_ranges,omitempty"`

	// The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: 30 (normal probing) and 10 (fast probing). The default value is 30.
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// The path used by the monitoring checks. Required when protocol is set to HTTP or HTTPS - cannot be set when protocol is set to TCP.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port number used by the monitoring checks.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol used by the monitoring checks, supported values are HTTP, HTTPS and TCP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If interval_in_seconds is set to 30, then timeout_in_seconds can be between 5 and 10. The default value is 10. If interval_in_seconds is set to 10, then valid values are between 5 and 9 and timeout_in_seconds is required.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between 0 and 9. The default value is 3
	ToleratedNumberOfFailures *float64 `json:"toleratedNumberOfFailures,omitempty" tf:"tolerated_number_of_failures,omitempty"`
}

type MonitorConfigParameters struct {

	// One or more custom_header blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// A list of status code ranges in the format of 100-101.
	// +kubebuilder:validation:Optional
	ExpectedStatusCodeRanges []*string `json:"expectedStatusCodeRanges,omitempty" tf:"expected_status_code_ranges,omitempty"`

	// The interval used to check the endpoint health from a Traffic Manager probing agent. You can specify two values here: 30 (normal probing) and 10 (fast probing). The default value is 30.
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// The path used by the monitoring checks. Required when protocol is set to HTTP or HTTPS - cannot be set when protocol is set to TCP.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port number used by the monitoring checks.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The protocol used by the monitoring checks, supported values are HTTP, HTTPS and TCP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint. If interval_in_seconds is set to 30, then timeout_in_seconds can be between 5 and 10. The default value is 10. If interval_in_seconds is set to 10, then valid values are between 5 and 9 and timeout_in_seconds is required.
	// +kubebuilder:validation:Optional
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. Valid values are between 0 and 9. The default value is 3
	// +kubebuilder:validation:Optional
	ToleratedNumberOfFailures *float64 `json:"toleratedNumberOfFailures,omitempty" tf:"tolerated_number_of_failures,omitempty"`
}

type TrafficManagerProfileInitParameters struct {

	// This block specifies the DNS configuration of the Profile. One dns_config block as defined below.
	DNSConfig *DNSConfigInitParameters `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from 1 to 8.
	MaxReturn *float64 `json:"maxReturn,omitempty" tf:"max_return,omitempty"`

	// This block specifies the Endpoint monitoring configuration for the Profile. One monitor_config block as defined below.
	MonitorConfig *MonitorConfigInitParameters `json:"monitorConfig,omitempty" tf:"monitor_config,omitempty"`

	// The status of the profile, can be set to either Enabled or Disabled. Defaults to Enabled.
	ProfileStatus *string `json:"profileStatus,omitempty" tf:"profile_status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the algorithm used to route traffic. Possible values are Geographic, Weighted, Performance, Priority, Subnet and MultiValue.
	TrafficRoutingMethod *string `json:"trafficRoutingMethod,omitempty" tf:"traffic_routing_method,omitempty"`

	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `json:"trafficViewEnabled,omitempty" tf:"traffic_view_enabled,omitempty"`
}

type TrafficManagerProfileObservation struct {

	// This block specifies the DNS configuration of the Profile. One dns_config block as defined below.
	DNSConfig *DNSConfigObservation `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// The FQDN of the created Profile.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The ID of the Traffic Manager Profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from 1 to 8.
	MaxReturn *float64 `json:"maxReturn,omitempty" tf:"max_return,omitempty"`

	// This block specifies the Endpoint monitoring configuration for the Profile. One monitor_config block as defined below.
	MonitorConfig *MonitorConfigObservation `json:"monitorConfig,omitempty" tf:"monitor_config,omitempty"`

	// The status of the profile, can be set to either Enabled or Disabled. Defaults to Enabled.
	ProfileStatus *string `json:"profileStatus,omitempty" tf:"profile_status,omitempty"`

	// The name of the resource group in which to create the Traffic Manager profile. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the algorithm used to route traffic. Possible values are Geographic, Weighted, Performance, Priority, Subnet and MultiValue.
	TrafficRoutingMethod *string `json:"trafficRoutingMethod,omitempty" tf:"traffic_routing_method,omitempty"`

	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	TrafficViewEnabled *bool `json:"trafficViewEnabled,omitempty" tf:"traffic_view_enabled,omitempty"`
}

type TrafficManagerProfileParameters struct {

	// This block specifies the DNS configuration of the Profile. One dns_config block as defined below.
	// +kubebuilder:validation:Optional
	DNSConfig *DNSConfigParameters `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// The amount of endpoints to return for DNS queries to this Profile. Possible values range from 1 to 8.
	// +kubebuilder:validation:Optional
	MaxReturn *float64 `json:"maxReturn,omitempty" tf:"max_return,omitempty"`

	// This block specifies the Endpoint monitoring configuration for the Profile. One monitor_config block as defined below.
	// +kubebuilder:validation:Optional
	MonitorConfig *MonitorConfigParameters `json:"monitorConfig,omitempty" tf:"monitor_config,omitempty"`

	// The status of the profile, can be set to either Enabled or Disabled. Defaults to Enabled.
	// +kubebuilder:validation:Optional
	ProfileStatus *string `json:"profileStatus,omitempty" tf:"profile_status,omitempty"`

	// The name of the resource group in which to create the Traffic Manager profile. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the algorithm used to route traffic. Possible values are Geographic, Weighted, Performance, Priority, Subnet and MultiValue.
	// +kubebuilder:validation:Optional
	TrafficRoutingMethod *string `json:"trafficRoutingMethod,omitempty" tf:"traffic_routing_method,omitempty"`

	// Indicates whether Traffic View is enabled for the Traffic Manager profile.
	// +kubebuilder:validation:Optional
	TrafficViewEnabled *bool `json:"trafficViewEnabled,omitempty" tf:"traffic_view_enabled,omitempty"`
}

// TrafficManagerProfileSpec defines the desired state of TrafficManagerProfile
type TrafficManagerProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficManagerProfileParameters `json:"forProvider"`
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
	InitProvider TrafficManagerProfileInitParameters `json:"initProvider,omitempty"`
}

// TrafficManagerProfileStatus defines the observed state of TrafficManagerProfile.
type TrafficManagerProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficManagerProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TrafficManagerProfile is the Schema for the TrafficManagerProfiles API. Manages a Traffic Manager Profile.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TrafficManagerProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsConfig) || (has(self.initProvider) && has(self.initProvider.dnsConfig))",message="spec.forProvider.dnsConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitorConfig) || (has(self.initProvider) && has(self.initProvider.monitorConfig))",message="spec.forProvider.monitorConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.trafficRoutingMethod) || (has(self.initProvider) && has(self.initProvider.trafficRoutingMethod))",message="spec.forProvider.trafficRoutingMethod is a required parameter"
	Spec   TrafficManagerProfileSpec   `json:"spec"`
	Status TrafficManagerProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerProfileList contains a list of TrafficManagerProfiles
type TrafficManagerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficManagerProfile `json:"items"`
}

// Repository type metadata.
var (
	TrafficManagerProfile_Kind             = "TrafficManagerProfile"
	TrafficManagerProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficManagerProfile_Kind}.String()
	TrafficManagerProfile_KindAPIVersion   = TrafficManagerProfile_Kind + "." + CRDGroupVersion.String()
	TrafficManagerProfile_GroupVersionKind = CRDGroupVersion.WithKind(TrafficManagerProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficManagerProfile{}, &TrafficManagerProfileList{})
}
