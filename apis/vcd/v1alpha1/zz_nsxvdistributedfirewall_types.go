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

type ApplicationInitParameters struct {

	// The destination port used by the application, if any
	// Destination port for this application. Leaving it empty means 'any' port
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of application (Application, ApplicationGroup)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// , ports, source_port, destination_port
	// Protocol of the application (one of TCP, UDP, ICMP) (When not using name/value)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The source port used by the application, if any
	// Source port for this application. Leaving it empty means 'any' port
	SourcePort *string `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of application
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the application
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ApplicationObservation struct {

	// The destination port used by the application, if any
	// Destination port for this application. Leaving it empty means 'any' port
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of application (Application, ApplicationGroup)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// , ports, source_port, destination_port
	// Protocol of the application (one of TCP, UDP, ICMP) (When not using name/value)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The source port used by the application, if any
	// Source port for this application. Leaving it empty means 'any' port
	SourcePort *string `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of application
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the application
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ApplicationParameters struct {

	// The destination port used by the application, if any
	// Destination port for this application. Leaving it empty means 'any' port
	// +kubebuilder:validation:Optional
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of application (Application, ApplicationGroup)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// , ports, source_port, destination_port
	// Protocol of the application (one of TCP, UDP, ICMP) (When not using name/value)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The source port used by the application, if any
	// Source port for this application. Leaving it empty means 'any' port
	// +kubebuilder:validation:Optional
	SourcePort *string `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of application
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the application
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AppliedToInitParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the applied-to entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the applied-to entity (one of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the applied-to entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AppliedToObservation struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the applied-to entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the applied-to entity (one of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the applied-to entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AppliedToParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the applied-to entity
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the applied-to entity (one of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the applied-to entity
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DestinationInitParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the destination entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the destination entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the destination entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DestinationObservation struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the destination entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the destination entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the destination entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DestinationParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the destination entity
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the destination entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the destination entity
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type NsxvDistributedFirewallInitParameters struct {

	// One or more blocks with Firewall Rule definitions. Order
	// defines firewall rule precedence. If no rules are defined, all will be removed from the firewall
	// Ordered list of distributed firewall rules
	Rule []NsxvDistributedFirewallRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The ID of VDC to manage the Distributed Firewall in. Can be looked up using a vcd_org_vdc data source
	// The ID of VDC
	VdcID *string `json:"vdcId,omitempty" tf:"vdc_id,omitempty"`
}

type NsxvDistributedFirewallObservation struct {

	// Shows whether the NSX-V Distributed Firewall is enabled.
	// Shows whether the NSX-V distributed firewall is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more blocks with Firewall Rule definitions. Order
	// defines firewall rule precedence. If no rules are defined, all will be removed from the firewall
	// Ordered list of distributed firewall rules
	Rule []NsxvDistributedFirewallRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// The ID of VDC to manage the Distributed Firewall in. Can be looked up using a vcd_org_vdc data source
	// The ID of VDC
	VdcID *string `json:"vdcId,omitempty" tf:"vdc_id,omitempty"`
}

type NsxvDistributedFirewallParameters struct {

	// One or more blocks with Firewall Rule definitions. Order
	// defines firewall rule precedence. If no rules are defined, all will be removed from the firewall
	// Ordered list of distributed firewall rules
	// +kubebuilder:validation:Optional
	Rule []NsxvDistributedFirewallRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The ID of VDC to manage the Distributed Firewall in. Can be looked up using a vcd_org_vdc data source
	// The ID of VDC
	// +kubebuilder:validation:Optional
	VdcID *string `json:"vdcId,omitempty" tf:"vdc_id,omitempty"`
}

type NsxvDistributedFirewallRuleInitParameters struct {

	// Defines if it should allow or deny traffic
	// Action of the rule (allow, deny)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional set of applications to use for this rule. See below for Application objects
	// Application definitions for this rule. An empty value means 'any'
	Application []ApplicationInitParameters `json:"application,omitempty" tf:"application,omitempty"`

	// A set of objects to which the rule applies. See below for Source or destination objects
	// List of elements to which this rule applies
	AppliedTo []AppliedToInitParameters `json:"appliedTo,omitempty" tf:"applied_to,omitempty"`

	// A set of destination objects. See below for source or destination objects. Leaving it empty matches any (all)
	// List of destination traffic for this rule. An empty value means 'any'
	Destination []DestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// One of in, out, or inout
	// Direction of the rule (in, out, inout)
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Defines if the rule is enabled (default true)
	// Whether the rule is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// - reverses value of destination for the rule to match everything except specified objects
	// If true, the content of the destination elements is reversed
	ExcludeDestination *bool `json:"excludeDestination,omitempty" tf:"exclude_destination,omitempty"`

	// - reverses value of source for the rule to match everything except specified objects
	// If true, the content of the source elements is reversed
	ExcludeSource *bool `json:"excludeSource,omitempty" tf:"exclude_source,omitempty"`

	// Whether the rule traffic is logged
	Logged *bool `json:"logged,omitempty" tf:"logged,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Firewall Rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Packet type of the rule (any, ipv4, ipv6)
	PacketType *string `json:"packetType,omitempty" tf:"packet_type,omitempty"`

	// A set of source objects. See below for source or destination objects
	// Leaving it empty matches any (all)
	// List of source traffic for this rule. An empty value means 'any'
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type NsxvDistributedFirewallRuleObservation struct {

	// Defines if it should allow or deny traffic
	// Action of the rule (allow, deny)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional set of applications to use for this rule. See below for Application objects
	// Application definitions for this rule. An empty value means 'any'
	Application []ApplicationObservation `json:"application,omitempty" tf:"application,omitempty"`

	// A set of objects to which the rule applies. See below for Source or destination objects
	// List of elements to which this rule applies
	AppliedTo []AppliedToObservation `json:"appliedTo,omitempty" tf:"applied_to,omitempty"`

	// A set of destination objects. See below for source or destination objects. Leaving it empty matches any (all)
	// List of destination traffic for this rule. An empty value means 'any'
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// One of in, out, or inout
	// Direction of the rule (in, out, inout)
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Defines if the rule is enabled (default true)
	// Whether the rule is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// - reverses value of destination for the rule to match everything except specified objects
	// If true, the content of the destination elements is reversed
	ExcludeDestination *bool `json:"excludeDestination,omitempty" tf:"exclude_destination,omitempty"`

	// - reverses value of source for the rule to match everything except specified objects
	// If true, the content of the source elements is reversed
	ExcludeSource *bool `json:"excludeSource,omitempty" tf:"exclude_source,omitempty"`

	// Firewall Rule ID
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the rule traffic is logged
	Logged *bool `json:"logged,omitempty" tf:"logged,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Firewall Rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Packet type of the rule (any, ipv4, ipv6)
	PacketType *string `json:"packetType,omitempty" tf:"packet_type,omitempty"`

	// A set of source objects. See below for source or destination objects
	// Leaving it empty matches any (all)
	// List of source traffic for this rule. An empty value means 'any'
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type NsxvDistributedFirewallRuleParameters struct {

	// Defines if it should allow or deny traffic
	// Action of the rule (allow, deny)
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// An optional set of applications to use for this rule. See below for Application objects
	// Application definitions for this rule. An empty value means 'any'
	// +kubebuilder:validation:Optional
	Application []ApplicationParameters `json:"application,omitempty" tf:"application,omitempty"`

	// A set of objects to which the rule applies. See below for Source or destination objects
	// List of elements to which this rule applies
	// +kubebuilder:validation:Optional
	AppliedTo []AppliedToParameters `json:"appliedTo" tf:"applied_to,omitempty"`

	// A set of destination objects. See below for source or destination objects. Leaving it empty matches any (all)
	// List of destination traffic for this rule. An empty value means 'any'
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// One of in, out, or inout
	// Direction of the rule (in, out, inout)
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Defines if the rule is enabled (default true)
	// Whether the rule is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// - reverses value of destination for the rule to match everything except specified objects
	// If true, the content of the destination elements is reversed
	// +kubebuilder:validation:Optional
	ExcludeDestination *bool `json:"excludeDestination,omitempty" tf:"exclude_destination,omitempty"`

	// - reverses value of source for the rule to match everything except specified objects
	// If true, the content of the source elements is reversed
	// +kubebuilder:validation:Optional
	ExcludeSource *bool `json:"excludeSource,omitempty" tf:"exclude_source,omitempty"`

	// Whether the rule traffic is logged
	// +kubebuilder:validation:Optional
	Logged *bool `json:"logged,omitempty" tf:"logged,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Firewall Rule name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Packet type of the rule (any, ipv4, ipv6)
	// +kubebuilder:validation:Optional
	PacketType *string `json:"packetType,omitempty" tf:"packet_type,omitempty"`

	// A set of source objects. See below for source or destination objects
	// Leaving it empty matches any (all)
	// List of source traffic for this rule. An empty value means 'any'
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SourceInitParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the source entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the source entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the source entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SourceObservation struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the source entity
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the source entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the source entity
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SourceParameters struct {

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Name of the source entity
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// is the type of the object. One of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address.
	// Note that the case of the type identifiers are relevant. Using IpSet instead of IPSet results in an error.
	// Also note that Ipv4Address allows any of:
	// Type of the source entity (one of Network, Edge, VirtualMachine, IpSet, VDC, Ipv4Address)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// - When using a named object (such a VM or a network), this field will have the object ID. For a literal
	// object, such as an IP or IP range, this will be the text of the IP reference.
	// Value of the source entity
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// NsxvDistributedFirewallSpec defines the desired state of NsxvDistributedFirewall
type NsxvDistributedFirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxvDistributedFirewallParameters `json:"forProvider"`
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
	InitProvider NsxvDistributedFirewallInitParameters `json:"initProvider,omitempty"`
}

// NsxvDistributedFirewallStatus defines the observed state of NsxvDistributedFirewall.
type NsxvDistributedFirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxvDistributedFirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxvDistributedFirewall is the Schema for the NsxvDistributedFirewalls API. The NSX-V Distributed Firewall allows user to segment organization virtual data center entities, such as virtual machines, edges, networks, based on several attributes
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxvDistributedFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vdcId) || (has(self.initProvider) && has(self.initProvider.vdcId))",message="spec.forProvider.vdcId is a required parameter"
	Spec   NsxvDistributedFirewallSpec   `json:"spec"`
	Status NsxvDistributedFirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxvDistributedFirewallList contains a list of NsxvDistributedFirewalls
type NsxvDistributedFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxvDistributedFirewall `json:"items"`
}

// Repository type metadata.
var (
	NsxvDistributedFirewall_Kind             = "NsxvDistributedFirewall"
	NsxvDistributedFirewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxvDistributedFirewall_Kind}.String()
	NsxvDistributedFirewall_KindAPIVersion   = NsxvDistributedFirewall_Kind + "." + CRDGroupVersion.String()
	NsxvDistributedFirewall_GroupVersionKind = CRDGroupVersion.WithKind(NsxvDistributedFirewall_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxvDistributedFirewall{}, &NsxvDistributedFirewallList{})
}
