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
	// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroups to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsFinalizer string = "NextHopGroups.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupBlackhole struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupBlackhole struct {
	// +kubebuilder:default:=false
	GenerateIcmp *bool `json:"generate-icmp,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetectionEnableBfd struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetectionEnableBfd struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16384
	RemoteDiscriminator *uint32 `json:"remote-discriminator,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16384
	LocalDiscriminator *uint32 `json:"local-discriminator,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetection struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetection struct {
	EnableBfd *SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetectionEnableBfd `json:"enable-bfd,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthop struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthop struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState       *string                                                                          `json:"admin-state"`
	FailureDetection *SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthopFailureDetection `json:"failure-detection,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	IpAddress            *string `json:"ip-address,omitempty"`
	PushedMplsLabelStack *string `json:"pushed-mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Resolve *bool `json:"resolve,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroup struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                                            `json:"admin-state"`
	Blackhole  *SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupBlackhole `json:"blackhole,omitempty"`
	Nexthop    []*SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroupNexthop `json:"nexthop,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroups struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroups struct {
	Group []*SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsGroup `json:"group,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsSpec struct {
	SrlNokiaNetworkInstanceName                         *string                                              `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceNextHopGroups *SrlNokiaNetworkInstanceNetworkInstanceNextHopGroups `json:"next-hop-groups"`
}

// SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupss API
type K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupss
type K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsList{})
}
