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
	// SrlNetworkinstanceProtocolsOspfFinalizer is the name of the finalizer added to
	// SrlNetworkinstanceProtocolsOspf to block delete operations until the physical node can be
	// deprovisioned.
	SrlNetworkinstanceProtocolsOspfFinalizer string = "Ospf.srlinux.henderiw.be"
)

// NetworkinstanceProtocolsOspfInstanceAreaAreaRange struct
type NetworkinstanceProtocolsOspfInstanceAreaAreaRange struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct {
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct {
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type   *string `json:"type,omitempty"`
	Detail *string `json:"detail,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct {
	Adjacencies *string                                                                   `json:"adjacencies,omitempty"`
	Interfaces  *string                                                                   `json:"interfaces,omitempty"`
	Packet      *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket `json:"packet,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct
type NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaInterface struct
type NetworkinstanceProtocolsOspfInstanceAreaInterface struct {
	// +kubebuilder:default:=true
	AdvertiseRouterCapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=40
	DeadInterval *uint32 `json:"dead-interval,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`except-own-rtrlsa`;`except-own-rtrlsa-and-defaults`;`none`
	// +kubebuilder:default:=none
	LsaFilterOut *string `json:"lsa-filter-out,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=5
	RetransmitInterval *uint32 `json:"retransmit-interval,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=1
	TransitDelay *uint32 `json:"transit-delay,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=512
	// +kubebuilder:validation:Maximum=9486
	Mtu     *uint32 `json:"mtu,omitempty"`
	Passive *bool   `json:"passive,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Priority         *uint16                                                            `json:"priority,omitempty"`
	InterfaceName    *string                                                            `json:"interface-name"`
	FailureDetection *NetworkinstanceProtocolsOspfInstanceAreaInterfaceFailureDetection `json:"failure-detection,omitempty"`
	TraceOptions     *NetworkinstanceProtocolsOspfInstanceAreaInterfaceTraceOptions     `json:"trace-options,omitempty"`
	// +kubebuilder:default:=true
	AdvertiseSubnet *bool                                                            `json:"advertise-subnet,omitempty"`
	Authentication  *NetworkinstanceProtocolsOspfInstanceAreaInterfaceAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	InterfaceType *string `json:"interface-type,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Metric *uint16 `json:"metric,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct
type NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct {
	// +kubebuilder:default:=true
	AdjacencyCheck *bool `json:"adjacency-check,omitempty"`
	// +kubebuilder:default:=false
	TypeNssa *bool `json:"type-nssa,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaNssa struct
type NetworkinstanceProtocolsOspfInstanceAreaNssa struct {
	Summaries             *bool                                                              `json:"summaries,omitempty"`
	AreaRange             []*NetworkinstanceProtocolsOspfInstanceAreaNssaAreaRange           `json:"area-range,omitempty"`
	OriginateDefaultRoute *NetworkinstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute `json:"originate-default-route,omitempty"`
	RedistributeExternal  *bool                                                              `json:"redistribute-external,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAreaStub struct
type NetworkinstanceProtocolsOspfInstanceAreaStub struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	DefaultMetric *uint16 `json:"default-metric,omitempty"`
	Summaries     *bool   `json:"summaries,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceArea struct
type NetworkinstanceProtocolsOspfInstanceArea struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	AreaId             *string                                       `json:"area-id"`
	BlackholeAggregate *bool                                         `json:"blackhole-aggregate,omitempty"`
	ExportPolicy       *string                                       `json:"export-policy,omitempty"`
	Nssa               *NetworkinstanceProtocolsOspfInstanceAreaNssa `json:"nssa,omitempty"`
	// +kubebuilder:default:=true
	AdvertiseRouterCapability *bool                                                `json:"advertise-router-capability,omitempty"`
	AreaRange                 []*NetworkinstanceProtocolsOspfInstanceAreaAreaRange `json:"area-range,omitempty"`
	// +kubebuilder:default:=false
	BgpLsExclude *bool                                                `json:"bgp-ls-exclude,omitempty"`
	Interface    []*NetworkinstanceProtocolsOspfInstanceAreaInterface `json:"interface,omitempty"`
	Stub         *NetworkinstanceProtocolsOspfInstanceAreaStub        `json:"stub,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceAsbr struct
type NetworkinstanceProtocolsOspfInstanceAsbr struct {
	// +kubebuilder:default:=none
	TracePath *string `json:"trace-path,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceExportLimit struct
type NetworkinstanceProtocolsOspfInstanceExportLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	LogPercent *uint32 `json:"log-percent,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=0
	Number *uint32 `json:"number"`
}

// NetworkinstanceProtocolsOspfInstanceExternalDbOverflow struct
type NetworkinstanceProtocolsOspfInstanceExternalDbOverflow struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Interval *uint32 `json:"interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Limit *uint32 `json:"limit,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceGracefulRestart struct
type NetworkinstanceProtocolsOspfInstanceGracefulRestart struct {
	// +kubebuilder:default:=false
	HelperMode *bool `json:"helper-mode,omitempty"`
	// +kubebuilder:default:=false
	StrictLsaChecking *bool `json:"strict-lsa-checking,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct
type NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct {
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Timeout *uint32 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct
type NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct {
	LogOnly *bool `json:"log-only,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxLsaCount *uint32 `json:"max-lsa-count,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	OverloadTimeout *uint16 `json:"overload-timeout,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	WarningThreshold *uint8 `json:"warning-threshold,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceOverload struct
type NetworkinstanceProtocolsOspfInstanceOverload struct {
	// +kubebuilder:default:=false
	OverloadIncludeExtStub *bool                                                       `json:"overload-include-ext-stub,omitempty"`
	OverloadOnBoot         *NetworkinstanceProtocolsOspfInstanceOverloadOverloadOnBoot `json:"overload-on-boot,omitempty"`
	RtrAdvLsaLimit         *NetworkinstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit `json:"rtr-adv-lsa-limit,omitempty"`
	// +kubebuilder:default:=false
	Active *bool `json:"active,omitempty"`
	// +kubebuilder:default:=false
	OverloadIncludeExt1 *bool `json:"overload-include-ext-1,omitempty"`
	// +kubebuilder:default:=false
	OverloadIncludeExt2 *bool `json:"overload-include-ext-2,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct
type NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall struct
type NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall struct {
	BgpLs *NetworkinstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate struct
type NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate struct {
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	LsaInitialWait *uint32 `json:"lsa-initial-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	LsaSecondWait *uint32 `json:"lsa-second-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=5000
	MaxLsaWait *uint32 `json:"max-lsa-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimersSpfWait struct
type NetworkinstanceProtocolsOspfInstanceTimersSpfWait struct {
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	SpfInitialWait *uint32 `json:"spf-initial-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	SpfMaxWait *uint32 `json:"spf-max-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	SpfSecondWait *uint32 `json:"spf-second-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTimers struct
type NetworkinstanceProtocolsOspfInstanceTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	IncrementalSpfWait *uint32 `json:"incremental-spf-wait,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	LsaAccumulate *uint32 `json:"lsa-accumulate,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=1000
	LsaArrival  *uint32                                                `json:"lsa-arrival,omitempty"`
	LsaGenerate *NetworkinstanceProtocolsOspfInstanceTimersLsaGenerate `json:"lsa-generate,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	RedistributeDelay *uint32                                            `json:"redistribute-delay,omitempty"`
	SpfWait           *NetworkinstanceProtocolsOspfInstanceTimersSpfWait `json:"spf-wait,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	LinkStateId *string `json:"link-state-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId *string `json:"router-id,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`external`;`inter-area-prefix`;`inter-area-router`;`intra-area-prefix`;`network`;`nssa`;`opaque`;`router`;`summary`
	Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier,omitempty"`
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type *string `json:"type,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace struct
type NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace struct {
	Interfaces      *string                                                      `json:"interfaces,omitempty"`
	Lsdb            *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceLsdb   `json:"lsdb,omitempty"`
	Misc            *string                                                      `json:"misc,omitempty"`
	Packet          *NetworkinstanceProtocolsOspfInstanceTraceOptionsTracePacket `json:"packet,omitempty"`
	Routes          *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceRoutes `json:"routes,omitempty"`
	Spf             *NetworkinstanceProtocolsOspfInstanceTraceOptionsTraceSpf    `json:"spf,omitempty"`
	Adjacencies     *string                                                      `json:"adjacencies,omitempty"`
	GracefulRestart *string                                                      `json:"graceful-restart,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTraceOptions struct
type NetworkinstanceProtocolsOspfInstanceTraceOptions struct {
	Trace *NetworkinstanceProtocolsOspfInstanceTraceOptionsTrace `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsOspfInstanceTrafficEngineering struct
type NetworkinstanceProtocolsOspfInstanceTrafficEngineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// NetworkinstanceProtocolsOspfInstance struct
type NetworkinstanceProtocolsOspfInstance struct {
	ExportLimit *NetworkinstanceProtocolsOspfInstanceExportLimit `json:"export-limit,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8000000000
	// +kubebuilder:default:=400000000
	ReferenceBandwidth *uint64                                     `json:"reference-bandwidth,omitempty"`
	Timers             *NetworkinstanceProtocolsOspfInstanceTimers `json:"timers,omitempty"`
	Version            *string                                     `json:"version"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string                                   `json:"admin-state,omitempty"`
	Asbr       *NetworkinstanceProtocolsOspfInstanceAsbr `json:"asbr,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxEcmpPaths       *uint8                                                  `json:"max-ecmp-paths,omitempty"`
	ExternalDbOverflow *NetworkinstanceProtocolsOspfInstanceExternalDbOverflow `json:"external-db-overflow,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=150
	ExternalPreference *uint8                                                 `json:"external-preference,omitempty"`
	GracefulRestart    *NetworkinstanceProtocolsOspfInstanceGracefulRestart   `json:"graceful-restart,omitempty"`
	TeDatabaseInstall  *NetworkinstanceProtocolsOspfInstanceTeDatabaseInstall `json:"te-database-install,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsOspfInstanceTraceOptions      `json:"trace-options,omitempty"`
	// +kubebuilder:validation:Enum=`area`;`as`;`false`;`link`
	AdvertiseRouterCapability *string                                     `json:"advertise-router-capability,omitempty"`
	Area                      []*NetworkinstanceProtocolsOspfInstanceArea `json:"area,omitempty"`
	ExportPolicy              *string                                     `json:"export-policy,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Preference *uint8 `json:"preference,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId           *string                                                 `json:"router-id,omitempty"`
	TrafficEngineering *NetworkinstanceProtocolsOspfInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	AddressFamily      *string                                                 `json:"address-family,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	InstanceId *uint32                                       `json:"instance-id,omitempty"`
	Overload   *NetworkinstanceProtocolsOspfInstanceOverload `json:"overload,omitempty"`
}

// NetworkinstanceProtocolsOspf struct
type NetworkinstanceProtocolsOspf struct {
	Instance []*NetworkinstanceProtocolsOspfInstance `json:"instance,omitempty"`
}

// SrlNetworkinstanceProtocolsOspfSpec struct
type SrlNetworkinstanceProtocolsOspfSpec struct {
	SrlNokiaNetworkInstanceName     *string                       `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsOspf *NetworkinstanceProtocolsOspf `json:"ospf"`
}

// SrlNetworkinstanceProtocolsOspfStatus struct
type SrlNetworkinstanceProtocolsOspfStatus struct {
	// ConfigurationDependencyTargetNotFound identifies if the target of the resource object is missing or not
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyTargetFound *TargetFoundStatus `json:"configurationDependencyTargetFound,omitempty"`

	// ConfigurationDependencyInternalLeafrefValidationStatus identifies the status of the local LeafRef Validation of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyInternalLeafrefValidationStatus *ValidationStatus `json:"configurationDependencyInternalLeafrefValidationStatus,omitempty"`

	// ConfigurationDependencyInternalLeafrefValidationDetails defines the validation details of the resource object
	ConfigurationDependencyInternalLeafrefValidationDetails map[string]*ValidationDetails `json:"internalLeafrefValidationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNetworkinstanceProtocolsOspfSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlNetworkinstanceProtocolsOspf is the Schema for the SrlNetworkinstanceProtocolsOspfs API
type SrlNetworkinstanceProtocolsOspf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNetworkinstanceProtocolsOspfSpec   `json:"spec,omitempty"`
	Status SrlNetworkinstanceProtocolsOspfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsOspfList contains a list of SrlNetworkinstanceProtocolsOspfs
type SrlNetworkinstanceProtocolsOspfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsOspf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsOspf{}, &SrlNetworkinstanceProtocolsOspfList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlNetworkinstanceProtocolsOspf) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNetworkinstanceProtocolsOspf",
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

func (o *SrlNetworkinstanceProtocolsOspf) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlNetworkinstanceProtocolsOspf) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
