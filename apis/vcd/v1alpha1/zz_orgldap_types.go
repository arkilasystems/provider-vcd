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

type CustomSettingsInitParameters struct {

	// Authentication method: one of SIMPLE, MD5DIGEST, NTLM
	// authentication method: one of SIMPLE, MD5DIGEST, NTLM
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// LDAP search base
	// LDAP search base
	BaseDistinguishedName *string `json:"baseDistinguishedName,omitempty" tf:"base_distinguished_name,omitempty"`

	// Type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	// type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// Group settings when ldap_mode is CUSTOM See Group Attributes below for details
	// Group settings when `ldap_mode` is CUSTOM
	GroupAttributes []GroupAttributesInitParameters `json:"groupAttributes,omitempty" tf:"group_attributes,omitempty"`

	// True if the LDAP service requires an SSL connection
	// True if the LDAP service requires an SSL connection
	IsSSL *bool `json:"isSsl,omitempty" tf:"is_ssl,omitempty"`

	// Password for the user identified by UserName. This value is never returned by GET.
	// It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
	// Password for the user identified by UserName. This value is never returned by GET. It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port number of the LDAP server (usually 389 for LDAP, 636 for LDAPS)
	// Port number for LDAP service
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The IP address or host name of the server providing the LDAP service
	// host name or IP of the LDAP server
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// User settings when ldap_mode is CUSTOM See User Attributes below for details
	// User settings when `ldap_mode` is CUSTOM
	UserAttributes []UserAttributesInitParameters `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs (for example: cn="ldap-admin", c="example", dc="com")
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CustomSettingsObservation struct {

	// Authentication method: one of SIMPLE, MD5DIGEST, NTLM
	// authentication method: one of SIMPLE, MD5DIGEST, NTLM
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// LDAP search base
	// LDAP search base
	BaseDistinguishedName *string `json:"baseDistinguishedName,omitempty" tf:"base_distinguished_name,omitempty"`

	// Type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	// type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// Group settings when ldap_mode is CUSTOM See Group Attributes below for details
	// Group settings when `ldap_mode` is CUSTOM
	GroupAttributes []GroupAttributesObservation `json:"groupAttributes,omitempty" tf:"group_attributes,omitempty"`

	// True if the LDAP service requires an SSL connection
	// True if the LDAP service requires an SSL connection
	IsSSL *bool `json:"isSsl,omitempty" tf:"is_ssl,omitempty"`

	// Port number of the LDAP server (usually 389 for LDAP, 636 for LDAPS)
	// Port number for LDAP service
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The IP address or host name of the server providing the LDAP service
	// host name or IP of the LDAP server
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// User settings when ldap_mode is CUSTOM See User Attributes below for details
	// User settings when `ldap_mode` is CUSTOM
	UserAttributes []UserAttributesObservation `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs (for example: cn="ldap-admin", c="example", dc="com")
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CustomSettingsParameters struct {

	// Authentication method: one of SIMPLE, MD5DIGEST, NTLM
	// authentication method: one of SIMPLE, MD5DIGEST, NTLM
	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod" tf:"authentication_method,omitempty"`

	// LDAP search base
	// LDAP search base
	// +kubebuilder:validation:Optional
	BaseDistinguishedName *string `json:"baseDistinguishedName,omitempty" tf:"base_distinguished_name,omitempty"`

	// Type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	// type of connector: one of OPEN_LDAP, ACTIVE_DIRECTORY
	// +kubebuilder:validation:Optional
	ConnectorType *string `json:"connectorType" tf:"connector_type,omitempty"`

	// Group settings when ldap_mode is CUSTOM See Group Attributes below for details
	// Group settings when `ldap_mode` is CUSTOM
	// +kubebuilder:validation:Optional
	GroupAttributes []GroupAttributesParameters `json:"groupAttributes" tf:"group_attributes,omitempty"`

	// True if the LDAP service requires an SSL connection
	// True if the LDAP service requires an SSL connection
	// +kubebuilder:validation:Optional
	IsSSL *bool `json:"isSsl,omitempty" tf:"is_ssl,omitempty"`

	// Password for the user identified by UserName. This value is never returned by GET.
	// It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
	// Password for the user identified by UserName. This value is never returned by GET. It is inspected on create and modify. On modify, the absence of this element indicates that the password should not be changed
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port number of the LDAP server (usually 389 for LDAP, 636 for LDAPS)
	// Port number for LDAP service
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The IP address or host name of the server providing the LDAP service
	// host name or IP of the LDAP server
	// +kubebuilder:validation:Optional
	Server *string `json:"server" tf:"server,omitempty"`

	// User settings when ldap_mode is CUSTOM See User Attributes below for details
	// User settings when `ldap_mode` is CUSTOM
	// +kubebuilder:validation:Optional
	UserAttributes []UserAttributesParameters `json:"userAttributes" tf:"user_attributes,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs (for example: cn="ldap-admin", c="example", dc="com")
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type GroupAttributesInitParameters struct {

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP group attribute used to identify a group member
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a group as a member of another group. For example, dn
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier,omitempty" tf:"group_membership_identifier,omitempty"`

	// LDAP attribute to use when getting the members of a group. For example, member
	// LDAP attribute to use when getting the members of a group. For example, member
	Membership *string `json:"membership,omitempty" tf:"membership,omitempty"`

	// LDAP attribute to use for the group name. For example, cn
	// LDAP attribute to use for the group name. For example, cn
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported groups are members. For example, group
	ObjectClass *string `json:"objectClass,omitempty" tf:"object_class,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a group. For example, objectGuid
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type GroupAttributesObservation struct {

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP group attribute used to identify a group member
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a group as a member of another group. For example, dn
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier,omitempty" tf:"group_membership_identifier,omitempty"`

	// LDAP attribute to use when getting the members of a group. For example, member
	// LDAP attribute to use when getting the members of a group. For example, member
	Membership *string `json:"membership,omitempty" tf:"membership,omitempty"`

	// LDAP attribute to use for the group name. For example, cn
	// LDAP attribute to use for the group name. For example, cn
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported groups are members. For example, group
	ObjectClass *string `json:"objectClass,omitempty" tf:"object_class,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a group. For example, objectGuid
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type GroupAttributesParameters struct {

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP group attribute used to identify a group member
	// +kubebuilder:validation:Optional
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a group as a member of another group. For example, dn
	// +kubebuilder:validation:Optional
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier" tf:"group_membership_identifier,omitempty"`

	// LDAP attribute to use when getting the members of a group. For example, member
	// LDAP attribute to use when getting the members of a group. For example, member
	// +kubebuilder:validation:Optional
	Membership *string `json:"membership" tf:"membership,omitempty"`

	// LDAP attribute to use for the group name. For example, cn
	// LDAP attribute to use for the group name. For example, cn
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported groups are members. For example, group
	// +kubebuilder:validation:Optional
	ObjectClass *string `json:"objectClass" tf:"object_class,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a group. For example, objectGuid
	// +kubebuilder:validation:Optional
	UniqueIdentifier *string `json:"uniqueIdentifier" tf:"unique_identifier,omitempty"`
}

type OrgLdapInitParameters struct {

	// LDAP server configuration. Becomes mandatory if ldap_mode is set to CUSTOM. See Custom Settings below for details
	// Custom settings when `ldap_mode` is CUSTOM
	CustomSettings []CustomSettingsInitParameters `json:"customSettings,omitempty" tf:"custom_settings,omitempty"`

	// If ldap_mode is SYSTEM, specifies an LDAP attribute=value pair to use for OU (organizational unit)
	// If ldap_mode is SYSTEM, specifies a LDAP attribute=value pair to use for OU (organizational unit)
	CustomUserOu *string `json:"customUserOu,omitempty" tf:"custom_user_ou,omitempty"`

	// One of NONE, CUSTOM, SYSTEM. Note that using NONE has the effect of removing the LDAP settings
	// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
	LdapMode *string `json:"ldapMode,omitempty" tf:"ldap_mode,omitempty"`

	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	// Organization ID
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type OrgLdapObservation struct {

	// LDAP server configuration. Becomes mandatory if ldap_mode is set to CUSTOM. See Custom Settings below for details
	// Custom settings when `ldap_mode` is CUSTOM
	CustomSettings []CustomSettingsObservation `json:"customSettings,omitempty" tf:"custom_settings,omitempty"`

	// If ldap_mode is SYSTEM, specifies an LDAP attribute=value pair to use for OU (organizational unit)
	// If ldap_mode is SYSTEM, specifies a LDAP attribute=value pair to use for OU (organizational unit)
	CustomUserOu *string `json:"customUserOu,omitempty" tf:"custom_user_ou,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One of NONE, CUSTOM, SYSTEM. Note that using NONE has the effect of removing the LDAP settings
	// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
	LdapMode *string `json:"ldapMode,omitempty" tf:"ldap_mode,omitempty"`

	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	// Organization ID
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type OrgLdapParameters struct {

	// LDAP server configuration. Becomes mandatory if ldap_mode is set to CUSTOM. See Custom Settings below for details
	// Custom settings when `ldap_mode` is CUSTOM
	// +kubebuilder:validation:Optional
	CustomSettings []CustomSettingsParameters `json:"customSettings,omitempty" tf:"custom_settings,omitempty"`

	// If ldap_mode is SYSTEM, specifies an LDAP attribute=value pair to use for OU (organizational unit)
	// If ldap_mode is SYSTEM, specifies a LDAP attribute=value pair to use for OU (organizational unit)
	// +kubebuilder:validation:Optional
	CustomUserOu *string `json:"customUserOu,omitempty" tf:"custom_user_ou,omitempty"`

	// One of NONE, CUSTOM, SYSTEM. Note that using NONE has the effect of removing the LDAP settings
	// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
	// +kubebuilder:validation:Optional
	LdapMode *string `json:"ldapMode,omitempty" tf:"ldap_mode,omitempty"`

	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	// Organization ID
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type UserAttributesInitParameters struct {

	// LDAP attribute to use for the user's full name. For example, displayName
	// LDAP attribute to use for the user's full name. For example, displayName
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// LDAP attribute to use for the user's email address. For example, mail
	// LDAP attribute to use for the user's email address. For example, mail
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// LDAP attribute to use for the user's given name. For example, givenName
	// LDAP attribute to use for the user's given name. For example, givenName
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a user as a member of a group. For example, dn
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier,omitempty" tf:"group_membership_identifier,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported users are members. For example, user or person
	ObjectClass *string `json:"objectClass,omitempty" tf:"object_class,omitempty"`

	// LDAP attribute to use for the user's surname. For example, sn
	// LDAP attribute to use for the user's surname. For example, sn
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	Telephone *string `json:"telephone,omitempty" tf:"telephone,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserAttributesObservation struct {

	// LDAP attribute to use for the user's full name. For example, displayName
	// LDAP attribute to use for the user's full name. For example, displayName
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// LDAP attribute to use for the user's email address. For example, mail
	// LDAP attribute to use for the user's email address. For example, mail
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// LDAP attribute to use for the user's given name. For example, givenName
	// LDAP attribute to use for the user's given name. For example, givenName
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a user as a member of a group. For example, dn
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier,omitempty" tf:"group_membership_identifier,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported users are members. For example, user or person
	ObjectClass *string `json:"objectClass,omitempty" tf:"object_class,omitempty"`

	// LDAP attribute to use for the user's surname. For example, sn
	// LDAP attribute to use for the user's surname. For example, sn
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	Telephone *string `json:"telephone,omitempty" tf:"telephone,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserAttributesParameters struct {

	// LDAP attribute to use for the user's full name. For example, displayName
	// LDAP attribute to use for the user's full name. For example, displayName
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// LDAP attribute to use for the user's email address. For example, mail
	// LDAP attribute to use for the user's email address. For example, mail
	// +kubebuilder:validation:Optional
	Email *string `json:"email" tf:"email,omitempty"`

	// LDAP attribute to use for the user's given name. For example, givenName
	// LDAP attribute to use for the user's given name. For example, givenName
	// +kubebuilder:validation:Optional
	GivenName *string `json:"givenName" tf:"given_name,omitempty"`

	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// LDAP attribute that returns the identifiers of all the groups of which the user is a member
	// +kubebuilder:validation:Optional
	GroupBackLinkIdentifier *string `json:"groupBackLinkIdentifier,omitempty" tf:"group_back_link_identifier,omitempty"`

	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// LDAP attribute that identifies a user as a member of a group. For example, dn
	// +kubebuilder:validation:Optional
	GroupMembershipIdentifier *string `json:"groupMembershipIdentifier" tf:"group_membership_identifier,omitempty"`

	// LDAP objectClass of which imported users are members. For example, user or person
	// LDAP objectClass of which imported users are members. For example, user or person
	// +kubebuilder:validation:Optional
	ObjectClass *string `json:"objectClass" tf:"object_class,omitempty"`

	// LDAP attribute to use for the user's surname. For example, sn
	// LDAP attribute to use for the user's surname. For example, sn
	// +kubebuilder:validation:Optional
	Surname *string `json:"surname" tf:"surname,omitempty"`

	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	// LDAP attribute to use for the user's telephone number. For example, telephoneNumber
	// +kubebuilder:validation:Optional
	Telephone *string `json:"telephone" tf:"telephone,omitempty"`

	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// LDAP attribute to use as the unique identifier for a user. For example, objectGuid
	// +kubebuilder:validation:Optional
	UniqueIdentifier *string `json:"uniqueIdentifier" tf:"unique_identifier,omitempty"`

	// Username to use when logging in to LDAP, specified using LDAP attribute=value pairs
	// (for example: cn="ldap-admin", c="example", dc="com")
	// LDAP attribute to use when looking up a user name to import. For example, userPrincipalName or samAccountName
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

// OrgLdapSpec defines the desired state of OrgLdap
type OrgLdapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrgLdapParameters `json:"forProvider"`
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
	InitProvider OrgLdapInitParameters `json:"initProvider,omitempty"`
}

// OrgLdapStatus defines the observed state of OrgLdap.
type OrgLdapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrgLdapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrgLdap is the Schema for the OrgLdaps API. Provides a VMware Cloud Director Organization LDAP resource. This can be used to create, delete, and update LDAP configuration for an organization .
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type OrgLdap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ldapMode) || (has(self.initProvider) && has(self.initProvider.ldapMode))",message="spec.forProvider.ldapMode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.orgId) || (has(self.initProvider) && has(self.initProvider.orgId))",message="spec.forProvider.orgId is a required parameter"
	Spec   OrgLdapSpec   `json:"spec"`
	Status OrgLdapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrgLdapList contains a list of OrgLdaps
type OrgLdapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgLdap `json:"items"`
}

// Repository type metadata.
var (
	OrgLdap_Kind             = "OrgLdap"
	OrgLdap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrgLdap_Kind}.String()
	OrgLdap_KindAPIVersion   = OrgLdap_Kind + "." + CRDGroupVersion.String()
	OrgLdap_GroupVersionKind = CRDGroupVersion.WithKind(OrgLdap_Kind)
)

func init() {
	SchemeBuilder.Register(&OrgLdap{}, &OrgLdapList{})
}
