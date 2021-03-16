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
	// SrlNokiaInterfacesInterfaceSubinterfaceFinalizer is the name of the finalizer added to
	// SrlNokiaInterfacesInterfaceSubinterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaInterfacesInterfaceSubinterfaceFinalizer string = "Subinterface.srlinux.henderiw.be"
)

// SrlNokiaInterfacesInterfaceSubinterfaceAclInput struct
type SrlNokiaInterfacesInterfaceSubinterfaceAclInput struct {
	Ipv4Filter *string `json:"ipv4-filter,omitempty"`
	Ipv6Filter *string `json:"ipv6-filter,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceAclOutput struct
type SrlNokiaInterfacesInterfaceSubinterfaceAclOutput struct {
	Ipv4Filter *string `json:"ipv4-filter,omitempty"`
	Ipv6Filter *string `json:"ipv6-filter,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceAcl struct
type SrlNokiaInterfacesInterfaceSubinterfaceAcl struct {
	Input  *SrlNokiaInterfacesInterfaceSubinterfaceAclInput  `json:"input,omitempty"`
	Output *SrlNokiaInterfacesInterfaceSubinterfaceAclOutput `json:"output,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceAnycastGw struct
type SrlNokiaInterfacesInterfaceSubinterfaceAnycastGw struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	AnycastGwMac *string `json:"anycast-gw-mac,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	VirtualRouterId *uint8 `json:"virtual-router-id,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacDuplication struct
type SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacDuplication struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`;`use-net-instance-action`
	// +kubebuilder:default:=use-net-instance-action
	Action *string `json:"action"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearningAging struct
type SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearningAging struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearning struct
type SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearning struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                                             `json:"admin-state"`
	Aging      *SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearningAging `json:"aging,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLimit struct
type SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLimit struct {
	// +kubebuilder:validation:Minimum=6
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	WarningThresholdPct *int32 `json:"warning-threshold-pct,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	MaximumEntries *int32 `json:"maximum-entries,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceBridgeTable struct
type SrlNokiaInterfacesInterfaceSubinterfaceBridgeTable struct {
	// +kubebuilder:default:=false
	DiscardUnknownSrcMac *bool                                                             `json:"discard-unknown-src-mac,omitempty"`
	MacDuplication       *SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacDuplication `json:"mac-duplication,omitempty"`
	MacLearning          *SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLearning    `json:"mac-learning,omitempty"`
	MacLimit             *SrlNokiaInterfacesInterfaceSubinterfaceBridgeTableMacLimit       `json:"mac-limit,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4Address struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4Address struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	IpPrefix  *string `json:"ip-prefix"`
	AnycastGw *bool   `json:"anycast-gw,omitempty"`
	Primary   *string `json:"primary,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpnAdvertise struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpnAdvertise struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	RouteType *string `json:"route-type"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AdminTag *uint32 `json:"admin-tag,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpn struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpn struct {
	Advertise []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpnAdvertise `json:"advertise,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoutePopulate struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoutePopulate struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	RouteType *string `json:"route-type"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	AdminTag *uint32 `json:"admin-tag,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoute struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoute struct {
	Populate []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoutePopulate `json:"populate,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpNeighbor struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	LinkLayerAddress *string `json:"link-layer-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4Address *string `json:"ipv4-address"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4Arp struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4Arp struct {
	HostRoute *SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpHostRoute `json:"host-route,omitempty"`
	// +kubebuilder:default:=false
	LearnUnsolicited *bool                                                     `json:"learn-unsolicited,omitempty"`
	Neighbor         []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpNeighbor `json:"neighbor,omitempty"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	Timeout *uint16 `json:"timeout,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Debug *string `json:"debug"`
	// +kubebuilder:default:=true
	DuplicateAddressDetection *bool                                               `json:"duplicate-address-detection,omitempty"`
	Evpn                      *SrlNokiaInterfacesInterfaceSubinterfaceIpv4ArpEvpn `json:"evpn,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClientTraceOptions struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClientTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClient struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClient struct {
	TraceOptions *SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClientTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelayTraceOptions struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelayTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelay struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelay struct {
	// +kubebuilder:default:=false
	UseGiAddrAsSrcIpAddr *bool `json:"use-gi-addr-as-src-ip-addr,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	GiAddress *string `json:"gi-address,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`circuit-id`;`remote-id`
	Option *string `json:"option"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server       *string                                                           `json:"server,omitempty"`
	TraceOptions *SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelayTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface struct {
	Interface *string `json:"interface"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	PriorityDecrement *uint8 `json:"priority-decrement,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking struct {
	TrackInterface []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTrackingTrackInterface `json:"track-interface,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics struct {
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroup struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroup struct {
	Authentication *SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	MasterInheritInterval *bool                                                               `json:"master-inherit-interval,omitempty"`
	Statistics            *SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupStatistics `json:"statistics,omitempty"`
	AcceptMode            *bool                                                               `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	AdvertiseInterval *uint16 `json:"advertise-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=2
	Version *uint8 `json:"version,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	VirtualRouterId *uint8 `json:"virtual-router-id"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState        *string                                                                    `json:"admin-state"`
	InterfaceTracking *SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroupInterfaceTracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	InitDelay *uint16 `json:"init-delay,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	OperInterval *uint16 `json:"oper-interval,omitempty"`
	Preempt      *bool   `json:"preempt,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	PreemptDelay *uint16 `json:"preempt-delay,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority *uint8 `json:"priority,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	VirtualAddress *string `json:"virtual-address,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4Vrrp struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4Vrrp struct {
	VrrpGroup []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4VrrpVrrpGroup `json:"vrrp-group,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv4 struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv4 struct {
	DhcpRelay *SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpRelay `json:"dhcp-relay,omitempty"`
	Vrrp      *SrlNokiaInterfacesInterfaceSubinterfaceIpv4Vrrp      `json:"vrrp,omitempty"`
	Address   []*SrlNokiaInterfacesInterfaceSubinterfaceIpv4Address `json:"address,omitempty"`
	// +kubebuilder:default:=false
	AllowDirectedBroadcast *bool                                                  `json:"allow-directed-broadcast,omitempty"`
	Arp                    *SrlNokiaInterfacesInterfaceSubinterfaceIpv4Arp        `json:"arp,omitempty"`
	DhcpClient             *SrlNokiaInterfacesInterfaceSubinterfaceIpv4DhcpClient `json:"dhcp-client,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6Address struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6Address struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix  *string `json:"ip-prefix"`
	AnycastGw *bool   `json:"anycast-gw,omitempty"`
	Primary   *string `json:"primary,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClientTraceOptions struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClientTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClient struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClient struct {
	TraceOptions *SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClientTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelayTraceOptions struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelayTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Trace *string `json:"trace"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelay struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelay struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`interface-id`;`remote-id`
	Option *string `json:"option"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	Server *string `json:"server,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	SourceAddress *string                                                           `json:"source-address,omitempty"`
	TraceOptions  *SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelayTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`dynamic`;`static`
	RouteType *string `json:"route-type"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AdminTag *uint32 `json:"admin-tag,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpn struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpn struct {
	Advertise []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpnAdvertise `json:"advertise,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`dynamic`;`evpn`;`static`
	RouteType *string `json:"route-type"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	AdminTag *uint32 `json:"admin-tag,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute struct {
	Populate []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoutePopulate `json:"populate,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6Address *string `json:"ipv6-address"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	LinkLayerAddress *string `json:"link-layer-address"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscovery struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscovery struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`both`;`global`;`link-local`;`none`
	// +kubebuilder:default:=none
	LearnUnsolicited *string                                                                 `json:"learn-unsolicited"`
	Neighbor         []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryNeighbor `json:"neighbor,omitempty"`
	// +kubebuilder:validation:Minimum=30
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=30
	ReachableTime *uint32 `json:"reachable-time,omitempty"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=14400
	StaleTime *uint32 `json:"stale-time,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Debug *string `json:"debug"`
	// +kubebuilder:default:=true
	DuplicateAddressDetection *bool                                                                  `json:"duplicate-address-detection,omitempty"`
	Evpn                      *SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryEvpn      `json:"evpn,omitempty"`
	HostRoute                 *SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscoveryHostRoute `json:"host-route,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix struct {
	// +kubebuilder:default:=604800
	PreferredLifetime *uint32 `json:"preferred-lifetime,omitempty"`
	// +kubebuilder:default:=2592000
	ValidLifetime *uint32 `json:"valid-lifetime,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Ipv6Prefix *string `json:"ipv6-prefix"`
	// +kubebuilder:default:=false
	AutonomousFlag *bool `json:"autonomous-flag,omitempty"`
	// +kubebuilder:default:=false
	OnLinkFlag *bool `json:"on-link-flag,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRole struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRole struct {
	Prefix []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRolePrefix `json:"prefix,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600000
	// +kubebuilder:default:=0
	ReachableTime *uint32 `json:"reachable-time,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1800000
	// +kubebuilder:default:=0
	RetransmitTime *uint32 `json:"retransmit-time,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=64
	CurrentHopLimit *uint8 `json:"current-hop-limit,omitempty"`
	// +kubebuilder:default:=false
	ManagedConfigurationFlag *bool `json:"managed-configuration-flag,omitempty"`
	// +kubebuilder:validation:Minimum=3
	// +kubebuilder:validation:Maximum=1350
	// +kubebuilder:default:=200
	MinAdvertisementInterval *uint16 `json:"min-advertisement-interval,omitempty"`
	// +kubebuilder:default:=false
	OtherConfigurationFlag *bool `json:"other-configuration-flag,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=9000
	// +kubebuilder:default:=1800
	RouterLifetime *uint16 `json:"router-lifetime,omitempty"`
	// +kubebuilder:validation:Minimum=1280
	// +kubebuilder:validation:Maximum=9486
	IpMtu *uint16 `json:"ip-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=600
	MaxAdvertisementInterval *uint16 `json:"max-advertisement-interval,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisement struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisement struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`messages`
	Debug      *string                                                                   `json:"debug"`
	RouterRole *SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisementRouterRole `json:"router-role,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	PriorityDecrement *uint8  `json:"priority-decrement,omitempty"`
	Interface         *string `json:"interface"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking struct {
	TrackInterface []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTrackingTrackInterface `json:"track-interface,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics struct {
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroup struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroup struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	InitDelay         *uint16                                                                    `json:"init-delay,omitempty"`
	InterfaceTracking *SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupInterfaceTracking `json:"interface-tracking,omitempty"`
	// +kubebuilder:default:=false
	MasterInheritInterval *bool `json:"master-inherit-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	VirtualRouterId *uint8 `json:"virtual-router-id"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=3
	Version *uint8 `json:"version,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1000
	AdvertiseInterval *uint16 `json:"advertise-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	OperInterval *uint16 `json:"oper-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=100
	Priority   *uint8                                                              `json:"priority,omitempty"`
	Statistics *SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupStatistics `json:"statistics,omitempty"`
	AcceptMode *bool                                                               `json:"accept-mode,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState     *string                                                                 `json:"admin-state"`
	Authentication *SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroupAuthentication `json:"authentication,omitempty"`
	Preempt        *bool                                                                   `json:"preempt,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	PreemptDelay *uint16 `json:"preempt-delay,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	VirtualAddress *string `json:"virtual-address,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6Vrrp struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6Vrrp struct {
	VrrpGroup []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6VrrpVrrpGroup `json:"vrrp-group,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceIpv6 struct
type SrlNokiaInterfacesInterfaceSubinterfaceIpv6 struct {
	Vrrp                *SrlNokiaInterfacesInterfaceSubinterfaceIpv6Vrrp                `json:"vrrp,omitempty"`
	Address             []*SrlNokiaInterfacesInterfaceSubinterfaceIpv6Address           `json:"address,omitempty"`
	DhcpClient          *SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpClient          `json:"dhcp-client,omitempty"`
	DhcpRelay           *SrlNokiaInterfacesInterfaceSubinterfaceIpv6DhcpRelay           `json:"dhcp-relay,omitempty"`
	NeighborDiscovery   *SrlNokiaInterfacesInterfaceSubinterfaceIpv6NeighborDiscovery   `json:"neighbor-discovery,omitempty"`
	RouterAdvertisement *SrlNokiaInterfacesInterfaceSubinterfaceIpv6RouterAdvertisement `json:"router-advertisement,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceLocalMirrorDestination struct
type SrlNokiaInterfacesInterfaceSubinterfaceLocalMirrorDestination struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceQosInputClassifiers struct
type SrlNokiaInterfacesInterfaceSubinterfaceQosInputClassifiers struct {
	Ipv4Dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6Dscp         *string `json:"ipv6-dscp,omitempty"`
	MplsTrafficClass *string `json:"mpls-traffic-class,omitempty"`
	Dscp             *string `json:"dscp,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceQosInput struct
type SrlNokiaInterfacesInterfaceSubinterfaceQosInput struct {
	Classifiers *SrlNokiaInterfacesInterfaceSubinterfaceQosInputClassifiers `json:"classifiers,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceQosOutputRewriteRules struct
type SrlNokiaInterfacesInterfaceSubinterfaceQosOutputRewriteRules struct {
	Dscp             *string `json:"dscp,omitempty"`
	Ipv4Dscp         *string `json:"ipv4-dscp,omitempty"`
	Ipv6Dscp         *string `json:"ipv6-dscp,omitempty"`
	MplsTrafficClass *string `json:"mpls-traffic-class,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceQosOutput struct
type SrlNokiaInterfacesInterfaceSubinterfaceQosOutput struct {
	RewriteRules *SrlNokiaInterfacesInterfaceSubinterfaceQosOutputRewriteRules `json:"rewrite-rules,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceQos struct
type SrlNokiaInterfacesInterfaceSubinterfaceQos struct {
	Output *SrlNokiaInterfacesInterfaceSubinterfaceQosOutput `json:"output,omitempty"`
	Input  *SrlNokiaInterfacesInterfaceSubinterfaceQosInput  `json:"input,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapSingleTagged struct
type SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapSingleTagged struct {
	VlanId *string `json:"vlan-id,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapUntagged struct
type SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapUntagged struct {
}

// SrlNokiaInterfacesInterfaceSubinterfaceVlanEncap struct
type SrlNokiaInterfacesInterfaceSubinterfaceVlanEncap struct {
	SingleTagged *SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapSingleTagged `json:"single-tagged,omitempty"`
	Untagged     *SrlNokiaInterfacesInterfaceSubinterfaceVlanEncapUntagged     `json:"untagged,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceVlan struct
type SrlNokiaInterfacesInterfaceSubinterfaceVlan struct {
	Encap *SrlNokiaInterfacesInterfaceSubinterfaceVlanEncap `json:"encap,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterface struct
type SrlNokiaInterfacesInterfaceSubinterface struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState             *string                                                        `json:"admin-state"`
	AnycastGw              *SrlNokiaInterfacesInterfaceSubinterfaceAnycastGw              `json:"anycast-gw,omitempty"`
	Ipv4                   *SrlNokiaInterfacesInterfaceSubinterfaceIpv4                   `json:"ipv4,omitempty"`
	LocalMirrorDestination *SrlNokiaInterfacesInterfaceSubinterfaceLocalMirrorDestination `json:"local-mirror-destination,omitempty"`
	Acl                    *SrlNokiaInterfacesInterfaceSubinterfaceAcl                    `json:"acl,omitempty"`
	Ipv6                   *SrlNokiaInterfacesInterfaceSubinterfaceIpv6                   `json:"ipv6,omitempty"`
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	L2Mtu *uint16                                     `json:"l2-mtu,omitempty"`
	Qos   *SrlNokiaInterfacesInterfaceSubinterfaceQos `json:"qos,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=9999
	Index       *uint32                                             `json:"index"`
	BridgeTable *SrlNokiaInterfacesInterfaceSubinterfaceBridgeTable `json:"bridge-table,omitempty"`
	// +kubebuilder:validation:Minimum=1280
	// +kubebuilder:validation:Maximum=9486
	IpMtu *uint16                                      `json:"ip-mtu,omitempty"`
	Vlan  *SrlNokiaInterfacesInterfaceSubinterfaceVlan `json:"vlan,omitempty"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceSpec struct
type SrlNokiaInterfacesInterfaceSubinterfaceSpec struct {
	SrlNokiaInterfaceName                   *string                                    `json:"interface-name"`
	SrlNokiaInterfacesInterfaceSubinterface *[]SrlNokiaInterfacesInterfaceSubinterface `json:"subinterface"`
}

// SrlNokiaInterfacesInterfaceSubinterfaceStatus struct
type SrlNokiaInterfacesInterfaceSubinterfaceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaInterfacesInterfaceSubinterface is the Schema for the K8sSrlNokiaInterfacesInterfaceSubinterfaces API
type K8sSrlNokiaInterfacesInterfaceSubinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaInterfacesInterfaceSubinterfaceSpec   `json:"spec,omitempty"`
	Status SrlNokiaInterfacesInterfaceSubinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaInterfacesInterfaceSubinterfaceList contains a list of K8sSrlNokiaInterfacesInterfaceSubinterfaces
type K8sSrlNokiaInterfacesInterfaceSubinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaInterfacesInterfaceSubinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaInterfacesInterfaceSubinterface{}, &K8sSrlNokiaInterfacesInterfaceSubinterfaceList{})
}
