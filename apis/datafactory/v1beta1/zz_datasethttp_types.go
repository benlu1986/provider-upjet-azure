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

type DataSetHTTPInitParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`

	// The body for the HTTP request.
	RequestBody *string `json:"requestBody,omitempty" tf:"request_body,omitempty"`

	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetHTTPSchemaColumnInitParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetHTTPObservation struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`

	// The body for the HTTP request.
	RequestBody *string `json:"requestBody,omitempty" tf:"request_body,omitempty"`

	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetHTTPSchemaColumnObservation `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetHTTPParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

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

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The relative URL based on the URL in the HTTP Linked Service.
	// +kubebuilder:validation:Optional
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`

	// The body for the HTTP request.
	// +kubebuilder:validation:Optional
	RequestBody *string `json:"requestBody,omitempty" tf:"request_body,omitempty"`

	// The HTTP method for the HTTP request. (e.g. GET, POST)
	// +kubebuilder:validation:Optional
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// A schema_column block as defined below.
	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetHTTPSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetHTTPSchemaColumnInitParameters struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetHTTPSchemaColumnObservation struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetHTTPSchemaColumnParameters struct {

	// The description of the column.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetHTTPSpec defines the desired state of DataSetHTTP
type DataSetHTTPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetHTTPParameters `json:"forProvider"`
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
	InitProvider DataSetHTTPInitParameters `json:"initProvider,omitempty"`
}

// DataSetHTTPStatus defines the observed state of DataSetHTTP.
type DataSetHTTPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetHTTPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetHTTP is the Schema for the DataSetHTTPs API. Manages an Azure Delimited Text Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetHTTP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetHTTPSpec   `json:"spec"`
	Status            DataSetHTTPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetHTTPList contains a list of DataSetHTTPs
type DataSetHTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetHTTP `json:"items"`
}

// Repository type metadata.
var (
	DataSetHTTP_Kind             = "DataSetHTTP"
	DataSetHTTP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetHTTP_Kind}.String()
	DataSetHTTP_KindAPIVersion   = DataSetHTTP_Kind + "." + CRDGroupVersion.String()
	DataSetHTTP_GroupVersionKind = CRDGroupVersion.WithKind(DataSetHTTP_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetHTTP{}, &DataSetHTTPList{})
}
