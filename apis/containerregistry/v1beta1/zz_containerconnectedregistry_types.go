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

type ContainerConnectedRegistryInitParameters struct {

	// Should the log auditing be enabled?
	AuditLogEnabled *bool `json:"auditLogEnabled,omitempty" tf:"audit_log_enabled,omitempty"`

	// Specifies a list of IDs of Container Registry Tokens, which are meant to be used by the clients to connect to the Connected Registry.
	ClientTokenIds []*string `json:"clientTokenIds,omitempty" tf:"client_token_ids,omitempty"`

	// The verbosity of the logs. Possible values are None, Debug, Information, Warning and Error.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// The mode of the Connected Registry. Possible values are Mirror, ReadOnly, ReadWrite and Registry. Changing this forces a new Container Connected Registry to be created.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// One or more notification blocks as defined below.
	Notification []NotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the parent registry. This can be either a Container Registry ID or a Connected Registry ID. Changing this forces a new Container Connected Registry to be created.
	ParentRegistryID *string `json:"parentRegistryId,omitempty" tf:"parent_registry_id,omitempty"`

	// The period of time (in form of ISO8601) for which a message is available to sync before it is expired. Allowed range is from P1D to P90D.
	SyncMessageTTL *string `json:"syncMessageTtl,omitempty" tf:"sync_message_ttl,omitempty"`

	// The cron expression indicating the schedule that the Connected Registry will sync with its parent. Defaults to * * * * *.
	SyncSchedule *string `json:"syncSchedule,omitempty" tf:"sync_schedule,omitempty"`

	// The time window (in form of ISO8601) during which sync is enabled for each schedule occurrence. Allowed range is from PT3H to P7D.
	SyncWindow *string `json:"syncWindow,omitempty" tf:"sync_window,omitempty"`
}

type ContainerConnectedRegistryObservation struct {

	// Should the log auditing be enabled?
	AuditLogEnabled *bool `json:"auditLogEnabled,omitempty" tf:"audit_log_enabled,omitempty"`

	// Specifies a list of IDs of Container Registry Tokens, which are meant to be used by the clients to connect to the Connected Registry.
	ClientTokenIds []*string `json:"clientTokenIds,omitempty" tf:"client_token_ids,omitempty"`

	// The ID of the Container Registry that this Connected Registry will reside in. Changing this forces a new Container Connected Registry to be created.
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// The ID of the Container Connected Registry.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The verbosity of the logs. Possible values are None, Debug, Information, Warning and Error.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// The mode of the Connected Registry. Possible values are Mirror, ReadOnly, ReadWrite and Registry. Changing this forces a new Container Connected Registry to be created.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// One or more notification blocks as defined below.
	Notification []NotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the parent registry. This can be either a Container Registry ID or a Connected Registry ID. Changing this forces a new Container Connected Registry to be created.
	ParentRegistryID *string `json:"parentRegistryId,omitempty" tf:"parent_registry_id,omitempty"`

	// The period of time (in form of ISO8601) for which a message is available to sync before it is expired. Allowed range is from P1D to P90D.
	SyncMessageTTL *string `json:"syncMessageTtl,omitempty" tf:"sync_message_ttl,omitempty"`

	// The cron expression indicating the schedule that the Connected Registry will sync with its parent. Defaults to * * * * *.
	SyncSchedule *string `json:"syncSchedule,omitempty" tf:"sync_schedule,omitempty"`

	// The ID of the Container Registry Token which is used for synchronizing the Connected Registry. Changing this forces a new Container Connected Registry to be created.
	SyncTokenID *string `json:"syncTokenId,omitempty" tf:"sync_token_id,omitempty"`

	// The time window (in form of ISO8601) during which sync is enabled for each schedule occurrence. Allowed range is from PT3H to P7D.
	SyncWindow *string `json:"syncWindow,omitempty" tf:"sync_window,omitempty"`
}

type ContainerConnectedRegistryParameters struct {

	// Should the log auditing be enabled?
	// +kubebuilder:validation:Optional
	AuditLogEnabled *bool `json:"auditLogEnabled,omitempty" tf:"audit_log_enabled,omitempty"`

	// Specifies a list of IDs of Container Registry Tokens, which are meant to be used by the clients to connect to the Connected Registry.
	// +kubebuilder:validation:Optional
	ClientTokenIds []*string `json:"clientTokenIds,omitempty" tf:"client_token_ids,omitempty"`

	// The ID of the Container Registry that this Connected Registry will reside in. Changing this forces a new Container Connected Registry to be created.
	// +crossplane:generate:reference:type=Registry
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// Reference to a Registry to populate containerRegistryId.
	// +kubebuilder:validation:Optional
	ContainerRegistryIDRef *v1.Reference `json:"containerRegistryIdRef,omitempty" tf:"-"`

	// Selector for a Registry to populate containerRegistryId.
	// +kubebuilder:validation:Optional
	ContainerRegistryIDSelector *v1.Selector `json:"containerRegistryIdSelector,omitempty" tf:"-"`

	// The verbosity of the logs. Possible values are None, Debug, Information, Warning and Error.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// The mode of the Connected Registry. Possible values are Mirror, ReadOnly, ReadWrite and Registry. Changing this forces a new Container Connected Registry to be created.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// One or more notification blocks as defined below.
	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the parent registry. This can be either a Container Registry ID or a Connected Registry ID. Changing this forces a new Container Connected Registry to be created.
	// +kubebuilder:validation:Optional
	ParentRegistryID *string `json:"parentRegistryId,omitempty" tf:"parent_registry_id,omitempty"`

	// The period of time (in form of ISO8601) for which a message is available to sync before it is expired. Allowed range is from P1D to P90D.
	// +kubebuilder:validation:Optional
	SyncMessageTTL *string `json:"syncMessageTtl,omitempty" tf:"sync_message_ttl,omitempty"`

	// The cron expression indicating the schedule that the Connected Registry will sync with its parent. Defaults to * * * * *.
	// +kubebuilder:validation:Optional
	SyncSchedule *string `json:"syncSchedule,omitempty" tf:"sync_schedule,omitempty"`

	// The ID of the Container Registry Token which is used for synchronizing the Connected Registry. Changing this forces a new Container Connected Registry to be created.
	// +crossplane:generate:reference:type=Token
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SyncTokenID *string `json:"syncTokenId,omitempty" tf:"sync_token_id,omitempty"`

	// Reference to a Token to populate syncTokenId.
	// +kubebuilder:validation:Optional
	SyncTokenIDRef *v1.Reference `json:"syncTokenIdRef,omitempty" tf:"-"`

	// Selector for a Token to populate syncTokenId.
	// +kubebuilder:validation:Optional
	SyncTokenIDSelector *v1.Selector `json:"syncTokenIdSelector,omitempty" tf:"-"`

	// The time window (in form of ISO8601) during which sync is enabled for each schedule occurrence. Allowed range is from PT3H to P7D.
	// +kubebuilder:validation:Optional
	SyncWindow *string `json:"syncWindow,omitempty" tf:"sync_window,omitempty"`
}

type NotificationInitParameters struct {

	// The action of the artifact that wants to be subscribed for the Connected Registry. Possible values are push, delete and * (i.e. any).
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The digest of the artifact that wants to be subscribed for the Connected Registry.
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// The name of the artifact that wants to be subscribed for the Connected Registry.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tag of the artifact that wants to be subscribed for the Connected Registry.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotificationObservation struct {

	// The action of the artifact that wants to be subscribed for the Connected Registry. Possible values are push, delete and * (i.e. any).
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The digest of the artifact that wants to be subscribed for the Connected Registry.
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// The name of the artifact that wants to be subscribed for the Connected Registry.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tag of the artifact that wants to be subscribed for the Connected Registry.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotificationParameters struct {

	// The action of the artifact that wants to be subscribed for the Connected Registry. Possible values are push, delete and * (i.e. any).
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The digest of the artifact that wants to be subscribed for the Connected Registry.
	// +kubebuilder:validation:Optional
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// The name of the artifact that wants to be subscribed for the Connected Registry.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The tag of the artifact that wants to be subscribed for the Connected Registry.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// ContainerConnectedRegistrySpec defines the desired state of ContainerConnectedRegistry
type ContainerConnectedRegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerConnectedRegistryParameters `json:"forProvider"`
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
	InitProvider ContainerConnectedRegistryInitParameters `json:"initProvider,omitempty"`
}

// ContainerConnectedRegistryStatus defines the observed state of ContainerConnectedRegistry.
type ContainerConnectedRegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerConnectedRegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerConnectedRegistry is the Schema for the ContainerConnectedRegistrys API. Manages a Container Connected Registry.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContainerConnectedRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerConnectedRegistrySpec   `json:"spec"`
	Status            ContainerConnectedRegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerConnectedRegistryList contains a list of ContainerConnectedRegistrys
type ContainerConnectedRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerConnectedRegistry `json:"items"`
}

// Repository type metadata.
var (
	ContainerConnectedRegistry_Kind             = "ContainerConnectedRegistry"
	ContainerConnectedRegistry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerConnectedRegistry_Kind}.String()
	ContainerConnectedRegistry_KindAPIVersion   = ContainerConnectedRegistry_Kind + "." + CRDGroupVersion.String()
	ContainerConnectedRegistry_GroupVersionKind = CRDGroupVersion.WithKind(ContainerConnectedRegistry_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerConnectedRegistry{}, &ContainerConnectedRegistryList{})
}
