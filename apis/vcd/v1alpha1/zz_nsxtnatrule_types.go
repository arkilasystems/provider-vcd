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

type NsxtNatRuleInitParameters struct {

	// Application Port Profile to which to apply the rule. The
	// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
	// gateway to connect to the internal network.  Can be looked up using vcd_nsxt_app_port_profile
	// data source or created using vcd_nsxt_app_port_profile resource
	// Application Port Profile to apply for this rule
	AppPortProfileID *string `json:"appPortProfileId,omitempty" tf:"app_port_profile_id,omitempty"`

	// An optional description of the NAT rule
	// Description of NAT rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// For DNAT only. This represents the external port number or port range when doing
	// DNAT port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
	// for the given IPs selected will be translated.
	// For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
	DnatExternalPort *string `json:"dnatExternalPort,omitempty" tf:"dnat_external_port,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway name in which NAT Rule is located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Enables or disables NAT rule (default true)
	// Enables or disables this rule
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The external address for the NAT Rule. This must be supplied as a single IP or Network
	// CIDR. For a DNAT rule, this is the external facing IP Address for incoming traffic. For an SNAT rule, this is the
	// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
	// Edge Gateway. For a REFLEXIVE rule, these are the external facing IPs.
	// IP address or CIDR of external network
	ExternalAddress *string `json:"externalAddress,omitempty" tf:"external_address,omitempty"`

	// You can set a firewall match rule to determine how
	// firewall is applied during NAT. One of MATCH_INTERNAL_ADDRESS, MATCH_EXTERNAL_ADDRESS,
	// BYPASS
	// VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of 'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
	FirewallMatch *string `json:"firewallMatch,omitempty" tf:"firewall_match,omitempty"`

	// The internal address for the NAT Rule. This must be supplied as a single IP or
	// Network CIDR. For a DNAT rule, this is the internal IP address for incoming traffic. For an SNAT rule, this is the
	// internal IP Address for outgoing traffic. For a REFLEXIVE rule, these are the internal IPs.
	// These IPs are typically the Private IPs that are allocated to workloads.
	// IP address or CIDR of the virtual machines for which you are configuring NAT
	InternalAddress *string `json:"internalAddress,omitempty" tf:"internal_address,omitempty"`

	// Enable to have the address translation performed by this rule logged
	// (default false). Note User might lack rights (Organization Administrator role by default
	// is missing Gateway -> Configure System Logging right) to enable logging, but API does not
	// return error and it is not possible to validate it.
	// Enable logging when this rule is applied
	Logging *bool `json:"logging,omitempty" tf:"logging,omitempty"`

	// A name for NAT rule
	// Name of NAT rule
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// if an address has multiple NAT rules, you can assign these
	// rules different priorities to determine the order in which they are applied. A lower value means a
	// higher priority for this rule.
	// VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a higher precedence for this rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One of DNAT, NO_DNAT, SNAT, NO_SNAT, REFLEXIVE
	// Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// For SNAT only. The destination addresses to match in the SNAT Rule. This
	// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
	// destination network.
	// For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain or an IP address range in CIDR format.
	SnatDestinationAddress *string `json:"snatDestinationAddress,omitempty" tf:"snat_destination_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtNatRuleObservation struct {

	// Application Port Profile to which to apply the rule. The
	// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
	// gateway to connect to the internal network.  Can be looked up using vcd_nsxt_app_port_profile
	// data source or created using vcd_nsxt_app_port_profile resource
	// Application Port Profile to apply for this rule
	AppPortProfileID *string `json:"appPortProfileId,omitempty" tf:"app_port_profile_id,omitempty"`

	// An optional description of the NAT rule
	// Description of NAT rule
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// For DNAT only. This represents the external port number or port range when doing
	// DNAT port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
	// for the given IPs selected will be translated.
	// For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
	DnatExternalPort *string `json:"dnatExternalPort,omitempty" tf:"dnat_external_port,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway name in which NAT Rule is located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Enables or disables NAT rule (default true)
	// Enables or disables this rule
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The external address for the NAT Rule. This must be supplied as a single IP or Network
	// CIDR. For a DNAT rule, this is the external facing IP Address for incoming traffic. For an SNAT rule, this is the
	// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
	// Edge Gateway. For a REFLEXIVE rule, these are the external facing IPs.
	// IP address or CIDR of external network
	ExternalAddress *string `json:"externalAddress,omitempty" tf:"external_address,omitempty"`

	// You can set a firewall match rule to determine how
	// firewall is applied during NAT. One of MATCH_INTERNAL_ADDRESS, MATCH_EXTERNAL_ADDRESS,
	// BYPASS
	// VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of 'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
	FirewallMatch *string `json:"firewallMatch,omitempty" tf:"firewall_match,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The internal address for the NAT Rule. This must be supplied as a single IP or
	// Network CIDR. For a DNAT rule, this is the internal IP address for incoming traffic. For an SNAT rule, this is the
	// internal IP Address for outgoing traffic. For a REFLEXIVE rule, these are the internal IPs.
	// These IPs are typically the Private IPs that are allocated to workloads.
	// IP address or CIDR of the virtual machines for which you are configuring NAT
	InternalAddress *string `json:"internalAddress,omitempty" tf:"internal_address,omitempty"`

	// Enable to have the address translation performed by this rule logged
	// (default false). Note User might lack rights (Organization Administrator role by default
	// is missing Gateway -> Configure System Logging right) to enable logging, but API does not
	// return error and it is not possible to validate it.
	// Enable logging when this rule is applied
	Logging *bool `json:"logging,omitempty" tf:"logging,omitempty"`

	// A name for NAT rule
	// Name of NAT rule
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// if an address has multiple NAT rules, you can assign these
	// rules different priorities to determine the order in which they are applied. A lower value means a
	// higher priority for this rule.
	// VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a higher precedence for this rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One of DNAT, NO_DNAT, SNAT, NO_SNAT, REFLEXIVE
	// Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// For SNAT only. The destination addresses to match in the SNAT Rule. This
	// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
	// destination network.
	// For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain or an IP address range in CIDR format.
	SnatDestinationAddress *string `json:"snatDestinationAddress,omitempty" tf:"snat_destination_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtNatRuleParameters struct {

	// Application Port Profile to which to apply the rule. The
	// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
	// gateway to connect to the internal network.  Can be looked up using vcd_nsxt_app_port_profile
	// data source or created using vcd_nsxt_app_port_profile resource
	// Application Port Profile to apply for this rule
	// +kubebuilder:validation:Optional
	AppPortProfileID *string `json:"appPortProfileId,omitempty" tf:"app_port_profile_id,omitempty"`

	// An optional description of the NAT rule
	// Description of NAT rule
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// For DNAT only. This represents the external port number or port range when doing
	// DNAT port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
	// for the given IPs selected will be translated.
	// For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
	// +kubebuilder:validation:Optional
	DnatExternalPort *string `json:"dnatExternalPort,omitempty" tf:"dnat_external_port,omitempty"`

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway name in which NAT Rule is located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Enables or disables NAT rule (default true)
	// Enables or disables this rule
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The external address for the NAT Rule. This must be supplied as a single IP or Network
	// CIDR. For a DNAT rule, this is the external facing IP Address for incoming traffic. For an SNAT rule, this is the
	// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
	// Edge Gateway. For a REFLEXIVE rule, these are the external facing IPs.
	// IP address or CIDR of external network
	// +kubebuilder:validation:Optional
	ExternalAddress *string `json:"externalAddress,omitempty" tf:"external_address,omitempty"`

	// You can set a firewall match rule to determine how
	// firewall is applied during NAT. One of MATCH_INTERNAL_ADDRESS, MATCH_EXTERNAL_ADDRESS,
	// BYPASS
	// VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of 'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
	// +kubebuilder:validation:Optional
	FirewallMatch *string `json:"firewallMatch,omitempty" tf:"firewall_match,omitempty"`

	// The internal address for the NAT Rule. This must be supplied as a single IP or
	// Network CIDR. For a DNAT rule, this is the internal IP address for incoming traffic. For an SNAT rule, this is the
	// internal IP Address for outgoing traffic. For a REFLEXIVE rule, these are the internal IPs.
	// These IPs are typically the Private IPs that are allocated to workloads.
	// IP address or CIDR of the virtual machines for which you are configuring NAT
	// +kubebuilder:validation:Optional
	InternalAddress *string `json:"internalAddress,omitempty" tf:"internal_address,omitempty"`

	// Enable to have the address translation performed by this rule logged
	// (default false). Note User might lack rights (Organization Administrator role by default
	// is missing Gateway -> Configure System Logging right) to enable logging, but API does not
	// return error and it is not possible to validate it.
	// Enable logging when this rule is applied
	// +kubebuilder:validation:Optional
	Logging *bool `json:"logging,omitempty" tf:"logging,omitempty"`

	// A name for NAT rule
	// Name of NAT rule
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// if an address has multiple NAT rules, you can assign these
	// rules different priorities to determine the order in which they are applied. A lower value means a
	// higher priority for this rule.
	// VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a higher precedence for this rule.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One of DNAT, NO_DNAT, SNAT, NO_SNAT, REFLEXIVE
	// Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// For SNAT only. The destination addresses to match in the SNAT Rule. This
	// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
	// destination network.
	// For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain or an IP address range in CIDR format.
	// +kubebuilder:validation:Optional
	SnatDestinationAddress *string `json:"snatDestinationAddress,omitempty" tf:"snat_destination_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxtNatRuleSpec defines the desired state of NsxtNatRule
type NsxtNatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtNatRuleParameters `json:"forProvider"`
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
	InitProvider NsxtNatRuleInitParameters `json:"initProvider,omitempty"`
}

// NsxtNatRuleStatus defines the observed state of NsxtNatRule.
type NsxtNatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtNatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxtNatRule is the Schema for the NsxtNatRules API. Provides a resource to manage NSX-T NAT rules. To change the source IP address from a private to a public IP address, you create a source NAT (SNAT) rule. To change the destination IP address from a public to a private IP address, you create a destination NAT (DNAT) rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtNatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGatewayId) || (has(self.initProvider) && has(self.initProvider.edgeGatewayId))",message="spec.forProvider.edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleType) || (has(self.initProvider) && has(self.initProvider.ruleType))",message="spec.forProvider.ruleType is a required parameter"
	Spec   NsxtNatRuleSpec   `json:"spec"`
	Status NsxtNatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtNatRuleList contains a list of NsxtNatRules
type NsxtNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtNatRule `json:"items"`
}

// Repository type metadata.
var (
	NsxtNatRule_Kind             = "NsxtNatRule"
	NsxtNatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtNatRule_Kind}.String()
	NsxtNatRule_KindAPIVersion   = NsxtNatRule_Kind + "." + CRDGroupVersion.String()
	NsxtNatRule_GroupVersionKind = CRDGroupVersion.WithKind(NsxtNatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtNatRule{}, &NsxtNatRuleList{})
}
