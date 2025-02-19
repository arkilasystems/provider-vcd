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

type SubscribedCatalogInitParameters struct {

	// When true, the subscribed catalog will attempt canceling failed tasks.
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks *bool `json:"cancelFailedTasks,omitempty" tf:"cancel_failed_tasks,omitempty"`

	// When destroying use delete_force=true with delete_recursive=true to remove a catalog and any objects it contains, regardless of their state.
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// When destroying use delete_recursive=true to remove the catalog and any objects it contains that are in a state that normally allows removal.
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be false if the user configured in the provider is the System administrator.
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	MakeLocalCopy *bool `json:"makeLocalCopy,omitempty" tf:"make_local_copy,omitempty"`

	// Catalog name
	// The name of the catalog
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level.
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Allows to set specific storage profile to be used for catalog.
	// Optional storage profile ID
	StorageProfileID *string `json:"storageProfileId,omitempty" tf:"storage_profile_id,omitempty"`

	// if true, saves the list of tasks to a file for later update.
	// If true, saves list of tasks to file for later update
	StoreTasks *bool `json:"storeTasks,omitempty" tf:"store_tasks,omitempty"`

	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPasswordSecretRef *v1.SecretKeySelector `json:"subscriptionPasswordSecretRef,omitempty" tf:"-"`

	// The URL to subscribe to the external catalog.
	// The URL to subscribe to the external catalog.
	SubscriptionURL *string `json:"subscriptionUrl,omitempty" tf:"subscription_url,omitempty"`

	// If true, synchronise this catalog and all items.
	// If true, synchronise this catalog and all items
	SyncAll *bool `json:"syncAll,omitempty" tf:"sync_all,omitempty"`

	// If true, synchronise all media items. Not to be used when sync_all is set.
	// If true, synchronises all media items
	SyncAllMediaItems *bool `json:"syncAllMediaItems,omitempty" tf:"sync_all_media_items,omitempty"`

	// If true, synchronise all vApp templates. Not to be used when sync_all is set.
	// If true, synchronises all vApp templates
	SyncAllVappTemplates *bool `json:"syncAllVappTemplates,omitempty" tf:"sync_all_vapp_templates,omitempty"`

	// If true, synchronise this catalog. Not to be used when sync_all is set. This operation fetches the list of items. If make_local_copy is set, it also synchronises all the items.
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also fetches the items data.
	SyncCatalog *bool `json:"syncCatalog,omitempty" tf:"sync_catalog,omitempty"`

	// Synchronise a list of media items. Not to be used when sync_all or sync_all_media_items are set.
	// Synchronises media items from this list of names.
	SyncMediaItems []*string `json:"syncMediaItems,omitempty" tf:"sync_media_items,omitempty"`

	// Boolean value that shows if sync should be performed on every refresh.
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `json:"syncOnRefresh,omitempty" tf:"sync_on_refresh,omitempty"`

	// Synchronise a list of vApp templates. Not to be used when sync_all or sync_all_vapp_templates are set.
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates []*string `json:"syncVappTemplates,omitempty" tf:"sync_vapp_templates,omitempty"`
}

type SubscribedCatalogObservation struct {

	// When true, the subscribed catalog will attempt canceling failed tasks.
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks *bool `json:"cancelFailedTasks,omitempty" tf:"cancel_failed_tasks,omitempty"`

	// Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
	// Catalog version number. Inherited from publishing catalog and updated on sync.
	CatalogVersion *float64 `json:"catalogVersion,omitempty" tf:"catalog_version,omitempty"`

	// Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
	// Time stamp of when the catalog was created
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// When destroying use delete_force=true with delete_recursive=true to remove a catalog and any objects it contains, regardless of their state.
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// When destroying use delete_recursive=true to remove the catalog and any objects it contains that are in a state that normally allows removal.
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// Description of catalog. This is inherited from the publishing catalog and updated on sync.
	// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
	// List of failed synchronization tasks
	FailedTasks []*string `json:"failedTasks,omitempty" tf:"failed_tasks,omitempty"`

	// the catalog's Hyper reference.
	// Catalog HREF
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (v3.8.1+) Indicates if this catalog was created in the current organization.
	// True if this catalog belongs to the current organization.
	IsLocal *bool `json:"isLocal,omitempty" tf:"is_local,omitempty"`

	// Indicates if this catalog is available for subscription. (Always false)
	// True if this catalog is published. (Always false)
	IsPublished *bool `json:"isPublished,omitempty" tf:"is_published,omitempty"`

	// Indicates if the catalog is shared.
	// True if this catalog is shared.
	IsShared *bool `json:"isShared,omitempty" tf:"is_shared,omitempty"`

	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be false if the user configured in the provider is the System administrator.
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	MakeLocalCopy *bool `json:"makeLocalCopy,omitempty" tf:"make_local_copy,omitempty"`

	// List of media item names in this catalog, in alphabetical order.
	// List of Media items in this catalog
	MediaItemList []*string `json:"mediaItemList,omitempty" tf:"media_item_list,omitempty"`

	// (Available until VCD 10.5) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
	// Key and value pairs for catalog metadata. Inherited from publishing catalog
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Catalog name
	// The name of the catalog
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of media items available in this catalog.
	// Number of Media items this catalog contains.
	NumberOfMedia *float64 `json:"numberOfMedia,omitempty" tf:"number_of_media,omitempty"`

	// Number of vApp templates available in this catalog.
	// Number of vApp templates this catalog contains.
	NumberOfVappTemplates *float64 `json:"numberOfVappTemplates,omitempty" tf:"number_of_vapp_templates,omitempty"`

	// The name of organization to use, optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level.
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Owner of the catalog.
	// Owner name from the catalog.
	OwnerName *string `json:"ownerName,omitempty" tf:"owner_name,omitempty"`

	// Shows if the catalog is published, if it is a subscription from another one or none of those. (Always SUBSCRIBED)
	// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always SUBSCRIBED)
	PublishSubscriptionType *string `json:"publishSubscriptionType,omitempty" tf:"publish_subscription_type,omitempty"`

	// List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
	// List of running synchronization tasks
	RunningTasks []*string `json:"runningTasks,omitempty" tf:"running_tasks,omitempty"`

	// Allows to set specific storage profile to be used for catalog.
	// Optional storage profile ID
	StorageProfileID *string `json:"storageProfileId,omitempty" tf:"storage_profile_id,omitempty"`

	// if true, saves the list of tasks to a file for later update.
	// If true, saves list of tasks to file for later update
	StoreTasks *bool `json:"storeTasks,omitempty" tf:"store_tasks,omitempty"`

	// The URL to subscribe to the external catalog.
	// The URL to subscribe to the external catalog.
	SubscriptionURL *string `json:"subscriptionUrl,omitempty" tf:"subscription_url,omitempty"`

	// If true, synchronise this catalog and all items.
	// If true, synchronise this catalog and all items
	SyncAll *bool `json:"syncAll,omitempty" tf:"sync_all,omitempty"`

	// If true, synchronise all media items. Not to be used when sync_all is set.
	// If true, synchronises all media items
	SyncAllMediaItems *bool `json:"syncAllMediaItems,omitempty" tf:"sync_all_media_items,omitempty"`

	// If true, synchronise all vApp templates. Not to be used when sync_all is set.
	// If true, synchronises all vApp templates
	SyncAllVappTemplates *bool `json:"syncAllVappTemplates,omitempty" tf:"sync_all_vapp_templates,omitempty"`

	// If true, synchronise this catalog. Not to be used when sync_all is set. This operation fetches the list of items. If make_local_copy is set, it also synchronises all the items.
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also fetches the items data.
	SyncCatalog *bool `json:"syncCatalog,omitempty" tf:"sync_catalog,omitempty"`

	// Synchronise a list of media items. Not to be used when sync_all or sync_all_media_items are set.
	// Synchronises media items from this list of names.
	SyncMediaItems []*string `json:"syncMediaItems,omitempty" tf:"sync_media_items,omitempty"`

	// Boolean value that shows if sync should be performed on every refresh.
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `json:"syncOnRefresh,omitempty" tf:"sync_on_refresh,omitempty"`

	// Synchronise a list of vApp templates. Not to be used when sync_all or sync_all_vapp_templates are set.
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates []*string `json:"syncVappTemplates,omitempty" tf:"sync_vapp_templates,omitempty"`

	// Where the running tasks IDs have been stored. Only if store_tasks is set.
	// Where the running tasks IDs have been stored. Only if `store_tasks` is set
	TasksFileName *string `json:"tasksFileName,omitempty" tf:"tasks_file_name,omitempty"`

	// List of vApp template names in this catalog, in alphabetical order.
	// List of catalog items in this catalog
	VappTemplateList []*string `json:"vappTemplateList,omitempty" tf:"vapp_template_list,omitempty"`
}

type SubscribedCatalogParameters struct {

	// When true, the subscribed catalog will attempt canceling failed tasks.
	// When true, the subscribed catalog will attempt canceling failed tasks
	// +kubebuilder:validation:Optional
	CancelFailedTasks *bool `json:"cancelFailedTasks,omitempty" tf:"cancel_failed_tasks,omitempty"`

	// When destroying use delete_force=true with delete_recursive=true to remove a catalog and any objects it contains, regardless of their state.
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains, regardless of their state.
	// +kubebuilder:validation:Optional
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// When destroying use delete_recursive=true to remove the catalog and any objects it contains that are in a state that normally allows removal.
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal.
	// +kubebuilder:validation:Optional
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be false if the user configured in the provider is the System administrator.
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local copy of catalog items unless a sync operation is performed.
	// +kubebuilder:validation:Optional
	MakeLocalCopy *bool `json:"makeLocalCopy,omitempty" tf:"make_local_copy,omitempty"`

	// Catalog name
	// The name of the catalog
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level.
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Allows to set specific storage profile to be used for catalog.
	// Optional storage profile ID
	// +kubebuilder:validation:Optional
	StorageProfileID *string `json:"storageProfileId,omitempty" tf:"storage_profile_id,omitempty"`

	// if true, saves the list of tasks to a file for later update.
	// If true, saves list of tasks to file for later update
	// +kubebuilder:validation:Optional
	StoreTasks *bool `json:"storeTasks,omitempty" tf:"store_tasks,omitempty"`

	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	// +kubebuilder:validation:Optional
	SubscriptionPasswordSecretRef *v1.SecretKeySelector `json:"subscriptionPasswordSecretRef,omitempty" tf:"-"`

	// The URL to subscribe to the external catalog.
	// The URL to subscribe to the external catalog.
	// +kubebuilder:validation:Optional
	SubscriptionURL *string `json:"subscriptionUrl,omitempty" tf:"subscription_url,omitempty"`

	// If true, synchronise this catalog and all items.
	// If true, synchronise this catalog and all items
	// +kubebuilder:validation:Optional
	SyncAll *bool `json:"syncAll,omitempty" tf:"sync_all,omitempty"`

	// If true, synchronise all media items. Not to be used when sync_all is set.
	// If true, synchronises all media items
	// +kubebuilder:validation:Optional
	SyncAllMediaItems *bool `json:"syncAllMediaItems,omitempty" tf:"sync_all_media_items,omitempty"`

	// If true, synchronise all vApp templates. Not to be used when sync_all is set.
	// If true, synchronises all vApp templates
	// +kubebuilder:validation:Optional
	SyncAllVappTemplates *bool `json:"syncAllVappTemplates,omitempty" tf:"sync_all_vapp_templates,omitempty"`

	// If true, synchronise this catalog. Not to be used when sync_all is set. This operation fetches the list of items. If make_local_copy is set, it also synchronises all the items.
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also fetches the items data.
	// +kubebuilder:validation:Optional
	SyncCatalog *bool `json:"syncCatalog,omitempty" tf:"sync_catalog,omitempty"`

	// Synchronise a list of media items. Not to be used when sync_all or sync_all_media_items are set.
	// Synchronises media items from this list of names.
	// +kubebuilder:validation:Optional
	SyncMediaItems []*string `json:"syncMediaItems,omitempty" tf:"sync_media_items,omitempty"`

	// Boolean value that shows if sync should be performed on every refresh.
	// Boolean value that shows if sync should be performed on every refresh.
	// +kubebuilder:validation:Optional
	SyncOnRefresh *bool `json:"syncOnRefresh,omitempty" tf:"sync_on_refresh,omitempty"`

	// Synchronise a list of vApp templates. Not to be used when sync_all or sync_all_vapp_templates are set.
	// Synchronises vApp templates from this list of names.
	// +kubebuilder:validation:Optional
	SyncVappTemplates []*string `json:"syncVappTemplates,omitempty" tf:"sync_vapp_templates,omitempty"`
}

// SubscribedCatalogSpec defines the desired state of SubscribedCatalog
type SubscribedCatalogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscribedCatalogParameters `json:"forProvider"`
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
	InitProvider SubscribedCatalogInitParameters `json:"initProvider,omitempty"`
}

// SubscribedCatalogStatus defines the observed state of SubscribedCatalog.
type SubscribedCatalogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscribedCatalogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubscribedCatalog is the Schema for the SubscribedCatalogs API. Provides a VMware Cloud Director subscribed catalog resource. This can be used to create, update, and delete a subscribed catalog.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type SubscribedCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionUrl) || (has(self.initProvider) && has(self.initProvider.subscriptionUrl))",message="spec.forProvider.subscriptionUrl is a required parameter"
	Spec   SubscribedCatalogSpec   `json:"spec"`
	Status SubscribedCatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscribedCatalogList contains a list of SubscribedCatalogs
type SubscribedCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscribedCatalog `json:"items"`
}

// Repository type metadata.
var (
	SubscribedCatalog_Kind             = "SubscribedCatalog"
	SubscribedCatalog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscribedCatalog_Kind}.String()
	SubscribedCatalog_KindAPIVersion   = SubscribedCatalog_Kind + "." + CRDGroupVersion.String()
	SubscribedCatalog_GroupVersionKind = CRDGroupVersion.WithKind(SubscribedCatalog_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscribedCatalog{}, &SubscribedCatalogList{})
}
