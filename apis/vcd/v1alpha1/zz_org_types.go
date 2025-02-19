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

type AccountLockoutInitParameters struct {

	// Whether account lockout is enabled or not
	// Whether account lockout is enabled or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of login attempts that will trigger an account lockout for the given user
	// Number of login attempts that will trigger an account lockout for the given user
	InvalidLoginsBeforeLockout *float64 `json:"invalidLoginsBeforeLockout,omitempty" tf:"invalid_logins_before_lockout,omitempty"`

	// Once a user is locked out, they will not be able to log back in for this time period
	// Once a user is locked out, they will not be able to log back in for this time period
	LockoutIntervalMinutes *float64 `json:"lockoutIntervalMinutes,omitempty" tf:"lockout_interval_minutes,omitempty"`
}

type AccountLockoutObservation struct {

	// Whether account lockout is enabled or not
	// Whether account lockout is enabled or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of login attempts that will trigger an account lockout for the given user
	// Number of login attempts that will trigger an account lockout for the given user
	InvalidLoginsBeforeLockout *float64 `json:"invalidLoginsBeforeLockout,omitempty" tf:"invalid_logins_before_lockout,omitempty"`

	// Once a user is locked out, they will not be able to log back in for this time period
	// Once a user is locked out, they will not be able to log back in for this time period
	LockoutIntervalMinutes *float64 `json:"lockoutIntervalMinutes,omitempty" tf:"lockout_interval_minutes,omitempty"`
}

type AccountLockoutParameters struct {

	// Whether account lockout is enabled or not
	// Whether account lockout is enabled or not
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Number of login attempts that will trigger an account lockout for the given user
	// Number of login attempts that will trigger an account lockout for the given user
	// +kubebuilder:validation:Optional
	InvalidLoginsBeforeLockout *float64 `json:"invalidLoginsBeforeLockout" tf:"invalid_logins_before_lockout,omitempty"`

	// Once a user is locked out, they will not be able to log back in for this time period
	// Once a user is locked out, they will not be able to log back in for this time period
	// +kubebuilder:validation:Optional
	LockoutIntervalMinutes *float64 `json:"lockoutIntervalMinutes" tf:"lockout_interval_minutes,omitempty"`
}

type OrgInitParameters struct {

	// (v3.14+) Defines account lockout properties in this organization:
	// Defines account lockout properties in this organization
	AccountLockout []AccountLockoutInitParameters `json:"accountLockout,omitempty" tf:"account_lockout,omitempty"`

	// True if this organization is allowed to share catalogs. Default is true.
	// True if this organization is allowed to share catalogs.
	CanPublishCatalogs *bool `json:"canPublishCatalogs,omitempty" tf:"can_publish_catalogs,omitempty"`

	// True if this organization is allowed to publish external catalogs. Default is false.
	// True if this organization is allowed to publish external catalogs.
	CanPublishExternalCatalogs *bool `json:"canPublishExternalCatalogs,omitempty" tf:"can_publish_external_catalogs,omitempty"`

	// True if this organization is allowed to subscribe to external catalogs. Default is false.
	// True if this organization is allowed to subscribe to external catalogs.
	CanSubscribeExternalCatalogs *bool `json:"canSubscribeExternalCatalogs,omitempty" tf:"can_subscribe_external_catalogs,omitempty"`

	// Specifies this organization's default for virtual machine boot delay after power on. Default is 0.
	// Specifies this organization's default for virtual machine boot delay after power on.
	DelayAfterPowerOnSeconds *float64 `json:"delayAfterPowerOnSeconds,omitempty" tf:"delay_after_power_on_seconds,omitempty"`

	// Pass delete_force=true and delete_recursive=true to remove an organization or VDC and any objects it contains, regardless of their state. Default is false
	// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains, regardless of their state.
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// Pass delete_recursive=true as query parameter to remove an organization or VDC and any objects it contains that are in a state that normally allows removal. Default is false
	// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
	DeployedVMQuota *float64 `json:"deployedVmQuota,omitempty" tf:"deployed_vm_quota,omitempty"`

	// Org description. Default is empty.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Org full name
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	// True if this organization is enabled (allows login and all other operations). Default is true.
	// True if this organization is enabled (allows login and all other operations).
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (Deprecated; v3.6+) Use metadata_entry instead. Key value map of metadata to assign to this organization.
	// Key value map of metadata to assign to this organization. Key and value can be any string.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given Organization
	MetadataEntry []OrgMetadataEntryInitParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// Org name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. (0 = unlimited)
	StoredVMQuota *float64 `json:"storedVmQuota,omitempty" tf:"stored_vm_quota,omitempty"`

	// Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
	// Defines lease parameters for vApps created in this organization
	VappLease []VappLeaseInitParameters `json:"vappLease,omitempty" tf:"vapp_lease,omitempty"`

	// Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
	// Defines lease parameters for vApp templates created in this organization
	VappTemplateLease []VappTemplateLeaseInitParameters `json:"vappTemplateLease,omitempty" tf:"vapp_template_lease,omitempty"`
}

type OrgMetadataEntryInitParameters struct {

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

type OrgMetadataEntryObservation struct {

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

type OrgMetadataEntryParameters struct {

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

type OrgObservation struct {

	// (v3.14+) Defines account lockout properties in this organization:
	// Defines account lockout properties in this organization
	AccountLockout []AccountLockoutObservation `json:"accountLockout,omitempty" tf:"account_lockout,omitempty"`

	// True if this organization is allowed to share catalogs. Default is true.
	// True if this organization is allowed to share catalogs.
	CanPublishCatalogs *bool `json:"canPublishCatalogs,omitempty" tf:"can_publish_catalogs,omitempty"`

	// True if this organization is allowed to publish external catalogs. Default is false.
	// True if this organization is allowed to publish external catalogs.
	CanPublishExternalCatalogs *bool `json:"canPublishExternalCatalogs,omitempty" tf:"can_publish_external_catalogs,omitempty"`

	// True if this organization is allowed to subscribe to external catalogs. Default is false.
	// True if this organization is allowed to subscribe to external catalogs.
	CanSubscribeExternalCatalogs *bool `json:"canSubscribeExternalCatalogs,omitempty" tf:"can_subscribe_external_catalogs,omitempty"`

	// Specifies this organization's default for virtual machine boot delay after power on. Default is 0.
	// Specifies this organization's default for virtual machine boot delay after power on.
	DelayAfterPowerOnSeconds *float64 `json:"delayAfterPowerOnSeconds,omitempty" tf:"delay_after_power_on_seconds,omitempty"`

	// Pass delete_force=true and delete_recursive=true to remove an organization or VDC and any objects it contains, regardless of their state. Default is false
	// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains, regardless of their state.
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// Pass delete_recursive=true as query parameter to remove an organization or VDC and any objects it contains that are in a state that normally allows removal. Default is false
	// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
	DeployedVMQuota *float64 `json:"deployedVmQuota,omitempty" tf:"deployed_vm_quota,omitempty"`

	// Org description. Default is empty.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Org full name
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// True if this organization is enabled (allows login and all other operations). Default is true.
	// True if this organization is enabled (allows login and all other operations).
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (v3.11+) List of catalogs (sorted alphabetically), owned or shared, available to this organization.
	// List of catalogs, owned or shared, available to this organization
	// +listType=set
	ListOfCatalogs []*string `json:"listOfCatalogs,omitempty" tf:"list_of_catalogs,omitempty"`

	// (v3.11+) List of VDCs (sorted alphabetically), owned or shared, available to this organization.
	// List of VDCs, owned or shared, available to this organization
	// +listType=set
	ListOfVdcs []*string `json:"listOfVdcs,omitempty" tf:"list_of_vdcs,omitempty"`

	// (Deprecated; v3.6+) Use metadata_entry instead. Key value map of metadata to assign to this organization.
	// Key value map of metadata to assign to this organization. Key and value can be any string.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given Organization
	MetadataEntry []OrgMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// Org name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (v3.11+) Number of catalogs owned or shared, available to this organization.
	// Number of catalogs, owned or shared, available to this organization
	NumberOfCatalogs *float64 `json:"numberOfCatalogs,omitempty" tf:"number_of_catalogs,omitempty"`

	// (v3.11+) Number of VDCs owned or shared, available to this organization.
	// Number of VDCs, owned or shared, available to this organization
	NumberOfVdcs *float64 `json:"numberOfVdcs,omitempty" tf:"number_of_vdcs,omitempty"`

	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. (0 = unlimited)
	StoredVMQuota *float64 `json:"storedVmQuota,omitempty" tf:"stored_vm_quota,omitempty"`

	// Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
	// Defines lease parameters for vApps created in this organization
	VappLease []VappLeaseObservation `json:"vappLease,omitempty" tf:"vapp_lease,omitempty"`

	// Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
	// Defines lease parameters for vApp templates created in this organization
	VappTemplateLease []VappTemplateLeaseObservation `json:"vappTemplateLease,omitempty" tf:"vapp_template_lease,omitempty"`
}

type OrgParameters struct {

	// (v3.14+) Defines account lockout properties in this organization:
	// Defines account lockout properties in this organization
	// +kubebuilder:validation:Optional
	AccountLockout []AccountLockoutParameters `json:"accountLockout,omitempty" tf:"account_lockout,omitempty"`

	// True if this organization is allowed to share catalogs. Default is true.
	// True if this organization is allowed to share catalogs.
	// +kubebuilder:validation:Optional
	CanPublishCatalogs *bool `json:"canPublishCatalogs,omitempty" tf:"can_publish_catalogs,omitempty"`

	// True if this organization is allowed to publish external catalogs. Default is false.
	// True if this organization is allowed to publish external catalogs.
	// +kubebuilder:validation:Optional
	CanPublishExternalCatalogs *bool `json:"canPublishExternalCatalogs,omitempty" tf:"can_publish_external_catalogs,omitempty"`

	// True if this organization is allowed to subscribe to external catalogs. Default is false.
	// True if this organization is allowed to subscribe to external catalogs.
	// +kubebuilder:validation:Optional
	CanSubscribeExternalCatalogs *bool `json:"canSubscribeExternalCatalogs,omitempty" tf:"can_subscribe_external_catalogs,omitempty"`

	// Specifies this organization's default for virtual machine boot delay after power on. Default is 0.
	// Specifies this organization's default for virtual machine boot delay after power on.
	// +kubebuilder:validation:Optional
	DelayAfterPowerOnSeconds *float64 `json:"delayAfterPowerOnSeconds,omitempty" tf:"delay_after_power_on_seconds,omitempty"`

	// Pass delete_force=true and delete_recursive=true to remove an organization or VDC and any objects it contains, regardless of their state. Default is false
	// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains, regardless of their state.
	// +kubebuilder:validation:Optional
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// Pass delete_recursive=true as query parameter to remove an organization or VDC and any objects it contains that are in a state that normally allows removal. Default is false
	// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that normally allows removal.
	// +kubebuilder:validation:Optional
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
	// +kubebuilder:validation:Optional
	DeployedVMQuota *float64 `json:"deployedVmQuota,omitempty" tf:"deployed_vm_quota,omitempty"`

	// Org description. Default is empty.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Org full name
	// +kubebuilder:validation:Optional
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	// True if this organization is enabled (allows login and all other operations). Default is true.
	// True if this organization is enabled (allows login and all other operations).
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// (Deprecated; v3.6+) Use metadata_entry instead. Key value map of metadata to assign to this organization.
	// Key value map of metadata to assign to this organization. Key and value can be any string.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A set of metadata entries to assign. See Metadata section for details.
	// Metadata entries for the given Organization
	// +kubebuilder:validation:Optional
	MetadataEntry []OrgMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// Org name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. Default is unlimited (0)
	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. (0 = unlimited)
	// +kubebuilder:validation:Optional
	StoredVMQuota *float64 `json:"storedVmQuota,omitempty" tf:"stored_vm_quota,omitempty"`

	// Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
	// Defines lease parameters for vApps created in this organization
	// +kubebuilder:validation:Optional
	VappLease []VappLeaseParameters `json:"vappLease,omitempty" tf:"vapp_lease,omitempty"`

	// Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
	// Defines lease parameters for vApp templates created in this organization
	// +kubebuilder:validation:Optional
	VappTemplateLease []VappTemplateLeaseParameters `json:"vappTemplateLease,omitempty" tf:"vapp_template_lease,omitempty"`
}

type VappLeaseInitParameters struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration,omitempty" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires. Values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 604800 (7 days) but may vary depending on vCD version
	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires
	MaximumRuntimeLeaseInSec *float64 `json:"maximumRuntimeLeaseInSec,omitempty" tf:"maximum_runtime_lease_in_sec,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec,omitempty" tf:"maximum_storage_lease_in_sec,omitempty"`

	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
	// Note: Default when the whole vapp_lease block is omitted is false
	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires
	PowerOffOnRuntimeLeaseExpiration *bool `json:"powerOffOnRuntimeLeaseExpiration,omitempty" tf:"power_off_on_runtime_lease_expiration,omitempty"`
}

type VappLeaseObservation struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration,omitempty" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires. Values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 604800 (7 days) but may vary depending on vCD version
	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires
	MaximumRuntimeLeaseInSec *float64 `json:"maximumRuntimeLeaseInSec,omitempty" tf:"maximum_runtime_lease_in_sec,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec,omitempty" tf:"maximum_storage_lease_in_sec,omitempty"`

	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
	// Note: Default when the whole vapp_lease block is omitted is false
	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires
	PowerOffOnRuntimeLeaseExpiration *bool `json:"powerOffOnRuntimeLeaseExpiration,omitempty" tf:"power_off_on_runtime_lease_expiration,omitempty"`
}

type VappLeaseParameters struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// +kubebuilder:validation:Optional
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires. Values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 604800 (7 days) but may vary depending on vCD version
	// How long vApps can run before they are automatically stopped (in seconds). 0 means never expires
	// +kubebuilder:validation:Optional
	MaximumRuntimeLeaseInSec *float64 `json:"maximumRuntimeLeaseInSec" tf:"maximum_runtime_lease_in_sec,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires
	// +kubebuilder:validation:Optional
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec" tf:"maximum_storage_lease_in_sec,omitempty"`

	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
	// Note: Default when the whole vapp_lease block is omitted is false
	// When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires
	// +kubebuilder:validation:Optional
	PowerOffOnRuntimeLeaseExpiration *bool `json:"powerOffOnRuntimeLeaseExpiration" tf:"power_off_on_runtime_lease_expiration,omitempty"`
}

type VappTemplateLeaseInitParameters struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration,omitempty" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long vApp templates are available before being automatically cleaned up (in seconds). 0 means never expires
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec,omitempty" tf:"maximum_storage_lease_in_sec,omitempty"`
}

type VappTemplateLeaseObservation struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration,omitempty" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long vApp templates are available before being automatically cleaned up (in seconds). 0 means never expires
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec,omitempty" tf:"maximum_storage_lease_in_sec,omitempty"`
}

type VappTemplateLeaseParameters struct {

	// If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
	// Note: Default when the whole vapp_lease block is omitted is false
	// If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
	// +kubebuilder:validation:Optional
	DeleteOnStorageLeaseExpiration *bool `json:"deleteOnStorageLeaseExpiration" tf:"delete_on_storage_lease_expiration,omitempty"`

	// How long stopped vApps are available before being automatically cleaned up (in seconds). 0 means never expires. Regular values accepted from 3600+
	// Note: Default when the whole vapp_lease block is omitted is 2592000 (30 days) but may vary depending on vCD version
	// How long vApp templates are available before being automatically cleaned up (in seconds). 0 means never expires
	// +kubebuilder:validation:Optional
	MaximumStorageLeaseInSec *float64 `json:"maximumStorageLeaseInSec" tf:"maximum_storage_lease_in_sec,omitempty"`
}

// OrgSpec defines the desired state of Org
type OrgSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrgParameters `json:"forProvider"`
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
	InitProvider OrgInitParameters `json:"initProvider,omitempty"`
}

// OrgStatus defines the observed state of Org.
type OrgStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrgObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Org is the Schema for the Orgs API. Provides a VMware Cloud Director Organization resource. This can be used to create  delete, and update an organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type Org struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fullName) || (has(self.initProvider) && has(self.initProvider.fullName))",message="spec.forProvider.fullName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OrgSpec   `json:"spec"`
	Status OrgStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrgList contains a list of Orgs
type OrgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Org `json:"items"`
}

// Repository type metadata.
var (
	Org_Kind             = "Org"
	Org_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Org_Kind}.String()
	Org_KindAPIVersion   = Org_Kind + "." + CRDGroupVersion.String()
	Org_GroupVersionKind = CRDGroupVersion.WithKind(Org_Kind)
)

func init() {
	SchemeBuilder.Register(&Org{}, &OrgList{})
}
