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

type IpSpaceUplinkInitParameters struct {

	// A set of Tier-0 Router Interfaces to associate with this uplink
	// +listType=set
	AssociatedInterfaceIds []*string `json:"associatedInterfaceIds,omitempty" tf:"associated_interface_ids,omitempty"`

	// An optional description for IP Space Uplink
	// IP Space Uplink description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External Network ID For IP Space Uplink configuration
	// External Network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// IP Space ID configuration
	// IP Space ID
	IPSpaceID *string `json:"ipSpaceId,omitempty" tf:"ip_space_id,omitempty"`

	// A tenant facing name for IP Space Uplink
	// Tenant facing name for IP Space Uplink
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IpSpaceUplinkObservation struct {

	// A set of Tier-0 Router Interfaces to associate with this uplink
	// +listType=set
	AssociatedInterfaceIds []*string `json:"associatedInterfaceIds,omitempty" tf:"associated_interface_ids,omitempty"`

	// An optional description for IP Space Uplink
	// IP Space Uplink description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External Network ID For IP Space Uplink configuration
	// External Network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP Space ID configuration
	// IP Space ID
	IPSpaceID *string `json:"ipSpaceId,omitempty" tf:"ip_space_id,omitempty"`

	// Backing IP Space type
	// IP Space Type
	IPSpaceType *string `json:"ipSpaceType,omitempty" tf:"ip_space_type,omitempty"`

	// A tenant facing name for IP Space Uplink
	// Tenant facing name for IP Space Uplink
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Status of IP Space Uplink
	// IP Space Status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type IpSpaceUplinkParameters struct {

	// A set of Tier-0 Router Interfaces to associate with this uplink
	// +kubebuilder:validation:Optional
	// +listType=set
	AssociatedInterfaceIds []*string `json:"associatedInterfaceIds,omitempty" tf:"associated_interface_ids,omitempty"`

	// An optional description for IP Space Uplink
	// IP Space Uplink description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External Network ID For IP Space Uplink configuration
	// External Network ID
	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// IP Space ID configuration
	// IP Space ID
	// +kubebuilder:validation:Optional
	IPSpaceID *string `json:"ipSpaceId,omitempty" tf:"ip_space_id,omitempty"`

	// A tenant facing name for IP Space Uplink
	// Tenant facing name for IP Space Uplink
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// IpSpaceUplinkSpec defines the desired state of IpSpaceUplink
type IpSpaceUplinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpSpaceUplinkParameters `json:"forProvider"`
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
	InitProvider IpSpaceUplinkInitParameters `json:"initProvider,omitempty"`
}

// IpSpaceUplinkStatus defines the observed state of IpSpaceUplink.
type IpSpaceUplinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpSpaceUplinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IpSpaceUplink is the Schema for the IpSpaceUplinks API. Provides a resource to manage IP Space Uplinks in External Networks (Provider Gateways).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type IpSpaceUplink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.externalNetworkId) || (has(self.initProvider) && has(self.initProvider.externalNetworkId))",message="spec.forProvider.externalNetworkId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipSpaceId) || (has(self.initProvider) && has(self.initProvider.ipSpaceId))",message="spec.forProvider.ipSpaceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IpSpaceUplinkSpec   `json:"spec"`
	Status IpSpaceUplinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpSpaceUplinkList contains a list of IpSpaceUplinks
type IpSpaceUplinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpSpaceUplink `json:"items"`
}

// Repository type metadata.
var (
	IpSpaceUplink_Kind             = "IpSpaceUplink"
	IpSpaceUplink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpSpaceUplink_Kind}.String()
	IpSpaceUplink_KindAPIVersion   = IpSpaceUplink_Kind + "." + CRDGroupVersion.String()
	IpSpaceUplink_GroupVersionKind = CRDGroupVersion.WithKind(IpSpaceUplink_Kind)
)

func init() {
	SchemeBuilder.Register(&IpSpaceUplink{}, &IpSpaceUplinkList{})
}
