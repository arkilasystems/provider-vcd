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

type CaptureVappInitParameters struct {

	// Defines if Trusted Platform Module should be copied (false) or created (true). Default 'false'. VCD 10.4.2+
	CopyTpmOnInstantiate *bool `json:"copyTpmOnInstantiate,omitempty" tf:"copy_tpm_on_instantiate,omitempty"`

	// Default false - means "Make identical copy". true
	// means "Customize VM settings". Note true can only be set when source vApp is powered off
	// Marks if instantiating applies customization settings ('true'). Default is 'false` - create an identical copy.
	CustomizeOnInstantiate *bool `json:"customizeOnInstantiate,omitempty" tf:"customize_on_instantiate,omitempty"`

	// Optionally newly created template can overwrite. It can
	// either be id of vcd_catalog_item resource or catalog_item_id of
	// vcd_catalog_vapp_template resource
	// An existing catalog item ID to overwrite
	OverwriteCatalogItemID *string `json:"overwriteCatalogItemId,omitempty" tf:"overwrite_catalog_item_id,omitempty"`

	// Source vApp ID (can be referenced by vcd_vapp.id or
	// vcd_vm.vapp_id/vcd_vapp_vm.vapp_id)
	// Source vApp ID (can be a vApp ID or 'vapp_id' field of standalone VM 'vcd_vm')
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`
}

type CaptureVappObservation struct {

	// Defines if Trusted Platform Module should be copied (false) or created (true). Default 'false'. VCD 10.4.2+
	CopyTpmOnInstantiate *bool `json:"copyTpmOnInstantiate,omitempty" tf:"copy_tpm_on_instantiate,omitempty"`

	// Default false - means "Make identical copy". true
	// means "Customize VM settings". Note true can only be set when source vApp is powered off
	// Marks if instantiating applies customization settings ('true'). Default is 'false` - create an identical copy.
	CustomizeOnInstantiate *bool `json:"customizeOnInstantiate,omitempty" tf:"customize_on_instantiate,omitempty"`

	// Optionally newly created template can overwrite. It can
	// either be id of vcd_catalog_item resource or catalog_item_id of
	// vcd_catalog_vapp_template resource
	// An existing catalog item ID to overwrite
	OverwriteCatalogItemID *string `json:"overwriteCatalogItemId,omitempty" tf:"overwrite_catalog_item_id,omitempty"`

	// Source vApp ID (can be referenced by vcd_vapp.id or
	// vcd_vm.vapp_id/vcd_vapp_vm.vapp_id)
	// Source vApp ID (can be a vApp ID or 'vapp_id' field of standalone VM 'vcd_vm')
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`
}

type CaptureVappParameters struct {

	// Defines if Trusted Platform Module should be copied (false) or created (true). Default 'false'. VCD 10.4.2+
	// +kubebuilder:validation:Optional
	CopyTpmOnInstantiate *bool `json:"copyTpmOnInstantiate,omitempty" tf:"copy_tpm_on_instantiate,omitempty"`

	// Default false - means "Make identical copy". true
	// means "Customize VM settings". Note true can only be set when source vApp is powered off
	// Marks if instantiating applies customization settings ('true'). Default is 'false` - create an identical copy.
	// +kubebuilder:validation:Optional
	CustomizeOnInstantiate *bool `json:"customizeOnInstantiate,omitempty" tf:"customize_on_instantiate,omitempty"`

	// Optionally newly created template can overwrite. It can
	// either be id of vcd_catalog_item resource or catalog_item_id of
	// vcd_catalog_vapp_template resource
	// An existing catalog item ID to overwrite
	// +kubebuilder:validation:Optional
	OverwriteCatalogItemID *string `json:"overwriteCatalogItemId,omitempty" tf:"overwrite_catalog_item_id,omitempty"`

	// Source vApp ID (can be referenced by vcd_vapp.id or
	// vcd_vm.vapp_id/vcd_vapp_vm.vapp_id)
	// Source vApp ID (can be a vApp ID or 'vapp_id' field of standalone VM 'vcd_vm')
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`
}

type CatalogvAppTemplateInitParameters struct {

	// A configuration block to create template from existing
	// vApp (Standalone VM or vApp)
	// Provides configuration options for creating a vApp Template from existing vApp
	CaptureVapp []CaptureVappInitParameters `json:"captureVapp,omitempty" tf:"capture_vapp,omitempty"`

	// ID of the Catalog where to upload the OVA file
	// ID of the Catalog where to upload the OVA file
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Description of the vApp Template. Not to be used with ovf_url when target OVA has a description
	// Description of the vApp Template. Not to be used with `ovf_url` when target OVA has a description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The information about the vApp Template lease. It includes the field below. When this section is
	// included, the field is mandatory. If lease value is higher than the one allowed for the whole Org, we get an error
	// Defines lease parameters for this vApp template
	Lease []LeaseInitParameters `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key/value map of metadata to assign to the associated vApp Template
	// Key and value pairs for the metadata of this vApp Template
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp Template
	MetadataEntry []CatalogvAppTemplateMetadataEntryInitParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// vApp Template name in Catalog
	// vApp Template name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Absolute or relative path to file to upload
	// Absolute or relative path to OVA
	OvaPath *string `json:"ovaPath,omitempty" tf:"ova_path,omitempty"`

	// URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
	// URL of OVF file
	OvfURL *string `json:"ovfUrl,omitempty" tf:"ovf_url,omitempty"`

	// - Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB
	// Size of upload file piece size in megabytes
	UploadPieceSize *float64 `json:"uploadPieceSize,omitempty" tf:"upload_piece_size,omitempty"`
}

type CatalogvAppTemplateMetadataEntryInitParameters struct {

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

type CatalogvAppTemplateMetadataEntryObservation struct {

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

type CatalogvAppTemplateMetadataEntryParameters struct {

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

type CatalogvAppTemplateObservation struct {

	// A configuration block to create template from existing
	// vApp (Standalone VM or vApp)
	// Provides configuration options for creating a vApp Template from existing vApp
	CaptureVapp []CaptureVappObservation `json:"captureVapp,omitempty" tf:"capture_vapp,omitempty"`

	// ID of the Catalog where to upload the OVA file
	// ID of the Catalog where to upload the OVA file
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Catalog Item ID
	// Catalog Item ID of this vApp template
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// Timestamp of when the vApp Template was created
	// Timestamp of when the vApp Template was created
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// Description of the vApp Template. Not to be used with ovf_url when target OVA has a description
	// Description of the vApp Template. Not to be used with `ovf_url` when target OVA has a description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key/value map of metadata to assign to the associated vApp Template
	// A map that contains metadata that is automatically added by VCD (10.5.1+) and provides details on the origin of the VM
	// +mapType=granular
	InheritedMetadata map[string]*string `json:"inheritedMetadata,omitempty" tf:"inherited_metadata,omitempty"`

	// The information about the vApp Template lease. It includes the field below. When this section is
	// included, the field is mandatory. If lease value is higher than the one allowed for the whole Org, we get an error
	// Defines lease parameters for this vApp template
	Lease []LeaseObservation `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key/value map of metadata to assign to the associated vApp Template
	// Key and value pairs for the metadata of this vApp Template
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp Template
	MetadataEntry []CatalogvAppTemplateMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// vApp Template name in Catalog
	// vApp Template name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Absolute or relative path to file to upload
	// Absolute or relative path to OVA
	OvaPath *string `json:"ovaPath,omitempty" tf:"ova_path,omitempty"`

	// URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
	// URL of OVF file
	OvfURL *string `json:"ovfUrl,omitempty" tf:"ovf_url,omitempty"`

	// - Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB
	// Size of upload file piece size in megabytes
	UploadPieceSize *float64 `json:"uploadPieceSize,omitempty" tf:"upload_piece_size,omitempty"`

	// Set of VM names within the vApp template
	// Set of VM names within the vApp template
	// +listType=set
	VMNames []*string `json:"vmNames,omitempty" tf:"vm_names,omitempty"`

	// The VDC ID to which this vApp Template belongs
	// ID of the VDC to which the vApp Template belongs
	VdcID *string `json:"vdcId,omitempty" tf:"vdc_id,omitempty"`
}

type CatalogvAppTemplateParameters struct {

	// A configuration block to create template from existing
	// vApp (Standalone VM or vApp)
	// Provides configuration options for creating a vApp Template from existing vApp
	// +kubebuilder:validation:Optional
	CaptureVapp []CaptureVappParameters `json:"captureVapp,omitempty" tf:"capture_vapp,omitempty"`

	// ID of the Catalog where to upload the OVA file
	// ID of the Catalog where to upload the OVA file
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Description of the vApp Template. Not to be used with ovf_url when target OVA has a description
	// Description of the vApp Template. Not to be used with `ovf_url` when target OVA has a description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The information about the vApp Template lease. It includes the field below. When this section is
	// included, the field is mandatory. If lease value is higher than the one allowed for the whole Org, we get an error
	// Defines lease parameters for this vApp template
	// +kubebuilder:validation:Optional
	Lease []LeaseParameters `json:"lease,omitempty" tf:"lease,omitempty"`

	// (Deprecated) Use metadata_entry instead. Key/value map of metadata to assign to the associated vApp Template
	// Key and value pairs for the metadata of this vApp Template
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given vApp Template
	// +kubebuilder:validation:Optional
	MetadataEntry []CatalogvAppTemplateMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// vApp Template name in Catalog
	// vApp Template name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Absolute or relative path to file to upload
	// Absolute or relative path to OVA
	// +kubebuilder:validation:Optional
	OvaPath *string `json:"ovaPath,omitempty" tf:"ova_path,omitempty"`

	// URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
	// URL of OVF file
	// +kubebuilder:validation:Optional
	OvfURL *string `json:"ovfUrl,omitempty" tf:"ovf_url,omitempty"`

	// - Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB
	// Size of upload file piece size in megabytes
	// +kubebuilder:validation:Optional
	UploadPieceSize *float64 `json:"uploadPieceSize,omitempty" tf:"upload_piece_size,omitempty"`
}

type LeaseInitParameters struct {

	// How long the vApp Template is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by parent Org). Regular values accepted from 3600+.
	// How long the vApp template is available before being automatically deleted or marked as expired. 0 means never expires (or expires at the maximum limit provided by the parent Org)
	StorageLeaseInSec *float64 `json:"storageLeaseInSec,omitempty" tf:"storage_lease_in_sec,omitempty"`
}

type LeaseObservation struct {

	// How long the vApp Template is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by parent Org). Regular values accepted from 3600+.
	// How long the vApp template is available before being automatically deleted or marked as expired. 0 means never expires (or expires at the maximum limit provided by the parent Org)
	StorageLeaseInSec *float64 `json:"storageLeaseInSec,omitempty" tf:"storage_lease_in_sec,omitempty"`
}

type LeaseParameters struct {

	// How long the vApp Template is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by parent Org). Regular values accepted from 3600+.
	// How long the vApp template is available before being automatically deleted or marked as expired. 0 means never expires (or expires at the maximum limit provided by the parent Org)
	// +kubebuilder:validation:Optional
	StorageLeaseInSec *float64 `json:"storageLeaseInSec" tf:"storage_lease_in_sec,omitempty"`
}

// CatalogvAppTemplateSpec defines the desired state of CatalogvAppTemplate
type CatalogvAppTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogvAppTemplateParameters `json:"forProvider"`
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
	InitProvider CatalogvAppTemplateInitParameters `json:"initProvider,omitempty"`
}

// CatalogvAppTemplateStatus defines the observed state of CatalogvAppTemplate.
type CatalogvAppTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogvAppTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CatalogvAppTemplate is the Schema for the CatalogvAppTemplates API. Provides a VMware Cloud Director vApp Template resource. This can be used to upload and delete OVA files inside a catalog.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type CatalogvAppTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogId) || (has(self.initProvider) && has(self.initProvider.catalogId))",message="spec.forProvider.catalogId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CatalogvAppTemplateSpec   `json:"spec"`
	Status CatalogvAppTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogvAppTemplateList contains a list of CatalogvAppTemplates
type CatalogvAppTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogvAppTemplate `json:"items"`
}

// Repository type metadata.
var (
	CatalogvAppTemplate_Kind             = "CatalogvAppTemplate"
	CatalogvAppTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogvAppTemplate_Kind}.String()
	CatalogvAppTemplate_KindAPIVersion   = CatalogvAppTemplate_Kind + "." + CRDGroupVersion.String()
	CatalogvAppTemplate_GroupVersionKind = CRDGroupVersion.WithKind(CatalogvAppTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogvAppTemplate{}, &CatalogvAppTemplateList{})
}
