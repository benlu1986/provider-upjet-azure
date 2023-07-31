/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this CloudApplicationLiveView.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *CloudApplicationLiveView) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this CloudApplicationLiveView.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *CloudApplicationLiveView) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this CloudApplicationLiveView.
func (mg *CloudApplicationLiveView) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
