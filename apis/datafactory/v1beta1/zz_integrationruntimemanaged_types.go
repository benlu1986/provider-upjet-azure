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

type IntegrationRuntimeManagedCatalogInfoInitParameters struct {

	// Administrator login name for the SQL Server.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Pricing tier for the database that will be created for the SSIS catalog. Valid values are: Basic, Standard, Premium and PremiumRS.
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	// The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
	ServerEndpoint *string `json:"serverEndpoint,omitempty" tf:"server_endpoint,omitempty"`
}

type IntegrationRuntimeManagedCatalogInfoObservation struct {

	// Administrator login name for the SQL Server.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Pricing tier for the database that will be created for the SSIS catalog. Valid values are: Basic, Standard, Premium and PremiumRS.
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	// The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
	ServerEndpoint *string `json:"serverEndpoint,omitempty" tf:"server_endpoint,omitempty"`
}

type IntegrationRuntimeManagedCatalogInfoParameters struct {

	// Administrator login name for the SQL Server.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Administrator login password for the SQL Server.
	// +kubebuilder:validation:Optional
	AdministratorPasswordSecretRef *v1.SecretKeySelector `json:"administratorPasswordSecretRef,omitempty" tf:"-"`

	// Pricing tier for the database that will be created for the SSIS catalog. Valid values are: Basic, Standard, Premium and PremiumRS.
	// +kubebuilder:validation:Optional
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	// The endpoint of an Azure SQL Server that will be used to host the SSIS catalog.
	// +kubebuilder:validation:Optional
	ServerEndpoint *string `json:"serverEndpoint" tf:"server_endpoint,omitempty"`
}

type IntegrationRuntimeManagedCustomSetupScriptInitParameters struct {

	// The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See https://docs.microsoft.com/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup for more information.
	BlobContainerURI *string `json:"blobContainerUri,omitempty" tf:"blob_container_uri,omitempty"`
}

type IntegrationRuntimeManagedCustomSetupScriptObservation struct {

	// The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See https://docs.microsoft.com/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup for more information.
	BlobContainerURI *string `json:"blobContainerUri,omitempty" tf:"blob_container_uri,omitempty"`
}

type IntegrationRuntimeManagedCustomSetupScriptParameters struct {

	// The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See https://docs.microsoft.com/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup for more information.
	// +kubebuilder:validation:Optional
	BlobContainerURI *string `json:"blobContainerUri" tf:"blob_container_uri,omitempty"`

	// A container SAS token that gives access to the files. See https://docs.microsoft.com/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup for more information.
	// +kubebuilder:validation:Required
	SASTokenSecretRef v1.SecretKeySelector `json:"sasTokenSecretRef" tf:"-"`
}

type IntegrationRuntimeManagedInitParameters struct {

	// A catalog_info block as defined below.
	CatalogInfo []IntegrationRuntimeManagedCatalogInfoInitParameters `json:"catalogInfo,omitempty" tf:"catalog_info,omitempty"`

	// A custom_setup_script block as defined below.
	CustomSetupScript []IntegrationRuntimeManagedCustomSetupScriptInitParameters `json:"customSetupScript,omitempty" tf:"custom_setup_script,omitempty"`

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Managed Integration Runtime edition. Valid values are Standard and Enterprise. Defaults to Standard.
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// The type of the license that is used. Valid values are LicenseIncluded and BasePrice. Defaults to LicenseIncluded.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the maximum parallel executions per node. Defaults to 1. Max is 16.
	MaxParallelExecutionsPerNode *float64 `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node,omitempty"`

	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: Standard_D2_v3, Standard_D4_v3, Standard_D8_v3, Standard_D16_v3, Standard_D32_v3, Standard_D64_v3, Standard_E2_v3, Standard_E4_v3, Standard_E8_v3, Standard_E16_v3, Standard_E32_v3, Standard_E64_v3, Standard_D1_v2, Standard_D2_v2, Standard_D3_v2, Standard_D4_v2, Standard_A4_v2 and Standard_A8_v2
	NodeSize *string `json:"nodeSize,omitempty" tf:"node_size,omitempty"`

	// Number of nodes for the Managed Integration Runtime. Max is 10. Defaults to 1.
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// A vnet_integration block as defined below.
	VnetIntegration []IntegrationRuntimeManagedVnetIntegrationInitParameters `json:"vnetIntegration,omitempty" tf:"vnet_integration,omitempty"`
}

type IntegrationRuntimeManagedObservation struct {

	// A catalog_info block as defined below.
	CatalogInfo []IntegrationRuntimeManagedCatalogInfoObservation `json:"catalogInfo,omitempty" tf:"catalog_info,omitempty"`

	// A custom_setup_script block as defined below.
	CustomSetupScript []IntegrationRuntimeManagedCustomSetupScriptObservation `json:"customSetupScript,omitempty" tf:"custom_setup_script,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Integration runtime description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Managed Integration Runtime edition. Valid values are Standard and Enterprise. Defaults to Standard.
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// The ID of the Data Factory Integration Managed Runtime.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of the license that is used. Valid values are LicenseIncluded and BasePrice. Defaults to LicenseIncluded.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the maximum parallel executions per node. Defaults to 1. Max is 16.
	MaxParallelExecutionsPerNode *float64 `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node,omitempty"`

	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: Standard_D2_v3, Standard_D4_v3, Standard_D8_v3, Standard_D16_v3, Standard_D32_v3, Standard_D64_v3, Standard_E2_v3, Standard_E4_v3, Standard_E8_v3, Standard_E16_v3, Standard_E32_v3, Standard_E64_v3, Standard_D1_v2, Standard_D2_v2, Standard_D3_v2, Standard_D4_v2, Standard_A4_v2 and Standard_A8_v2
	NodeSize *string `json:"nodeSize,omitempty" tf:"node_size,omitempty"`

	// Number of nodes for the Managed Integration Runtime. Max is 10. Defaults to 1.
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// A vnet_integration block as defined below.
	VnetIntegration []IntegrationRuntimeManagedVnetIntegrationObservation `json:"vnetIntegration,omitempty" tf:"vnet_integration,omitempty"`
}

type IntegrationRuntimeManagedParameters struct {

	// A catalog_info block as defined below.
	// +kubebuilder:validation:Optional
	CatalogInfo []IntegrationRuntimeManagedCatalogInfoParameters `json:"catalogInfo,omitempty" tf:"catalog_info,omitempty"`

	// A custom_setup_script block as defined below.
	// +kubebuilder:validation:Optional
	CustomSetupScript []IntegrationRuntimeManagedCustomSetupScriptParameters `json:"customSetupScript,omitempty" tf:"custom_setup_script,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// Integration runtime description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Managed Integration Runtime edition. Valid values are Standard and Enterprise. Defaults to Standard.
	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// The type of the license that is used. Valid values are LicenseIncluded and BasePrice. Defaults to LicenseIncluded.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the maximum parallel executions per node. Defaults to 1. Max is 16.
	// +kubebuilder:validation:Optional
	MaxParallelExecutionsPerNode *float64 `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node,omitempty"`

	// The size of the nodes on which the Managed Integration Runtime runs. Valid values are: Standard_D2_v3, Standard_D4_v3, Standard_D8_v3, Standard_D16_v3, Standard_D32_v3, Standard_D64_v3, Standard_E2_v3, Standard_E4_v3, Standard_E8_v3, Standard_E16_v3, Standard_E32_v3, Standard_E64_v3, Standard_D1_v2, Standard_D2_v2, Standard_D3_v2, Standard_D4_v2, Standard_A4_v2 and Standard_A8_v2
	// +kubebuilder:validation:Optional
	NodeSize *string `json:"nodeSize,omitempty" tf:"node_size,omitempty"`

	// Number of nodes for the Managed Integration Runtime. Max is 10. Defaults to 1.
	// +kubebuilder:validation:Optional
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`

	// A vnet_integration block as defined below.
	// +kubebuilder:validation:Optional
	VnetIntegration []IntegrationRuntimeManagedVnetIntegrationParameters `json:"vnetIntegration,omitempty" tf:"vnet_integration,omitempty"`
}

type IntegrationRuntimeManagedVnetIntegrationInitParameters struct {

	// ID of the virtual network to which the nodes of the Managed Integration Runtime will be added.
	VnetID *string `json:"vnetId,omitempty" tf:"vnet_id,omitempty"`
}

type IntegrationRuntimeManagedVnetIntegrationObservation struct {

	// Name of the subnet to which the nodes of the Managed Integration Runtime will be added.
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`

	// ID of the virtual network to which the nodes of the Managed Integration Runtime will be added.
	VnetID *string `json:"vnetId,omitempty" tf:"vnet_id,omitempty"`
}

type IntegrationRuntimeManagedVnetIntegrationParameters struct {

	// Name of the subnet to which the nodes of the Managed Integration Runtime will be added.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`

	// Reference to a Subnet in network to populate subnetName.
	// +kubebuilder:validation:Optional
	SubnetNameRef *v1.Reference `json:"subnetNameRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetName.
	// +kubebuilder:validation:Optional
	SubnetNameSelector *v1.Selector `json:"subnetNameSelector,omitempty" tf:"-"`

	// ID of the virtual network to which the nodes of the Managed Integration Runtime will be added.
	// +kubebuilder:validation:Optional
	VnetID *string `json:"vnetId" tf:"vnet_id,omitempty"`
}

// IntegrationRuntimeManagedSpec defines the desired state of IntegrationRuntimeManaged
type IntegrationRuntimeManagedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationRuntimeManagedParameters `json:"forProvider"`
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
	InitProvider IntegrationRuntimeManagedInitParameters `json:"initProvider,omitempty"`
}

// IntegrationRuntimeManagedStatus defines the observed state of IntegrationRuntimeManaged.
type IntegrationRuntimeManagedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationRuntimeManagedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeManaged is the Schema for the IntegrationRuntimeManageds API. Manages an Azure Data Factory Managed Integration Runtime.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IntegrationRuntimeManaged struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeSize) || has(self.initProvider.nodeSize)",message="nodeSize is a required parameter"
	Spec   IntegrationRuntimeManagedSpec   `json:"spec"`
	Status IntegrationRuntimeManagedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationRuntimeManagedList contains a list of IntegrationRuntimeManageds
type IntegrationRuntimeManagedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationRuntimeManaged `json:"items"`
}

// Repository type metadata.
var (
	IntegrationRuntimeManaged_Kind             = "IntegrationRuntimeManaged"
	IntegrationRuntimeManaged_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationRuntimeManaged_Kind}.String()
	IntegrationRuntimeManaged_KindAPIVersion   = IntegrationRuntimeManaged_Kind + "." + CRDGroupVersion.String()
	IntegrationRuntimeManaged_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationRuntimeManaged_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationRuntimeManaged{}, &IntegrationRuntimeManagedList{})
}
