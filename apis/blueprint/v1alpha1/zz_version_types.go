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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VersionObservation struct {
	BlueprintDescription *string `json:"blueprintDescription,omitempty" tf:"blueprint_description,omitempty"`

	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`

	Valid *string `json:"valid,omitempty" tf:"valid,omitempty"`
}

type VersionParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-vra/apis/blueprint/v1alpha1.Blueprint
	// +kubebuilder:validation:Optional
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// +kubebuilder:validation:Optional
	BlueprintIDRef *v1.Reference `json:"blueprintIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BlueprintIDSelector *v1.Selector `json:"blueprintIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ChangeLog *string `json:"changeLog,omitempty" tf:"change_log,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Release *bool `json:"release,omitempty" tf:"release,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// VersionSpec defines the desired state of Version
type VersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VersionParameters `json:"forProvider"`
}

// VersionStatus defines the observed state of Version.
type VersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Version is the Schema for the Versions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VersionSpec   `json:"spec"`
	Status            VersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VersionList contains a list of Versions
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Version `json:"items"`
}

// Repository type metadata.
var (
	Version_Kind             = "Version"
	Version_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Version_Kind}.String()
	Version_KindAPIVersion   = Version_Kind + "." + CRDGroupVersion.String()
	Version_GroupVersionKind = CRDGroupVersion.WithKind(Version_Kind)
)

func init() {
	SchemeBuilder.Register(&Version{}, &VersionList{})
}