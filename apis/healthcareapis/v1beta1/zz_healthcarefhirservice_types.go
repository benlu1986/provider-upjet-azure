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

type CorsInitParameters struct {

	// A set of headers to be allowed via CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// If credentials are allowed via CORS.
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// The max age to be allowed via CORS.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsObservation struct {

	// A set of headers to be allowed via CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// If credentials are allowed via CORS.
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// The max age to be allowed via CORS.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsParameters struct {

	// A set of headers to be allowed via CORS.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// If credentials are allowed via CORS.
	// +kubebuilder:validation:Optional
	CredentialsAllowed *bool `json:"credentialsAllowed,omitempty" tf:"credentials_allowed,omitempty"`

	// The max age to be allowed via CORS.
	// +kubebuilder:validation:Optional
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type HealthcareFHIRServiceAuthenticationInitParameters struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://<name>.fhir.azurehealthcareapis.com.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// Whether smart proxy is enabled.
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type HealthcareFHIRServiceAuthenticationObservation struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://<name>.fhir.azurehealthcareapis.com.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// Whether smart proxy is enabled.
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type HealthcareFHIRServiceAuthenticationParameters struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://<name>.fhir.azurehealthcareapis.com.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	// +kubebuilder:validation:Optional
	Authority *string `json:"authority" tf:"authority,omitempty"`

	// Whether smart proxy is enabled.
	// +kubebuilder:validation:Optional
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type HealthcareFHIRServiceIdentityInitParameters struct {

	// A list of one or more Resource IDs for User Assigned Managed identities to assign. Required when type is set to UserAssigned.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of managed identity to assign. Possible values are UserAssigned and SystemAssigned
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareFHIRServiceIdentityObservation struct {

	// A list of one or more Resource IDs for User Assigned Managed identities to assign. Required when type is set to UserAssigned.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The ID of the Healthcare FHIR Service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Healthcare FHIR Service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of managed identity to assign. Possible values are UserAssigned and SystemAssigned
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareFHIRServiceIdentityParameters struct {

	// A list of one or more Resource IDs for User Assigned Managed identities to assign. Required when type is set to UserAssigned.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of managed identity to assign. Possible values are UserAssigned and SystemAssigned
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type HealthcareFHIRServiceInitParameters struct {

	// A list of the access policies of the service instance.
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication block as defined below.
	Authentication []HealthcareFHIRServiceAuthenticationInitParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A list of azure container registry settings used for convert data operation of the service instance.
	ContainerRegistryLoginServerURL []*string `json:"containerRegistryLoginServerUrl,omitempty" tf:"container_registry_login_server_url,omitempty"`

	// A cors block as defined below.
	Cors []CorsInitParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// An identity block as defined below.
	Identity []HealthcareFHIRServiceIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the kind of the Healthcare FHIR Service. Possible values are: fhir-Stu3 and fhir-R4. Defaults to fhir-R4. Changing this forces a new Healthcare FHIR Service to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Azure Region where the Healthcare FHIR Service should be created. Changing this forces a new Healthcare FHIR Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of objects describing OCI artifacts for export as defined below.
	OciArtifact []OciArtifactInitParameters `json:"ociArtifact,omitempty" tf:"oci_artifact,omitempty"`

	// A mapping of tags to assign to the Healthcare FHIR Service.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareFHIRServiceObservation struct {

	// A list of the access policies of the service instance.
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication block as defined below.
	Authentication []HealthcareFHIRServiceAuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A list of azure container registry settings used for convert data operation of the service instance.
	ContainerRegistryLoginServerURL []*string `json:"containerRegistryLoginServerUrl,omitempty" tf:"container_registry_login_server_url,omitempty"`

	// A cors block as defined below.
	Cors []CorsObservation `json:"cors,omitempty" tf:"cors,omitempty"`

	// The ID of the Healthcare FHIR Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []HealthcareFHIRServiceIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the kind of the Healthcare FHIR Service. Possible values are: fhir-Stu3 and fhir-R4. Defaults to fhir-R4. Changing this forces a new Healthcare FHIR Service to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Azure Region where the Healthcare FHIR Service should be created. Changing this forces a new Healthcare FHIR Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of objects describing OCI artifacts for export as defined below.
	OciArtifact []OciArtifactObservation `json:"ociArtifact,omitempty" tf:"oci_artifact,omitempty"`

	// Whether public networks access is enabled.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the name of the Resource Group in which to create the Healthcare FHIR Service. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Healthcare FHIR Service.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of the Healthcare Workspace where the Healthcare FHIR Service should exist. Changing this forces a new Healthcare FHIR Service to be created.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type HealthcareFHIRServiceParameters struct {

	// A list of the access policies of the service instance.
	// +kubebuilder:validation:Optional
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication block as defined below.
	// +kubebuilder:validation:Optional
	Authentication []HealthcareFHIRServiceAuthenticationParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	// +kubebuilder:validation:Optional
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A list of azure container registry settings used for convert data operation of the service instance.
	// +kubebuilder:validation:Optional
	ContainerRegistryLoginServerURL []*string `json:"containerRegistryLoginServerUrl,omitempty" tf:"container_registry_login_server_url,omitempty"`

	// A cors block as defined below.
	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []HealthcareFHIRServiceIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the kind of the Healthcare FHIR Service. Possible values are: fhir-Stu3 and fhir-R4. Defaults to fhir-R4. Changing this forces a new Healthcare FHIR Service to be created.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the Azure Region where the Healthcare FHIR Service should be created. Changing this forces a new Healthcare FHIR Service to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of objects describing OCI artifacts for export as defined below.
	// +kubebuilder:validation:Optional
	OciArtifact []OciArtifactParameters `json:"ociArtifact,omitempty" tf:"oci_artifact,omitempty"`

	// Specifies the name of the Resource Group in which to create the Healthcare FHIR Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Healthcare FHIR Service.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of the Healthcare Workspace where the Healthcare FHIR Service should exist. Changing this forces a new Healthcare FHIR Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/healthcareapis/v1beta1.HealthcareWorkspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type OciArtifactInitParameters struct {

	// A digest of an image within Azure container registry used for export operations of the service instance to narrow the artifacts down.
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// An image within Azure container registry used for export operations of the service instance.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// An Azure container registry used for export operations of the service instance.
	LoginServer *string `json:"loginServer,omitempty" tf:"login_server,omitempty"`
}

type OciArtifactObservation struct {

	// A digest of an image within Azure container registry used for export operations of the service instance to narrow the artifacts down.
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// An image within Azure container registry used for export operations of the service instance.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// An Azure container registry used for export operations of the service instance.
	LoginServer *string `json:"loginServer,omitempty" tf:"login_server,omitempty"`
}

type OciArtifactParameters struct {

	// A digest of an image within Azure container registry used for export operations of the service instance to narrow the artifacts down.
	// +kubebuilder:validation:Optional
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// An image within Azure container registry used for export operations of the service instance.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// An Azure container registry used for export operations of the service instance.
	// +kubebuilder:validation:Optional
	LoginServer *string `json:"loginServer" tf:"login_server,omitempty"`
}

// HealthcareFHIRServiceSpec defines the desired state of HealthcareFHIRService
type HealthcareFHIRServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareFHIRServiceParameters `json:"forProvider"`
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
	InitProvider HealthcareFHIRServiceInitParameters `json:"initProvider,omitempty"`
}

// HealthcareFHIRServiceStatus defines the observed state of HealthcareFHIRService.
type HealthcareFHIRServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareFHIRServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareFHIRService is the Schema for the HealthcareFHIRServices API. Manages a Healthcare FHIR (Fast Healthcare Interoperability Resources) Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareFHIRService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authentication) || has(self.initProvider.authentication)",message="authentication is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   HealthcareFHIRServiceSpec   `json:"spec"`
	Status HealthcareFHIRServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareFHIRServiceList contains a list of HealthcareFHIRServices
type HealthcareFHIRServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareFHIRService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareFHIRService_Kind             = "HealthcareFHIRService"
	HealthcareFHIRService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareFHIRService_Kind}.String()
	HealthcareFHIRService_KindAPIVersion   = HealthcareFHIRService_Kind + "." + CRDGroupVersion.String()
	HealthcareFHIRService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareFHIRService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareFHIRService{}, &HealthcareFHIRServiceList{})
}
