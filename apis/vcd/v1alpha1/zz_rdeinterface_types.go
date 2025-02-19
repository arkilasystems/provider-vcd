// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RdeInterfaceInitParameters struct {

	// The name of the RDE Interface.
	// The name of the Runtime Defined Entity Interface
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// A unique namespace associated with the Runtime Defined Entity Interface. Combination of `vendor`, `nss` and `version` must be unique
	Nss *string `json:"nss,omitempty" tf:"nss,omitempty"`

	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// The vendor name. Combination of `vendor`, `nss` and `version` must be unique
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`

	// The version of the RDE Interface. Must follow semantic versioning syntax.
	// The Runtime Defined Entity Interface's version. The version must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RdeInterfaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the RDE Interface.
	// The name of the Runtime Defined Entity Interface
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// A unique namespace associated with the Runtime Defined Entity Interface. Combination of `vendor`, `nss` and `version` must be unique
	Nss *string `json:"nss,omitempty" tf:"nss,omitempty"`

	// Specifies if the RDE Interface can be only read.
	// True if the Runtime Defined Entity Interface cannot be modified
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`

	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// The vendor name. Combination of `vendor`, `nss` and `version` must be unique
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`

	// The version of the RDE Interface. Must follow semantic versioning syntax.
	// The Runtime Defined Entity Interface's version. The version must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RdeInterfaceParameters struct {

	// The name of the RDE Interface.
	// The name of the Runtime Defined Entity Interface
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// A unique namespace associated with the Runtime Defined Entity Interface. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Nss *string `json:"nss,omitempty" tf:"nss,omitempty"`

	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	// The vendor name. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`

	// The version of the RDE Interface. Must follow semantic versioning syntax.
	// The Runtime Defined Entity Interface's version. The version must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// RdeInterfaceSpec defines the desired state of RdeInterface
type RdeInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdeInterfaceParameters `json:"forProvider"`
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
	InitProvider RdeInterfaceInitParameters `json:"initProvider,omitempty"`
}

// RdeInterfaceStatus defines the observed state of RdeInterface.
type RdeInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdeInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RdeInterface is the Schema for the RdeInterfaces API. Provides the capability of creating, updating, and deleting Runtime Defined Entity Interfaces in VMware Cloud Director.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type RdeInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nss) || (has(self.initProvider) && has(self.initProvider.nss))",message="spec.forProvider.nss is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vendor) || (has(self.initProvider) && has(self.initProvider.vendor))",message="spec.forProvider.vendor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   RdeInterfaceSpec   `json:"spec"`
	Status RdeInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdeInterfaceList contains a list of RdeInterfaces
type RdeInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdeInterface `json:"items"`
}

// Repository type metadata.
var (
	RdeInterface_Kind             = "RdeInterface"
	RdeInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RdeInterface_Kind}.String()
	RdeInterface_KindAPIVersion   = RdeInterface_Kind + "." + CRDGroupVersion.String()
	RdeInterface_GroupVersionKind = CRDGroupVersion.WithKind(RdeInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&RdeInterface{}, &RdeInterfaceList{})
}
