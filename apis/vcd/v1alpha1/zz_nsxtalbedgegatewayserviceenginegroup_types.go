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

type NsxtAlbEdgegatewayServiceEngineGroupInitParameters struct {

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source.
	// Edge Gateway ID in which ALB Service Engine Group should be located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Maximum amount of Virtual Services to run on this Service Engine Group. Only for
	// Shared Service Engine Groups.
	// Maximum number of virtual services to be used in this Service Engine Group
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved Virtual Services for this Edge Gateway. Only for Shared
	// Service Engine Groups.
	// Number of reserved virtual services for this Service Engine Group
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// An ID of NSX-T Service Engine Group. Can be looked up using
	// vcd_nsxt_alb_service_engine_group data
	// source.
	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtAlbEdgegatewayServiceEngineGroupObservation struct {

	// Number of deployed Virtual Services on this Service Engine Group.
	// Number of deployed virtual services for this Service Engine Group
	DeployedVirtualServices *float64 `json:"deployedVirtualServices,omitempty" tf:"deployed_virtual_services,omitempty"`

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source.
	// Edge Gateway ID in which ALB Service Engine Group should be located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum amount of Virtual Services to run on this Service Engine Group. Only for
	// Shared Service Engine Groups.
	// Maximum number of virtual services to be used in this Service Engine Group
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved Virtual Services for this Edge Gateway. Only for Shared
	// Service Engine Groups.
	// Number of reserved virtual services for this Service Engine Group
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// An ID of NSX-T Service Engine Group. Can be looked up using
	// vcd_nsxt_alb_service_engine_group data
	// source.
	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// Service Engine Group Name which is attached to NSX-T Edge Gateway
	ServiceEngineGroupName *string `json:"serviceEngineGroupName,omitempty" tf:"service_engine_group_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtAlbEdgegatewayServiceEngineGroupParameters struct {

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source.
	// Edge Gateway ID in which ALB Service Engine Group should be located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Maximum amount of Virtual Services to run on this Service Engine Group. Only for
	// Shared Service Engine Groups.
	// Maximum number of virtual services to be used in this Service Engine Group
	// +kubebuilder:validation:Optional
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved Virtual Services for this Edge Gateway. Only for Shared
	// Service Engine Groups.
	// Number of reserved virtual services for this Service Engine Group
	// +kubebuilder:validation:Optional
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// An ID of NSX-T Service Engine Group. Can be looked up using
	// vcd_nsxt_alb_service_engine_group data
	// source.
	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	// +kubebuilder:validation:Optional
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxtAlbEdgegatewayServiceEngineGroupSpec defines the desired state of NsxtAlbEdgegatewayServiceEngineGroup
type NsxtAlbEdgegatewayServiceEngineGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtAlbEdgegatewayServiceEngineGroupParameters `json:"forProvider"`
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
	InitProvider NsxtAlbEdgegatewayServiceEngineGroupInitParameters `json:"initProvider,omitempty"`
}

// NsxtAlbEdgegatewayServiceEngineGroupStatus defines the observed state of NsxtAlbEdgegatewayServiceEngineGroup.
type NsxtAlbEdgegatewayServiceEngineGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtAlbEdgegatewayServiceEngineGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxtAlbEdgegatewayServiceEngineGroup is the Schema for the NsxtAlbEdgegatewayServiceEngineGroups API. Provides a resource to manage ALB Service Engine Group assignment to Edge Gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtAlbEdgegatewayServiceEngineGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGatewayId) || (has(self.initProvider) && has(self.initProvider.edgeGatewayId))",message="spec.forProvider.edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceEngineGroupId) || (has(self.initProvider) && has(self.initProvider.serviceEngineGroupId))",message="spec.forProvider.serviceEngineGroupId is a required parameter"
	Spec   NsxtAlbEdgegatewayServiceEngineGroupSpec   `json:"spec"`
	Status NsxtAlbEdgegatewayServiceEngineGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbEdgegatewayServiceEngineGroupList contains a list of NsxtAlbEdgegatewayServiceEngineGroups
type NsxtAlbEdgegatewayServiceEngineGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtAlbEdgegatewayServiceEngineGroup `json:"items"`
}

// Repository type metadata.
var (
	NsxtAlbEdgegatewayServiceEngineGroup_Kind             = "NsxtAlbEdgegatewayServiceEngineGroup"
	NsxtAlbEdgegatewayServiceEngineGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtAlbEdgegatewayServiceEngineGroup_Kind}.String()
	NsxtAlbEdgegatewayServiceEngineGroup_KindAPIVersion   = NsxtAlbEdgegatewayServiceEngineGroup_Kind + "." + CRDGroupVersion.String()
	NsxtAlbEdgegatewayServiceEngineGroup_GroupVersionKind = CRDGroupVersion.WithKind(NsxtAlbEdgegatewayServiceEngineGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtAlbEdgegatewayServiceEngineGroup{}, &NsxtAlbEdgegatewayServiceEngineGroupList{})
}
