// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/arkilasystems/provider-vcd/internal/controller/cluster/providerconfig"
	apitoken "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/apitoken"
	catalog "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/catalog"
	catalogitem "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/catalogitem"
	catalogmedia "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/catalogmedia"
	catalogvapptemplate "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/catalogvapptemplate"
	clonedvapp "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/clonedvapp"
	edgegateway "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/edgegateway"
	edgegatewaysettings "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/edgegatewaysettings"
	edgegatewayvpn "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/edgegatewayvpn"
	externalnetwork "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/externalnetwork"
	externalnetworkv2 "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/externalnetworkv2"
	globalrole "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/globalrole"
	independentdisk "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/independentdisk"
	insertedmedia "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/insertedmedia"
	ipspace "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/ipspace"
	ipspacecustomquota "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/ipspacecustomquota"
	ipspaceipallocation "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/ipspaceipallocation"
	ipspaceuplink "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/ipspaceuplink"
	lbappprofile "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/lbappprofile"
	lbapprule "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/lbapprule"
	lbserverpool "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/lbserverpool"
	lbservicemonitor "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/lbservicemonitor"
	lbvirtualserver "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/lbvirtualserver"
	networkdirect "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/networkdirect"
	networkisolated "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/networkisolated"
	networkisolatedv2 "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/networkisolatedv2"
	networkrouted "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/networkrouted"
	networkroutedv2 "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/networkroutedv2"
	nsxtalbcloud "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbcloud"
	nsxtalbcontroller "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbcontroller"
	nsxtalbedgegatewayserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbedgegatewayserviceenginegroup"
	nsxtalbpool "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbpool"
	nsxtalbserviceenginegroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbserviceenginegroup"
	nsxtalbsettings "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbsettings"
	nsxtalbvirtualservice "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtalbvirtualservice"
	nsxtappportprofile "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtappportprofile"
	nsxtdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtdistributedfirewall"
	nsxtdistributedfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtdistributedfirewallrule"
	nsxtdynamicsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtdynamicsecuritygroup"
	nsxtedgegateway "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegateway"
	nsxtedgegatewaybgpconfiguration "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewaybgpconfiguration"
	nsxtedgegatewaybgpipprefixlist "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewaybgpipprefixlist"
	nsxtedgegatewaybgpneighbor "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewaybgpneighbor"
	nsxtedgegatewaydhcpv6 "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewaydhcpv6"
	nsxtedgegatewayratelimit "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewayratelimit"
	nsxtedgegatewaystaticroute "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtedgegatewaystaticroute"
	nsxtfirewall "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtfirewall"
	nsxtipsecvpntunnel "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtipsecvpntunnel"
	nsxtipset "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtipset"
	nsxtnatrule "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtnatrule"
	nsxtnetworkdhcp "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtnetworkdhcp"
	nsxtnetworkdhcpbinding "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtnetworkdhcpbinding"
	nsxtnetworkimported "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtnetworkimported"
	nsxtrouteadvertisement "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtrouteadvertisement"
	nsxtsecuritygroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxtsecuritygroup"
	nsxvdhcprelay "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvdhcprelay"
	nsxvdistributedfirewall "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvdistributedfirewall"
	nsxvdnat "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvdnat"
	nsxvfirewallrule "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvfirewallrule"
	nsxvipset "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvipset"
	nsxvsnat "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/nsxvsnat"
	org "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/org"
	orggroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orggroup"
	orgldap "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orgldap"
	orgsaml "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orgsaml"
	orguser "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orguser"
	orgvdc "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orgvdc"
	orgvdcaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/orgvdcaccesscontrol"
	providervdc "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/providervdc"
	rde "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rde"
	rdeinterface "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rdeinterface"
	rdeinterfacebehavior "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rdeinterfacebehavior"
	rdetype "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rdetype"
	rdetypebehavior "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rdetypebehavior"
	rdetypebehavioracl "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rdetypebehavioracl"
	rightsbundle "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/rightsbundle"
	role "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/role"
	securitytag "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/securitytag"
	serviceaccount "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/serviceaccount"
	subscribedcatalog "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/subscribedcatalog"
	uiplugin "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/uiplugin"
	vapp "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vapp"
	vappaccesscontrol "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappaccesscontrol"
	vappfirewallrules "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappfirewallrules"
	vappnatrules "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappnatrules"
	vappnetwork "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappnetwork"
	vapporgnetwork "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vapporgnetwork"
	vappstaticrouting "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappstaticrouting"
	vappvm "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vappvm"
	vdcgroup "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vdcgroup"
	vm "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vm"
	vmaffinityrule "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vmaffinityrule"
	vminternaldisk "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vminternaldisk"
	vmplacementpolicy "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vmplacementpolicy"
	vmsizingpolicy "github.com/arkilasystems/provider-vcd/internal/controller/cluster/vcd/vmsizingpolicy"
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
