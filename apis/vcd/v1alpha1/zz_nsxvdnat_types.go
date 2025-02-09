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

type NsxvDnatInitParameters struct {

	// Free text description.
	// NAT rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the DNAT rule.
	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Only when protocol is set to icmp. One of any,
	// address-mask-request, address-mask-reply, destination-unreachable, echo-request,
	// echo-reply, parameter-problem, redirect, router-advertisement, router-solicitation,
	// source-quench, time-exceeded, timestamp-request, timestamp-reply. Default any
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`, `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`, `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the DNAT rule.
	// Org or external network name
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Type of the network on which to apply the DNAT rule. Possible values
	// org or ext.
	// Network type. One of 'ext', 'org'
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// IP address, range or subnet. This address must be the public IP
	// address of the edge gateway for which you are configuring the DNAT rule. In the packet being
	// inspected, this IP address or range would be those that appear as the destination IP address of the
	// packet. These packet destination addresses are the ones translated by this DNAT rule.
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// Select the port or port range that the incoming traffic uses on the
	// edge gateway to connect to the internal network on which the virtual machines are connected. This
	// selection is not available when the Protocol is set to icmp or any. Default any.
	// Original port. This is the destination port for DNAT rules
	OriginalPort *string `json:"originalPort,omitempty" tf:"original_port,omitempty"`

	// Select the protocol to which the rule applies. One of tcp, udp,
	// icmp, any. Default any
	// protocols, select Any.
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. IP addresses to which destination
	// addresses on inbound packets will be translated. These addresses are the IP addresses of the one or
	// more virtual machines for which you are configuring DNAT so that they can receive traffic from the
	// external network.
	// Translated address or address range
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// Select the port or port range that inbound traffic is connecting
	// to on the virtual machines on the internal network. These ports are the ones into which the DNAT
	// rule is translating for the packets inbound to the virtual machines.
	// Translated port
	TranslatedPort *string `json:"translatedPort,omitempty" tf:"translated_port,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxvDnatObservation struct {

	// Free text description.
	// NAT rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the DNAT rule.
	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Only when protocol is set to icmp. One of any,
	// address-mask-request, address-mask-reply, destination-unreachable, echo-request,
	// echo-reply, parameter-problem, redirect, router-advertisement, router-solicitation,
	// source-quench, time-exceeded, timestamp-request, timestamp-reply. Default any
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`, `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`, `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the DNAT rule.
	// Org or external network name
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Type of the network on which to apply the DNAT rule. Possible values
	// org or ext.
	// Network type. One of 'ext', 'org'
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// IP address, range or subnet. This address must be the public IP
	// address of the edge gateway for which you are configuring the DNAT rule. In the packet being
	// inspected, this IP address or range would be those that appear as the destination IP address of the
	// packet. These packet destination addresses are the ones translated by this DNAT rule.
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// Select the port or port range that the incoming traffic uses on the
	// edge gateway to connect to the internal network on which the virtual machines are connected. This
	// selection is not available when the Protocol is set to icmp or any. Default any.
	// Original port. This is the destination port for DNAT rules
	OriginalPort *string `json:"originalPort,omitempty" tf:"original_port,omitempty"`

	// Select the protocol to which the rule applies. One of tcp, udp,
	// icmp, any. Default any
	// protocols, select Any.
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. IP addresses to which destination
	// addresses on inbound packets will be translated. These addresses are the IP addresses of the one or
	// more virtual machines for which you are configuring DNAT so that they can receive traffic from the
	// external network.
	// Translated address or address range
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// Select the port or port range that inbound traffic is connecting
	// to on the virtual machines on the internal network. These ports are the ones into which the DNAT
	// rule is translating for the packets inbound to the virtual machines.
	// Translated port
	TranslatedPort *string `json:"translatedPort,omitempty" tf:"translated_port,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxvDnatParameters struct {

	// Free text description.
	// NAT rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the DNAT rule.
	// Edge gateway name in which NAT Rule is located
	// +kubebuilder:validation:Optional
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Only when protocol is set to icmp. One of any,
	// address-mask-request, address-mask-reply, destination-unreachable, echo-request,
	// echo-reply, parameter-problem, redirect, router-advertisement, router-solicitation,
	// source-quench, time-exceeded, timestamp-request, timestamp-reply. Default any
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`, `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`, `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	// +kubebuilder:validation:Optional
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	// +kubebuilder:validation:Optional
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the DNAT rule.
	// Org or external network name
	// +kubebuilder:validation:Optional
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Type of the network on which to apply the DNAT rule. Possible values
	// org or ext.
	// Network type. One of 'ext', 'org'
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// IP address, range or subnet. This address must be the public IP
	// address of the edge gateway for which you are configuring the DNAT rule. In the packet being
	// inspected, this IP address or range would be those that appear as the destination IP address of the
	// packet. These packet destination addresses are the ones translated by this DNAT rule.
	// Original address or address range. This is the the destination address for DNAT rules.
	// +kubebuilder:validation:Optional
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// Select the port or port range that the incoming traffic uses on the
	// edge gateway to connect to the internal network on which the virtual machines are connected. This
	// selection is not available when the Protocol is set to icmp or any. Default any.
	// Original port. This is the destination port for DNAT rules
	// +kubebuilder:validation:Optional
	OriginalPort *string `json:"originalPort,omitempty" tf:"original_port,omitempty"`

	// Select the protocol to which the rule applies. One of tcp, udp,
	// icmp, any. Default any
	// protocols, select Any.
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	// +kubebuilder:validation:Optional
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. IP addresses to which destination
	// addresses on inbound packets will be translated. These addresses are the IP addresses of the one or
	// more virtual machines for which you are configuring DNAT so that they can receive traffic from the
	// external network.
	// Translated address or address range
	// +kubebuilder:validation:Optional
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// Select the port or port range that inbound traffic is connecting
	// to on the virtual machines on the internal network. These ports are the ones into which the DNAT
	// rule is translating for the packets inbound to the virtual machines.
	// Translated port
	// +kubebuilder:validation:Optional
	TranslatedPort *string `json:"translatedPort,omitempty" tf:"translated_port,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxvDnatSpec defines the desired state of NsxvDnat
type NsxvDnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxvDnatParameters `json:"forProvider"`
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
	InitProvider NsxvDnatInitParameters `json:"initProvider,omitempty"`
}

// NsxvDnatStatus defines the observed state of NsxvDnat.
type NsxvDnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxvDnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxvDnat is the Schema for the NsxvDnats API. Provides a VMware Cloud Director DNAT resource for advanced edge gateways (NSX-V). This can be used to create, modify, and delete destination NATs to map an external IP/port to an internal IP/port.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxvDnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGateway) || (has(self.initProvider) && has(self.initProvider.edgeGateway))",message="spec.forProvider.edgeGateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkName) || (has(self.initProvider) && has(self.initProvider.networkName))",message="spec.forProvider.networkName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkType) || (has(self.initProvider) && has(self.initProvider.networkType))",message="spec.forProvider.networkType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originalAddress) || (has(self.initProvider) && has(self.initProvider.originalAddress))",message="spec.forProvider.originalAddress is a required parameter"
	Spec   NsxvDnatSpec   `json:"spec"`
	Status NsxvDnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxvDnatList contains a list of NsxvDnats
type NsxvDnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxvDnat `json:"items"`
}

// Repository type metadata.
var (
	NsxvDnat_Kind             = "NsxvDnat"
	NsxvDnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxvDnat_Kind}.String()
	NsxvDnat_KindAPIVersion   = NsxvDnat_Kind + "." + CRDGroupVersion.String()
	NsxvDnat_GroupVersionKind = CRDGroupVersion.WithKind(NsxvDnat_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxvDnat{}, &NsxvDnatList{})
}
