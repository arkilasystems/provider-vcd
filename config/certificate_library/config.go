package certificatelibrary

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_certificate_library", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "CertificateLibrary"
		r.Version = "v1alpha1"
	})
}
