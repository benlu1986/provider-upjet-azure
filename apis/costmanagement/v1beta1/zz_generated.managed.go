/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ResourceGroupCostManagementExport.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ResourceGroupCostManagementExport) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ResourceGroupCostManagementExport.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ResourceGroupCostManagementExport) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceGroupCostManagementExport.
func (mg *ResourceGroupCostManagementExport) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SubscriptionCostManagementExport.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SubscriptionCostManagementExport) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SubscriptionCostManagementExport.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SubscriptionCostManagementExport) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SubscriptionCostManagementExport.
func (mg *SubscriptionCostManagementExport) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
