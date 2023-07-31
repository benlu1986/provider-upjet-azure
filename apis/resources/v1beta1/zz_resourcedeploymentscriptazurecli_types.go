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

type ContainerInitParameters struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type ContainerObservation struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type ContainerParameters struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	// +kubebuilder:validation:Optional
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type EnvironmentVariableInitParameters struct {

	// Specifies the name of the environment variable.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnvironmentVariableObservation struct {

	// Specifies the name of the environment variable.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnvironmentVariableParameters struct {

	// Specifies the name of the environment variable.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the secure environment variable.
	// +kubebuilder:validation:Optional
	SecureValueSecretRef *v1.SecretKeySelector `json:"secureValueSecretRef,omitempty" tf:"-"`

	// Specifies the value of the environment variable.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IdentityInitParameters struct {

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies the list of user-assigned managed identity IDs associated with the resource. Changing this forces a new resource to be created.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies the list of user-assigned managed identity IDs associated with the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// References to UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsRefs []v1.Reference `json:"identityIdsRefs,omitempty" tf:"-"`

	// Selector for a list of UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsSelector *v1.Selector `json:"identityIdsSelector,omitempty" tf:"-"`

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceDeploymentScriptAzureCliInitParameters struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Container []ContainerInitParameters `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	EnvironmentVariable []EnvironmentVariableInitParameters `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	StorageAccount []StorageAccountInitParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure CLI module version to be used. The supported versions are 2.0.77, 2.0.78, 2.0.79, 2.0.80, 2.0.81, 2.1.0, 2.10.0, 2.10.1, 2.11.0, 2.11.1, 2.12.0, 2.12.1, 2.13.0, 2.14.0, 2.14.1, 2.14.2, 2.15.0, 2.15.1, 2.16.0, 2.17.0, 2.17.1, 2.18.0, 2.19.0, 2.19.1, 2.2.0, 2.20.0, 2.21.0, 2.22.0, 2.22.1, 2.23.0, 2.24.0, 2.24.1, 2.24.2, 2.25.0, 2.26.0, 2.26.1, 2.27.0, 2.27.1, 2.27.2, 2.28.0, 2.29.0, 2.29.1, 2.29.2, 2.3.0, 2.3.1, 2.30.0, 2.31.0, 2.32.0, 2.33.0, 2.33.1, 2.34.0, 2.34.1, 2.35.0, 2.36.0, 2.37.0, 2.38.0, 2.39.0, 2.4.0, 2.40.0, 2.41.0, 2.5.0, 2.5.1, 2.6.0, 2.7.0, 2.8.0, 2.9.0, 2.9.1. Changing this forces a new Resource Deployment Script to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ResourceDeploymentScriptAzureCliObservation struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Container []ContainerObservation `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	EnvironmentVariable []EnvironmentVariableObservation `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// The ID of the Resource Deployment Script.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of script outputs.
	Outputs *string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Specifies the name of the Resource Group where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	StorageAccount []StorageAccountObservation `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure CLI module version to be used. The supported versions are 2.0.77, 2.0.78, 2.0.79, 2.0.80, 2.0.81, 2.1.0, 2.10.0, 2.10.1, 2.11.0, 2.11.1, 2.12.0, 2.12.1, 2.13.0, 2.14.0, 2.14.1, 2.14.2, 2.15.0, 2.15.1, 2.16.0, 2.17.0, 2.17.1, 2.18.0, 2.19.0, 2.19.1, 2.2.0, 2.20.0, 2.21.0, 2.22.0, 2.22.1, 2.23.0, 2.24.0, 2.24.1, 2.24.2, 2.25.0, 2.26.0, 2.26.1, 2.27.0, 2.27.1, 2.27.2, 2.28.0, 2.29.0, 2.29.1, 2.29.2, 2.3.0, 2.3.1, 2.30.0, 2.31.0, 2.32.0, 2.33.0, 2.33.1, 2.34.0, 2.34.1, 2.35.0, 2.36.0, 2.37.0, 2.38.0, 2.39.0, 2.4.0, 2.40.0, 2.41.0, 2.5.0, 2.5.1, 2.6.0, 2.7.0, 2.8.0, 2.9.0, 2.9.1. Changing this forces a new Resource Deployment Script to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ResourceDeploymentScriptAzureCliParameters struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Container []ContainerParameters `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	EnvironmentVariable []EnvironmentVariableParameters `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Specifies the name of the Resource Group where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	StorageAccount []StorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure CLI module version to be used. The supported versions are 2.0.77, 2.0.78, 2.0.79, 2.0.80, 2.0.81, 2.1.0, 2.10.0, 2.10.1, 2.11.0, 2.11.1, 2.12.0, 2.12.1, 2.13.0, 2.14.0, 2.14.1, 2.14.2, 2.15.0, 2.15.1, 2.16.0, 2.17.0, 2.17.1, 2.18.0, 2.19.0, 2.19.1, 2.2.0, 2.20.0, 2.21.0, 2.22.0, 2.22.1, 2.23.0, 2.24.0, 2.24.1, 2.24.2, 2.25.0, 2.26.0, 2.26.1, 2.27.0, 2.27.1, 2.27.2, 2.28.0, 2.29.0, 2.29.1, 2.29.2, 2.3.0, 2.3.1, 2.30.0, 2.31.0, 2.32.0, 2.33.0, 2.33.1, 2.34.0, 2.34.1, 2.35.0, 2.36.0, 2.37.0, 2.38.0, 2.39.0, 2.4.0, 2.40.0, 2.41.0, 2.5.0, 2.5.1, 2.6.0, 2.7.0, 2.8.0, 2.9.0, 2.9.1. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type StorageAccountInitParameters struct {

	// Specifies the storage account name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StorageAccountObservation struct {

	// Specifies the storage account name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StorageAccountParameters struct {

	// Specifies the storage account access key.
	// +kubebuilder:validation:Required
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// Specifies the storage account name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ResourceDeploymentScriptAzureCliSpec defines the desired state of ResourceDeploymentScriptAzureCli
type ResourceDeploymentScriptAzureCliSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceDeploymentScriptAzureCliParameters `json:"forProvider"`
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
	InitProvider ResourceDeploymentScriptAzureCliInitParameters `json:"initProvider,omitempty"`
}

// ResourceDeploymentScriptAzureCliStatus defines the observed state of ResourceDeploymentScriptAzureCli.
type ResourceDeploymentScriptAzureCliStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceDeploymentScriptAzureCliObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDeploymentScriptAzureCli is the Schema for the ResourceDeploymentScriptAzureClis API. Manages a Resource Deployment Script of Azure Cli.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceDeploymentScriptAzureCli struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionInterval) || has(self.initProvider.retentionInterval)",message="retentionInterval is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || has(self.initProvider.version)",message="version is a required parameter"
	Spec   ResourceDeploymentScriptAzureCliSpec   `json:"spec"`
	Status ResourceDeploymentScriptAzureCliStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDeploymentScriptAzureCliList contains a list of ResourceDeploymentScriptAzureClis
type ResourceDeploymentScriptAzureCliList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceDeploymentScriptAzureCli `json:"items"`
}

// Repository type metadata.
var (
	ResourceDeploymentScriptAzureCli_Kind             = "ResourceDeploymentScriptAzureCli"
	ResourceDeploymentScriptAzureCli_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceDeploymentScriptAzureCli_Kind}.String()
	ResourceDeploymentScriptAzureCli_KindAPIVersion   = ResourceDeploymentScriptAzureCli_Kind + "." + CRDGroupVersion.String()
	ResourceDeploymentScriptAzureCli_GroupVersionKind = CRDGroupVersion.WithKind(ResourceDeploymentScriptAzureCli_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceDeploymentScriptAzureCli{}, &ResourceDeploymentScriptAzureCliList{})
}
