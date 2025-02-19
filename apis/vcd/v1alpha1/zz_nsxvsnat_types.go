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

type NsxvSnatInitParameters struct {

	// Free text description.
	// NAT rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the SNAT rule.
	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the SNAT rule.
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

	// IP address, range or subnet. These addresses are the IP addresses
	// of one or more virtual machines for which you are configuring the SNAT rule so that they can send
	// traffic to the external network.
	// Original address or address range. This is the the source address for SNAT rules
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. This address is always the public
	// IP address of the gateway for which you are configuring the SNAT rule. Specifies the IP address to
	// which source addresses (the virtual machines) on outbound packets are translated to when they send
	// traffic to the external network.
	// Translated address or address range
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxvSnatObservation struct {

	// Free text description.
	// NAT rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the SNAT rule.
	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the SNAT rule.
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

	// IP address, range or subnet. These addresses are the IP addresses
	// of one or more virtual machines for which you are configuring the SNAT rule so that they can send
	// traffic to the external network.
	// Original address or address range. This is the the source address for SNAT rules
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. This address is always the public
	// IP address of the gateway for which you are configuring the SNAT rule. Specifies the IP address to
	// which source addresses (the virtual machines) on outbound packets are translated to when they send
	// traffic to the external network.
	// Translated address or address range
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxvSnatParameters struct {

	// Free text description.
	// NAT rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the edge gateway on which to apply the SNAT rule.
	// Edge gateway name in which NAT Rule is located
	// +kubebuilder:validation:Optional
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Defines if the rule is enabaled. Default true.
	// Whether the rule should be enabled. Default 'true'
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Defines if the logging for this rule is enabaled. Default false.
	// Whether logging should be enabled for this rule. Default 'false'
	// +kubebuilder:validation:Optional
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The name of the network on which to apply the SNAT rule.
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

	// IP address, range or subnet. These addresses are the IP addresses
	// of one or more virtual machines for which you are configuring the SNAT rule so that they can send
	// traffic to the external network.
	// Original address or address range. This is the the source address for SNAT rules
	// +kubebuilder:validation:Optional
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// This can be used to specify user-controlled rule tag. If not specified,
	// it will report rule ID after creation. Must be between 65537-131072.
	// Optional. Allows to set custom rule tag
	// +kubebuilder:validation:Optional
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Possible values - user, internal_high.
	// Read only. Possible values 'user', 'internal_high'
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// IP address, range or subnet. This address is always the public
	// IP address of the gateway for which you are configuring the SNAT rule. Specifies the IP address to
	// which source addresses (the virtual machines) on outbound packets are translated to when they send
	// traffic to the external network.
	// Translated address or address range
	// +kubebuilder:validation:Optional
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxvSnatSpec defines the desired state of NsxvSnat
type NsxvSnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxvSnatParameters `json:"forProvider"`
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
	InitProvider NsxvSnatInitParameters `json:"initProvider,omitempty"`
}

// NsxvSnatStatus defines the observed state of NsxvSnat.
type NsxvSnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxvSnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxvSnat is the Schema for the NsxvSnats API. Provides a VMware Cloud Director SNAT resource for advanced edge gateways (NSX-V). This can be used to create, modify, and delete source NATs to allow vApps to send external traffic.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxvSnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGateway) || (has(self.initProvider) && has(self.initProvider.edgeGateway))",message="spec.forProvider.edgeGateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkName) || (has(self.initProvider) && has(self.initProvider.networkName))",message="spec.forProvider.networkName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkType) || (has(self.initProvider) && has(self.initProvider.networkType))",message="spec.forProvider.networkType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originalAddress) || (has(self.initProvider) && has(self.initProvider.originalAddress))",message="spec.forProvider.originalAddress is a required parameter"
	Spec   NsxvSnatSpec   `json:"spec"`
	Status NsxvSnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxvSnatList contains a list of NsxvSnats
type NsxvSnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxvSnat `json:"items"`
}

// Repository type metadata.
var (
	NsxvSnat_Kind             = "NsxvSnat"
	NsxvSnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxvSnat_Kind}.String()
	NsxvSnat_KindAPIVersion   = NsxvSnat_Kind + "." + CRDGroupVersion.String()
	NsxvSnat_GroupVersionKind = CRDGroupVersion.WithKind(NsxvSnat_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxvSnat{}, &NsxvSnatList{})
}
