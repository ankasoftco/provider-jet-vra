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

// Code generated by terrajet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-vra/apis/blockdevice/v1alpha1"
	v1alpha1blueprint "github.com/crossplane-contrib/provider-jet-vra/apis/blueprint/v1alpha1"
	v1alpha1catalogitementitlement "github.com/crossplane-contrib/provider-jet-vra/apis/catalogitementitlement/v1alpha1"
	v1alpha1catalogsourceblueprint "github.com/crossplane-contrib/provider-jet-vra/apis/catalogsourceblueprint/v1alpha1"
	v1alpha1catalogsourceentitlement "github.com/crossplane-contrib/provider-jet-vra/apis/catalogsourceentitlement/v1alpha1"
	v1alpha1cloudaccount "github.com/crossplane-contrib/provider-jet-vra/apis/cloudaccount/v1alpha1"
	v1alpha1contentsource "github.com/crossplane-contrib/provider-jet-vra/apis/contentsource/v1alpha1"
	v1alpha1deployment "github.com/crossplane-contrib/provider-jet-vra/apis/deployment/v1alpha1"
	v1alpha1fabric "github.com/crossplane-contrib/provider-jet-vra/apis/fabric/v1alpha1"
	v1alpha1flavorprofile "github.com/crossplane-contrib/provider-jet-vra/apis/flavorprofile/v1alpha1"
	v1alpha1imageprofile "github.com/crossplane-contrib/provider-jet-vra/apis/imageprofile/v1alpha1"
	v1alpha1loadbalancer "github.com/crossplane-contrib/provider-jet-vra/apis/loadbalancer/v1alpha1"
	v1alpha1machine "github.com/crossplane-contrib/provider-jet-vra/apis/machine/v1alpha1"
	v1alpha1network "github.com/crossplane-contrib/provider-jet-vra/apis/network/v1alpha1"
	v1alpha1project "github.com/crossplane-contrib/provider-jet-vra/apis/project/v1alpha1"
	v1alpha1storage "github.com/crossplane-contrib/provider-jet-vra/apis/storage/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-vra/apis/v1alpha1"
	v1alpha1zone "github.com/crossplane-contrib/provider-jet-vra/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1blueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogitementitlement.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceblueprint.SchemeBuilder.AddToScheme,
		v1alpha1catalogsourceentitlement.SchemeBuilder.AddToScheme,
		v1alpha1cloudaccount.SchemeBuilder.AddToScheme,
		v1alpha1contentsource.SchemeBuilder.AddToScheme,
		v1alpha1deployment.SchemeBuilder.AddToScheme,
		v1alpha1fabric.SchemeBuilder.AddToScheme,
		v1alpha1flavorprofile.SchemeBuilder.AddToScheme,
		v1alpha1imageprofile.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1machine.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1zone.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
