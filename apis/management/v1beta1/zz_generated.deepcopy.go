//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroup) DeepCopyInto(out *ManagementGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroup.
func (in *ManagementGroup) DeepCopy() *ManagementGroup {
	if in == nil {
		return nil
	}
	out := new(ManagementGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupInitParameters) DeepCopyInto(out *ManagementGroupInitParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionIds != nil {
		in, out := &in.SubscriptionIds, &out.SubscriptionIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupInitParameters.
func (in *ManagementGroupInitParameters) DeepCopy() *ManagementGroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupList) DeepCopyInto(out *ManagementGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagementGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupList.
func (in *ManagementGroupList) DeepCopy() *ManagementGroupList {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupObservation) DeepCopyInto(out *ManagementGroupObservation) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ParentManagementGroupID != nil {
		in, out := &in.ParentManagementGroupID, &out.ParentManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionIds != nil {
		in, out := &in.SubscriptionIds, &out.SubscriptionIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupObservation.
func (in *ManagementGroupObservation) DeepCopy() *ManagementGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupParameters) DeepCopyInto(out *ManagementGroupParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ParentManagementGroupID != nil {
		in, out := &in.ParentManagementGroupID, &out.ParentManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ParentManagementGroupIDRef != nil {
		in, out := &in.ParentManagementGroupIDRef, &out.ParentManagementGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ParentManagementGroupIDSelector != nil {
		in, out := &in.ParentManagementGroupIDSelector, &out.ParentManagementGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriptionIds != nil {
		in, out := &in.SubscriptionIds, &out.SubscriptionIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupParameters.
func (in *ManagementGroupParameters) DeepCopy() *ManagementGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSpec) DeepCopyInto(out *ManagementGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSpec.
func (in *ManagementGroupSpec) DeepCopy() *ManagementGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupStatus) DeepCopyInto(out *ManagementGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupStatus.
func (in *ManagementGroupStatus) DeepCopy() *ManagementGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociation) DeepCopyInto(out *ManagementGroupSubscriptionAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociation.
func (in *ManagementGroupSubscriptionAssociation) DeepCopy() *ManagementGroupSubscriptionAssociation {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementGroupSubscriptionAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationInitParameters) DeepCopyInto(out *ManagementGroupSubscriptionAssociationInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationInitParameters.
func (in *ManagementGroupSubscriptionAssociationInitParameters) DeepCopy() *ManagementGroupSubscriptionAssociationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationList) DeepCopyInto(out *ManagementGroupSubscriptionAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagementGroupSubscriptionAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationList.
func (in *ManagementGroupSubscriptionAssociationList) DeepCopy() *ManagementGroupSubscriptionAssociationList {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementGroupSubscriptionAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationObservation) DeepCopyInto(out *ManagementGroupSubscriptionAssociationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationObservation.
func (in *ManagementGroupSubscriptionAssociationObservation) DeepCopy() *ManagementGroupSubscriptionAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationParameters) DeepCopyInto(out *ManagementGroupSubscriptionAssociationParameters) {
	*out = *in
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupIDRef != nil {
		in, out := &in.ManagementGroupIDRef, &out.ManagementGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagementGroupIDSelector != nil {
		in, out := &in.ManagementGroupIDSelector, &out.ManagementGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionIDRef != nil {
		in, out := &in.SubscriptionIDRef, &out.SubscriptionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriptionIDSelector != nil {
		in, out := &in.SubscriptionIDSelector, &out.SubscriptionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationParameters.
func (in *ManagementGroupSubscriptionAssociationParameters) DeepCopy() *ManagementGroupSubscriptionAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationSpec) DeepCopyInto(out *ManagementGroupSubscriptionAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	out.InitProvider = in.InitProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationSpec.
func (in *ManagementGroupSubscriptionAssociationSpec) DeepCopy() *ManagementGroupSubscriptionAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementGroupSubscriptionAssociationStatus) DeepCopyInto(out *ManagementGroupSubscriptionAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementGroupSubscriptionAssociationStatus.
func (in *ManagementGroupSubscriptionAssociationStatus) DeepCopy() *ManagementGroupSubscriptionAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ManagementGroupSubscriptionAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
