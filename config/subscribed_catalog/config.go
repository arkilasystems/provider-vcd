package subscribedcatalog

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_subscribed_catalog", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "SubscribedCatalog"
		r.Version = "v1alpha1"
	})
}
