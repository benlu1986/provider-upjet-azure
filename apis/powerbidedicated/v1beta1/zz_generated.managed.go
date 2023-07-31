/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PowerBIEmbedded.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PowerBIEmbedded) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PowerBIEmbedded.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PowerBIEmbedded) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PowerBIEmbedded.
func (mg *PowerBIEmbedded) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
