// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=hdinsight.azure.upbound.io
// +versionName=v1beta2
package v1beta2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	CRDGroup   = "hdinsight.azure.upbound.io"
	CRDVersion = "v1beta2"
)

var (
	// CRDGroupVersion is the API Group Version used to register the objects
	CRDGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: CRDGroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
