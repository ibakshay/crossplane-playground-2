/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// BmwParameters are the configurable fields of a Bmw.
type BmwParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// BmwObservation are the observable fields of a Bmw.
type BmwObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A BmwSpec defines the desired state of a Bmw.
type BmwSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BmwParameters `json:"forProvider"`
}

// A BmwStatus represents the observed state of a Bmw.
type BmwStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BmwObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Bmw is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,myplayground}
type Bmw struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BmwSpec   `json:"spec"`
	Status BmwStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BmwList contains a list of Bmw
type BmwList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bmw `json:"items"`
}

// Bmw type metadata.
var (
	BmwKind             = reflect.TypeOf(Bmw{}).Name()
	BmwGroupKind        = schema.GroupKind{Group: Group, Kind: BmwKind}.String()
	BmwKindAPIVersion   = BmwKind + "." + SchemeGroupVersion.String()
	BmwGroupVersionKind = SchemeGroupVersion.WithKind(BmwKind)
)

func init() {
	SchemeBuilder.Register(&Bmw{}, &BmwList{})
}
