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

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Managed Instance. Required when type is set to UserAssigned.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Managed Instance. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Managed Instance. Required when type is set to UserAssigned.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Identity of this SQL Managed Instance.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Identity of this SQL Managed Instance.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Managed Instance. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this SQL Managed Instance. Required when type is set to UserAssigned.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this SQL Managed Instance. Possible values are SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MSSQLManagedInstanceInitParameters struct {

	// The administrator login name for the new SQL Managed Instance. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Specifies how the SQL Managed Instance will be collated. Default value is SQL_Latin1_General_CP1_CI_AS. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// What type of license the Managed Instance will use. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the SQL Managed Instance. Valid values include SQL_Default or an Azure Location in the format SQL_{Location}_MI_{Size}(for example SQL_EastUS_MI_1). Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The Minimum TLS Version. Default value is 1.2 Valid values include 1.0, 1.1, 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Specifies how the SQL Managed Instance will be accessed. Default value is Default. Valid values include Default, Proxy, and Redirect.
	ProxyOverride *string `json:"proxyOverride,omitempty" tf:"proxy_override,omitempty"`

	// Is the public data endpoint enabled? Default value is false.
	PublicDataEndpointEnabled *bool `json:"publicDataEndpointEnabled,omitempty" tf:"public_data_endpoint_enabled,omitempty"`

	// Specifies the SKU Name for the SQL Managed Instance. Valid values include GP_Gen4, GP_Gen5, GP_Gen8IM, GP_Gen8IH, BC_Gen4, BC_Gen5, BC_Gen8IM or BC_Gen8IH.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Changing this forces a new resource to be created. Possible values are GRS, LRS and ZRS. The default value is GRS.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// Maximum storage space for the SQL Managed instance. This should be a multiple of 32 (GB).
	StorageSizeInGb *float64 `json:"storageSizeInGb,omitempty" tf:"storage_size_in_gb,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is UTC. Changing this forces a new resource to be created.
	TimezoneID *string `json:"timezoneId,omitempty" tf:"timezone_id,omitempty"`

	// Number of cores that should be assigned to the SQL Managed Instance. Values can be 8, 16, or 24 for Gen4 SKUs, or 4, 8, 16, 24, 32, 40, 64, or 80 for Gen5 SKUs.
	Vcores *float64 `json:"vcores,omitempty" tf:"vcores,omitempty"`
}

type MSSQLManagedInstanceObservation struct {

	// The administrator login name for the new SQL Managed Instance. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Specifies how the SQL Managed Instance will be collated. Default value is SQL_Latin1_General_CP1_CI_AS. Changing this forces a new resource to be created.
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The ID of the SQL Managed Instance which will share the DNS zone. This is a prerequisite for creating an azurerm_sql_managed_instance_failover_group. Setting this after creation forces a new resource to be created.
	DNSZonePartnerID *string `json:"dnsZonePartnerId,omitempty" tf:"dns_zone_partner_id,omitempty"`

	// The fully qualified domain name of the Azure Managed SQL Instance
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The SQL Managed Instance ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// What type of license the Managed Instance will use. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the SQL Managed Instance. Valid values include SQL_Default or an Azure Location in the format SQL_{Location}_MI_{Size}(for example SQL_EastUS_MI_1). Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The Minimum TLS Version. Default value is 1.2 Valid values include 1.0, 1.1, 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Specifies how the SQL Managed Instance will be accessed. Default value is Default. Valid values include Default, Proxy, and Redirect.
	ProxyOverride *string `json:"proxyOverride,omitempty" tf:"proxy_override,omitempty"`

	// Is the public data endpoint enabled? Default value is false.
	PublicDataEndpointEnabled *bool `json:"publicDataEndpointEnabled,omitempty" tf:"public_data_endpoint_enabled,omitempty"`

	// The name of the resource group in which to create the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the SKU Name for the SQL Managed Instance. Valid values include GP_Gen4, GP_Gen5, GP_Gen8IM, GP_Gen8IH, BC_Gen4, BC_Gen5, BC_Gen8IM or BC_Gen8IH.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Changing this forces a new resource to be created. Possible values are GRS, LRS and ZRS. The default value is GRS.
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// Maximum storage space for the SQL Managed instance. This should be a multiple of 32 (GB).
	StorageSizeInGb *float64 `json:"storageSizeInGb,omitempty" tf:"storage_size_in_gb,omitempty"`

	// The subnet resource id that the SQL Managed Instance will be associated with. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is UTC. Changing this forces a new resource to be created.
	TimezoneID *string `json:"timezoneId,omitempty" tf:"timezone_id,omitempty"`

	// Number of cores that should be assigned to the SQL Managed Instance. Values can be 8, 16, or 24 for Gen4 SKUs, or 4, 8, 16, 24, 32, 40, 64, or 80 for Gen5 SKUs.
	Vcores *float64 `json:"vcores,omitempty" tf:"vcores,omitempty"`
}

type MSSQLManagedInstanceParameters struct {

	// The administrator login name for the new SQL Managed Instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The password associated with the administrator_login user. Needs to comply with Azure's Password Policy
	// +kubebuilder:validation:Optional
	AdministratorLoginPasswordSecretRef v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef" tf:"-"`

	// Specifies how the SQL Managed Instance will be collated. Default value is SQL_Latin1_General_CP1_CI_AS. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// The ID of the SQL Managed Instance which will share the DNS zone. This is a prerequisite for creating an azurerm_sql_managed_instance_failover_group. Setting this after creation forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DNSZonePartnerID *string `json:"dnsZonePartnerId,omitempty" tf:"dns_zone_partner_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate dnsZonePartnerId.
	// +kubebuilder:validation:Optional
	DNSZonePartnerIDRef *v1.Reference `json:"dnsZonePartnerIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate dnsZonePartnerId.
	// +kubebuilder:validation:Optional
	DNSZonePartnerIDSelector *v1.Selector `json:"dnsZonePartnerIdSelector,omitempty" tf:"-"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// What type of license the Managed Instance will use. Possible values are LicenseIncluded and BasePrice.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the SQL Managed Instance. Valid values include SQL_Default or an Azure Location in the format SQL_{Location}_MI_{Size}(for example SQL_EastUS_MI_1). Defaults to SQL_Default.
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The Minimum TLS Version. Default value is 1.2 Valid values include 1.0, 1.1, 1.2.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Specifies how the SQL Managed Instance will be accessed. Default value is Default. Valid values include Default, Proxy, and Redirect.
	// +kubebuilder:validation:Optional
	ProxyOverride *string `json:"proxyOverride,omitempty" tf:"proxy_override,omitempty"`

	// Is the public data endpoint enabled? Default value is false.
	// +kubebuilder:validation:Optional
	PublicDataEndpointEnabled *bool `json:"publicDataEndpointEnabled,omitempty" tf:"public_data_endpoint_enabled,omitempty"`

	// The name of the resource group in which to create the SQL Managed Instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU Name for the SQL Managed Instance. Valid values include GP_Gen4, GP_Gen5, GP_Gen8IM, GP_Gen8IH, BC_Gen4, BC_Gen5, BC_Gen8IM or BC_Gen8IH.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the storage account type used to store backups for this database. Changing this forces a new resource to be created. Possible values are GRS, LRS and ZRS. The default value is GRS.
	// +kubebuilder:validation:Optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`

	// Maximum storage space for the SQL Managed instance. This should be a multiple of 32 (GB).
	// +kubebuilder:validation:Optional
	StorageSizeInGb *float64 `json:"storageSizeInGb,omitempty" tf:"storage_size_in_gb,omitempty"`

	// The subnet resource id that the SQL Managed Instance will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The TimeZone ID that the SQL Managed Instance will be operating in. Default value is UTC. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TimezoneID *string `json:"timezoneId,omitempty" tf:"timezone_id,omitempty"`

	// Number of cores that should be assigned to the SQL Managed Instance. Values can be 8, 16, or 24 for Gen4 SKUs, or 4, 8, 16, 24, 32, 40, 64, or 80 for Gen5 SKUs.
	// +kubebuilder:validation:Optional
	Vcores *float64 `json:"vcores,omitempty" tf:"vcores,omitempty"`
}

// MSSQLManagedInstanceSpec defines the desired state of MSSQLManagedInstance
type MSSQLManagedInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedInstanceParameters `json:"forProvider"`
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
	InitProvider MSSQLManagedInstanceInitParameters `json:"initProvider,omitempty"`
}

// MSSQLManagedInstanceStatus defines the observed state of MSSQLManagedInstance.
type MSSQLManagedInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstance is the Schema for the MSSQLManagedInstances API. Manages a Microsoft SQL Azure Managed Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.administratorLogin) || has(self.initProvider.administratorLogin)",message="administratorLogin is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.administratorLoginPasswordSecretRef)",message="administratorLoginPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.licenseType) || has(self.initProvider.licenseType)",message="licenseType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSizeInGb) || has(self.initProvider.storageSizeInGb)",message="storageSizeInGb is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vcores) || has(self.initProvider.vcores)",message="vcores is a required parameter"
	Spec   MSSQLManagedInstanceSpec   `json:"spec"`
	Status MSSQLManagedInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstanceList contains a list of MSSQLManagedInstances
type MSSQLManagedInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedInstance `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedInstance_Kind             = "MSSQLManagedInstance"
	MSSQLManagedInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedInstance_Kind}.String()
	MSSQLManagedInstance_KindAPIVersion   = MSSQLManagedInstance_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedInstance_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedInstance{}, &MSSQLManagedInstanceList{})
}
