/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/crossplane-contrib/provider-jet-vra/config/fabric_datastore_vsphere"
	"github.com/crossplane-contrib/provider-jet-vra/config/fabric_network_vsphere"

	catalog_item_entitlement "github.com/crossplane-contrib/provider-jet-vra/config/catalog_item_entitlement"
	catalog_source_blueprint "github.com/crossplane-contrib/provider-jet-vra/config/catalog_source_blueprint"
	catalog_source_entitlement "github.com/crossplane-contrib/provider-jet-vra/config/catalog_source_entitlement"
	content_source "github.com/crossplane-contrib/provider-jet-vra/config/content_source"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-vra/config/block_device"
	"github.com/crossplane-contrib/provider-jet-vra/config/block_device_snapshot"
	"github.com/crossplane-contrib/provider-jet-vra/config/blueprint"
	"github.com/crossplane-contrib/provider-jet-vra/config/deployment"
	"github.com/crossplane-contrib/provider-jet-vra/config/fabric_compute"
	"github.com/crossplane-contrib/provider-jet-vra/config/project"
	"github.com/crossplane-contrib/provider-jet-vra/config/zone"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/crossplane-contrib/provider-jet-vra"
)

//go:embed schema.json
var providerSchema string

// IncludedResources s
var IncludedResources = []string{
	"vra_deployment$",
	"vra_blueprint$",
	"vra_blueprint_version$",
	"vra_catalog_item_entitlement$",
	"vra_catalog_source_blueprint$",
	"vra_catalog_source_entitlement$",
	"vra_content_source$",
	"vra_project$",
	"vra_zone$",
	"block_device$",
	"block_device_snapshot$",
	"fabric_compute$",
	"fabric_datastore_vsphere$",
	"fabric_network_vsphere$",
}

// skipList
var skipList = []string{}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithShortName("vrajet"),
		tjconfig.WithRootGroup("vra.jet.crossplane.io"),
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		deployment.Configure,
		blueprint.Configure,
		catalog_item_entitlement.Configure,
		catalog_source_blueprint.Configure,
		catalog_source_entitlement.Configure,
		content_source.Configure,
		project.Configure,
		zone.Configure,
		block_device.Configure,
		block_device_snapshot.Configure,
		fabric_compute.Configure,
		fabric_datastore_vsphere.Configure,
		fabric_network_vsphere.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}