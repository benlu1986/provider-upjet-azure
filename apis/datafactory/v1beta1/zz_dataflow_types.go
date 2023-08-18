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

type DataFlowInitParameters struct {

	// List of tags that can be used for describing the Data Factory Data Flow.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Data Flow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Data Flow is in. If not specified, the Data Flow will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The script for the Data Factory Data Flow.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`

	// The script lines for the Data Factory Data Flow.
	ScriptLines []*string `json:"scriptLines,omitempty" tf:"script_lines,omitempty"`

	// One or more sink blocks as defined below.
	Sink []SinkInitParameters `json:"sink,omitempty" tf:"sink,omitempty"`

	// One or more source blocks as defined below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`

	// One or more transformation blocks as defined below.
	Transformation []TransformationInitParameters `json:"transformation,omitempty" tf:"transformation,omitempty"`
}

type DataFlowObservation struct {

	// List of tags that can be used for describing the Data Factory Data Flow.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The ID of Data Factory in which to associate the Data Flow with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Data Flow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Data Flow is in. If not specified, the Data Flow will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The ID of the Data Factory Data Flow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The script for the Data Factory Data Flow.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`

	// The script lines for the Data Factory Data Flow.
	ScriptLines []*string `json:"scriptLines,omitempty" tf:"script_lines,omitempty"`

	// One or more sink blocks as defined below.
	Sink []SinkObservation `json:"sink,omitempty" tf:"sink,omitempty"`

	// One or more source blocks as defined below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// One or more transformation blocks as defined below.
	Transformation []TransformationObservation `json:"transformation,omitempty" tf:"transformation,omitempty"`
}

type DataFlowParameters struct {

	// List of tags that can be used for describing the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The ID of Data Factory in which to associate the Data Flow with. Changing this forces a new resource.
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

	// The description for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Data Flow is in. If not specified, the Data Flow will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The script for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Script *string `json:"script,omitempty" tf:"script,omitempty"`

	// The script lines for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	ScriptLines []*string `json:"scriptLines,omitempty" tf:"script_lines,omitempty"`

	// One or more sink blocks as defined below.
	// +kubebuilder:validation:Optional
	Sink []SinkParameters `json:"sink,omitempty" tf:"sink,omitempty"`

	// One or more source blocks as defined below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// One or more transformation blocks as defined below.
	// +kubebuilder:validation:Optional
	Transformation []TransformationParameters `json:"transformation,omitempty" tf:"transformation,omitempty"`
}

type DataSetInitParameters struct {

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type DataSetObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type DataSetParameters struct {

	// The name for the Data Flow transformation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.DataSetJSON
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FlowletInitParameters struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FlowletObservation struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type RejectedLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type RejectedLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type RejectedLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SchemaLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SchemaLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SchemaLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkInitParameters struct {

	// A dataset block as defined below.
	DataSet []DataSetInitParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []FlowletInitParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []SinkLinkedServiceInitParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	RejectedLinkedService []RejectedLinkedServiceInitParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	SchemaLinkedService []SchemaLinkedServiceInitParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SinkLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkObservation struct {

	// A dataset block as defined below.
	DataSet []DataSetObservation `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []FlowletObservation `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []SinkLinkedServiceObservation `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	RejectedLinkedService []RejectedLinkedServiceObservation `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	SchemaLinkedService []SchemaLinkedServiceObservation `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SinkParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []DataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []FlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []SinkLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	RejectedLinkedService []RejectedLinkedServiceParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	SchemaLinkedService []SchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceDataSetInitParameters struct {

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceDataSetObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceDataSetParameters struct {

	// The name for the Data Flow transformation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.DataSetJSON
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceFlowletInitParameters struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceFlowletObservation struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceFlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceInitParameters struct {

	// A dataset block as defined below.
	DataSet []SourceDataSetInitParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []SourceFlowletInitParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []SourceLinkedServiceInitParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	RejectedLinkedService []SourceRejectedLinkedServiceInitParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	SchemaLinkedService []SourceSchemaLinkedServiceInitParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceObservation struct {

	// A dataset block as defined below.
	DataSet []SourceDataSetObservation `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []SourceFlowletObservation `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []SourceLinkedServiceObservation `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	RejectedLinkedService []SourceRejectedLinkedServiceObservation `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	SchemaLinkedService []SourceSchemaLinkedServiceObservation `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []SourceDataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []SourceFlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []SourceLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	RejectedLinkedService []SourceRejectedLinkedServiceParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	SchemaLinkedService []SourceSchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceRejectedLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceRejectedLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceRejectedLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceSchemaLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceSchemaLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceSchemaLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationDataSetInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationDataSetObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationDataSetParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationFlowletInitParameters struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationFlowletObservation struct {

	// Specifies the reference data flow parameters from dataset.
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationFlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationInitParameters struct {

	// A dataset block as defined below.
	DataSet []TransformationDataSetInitParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow transformation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []TransformationFlowletInitParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []TransformationLinkedServiceInitParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TransformationLinkedServiceInitParameters struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationLinkedServiceObservation struct {

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationObservation struct {

	// A dataset block as defined below.
	DataSet []TransformationDataSetObservation `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow transformation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	Flowlet []TransformationFlowletObservation `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	LinkedService []TransformationLinkedServiceObservation `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow transformation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TransformationParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []TransformationDataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []TransformationFlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []TransformationLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// DataFlowSpec defines the desired state of DataFlow
type DataFlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFlowParameters `json:"forProvider"`
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
	InitProvider DataFlowInitParameters `json:"initProvider,omitempty"`
}

// DataFlowStatus defines the observed state of DataFlow.
type DataFlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFlow is the Schema for the DataFlows API. Manages a Data Flow inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sink) || has(self.initProvider.sink)",message="sink is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || has(self.initProvider.source)",message="source is a required parameter"
	Spec   DataFlowSpec   `json:"spec"`
	Status DataFlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFlowList contains a list of DataFlows
type DataFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFlow `json:"items"`
}

// Repository type metadata.
var (
	DataFlow_Kind             = "DataFlow"
	DataFlow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataFlow_Kind}.String()
	DataFlow_KindAPIVersion   = DataFlow_Kind + "." + CRDGroupVersion.String()
	DataFlow_GroupVersionKind = CRDGroupVersion.WithKind(DataFlow_Kind)
)

func init() {
	SchemeBuilder.Register(&DataFlow{}, &DataFlowList{})
}
