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

type ExternalNetworkV2IPScopeInitParameters struct {

	// A FQDN for the virtual machines on this network. Provider version
	// v3.9+ also supports NSX-T
	// DNS suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Primary DNS server. Provider version v3.9+ also supports NSX-T
	// Primary DNS server
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Secondary DNS server. Provider version v3.9+ also supports NSX-T
	// Secondary DNS server
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// Default is true.
	// If subnet is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Gateway of the network.
	// Gateway of the network
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Network prefix.
	// Network mask
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// IP ranges used for static pool allocation in the network.  See IP Pool below for details.
	// IP ranges used for static pool allocation in the network
	StaticIPPool []IPScopeStaticIPPoolInitParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`
}

type ExternalNetworkV2IPScopeObservation struct {

	// A FQDN for the virtual machines on this network. Provider version
	// v3.9+ also supports NSX-T
	// DNS suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Primary DNS server. Provider version v3.9+ also supports NSX-T
	// Primary DNS server
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Secondary DNS server. Provider version v3.9+ also supports NSX-T
	// Secondary DNS server
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// Default is true.
	// If subnet is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Gateway of the network.
	// Gateway of the network
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Network prefix.
	// Network mask
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// IP ranges used for static pool allocation in the network.  See IP Pool below for details.
	// IP ranges used for static pool allocation in the network
	StaticIPPool []IPScopeStaticIPPoolObservation `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`
}

type ExternalNetworkV2IPScopeParameters struct {

	// A FQDN for the virtual machines on this network. Provider version
	// v3.9+ also supports NSX-T
	// DNS suffix
	// +kubebuilder:validation:Optional
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Primary DNS server. Provider version v3.9+ also supports NSX-T
	// Primary DNS server
	// +kubebuilder:validation:Optional
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Secondary DNS server. Provider version v3.9+ also supports NSX-T
	// Secondary DNS server
	// +kubebuilder:validation:Optional
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// Default is true.
	// If subnet is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Gateway of the network.
	// Gateway of the network
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// Network prefix.
	// Network mask
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength" tf:"prefix_length,omitempty"`

	// IP ranges used for static pool allocation in the network.  See IP Pool below for details.
	// IP ranges used for static pool allocation in the network
	// +kubebuilder:validation:Optional
	StaticIPPool []IPScopeStaticIPPoolParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`
}

type ExternalNetworkV2InitParameters struct {

	// An Org ID that this network should be
	// dedicated to. Only applicable when use_ip_spaces=true
	// Dedicate this External Network to an Org ID (only with IP Spaces, VCD 10.4.1+)
	DedicatedOrgID *string `json:"dedicatedOrgId,omitempty" tf:"dedicated_org_id,omitempty"`

	// Network friendly description
	// Network description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more IP scopes for the network. See IP Scope below for details.
	// A set of IP scopes for the network
	IPScope []ExternalNetworkV2IPScopeInitParameters `json:"ipScope,omitempty" tf:"ip_scope,omitempty"`

	// Configure intentions for
	// NAT and Firewall rule configuration:
	// Defines intentions to configure NAT and Firewall rules (only with IP Spaces, VCD 10.5.1+) One of `PROVIDER_GATEWAY`,`EDGE_GATEWAY`,`PROVIDER_AND_EDGE_GATEWAY`
	NATAndFirewallServiceIntention *string `json:"natAndFirewallServiceIntention,omitempty" tf:"nat_and_firewall_service_intention,omitempty"`

	// A unique name for the network
	// Network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NSX-T network definition. See NSX-T Network below for details.
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork []NsxtNetworkInitParameters `json:"nsxtNetwork,omitempty" tf:"nsxt_network,omitempty"`

	// Configure intentions for
	// Org VDC network Route Advertisement:
	// Defines intentions to configure route advertisement (only with IP Spaces, VCD 10.5.1+) One of `IP_SPACE_UPLINKS_ADVERTISED_STRICT`,`IP_SPACE_UPLINKS_ADVERTISED_FLEXIBLE`,`ALL_NETWORKS_ADVERTISED`
	RouteAdvertisementIntention *string `json:"routeAdvertisementIntention,omitempty" tf:"route_advertisement_intention,omitempty"`

	// Defines if the network uses IP Spaces. Do
	// not specify ip_scope when using IP Spaces. (default false)
	// Enables IP Spaces for this network (default 'false'). VCD 10.4.1+
	UseIPSpaces *bool `json:"useIpSpaces,omitempty" tf:"use_ip_spaces,omitempty"`

	// One or more blocks of vSphere Network..
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system.
	VsphereNetwork []ExternalNetworkV2VsphereNetworkInitParameters `json:"vsphereNetwork,omitempty" tf:"vsphere_network,omitempty"`
}

type ExternalNetworkV2Observation struct {

	// An Org ID that this network should be
	// dedicated to. Only applicable when use_ip_spaces=true
	// Dedicate this External Network to an Org ID (only with IP Spaces, VCD 10.4.1+)
	DedicatedOrgID *string `json:"dedicatedOrgId,omitempty" tf:"dedicated_org_id,omitempty"`

	// Network friendly description
	// Network description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more IP scopes for the network. See IP Scope below for details.
	// A set of IP scopes for the network
	IPScope []ExternalNetworkV2IPScopeObservation `json:"ipScope,omitempty" tf:"ip_scope,omitempty"`

	// Configure intentions for
	// NAT and Firewall rule configuration:
	// Defines intentions to configure NAT and Firewall rules (only with IP Spaces, VCD 10.5.1+) One of `PROVIDER_GATEWAY`,`EDGE_GATEWAY`,`PROVIDER_AND_EDGE_GATEWAY`
	NATAndFirewallServiceIntention *string `json:"natAndFirewallServiceIntention,omitempty" tf:"nat_and_firewall_service_intention,omitempty"`

	// A unique name for the network
	// Network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NSX-T network definition. See NSX-T Network below for details.
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork []NsxtNetworkObservation `json:"nsxtNetwork,omitempty" tf:"nsxt_network,omitempty"`

	// Configure intentions for
	// Org VDC network Route Advertisement:
	// Defines intentions to configure route advertisement (only with IP Spaces, VCD 10.5.1+) One of `IP_SPACE_UPLINKS_ADVERTISED_STRICT`,`IP_SPACE_UPLINKS_ADVERTISED_FLEXIBLE`,`ALL_NETWORKS_ADVERTISED`
	RouteAdvertisementIntention *string `json:"routeAdvertisementIntention,omitempty" tf:"route_advertisement_intention,omitempty"`

	// Defines if the network uses IP Spaces. Do
	// not specify ip_scope when using IP Spaces. (default false)
	// Enables IP Spaces for this network (default 'false'). VCD 10.4.1+
	UseIPSpaces *bool `json:"useIpSpaces,omitempty" tf:"use_ip_spaces,omitempty"`

	// One or more blocks of vSphere Network..
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system.
	VsphereNetwork []ExternalNetworkV2VsphereNetworkObservation `json:"vsphereNetwork,omitempty" tf:"vsphere_network,omitempty"`
}

type ExternalNetworkV2Parameters struct {

	// An Org ID that this network should be
	// dedicated to. Only applicable when use_ip_spaces=true
	// Dedicate this External Network to an Org ID (only with IP Spaces, VCD 10.4.1+)
	// +kubebuilder:validation:Optional
	DedicatedOrgID *string `json:"dedicatedOrgId,omitempty" tf:"dedicated_org_id,omitempty"`

	// Network friendly description
	// Network description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more IP scopes for the network. See IP Scope below for details.
	// A set of IP scopes for the network
	// +kubebuilder:validation:Optional
	IPScope []ExternalNetworkV2IPScopeParameters `json:"ipScope,omitempty" tf:"ip_scope,omitempty"`

	// Configure intentions for
	// NAT and Firewall rule configuration:
	// Defines intentions to configure NAT and Firewall rules (only with IP Spaces, VCD 10.5.1+) One of `PROVIDER_GATEWAY`,`EDGE_GATEWAY`,`PROVIDER_AND_EDGE_GATEWAY`
	// +kubebuilder:validation:Optional
	NATAndFirewallServiceIntention *string `json:"natAndFirewallServiceIntention,omitempty" tf:"nat_and_firewall_service_intention,omitempty"`

	// A unique name for the network
	// Network name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NSX-T network definition. See NSX-T Network below for details.
	// Reference to NSX-T Tier-0 router or segment and manager
	// +kubebuilder:validation:Optional
	NsxtNetwork []NsxtNetworkParameters `json:"nsxtNetwork,omitempty" tf:"nsxt_network,omitempty"`

	// Configure intentions for
	// Org VDC network Route Advertisement:
	// Defines intentions to configure route advertisement (only with IP Spaces, VCD 10.5.1+) One of `IP_SPACE_UPLINKS_ADVERTISED_STRICT`,`IP_SPACE_UPLINKS_ADVERTISED_FLEXIBLE`,`ALL_NETWORKS_ADVERTISED`
	// +kubebuilder:validation:Optional
	RouteAdvertisementIntention *string `json:"routeAdvertisementIntention,omitempty" tf:"route_advertisement_intention,omitempty"`

	// Defines if the network uses IP Spaces. Do
	// not specify ip_scope when using IP Spaces. (default false)
	// Enables IP Spaces for this network (default 'false'). VCD 10.4.1+
	// +kubebuilder:validation:Optional
	UseIPSpaces *bool `json:"useIpSpaces,omitempty" tf:"use_ip_spaces,omitempty"`

	// One or more blocks of vSphere Network..
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system.
	// +kubebuilder:validation:Optional
	VsphereNetwork []ExternalNetworkV2VsphereNetworkParameters `json:"vsphereNetwork,omitempty" tf:"vsphere_network,omitempty"`
}

type ExternalNetworkV2VsphereNetworkInitParameters struct {

	// vSphere portgroup ID. Can be looked up using  vcd_portgroup data source.
	// The name of the port group
	PortgroupID *string `json:"portgroupId,omitempty" tf:"portgroup_id,omitempty"`

	// vCenter ID. Can be looked up using vcd_vcenter data source.
	// The vCenter server name
	VcenterID *string `json:"vcenterId,omitempty" tf:"vcenter_id,omitempty"`
}

type ExternalNetworkV2VsphereNetworkObservation struct {

	// vSphere portgroup ID. Can be looked up using  vcd_portgroup data source.
	// The name of the port group
	PortgroupID *string `json:"portgroupId,omitempty" tf:"portgroup_id,omitempty"`

	// vCenter ID. Can be looked up using vcd_vcenter data source.
	// The vCenter server name
	VcenterID *string `json:"vcenterId,omitempty" tf:"vcenter_id,omitempty"`
}

type ExternalNetworkV2VsphereNetworkParameters struct {

	// vSphere portgroup ID. Can be looked up using  vcd_portgroup data source.
	// The name of the port group
	// +kubebuilder:validation:Optional
	PortgroupID *string `json:"portgroupId" tf:"portgroup_id,omitempty"`

	// vCenter ID. Can be looked up using vcd_vcenter data source.
	// The vCenter server name
	// +kubebuilder:validation:Optional
	VcenterID *string `json:"vcenterId" tf:"vcenter_id,omitempty"`
}

type IPScopeStaticIPPoolInitParameters struct {

	// End address of the IP range
	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type IPScopeStaticIPPoolObservation struct {

	// End address of the IP range
	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type IPScopeStaticIPPoolParameters struct {

	// End address of the IP range
	// End address of the IP range
	// +kubebuilder:validation:Optional
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// Start address of the IP range
	// +kubebuilder:validation:Optional
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type NsxtNetworkInitParameters struct {

	// NSX-T manager ID. Can be looked up using vcd_nsxt_manager data source.
	// ID of NSX-T manager
	NsxtManagerID *string `json:"nsxtManagerId,omitempty" tf:"nsxt_manager_id,omitempty"`

	// Existing NSX-T segment name.
	// Name of NSX-T segment (for NSX-T segment backed external network)
	NsxtSegmentName *string `json:"nsxtSegmentName,omitempty" tf:"nsxt_segment_name,omitempty"`

	// NSX-T Tier-0 router ID. Can be looked up using
	// vcd_nsxt_tier0_router data source.
	// ID of NSX-T Tier-0 router (for T0 gateway backed external network)
	NsxtTier0RouterID *string `json:"nsxtTier0RouterId,omitempty" tf:"nsxt_tier0_router_id,omitempty"`
}

type NsxtNetworkObservation struct {

	// NSX-T manager ID. Can be looked up using vcd_nsxt_manager data source.
	// ID of NSX-T manager
	NsxtManagerID *string `json:"nsxtManagerId,omitempty" tf:"nsxt_manager_id,omitempty"`

	// Existing NSX-T segment name.
	// Name of NSX-T segment (for NSX-T segment backed external network)
	NsxtSegmentName *string `json:"nsxtSegmentName,omitempty" tf:"nsxt_segment_name,omitempty"`

	// NSX-T Tier-0 router ID. Can be looked up using
	// vcd_nsxt_tier0_router data source.
	// ID of NSX-T Tier-0 router (for T0 gateway backed external network)
	NsxtTier0RouterID *string `json:"nsxtTier0RouterId,omitempty" tf:"nsxt_tier0_router_id,omitempty"`
}

type NsxtNetworkParameters struct {

	// NSX-T manager ID. Can be looked up using vcd_nsxt_manager data source.
	// ID of NSX-T manager
	// +kubebuilder:validation:Optional
	NsxtManagerID *string `json:"nsxtManagerId" tf:"nsxt_manager_id,omitempty"`

	// Existing NSX-T segment name.
	// Name of NSX-T segment (for NSX-T segment backed external network)
	// +kubebuilder:validation:Optional
	NsxtSegmentName *string `json:"nsxtSegmentName,omitempty" tf:"nsxt_segment_name,omitempty"`

	// NSX-T Tier-0 router ID. Can be looked up using
	// vcd_nsxt_tier0_router data source.
	// ID of NSX-T Tier-0 router (for T0 gateway backed external network)
	// +kubebuilder:validation:Optional
	NsxtTier0RouterID *string `json:"nsxtTier0RouterId,omitempty" tf:"nsxt_tier0_router_id,omitempty"`
}

// ExternalNetworkV2Spec defines the desired state of ExternalNetworkV2
type ExternalNetworkV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalNetworkV2Parameters `json:"forProvider"`
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
	InitProvider ExternalNetworkV2InitParameters `json:"initProvider,omitempty"`
}

// ExternalNetworkV2Status defines the observed state of ExternalNetworkV2.
type ExternalNetworkV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalNetworkV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ExternalNetworkV2 is the Schema for the ExternalNetworkV2s API. Provides a VMware Cloud Director External Network resource (version 2). New version of this resource uses new VCD API and is capable of creating NSX-T backed external networks as well as port group backed ones.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type ExternalNetworkV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ExternalNetworkV2Spec   `json:"spec"`
	Status ExternalNetworkV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalNetworkV2List contains a list of ExternalNetworkV2s
type ExternalNetworkV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalNetworkV2 `json:"items"`
}

// Repository type metadata.
var (
	ExternalNetworkV2_Kind             = "ExternalNetworkV2"
	ExternalNetworkV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalNetworkV2_Kind}.String()
	ExternalNetworkV2_KindAPIVersion   = ExternalNetworkV2_Kind + "." + CRDGroupVersion.String()
	ExternalNetworkV2_GroupVersionKind = CRDGroupVersion.WithKind(ExternalNetworkV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalNetworkV2{}, &ExternalNetworkV2List{})
}
