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
	// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutes to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesFinalizer string = "StaticRoutes.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesRoute struct
type SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesRoute struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8 `json:"preference,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=1
	Metric       *uint32 `json:"metric,omitempty"`
	NextHopGroup *string `json:"next-hop-group,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutes struct
type SrlNokiaNetworkInstanceNetworkInstanceStaticRoutes struct {
	Route []*SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesRoute `json:"route,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesSpec struct {
	SrlNokiaNetworkInstanceName                        *string                                             `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceStaticRoutes *SrlNokiaNetworkInstanceNetworkInstanceStaticRoutes `json:"static-routes"`
}

// SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutess API
type K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceStaticRoutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutesList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutess
type K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutesList{})
}