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

type GlobalRoleInitParameters struct {

	// A description of the global role
	// Global role description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the global role.
	// Name of global role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, publishes the global role to all tenants
	// When true, publishes the global role to all tenants
	PublishToAllTenants *bool `json:"publishToAllTenants,omitempty" tf:"publish_to_all_tenants,omitempty"`

	// List of rights assigned to this role
	// list of rights assigned to this global role
	// +listType=set
	Rights []*string `json:"rights,omitempty" tf:"rights,omitempty"`

	// List of tenants to which this global role gets published. Ignored if publish_to_all_tenants is true.
	// list of tenants to which this global role is published
	// +listType=set
	Tenants []*string `json:"tenants,omitempty" tf:"tenants,omitempty"`
}

type GlobalRoleObservation struct {

	// Key used for internationalization
	// Key used for internationalization
	BundleKey *string `json:"bundleKey,omitempty" tf:"bundle_key,omitempty"`

	// A description of the global role
	// Global role description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the global role.
	// Name of global role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, publishes the global role to all tenants
	// When true, publishes the global role to all tenants
	PublishToAllTenants *bool `json:"publishToAllTenants,omitempty" tf:"publish_to_all_tenants,omitempty"`

	// Whether this global role is read-only
	// Whether this global role is read-only
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// List of rights assigned to this role
	// list of rights assigned to this global role
	// +listType=set
	Rights []*string `json:"rights,omitempty" tf:"rights,omitempty"`

	// List of tenants to which this global role gets published. Ignored if publish_to_all_tenants is true.
	// list of tenants to which this global role is published
	// +listType=set
	Tenants []*string `json:"tenants,omitempty" tf:"tenants,omitempty"`
}

type GlobalRoleParameters struct {

	// A description of the global role
	// Global role description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the global role.
	// Name of global role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, publishes the global role to all tenants
	// When true, publishes the global role to all tenants
	// +kubebuilder:validation:Optional
	PublishToAllTenants *bool `json:"publishToAllTenants,omitempty" tf:"publish_to_all_tenants,omitempty"`

	// List of rights assigned to this role
	// list of rights assigned to this global role
	// +kubebuilder:validation:Optional
	// +listType=set
	Rights []*string `json:"rights,omitempty" tf:"rights,omitempty"`

	// List of tenants to which this global role gets published. Ignored if publish_to_all_tenants is true.
	// list of tenants to which this global role is published
	// +kubebuilder:validation:Optional
	// +listType=set
	Tenants []*string `json:"tenants,omitempty" tf:"tenants,omitempty"`
}

// GlobalRoleSpec defines the desired state of GlobalRole
type GlobalRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalRoleParameters `json:"forProvider"`
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
	InitProvider GlobalRoleInitParameters `json:"initProvider,omitempty"`
}

// GlobalRoleStatus defines the observed state of GlobalRole.
type GlobalRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalRole is the Schema for the GlobalRoles API. Provides a VMware Cloud Director global role. This can be used to create, modify, and delete global roles.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type GlobalRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publishToAllTenants) || (has(self.initProvider) && has(self.initProvider.publishToAllTenants))",message="spec.forProvider.publishToAllTenants is a required parameter"
	Spec   GlobalRoleSpec   `json:"spec"`
	Status GlobalRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalRoleList contains a list of GlobalRoles
type GlobalRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalRole `json:"items"`
}

// Repository type metadata.
var (
	GlobalRole_Kind             = "GlobalRole"
	GlobalRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalRole_Kind}.String()
	GlobalRole_KindAPIVersion   = GlobalRole_Kind + "." + CRDGroupVersion.String()
	GlobalRole_GroupVersionKind = CRDGroupVersion.WithKind(GlobalRole_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalRole{}, &GlobalRoleList{})
}
