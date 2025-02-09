// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/arkilasystems/provider-vcd/internal/controller/providerconfig"
	apitoken "github.com/arkilasystems/provider-vcd/internal/controller/vcd/apitoken"
	catalog "github.com/arkilasystems/provider-vcd/internal/controller/vcd/catalog"
	catalogitem "github.com/arkilasystems/provider-vcd/internal/controller/vcd/catalogitem"
	catalogmedia "github.com/arkilasystems/provider-vcd/internal/controller/vcd/catalogmedia"
	catalogvapptemplate "github.com/arkilasystems/provider-vcd/internal/controller/vcd/catalogvapptemplate"
	clonedvapp "github.com/arkilasystems/provider-vcd/internal/controller/vcd/clonedvapp"
	edgegateway "github.com/arkilasystems/provider-vcd/internal/controller/vcd/edgegateway"
	edgegatewaysettings "github.com/arkilasystems/provider-vcd/internal/controller/vcd/edgegatewaysettings"
	edgegatewayvpn "github.com/arkilasystems/provider-vcd/internal/controller/vcd/edgegatewayvpn"
	externalnetwork "github.com/arkilasystems/provider-vcd/internal/controller/vcd/externalnetwork"
	externalnetworkv2 "github.com/arkilasystems/provider-vcd/internal/controller/vcd/externalnetworkv2"
	globalrole "github.com/arkilasystems/provider-vcd/internal/controller/vcd/globalrole"
	independentdisk "github.com/arkilasystems/provider-vcd/internal/controller/vcd/independentdisk"
	insertedmedia "github.com/arkilasystems/provider-vcd/internal/controller/vcd/insertedmedia"
	ipspace "github.com/arkilasystems/provider-vcd/internal/controller/vcd/ipspace"
	ipspacecustomquota "github.com/arkilasystems/provider-vcd/internal/controller/vcd/ipspacecustomquota"
	ipspaceipallocation "github.com/arkilasystems/provider-vcd/internal/controller/vcd/ipspaceipallocation"
	ipspaceuplink "github.com/arkilasystems/provider-vcd/internal/controller/vcd/ipspaceuplink"
	lbappprofile "github.com/arkilasystems/provider-vcd/internal/controller/vcd/lbappprofile"
	lbapprule "github.com/arkilasystems/provider-vcd/internal/controller/vcd/lbapprule"
	lbserverpool "github.com/arkilasystems/provider-vcd/internal/controller/vcd/lbserverpool"
	lbservicemonitor "github.com/arkilasystems/provider-vcd/internal/controller/vcd/lbservicemonitor"
	lbvirtualserver "github.com/arkilasystems/provider-vcd/internal/controller/vcd/lbvirtualserver"
	networkdirect "github.com/arkilasystems/provider-vcd/internal/controller/vcd/networkdirect"
	networkisolated "github.com/arkilasystems/provider-vcd/internal/controller/vcd/networkisolated"
	networkisolatedv2 "github.com/arkilasystems/provider-vcd/internal/controller/vcd/networkisolatedv2"
	networkrouted "github.com/arkilasystems/provider-vcd/internal/controller/vcd/networkrouted"
	networkroutedv2 "github.com/arkilasystems/provider-vcd/internal/controller/vcd/networkroutedv2"
	nsxtalbcloud "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbcloud"
	nsxtalbcontroller "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbcontroller"
	nsxtalbedgegatewayserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbedgegatewayserviceenginegroup"
	nsxtalbpool "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbpool"
	nsxtalbserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbserviceenginegroup"
	nsxtalbsettings "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbsettings"
	nsxtalbvirtualservice "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtalbvirtualservice"
	nsxtappportprofile "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtappportprofile"
	nsxtdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtdistributedfirewall"
	nsxtdistributedfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtdistributedfirewallrule"
	nsxtdynamicsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtdynamicsecuritygroup"
	nsxtedgegateway "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegateway"
	nsxtedgegatewaybgpconfiguration "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpconfiguration"
	nsxtedgegatewaybgpipprefixlist "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpipprefixlist"
	nsxtedgegatewaybgpneighbor "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpneighbor"
	nsxtedgegatewaydhcpv6 "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewaydhcpv6"
	nsxtedgegatewayratelimit "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewayratelimit"
	nsxtedgegatewaystaticroute "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtedgegatewaystaticroute"
	nsxtfirewall "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtfirewall"
	nsxtipsecvpntunnel "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtipsecvpntunnel"
	nsxtipset "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtipset"
	nsxtnatrule "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtnatrule"
	nsxtnetworkdhcp "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtnetworkdhcp"
	nsxtnetworkdhcpbinding "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtnetworkdhcpbinding"
	nsxtnetworkimported "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtnetworkimported"
	nsxtrouteadvertisement "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtrouteadvertisement"
	nsxtsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxtsecuritygroup"
	nsxvdhcprelay "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvdhcprelay"
	nsxvdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvdistributedfirewall"
	nsxvdnat "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvdnat"
	nsxvfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvfirewallrule"
	nsxvipset "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvipset"
	nsxvsnat "github.com/arkilasystems/provider-vcd/internal/controller/vcd/nsxvsnat"
	org "github.com/arkilasystems/provider-vcd/internal/controller/vcd/org"
	orggroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orggroup"
	orgldap "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orgldap"
	orgsaml "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orgsaml"
	orguser "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orguser"
	orgvdc "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orgvdc"
	orgvdcaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/vcd/orgvdcaccesscontrol"
	providervdc "github.com/arkilasystems/provider-vcd/internal/controller/vcd/providervdc"
	rde "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rde"
	rdeinterface "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rdeinterface"
	rdeinterfacebehavior "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rdeinterfacebehavior"
	rdetype "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rdetype"
	rdetypebehavior "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rdetypebehavior"
	rdetypebehavioracl "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rdetypebehavioracl"
	rightsbundle "github.com/arkilasystems/provider-vcd/internal/controller/vcd/rightsbundle"
	role "github.com/arkilasystems/provider-vcd/internal/controller/vcd/role"
	securitytag "github.com/arkilasystems/provider-vcd/internal/controller/vcd/securitytag"
	serviceaccount "github.com/arkilasystems/provider-vcd/internal/controller/vcd/serviceaccount"
	subscribedcatalog "github.com/arkilasystems/provider-vcd/internal/controller/vcd/subscribedcatalog"
	uiplugin "github.com/arkilasystems/provider-vcd/internal/controller/vcd/uiplugin"
	vapp "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vapp"
	vappaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappaccesscontrol"
	vappfirewallrules "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappfirewallrules"
	vappnatrules "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappnatrules"
	vappnetwork "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappnetwork"
	vapporgnetwork "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vapporgnetwork"
	vappstaticrouting "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappstaticrouting"
	vappvm "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vappvm"
	vdcgroup "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vdcgroup"
	vm "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vm"
	vmaffinityrule "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vmaffinityrule"
	vminternaldisk "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vminternaldisk"
	vmplacementpolicy "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vmplacementpolicy"
	vmsizingpolicy "github.com/arkilasystems/provider-vcd/internal/controller/vcd/vmsizingpolicy"
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
