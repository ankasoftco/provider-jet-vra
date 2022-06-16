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

type AccountNsxvLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type AccountNsxvLinksParameters struct {
}

type AccountNsxvObservation struct {
	AssociatedCloudAccountIds []*string `json:"associatedCloudAccountIds,omitempty" tf:"associated_cloud_account_ids,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Links []AccountNsxvLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type AccountNsxvParameters struct {

	// +kubebuilder:validation:Optional
	AcceptSelfSignedCert *bool `json:"acceptSelfSignedCert,omitempty" tf:"accept_self_signed_cert,omitempty"`

	// +kubebuilder:validation:Optional
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []AccountNsxvTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type AccountNsxvTagsObservation struct {
}

type AccountNsxvTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// AccountNsxvSpec defines the desired state of AccountNsxv
type AccountNsxvSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountNsxvParameters `json:"forProvider"`
}

// AccountNsxvStatus defines the observed state of AccountNsxv.
type AccountNsxvStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountNsxvObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNsxv is the Schema for the AccountNsxvs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type AccountNsxv struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountNsxvSpec   `json:"spec"`
	Status            AccountNsxvStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNsxvList contains a list of AccountNsxvs
type AccountNsxvList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountNsxv `json:"items"`
}

// Repository type metadata.
var (
	AccountNsxv_Kind             = "AccountNsxv"
	AccountNsxv_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountNsxv_Kind}.String()
	AccountNsxv_KindAPIVersion   = AccountNsxv_Kind + "." + CRDGroupVersion.String()
	AccountNsxv_GroupVersionKind = CRDGroupVersion.WithKind(AccountNsxv_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountNsxv{}, &AccountNsxvList{})
}