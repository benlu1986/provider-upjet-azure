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

type AppServicePlanInitParameters struct {

	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// Whether to create a xenon App Service Plan.
	IsXenon *bool `json:"isXenon,omitempty" tf:"is_xenon,omitempty"`

	// The kind of the App Service Plan to create. Possible values are Windows (also available as App), Linux, elastic (for Premium Consumption), xenon and FunctionApp (for a Consumption Plan). Defaults to Windows. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// Can Apps assigned to this App Service Plan be scaled independently? If set to false apps assigned to this plan will scale to all instances of the plan.
	PerSiteScaling *bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling,omitempty"`

	// Is this App Service Plan Reserved.
	Reserved *bool `json:"reserved,omitempty" tf:"reserved,omitempty"`

	// A sku block as documented below.
	Sku *SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if the App Service Plan should be Zone Redundant. Changing this forces a new resource to be created.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type AppServicePlanObservation struct {

	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// The ID of the App Service Plan component.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to create a xenon App Service Plan.
	IsXenon *bool `json:"isXenon,omitempty" tf:"is_xenon,omitempty"`

	// The kind of the App Service Plan to create. Possible values are Windows (also available as App), Linux, elastic (for Premium Consumption), xenon and FunctionApp (for a Consumption Plan). Defaults to Windows. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// The maximum number of workers supported with the App Service Plan's sku.
	MaximumNumberOfWorkers *float64 `json:"maximumNumberOfWorkers,omitempty" tf:"maximum_number_of_workers,omitempty"`

	// Can Apps assigned to this App Service Plan be scaled independently? If set to false apps assigned to this plan will scale to all instances of the plan.
	PerSiteScaling *bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling,omitempty"`

	// Is this App Service Plan Reserved.
	Reserved *bool `json:"reserved,omitempty" tf:"reserved,omitempty"`

	// The name of the resource group in which to create the App Service Plan component. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A sku block as documented below.
	Sku *SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if the App Service Plan should be Zone Redundant. Changing this forces a new resource to be created.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type AppServicePlanParameters struct {

	// The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// Whether to create a xenon App Service Plan.
	// +kubebuilder:validation:Optional
	IsXenon *bool `json:"isXenon,omitempty" tf:"is_xenon,omitempty"`

	// The kind of the App Service Plan to create. Possible values are Windows (also available as App), Linux, elastic (for Premium Consumption), xenon and FunctionApp (for a Consumption Plan). Defaults to Windows. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
	// +kubebuilder:validation:Optional
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// Can Apps assigned to this App Service Plan be scaled independently? If set to false apps assigned to this plan will scale to all instances of the plan.
	// +kubebuilder:validation:Optional
	PerSiteScaling *bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling,omitempty"`

	// Is this App Service Plan Reserved.
	// +kubebuilder:validation:Optional
	Reserved *bool `json:"reserved,omitempty" tf:"reserved,omitempty"`

	// The name of the resource group in which to create the App Service Plan component. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sku block as documented below.
	// +kubebuilder:validation:Optional
	Sku *SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if the App Service Plan should be Zone Redundant. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type SkuInitParameters struct {

	// Specifies the number of workers associated with this App Service Plan.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies the plan's instance size.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the plan's pricing tier.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SkuObservation struct {

	// Specifies the number of workers associated with this App Service Plan.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies the plan's instance size.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the plan's pricing tier.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SkuParameters struct {

	// Specifies the number of workers associated with this App Service Plan.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Specifies the plan's instance size.
	// +kubebuilder:validation:Optional
	Size *string `json:"size" tf:"size,omitempty"`

	// Specifies the plan's pricing tier.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// AppServicePlanSpec defines the desired state of AppServicePlan
type AppServicePlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppServicePlanParameters `json:"forProvider"`
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
	InitProvider AppServicePlanInitParameters `json:"initProvider,omitempty"`
}

// AppServicePlanStatus defines the observed state of AppServicePlan.
type AppServicePlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppServicePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AppServicePlan is the Schema for the AppServicePlans API. Manages an App Service Plan component.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServicePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   AppServicePlanSpec   `json:"spec"`
	Status AppServicePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServicePlanList contains a list of AppServicePlans
type AppServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServicePlan `json:"items"`
}

// Repository type metadata.
var (
	AppServicePlan_Kind             = "AppServicePlan"
	AppServicePlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppServicePlan_Kind}.String()
	AppServicePlan_KindAPIVersion   = AppServicePlan_Kind + "." + CRDGroupVersion.String()
	AppServicePlan_GroupVersionKind = CRDGroupVersion.WithKind(AppServicePlan_Kind)
)

func init() {
	SchemeBuilder.Register(&AppServicePlan{}, &AppServicePlanList{})
}
