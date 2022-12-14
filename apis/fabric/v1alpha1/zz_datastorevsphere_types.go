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

type DatastoreVsphereLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type DatastoreVsphereLinksParameters struct {
}

type DatastoreVsphereObservation struct {
	CloudAccountIds []*string `json:"cloudAccountIds,omitempty" tf:"cloud_account_ids,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	FreeSizeGb *string `json:"freeSizeGb,omitempty" tf:"free_size_gb,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []DatastoreVsphereLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type DatastoreVsphereParameters struct {

	// +kubebuilder:validation:Optional
	Tags []DatastoreVsphereTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DatastoreVsphereTagsObservation struct {
}

type DatastoreVsphereTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// DatastoreVsphereSpec defines the desired state of DatastoreVsphere
type DatastoreVsphereSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatastoreVsphereParameters `json:"forProvider"`
}

// DatastoreVsphereStatus defines the observed state of DatastoreVsphere.
type DatastoreVsphereStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatastoreVsphereObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatastoreVsphere is the Schema for the DatastoreVspheres API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type DatastoreVsphere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatastoreVsphereSpec   `json:"spec"`
	Status            DatastoreVsphereStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatastoreVsphereList contains a list of DatastoreVspheres
type DatastoreVsphereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastoreVsphere `json:"items"`
}

// Repository type metadata.
var (
	DatastoreVsphere_Kind             = "DatastoreVsphere"
	DatastoreVsphere_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatastoreVsphere_Kind}.String()
	DatastoreVsphere_KindAPIVersion   = DatastoreVsphere_Kind + "." + CRDGroupVersion.String()
	DatastoreVsphere_GroupVersionKind = CRDGroupVersion.WithKind(DatastoreVsphere_Kind)
)

func init() {
	SchemeBuilder.Register(&DatastoreVsphere{}, &DatastoreVsphereList{})
}
