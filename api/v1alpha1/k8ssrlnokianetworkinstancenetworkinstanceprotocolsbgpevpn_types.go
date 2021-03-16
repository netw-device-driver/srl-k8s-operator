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
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnFinalizer string = "BgpEvpn.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	OriginatingIp *string `json:"originating-ip,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct {
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop        *string                                                                                           `json:"next-hop,omitempty"`
	InclusiveMcast *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast `json:"inclusive-mcast,omitempty"`
	MacIp          *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp          `json:"mac-ip,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct {
	// +kubebuilder:default:=false
	AdvertiseGatewayMac *bool `json:"advertise-gateway-mac,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct {
	MacIp *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp `json:"mac-ip,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutes struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutes struct {
	BridgeTable *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable `json:"bridge-table,omitempty"`
	RouteTable  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable  `json:"route-table,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstance struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstance struct {
	Id *string `json:"id"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	DefaultAdminTag *uint32 `json:"default-admin-tag,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=1
	Ecmp *uint8 `json:"ecmp,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`vxlan`
	// +kubebuilder:default:=vxlan
	EncapsulationType *string `json:"encapsulation-type"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Evi            *uint32                                                                  `json:"evi"`
	Routes         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstanceRoutes `json:"routes,omitempty"`
	VxlanInterface *string                                                                  `json:"vxlan-interface,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn struct {
	BgpInstance []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnSpec struct {
	SrlNokiaNetworkInstanceName                            *string                                                 `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn `json:"bgp-evpn"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpns API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpns
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnList{})
}
