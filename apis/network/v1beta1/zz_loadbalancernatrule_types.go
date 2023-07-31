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

type LoadBalancerNatRuleInitParameters struct {

	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	FrontendPortEnd *float64 `json:"frontendPortEnd,omitempty" tf:"frontend_port_end,omitempty"`

	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	FrontendPortStart *float64 `json:"frontendPortStart,omitempty" tf:"frontend_port_start,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30 minutes. Defaults to 4 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Udp, Tcp or All.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type LoadBalancerNatRuleObservation struct {

	// Specifies a reference to backendAddressPool resource.
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	// The ID of the Load Balancer NAT Rule.
	BackendIPConfigurationID *string `json:"backendIpConfigurationId,omitempty" tf:"backend_ip_configuration_id,omitempty"`

	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The ID of the Load Balancer NAT Rule.
	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	FrontendPortEnd *float64 `json:"frontendPortEnd,omitempty" tf:"frontend_port_end,omitempty"`

	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	FrontendPortStart *float64 `json:"frontendPortStart,omitempty" tf:"frontend_port_start,omitempty"`

	// The ID of the Load Balancer NAT Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30 minutes. Defaults to 4 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the NAT Rule. Changing this forces a new resource to be created.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Udp, Tcp or All.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type LoadBalancerNatRuleParameters struct {

	// Specifies a reference to backendAddressPool resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.LoadBalancerBackendAddressPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackendAddressPoolID *string `json:"backendAddressPoolId,omitempty" tf:"backend_address_pool_id,omitempty"`

	// Reference to a LoadBalancerBackendAddressPool in network to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDRef *v1.Reference `json:"backendAddressPoolIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerBackendAddressPool in network to populate backendAddressPoolId.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIDSelector *v1.Selector `json:"backendAddressPoolIdSelector,omitempty" tf:"-"`

	// The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration exposing this rule.
	// +kubebuilder:validation:Optional
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
	// +kubebuilder:validation:Optional
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	// +kubebuilder:validation:Optional
	FrontendPortEnd *float64 `json:"frontendPortEnd,omitempty" tf:"frontend_port_end,omitempty"`

	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534, inclusive.
	// +kubebuilder:validation:Optional
	FrontendPortStart *float64 `json:"frontendPortStart,omitempty" tf:"frontend_port_start,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 30 minutes. Defaults to 4 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// The ID of the Load Balancer in which to create the NAT Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// The transport protocol for the external endpoint. Possible values are Udp, Tcp or All.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// LoadBalancerNatRuleSpec defines the desired state of LoadBalancerNatRule
type LoadBalancerNatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerNatRuleParameters `json:"forProvider"`
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
	InitProvider LoadBalancerNatRuleInitParameters `json:"initProvider,omitempty"`
}

// LoadBalancerNatRuleStatus defines the observed state of LoadBalancerNatRule.
type LoadBalancerNatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerNatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatRule is the Schema for the LoadBalancerNatRules API. Manages a Load Balancer NAT Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerNatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backendPort) || has(self.initProvider.backendPort)",message="backendPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.frontendIpConfigurationName) || has(self.initProvider.frontendIpConfigurationName)",message="frontendIpConfigurationName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || has(self.initProvider.protocol)",message="protocol is a required parameter"
	Spec   LoadBalancerNatRuleSpec   `json:"spec"`
	Status LoadBalancerNatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerNatRuleList contains a list of LoadBalancerNatRules
type LoadBalancerNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerNatRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerNatRule_Kind             = "LoadBalancerNatRule"
	LoadBalancerNatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerNatRule_Kind}.String()
	LoadBalancerNatRule_KindAPIVersion   = LoadBalancerNatRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerNatRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerNatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerNatRule{}, &LoadBalancerNatRuleList{})
}
