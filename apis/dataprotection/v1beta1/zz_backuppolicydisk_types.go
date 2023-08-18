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

type BackupPolicyDiskInitParameters struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval . Changing this forces a new Backup Policy Disk to be created.
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The duration of default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	DefaultRetentionDuration *string `json:"defaultRetentionDuration,omitempty" tf:"default_retention_duration,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Disk to be created.
	RetentionRule []RetentionRuleInitParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`
}

type BackupPolicyDiskObservation struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval . Changing this forces a new Backup Policy Disk to be created.
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The duration of default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	DefaultRetentionDuration *string `json:"defaultRetentionDuration,omitempty" tf:"default_retention_duration,omitempty"`

	// The ID of the Backup Policy Disk.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Disk to be created.
	RetentionRule []RetentionRuleObservation `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Disk should exist. Changing this forces a new Backup Policy Disk to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupPolicyDiskParameters struct {

	// Specifies a list of repeating time interval. It should follow ISO 8601 repeating time interval . Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals,omitempty" tf:"backup_repeating_time_intervals,omitempty"`

	// The duration of default retention rule. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	DefaultRetentionDuration *string `json:"defaultRetentionDuration,omitempty" tf:"default_retention_duration,omitempty"`

	// One or more retention_rule blocks as defined below. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	RetentionRule []RetentionRuleParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Disk should exist. Changing this forces a new Backup Policy Disk to be created.
	// +crossplane:generate:reference:type=BackupVault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

type CriteriaInitParameters struct {

	// Possible values are FirstOfDay and FirstOfWeek. Changing this forces a new Backup Policy Disk to be created.
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`
}

type CriteriaObservation struct {

	// Possible values are FirstOfDay and FirstOfWeek. Changing this forces a new Backup Policy Disk to be created.
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`
}

type CriteriaParameters struct {

	// Possible values are FirstOfDay and FirstOfWeek. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`
}

type RetentionRuleInitParameters struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Disk to be created.
	Criteria []CriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type RetentionRuleObservation struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Disk to be created.
	Criteria []CriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type RetentionRuleParameters struct {

	// A criteria block as defined below. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	Criteria []CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

// BackupPolicyDiskSpec defines the desired state of BackupPolicyDisk
type BackupPolicyDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyDiskParameters `json:"forProvider"`
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
	InitProvider BackupPolicyDiskInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyDiskStatus defines the observed state of BackupPolicyDisk.
type BackupPolicyDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyDisk is the Schema for the BackupPolicyDisks API. Manages a Backup Policy Disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backupRepeatingTimeIntervals) || has(self.initProvider.backupRepeatingTimeIntervals)",message="backupRepeatingTimeIntervals is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultRetentionDuration) || has(self.initProvider.defaultRetentionDuration)",message="defaultRetentionDuration is a required parameter"
	Spec   BackupPolicyDiskSpec   `json:"spec"`
	Status BackupPolicyDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyDiskList contains a list of BackupPolicyDisks
type BackupPolicyDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyDisk `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyDisk_Kind             = "BackupPolicyDisk"
	BackupPolicyDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyDisk_Kind}.String()
	BackupPolicyDisk_KindAPIVersion   = BackupPolicyDisk_Kind + "." + CRDGroupVersion.String()
	BackupPolicyDisk_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyDisk{}, &BackupPolicyDiskList{})
}
