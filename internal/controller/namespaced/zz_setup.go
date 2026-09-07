// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/providerconfig"
	apitoken "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/apitoken"
	catalog "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/catalog"
	catalogitem "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/catalogitem"
	catalogmedia "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/catalogmedia"
	catalogvapptemplate "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/catalogvapptemplate"
	clonedvapp "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/clonedvapp"
	edgegateway "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/edgegateway"
	edgegatewaysettings "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/edgegatewaysettings"
	edgegatewayvpn "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/edgegatewayvpn"
	externalnetwork "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/externalnetwork"
	externalnetworkv2 "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/externalnetworkv2"
	globalrole "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/globalrole"
	independentdisk "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/independentdisk"
	insertedmedia "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/insertedmedia"
	ipspace "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/ipspace"
	ipspacecustomquota "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/ipspacecustomquota"
	ipspaceipallocation "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/ipspaceipallocation"
	ipspaceuplink "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/ipspaceuplink"
	lbappprofile "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/lbappprofile"
	lbapprule "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/lbapprule"
	lbserverpool "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/lbserverpool"
	lbservicemonitor "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/lbservicemonitor"
	lbvirtualserver "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/lbvirtualserver"
	networkdirect "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/networkdirect"
	networkisolated "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/networkisolated"
	networkisolatedv2 "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/networkisolatedv2"
	networkrouted "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/networkrouted"
	networkroutedv2 "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/networkroutedv2"
	nsxtalbcloud "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbcloud"
	nsxtalbcontroller "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbcontroller"
	nsxtalbedgegatewayserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbedgegatewayserviceenginegroup"
	nsxtalbpool "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbpool"
	nsxtalbserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbserviceenginegroup"
	nsxtalbsettings "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbsettings"
	nsxtalbvirtualservice "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtalbvirtualservice"
	nsxtappportprofile "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtappportprofile"
	nsxtdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtdistributedfirewall"
	nsxtdistributedfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtdistributedfirewallrule"
	nsxtdynamicsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtdynamicsecuritygroup"
	nsxtedgegateway "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegateway"
	nsxtedgegatewaybgpconfiguration "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewaybgpconfiguration"
	nsxtedgegatewaybgpipprefixlist "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewaybgpipprefixlist"
	nsxtedgegatewaybgpneighbor "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewaybgpneighbor"
	nsxtedgegatewaydhcpv6 "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewaydhcpv6"
	nsxtedgegatewayratelimit "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewayratelimit"
	nsxtedgegatewaystaticroute "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtedgegatewaystaticroute"
	nsxtfirewall "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtfirewall"
	nsxtipsecvpntunnel "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtipsecvpntunnel"
	nsxtipset "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtipset"
	nsxtnatrule "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtnatrule"
	nsxtnetworkdhcp "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtnetworkdhcp"
	nsxtnetworkdhcpbinding "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtnetworkdhcpbinding"
	nsxtnetworkimported "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtnetworkimported"
	nsxtrouteadvertisement "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtrouteadvertisement"
	nsxtsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxtsecuritygroup"
	nsxvdhcprelay "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvdhcprelay"
	nsxvdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvdistributedfirewall"
	nsxvdnat "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvdnat"
	nsxvfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvfirewallrule"
	nsxvipset "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvipset"
	nsxvsnat "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/nsxvsnat"
	org "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/org"
	orggroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orggroup"
	orgldap "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orgldap"
	orgsaml "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orgsaml"
	orguser "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orguser"
	orgvdc "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orgvdc"
	orgvdcaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/orgvdcaccesscontrol"
	providervdc "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/providervdc"
	rde "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rde"
	rdeinterface "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rdeinterface"
	rdeinterfacebehavior "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rdeinterfacebehavior"
	rdetype "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rdetype"
	rdetypebehavior "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rdetypebehavior"
	rdetypebehavioracl "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rdetypebehavioracl"
	rightsbundle "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/rightsbundle"
	role "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/role"
	securitytag "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/securitytag"
	serviceaccount "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/serviceaccount"
	subscribedcatalog "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/subscribedcatalog"
	uiplugin "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/uiplugin"
	vapp "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vapp"
	vappaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappaccesscontrol"
	vappfirewallrules "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappfirewallrules"
	vappnatrules "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappnatrules"
	vappnetwork "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappnetwork"
	vapporgnetwork "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vapporgnetwork"
	vappstaticrouting "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappstaticrouting"
	vappvm "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vappvm"
	vdcgroup "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vdcgroup"
	vm "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vm"
	vmaffinityrule "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vmaffinityrule"
	vminternaldisk "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vminternaldisk"
	vmplacementpolicy "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vmplacementpolicy"
	vmsizingpolicy "github.com/arkilasystems/provider-vcd/internal/controller/namespaced/vcd/vmsizingpolicy"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		apitoken.Setup,
		catalog.Setup,
		catalogitem.Setup,
		catalogmedia.Setup,
		catalogvapptemplate.Setup,
		clonedvapp.Setup,
		edgegateway.Setup,
		edgegatewaysettings.Setup,
		edgegatewayvpn.Setup,
		externalnetwork.Setup,
		externalnetworkv2.Setup,
		globalrole.Setup,
		independentdisk.Setup,
		insertedmedia.Setup,
		ipspace.Setup,
		ipspacecustomquota.Setup,
		ipspaceipallocation.Setup,
		ipspaceuplink.Setup,
		lbappprofile.Setup,
		lbapprule.Setup,
		lbserverpool.Setup,
		lbservicemonitor.Setup,
		lbvirtualserver.Setup,
		networkdirect.Setup,
		networkisolated.Setup,
		networkisolatedv2.Setup,
		networkrouted.Setup,
		networkroutedv2.Setup,
		nsxtalbcloud.Setup,
		nsxtalbcontroller.Setup,
		nsxtalbedgegatewayserviceenginegroup.Setup,
		nsxtalbpool.Setup,
		nsxtalbserviceenginegroup.Setup,
		nsxtalbsettings.Setup,
		nsxtalbvirtualservice.Setup,
		nsxtappportprofile.Setup,
		nsxtdistributedfirewall.Setup,
		nsxtdistributedfirewallrule.Setup,
		nsxtdynamicsecuritygroup.Setup,
		nsxtedgegateway.Setup,
		nsxtedgegatewaybgpconfiguration.Setup,
		nsxtedgegatewaybgpipprefixlist.Setup,
		nsxtedgegatewaybgpneighbor.Setup,
		nsxtedgegatewaydhcpv6.Setup,
		nsxtedgegatewayratelimit.Setup,
		nsxtedgegatewaystaticroute.Setup,
		nsxtfirewall.Setup,
		nsxtipsecvpntunnel.Setup,
		nsxtipset.Setup,
		nsxtnatrule.Setup,
		nsxtnetworkdhcp.Setup,
		nsxtnetworkdhcpbinding.Setup,
		nsxtnetworkimported.Setup,
		nsxtrouteadvertisement.Setup,
		nsxtsecuritygroup.Setup,
		nsxvdhcprelay.Setup,
		nsxvdistributedfirewall.Setup,
		nsxvdnat.Setup,
		nsxvfirewallrule.Setup,
		nsxvipset.Setup,
		nsxvsnat.Setup,
		org.Setup,
		orggroup.Setup,
		orgldap.Setup,
		orgsaml.Setup,
		orguser.Setup,
		orgvdc.Setup,
		orgvdcaccesscontrol.Setup,
		providervdc.Setup,
		rde.Setup,
		rdeinterface.Setup,
		rdeinterfacebehavior.Setup,
		rdetype.Setup,
		rdetypebehavior.Setup,
		rdetypebehavioracl.Setup,
		rightsbundle.Setup,
		role.Setup,
		securitytag.Setup,
		serviceaccount.Setup,
		subscribedcatalog.Setup,
		uiplugin.Setup,
		vapp.Setup,
		vappaccesscontrol.Setup,
		vappfirewallrules.Setup,
		vappnatrules.Setup,
		vappnetwork.Setup,
		vapporgnetwork.Setup,
		vappstaticrouting.Setup,
		vappvm.Setup,
		vdcgroup.Setup,
		vm.Setup,
		vmaffinityrule.Setup,
		vminternaldisk.Setup,
		vmplacementpolicy.Setup,
		vmsizingpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		apitoken.SetupGated,
		catalog.SetupGated,
		catalogitem.SetupGated,
		catalogmedia.SetupGated,
		catalogvapptemplate.SetupGated,
		clonedvapp.SetupGated,
		edgegateway.SetupGated,
		edgegatewaysettings.SetupGated,
		edgegatewayvpn.SetupGated,
		externalnetwork.SetupGated,
		externalnetworkv2.SetupGated,
		globalrole.SetupGated,
		independentdisk.SetupGated,
		insertedmedia.SetupGated,
		ipspace.SetupGated,
		ipspacecustomquota.SetupGated,
		ipspaceipallocation.SetupGated,
		ipspaceuplink.SetupGated,
		lbappprofile.SetupGated,
		lbapprule.SetupGated,
		lbserverpool.SetupGated,
		lbservicemonitor.SetupGated,
		lbvirtualserver.SetupGated,
		networkdirect.SetupGated,
		networkisolated.SetupGated,
		networkisolatedv2.SetupGated,
		networkrouted.SetupGated,
		networkroutedv2.SetupGated,
		nsxtalbcloud.SetupGated,
		nsxtalbcontroller.SetupGated,
		nsxtalbedgegatewayserviceenginegroup.SetupGated,
		nsxtalbpool.SetupGated,
		nsxtalbserviceenginegroup.SetupGated,
		nsxtalbsettings.SetupGated,
		nsxtalbvirtualservice.SetupGated,
		nsxtappportprofile.SetupGated,
		nsxtdistributedfirewall.SetupGated,
		nsxtdistributedfirewallrule.SetupGated,
		nsxtdynamicsecuritygroup.SetupGated,
		nsxtedgegateway.SetupGated,
		nsxtedgegatewaybgpconfiguration.SetupGated,
		nsxtedgegatewaybgpipprefixlist.SetupGated,
		nsxtedgegatewaybgpneighbor.SetupGated,
		nsxtedgegatewaydhcpv6.SetupGated,
		nsxtedgegatewayratelimit.SetupGated,
		nsxtedgegatewaystaticroute.SetupGated,
		nsxtfirewall.SetupGated,
		nsxtipsecvpntunnel.SetupGated,
		nsxtipset.SetupGated,
		nsxtnatrule.SetupGated,
		nsxtnetworkdhcp.SetupGated,
		nsxtnetworkdhcpbinding.SetupGated,
		nsxtnetworkimported.SetupGated,
		nsxtrouteadvertisement.SetupGated,
		nsxtsecuritygroup.SetupGated,
		nsxvdhcprelay.SetupGated,
		nsxvdistributedfirewall.SetupGated,
		nsxvdnat.SetupGated,
		nsxvfirewallrule.SetupGated,
		nsxvipset.SetupGated,
		nsxvsnat.SetupGated,
		org.SetupGated,
		orggroup.SetupGated,
		orgldap.SetupGated,
		orgsaml.SetupGated,
		orguser.SetupGated,
		orgvdc.SetupGated,
		orgvdcaccesscontrol.SetupGated,
		providervdc.SetupGated,
		rde.SetupGated,
		rdeinterface.SetupGated,
		rdeinterfacebehavior.SetupGated,
		rdetype.SetupGated,
		rdetypebehavior.SetupGated,
		rdetypebehavioracl.SetupGated,
		rightsbundle.SetupGated,
		role.SetupGated,
		securitytag.SetupGated,
		serviceaccount.SetupGated,
		subscribedcatalog.SetupGated,
		uiplugin.SetupGated,
		vapp.SetupGated,
		vappaccesscontrol.SetupGated,
		vappfirewallrules.SetupGated,
		vappnatrules.SetupGated,
		vappnetwork.SetupGated,
		vapporgnetwork.SetupGated,
		vappstaticrouting.SetupGated,
		vappvm.SetupGated,
		vdcgroup.SetupGated,
		vm.SetupGated,
		vmaffinityrule.SetupGated,
		vminternaldisk.SetupGated,
		vmplacementpolicy.SetupGated,
		vmsizingpolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		providerconfig.SetupWebhookWithManager,
		apitoken.SetupWebhookWithManager,
		catalog.SetupWebhookWithManager,
		catalogitem.SetupWebhookWithManager,
		catalogmedia.SetupWebhookWithManager,
		catalogvapptemplate.SetupWebhookWithManager,
		clonedvapp.SetupWebhookWithManager,
		edgegateway.SetupWebhookWithManager,
		edgegatewaysettings.SetupWebhookWithManager,
		edgegatewayvpn.SetupWebhookWithManager,
		externalnetwork.SetupWebhookWithManager,
		externalnetworkv2.SetupWebhookWithManager,
		globalrole.SetupWebhookWithManager,
		independentdisk.SetupWebhookWithManager,
		insertedmedia.SetupWebhookWithManager,
		ipspace.SetupWebhookWithManager,
		ipspacecustomquota.SetupWebhookWithManager,
		ipspaceipallocation.SetupWebhookWithManager,
		ipspaceuplink.SetupWebhookWithManager,
		lbappprofile.SetupWebhookWithManager,
		lbapprule.SetupWebhookWithManager,
		lbserverpool.SetupWebhookWithManager,
		lbservicemonitor.SetupWebhookWithManager,
		lbvirtualserver.SetupWebhookWithManager,
		networkdirect.SetupWebhookWithManager,
		networkisolated.SetupWebhookWithManager,
		networkisolatedv2.SetupWebhookWithManager,
		networkrouted.SetupWebhookWithManager,
		networkroutedv2.SetupWebhookWithManager,
		nsxtalbcloud.SetupWebhookWithManager,
		nsxtalbcontroller.SetupWebhookWithManager,
		nsxtalbedgegatewayserviceenginegroup.SetupWebhookWithManager,
		nsxtalbpool.SetupWebhookWithManager,
		nsxtalbserviceenginegroup.SetupWebhookWithManager,
		nsxtalbsettings.SetupWebhookWithManager,
		nsxtalbvirtualservice.SetupWebhookWithManager,
		nsxtappportprofile.SetupWebhookWithManager,
		nsxtdistributedfirewall.SetupWebhookWithManager,
		nsxtdistributedfirewallrule.SetupWebhookWithManager,
		nsxtdynamicsecuritygroup.SetupWebhookWithManager,
		nsxtedgegateway.SetupWebhookWithManager,
		nsxtedgegatewaybgpconfiguration.SetupWebhookWithManager,
		nsxtedgegatewaybgpipprefixlist.SetupWebhookWithManager,
		nsxtedgegatewaybgpneighbor.SetupWebhookWithManager,
		nsxtedgegatewaydhcpv6.SetupWebhookWithManager,
		nsxtedgegatewayratelimit.SetupWebhookWithManager,
		nsxtedgegatewaystaticroute.SetupWebhookWithManager,
		nsxtfirewall.SetupWebhookWithManager,
		nsxtipsecvpntunnel.SetupWebhookWithManager,
		nsxtipset.SetupWebhookWithManager,
		nsxtnatrule.SetupWebhookWithManager,
		nsxtnetworkdhcp.SetupWebhookWithManager,
		nsxtnetworkdhcpbinding.SetupWebhookWithManager,
		nsxtnetworkimported.SetupWebhookWithManager,
		nsxtrouteadvertisement.SetupWebhookWithManager,
		nsxtsecuritygroup.SetupWebhookWithManager,
		nsxvdhcprelay.SetupWebhookWithManager,
		nsxvdistributedfirewall.SetupWebhookWithManager,
		nsxvdnat.SetupWebhookWithManager,
		nsxvfirewallrule.SetupWebhookWithManager,
		nsxvipset.SetupWebhookWithManager,
		nsxvsnat.SetupWebhookWithManager,
		org.SetupWebhookWithManager,
		orggroup.SetupWebhookWithManager,
		orgldap.SetupWebhookWithManager,
		orgsaml.SetupWebhookWithManager,
		orguser.SetupWebhookWithManager,
		orgvdc.SetupWebhookWithManager,
		orgvdcaccesscontrol.SetupWebhookWithManager,
		providervdc.SetupWebhookWithManager,
		rde.SetupWebhookWithManager,
		rdeinterface.SetupWebhookWithManager,
		rdeinterfacebehavior.SetupWebhookWithManager,
		rdetype.SetupWebhookWithManager,
		rdetypebehavior.SetupWebhookWithManager,
		rdetypebehavioracl.SetupWebhookWithManager,
		rightsbundle.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		securitytag.SetupWebhookWithManager,
		serviceaccount.SetupWebhookWithManager,
		subscribedcatalog.SetupWebhookWithManager,
		uiplugin.SetupWebhookWithManager,
		vapp.SetupWebhookWithManager,
		vappaccesscontrol.SetupWebhookWithManager,
		vappfirewallrules.SetupWebhookWithManager,
		vappnatrules.SetupWebhookWithManager,
		vappnetwork.SetupWebhookWithManager,
		vapporgnetwork.SetupWebhookWithManager,
		vappstaticrouting.SetupWebhookWithManager,
		vappvm.SetupWebhookWithManager,
		vdcgroup.SetupWebhookWithManager,
		vm.SetupWebhookWithManager,
		vmaffinityrule.SetupWebhookWithManager,
		vminternaldisk.SetupWebhookWithManager,
		vmplacementpolicy.SetupWebhookWithManager,
		vmsizingpolicy.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
