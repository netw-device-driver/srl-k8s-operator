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
	// SrlNetworkinstanceProtocolsBgpFinalizer is the name of the finalizer added to
	// SrlNetworkinstanceProtocolsBgp to block delete operations until the physical node can be
	// deprovisioned.
	SrlNetworkinstanceProtocolsBgpFinalizer string = "Bgp.srlinux.henderiw.be"
)

// NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	// +kubebuilder:default:=disabled
	Mode *string `json:"mode,omitempty"`
}

// NetworkinstanceProtocolsBgpAsPathOptions struct
type NetworkinstanceProtocolsBgpAsPathOptions struct {
	RemovePrivateAs *NetworkinstanceProtocolsBgpAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=0
	AllowOwnAs *uint8 `json:"allow-own-as,omitempty"`
}

// NetworkinstanceProtocolsBgpAuthentication struct
type NetworkinstanceProtocolsBgpAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpConvergence struct
type NetworkinstanceProtocolsBgpConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MinWaitToAdvertise *uint16 `json:"min-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch struct
type NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([1-9][0-9]*)|([1-9][0-9]*)\.\.([1-9][0-9]*)`
	AllowedPeerAs *string `json:"allowed-peer-as,omitempty"`
	PeerGroup     *string `json:"peer-group"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
}

// NetworkinstanceProtocolsBgpDynamicNeighborsAccept struct
type NetworkinstanceProtocolsBgpDynamicNeighborsAccept struct {
	Match []*NetworkinstanceProtocolsBgpDynamicNeighborsAcceptMatch `json:"match,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=0
	MaxSessions *uint16 `json:"max-sessions,omitempty"`
}

// NetworkinstanceProtocolsBgpDynamicNeighbors struct
type NetworkinstanceProtocolsBgpDynamicNeighbors struct {
	Accept *NetworkinstanceProtocolsBgpDynamicNeighborsAccept `json:"accept,omitempty"`
}

// NetworkinstanceProtocolsBgpEbgpDefaultPolicy struct
type NetworkinstanceProtocolsBgpEbgpDefaultPolicy struct {
	// +kubebuilder:default:=true
	ExportRejectAll *bool `json:"export-reject-all,omitempty"`
	// +kubebuilder:default:=true
	ImportRejectAll *bool `json:"import-reject-all,omitempty"`
}

// NetworkinstanceProtocolsBgpEvpnMultipath struct
type NetworkinstanceProtocolsBgpEvpnMultipath struct {
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

// NetworkinstanceProtocolsBgpEvpn struct
type NetworkinstanceProtocolsBgpEvpn struct {
	// +kubebuilder:default:=false
	RapidUpdate *bool `json:"rapid-update,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseIpv6NextHops *bool                                     `json:"advertise-ipv6-next-hops,omitempty"`
	KeepAllRoutes         *bool                                     `json:"keep-all-routes,omitempty"`
	Multipath             *NetworkinstanceProtocolsBgpEvpnMultipath `json:"multipath,omitempty"`
}

// NetworkinstanceProtocolsBgpFailureDetection struct
type NetworkinstanceProtocolsBgpFailureDetection struct {
	// +kubebuilder:default:=true
	FastFailover *bool `json:"fast-failover,omitempty"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// NetworkinstanceProtocolsBgpGracefulRestart struct
type NetworkinstanceProtocolsBgpGracefulRestart struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=360
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
}

// NetworkinstanceProtocolsBgpGroupAsPathOptions struct
type NetworkinstanceProtocolsBgpGroupAsPathOptions struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	AllowOwnAs      *uint8                                                        `json:"allow-own-as,omitempty"`
	RemovePrivateAs *NetworkinstanceProtocolsBgpGroupAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                         `json:"replace-peer-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupAuthentication struct
type NetworkinstanceProtocolsBgpGroupAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupEvpn struct
type NetworkinstanceProtocolsBgpGroupEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                          `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                            `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpGroupEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupFailureDetection struct
type NetworkinstanceProtocolsBgpGroupFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupGracefulRestart struct
type NetworkinstanceProtocolsBgpGroupGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16 `json:"stale-routes-time,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv4Unicast struct
type NetworkinstanceProtocolsBgpGroupIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                 `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                   `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpGroupIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
	ReceiveIpv6NextHops   *bool                                                   `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=90
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupIpv6Unicast struct
type NetworkinstanceProtocolsBgpGroupIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                 `json:"admin-state,omitempty"`
	PrefixLimit *NetworkinstanceProtocolsBgpGroupIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupLocalAs struct
type NetworkinstanceProtocolsBgpGroupLocalAs struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber *uint32 `json:"as-number"`
	// +kubebuilder:default:=true
	PrependGlobalAs *bool `json:"prepend-global-as,omitempty"`
	// +kubebuilder:default:=true
	PrependLocalAs *bool `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupRouteReflector struct
type NetworkinstanceProtocolsBgpGroupRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSendCommunity struct
type NetworkinstanceProtocolsBgpGroupSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupSendDefaultRoute struct
type NetworkinstanceProtocolsBgpGroupSendDefaultRoute struct {
	// +kubebuilder:default:=false
	Ipv4Unicast *bool `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:default:=false
	Ipv6Unicast  *bool   `json:"ipv6-unicast,omitempty"`
	ExportPolicy *string `json:"export-policy,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTimers struct
type NetworkinstanceProtocolsBgpGroupTimers struct {
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

// NetworkinstanceProtocolsBgpGroupTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpGroupTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
}

// NetworkinstanceProtocolsBgpGroupTraceOptions struct
type NetworkinstanceProtocolsBgpGroupTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpGroupTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpGroupTransport struct
type NetworkinstanceProtocolsBgpGroupTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	// +kubebuilder:default:=false
	PassiveMode *bool `json:"passive-mode,omitempty"`
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpGroup struct
type NetworkinstanceProtocolsBgpGroup struct {
	TraceOptions    *NetworkinstanceProtocolsBgpGroupTraceOptions    `json:"trace-options,omitempty"`
	ExportPolicy    *string                                          `json:"export-policy,omitempty"`
	GracefulRestart *NetworkinstanceProtocolsBgpGroupGracefulRestart `json:"graceful-restart,omitempty"`
	ImportPolicy    *string                                          `json:"import-policy,omitempty"`
	RouteReflector  *NetworkinstanceProtocolsBgpGroupRouteReflector  `json:"route-reflector,omitempty"`
	Timers          *NetworkinstanceProtocolsBgpGroupTimers          `json:"timers,omitempty"`
	Ipv4Unicast     *NetworkinstanceProtocolsBgpGroupIpv4Unicast     `json:"ipv4-unicast,omitempty"`
	LocalAs         []*NetworkinstanceProtocolsBgpGroupLocalAs       `json:"local-as,omitempty"`
	// +kubebuilder:default:=false
	NextHopSelf   *bool                                          `json:"next-hop-self,omitempty"`
	SendCommunity *NetworkinstanceProtocolsBgpGroupSendCommunity `json:"send-community,omitempty"`
	Transport     *NetworkinstanceProtocolsBgpGroupTransport     `json:"transport,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState    *string                                        `json:"admin-state,omitempty"`
	AsPathOptions *NetworkinstanceProtocolsBgpGroupAsPathOptions `json:"as-path-options,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	PeerAs           *uint32                                           `json:"peer-as,omitempty"`
	SendDefaultRoute *NetworkinstanceProtocolsBgpGroupSendDefaultRoute `json:"send-default-route,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	GroupName        *string                                           `json:"group-name"`
	Authentication   *NetworkinstanceProtocolsBgpGroupAuthentication   `json:"authentication,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpGroupEvpn             `json:"evpn,omitempty"`
	FailureDetection *NetworkinstanceProtocolsBgpGroupFailureDetection `json:"failure-detection,omitempty"`
	Ipv6Unicast      *NetworkinstanceProtocolsBgpGroupIpv6Unicast      `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4UnicastConvergence struct
type NetworkinstanceProtocolsBgpIpv4UnicastConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4UnicastMultipath struct
type NetworkinstanceProtocolsBgpIpv4UnicastMultipath struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel2 *uint32 `json:"max-paths-level-2,omitempty"`
	// +kubebuilder:default:=true
	AllowMultipleAs *bool `json:"allow-multiple-as,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPathsLevel1 *uint32 `json:"max-paths-level-1,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv4Unicast struct
type NetworkinstanceProtocolsBgpIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseIpv6NextHops *bool                                              `json:"advertise-ipv6-next-hops,omitempty"`
	Convergence           *NetworkinstanceProtocolsBgpIpv4UnicastConvergence `json:"convergence,omitempty"`
	Multipath             *NetworkinstanceProtocolsBgpIpv4UnicastMultipath   `json:"multipath,omitempty"`
	// +kubebuilder:default:=false
	ReceiveIpv6NextHops *bool `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6UnicastConvergence struct
type NetworkinstanceProtocolsBgpIpv6UnicastConvergence struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:default:=0
	MaxWaitToAdvertise *uint16 `json:"max-wait-to-advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpIpv6UnicastMultipath struct
type NetworkinstanceProtocolsBgpIpv6UnicastMultipath struct {
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

// NetworkinstanceProtocolsBgpIpv6Unicast struct
type NetworkinstanceProtocolsBgpIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState  *string                                            `json:"admin-state,omitempty"`
	Convergence *NetworkinstanceProtocolsBgpIpv6UnicastConvergence `json:"convergence,omitempty"`
	Multipath   *NetworkinstanceProtocolsBgpIpv6UnicastMultipath   `json:"multipath,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct
type NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs struct {
	// +kubebuilder:default:=false
	IgnorePeerAs *bool `json:"ignore-peer-as,omitempty"`
	// +kubebuilder:default:=false
	LeadingOnly *bool `json:"leading-only,omitempty"`
	// +kubebuilder:validation:Enum=`delete`;`disabled`;`replace`
	Mode *string `json:"mode"`
}

// NetworkinstanceProtocolsBgpNeighborAsPathOptions struct
type NetworkinstanceProtocolsBgpNeighborAsPathOptions struct {
	RemovePrivateAs *NetworkinstanceProtocolsBgpNeighborAsPathOptionsRemovePrivateAs `json:"remove-private-as,omitempty"`
	ReplacePeerAs   *bool                                                            `json:"replace-peer-as,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	AllowOwnAs *uint8 `json:"allow-own-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborAuthentication struct
type NetworkinstanceProtocolsBgpNeighborAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborEvpn struct
type NetworkinstanceProtocolsBgpNeighborEvpn struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                             `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                               `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpNeighborEvpnPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborFailureDetection struct
type NetworkinstanceProtocolsBgpNeighborFailureDetection struct {
	EnableBfd    *bool `json:"enable-bfd,omitempty"`
	FastFailover *bool `json:"fast-failover,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct
type NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborGracefulRestart struct
type NetworkinstanceProtocolsBgpNeighborGracefulRestart struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	StaleRoutesTime *uint16                                                        `json:"stale-routes-time,omitempty"`
	WarmRestart     *NetworkinstanceProtocolsBgpNeighborGracefulRestartWarmRestart `json:"warm-restart,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv4Unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState            *string                                                    `json:"admin-state,omitempty"`
	AdvertiseIpv6NextHops *bool                                                      `json:"advertise-ipv6-next-hops,omitempty"`
	PrefixLimit           *NetworkinstanceProtocolsBgpNeighborIpv4UnicastPrefixLimit `json:"prefix-limit,omitempty"`
	ReceiveIpv6NextHops   *bool                                                      `json:"receive-ipv6-next-hops,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct
type NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxReceivedRoutes *uint32 `json:"max-received-routes,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	WarningThresholdPct *uint8 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborIpv6Unicast struct
type NetworkinstanceProtocolsBgpNeighborIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState  *string                                                    `json:"admin-state,omitempty"`
	PrefixLimit *NetworkinstanceProtocolsBgpNeighborIpv6UnicastPrefixLimit `json:"prefix-limit,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborLocalAs struct
type NetworkinstanceProtocolsBgpNeighborLocalAs struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber        *uint32 `json:"as-number"`
	PrependGlobalAs *bool   `json:"prepend-global-as,omitempty"`
	PrependLocalAs  *bool   `json:"prepend-local-as,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborRouteReflector struct
type NetworkinstanceProtocolsBgpNeighborRouteReflector struct {
	Client *bool `json:"client,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSendCommunity struct
type NetworkinstanceProtocolsBgpNeighborSendCommunity struct {
	Large    *bool `json:"large,omitempty"`
	Standard *bool `json:"standard,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborSendDefaultRoute struct
type NetworkinstanceProtocolsBgpNeighborSendDefaultRoute struct {
	Ipv6Unicast  *bool   `json:"ipv6-unicast,omitempty"`
	ExportPolicy *string `json:"export-policy,omitempty"`
	Ipv4Unicast  *bool   `json:"ipv4-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTimers struct
type NetworkinstanceProtocolsBgpNeighborTimers struct {
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

// NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTraceOptions struct
type NetworkinstanceProtocolsBgpNeighborTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpNeighborTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighborTransport struct
type NetworkinstanceProtocolsBgpNeighborTransport struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	PassiveMode  *bool   `json:"passive-mode,omitempty"`
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgpNeighbor struct
type NetworkinstanceProtocolsBgpNeighbor struct {
	RouteReflector *NetworkinstanceProtocolsBgpNeighborRouteReflector `json:"route-reflector,omitempty"`
	Timers         *NetworkinstanceProtocolsBgpNeighborTimers         `json:"timers,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	PeerAddress *string `json:"peer-address"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                       `json:"admin-state,omitempty"`
	LocalAs    []*NetworkinstanceProtocolsBgpNeighborLocalAs `json:"local-as,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	LocalPreference *uint32 `json:"local-preference,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	PeerAs           *uint32                                              `json:"peer-as,omitempty"`
	PeerGroup        *string                                              `json:"peer-group"`
	Transport        *NetworkinstanceProtocolsBgpNeighborTransport        `json:"transport,omitempty"`
	AsPathOptions    *NetworkinstanceProtocolsBgpNeighborAsPathOptions    `json:"as-path-options,omitempty"`
	Evpn             *NetworkinstanceProtocolsBgpNeighborEvpn             `json:"evpn,omitempty"`
	Ipv4Unicast      *NetworkinstanceProtocolsBgpNeighborIpv4Unicast      `json:"ipv4-unicast,omitempty"`
	NextHopSelf      *bool                                                `json:"next-hop-self,omitempty"`
	SendCommunity    *NetworkinstanceProtocolsBgpNeighborSendCommunity    `json:"send-community,omitempty"`
	TraceOptions     *NetworkinstanceProtocolsBgpNeighborTraceOptions     `json:"trace-options,omitempty"`
	FailureDetection *NetworkinstanceProtocolsBgpNeighborFailureDetection `json:"failure-detection,omitempty"`
	ImportPolicy     *string                                              `json:"import-policy,omitempty"`
	SendDefaultRoute *NetworkinstanceProtocolsBgpNeighborSendDefaultRoute `json:"send-default-route,omitempty"`
	Authentication   *NetworkinstanceProtocolsBgpNeighborAuthentication   `json:"authentication,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description     *string                                             `json:"description,omitempty"`
	ExportPolicy    *string                                             `json:"export-policy,omitempty"`
	GracefulRestart *NetworkinstanceProtocolsBgpNeighborGracefulRestart `json:"graceful-restart,omitempty"`
	Ipv6Unicast     *NetworkinstanceProtocolsBgpNeighborIpv6Unicast     `json:"ipv6-unicast,omitempty"`
}

// NetworkinstanceProtocolsBgpPreference struct
type NetworkinstanceProtocolsBgpPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ebgp *uint8 `json:"ebgp,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=170
	Ibgp *uint8 `json:"ibgp,omitempty"`
}

// NetworkinstanceProtocolsBgpRouteAdvertisement struct
type NetworkinstanceProtocolsBgpRouteAdvertisement struct {
	// +kubebuilder:default:=false
	RapidWithdrawal *bool `json:"rapid-withdrawal,omitempty"`
	// +kubebuilder:default:=true
	WaitForFibInstall *bool `json:"wait-for-fib-install,omitempty"`
}

// NetworkinstanceProtocolsBgpRouteReflector struct
type NetworkinstanceProtocolsBgpRouteReflector struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	ClusterId *string `json:"cluster-id,omitempty"`
	// +kubebuilder:default:=false
	Client *bool `json:"client,omitempty"`
}

// NetworkinstanceProtocolsBgpSendCommunity struct
type NetworkinstanceProtocolsBgpSendCommunity struct {
	// +kubebuilder:default:=true
	Standard *bool `json:"standard,omitempty"`
	// +kubebuilder:default:=true
	Large *bool `json:"large,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceOptionsFlag struct
type NetworkinstanceProtocolsBgpTraceOptionsFlag struct {
	// +kubebuilder:validation:Enum=`events`;`graceful-restart`;`keepalive`;`notification`;`open`;`packets`;`route`;`socket`;`timers`;`update`
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`detail`;`receive`;`send`
	Modifier *string `json:"modifier,omitempty"`
}

// NetworkinstanceProtocolsBgpTraceOptions struct
type NetworkinstanceProtocolsBgpTraceOptions struct {
	Flag []*NetworkinstanceProtocolsBgpTraceOptionsFlag `json:"flag,omitempty"`
}

// NetworkinstanceProtocolsBgpTransport struct
type NetworkinstanceProtocolsBgpTransport struct {
	// +kubebuilder:validation:Minimum=536
	// +kubebuilder:validation:Maximum=9446
	// +kubebuilder:default:=1024
	TcpMss *uint16 `json:"tcp-mss,omitempty"`
}

// NetworkinstanceProtocolsBgp struct
type NetworkinstanceProtocolsBgp struct {
	AsPathOptions     *NetworkinstanceProtocolsBgpAsPathOptions     `json:"as-path-options,omitempty"`
	EbgpDefaultPolicy *NetworkinstanceProtocolsBgpEbgpDefaultPolicy `json:"ebgp-default-policy,omitempty"`
	Evpn              *NetworkinstanceProtocolsBgpEvpn              `json:"evpn,omitempty"`
	ExportPolicy      *string                                       `json:"export-policy,omitempty"`
	FailureDetection  *NetworkinstanceProtocolsBgpFailureDetection  `json:"failure-detection,omitempty"`
	GracefulRestart   *NetworkinstanceProtocolsBgpGracefulRestart   `json:"graceful-restart,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RouterId *string `json:"router-id"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState     *string                                    `json:"admin-state,omitempty"`
	Authentication *NetworkinstanceProtocolsBgpAuthentication `json:"authentication,omitempty"`
	Convergence    *NetworkinstanceProtocolsBgpConvergence    `json:"convergence,omitempty"`
	Ipv6Unicast    *NetworkinstanceProtocolsBgpIpv6Unicast    `json:"ipv6-unicast,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=100
	LocalPreference    *uint32                                        `json:"local-preference,omitempty"`
	Neighbor           []*NetworkinstanceProtocolsBgpNeighbor         `json:"neighbor,omitempty"`
	Preference         *NetworkinstanceProtocolsBgpPreference         `json:"preference,omitempty"`
	SendCommunity      *NetworkinstanceProtocolsBgpSendCommunity      `json:"send-community,omitempty"`
	Group              []*NetworkinstanceProtocolsBgpGroup            `json:"group,omitempty"`
	ImportPolicy       *string                                        `json:"import-policy,omitempty"`
	RouteAdvertisement *NetworkinstanceProtocolsBgpRouteAdvertisement `json:"route-advertisement,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AutonomousSystem *uint32                                      `json:"autonomous-system"`
	DynamicNeighbors *NetworkinstanceProtocolsBgpDynamicNeighbors `json:"dynamic-neighbors,omitempty"`
	Ipv4Unicast      *NetworkinstanceProtocolsBgpIpv4Unicast      `json:"ipv4-unicast,omitempty"`
	RouteReflector   *NetworkinstanceProtocolsBgpRouteReflector   `json:"route-reflector,omitempty"`
	TraceOptions     *NetworkinstanceProtocolsBgpTraceOptions     `json:"trace-options,omitempty"`
	Transport        *NetworkinstanceProtocolsBgpTransport        `json:"transport,omitempty"`
}

// SrlNetworkinstanceProtocolsBgpSpec struct
type SrlNetworkinstanceProtocolsBgpSpec struct {
	SrlNokiaNetworkInstanceName    *string                      `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsBgp *NetworkinstanceProtocolsBgp `json:"bgp"`
}

// SrlNetworkinstanceProtocolsBgpStatus struct
type SrlNetworkinstanceProtocolsBgpStatus struct {
	// ConfigurationDependencyTargetNotFound identifies if the target of the resource object is missing or not
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyTargetFound *TargetFoundStatus `json:"configurationDependencyTargetFound,omitempty"`

	// ConfigurationDependencyValidationStatus identifies the status of the LeafRef Validation of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyValidationStatus *ValidationStatus `json:"configurationDependencyValidationStatus,omitempty"`

	// ConfigurationDependencyValidationDetails defines the validation details of the resource object
	ConfigurationDependencyValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNetworkinstanceProtocolsBgpSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlNetworkinstanceProtocolsBgp is the Schema for the SrlNetworkinstanceProtocolsBgps API
type SrlNetworkinstanceProtocolsBgp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNetworkinstanceProtocolsBgpSpec   `json:"spec,omitempty"`
	Status SrlNetworkinstanceProtocolsBgpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpList contains a list of SrlNetworkinstanceProtocolsBgps
type SrlNetworkinstanceProtocolsBgpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgp{}, &SrlNetworkinstanceProtocolsBgpList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlNetworkinstanceProtocolsBgp) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNetworkinstanceProtocolsBgp",
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

func (o *SrlNetworkinstanceProtocolsBgp) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlNetworkinstanceProtocolsBgp) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
