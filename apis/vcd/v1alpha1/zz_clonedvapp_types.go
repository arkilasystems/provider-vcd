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

type ClonedvAppInitParameters struct {

	// A boolean value of true or false stating if the source entity should be deleted after creation.
	// A source vApp can only be deleted if it is fully powered off.
	// If true, it will delete the source (vApp or template) after creating the new vApp
	DeleteSource *bool `json:"deleteSource,omitempty" tf:"delete_source,omitempty"`

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false.
	// A boolean value stating if this vApp should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The ID of the source to use.
	// The identifier of the source to use for the creation of this vApp
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// The type of the source to use: one of template or vapp.
	// The type of the source to use for the creation of this vApp (one of 'vapp' or 'template')
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type ClonedvAppObservation struct {

	// A boolean value of true or false stating if the source entity should be deleted after creation.
	// A source vApp can only be deleted if it is fully powered off.
	// If true, it will delete the source (vApp or template) after creating the new vApp
	DeleteSource *bool `json:"deleteSource,omitempty" tf:"delete_source,omitempty"`

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Computed) The vApp Hyper Reference.
	// vApp Hyper Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false.
	// A boolean value stating if this vApp should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The ID of the source to use.
	// The identifier of the source to use for the creation of this vApp
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// The type of the source to use: one of template or vapp.
	// The type of the source to use for the creation of this vApp (one of 'vapp' or 'template')
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// (Computed) The vApp status as a numeric code.
	// Shows the status code of the vApp
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// (Computed) The vApp status as text.
	// Shows the status of the vApp
	StatusText *string `json:"statusText,omitempty" tf:"status_text,omitempty"`

	// (Computed) The list of VM names included in this vApp, in alphabetic order.
	// List of VMs contained in this vApp (in alphabetic order)
	VMList []*string `json:"vmList,omitempty" tf:"vm_list,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type ClonedvAppParameters struct {

	// A boolean value of true or false stating if the source entity should be deleted after creation.
	// A source vApp can only be deleted if it is fully powered off.
	// If true, it will delete the source (vApp or template) after creating the new vApp
	// +kubebuilder:validation:Optional
	DeleteSource *bool `json:"deleteSource,omitempty" tf:"delete_source,omitempty"`

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false.
	// A boolean value stating if this vApp should be powered on
	// +kubebuilder:validation:Optional
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The ID of the source to use.
	// The identifier of the source to use for the creation of this vApp
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// The type of the source to use: one of template or vapp.
	// The type of the source to use for the creation of this vApp (one of 'vapp' or 'template')
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// ClonedvAppSpec defines the desired state of ClonedvApp
type ClonedvAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClonedvAppParameters `json:"forProvider"`
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
	InitProvider ClonedvAppInitParameters `json:"initProvider,omitempty"`
}

// ClonedvAppStatus defines the observed state of ClonedvApp.
type ClonedvAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClonedvAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClonedvApp is the Schema for the ClonedvApps API. Provides a VMware Cloud Director Cloned vApp resource. This can be used to create vApps from either a vApp template or another vApp.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type ClonedvApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceId) || (has(self.initProvider) && has(self.initProvider.sourceId))",message="spec.forProvider.sourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceType) || (has(self.initProvider) && has(self.initProvider.sourceType))",message="spec.forProvider.sourceType is a required parameter"
	Spec   ClonedvAppSpec   `json:"spec"`
	Status ClonedvAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClonedvAppList contains a list of ClonedvApps
type ClonedvAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClonedvApp `json:"items"`
}

// Repository type metadata.
var (
	ClonedvApp_Kind             = "ClonedvApp"
	ClonedvApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClonedvApp_Kind}.String()
	ClonedvApp_KindAPIVersion   = ClonedvApp_Kind + "." + CRDGroupVersion.String()
	ClonedvApp_GroupVersionKind = CRDGroupVersion.WithKind(ClonedvApp_Kind)
)

func init() {
	SchemeBuilder.Register(&ClonedvApp{}, &ClonedvAppList{})
}
