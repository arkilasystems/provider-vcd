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

type AllocatedIpsInitParameters struct {

	// - End IP address of a range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// - Start IP address of a range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type AllocatedIpsObservation struct {

	// - End IP address of a range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// - Start IP address of a range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type AllocatedIpsParameters struct {

	// - End IP address of a range
	// +kubebuilder:validation:Optional
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// - Start IP address of a range
	// +kubebuilder:validation:Optional
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type NsxtEdgeGatewayExternalNetworkInitParameters struct {

	// Number of allocated IPs
	// Number of allocated IPs
	AllocatedIPCount *float64 `json:"allocatedIpCount,omitempty" tf:"allocated_ip_count,omitempty"`

	// - ID of NSX-T Segment backed external network
	// NSX-T Segment backed External Network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway IP Address
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should fall
	// in provided gateway and prefix_length
	// Primary IP address for the Edge Gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type NsxtEdgeGatewayExternalNetworkObservation struct {

	// Number of allocated IPs
	// Number of allocated IPs
	AllocatedIPCount *float64 `json:"allocatedIpCount,omitempty" tf:"allocated_ip_count,omitempty"`

	// - ID of NSX-T Segment backed external network
	// NSX-T Segment backed External Network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway IP Address
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should fall
	// in provided gateway and prefix_length
	// Primary IP address for the Edge Gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type NsxtEdgeGatewayExternalNetworkParameters struct {

	// Number of allocated IPs
	// Number of allocated IPs
	// +kubebuilder:validation:Optional
	AllocatedIPCount *float64 `json:"allocatedIpCount" tf:"allocated_ip_count,omitempty"`

	// - ID of NSX-T Segment backed external network
	// NSX-T Segment backed External Network ID
	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId" tf:"external_network_id,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway IP Address
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24)
	// Prefix length for a subnet (e.g. 24)
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should fall
	// in provided gateway and prefix_length
	// Primary IP address for the Edge Gateway - will be auto-assigned if not defined
	// +kubebuilder:validation:Optional
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type NsxtEdgeGatewayInitParameters struct {

	// Dedicating the external network will enable Route Advertisement for this Edge Gateway. Default false.
	// Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
	DedicateExternalNetwork *bool `json:"dedicateExternalNetwork,omitempty" tf:"dedicate_external_network,omitempty"`

	// ACTIVE_STANDBY (default) or
	// DISTRIBUTED_ONLY (VCD 10.6+). Distributed-only does not provide services that run on service
	// router such as firewalling, NAT, VPN, DNS forwarding or static routes, instead, the distributed
	// nature guarantees high north-south data throughput.
	// Edge Gateway deployment mode. One of 'DISTRIBUTED_ONLY', 'ACTIVE_STANDBY'. Default 'ACTIVE_STANDBY'. VCD 10.6+
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific Edge Cluster ID if required
	// Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
	EdgeClusterID *string `json:"edgeClusterId,omitempty" tf:"edge_cluster_id,omitempty"`

	// attaches NSX-T Segment backed External
	// Networks with a given configuration block. It does not
	// support IP Spaces.
	// Additional NSX-T Segment Backed networks to attach
	ExternalNetwork []NsxtEdgeGatewayExternalNetworkInitParameters `json:"externalNetwork,omitempty" tf:"external_network,omitempty"`

	// An external network ID. Note: Data source vcd_external_network_v2
	// can be used to lookup ID by name.
	// External network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// Sets a limit of IPs to count for
	// used_ip_count and unused_ip_count attributes to avoid exhausting compute resource while
	// counting IPs in large IPv6 subnets. It does not affect operation of Edge Gateway configuration,
	// only IP count reporting. Defaults to 1000000, update is a no-op, but will affect newly read
	// data. While it is unlikely that a single Edge Gateway can effectively manage more IPs, one can
	// specify 0 for unlimited value.
	// How many maximum IPs should be reported in 'used_ipcount' and 'unused_ip_count'. Default 1000000, 0 - unlimited
	IPCountReadLimit *float64 `json:"ipCountReadLimit,omitempty" tf:"ip_count_read_limit,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Non-Distributed routing will allow
	// tenants the option of connecting Org Vdc networks to the Service
	// Router. This would
	// force all VM traffic through the service router for that network. Default false.
	// A flag indicating whether non-distributed routing is enabled or not (`false` by default)
	NonDistributedRoutingEnabled *bool `json:"nonDistributedRoutingEnabled,omitempty" tf:"non_distributed_routing_enabled,omitempty"`

	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The ID of VDC or VDC Group. Note: Data sources
	// vcd_vdc_group or
	// vcd_org_vdc can be used to lookup IDs by
	// name.
	// ID of VDC or VDC Group
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// If owner_id is a VDC Group, by default Edge
	// Gateway will be created in random member VDC and moved to destination VDC Group. This field allows
	// to specify initial VDC for Edge Gateway (this can define Egress location of traffic in the VDC
	// Group) Note: It can only be used when owner_id is a VDC Group.
	// Optional ID of starting VDC if the 'owner_id' is a VDC Group
	StartingVdcID *string `json:"startingVdcId,omitempty" tf:"starting_vdc_id,omitempty"`

	// One or more subnets defined for Edge Gateway. One of
	// subnet, subnet_with_total_ip_count or subnet_with_ip_count is required unless parent
	// network is backed by IP Spaces. Read more in IP allocation modes section.
	// One or more blocks with external network information to be attached to this gateway's interface including IP allocation ranges
	Subnet []NsxtEdgeGatewaySubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (v3.9+) One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Auto allocation of subnets by using per subnet IP allocation counts
	SubnetWithIPCount []SubnetWithIPCountInitParameters `json:"subnetWithIpCount,omitempty" tf:"subnet_with_ip_count,omitempty"`

	// One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Subnet definitions for this Edge Gateway. IP allocation is controlled using 'total_allocated_ip_count'
	SubnetWithTotalIPCount []SubnetWithTotalIPCountInitParameters `json:"subnetWithTotalIpCount,omitempty" tf:"subnet_with_total_ip_count,omitempty"`

	// Required with subnet_with_total_ip_count. It is
	// read-only attribute with other other allocation models subnet and subnet_with_ip_count.
	// Note. It sets or reports IP count only for NSX-T Tier 0 backed external network Uplink.
	// Total number of IP addresses allocated for this gateway from Tier0 uplink. Can be set with 'subnet_with_total_ip_count' definitions only
	TotalAllocatedIPCount *float64 `json:"totalAllocatedIpCount,omitempty" tf:"total_allocated_ip_count,omitempty"`

	// Deprecated in favor of owner_id. The name of VDC that owns the edge
	// gateway. Can be inherited from provider configuration if not defined here.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtEdgeGatewayObservation struct {

	// Dedicating the external network will enable Route Advertisement for this Edge Gateway. Default false.
	// Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
	DedicateExternalNetwork *bool `json:"dedicateExternalNetwork,omitempty" tf:"dedicate_external_network,omitempty"`

	// ACTIVE_STANDBY (default) or
	// DISTRIBUTED_ONLY (VCD 10.6+). Distributed-only does not provide services that run on service
	// router such as firewalling, NAT, VPN, DNS forwarding or static routes, instead, the distributed
	// nature guarantees high north-south data throughput.
	// Edge Gateway deployment mode. One of 'DISTRIBUTED_ONLY', 'ACTIVE_STANDBY'. Default 'ACTIVE_STANDBY'. VCD 10.6+
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific Edge Cluster ID if required
	// Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
	EdgeClusterID *string `json:"edgeClusterId,omitempty" tf:"edge_cluster_id,omitempty"`

	// attaches NSX-T Segment backed External
	// Networks with a given configuration block. It does not
	// support IP Spaces.
	// Additional NSX-T Segment Backed networks to attach
	ExternalNetwork []NsxtEdgeGatewayExternalNetworkObservation `json:"externalNetwork,omitempty" tf:"external_network,omitempty"`

	// Total allocated IP count in attached NSX-T Segment backed
	// external networks
	// Total number of IPs allocated for this Gateway from NSX-T Segment backed External Network uplinks
	ExternalNetworkAllocatedIPCount *float64 `json:"externalNetworkAllocatedIpCount,omitempty" tf:"external_network_allocated_ip_count,omitempty"`

	// An external network ID. Note: Data source vcd_external_network_v2
	// can be used to lookup ID by name.
	// External network ID
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Sets a limit of IPs to count for
	// used_ip_count and unused_ip_count attributes to avoid exhausting compute resource while
	// counting IPs in large IPv6 subnets. It does not affect operation of Edge Gateway configuration,
	// only IP count reporting. Defaults to 1000000, update is a no-op, but will affect newly read
	// data. While it is unlikely that a single Edge Gateway can effectively manage more IPs, one can
	// specify 0 for unlimited value.
	// How many maximum IPs should be reported in 'used_ipcount' and 'unused_ip_count'. Default 1000000, 0 - unlimited
	IPCountReadLimit *float64 `json:"ipCountReadLimit,omitempty" tf:"ip_count_read_limit,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Non-Distributed routing will allow
	// tenants the option of connecting Org Vdc networks to the Service
	// Router. This would
	// force all VM traffic through the service router for that network. Default false.
	// A flag indicating whether non-distributed routing is enabled or not (`false` by default)
	NonDistributedRoutingEnabled *bool `json:"nonDistributedRoutingEnabled,omitempty" tf:"non_distributed_routing_enabled,omitempty"`

	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The ID of VDC or VDC Group. Note: Data sources
	// vcd_vdc_group or
	// vcd_org_vdc can be used to lookup IDs by
	// name.
	// ID of VDC or VDC Group
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Primary IP address exposed for an easy access without nesting.
	// Primary IP address of edge gateway. Read-only (can be specified in specific subnet)
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`

	// If owner_id is a VDC Group, by default Edge
	// Gateway will be created in random member VDC and moved to destination VDC Group. This field allows
	// to specify initial VDC for Edge Gateway (this can define Egress location of traffic in the VDC
	// Group) Note: It can only be used when owner_id is a VDC Group.
	// Optional ID of starting VDC if the 'owner_id' is a VDC Group
	StartingVdcID *string `json:"startingVdcId,omitempty" tf:"starting_vdc_id,omitempty"`

	// One or more subnets defined for Edge Gateway. One of
	// subnet, subnet_with_total_ip_count or subnet_with_ip_count is required unless parent
	// network is backed by IP Spaces. Read more in IP allocation modes section.
	// One or more blocks with external network information to be attached to this gateway's interface including IP allocation ranges
	Subnet []NsxtEdgeGatewaySubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (v3.9+) One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Auto allocation of subnets by using per subnet IP allocation counts
	SubnetWithIPCount []SubnetWithIPCountObservation `json:"subnetWithIpCount,omitempty" tf:"subnet_with_ip_count,omitempty"`

	// One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Subnet definitions for this Edge Gateway. IP allocation is controlled using 'total_allocated_ip_count'
	SubnetWithTotalIPCount []SubnetWithTotalIPCountObservation `json:"subnetWithTotalIpCount,omitempty" tf:"subnet_with_total_ip_count,omitempty"`

	// Required with subnet_with_total_ip_count. It is
	// read-only attribute with other other allocation models subnet and subnet_with_ip_count.
	// Note. It sets or reports IP count only for NSX-T Tier 0 backed external network Uplink.
	// Total number of IP addresses allocated for this gateway from Tier0 uplink. Can be set with 'subnet_with_total_ip_count' definitions only
	TotalAllocatedIPCount *float64 `json:"totalAllocatedIpCount,omitempty" tf:"total_allocated_ip_count,omitempty"`

	// Unused IP count in this Edge Gateway (for all uplinks). Note: it is not
	// exposed when using IP Spaces.
	// Number of unused IP addresses
	UnusedIPCount *float64 `json:"unusedIpCount,omitempty" tf:"unused_ip_count,omitempty"`

	// Boolean value that hints if the NSX-T Edge Gateway uses IP Spaces
	// Boolean value that specifies that the Edge Gateway is using IP Spaces
	UseIPSpaces *bool `json:"useIpSpaces,omitempty" tf:"use_ip_spaces,omitempty"`

	// Used IP count in this Edge Gateway (for all uplinks). Note: it is not
	// exposed when using IP Spaces.
	// Number of used IP addresses
	UsedIPCount *float64 `json:"usedIpCount,omitempty" tf:"used_ip_count,omitempty"`

	// Deprecated in favor of owner_id. The name of VDC that owns the edge
	// gateway. Can be inherited from provider configuration if not defined here.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtEdgeGatewayParameters struct {

	// Dedicating the external network will enable Route Advertisement for this Edge Gateway. Default false.
	// Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
	// +kubebuilder:validation:Optional
	DedicateExternalNetwork *bool `json:"dedicateExternalNetwork,omitempty" tf:"dedicate_external_network,omitempty"`

	// ACTIVE_STANDBY (default) or
	// DISTRIBUTED_ONLY (VCD 10.6+). Distributed-only does not provide services that run on service
	// router such as firewalling, NAT, VPN, DNS forwarding or static routes, instead, the distributed
	// nature guarantees high north-south data throughput.
	// Edge Gateway deployment mode. One of 'DISTRIBUTED_ONLY', 'ACTIVE_STANDBY'. Default 'ACTIVE_STANDBY'. VCD 10.6+
	// +kubebuilder:validation:Optional
	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific Edge Cluster ID if required
	// Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
	// +kubebuilder:validation:Optional
	EdgeClusterID *string `json:"edgeClusterId,omitempty" tf:"edge_cluster_id,omitempty"`

	// attaches NSX-T Segment backed External
	// Networks with a given configuration block. It does not
	// support IP Spaces.
	// Additional NSX-T Segment Backed networks to attach
	// +kubebuilder:validation:Optional
	ExternalNetwork []NsxtEdgeGatewayExternalNetworkParameters `json:"externalNetwork,omitempty" tf:"external_network,omitempty"`

	// An external network ID. Note: Data source vcd_external_network_v2
	// can be used to lookup ID by name.
	// External network ID
	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// Sets a limit of IPs to count for
	// used_ip_count and unused_ip_count attributes to avoid exhausting compute resource while
	// counting IPs in large IPv6 subnets. It does not affect operation of Edge Gateway configuration,
	// only IP count reporting. Defaults to 1000000, update is a no-op, but will affect newly read
	// data. While it is unlikely that a single Edge Gateway can effectively manage more IPs, one can
	// specify 0 for unlimited value.
	// How many maximum IPs should be reported in 'used_ipcount' and 'unused_ip_count'. Default 1000000, 0 - unlimited
	// +kubebuilder:validation:Optional
	IPCountReadLimit *float64 `json:"ipCountReadLimit,omitempty" tf:"ip_count_read_limit,omitempty"`

	// A unique name for the edge gateway.
	// Edge Gateway name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Non-Distributed routing will allow
	// tenants the option of connecting Org Vdc networks to the Service
	// Router. This would
	// force all VM traffic through the service router for that network. Default false.
	// A flag indicating whether non-distributed routing is enabled or not (`false` by default)
	// +kubebuilder:validation:Optional
	NonDistributedRoutingEnabled *bool `json:"nonDistributedRoutingEnabled,omitempty" tf:"non_distributed_routing_enabled,omitempty"`

	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The ID of VDC or VDC Group. Note: Data sources
	// vcd_vdc_group or
	// vcd_org_vdc can be used to lookup IDs by
	// name.
	// ID of VDC or VDC Group
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// If owner_id is a VDC Group, by default Edge
	// Gateway will be created in random member VDC and moved to destination VDC Group. This field allows
	// to specify initial VDC for Edge Gateway (this can define Egress location of traffic in the VDC
	// Group) Note: It can only be used when owner_id is a VDC Group.
	// Optional ID of starting VDC if the 'owner_id' is a VDC Group
	// +kubebuilder:validation:Optional
	StartingVdcID *string `json:"startingVdcId,omitempty" tf:"starting_vdc_id,omitempty"`

	// One or more subnets defined for Edge Gateway. One of
	// subnet, subnet_with_total_ip_count or subnet_with_ip_count is required unless parent
	// network is backed by IP Spaces. Read more in IP allocation modes section.
	// One or more blocks with external network information to be attached to this gateway's interface including IP allocation ranges
	// +kubebuilder:validation:Optional
	Subnet []NsxtEdgeGatewaySubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (v3.9+) One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Auto allocation of subnets by using per subnet IP allocation counts
	// +kubebuilder:validation:Optional
	SubnetWithIPCount []SubnetWithIPCountParameters `json:"subnetWithIpCount,omitempty" tf:"subnet_with_ip_count,omitempty"`

	// One or more
	// subnets defined for Edge Gateway. One of subnet,
	// subnet_with_total_ip_count or subnet_with_ip_count is required unless parent network is
	// backed by IP Spaces. Read more in IP allocation modes section.
	// Subnet definitions for this Edge Gateway. IP allocation is controlled using 'total_allocated_ip_count'
	// +kubebuilder:validation:Optional
	SubnetWithTotalIPCount []SubnetWithTotalIPCountParameters `json:"subnetWithTotalIpCount,omitempty" tf:"subnet_with_total_ip_count,omitempty"`

	// Required with subnet_with_total_ip_count. It is
	// read-only attribute with other other allocation models subnet and subnet_with_ip_count.
	// Note. It sets or reports IP count only for NSX-T Tier 0 backed external network Uplink.
	// Total number of IP addresses allocated for this gateway from Tier0 uplink. Can be set with 'subnet_with_total_ip_count' definitions only
	// +kubebuilder:validation:Optional
	TotalAllocatedIPCount *float64 `json:"totalAllocatedIpCount,omitempty" tf:"total_allocated_ip_count,omitempty"`

	// Deprecated in favor of owner_id. The name of VDC that owns the edge
	// gateway. Can be inherited from provider configuration if not defined here.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtEdgeGatewaySubnetInitParameters struct {

	// One or more blocks of ip ranges
	// in the subnet to be allocated
	// Define one or more blocks to sub-allocate pools on the edge gateway
	AllocatedIps []AllocatedIpsInitParameters `json:"allocatedIps,omitempty" tf:"allocated_ips,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask
	// of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// - Primary IP address for edge gateway. Note: primary_ip must fall
	// into allocated_ips block range as otherwise plan will not be clean with a new range defined for
	// that particular block. There can only be one primary_ip defined for edge gateway.
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type NsxtEdgeGatewaySubnetObservation struct {

	// One or more blocks of ip ranges
	// in the subnet to be allocated
	// Define one or more blocks to sub-allocate pools on the edge gateway
	AllocatedIps []AllocatedIpsObservation `json:"allocatedIps,omitempty" tf:"allocated_ips,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask
	// of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// - Primary IP address for edge gateway. Note: primary_ip must fall
	// into allocated_ips block range as otherwise plan will not be clean with a new range defined for
	// that particular block. There can only be one primary_ip defined for edge gateway.
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type NsxtEdgeGatewaySubnetParameters struct {

	// One or more blocks of ip ranges
	// in the subnet to be allocated
	// Define one or more blocks to sub-allocate pools on the edge gateway
	// +kubebuilder:validation:Optional
	AllocatedIps []AllocatedIpsParameters `json:"allocatedIps,omitempty" tf:"allocated_ips,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask
	// of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength" tf:"prefix_length,omitempty"`

	// - Primary IP address for edge gateway. Note: primary_ip must fall
	// into allocated_ips block range as otherwise plan will not be clean with a new range defined for
	// that particular block. There can only be one primary_ip defined for edge gateway.
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	// +kubebuilder:validation:Optional
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithIPCountInitParameters struct {

	// Number of allocated IPs from that particular subnet
	// Number of IP addresses to allocate
	AllocatedIPCount *float64 `json:"allocatedIpCount,omitempty" tf:"allocated_ip_count,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithIPCountObservation struct {

	// Number of allocated IPs from that particular subnet
	// Number of IP addresses to allocate
	AllocatedIPCount *float64 `json:"allocatedIpCount,omitempty" tf:"allocated_ip_count,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithIPCountParameters struct {

	// Number of allocated IPs from that particular subnet
	// Number of IP addresses to allocate
	// +kubebuilder:validation:Optional
	AllocatedIPCount *float64 `json:"allocatedIpCount" tf:"allocated_ip_count,omitempty"`

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	// +kubebuilder:validation:Optional
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithTotalIPCountInitParameters struct {

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_total_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithTotalIPCountObservation struct {

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_total_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

type SubnetWithTotalIPCountParameters struct {

	// - Gateway for a subnet in external network
	// Gateway address for a subnet
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// Prefix length of a subnet in external network (e.g. 24 for netmask of 255.255.255.0)
	// Prefix length for a subnet (e.g. 24)
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength" tf:"prefix_length,omitempty"`

	// Exactly one Primary IP is required for an Edge Gateway. It should be
	// defined in any of the subnet_with_total_ip_count blocks
	// Primary IP address for the edge gateway - will be auto-assigned if not defined
	// +kubebuilder:validation:Optional
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`
}

// NsxtEdgeGatewaySpec defines the desired state of NsxtEdgeGateway
type NsxtEdgeGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtEdgeGatewayParameters `json:"forProvider"`
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
	InitProvider NsxtEdgeGatewayInitParameters `json:"initProvider,omitempty"`
}

// NsxtEdgeGatewayStatus defines the observed state of NsxtEdgeGateway.
type NsxtEdgeGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtEdgeGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxtEdgeGateway is the Schema for the NsxtEdgeGateways API. Provides a VMware Cloud Director NSX-T edge gateway. This can be used to create, update, and delete NSX-T edge gateways connected to external networks.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtEdgeGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.externalNetworkId) || (has(self.initProvider) && has(self.initProvider.externalNetworkId))",message="spec.forProvider.externalNetworkId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NsxtEdgeGatewaySpec   `json:"spec"`
	Status NsxtEdgeGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtEdgeGatewayList contains a list of NsxtEdgeGateways
type NsxtEdgeGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtEdgeGateway `json:"items"`
}

// Repository type metadata.
var (
	NsxtEdgeGateway_Kind             = "NsxtEdgeGateway"
	NsxtEdgeGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtEdgeGateway_Kind}.String()
	NsxtEdgeGateway_KindAPIVersion   = NsxtEdgeGateway_Kind + "." + CRDGroupVersion.String()
	NsxtEdgeGateway_GroupVersionKind = CRDGroupVersion.WithKind(NsxtEdgeGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtEdgeGateway{}, &NsxtEdgeGatewayList{})
}
