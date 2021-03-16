/*
Copyright 2020 Wim Henderickx.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SrlNokiaSystemSystemNameFinalizer is the name of the finalizer added to
	// SrlNokiaSystemSystemName to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaSystemSystemNameFinalizer string = "Name.srlinux.henderiw.be"
)

// SrlNokiaSystemSystemName struct
type SrlNokiaSystemSystemName struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	DomainName *string `json:"domain-name,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])`
	HostName *string `json:"host-name,omitempty"`
}

// SrlNokiaSystemSystemNameSpec struct
type SrlNokiaSystemSystemNameSpec struct {
	SrlNokiaSystemSystemName *SrlNokiaSystemSystemName `json:"name"`
}

// SrlNokiaSystemSystemNameStatus struct
type SrlNokiaSystemSystemNameStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaSystemSystemName is the Schema for the K8sSrlNokiaSystemSystemNames API
type K8sSrlNokiaSystemSystemName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaSystemSystemNameSpec   `json:"spec,omitempty"`
	Status SrlNokiaSystemSystemNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaSystemSystemNameList contains a list of K8sSrlNokiaSystemSystemNames
type K8sSrlNokiaSystemSystemNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaSystemSystemName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaSystemSystemName{}, &K8sSrlNokiaSystemSystemNameList{})
}
