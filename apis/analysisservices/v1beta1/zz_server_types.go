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

type IPv4FirewallRuleInitParameters struct {

	// Specifies the name of the firewall rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// End of the firewall rule range as IPv4 address.
	RangeEnd *string `json:"rangeEnd,omitempty" tf:"range_end,omitempty"`

	// Start of the firewall rule range as IPv4 address.
	RangeStart *string `json:"rangeStart,omitempty" tf:"range_start,omitempty"`
}

type IPv4FirewallRuleObservation struct {

	// Specifies the name of the firewall rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// End of the firewall rule range as IPv4 address.
	RangeEnd *string `json:"rangeEnd,omitempty" tf:"range_end,omitempty"`

	// Start of the firewall rule range as IPv4 address.
	RangeStart *string `json:"rangeStart,omitempty" tf:"range_start,omitempty"`
}

type IPv4FirewallRuleParameters struct {

	// Specifies the name of the firewall rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// End of the firewall rule range as IPv4 address.
	// +kubebuilder:validation:Optional
	RangeEnd *string `json:"rangeEnd" tf:"range_end,omitempty"`

	// Start of the firewall rule range as IPv4 address.
	// +kubebuilder:validation:Optional
	RangeStart *string `json:"rangeStart" tf:"range_start,omitempty"`
}

type ServerInitParameters struct {

	// List of email addresses of admin users.
	AdminUsers []*string `json:"adminUsers,omitempty" tf:"admin_users,omitempty"`

	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService *bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service,omitempty"`

	// One or more ipv4_firewall_rule block(s) as defined below.
	IPv4FirewallRule []IPv4FirewallRuleInitParameters `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule,omitempty"`

	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Analysis Services Server. Only lowercase Alphanumeric characters allowed, starting with a letter. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Controls how the read-write server is used in the query pool. If this value is set to All then read-write servers are also used for queries. Otherwise with ReadOnly these servers do not participate in query operations.
	QuerypoolConnectionMode *string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode,omitempty"`

	// SKU for the Analysis Services Server. Possible values are: D1, B1, B2, S0, S1, S2, S4, S8, S9, S8v2 and S9v2.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServerObservation struct {

	// List of email addresses of admin users.
	AdminUsers []*string `json:"adminUsers,omitempty" tf:"admin_users,omitempty"`

	// Indicates if the Power BI service is allowed to access or not.
	EnablePowerBiService *bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service,omitempty"`

	// The ID of the Analysis Services Server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more ipv4_firewall_rule block(s) as defined below.
	IPv4FirewallRule []IPv4FirewallRuleObservation `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule,omitempty"`

	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Analysis Services Server. Only lowercase Alphanumeric characters allowed, starting with a letter. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Controls how the read-write server is used in the query pool. If this value is set to All then read-write servers are also used for queries. Otherwise with ReadOnly these servers do not participate in query operations.
	QuerypoolConnectionMode *string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode,omitempty"`

	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The full name of the Analysis Services Server.
	ServerFullName *string `json:"serverFullName,omitempty" tf:"server_full_name,omitempty"`

	// SKU for the Analysis Services Server. Possible values are: D1, B1, B2, S0, S1, S2, S4, S8, S9, S8v2 and S9v2.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServerParameters struct {

	// List of email addresses of admin users.
	// +kubebuilder:validation:Optional
	AdminUsers []*string `json:"adminUsers,omitempty" tf:"admin_users,omitempty"`

	// URI and SAS token for a blob container to store backups.
	// +kubebuilder:validation:Optional
	BackupBlobContainerURISecretRef *v1.SecretKeySelector `json:"backupBlobContainerUriSecretRef,omitempty" tf:"-"`

	// Indicates if the Power BI service is allowed to access or not.
	// +kubebuilder:validation:Optional
	EnablePowerBiService *bool `json:"enablePowerBiService,omitempty" tf:"enable_power_bi_service,omitempty"`

	// One or more ipv4_firewall_rule block(s) as defined below.
	// +kubebuilder:validation:Optional
	IPv4FirewallRule []IPv4FirewallRuleParameters `json:"ipv4FirewallRule,omitempty" tf:"ipv4_firewall_rule,omitempty"`

	// The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Analysis Services Server. Only lowercase Alphanumeric characters allowed, starting with a letter. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Controls how the read-write server is used in the query pool. If this value is set to All then read-write servers are also used for queries. Otherwise with ReadOnly these servers do not participate in query operations.
	// +kubebuilder:validation:Optional
	QuerypoolConnectionMode *string `json:"querypoolConnectionMode,omitempty" tf:"querypool_connection_mode,omitempty"`

	// The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// SKU for the Analysis Services Server. Possible values are: D1, B1, B2, S0, S1, S2, S4, S8, S9, S8v2 and S9v2.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
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
	InitProvider ServerInitParameters `json:"initProvider,omitempty"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Manages an Analysis Services Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || has(self.initProvider.sku)",message="sku is a required parameter"
	Spec   ServerSpec   `json:"spec"`
	Status ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
