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

type VAppNetworkDHCPPoolInitParameters struct {

	// The default DHCP lease time to use. Defaults to 3600.
	DefaultLeaseTime *float64 `json:"defaultLeaseTime,omitempty" tf:"default_lease_time,omitempty"`

	// Allows to enable or disable service. Default is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The final address in the IP Range.
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The maximum DHCP lease time to use. Defaults to 7200.
	MaxLeaseTime *float64 `json:"maxLeaseTime,omitempty" tf:"max_lease_time,omitempty"`

	// The first address in the IP Range.
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type VAppNetworkDHCPPoolObservation struct {

	// The default DHCP lease time to use. Defaults to 3600.
	DefaultLeaseTime *float64 `json:"defaultLeaseTime,omitempty" tf:"default_lease_time,omitempty"`

	// Allows to enable or disable service. Default is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The final address in the IP Range.
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The maximum DHCP lease time to use. Defaults to 7200.
	MaxLeaseTime *float64 `json:"maxLeaseTime,omitempty" tf:"max_lease_time,omitempty"`

	// The first address in the IP Range.
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type VAppNetworkDHCPPoolParameters struct {

	// The default DHCP lease time to use. Defaults to 3600.
	// +kubebuilder:validation:Optional
	DefaultLeaseTime *float64 `json:"defaultLeaseTime,omitempty" tf:"default_lease_time,omitempty"`

	// Allows to enable or disable service. Default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The final address in the IP Range.
	// +kubebuilder:validation:Optional
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The maximum DHCP lease time to use. Defaults to 7200.
	// +kubebuilder:validation:Optional
	MaxLeaseTime *float64 `json:"maxLeaseTime,omitempty" tf:"max_lease_time,omitempty"`

	// The first address in the IP Range.
	// +kubebuilder:validation:Optional
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type VAppNetworkInitParameters struct {

	// A range of IPs to issue to virtual machines that don't have a static IP; see IP Pools below for details.
	// A range of IPs to issue to virtual machines that don't have a static IP
	DHCPPool []VAppNetworkDHCPPoolInitParameters `json:"dhcpPool,omitempty" tf:"dhcp_pool,omitempty"`

	// A FQDN for the virtual machines on this network.
	// DNS suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Description of vApp network
	// Optional description for the network
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// First DNS server to use.
	// Primary DNS server
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Second DNS server to use.
	// Secondary DNS server
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// The gateway for this network.
	// Gateway of the network
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// True if Network allows guest VLAN tagging.
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed *bool `json:"guestVlanAllowed,omitempty" tf:"guest_vlan_allowed,omitempty"`

	// A unique name for the network.
	// vApp network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Deprecated) Use prefix_length instead. The netmask for the new network.
	// Netmask address for a subnet.
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// An Org network name to which vApp network is connected. If not configured, then an isolated network is created.
	// org network name to which vapp network is connected
	OrgNetworkName *string `json:"orgNetworkName,omitempty" tf:"org_network_name,omitempty"`

	// The subnet prefix length for the network.
	// Prefix length for a subnet
	PrefixLength *string `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// VCD 10.4.1+ API prohibits removal of vApp
	// network from a powered on vApp. Set to true to power off the vApp during vApp network removal.
	// If the vApp's original state was powered on, it will be powered back on after removing the
	// network. (default false) Note. It only affects delete operation for the resource and will
	// never power cycle vApp during update operations. Changing this value will cause plan change, but
	// update will be a no-op operation.
	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	RebootVappOnRemoval *bool `json:"rebootVappOnRemoval,omitempty" tf:"reboot_vapp_on_removal,omitempty"`

	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIPMacEnabled *bool `json:"retainIpMacEnabled,omitempty" tf:"retain_ip_mac_enabled,omitempty"`

	// A range of IPs permitted to be used as static IPs for virtual machines; see IP Pools below for details.
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIPPool []VAppNetworkStaticIPPoolInitParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The vApp this network belongs to.
	// vApp to use
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppNetworkObservation struct {

	// A range of IPs to issue to virtual machines that don't have a static IP; see IP Pools below for details.
	// A range of IPs to issue to virtual machines that don't have a static IP
	DHCPPool []VAppNetworkDHCPPoolObservation `json:"dhcpPool,omitempty" tf:"dhcp_pool,omitempty"`

	// A FQDN for the virtual machines on this network.
	// DNS suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Description of vApp network
	// Optional description for the network
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// First DNS server to use.
	// Primary DNS server
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Second DNS server to use.
	// Secondary DNS server
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// The gateway for this network.
	// Gateway of the network
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// True if Network allows guest VLAN tagging.
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed *bool `json:"guestVlanAllowed,omitempty" tf:"guest_vlan_allowed,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the network.
	// vApp network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Deprecated) Use prefix_length instead. The netmask for the new network.
	// Netmask address for a subnet.
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// An Org network name to which vApp network is connected. If not configured, then an isolated network is created.
	// org network name to which vapp network is connected
	OrgNetworkName *string `json:"orgNetworkName,omitempty" tf:"org_network_name,omitempty"`

	// The subnet prefix length for the network.
	// Prefix length for a subnet
	PrefixLength *string `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// VCD 10.4.1+ API prohibits removal of vApp
	// network from a powered on vApp. Set to true to power off the vApp during vApp network removal.
	// If the vApp's original state was powered on, it will be powered back on after removing the
	// network. (default false) Note. It only affects delete operation for the resource and will
	// never power cycle vApp during update operations. Changing this value will cause plan change, but
	// update will be a no-op operation.
	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	RebootVappOnRemoval *bool `json:"rebootVappOnRemoval,omitempty" tf:"reboot_vapp_on_removal,omitempty"`

	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIPMacEnabled *bool `json:"retainIpMacEnabled,omitempty" tf:"retain_ip_mac_enabled,omitempty"`

	// A range of IPs permitted to be used as static IPs for virtual machines; see IP Pools below for details.
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIPPool []VAppNetworkStaticIPPoolObservation `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The vApp this network belongs to.
	// vApp to use
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppNetworkParameters struct {

	// A range of IPs to issue to virtual machines that don't have a static IP; see IP Pools below for details.
	// A range of IPs to issue to virtual machines that don't have a static IP
	// +kubebuilder:validation:Optional
	DHCPPool []VAppNetworkDHCPPoolParameters `json:"dhcpPool,omitempty" tf:"dhcp_pool,omitempty"`

	// A FQDN for the virtual machines on this network.
	// DNS suffix
	// +kubebuilder:validation:Optional
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Description of vApp network
	// Optional description for the network
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// First DNS server to use.
	// Primary DNS server
	// +kubebuilder:validation:Optional
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Second DNS server to use.
	// Secondary DNS server
	// +kubebuilder:validation:Optional
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// The gateway for this network.
	// Gateway of the network
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// True if Network allows guest VLAN tagging.
	// True if Network allows guest VLAN tagging
	// +kubebuilder:validation:Optional
	GuestVlanAllowed *bool `json:"guestVlanAllowed,omitempty" tf:"guest_vlan_allowed,omitempty"`

	// A unique name for the network.
	// vApp network name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Deprecated) Use prefix_length instead. The netmask for the new network.
	// Netmask address for a subnet.
	// +kubebuilder:validation:Optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// An Org network name to which vApp network is connected. If not configured, then an isolated network is created.
	// org network name to which vapp network is connected
	// +kubebuilder:validation:Optional
	OrgNetworkName *string `json:"orgNetworkName,omitempty" tf:"org_network_name,omitempty"`

	// The subnet prefix length for the network.
	// Prefix length for a subnet
	// +kubebuilder:validation:Optional
	PrefixLength *string `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// VCD 10.4.1+ API prohibits removal of vApp
	// network from a powered on vApp. Set to true to power off the vApp during vApp network removal.
	// If the vApp's original state was powered on, it will be powered back on after removing the
	// network. (default false) Note. It only affects delete operation for the resource and will
	// never power cycle vApp during update operations. Changing this value will cause plan change, but
	// update will be a no-op operation.
	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	// +kubebuilder:validation:Optional
	RebootVappOnRemoval *bool `json:"rebootVappOnRemoval,omitempty" tf:"reboot_vapp_on_removal,omitempty"`

	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	// +kubebuilder:validation:Optional
	RetainIPMacEnabled *bool `json:"retainIpMacEnabled,omitempty" tf:"retain_ip_mac_enabled,omitempty"`

	// A range of IPs permitted to be used as static IPs for virtual machines; see IP Pools below for details.
	// A range of IPs permitted to be used as static IPs for virtual machines
	// +kubebuilder:validation:Optional
	StaticIPPool []VAppNetworkStaticIPPoolParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The vApp this network belongs to.
	// vApp to use
	// +kubebuilder:validation:Optional
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppNetworkStaticIPPoolInitParameters struct {

	// The final address in the IP Range.
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The first address in the IP Range.
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type VAppNetworkStaticIPPoolObservation struct {

	// The final address in the IP Range.
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The first address in the IP Range.
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type VAppNetworkStaticIPPoolParameters struct {

	// The final address in the IP Range.
	// +kubebuilder:validation:Optional
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// The first address in the IP Range.
	// +kubebuilder:validation:Optional
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

// VAppNetworkSpec defines the desired state of VAppNetwork
type VAppNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VAppNetworkParameters `json:"forProvider"`
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
	InitProvider VAppNetworkInitParameters `json:"initProvider,omitempty"`
}

// VAppNetworkStatus defines the observed state of VAppNetwork.
type VAppNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VAppNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VAppNetwork is the Schema for the VAppNetworks API. Allows to provision a vApp network and optionally connect it to an existing Org VDC network.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VAppNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gateway) || (has(self.initProvider) && has(self.initProvider.gateway))",message="spec.forProvider.gateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vappName) || (has(self.initProvider) && has(self.initProvider.vappName))",message="spec.forProvider.vappName is a required parameter"
	Spec   VAppNetworkSpec   `json:"spec"`
	Status VAppNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VAppNetworkList contains a list of VAppNetworks
type VAppNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VAppNetwork `json:"items"`
}

// Repository type metadata.
var (
	VAppNetwork_Kind             = "VAppNetwork"
	VAppNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VAppNetwork_Kind}.String()
	VAppNetwork_KindAPIVersion   = VAppNetwork_Kind + "." + CRDGroupVersion.String()
	VAppNetwork_GroupVersionKind = CRDGroupVersion.WithKind(VAppNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&VAppNetwork{}, &VAppNetworkList{})
}
