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
	// SrlNokiaSystemSystemNetworkInstanceFinalizer is the name of the finalizer added to
	// SrlNokiaSystemSystemNetworkInstance to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaSystemSystemNetworkInstanceFinalizer string = "NetworkInstance.srlinux.henderiw.be"
)

// SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher struct {
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget struct {
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstance struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstance struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Id                 *uint8                                                                           `json:"id"`
	RouteDistinguisher *SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher `json:"route-distinguisher,omitempty"`
	RouteTarget        *SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget        `json:"route-target,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpn struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpn struct {
	BgpInstance []*SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlgCapabilities struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlg struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlg struct {
	Capabilities *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlgCapabilities `json:"capabilities,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlgCapabilities struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
	// +kubebuilder:default:=false
	NonRevertive *bool `json:"non-revertive,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlg struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlg struct {
	Capabilities *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlgCapabilities `json:"capabilities,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=32767
	PreferenceValue *uint32 `json:"preference-value,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithm struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithm struct {
	DefaultAlg    *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmDefaultAlg    `json:"default-alg,omitempty"`
	PreferenceAlg *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithmPreferenceAlg `json:"preference-alg,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`default`;`preference`
	// +kubebuilder:default:=default
	Type *string `json:"type"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionTimers struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ActivationTimer *uint32 `json:"activation-timer,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElection struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElection struct {
	Algorithm                        *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionAlgorithm `json:"algorithm,omitempty"`
	InterfaceStandbySignalingOnNonDf *bool                                                                                                          `json:"interface-standby-signaling-on-non-df,omitempty"`
	Timers                           *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElectionTimers    `json:"timers,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutesEthernetSegment struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutesEthernetSegment struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	OriginatingIp *string `json:"originating-ip"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutes struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutes struct {
	EthernetSegment *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutesEthernetSegment `json:"ethernet-segment,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop *string `json:"next-hop"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegment struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegment struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`all-active`;`single-active`
	// +kubebuilder:default:=all-active
	MultiHomingMode *string                                                                                           `json:"multi-homing-mode"`
	Routes          *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentRoutes `json:"routes,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string                                                                                               `json:"admin-state"`
	DfElection *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegmentDfElection `json:"df-election,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi       *string `json:"esi,omitempty"`
	Interface *string `json:"interface,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstance struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstance struct {
	Id              *string                                                                                       `json:"id"`
	EthernetSegment []*SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstanceEthernetSegment `json:"ethernet-segment,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsTimers struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	ActivationTimer *uint32 `json:"activation-timer,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=6000
	// +kubebuilder:default:=10
	BootTimer *uint32 `json:"boot-timer,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegments struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegments struct {
	BgpInstance []*SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsBgpInstance `json:"bgp-instance,omitempty"`
	Timers      *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegmentsTimers        `json:"timers,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocolsEvpn struct
type SrlNokiaSystemSystemNetworkInstanceProtocolsEvpn struct {
	EthernetSegments *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpnEthernetSegments `json:"ethernet-segments,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceProtocols struct
type SrlNokiaSystemSystemNetworkInstanceProtocols struct {
	BgpVpn *SrlNokiaSystemSystemNetworkInstanceProtocolsBgpVpn `json:"bgp-vpn,omitempty"`
	Evpn   *SrlNokiaSystemSystemNetworkInstanceProtocolsEvpn   `json:"evpn,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstance struct
type SrlNokiaSystemSystemNetworkInstance struct {
	Protocols *SrlNokiaSystemSystemNetworkInstanceProtocols `json:"protocols,omitempty"`
}

// SrlNokiaSystemSystemNetworkInstanceSpec struct
type SrlNokiaSystemSystemNetworkInstanceSpec struct {
	SrlNokiaSystemSystemNetworkInstance *SrlNokiaSystemSystemNetworkInstance `json:"network-instance"`
}

// SrlNokiaSystemSystemNetworkInstanceStatus struct
type SrlNokiaSystemSystemNetworkInstanceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaSystemSystemNetworkInstance is the Schema for the K8sSrlNokiaSystemSystemNetworkInstances API
type K8sSrlNokiaSystemSystemNetworkInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaSystemSystemNetworkInstanceSpec   `json:"spec,omitempty"`
	Status SrlNokiaSystemSystemNetworkInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaSystemSystemNetworkInstanceList contains a list of K8sSrlNokiaSystemSystemNetworkInstances
type K8sSrlNokiaSystemSystemNetworkInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaSystemSystemNetworkInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaSystemSystemNetworkInstance{}, &K8sSrlNokiaSystemSystemNetworkInstanceList{})
}
