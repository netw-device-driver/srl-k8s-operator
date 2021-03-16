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
	// SrlNokiaTunnelInterfacesTunnelInterfaceFinalizer is the name of the finalizer added to
	// SrlNokiaTunnelInterfacesTunnelInterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaTunnelInterfacesTunnelInterfaceFinalizer string = "TunnelInterface.srlinux.henderiw.be"
)

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable struct {
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	DestinationMac *string `json:"destination-mac,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState          *string                                                                                                          `json:"admin-state,omitempty"`
	InnerEthernetHeader *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                                                                                         `json:"admin-state,omitempty"`
	Destination []*SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination `json:"destination,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups struct {
	Group []*SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup `json:"group,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader struct {
	// +kubebuilder:default:=use-system-mac
	SourceMac *string `json:"source-mac,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress struct {
	DestinationGroups   *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups   `json:"destination-groups,omitempty"`
	InnerEthernetHeader *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:default:=use-system-ipv4-address
	SourceIp *string `json:"source-ip,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=99999999
	Index       *uint32                                                           `json:"index"`
	BridgeTable *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable `json:"bridge-table,omitempty"`
	Egress      *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress      `json:"egress,omitempty"`
	Ingress     *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress     `json:"ingress,omitempty"`
	Type        *string                                                           `json:"type"`
}

// SrlNokiaTunnelInterfacesTunnelInterface struct
type SrlNokiaTunnelInterfacesTunnelInterface struct {
	// +kubebuilder:validation:MinLength=6
	// +kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9]))`
	Name           *string                                                  `json:"name"`
	VxlanInterface []*SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface `json:"vxlan-interface,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceSpec struct
type SrlNokiaTunnelInterfacesTunnelInterfaceSpec struct {
	SrlNokiaTunnelInterfacesTunnelInterface *[]SrlNokiaTunnelInterfacesTunnelInterface `json:"tunnel-interface"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceStatus struct
type SrlNokiaTunnelInterfacesTunnelInterfaceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaTunnelInterfacesTunnelInterface is the Schema for the K8sSrlNokiaTunnelInterfacesTunnelInterfaces API
type K8sSrlNokiaTunnelInterfacesTunnelInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaTunnelInterfacesTunnelInterfaceSpec   `json:"spec,omitempty"`
	Status SrlNokiaTunnelInterfacesTunnelInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaTunnelInterfacesTunnelInterfaceList contains a list of K8sSrlNokiaTunnelInterfacesTunnelInterfaces
type K8sSrlNokiaTunnelInterfacesTunnelInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaTunnelInterfacesTunnelInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaTunnelInterfacesTunnelInterface{}, &K8sSrlNokiaTunnelInterfacesTunnelInterfaceList{})
}
