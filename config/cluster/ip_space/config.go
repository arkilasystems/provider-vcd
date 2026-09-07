package ipspace

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_ip_space", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpace"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_ip_space_custom_quota", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceCustomQuota"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_ip_space_ip_allocation", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceIpAllocation"
		r.Version = "v1alpha1"

		// The upstream Terraform provider's own schema description for
		// "type" contains a malformed double-backtick ("`FLOATING_IP``,").
		// Depending on the platform, upjet's markdown-to-doc-comment
		// conversion resolves that ambiguous sequence differently (a
		// literal double backtick on one platform, a Unicode smart quote
		// on another), which made code generation non-deterministic across
		// machines/CI runners. Fixed here at the schema level, before code
		// generation runs, since the live Terraform schema fetch re-derives
		// the same malformed text on every `make generate` regardless of
		// what's committed to config/schema.json.
		if typ, ok := r.TerraformResource.Schema["type"]; ok {
			typ.Description = "Type of allocation. One of `FLOATING_IP`, `IP_PREFIX`"
		}
	})

	p.AddResourceConfigurator("vcd_ip_space_uplink", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceUplink"
		r.Version = "v1alpha1"
	})
}
