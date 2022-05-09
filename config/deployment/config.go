package deployment

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_deployment", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "deployment"
	})
}
