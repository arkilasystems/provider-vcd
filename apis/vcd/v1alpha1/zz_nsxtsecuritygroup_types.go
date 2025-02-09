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

type NsxtSecurityGroupInitParameters struct {

	// An optional description of the Security Group
	// Security Group description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge Gateway ID in which security group is located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// A set of Org Network IDs
	// Set of Org VDC network IDs attached to this security group
	// +listType=set
	MemberOrgNetworkIds []*string `json:"memberOrgNetworkIds,omitempty" tf:"member_org_network_ids,omitempty"`

	// A unique name for Security Group
	// Security Group name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// (Deprecated; Optional) The name of VDC to use, optional if defined at provider level. Deprecated
	// in favor of edge_gateway_id field.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtSecurityGroupMemberVmsInitParameters struct {
}

type NsxtSecurityGroupMemberVmsObservation struct {

	// Member VM ID
	VMID *string `json:"vmId,omitempty" tf:"vm_id,omitempty"`

	// Member VM name
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// Parent vApp ID for member VM (empty for standalone VMs)
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// Parent vApp Name for member VM (empty for standalone VMs)
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`
}

type NsxtSecurityGroupMemberVmsParameters struct {
}

type NsxtSecurityGroupObservation struct {

	// An optional description of the Security Group
	// Security Group description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge Gateway ID in which security group is located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of Org Network IDs
	// Set of Org VDC network IDs attached to this security group
	// +listType=set
	MemberOrgNetworkIds []*string `json:"memberOrgNetworkIds,omitempty" tf:"member_org_network_ids,omitempty"`

	// A set of member VMs (if exist). see Member VMs below for details.
	// Set of VM IDs
	MemberVms []NsxtSecurityGroupMemberVmsObservation `json:"memberVms,omitempty" tf:"member_vms,omitempty"`

	// A unique name for Security Group
	// Security Group name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// ID of VDC or VDC Group
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// (Deprecated; Optional) The name of VDC to use, optional if defined at provider level. Deprecated
	// in favor of edge_gateway_id field.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtSecurityGroupParameters struct {

	// An optional description of the Security Group
	// Security Group description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge Gateway ID in which security group is located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// A set of Org Network IDs
	// Set of Org VDC network IDs attached to this security group
	// +kubebuilder:validation:Optional
	// +listType=set
	MemberOrgNetworkIds []*string `json:"memberOrgNetworkIds,omitempty" tf:"member_org_network_ids,omitempty"`

	// A unique name for Security Group
	// Security Group name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// (Deprecated; Optional) The name of VDC to use, optional if defined at provider level. Deprecated
	// in favor of edge_gateway_id field.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxtSecurityGroupSpec defines the desired state of NsxtSecurityGroup
type NsxtSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtSecurityGroupParameters `json:"forProvider"`
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
	InitProvider NsxtSecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// NsxtSecurityGroupStatus defines the observed state of NsxtSecurityGroup.
type NsxtSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxtSecurityGroup is the Schema for the NsxtSecurityGroups API. Provides a resource to manage NSX-T Security Group. Security Groups are groups of data center group networks to which distributed firewall rules apply. Grouping networks helps you to reduce the total number of distributed firewall rules to be created.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGatewayId) || (has(self.initProvider) && has(self.initProvider.edgeGatewayId))",message="spec.forProvider.edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NsxtSecurityGroupSpec   `json:"spec"`
	Status NsxtSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtSecurityGroupList contains a list of NsxtSecurityGroups
type NsxtSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	NsxtSecurityGroup_Kind             = "NsxtSecurityGroup"
	NsxtSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtSecurityGroup_Kind}.String()
	NsxtSecurityGroup_KindAPIVersion   = NsxtSecurityGroup_Kind + "." + CRDGroupVersion.String()
	NsxtSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(NsxtSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtSecurityGroup{}, &NsxtSecurityGroupList{})
}
