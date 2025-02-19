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

type HealthMonitorInitParameters struct {

	// Type of persistence profile. One of:
	// Type of health monitor. One of `HTTP`, `HTTPS`, `TCP`, `UDP`, `PING`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthMonitorObservation struct {

	// A name for ALB Pool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A boolean flag if the Health monitor is system defined.
	SystemDefined *bool `json:"systemDefined,omitempty" tf:"system_defined,omitempty"`

	// Type of persistence profile. One of:
	// Type of health monitor. One of `HTTP`, `HTTPS`, `TCP`, `UDP`, `PING`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthMonitorParameters struct {

	// Type of persistence profile. One of:
	// Type of health monitor. One of `HTTP`, `HTTPS`, `TCP`, `UDP`, `PING`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NsxtAlbPoolInitParameters struct {

	// Optional algorithm for choosing pool members (default LEAST_CONNECTIONS). Other options
	// contain ROUND_ROBIN, CONSISTENT_HASH (uses Source IP Address hash), FASTEST_RESPONSE, LEAST_LOAD,
	// FEWEST_SERVERS, RANDOM, FEWEST_TASKS, CORE_AFFINITY
	// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`, `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A set of CA Certificates to be used when validating certificates presented by the
	// pool members. Can be looked up using
	// vcd_library_certificate data source
	// A set of root certificate IDs to use when validating certificates presented by pool members
	// +listType=set
	CACertificateIds []*string `json:"caCertificateIds,omitempty" tf:"ca_certificate_ids,omitempty"`

	// Specifies whether to check the common name of the certificate presented by the pool
	// member
	// Boolean flag if common name check of the certificate should be enabled
	CnCheckEnabled *bool `json:"cnCheckEnabled,omitempty" tf:"cn_check_enabled,omitempty"`

	// Default Port defines destination server port used by the traffic sent to the member
	// (default 80)
	// Default Port defines destination server port used by the traffic sent to the member (default 80)
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description ALB Pool
	// Description of ALB Pool
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of domain names which will be used to verify the common names or subject alternative
	// names presented by the pool member certificates. It is performed only when common name check cn_check_enabled is
	// enabled
	// List of domain names which will be used to verify common names
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway ID in which ALB Pool should be created
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Boolean value if ALB Pool should be enabled (default true)
	// Boolean value if ALB Pool is enabled or not (default true)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// 1 (infinite)
	// Maximum time in minutes to gracefully disable pool member (default 1)
	GracefulTimeoutPeriod *float64 `json:"gracefulTimeoutPeriod,omitempty" tf:"graceful_timeout_period,omitempty"`

	// A block to define health monitor. Multiple can be used. See Health
	// monitor and example for usage details.
	HealthMonitor []HealthMonitorInitParameters `json:"healthMonitor,omitempty" tf:"health_monitor,omitempty"`

	// A block to define pool members. Multiple can be used. See
	// Member and example for usage details. Note only one of member,
	// member_group_id can be specified.
	// ALB Pool Members
	Member []NsxtAlbPoolMemberInitParameters `json:"member,omitempty" tf:"member,omitempty"`

	// A reference to NSX-T IP Set (vcd_nsxt_ip_set).
	// Note only one of member, member_group_id can be specified.
	// ID of Firewall Group to use for Pool Membership (VCD 10.4.1+)
	MemberGroupID *string `json:"memberGroupId,omitempty" tf:"member_group_id,omitempty"`

	// A name for ALB Pool
	// Name of ALB Pool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// defines if client traffic should be used to check if pool member is up or down
	// (default true)
	// Monitors if the traffic is accepted by node (default true)
	PassiveMonitoringEnabled *bool `json:"passiveMonitoringEnabled,omitempty" tf:"passive_monitoring_enabled,omitempty"`

	// Persistence profile will ensure that the same user sticks to the same server for a
	// desired duration of time. If the persistence profile is unmanaged by Cloud Director, updates that leave the values
	// unchanged will continue to use the same unmanaged profile. Any changes made to the persistence profile will cause
	// Cloud Director to switch the pool to a profile managed by Cloud Director. See Persistence
	// profile and example for usage details.
	PersistenceProfile []PersistenceProfileInitParameters `json:"persistenceProfile,omitempty" tf:"persistence_profile,omitempty"`

	// Will be turned on automatically when CA certificates are used
	// Enables SSL - Must be on when CA certificates are used
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtAlbPoolMemberInitParameters struct {

	// Boolean value if ALB Pool should be enabled (default true)
	// Defines if pool member is accepts traffic (default 'true')
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// IP address of pool member.
	// IP address of pool member
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port for receiving traffic - overrides the root value default_port for individual members
	// Member port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Ratio of selecting eligible servers in the pool (default 1)
	// Ratio of selecting eligible servers in the pool
	Ratio *float64 `json:"ratio,omitempty" tf:"ratio,omitempty"`
}

type NsxtAlbPoolMemberObservation struct {

	// human-readable member health description.
	// Detailed health message
	DetailedHealthMessage *string `json:"detailedHealthMessage,omitempty" tf:"detailed_health_message,omitempty"`

	// Boolean value if ALB Pool should be enabled (default true)
	// Defines if pool member is accepts traffic (default 'true')
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// one of UP, DOWN, DISABLED.
	// Health status
	HealthStatus *string `json:"healthStatus,omitempty" tf:"health_status,omitempty"`

	// IP address of pool member.
	// IP address of pool member
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A set of health monitors that marked the member as DOWN
	// Marked down by provides a set of health monitors that marked the service down
	// +listType=set
	MarkedDownBy []*string `json:"markedDownBy,omitempty" tf:"marked_down_by,omitempty"`

	// Port for receiving traffic - overrides the root value default_port for individual members
	// Member port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Ratio of selecting eligible servers in the pool (default 1)
	// Ratio of selecting eligible servers in the pool
	Ratio *float64 `json:"ratio,omitempty" tf:"ratio,omitempty"`
}

type NsxtAlbPoolMemberParameters struct {

	// Boolean value if ALB Pool should be enabled (default true)
	// Defines if pool member is accepts traffic (default 'true')
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// IP address of pool member.
	// IP address of pool member
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// Port for receiving traffic - overrides the root value default_port for individual members
	// Member port
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Ratio of selecting eligible servers in the pool (default 1)
	// Ratio of selecting eligible servers in the pool
	// +kubebuilder:validation:Optional
	Ratio *float64 `json:"ratio,omitempty" tf:"ratio,omitempty"`
}

type NsxtAlbPoolObservation struct {

	// Optional algorithm for choosing pool members (default LEAST_CONNECTIONS). Other options
	// contain ROUND_ROBIN, CONSISTENT_HASH (uses Source IP Address hash), FASTEST_RESPONSE, LEAST_LOAD,
	// FEWEST_SERVERS, RANDOM, FEWEST_TASKS, CORE_AFFINITY
	// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`, `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A set of associated Virtual Service IDs
	// IDs of associated virtual services
	// +listType=set
	AssociatedVirtualServiceIds []*string `json:"associatedVirtualServiceIds,omitempty" tf:"associated_virtual_service_ids,omitempty"`

	// A set of associated Virtual Service names
	// Names of associated virtual services
	// +listType=set
	AssociatedVirtualServices []*string `json:"associatedVirtualServices,omitempty" tf:"associated_virtual_services,omitempty"`

	// A set of CA Certificates to be used when validating certificates presented by the
	// pool members. Can be looked up using
	// vcd_library_certificate data source
	// A set of root certificate IDs to use when validating certificates presented by pool members
	// +listType=set
	CACertificateIds []*string `json:"caCertificateIds,omitempty" tf:"ca_certificate_ids,omitempty"`

	// Specifies whether to check the common name of the certificate presented by the pool
	// member
	// Boolean flag if common name check of the certificate should be enabled
	CnCheckEnabled *bool `json:"cnCheckEnabled,omitempty" tf:"cn_check_enabled,omitempty"`

	// Default Port defines destination server port used by the traffic sent to the member
	// (default 80)
	// Default Port defines destination server port used by the traffic sent to the member (default 80)
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description ALB Pool
	// Description of ALB Pool
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of domain names which will be used to verify the common names or subject alternative
	// names presented by the pool member certificates. It is performed only when common name check cn_check_enabled is
	// enabled
	// List of domain names which will be used to verify common names
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway ID in which ALB Pool should be created
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Boolean value if ALB Pool should be enabled (default true)
	// Boolean value if ALB Pool is enabled or not (default true)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of enabled members defined in the Pool
	// Number of enabled members in the pool
	EnabledMemberCount *float64 `json:"enabledMemberCount,omitempty" tf:"enabled_member_count,omitempty"`

	// 1 (infinite)
	// Maximum time in minutes to gracefully disable pool member (default 1)
	GracefulTimeoutPeriod *float64 `json:"gracefulTimeoutPeriod,omitempty" tf:"graceful_timeout_period,omitempty"`

	// Health message of ALB Pool
	// Health message
	HealthMessage *string `json:"healthMessage,omitempty" tf:"health_message,omitempty"`

	// A block to define health monitor. Multiple can be used. See Health
	// monitor and example for usage details.
	HealthMonitor []HealthMonitorObservation `json:"healthMonitor,omitempty" tf:"health_monitor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block to define pool members. Multiple can be used. See
	// Member and example for usage details. Note only one of member,
	// member_group_id can be specified.
	// ALB Pool Members
	Member []NsxtAlbPoolMemberObservation `json:"member,omitempty" tf:"member,omitempty"`

	// Total number of members defined in the Pool
	// Number of members in the pool
	MemberCount *float64 `json:"memberCount,omitempty" tf:"member_count,omitempty"`

	// A reference to NSX-T IP Set (vcd_nsxt_ip_set).
	// Note only one of member, member_group_id can be specified.
	// ID of Firewall Group to use for Pool Membership (VCD 10.4.1+)
	MemberGroupID *string `json:"memberGroupId,omitempty" tf:"member_group_id,omitempty"`

	// A name for ALB Pool
	// Name of ALB Pool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// defines if client traffic should be used to check if pool member is up or down
	// (default true)
	// Monitors if the traffic is accepted by node (default true)
	PassiveMonitoringEnabled *bool `json:"passiveMonitoringEnabled,omitempty" tf:"passive_monitoring_enabled,omitempty"`

	// Persistence profile will ensure that the same user sticks to the same server for a
	// desired duration of time. If the persistence profile is unmanaged by Cloud Director, updates that leave the values
	// unchanged will continue to use the same unmanaged profile. Any changes made to the persistence profile will cause
	// Cloud Director to switch the pool to a profile managed by Cloud Director. See Persistence
	// profile and example for usage details.
	PersistenceProfile []PersistenceProfileObservation `json:"persistenceProfile,omitempty" tf:"persistence_profile,omitempty"`

	// Will be turned on automatically when CA certificates are used
	// Enables SSL - Must be on when CA certificates are used
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// Number of members defined in the Pool that are accepting traffic
	// Number of members in the pool serving the traffic
	UpMemberCount *float64 `json:"upMemberCount,omitempty" tf:"up_member_count,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtAlbPoolParameters struct {

	// Optional algorithm for choosing pool members (default LEAST_CONNECTIONS). Other options
	// contain ROUND_ROBIN, CONSISTENT_HASH (uses Source IP Address hash), FASTEST_RESPONSE, LEAST_LOAD,
	// FEWEST_SERVERS, RANDOM, FEWEST_TASKS, CORE_AFFINITY
	// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`, `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A set of CA Certificates to be used when validating certificates presented by the
	// pool members. Can be looked up using
	// vcd_library_certificate data source
	// A set of root certificate IDs to use when validating certificates presented by pool members
	// +kubebuilder:validation:Optional
	// +listType=set
	CACertificateIds []*string `json:"caCertificateIds,omitempty" tf:"ca_certificate_ids,omitempty"`

	// Specifies whether to check the common name of the certificate presented by the pool
	// member
	// Boolean flag if common name check of the certificate should be enabled
	// +kubebuilder:validation:Optional
	CnCheckEnabled *bool `json:"cnCheckEnabled,omitempty" tf:"cn_check_enabled,omitempty"`

	// Default Port defines destination server port used by the traffic sent to the member
	// (default 80)
	// Default Port defines destination server port used by the traffic sent to the member (default 80)
	// +kubebuilder:validation:Optional
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description ALB Pool
	// Description of ALB Pool
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of domain names which will be used to verify the common names or subject alternative
	// names presented by the pool member certificates. It is performed only when common name check cn_check_enabled is
	// enabled
	// List of domain names which will be used to verify common names
	// +kubebuilder:validation:Optional
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source
	// Edge gateway ID in which ALB Pool should be created
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Boolean value if ALB Pool should be enabled (default true)
	// Boolean value if ALB Pool is enabled or not (default true)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// 1 (infinite)
	// Maximum time in minutes to gracefully disable pool member (default 1)
	// +kubebuilder:validation:Optional
	GracefulTimeoutPeriod *float64 `json:"gracefulTimeoutPeriod,omitempty" tf:"graceful_timeout_period,omitempty"`

	// A block to define health monitor. Multiple can be used. See Health
	// monitor and example for usage details.
	// +kubebuilder:validation:Optional
	HealthMonitor []HealthMonitorParameters `json:"healthMonitor,omitempty" tf:"health_monitor,omitempty"`

	// A block to define pool members. Multiple can be used. See
	// Member and example for usage details. Note only one of member,
	// member_group_id can be specified.
	// ALB Pool Members
	// +kubebuilder:validation:Optional
	Member []NsxtAlbPoolMemberParameters `json:"member,omitempty" tf:"member,omitempty"`

	// A reference to NSX-T IP Set (vcd_nsxt_ip_set).
	// Note only one of member, member_group_id can be specified.
	// ID of Firewall Group to use for Pool Membership (VCD 10.4.1+)
	// +kubebuilder:validation:Optional
	MemberGroupID *string `json:"memberGroupId,omitempty" tf:"member_group_id,omitempty"`

	// A name for ALB Pool
	// Name of ALB Pool
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// defines if client traffic should be used to check if pool member is up or down
	// (default true)
	// Monitors if the traffic is accepted by node (default true)
	// +kubebuilder:validation:Optional
	PassiveMonitoringEnabled *bool `json:"passiveMonitoringEnabled,omitempty" tf:"passive_monitoring_enabled,omitempty"`

	// Persistence profile will ensure that the same user sticks to the same server for a
	// desired duration of time. If the persistence profile is unmanaged by Cloud Director, updates that leave the values
	// unchanged will continue to use the same unmanaged profile. Any changes made to the persistence profile will cause
	// Cloud Director to switch the pool to a profile managed by Cloud Director. See Persistence
	// profile and example for usage details.
	// +kubebuilder:validation:Optional
	PersistenceProfile []PersistenceProfileParameters `json:"persistenceProfile,omitempty" tf:"persistence_profile,omitempty"`

	// Will be turned on automatically when CA certificates are used
	// Enables SSL - Must be on when CA certificates are used
	// +kubebuilder:validation:Optional
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type PersistenceProfileInitParameters struct {

	// Type of persistence profile. One of:
	// Type of persistence strategy. One of `CLIENT_IP`, `HTTP_COOKIE`, `CUSTOM_HTTP_HEADER`, `APP_COOKIE`, `TLS`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// is required for some type values: HTTP_COOKIE, CUSTOM_HTTP_HEADER, APP_COOKIE
	// Value of attribute based on persistence type
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PersistenceProfileObservation struct {

	// A name for ALB Pool
	// System generated name of persistence profile
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of persistence profile. One of:
	// Type of persistence strategy. One of `CLIENT_IP`, `HTTP_COOKIE`, `CUSTOM_HTTP_HEADER`, `APP_COOKIE`, `TLS`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// is required for some type values: HTTP_COOKIE, CUSTOM_HTTP_HEADER, APP_COOKIE
	// Value of attribute based on persistence type
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PersistenceProfileParameters struct {

	// Type of persistence profile. One of:
	// Type of persistence strategy. One of `CLIENT_IP`, `HTTP_COOKIE`, `CUSTOM_HTTP_HEADER`, `APP_COOKIE`, `TLS`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// is required for some type values: HTTP_COOKIE, CUSTOM_HTTP_HEADER, APP_COOKIE
	// Value of attribute based on persistence type
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// NsxtAlbPoolSpec defines the desired state of NsxtAlbPool
type NsxtAlbPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtAlbPoolParameters `json:"forProvider"`
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
	InitProvider NsxtAlbPoolInitParameters `json:"initProvider,omitempty"`
}

// NsxtAlbPoolStatus defines the observed state of NsxtAlbPool.
type NsxtAlbPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtAlbPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NsxtAlbPool is the Schema for the NsxtAlbPools API. Provides a resource to manage ALB Pools for particular NSX-T Edge Gateway. Pools maintain the list of servers assigned to them and perform health monitoring, load balancing, persistence. A pool may only be used or referenced by only one virtual service at a time.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtAlbPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.edgeGatewayId) || (has(self.initProvider) && has(self.initProvider.edgeGatewayId))",message="spec.forProvider.edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NsxtAlbPoolSpec   `json:"spec"`
	Status NsxtAlbPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbPoolList contains a list of NsxtAlbPools
type NsxtAlbPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtAlbPool `json:"items"`
}

// Repository type metadata.
var (
	NsxtAlbPool_Kind             = "NsxtAlbPool"
	NsxtAlbPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtAlbPool_Kind}.String()
	NsxtAlbPool_KindAPIVersion   = NsxtAlbPool_Kind + "." + CRDGroupVersion.String()
	NsxtAlbPool_GroupVersionKind = CRDGroupVersion.WithKind(NsxtAlbPool_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtAlbPool{}, &NsxtAlbPoolList{})
}
