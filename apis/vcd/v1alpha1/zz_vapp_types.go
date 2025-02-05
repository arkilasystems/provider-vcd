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

type VAppInitParameters struct {

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key value map of vApp guest properties
	// Key/value settings for guest properties. Will be picked up by new VMs when created.
	// +mapType=granular
	GuestProperties map[string]*string `json:"guestProperties,omitempty" tf:"guest_properties,omitempty"`

	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are silently reduced to the highest value allowed.
	// Defines lease parameters for this vApp
	Lease []VAppLeaseInitParameters `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since v2.2+ metadata is added directly to vApp instead of first VM in vApp)
	// Key value map of metadata to assign to this vApp. Key and value can be any string.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp
	MetadataEntry []VAppMetadataEntryInitParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false. Works only on update when vApp already has VMs.
	// A boolean value stating if this vApp should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppLeaseInitParameters struct {

	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires
	RuntimeLeaseInSec *float64 `json:"runtimeLeaseInSec,omitempty" tf:"runtime_lease_in_sec,omitempty"`

	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires
	StorageLeaseInSec *float64 `json:"storageLeaseInSec,omitempty" tf:"storage_lease_in_sec,omitempty"`
}

type VAppLeaseObservation struct {

	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires
	RuntimeLeaseInSec *float64 `json:"runtimeLeaseInSec,omitempty" tf:"runtime_lease_in_sec,omitempty"`

	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires
	StorageLeaseInSec *float64 `json:"storageLeaseInSec,omitempty" tf:"storage_lease_in_sec,omitempty"`
}

type VAppLeaseParameters struct {

	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires
	// +kubebuilder:validation:Optional
	RuntimeLeaseInSec *float64 `json:"runtimeLeaseInSec" tf:"runtime_lease_in_sec,omitempty"`

	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
	// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires
	// +kubebuilder:validation:Optional
	StorageLeaseInSec *float64 `json:"storageLeaseInSec" tf:"storage_lease_in_sec,omitempty"`
}

type VAppMetadataEntryInitParameters struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	// Domain for this metadata entry. true, if it belongs to SYSTEM. false, if it belongs to GENERAL
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry.
	// Key of this metadata entry. Required if the metadata entry is not empty
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: MetadataStringValue, MetadataNumberValue, MetadataDateTimeValue, MetadataBooleanValue.
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: PRIVATE (hidden), READONLY (read only), READWRITE (read/write).
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry.
	// Value of this metadata entry. Required if the metadata entry is not empty
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VAppMetadataEntryObservation struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	// Domain for this metadata entry. true, if it belongs to SYSTEM. false, if it belongs to GENERAL
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry.
	// Key of this metadata entry. Required if the metadata entry is not empty
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: MetadataStringValue, MetadataNumberValue, MetadataDateTimeValue, MetadataBooleanValue.
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: PRIVATE (hidden), READONLY (read only), READWRITE (read/write).
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry.
	// Value of this metadata entry. Required if the metadata entry is not empty
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VAppMetadataEntryParameters struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL.
	// Domain for this metadata entry. true, if it belongs to SYSTEM. false, if it belongs to GENERAL
	// +kubebuilder:validation:Optional
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry.
	// Key of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: MetadataStringValue, MetadataNumberValue, MetadataDateTimeValue, MetadataBooleanValue.
	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: PRIVATE (hidden), READONLY (read only), READWRITE (read/write).
	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	// +kubebuilder:validation:Optional
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry.
	// Value of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VAppObservation struct {

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key value map of vApp guest properties
	// Key/value settings for guest properties. Will be picked up by new VMs when created.
	// +mapType=granular
	GuestProperties map[string]*string `json:"guestProperties,omitempty" tf:"guest_properties,omitempty"`

	// (Computed) The vApp Hyper Reference.
	// vApp Hyper Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Computed; v3.11+; VCD 10.5.1+) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
	// details on the origin of the vApp (e.g. vapp.origin.id, vapp.origin.name, vapp.origin.type).
	// A map that contains metadata that is automatically added by VCD (10.5.1+) and provides details on the origin of the vApp
	// +mapType=granular
	InheritedMetadata map[string]*string `json:"inheritedMetadata,omitempty" tf:"inherited_metadata,omitempty"`

	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are silently reduced to the highest value allowed.
	// Defines lease parameters for this vApp
	Lease []VAppLeaseObservation `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since v2.2+ metadata is added directly to vApp instead of first VM in vApp)
	// Key value map of metadata to assign to this vApp. Key and value can be any string.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp
	MetadataEntry []VAppMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false. Works only on update when vApp already has VMs.
	// A boolean value stating if this vApp should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// (Computed; v2.5+) The vApp status as a numeric code.
	// Shows the status code of the vApp
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// (Computed; v2.5+) The vApp status as text.
	// Shows the status of the vApp
	StatusText *string `json:"statusText,omitempty" tf:"status_text,omitempty"`

	// (v3.13.0+) A list of VM names included in this vApp
	// List of VMs in this vApp
	VMNames []*string `json:"vmNames,omitempty" tf:"vm_names,omitempty"`

	// (v3.13.0+) A list of vApp network names included in this vApp
	// List of vApp networks connected to this vApp
	VappNetworkNames []*string `json:"vappNetworkNames,omitempty" tf:"vapp_network_names,omitempty"`

	// (v3.13.0+) A list of vApp Org network names included in this vApp
	// List of vApp Org networks connected to this vApp
	VappOrgNetworkNames []*string `json:"vappOrgNetworkNames,omitempty" tf:"vapp_org_network_names,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppParameters struct {

	// An optional description for the vApp, up to 256 characters.
	// Optional description of the vApp
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key value map of vApp guest properties
	// Key/value settings for guest properties. Will be picked up by new VMs when created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	GuestProperties map[string]*string `json:"guestProperties,omitempty" tf:"guest_properties,omitempty"`

	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are silently reduced to the highest value allowed.
	// Defines lease parameters for this vApp
	// +kubebuilder:validation:Optional
	Lease []VAppLeaseParameters `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since v2.2+ metadata is added directly to vApp instead of first VM in vApp)
	// Key value map of metadata to assign to this vApp. Key and value can be any string.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp
	// +kubebuilder:validation:Optional
	MetadataEntry []VAppMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for the vApp
	// A name for the vApp, unique withing the VDC
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on. Default is false. Works only on update when vApp already has VMs.
	// A boolean value stating if this vApp should be powered on
	// +kubebuilder:validation:Optional
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// VAppSpec defines the desired state of VApp
type VAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VAppParameters `json:"forProvider"`
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
	InitProvider VAppInitParameters `json:"initProvider,omitempty"`
}

// VAppStatus defines the observed state of VApp.
type VAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VApp is the Schema for the VApps API. Provides a VMware Cloud Director vApp resource. This can be used to create, modify, and delete vApps.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VAppSpec   `json:"spec"`
	Status VAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VAppList contains a list of VApps
type VAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VApp `json:"items"`
}

// Repository type metadata.
var (
	VApp_Kind             = "VApp"
	VApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VApp_Kind}.String()
	VApp_KindAPIVersion   = VApp_Kind + "." + CRDGroupVersion.String()
	VApp_GroupVersionKind = CRDGroupVersion.WithKind(VApp_Kind)
)

func init() {
	SchemeBuilder.Register(&VApp{}, &VAppList{})
}
