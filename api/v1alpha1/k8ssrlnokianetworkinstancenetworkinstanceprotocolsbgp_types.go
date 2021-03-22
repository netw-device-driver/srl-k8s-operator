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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpFinalizer string = "Bgp.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	// +kubebuilder:default:=disabled
	Mode *string `json:"mode,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptions struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AllowOwnAs      *uint8                                                                          `json:"allow-own-as,omitempty"`
	RemovePrivateAs *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpConvergence struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MinWaitToAdvertise *uint16 `json:"min-wait-to-advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAcceptMatch struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAcceptMatch struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([1-9][0-9]*)|([1-9][0-9]*)\.\.([1-9][0-9]*)`
	AllowedPeerAs *string `json:"allowed-peer-as,omitempty"`
	PeerGroup     *string `json:"peer-group"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAccept struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAccept struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	MaxSessions *uint16                                                                          `json:"max-sessions,omitempty"`
	Match       []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAcceptMatch `json:"match,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighbors struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighbors struct {
	Accept *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighborsAccept `json:"accept,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEbgpDefaultPolicy struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEbgpDefaultPolicy struct {
	// +kubebuilder:default:=true
	ExportRejectAll *bool `json:"export-reject-all,omitempty"`
	// +kubebuilder:default:=true
	ImportRejectAll *bool `json:"import-reject-all,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpnMultipath struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpnMultipath struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
	// +kubebuilder:default:=true
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpn struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpn struct {
	KeepAllRoutes *bool                                                            `json:"keep-all-routes,omitempty"`
	Multipath     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpnMultipath `json:"multipath,omitempty"`
	// +kubebuilder:default:=false
	RapidUpdate *bool `json:"rapid-update,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseIpv6NextHops *bool `json:"advertise-ipv6-next-hops,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpFailureDetection struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpFailureDetection struct {
	// +kubebuilder:default:=true
	FastFailover *bool `json:"fast-failover,omitempty"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGracefulRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=360
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptions struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                                               `json:"allow-own-as,omitempty"`
	RemovePrivateAs *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                                                `json:"replace-peer-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpnPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpnPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpn struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                                 `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                                   `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupFailureDetection struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupFailureDetection struct {
	FastFailover *bool `json:"fast-failover,omitempty"`
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupGracefulRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupGracefulRestart struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4Unicast struct {
	ReceiveIpv6NextHops *bool `json:"receive-ipv6-next-hops,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                                        `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                                          `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                                        `json:"admin-state,omitempty"`
	PrefixLimit *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupLocalAs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupLocalAs struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber *uint32 `json:"as-number"`
	// +kubebuilder:default:=true
	PrependGlobalAs *bool `json:"prepend-global-as,omitempty"`
	// +kubebuilder:default:=true
	PrependLocalAs *bool `json:"prepend-local-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupRouteReflector struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendCommunity struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendDefaultRoute struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendDefaultRoute struct {
	ExportPolicy *string `json:"export-policy,omitempty"`
	// +kubebuilder:default:=false
	Ipv4Unicast *bool `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:default:=false
	Ipv6Unicast *bool `json:"ipv6-unicast,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTimers struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=120
	ConnectRetry *uint16 `json:"connect-retry,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=90
	HoldTime *uint16 `json:"hold-time,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=21845
	KeepaliveInterval *uint16 `json:"keepalive-interval,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	MinimumAdvertisementInterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptionsFlag struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptions struct {
	Flag []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptionsFlag `json:"flag,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTransport struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTransport struct {
	// +kubebuilder:default:=false
	PassiveMode *bool `json:"passive-mode,omitempty"`
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroup struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	GroupName      *string                                                                `json:"group-name"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Description     *string                                                                 `json:"description,omitempty"`
	Evpn            *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupEvpn            `json:"evpn,omitempty"`
	ExportPolicy    *string                                                                 `json:"export-policy,omitempty"`
	GracefulRestart *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupGracefulRestart `json:"graceful-restart,omitempty"`
	Ipv4Unicast     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv4Unicast     `json:"ipv4-unicast,omitempty"`
	RouteReflector  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupRouteReflector  `json:"route-reflector,omitempty"`
	Transport       *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTransport       `json:"transport,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState    *string                                                               `json:"admin-state,omitempty"`
	AsPathOptions *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupAsPathOptions `json:"as-path-options,omitempty"`
	LocalAs       []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupLocalAs     `json:"local-as,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	PeerAs           *uint32                                                                  `json:"peer-as,omitempty"`
	SendDefaultRoute *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendDefaultRoute `json:"send-default-route,omitempty"`
	ImportPolicy     *string                                                                  `json:"import-policy,omitempty"`
	SendCommunity    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupSendCommunity    `json:"send-community,omitempty"`
	Timers           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTimers           `json:"timers,omitempty"`
	TraceOptions     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupTraceOptions     `json:"trace-options,omitempty"`
	FailureDetection *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupFailureDetection `json:"failure-detection,omitempty"`
	Ipv6Unicast      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroupIpv6Unicast      `json:"ipv6-unicast,omitempty"`
	// +kubebuilder:default:=false
	NextHopSelf *bool `json:"next-hop-self,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastConvergence struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastMultipath struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastMultipath struct {
	// +kubebuilder:default:=true
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4Unicast struct {
	Multipath *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastMultipath `json:"multipath,omitempty"`
	// +kubebuilder:default:=false
	ReceiveIpv6NextHops *bool `json:"receive-ipv6-next-hops,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseIpv6NextHops *bool                                                                     `json:"advertise-ipv6-next-hops,omitempty"`
	Convergence           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4UnicastConvergence `json:"convergence,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastConvergence struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastMultipath struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastMultipath struct {
	// +kubebuilder:default:=true
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState  *string                                                                   `json:"admin-state,omitempty"`
	Convergence *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastConvergence `json:"convergence,omitempty"`
	Multipath   *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6UnicastMultipath   `json:"multipath,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptions struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                                                  `json:"allow-own-as,omitempty"`
	RemovePrivateAs *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                                                   `json:"replace-peer-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpnPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpnPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpn struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                                    `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                                      `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborFailureDetection struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16                                                                               `json:"stale-routes-time,omitempty"`
	WarmRestart     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestartWarmRestart `json:"warm-restart,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                                           `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                                             `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
	ReceiveIpv6NextHops   *bool                                                                             `json:"receive-ipv6-next-hops,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                                           `json:"admin-state,omitempty"`
	PrefixLimit *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborLocalAs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborLocalAs struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber        *uint32 `json:"as-number"`
	PrependGlobalAs *bool   `json:"prepend-global-as,omitempty"`
	PrependLocalAs  *bool   `json:"prepend-local-as,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborRouteReflector struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendCommunity struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendCommunity struct {
	Standard *bool `json:"standard,omitempty"`
	Large    *bool `json:"large,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendDefaultRoute struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendDefaultRoute struct {
	ExportPolicy *string `json:"export-policy,omitempty"`
	Ipv4Unicast  *bool   `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast  *bool   `json:"ipv6-unicast,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	HoldTime *uint16 `json:"hold-time,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=21845
	KeepaliveInterval *uint16 `json:"keepalive-interval,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	MinimumAdvertisementInterval *uint16 `json:"minimum-advertisement-interval,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	ConnectRetry *uint16 `json:"connect-retry,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptionsFlag struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptions struct {
	Flag []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptionsFlag `json:"flag,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTransport struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTransport struct {
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	PassiveMode  *bool   `json:"passive-mode,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighbor struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighbor struct {
	GracefulRestart *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborGracefulRestart `json:"graceful-restart,omitempty"`
	ImportPolicy    *string                                                                    `json:"import-policy,omitempty"`
	LocalAs         []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborLocalAs       `json:"local-as,omitempty"`
	PeerGroup       *string                                                                    `json:"peer-group"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%!+(MISSING))?`
	// +kubebuilder:validation:Pattern=`(([^:]+:){6}(([^:]+:[^:]+)|(.*\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%!+(MISSING))?`
	// +kubebuilder:validation:Pattern=`([^%!](MISSING)+)(%!((MISSING)mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3})))?`
	PeerAddress      *string                                                                     `json:"peer-address"`
	AsPathOptions    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAsPathOptions    `json:"as-path-options,omitempty"`
	RouteReflector   *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborRouteReflector   `json:"route-reflector,omitempty"`
	Timers           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTimers           `json:"timers,omitempty"`
	SendCommunity    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendCommunity    `json:"send-community,omitempty"`
	SendDefaultRoute *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborSendDefaultRoute `json:"send-default-route,omitempty"`
	Transport        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTransport        `json:"transport,omitempty"`
	FailureDetection *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborFailureDetection `json:"failure-detection,omitempty"`
	Ipv4Unicast      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv4Unicast      `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	NextHopSelf     *bool   `json:"next-hop-self,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	PeerAs       *uint32                                                                 `json:"peer-as,omitempty"`
	Ipv6Unicast  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborIpv6Unicast  `json:"ipv6-unicast,omitempty"`
	TraceOptions *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborTraceOptions `json:"trace-options,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState     *string                                                                   `json:"admin-state,omitempty"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Description  *string                                                         `json:"description,omitempty"`
	Evpn         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighborEvpn `json:"evpn,omitempty"`
	ExportPolicy *string                                                         `json:"export-policy,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpPreference struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ebgp *uint8 `json:"ebgp,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ibgp *uint8 `json:"ibgp,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteAdvertisement struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteAdvertisement struct {
	// +kubebuilder:default:=false
	RapidWithdrawal *bool `json:"rapid-withdrawal,omitempty"`
	// +kubebuilder:default:=true
	WaitForFibInstall *bool `json:"wait-for-fib-install,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteReflector struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteReflector struct {
	// +kubebuilder:default:=false
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSendCommunity struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSendCommunity struct {
	// +kubebuilder:default:=true
	Large *bool `json:"large,omitempty"`
	// +kubebuilder:default:=true
	Standard *bool `json:"standard,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptionsFlag struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptions struct {
	Flag []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptionsFlag `json:"flag,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTransport struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTransport struct {
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	// +kubebuilder:default:=1024
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp struct {
	ImportPolicy       *string                                                               `json:"import-policy,omitempty"`
	Neighbor           []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpNeighbor         `json:"neighbor,omitempty"`
	Preference         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpPreference         `json:"preference,omitempty"`
	RouteAdvertisement *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteAdvertisement `json:"route-advertisement,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AutonomousSystem *uint32                                                        `json:"autonomous-system"`
	Convergence      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpConvergence `json:"convergence,omitempty"`
	Ipv6Unicast      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv6Unicast `json:"ipv6-unicast,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RouterId      *string                                                          `json:"router-id"`
	SendCommunity *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSendCommunity `json:"send-community,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState       *string                                                             `json:"admin-state,omitempty"`
	DynamicNeighbors *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpDynamicNeighbors `json:"dynamic-neighbors,omitempty"`
	FailureDetection *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpFailureDetection `json:"failure-detection,omitempty"`
	GracefulRestart  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGracefulRestart  `json:"graceful-restart,omitempty"`
	Group            []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpGroup          `json:"group,omitempty"`
	AsPathOptions    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAsPathOptions    `json:"as-path-options,omitempty"`
	ExportPolicy     *string                                                             `json:"export-policy,omitempty"`
	Evpn             *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEvpn             `json:"evpn,omitempty"`
	Ipv4Unicast      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpIpv4Unicast      `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=100
	LocalPreference   *uint32                                                              `json:"local-preference,omitempty"`
	RouteReflector    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpRouteReflector    `json:"route-reflector,omitempty"`
	TraceOptions      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTraceOptions      `json:"trace-options,omitempty"`
	Transport         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpTransport         `json:"transport,omitempty"`
	Authentication    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpAuthentication    `json:"authentication,omitempty"`
	EbgpDefaultPolicy *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpEbgpDefaultPolicy `json:"ebgp-default-policy,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSpec struct {
	SrlNokiaNetworkInstanceName                        *string                                             `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp `json:"bgp"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgps API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgps
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp",
			Namespace:  o.Namespace,
			Name:       o.Name,
			UID:        o.UID,
			APIVersion: GroupVersion.String(),
		},
		Reason:  reason,
		Message: message,
		Source: corev1.EventSource{
			Component: "srl-controller",
		},
		FirstTimestamp:      t,
		LastTimestamp:       t,
		Count:               1,
		Type:                corev1.EventTypeNormal,
		ReportingController: "srlinux.henderiw.be/srl-controller",
	}
}

func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
