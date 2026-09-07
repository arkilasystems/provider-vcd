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

	apiTokenNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/api_token"
	catalogNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/catalog"
	certificateLibraryNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/certificate_library"
	clonedvAppNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/cloned_vapp"
	edgegatewayNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/edgegateway"
	externalNetworkNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/external_network"
	globalRoleNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/global_role"
	independentDiskNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/independent_disk"
	insertedMediaNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/inserted_media"
	ipSpaceNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/ip_space"
	lbNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/lb"
	networkNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/network"
	nsxtNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/nsxt"
	nsxvNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/nsxv"
	orgNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/org"
	providerVdcNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/provider_vdc"
	rdeNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/rde"
	rightsBundleNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/rights_bundle"
	roleNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/role"
	securityTagNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/security_tag"
	serviceAccountNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/service_account"
	subscribedCatalogNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/subscribed_catalog"
	uiPluginNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/ui_plugin"
	vAppNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/vapp"
	vdcGroupNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/vdc_group"
	vmNamespaced "github.com/arkilasystems/provider-vcd/config/namespaced/vm"

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
		),
		// No v1alpha1 API package: that was only ever used for StoreConfig
		// (External Secret Stores), which is removed in Crossplane v2. Drop
		// it from upjet's default base package list rather than resurrect
		// an empty package.
		ujconfig.WithBasePackages(ujconfig.BasePackages{
			APIVersion:    []string{"v1beta1"},
			Controller:    ujconfig.DefaultBasePackages.Controller,
			ControllerMap: ujconfig.DefaultBasePackages.ControllerMap,
		}))

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

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("m.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}),
		// See the matching comment in GetProvider: no v1alpha1 API package.
		ujconfig.WithBasePackages(ujconfig.BasePackages{
			APIVersion:    []string{"v1beta1"},
			Controller:    ujconfig.DefaultBasePackages.Controller,
			ControllerMap: ujconfig.DefaultBasePackages.ControllerMap,
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apiTokenNamespaced.Configure,
		catalogNamespaced.Configure,
		certificateLibraryNamespaced.Configure,
		clonedvAppNamespaced.Configure,
		edgegatewayNamespaced.Configure,
		externalNetworkNamespaced.Configure,
		globalRoleNamespaced.Configure,
		independentDiskNamespaced.Configure,
		insertedMediaNamespaced.Configure,
		ipSpaceNamespaced.Configure,
		lbNamespaced.Configure,
		networkNamespaced.Configure,
		nsxtNamespaced.Configure,
		nsxvNamespaced.Configure,
		orgNamespaced.Configure,
		providerVdcNamespaced.Configure,
		rdeNamespaced.Configure,
		rightsBundleNamespaced.Configure,
		roleNamespaced.Configure,
		securityTagNamespaced.Configure,
		serviceAccountNamespaced.Configure,
		subscribedCatalogNamespaced.Configure,
		uiPluginNamespaced.Configure,
		vAppNamespaced.Configure,
		vdcGroupNamespaced.Configure,
		vmNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
