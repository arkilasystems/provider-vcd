/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	apiToken "github.com/arkilasystems/provider-vcd/config/cluster/api_token"
	catalog "github.com/arkilasystems/provider-vcd/config/cluster/catalog"
	certificateLibrary "github.com/arkilasystems/provider-vcd/config/cluster/certificate_library"
	clonedvApp "github.com/arkilasystems/provider-vcd/config/cluster/cloned_vapp"
	edgegateway "github.com/arkilasystems/provider-vcd/config/cluster/edgegateway"
	externalNetwork "github.com/arkilasystems/provider-vcd/config/cluster/external_network"
	globalRole "github.com/arkilasystems/provider-vcd/config/cluster/global_role"
	independentDisk "github.com/arkilasystems/provider-vcd/config/cluster/independent_disk"
	insertedMedia "github.com/arkilasystems/provider-vcd/config/cluster/inserted_media"
	ipSpace "github.com/arkilasystems/provider-vcd/config/cluster/ip_space"
	lb "github.com/arkilasystems/provider-vcd/config/cluster/lb"
	network "github.com/arkilasystems/provider-vcd/config/cluster/network"
	nsxt "github.com/arkilasystems/provider-vcd/config/cluster/nsxt"
	nsxv "github.com/arkilasystems/provider-vcd/config/cluster/nsxv"
	org "github.com/arkilasystems/provider-vcd/config/cluster/org"
	providerVdc "github.com/arkilasystems/provider-vcd/config/cluster/provider_vdc"
	rde "github.com/arkilasystems/provider-vcd/config/cluster/rde"
	rightsBundle "github.com/arkilasystems/provider-vcd/config/cluster/rights_bundle"
	role "github.com/arkilasystems/provider-vcd/config/cluster/role"
	securityTag "github.com/arkilasystems/provider-vcd/config/cluster/security_tag"
	serviceAccount "github.com/arkilasystems/provider-vcd/config/cluster/service_account"
	subscribedCatalog "github.com/arkilasystems/provider-vcd/config/cluster/subscribed_catalog"
	uiPlugin "github.com/arkilasystems/provider-vcd/config/cluster/ui_plugin"
	vApp "github.com/arkilasystems/provider-vcd/config/cluster/vapp"
	vdcGroup "github.com/arkilasystems/provider-vcd/config/cluster/vdc_group"
	vm "github.com/arkilasystems/provider-vcd/config/cluster/vm"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "vcd"
	modulePath     = "github.com/arkilasystems/provider-vcd"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apiToken.Configure,
		catalog.Configure,
		certificateLibrary.Configure,
		clonedvApp.Configure,
		edgegateway.Configure,
		externalNetwork.Configure,
		globalRole.Configure,
		independentDisk.Configure,
		insertedMedia.Configure,
		ipSpace.Configure,
		lb.Configure,
		network.Configure,
		nsxt.Configure,
		nsxv.Configure,
		org.Configure,
		providerVdc.Configure,
		rde.Configure,
		rightsBundle.Configure,
		role.Configure,
		securityTag.Configure,
		serviceAccount.Configure,
		subscribedCatalog.Configure,
		uiPlugin.Configure,
		vApp.Configure,
		vdcGroup.Configure,
		vm.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
