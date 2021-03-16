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
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxFinalizer string = "Linux.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct {
	// +kubebuilder:default:=true
	ExportNeighbors *bool `json:"export-neighbors,omitempty"`
	// +kubebuilder:default:=false
	ExportRoutes *bool `json:"export-routes,omitempty"`
	// +kubebuilder:default:=false
	ImportRoutes *bool `json:"import-routes,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec struct {
	SrlNokiaNetworkInstanceName                          *string                                               `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux *SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux `json:"linux"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxs API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxs
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList{})
}
