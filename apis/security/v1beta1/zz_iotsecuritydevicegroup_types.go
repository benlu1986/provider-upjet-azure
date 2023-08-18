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

type AllowRuleInitParameters struct {

	// Specifies which IP is not allowed to be connected to in current device group for inbound connection.
	ConnectionFromIpsNotAllowed []*string `json:"connectionFromIpsNotAllowed,omitempty" tf:"connection_from_ips_not_allowed,omitempty"`

	// Specifies which IP is not allowed to be connected to in current device group for outbound connection.
	ConnectionToIpsNotAllowed []*string `json:"connectionToIpsNotAllowed,omitempty" tf:"connection_to_ips_not_allowed,omitempty"`

	// Specifies which local user is not allowed to login in current device group.
	LocalUsersNotAllowed []*string `json:"localUsersNotAllowed,omitempty" tf:"local_users_not_allowed,omitempty"`

	// Specifies which process is not allowed to be executed in current device group.
	ProcessesNotAllowed []*string `json:"processesNotAllowed,omitempty" tf:"processes_not_allowed,omitempty"`
}

type AllowRuleObservation struct {

	// Specifies which IP is not allowed to be connected to in current device group for inbound connection.
	ConnectionFromIpsNotAllowed []*string `json:"connectionFromIpsNotAllowed,omitempty" tf:"connection_from_ips_not_allowed,omitempty"`

	// Specifies which IP is not allowed to be connected to in current device group for outbound connection.
	ConnectionToIpsNotAllowed []*string `json:"connectionToIpsNotAllowed,omitempty" tf:"connection_to_ips_not_allowed,omitempty"`

	// Specifies which local user is not allowed to login in current device group.
	LocalUsersNotAllowed []*string `json:"localUsersNotAllowed,omitempty" tf:"local_users_not_allowed,omitempty"`

	// Specifies which process is not allowed to be executed in current device group.
	ProcessesNotAllowed []*string `json:"processesNotAllowed,omitempty" tf:"processes_not_allowed,omitempty"`
}

type AllowRuleParameters struct {

	// Specifies which IP is not allowed to be connected to in current device group for inbound connection.
	// +kubebuilder:validation:Optional
	ConnectionFromIpsNotAllowed []*string `json:"connectionFromIpsNotAllowed,omitempty" tf:"connection_from_ips_not_allowed,omitempty"`

	// Specifies which IP is not allowed to be connected to in current device group for outbound connection.
	// +kubebuilder:validation:Optional
	ConnectionToIpsNotAllowed []*string `json:"connectionToIpsNotAllowed,omitempty" tf:"connection_to_ips_not_allowed,omitempty"`

	// Specifies which local user is not allowed to login in current device group.
	// +kubebuilder:validation:Optional
	LocalUsersNotAllowed []*string `json:"localUsersNotAllowed,omitempty" tf:"local_users_not_allowed,omitempty"`

	// Specifies which process is not allowed to be executed in current device group.
	// +kubebuilder:validation:Optional
	ProcessesNotAllowed []*string `json:"processesNotAllowed,omitempty" tf:"processes_not_allowed,omitempty"`
}

type IOTSecurityDeviceGroupInitParameters struct {

	// an allow_rule blocks as defined below.
	AllowRule []AllowRuleInitParameters `json:"allowRule,omitempty" tf:"allow_rule,omitempty"`

	// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more range_rule blocks as defined below.
	RangeRule []RangeRuleInitParameters `json:"rangeRule,omitempty" tf:"range_rule,omitempty"`
}

type IOTSecurityDeviceGroupObservation struct {

	// an allow_rule blocks as defined below.
	AllowRule []AllowRuleObservation `json:"allowRule,omitempty" tf:"allow_rule,omitempty"`

	// The ID of the Iot Security Device Group resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more range_rule blocks as defined below.
	RangeRule []RangeRuleObservation `json:"rangeRule,omitempty" tf:"range_rule,omitempty"`
}

type IOTSecurityDeviceGroupParameters struct {

	// an allow_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	AllowRule []AllowRuleParameters `json:"allowRule,omitempty" tf:"allow_rule,omitempty"`

	// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more range_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	RangeRule []RangeRuleParameters `json:"rangeRule,omitempty" tf:"range_rule,omitempty"`
}

type RangeRuleInitParameters struct {

	// Specifies the time range. represented in ISO 8601 duration format.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The maximum threshold in the given time window.
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// The minimum threshold in the given time window.
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// The type of supported rule type. Possible Values are ActiveConnectionsNotInAllowedRange, AmqpC2DMessagesNotInAllowedRange, MqttC2DMessagesNotInAllowedRange, HttpC2DMessagesNotInAllowedRange, AmqpC2DRejectedMessagesNotInAllowedRange, MqttC2DRejectedMessagesNotInAllowedRange, HttpC2DRejectedMessagesNotInAllowedRange, AmqpD2CMessagesNotInAllowedRange, MqttD2CMessagesNotInAllowedRange, HttpD2CMessagesNotInAllowedRange, DirectMethodInvokesNotInAllowedRange, FailedLocalLoginsNotInAllowedRange, FileUploadsNotInAllowedRange, QueuePurgesNotInAllowedRange, TwinUpdatesNotInAllowedRange and UnauthorizedOperationsNotInAllowedRange.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RangeRuleObservation struct {

	// Specifies the time range. represented in ISO 8601 duration format.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// The maximum threshold in the given time window.
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// The minimum threshold in the given time window.
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// The type of supported rule type. Possible Values are ActiveConnectionsNotInAllowedRange, AmqpC2DMessagesNotInAllowedRange, MqttC2DMessagesNotInAllowedRange, HttpC2DMessagesNotInAllowedRange, AmqpC2DRejectedMessagesNotInAllowedRange, MqttC2DRejectedMessagesNotInAllowedRange, HttpC2DRejectedMessagesNotInAllowedRange, AmqpD2CMessagesNotInAllowedRange, MqttD2CMessagesNotInAllowedRange, HttpD2CMessagesNotInAllowedRange, DirectMethodInvokesNotInAllowedRange, FailedLocalLoginsNotInAllowedRange, FileUploadsNotInAllowedRange, QueuePurgesNotInAllowedRange, TwinUpdatesNotInAllowedRange and UnauthorizedOperationsNotInAllowedRange.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RangeRuleParameters struct {

	// Specifies the time range. represented in ISO 8601 duration format.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// The maximum threshold in the given time window.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum threshold in the given time window.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min" tf:"min,omitempty"`

	// The type of supported rule type. Possible Values are ActiveConnectionsNotInAllowedRange, AmqpC2DMessagesNotInAllowedRange, MqttC2DMessagesNotInAllowedRange, HttpC2DMessagesNotInAllowedRange, AmqpC2DRejectedMessagesNotInAllowedRange, MqttC2DRejectedMessagesNotInAllowedRange, HttpC2DRejectedMessagesNotInAllowedRange, AmqpD2CMessagesNotInAllowedRange, MqttD2CMessagesNotInAllowedRange, HttpD2CMessagesNotInAllowedRange, DirectMethodInvokesNotInAllowedRange, FailedLocalLoginsNotInAllowedRange, FileUploadsNotInAllowedRange, QueuePurgesNotInAllowedRange, TwinUpdatesNotInAllowedRange and UnauthorizedOperationsNotInAllowedRange.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// IOTSecurityDeviceGroupSpec defines the desired state of IOTSecurityDeviceGroup
type IOTSecurityDeviceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTSecurityDeviceGroupParameters `json:"forProvider"`
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
	InitProvider IOTSecurityDeviceGroupInitParameters `json:"initProvider,omitempty"`
}

// IOTSecurityDeviceGroupStatus defines the observed state of IOTSecurityDeviceGroup.
type IOTSecurityDeviceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTSecurityDeviceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroup is the Schema for the IOTSecurityDeviceGroups API. Manages a Iot Security Device Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTSecurityDeviceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   IOTSecurityDeviceGroupSpec   `json:"spec"`
	Status IOTSecurityDeviceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTSecurityDeviceGroupList contains a list of IOTSecurityDeviceGroups
type IOTSecurityDeviceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTSecurityDeviceGroup `json:"items"`
}

// Repository type metadata.
var (
	IOTSecurityDeviceGroup_Kind             = "IOTSecurityDeviceGroup"
	IOTSecurityDeviceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTSecurityDeviceGroup_Kind}.String()
	IOTSecurityDeviceGroup_KindAPIVersion   = IOTSecurityDeviceGroup_Kind + "." + CRDGroupVersion.String()
	IOTSecurityDeviceGroup_GroupVersionKind = CRDGroupVersion.WithKind(IOTSecurityDeviceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTSecurityDeviceGroup{}, &IOTSecurityDeviceGroupList{})
}
