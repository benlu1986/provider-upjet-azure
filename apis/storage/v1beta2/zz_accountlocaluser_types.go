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

type AccountLocalUserInitParameters struct {

	// The home directory of the Storage Account Local User.
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`

	// One or more permission_scope blocks as defined below.
	PermissionScope []PermissionScopeInitParameters `json:"permissionScope,omitempty" tf:"permission_scope,omitempty"`

	// One or more ssh_authorized_key blocks as defined below.
	SSHAuthorizedKey []SSHAuthorizedKeyInitParameters `json:"sshAuthorizedKey,omitempty" tf:"ssh_authorized_key,omitempty"`

	// Specifies whether SSH Key Authentication is enabled. Defaults to false.
	SSHKeyEnabled *bool `json:"sshKeyEnabled,omitempty" tf:"ssh_key_enabled,omitempty"`

	// Specifies whether SSH Password Authentication is enabled. Defaults to false.
	SSHPasswordEnabled *bool `json:"sshPasswordEnabled,omitempty" tf:"ssh_password_enabled,omitempty"`
}

type AccountLocalUserObservation struct {

	// The home directory of the Storage Account Local User.
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`

	// The ID of the Storage Account Local User.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more permission_scope blocks as defined below.
	PermissionScope []PermissionScopeObservation `json:"permissionScope,omitempty" tf:"permission_scope,omitempty"`

	// One or more ssh_authorized_key blocks as defined below.
	SSHAuthorizedKey []SSHAuthorizedKeyObservation `json:"sshAuthorizedKey,omitempty" tf:"ssh_authorized_key,omitempty"`

	// Specifies whether SSH Key Authentication is enabled. Defaults to false.
	SSHKeyEnabled *bool `json:"sshKeyEnabled,omitempty" tf:"ssh_key_enabled,omitempty"`

	// Specifies whether SSH Password Authentication is enabled. Defaults to false.
	SSHPasswordEnabled *bool `json:"sshPasswordEnabled,omitempty" tf:"ssh_password_enabled,omitempty"`

	// The ID of the Storage Account that this Storage Account Local User resides in. Changing this forces a new Storage Account Local User to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type AccountLocalUserParameters struct {

	// The home directory of the Storage Account Local User.
	// +kubebuilder:validation:Optional
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`

	// One or more permission_scope blocks as defined below.
	// +kubebuilder:validation:Optional
	PermissionScope []PermissionScopeParameters `json:"permissionScope,omitempty" tf:"permission_scope,omitempty"`

	// One or more ssh_authorized_key blocks as defined below.
	// +kubebuilder:validation:Optional
	SSHAuthorizedKey []SSHAuthorizedKeyParameters `json:"sshAuthorizedKey,omitempty" tf:"ssh_authorized_key,omitempty"`

	// Specifies whether SSH Key Authentication is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	SSHKeyEnabled *bool `json:"sshKeyEnabled,omitempty" tf:"ssh_key_enabled,omitempty"`

	// Specifies whether SSH Password Authentication is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	SSHPasswordEnabled *bool `json:"sshPasswordEnabled,omitempty" tf:"ssh_password_enabled,omitempty"`

	// The ID of the Storage Account that this Storage Account Local User resides in. Changing this forces a new Storage Account Local User to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type PermissionScopeInitParameters struct {

	// A permissions block as defined below.
	Permissions *PermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The container name (when service is set to blob) or the file share name (when service is set to file), used by the Storage Account Local User.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Reference to a Container in storage to populate resourceName.
	// +kubebuilder:validation:Optional
	ResourceNameRef *v1.Reference `json:"resourceNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate resourceName.
	// +kubebuilder:validation:Optional
	ResourceNameSelector *v1.Selector `json:"resourceNameSelector,omitempty" tf:"-"`

	// The storage service used by this Storage Account Local User. Possible values are blob and file.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type PermissionScopeObservation struct {

	// A permissions block as defined below.
	Permissions *PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The container name (when service is set to blob) or the file share name (when service is set to file), used by the Storage Account Local User.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// The storage service used by this Storage Account Local User. Possible values are blob and file.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type PermissionScopeParameters struct {

	// A permissions block as defined below.
	// +kubebuilder:validation:Optional
	Permissions *PermissionsParameters `json:"permissions" tf:"permissions,omitempty"`

	// The container name (when service is set to blob) or the file share name (when service is set to file), used by the Storage Account Local User.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Reference to a Container in storage to populate resourceName.
	// +kubebuilder:validation:Optional
	ResourceNameRef *v1.Reference `json:"resourceNameRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate resourceName.
	// +kubebuilder:validation:Optional
	ResourceNameSelector *v1.Selector `json:"resourceNameSelector,omitempty" tf:"-"`

	// The storage service used by this Storage Account Local User. Possible values are blob and file.
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`
}

type PermissionsInitParameters struct {

	// (Defaults to 30 minutes) Used when creating the Storage Account Local User.
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// (Defaults to 30 minutes) Used when deleting the Storage Account Local User.
	Delete *bool `json:"delete,omitempty" tf:"delete,omitempty"`

	// Specifies if the Local User has the list permission for this scope. Defaults to false.
	List *bool `json:"list,omitempty" tf:"list,omitempty"`

	// (Defaults to 5 minutes) Used when retrieving the Storage Account Local User.
	Read *bool `json:"read,omitempty" tf:"read,omitempty"`

	// Specifies if the Local User has the write permission for this scope. Defaults to false.
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

type PermissionsObservation struct {

	// (Defaults to 30 minutes) Used when creating the Storage Account Local User.
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// (Defaults to 30 minutes) Used when deleting the Storage Account Local User.
	Delete *bool `json:"delete,omitempty" tf:"delete,omitempty"`

	// Specifies if the Local User has the list permission for this scope. Defaults to false.
	List *bool `json:"list,omitempty" tf:"list,omitempty"`

	// (Defaults to 5 minutes) Used when retrieving the Storage Account Local User.
	Read *bool `json:"read,omitempty" tf:"read,omitempty"`

	// Specifies if the Local User has the write permission for this scope. Defaults to false.
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

type PermissionsParameters struct {

	// (Defaults to 30 minutes) Used when creating the Storage Account Local User.
	// +kubebuilder:validation:Optional
	Create *bool `json:"create,omitempty" tf:"create,omitempty"`

	// (Defaults to 30 minutes) Used when deleting the Storage Account Local User.
	// +kubebuilder:validation:Optional
	Delete *bool `json:"delete,omitempty" tf:"delete,omitempty"`

	// Specifies if the Local User has the list permission for this scope. Defaults to false.
	// +kubebuilder:validation:Optional
	List *bool `json:"list,omitempty" tf:"list,omitempty"`

	// (Defaults to 5 minutes) Used when retrieving the Storage Account Local User.
	// +kubebuilder:validation:Optional
	Read *bool `json:"read,omitempty" tf:"read,omitempty"`

	// Specifies if the Local User has the write permission for this scope. Defaults to false.
	// +kubebuilder:validation:Optional
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

type SSHAuthorizedKeyInitParameters struct {

	// The description of this SSH authorized key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The public key value of this SSH authorized key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type SSHAuthorizedKeyObservation struct {

	// The description of this SSH authorized key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The public key value of this SSH authorized key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type SSHAuthorizedKeyParameters struct {

	// The description of this SSH authorized key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The public key value of this SSH authorized key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`
}

// AccountLocalUserSpec defines the desired state of AccountLocalUser
type AccountLocalUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountLocalUserParameters `json:"forProvider"`
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
	InitProvider AccountLocalUserInitParameters `json:"initProvider,omitempty"`
}

// AccountLocalUserStatus defines the observed state of AccountLocalUser.
type AccountLocalUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountLocalUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AccountLocalUser is the Schema for the AccountLocalUsers API. Manages a Storage Account Local User.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AccountLocalUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountLocalUserSpec   `json:"spec"`
	Status            AccountLocalUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountLocalUserList contains a list of AccountLocalUsers
type AccountLocalUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountLocalUser `json:"items"`
}

// Repository type metadata.
var (
	AccountLocalUser_Kind             = "AccountLocalUser"
	AccountLocalUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountLocalUser_Kind}.String()
	AccountLocalUser_KindAPIVersion   = AccountLocalUser_Kind + "." + CRDGroupVersion.String()
	AccountLocalUser_GroupVersionKind = CRDGroupVersion.WithKind(AccountLocalUser_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountLocalUser{}, &AccountLocalUserList{})
}
