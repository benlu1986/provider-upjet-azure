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

type ClusterPrincipalAssignmentInitParameters struct {

	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The type of the principal. Valid values include App, Group, User. Changing this forces a new resource to be created.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The cluster role assigned to the principal. Valid values include AllDatabasesAdmin and AllDatabasesViewer. Changing this forces a new resource to be created.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ClusterPrincipalAssignmentObservation struct {

	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// The ID of the Kusto Cluster Principal Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The name of the principal.
	PrincipalName *string `json:"principalName,omitempty" tf:"principal_name,omitempty"`

	// The type of the principal. Valid values include App, Group, User. Changing this forces a new resource to be created.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The cluster role assigned to the principal. Valid values include AllDatabasesAdmin and AllDatabasesViewer. Changing this forces a new resource to be created.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The name of the tenant.
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name,omitempty"`
}

type ClusterPrincipalAssignmentParameters struct {

	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// The object id of the principal. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The type of the principal. Valid values include App, Group, User. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

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

	// The cluster role assigned to the principal. Valid values include AllDatabasesAdmin and AllDatabasesViewer. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// ClusterPrincipalAssignmentSpec defines the desired state of ClusterPrincipalAssignment
type ClusterPrincipalAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterPrincipalAssignmentParameters `json:"forProvider"`
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
	InitProvider ClusterPrincipalAssignmentInitParameters `json:"initProvider,omitempty"`
}

// ClusterPrincipalAssignmentStatus defines the observed state of ClusterPrincipalAssignment.
type ClusterPrincipalAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterPrincipalAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterPrincipalAssignment is the Schema for the ClusterPrincipalAssignments API. Manages a Kusto Cluster Principal Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ClusterPrincipalAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalId) || has(self.initProvider.principalId)",message="principalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalType) || has(self.initProvider.principalType)",message="principalType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || has(self.initProvider.tenantId)",message="tenantId is a required parameter"
	Spec   ClusterPrincipalAssignmentSpec   `json:"spec"`
	Status ClusterPrincipalAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterPrincipalAssignmentList contains a list of ClusterPrincipalAssignments
type ClusterPrincipalAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPrincipalAssignment `json:"items"`
}

// Repository type metadata.
var (
	ClusterPrincipalAssignment_Kind             = "ClusterPrincipalAssignment"
	ClusterPrincipalAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterPrincipalAssignment_Kind}.String()
	ClusterPrincipalAssignment_KindAPIVersion   = ClusterPrincipalAssignment_Kind + "." + CRDGroupVersion.String()
	ClusterPrincipalAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ClusterPrincipalAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterPrincipalAssignment{}, &ClusterPrincipalAssignmentList{})
}
