package insertedmedia

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_inserted_media", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "InsertedMedia"
		r.Version = "v1alpha1"
	})
}
