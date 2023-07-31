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
func (in *AppServiceCertificateOrder) DeepCopyInto(out *AppServiceCertificateOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrder.
func (in *AppServiceCertificateOrder) DeepCopy() *AppServiceCertificateOrder {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppServiceCertificateOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderInitParameters) DeepCopyInto(out *AppServiceCertificateOrderInitParameters) {
	*out = *in
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.Csr != nil {
		in, out := &in.Csr, &out.Csr
		*out = new(string)
		**out = **in
	}
	if in.DistinguishedName != nil {
		in, out := &in.DistinguishedName, &out.DistinguishedName
		*out = new(string)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ProductType != nil {
		in, out := &in.ProductType, &out.ProductType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ValidityInYears != nil {
		in, out := &in.ValidityInYears, &out.ValidityInYears
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderInitParameters.
func (in *AppServiceCertificateOrderInitParameters) DeepCopy() *AppServiceCertificateOrderInitParameters {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderList) DeepCopyInto(out *AppServiceCertificateOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppServiceCertificateOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderList.
func (in *AppServiceCertificateOrderList) DeepCopy() *AppServiceCertificateOrderList {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppServiceCertificateOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderObservation) DeepCopyInto(out *AppServiceCertificateOrderObservation) {
	*out = *in
	if in.AppServiceCertificateNotRenewableReasons != nil {
		in, out := &in.AppServiceCertificateNotRenewableReasons, &out.AppServiceCertificateNotRenewableReasons
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]CertificatesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Csr != nil {
		in, out := &in.Csr, &out.Csr
		*out = new(string)
		**out = **in
	}
	if in.DistinguishedName != nil {
		in, out := &in.DistinguishedName, &out.DistinguishedName
		*out = new(string)
		**out = **in
	}
	if in.DomainVerificationToken != nil {
		in, out := &in.DomainVerificationToken, &out.DomainVerificationToken
		*out = new(string)
		**out = **in
	}
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IntermediateThumbprint != nil {
		in, out := &in.IntermediateThumbprint, &out.IntermediateThumbprint
		*out = new(string)
		**out = **in
	}
	if in.IsPrivateKeyExternal != nil {
		in, out := &in.IsPrivateKeyExternal, &out.IsPrivateKeyExternal
		*out = new(bool)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ProductType != nil {
		in, out := &in.ProductType, &out.ProductType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.RootThumbprint != nil {
		in, out := &in.RootThumbprint, &out.RootThumbprint
		*out = new(string)
		**out = **in
	}
	if in.SignedCertificateThumbprint != nil {
		in, out := &in.SignedCertificateThumbprint, &out.SignedCertificateThumbprint
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ValidityInYears != nil {
		in, out := &in.ValidityInYears, &out.ValidityInYears
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderObservation.
func (in *AppServiceCertificateOrderObservation) DeepCopy() *AppServiceCertificateOrderObservation {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderParameters) DeepCopyInto(out *AppServiceCertificateOrderParameters) {
	*out = *in
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.Csr != nil {
		in, out := &in.Csr, &out.Csr
		*out = new(string)
		**out = **in
	}
	if in.DistinguishedName != nil {
		in, out := &in.DistinguishedName, &out.DistinguishedName
		*out = new(string)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ProductType != nil {
		in, out := &in.ProductType, &out.ProductType
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ValidityInYears != nil {
		in, out := &in.ValidityInYears, &out.ValidityInYears
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderParameters.
func (in *AppServiceCertificateOrderParameters) DeepCopy() *AppServiceCertificateOrderParameters {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderSpec) DeepCopyInto(out *AppServiceCertificateOrderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderSpec.
func (in *AppServiceCertificateOrderSpec) DeepCopy() *AppServiceCertificateOrderSpec {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppServiceCertificateOrderStatus) DeepCopyInto(out *AppServiceCertificateOrderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppServiceCertificateOrderStatus.
func (in *AppServiceCertificateOrderStatus) DeepCopy() *AppServiceCertificateOrderStatus {
	if in == nil {
		return nil
	}
	out := new(AppServiceCertificateOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesInitParameters) DeepCopyInto(out *CertificatesInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesInitParameters.
func (in *CertificatesInitParameters) DeepCopy() *CertificatesInitParameters {
	if in == nil {
		return nil
	}
	out := new(CertificatesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesObservation) DeepCopyInto(out *CertificatesObservation) {
	*out = *in
	if in.CertificateName != nil {
		in, out := &in.CertificateName, &out.CertificateName
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultID != nil {
		in, out := &in.KeyVaultID, &out.KeyVaultID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultSecretName != nil {
		in, out := &in.KeyVaultSecretName, &out.KeyVaultSecretName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesObservation.
func (in *CertificatesObservation) DeepCopy() *CertificatesObservation {
	if in == nil {
		return nil
	}
	out := new(CertificatesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesParameters) DeepCopyInto(out *CertificatesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesParameters.
func (in *CertificatesParameters) DeepCopy() *CertificatesParameters {
	if in == nil {
		return nil
	}
	out := new(CertificatesParameters)
	in.DeepCopyInto(out)
	return out
}
