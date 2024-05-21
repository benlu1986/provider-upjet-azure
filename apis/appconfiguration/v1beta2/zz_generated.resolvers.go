// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Configuration) ResolveReferences( // ResolveReferences of this Configuration.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.IdentityClientID),
				Extract:      resource.ExtractParamPath("client_id", true),
				Reference:    mg.Spec.ForProvider.Encryption.IdentityClientIDRef,
				Selector:     mg.Spec.ForProvider.Encryption.IdentityClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.IdentityClientID")
		}
		mg.Spec.ForProvider.Encryption.IdentityClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.IdentityClientIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifier),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifierRef,
				Selector:     mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifierSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifier")
		}
		mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.KeyVaultKeyIdentifierRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.IdentityClientID),
				Extract:      resource.ExtractParamPath("client_id", true),
				Reference:    mg.Spec.InitProvider.Encryption.IdentityClientIDRef,
				Selector:     mg.Spec.InitProvider.Encryption.IdentityClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.IdentityClientID")
		}
		mg.Spec.InitProvider.Encryption.IdentityClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.IdentityClientIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifier),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifierRef,
				Selector:     mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifierSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifier")
		}
		mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.KeyVaultKeyIdentifierRef = rsp.ResolvedReference

	}

	return nil
}
