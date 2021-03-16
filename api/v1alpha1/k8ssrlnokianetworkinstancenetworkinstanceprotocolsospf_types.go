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
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfFinalizer string = "Ospf.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaAreaRange struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaAreaRange struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceFailureDetection struct {
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket struct {
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type *string `json:"type"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace struct {
	Interfaces  *string                                                                                          `json:"interfaces,omitempty"`
	Packet      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTracePacket `json:"packet,omitempty"`
	Adjacencies *string                                                                                          `json:"adjacencies,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptions struct {
	Trace *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptionsTrace `json:"trace,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterface struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterface struct {
	InterfaceName *string `json:"interface-name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState     *string                                                                                 `json:"admin-state"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceAuthentication `json:"authentication,omitempty"`
	Passive        *bool                                                                                   `json:"passive,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Priority *uint16 `json:"priority,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=5
	RetransmitInterval *uint32 `json:"retransmit-interval,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=1
	TransitDelay *uint32 `json:"transit-delay,omitempty"`
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=40
	DeadInterval     *uint32                                                                                   `json:"dead-interval,omitempty"`
	FailureDetection *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceFailureDetection `json:"failure-detection,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Metric *uint16 `json:"metric,omitempty"`
	// +kubebuilder:default:=true
	AdvertiseRouterCapability *bool `json:"advertise-router-capability,omitempty"`
	// +kubebuilder:default:=true
	AdvertiseSubnet *bool `json:"advertise-subnet,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	InterfaceType *string `json:"interface-type"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`all`;`except-own-rtrlsa`;`except-own-rtrlsa-and-defaults`;`none`
	// +kubebuilder:default:=none
	LsaFilterOut *string `json:"lsa-filter-out"`
	// +kubebuilder:validation:Minimum=512
	// +kubebuilder:validation:Maximum=9486
	Mtu          *uint32                                                                               `json:"mtu,omitempty"`
	TraceOptions *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterfaceTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaAreaRange struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaAreaRange struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefixMask *string `json:"ip-prefix-mask"`
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute struct {
	// +kubebuilder:default:=true
	AdjacencyCheck *bool `json:"adjacency-check,omitempty"`
	// +kubebuilder:default:=false
	TypeNssa *bool `json:"type-nssa,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssa struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssa struct {
	AreaRange             []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaAreaRange           `json:"area-range,omitempty"`
	OriginateDefaultRoute *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssaOriginateDefaultRoute `json:"originate-default-route,omitempty"`
	RedistributeExternal  *bool                                                                                     `json:"redistribute-external,omitempty"`
	Summaries             *bool                                                                                     `json:"summaries,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaStub struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaStub struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	DefaultMetric *uint16 `json:"default-metric,omitempty"`
	Summaries     *bool   `json:"summaries,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceArea struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceArea struct {
	// +kubebuilder:default:=false
	BgpLsExclude *bool                                                                       `json:"bgp-ls-exclude,omitempty"`
	ExportPolicy *string                                                                     `json:"export-policy,omitempty"`
	Interface    []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaInterface `json:"interface,omitempty"`
	Nssa         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaNssa        `json:"nssa,omitempty"`
	Stub         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaStub        `json:"stub,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`[0-9\.]*`
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%!+(BADINDEX))?`
	AreaId *string `json:"area-id"`
	// +kubebuilder:default:=true
	AdvertiseRouterCapability *bool                                                                       `json:"advertise-router-capability,omitempty"`
	AreaRange                 []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAreaAreaRange `json:"area-range,omitempty"`
	BlackholeAggregate        *bool                                                                       `json:"blackhole-aggregate,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAsbr struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAsbr struct {
	// +kubebuilder:default:=none
	TracePath *string `json:"trace-path,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExportLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExportLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	LogPercent *uint32 `json:"log-percent,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=0
	Number *uint32 `json:"number"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExternalDbOverflow struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExternalDbOverflow struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Interval *uint32 `json:"interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// +kubebuilder:default:=0
	Limit *uint32 `json:"limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceGracefulRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceGracefulRestart struct {
	// +kubebuilder:default:=false
	StrictLsaChecking *bool `json:"strict-lsa-checking,omitempty"`
	// +kubebuilder:default:=false
	HelperMode *bool `json:"helper-mode,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadOverloadOnBoot struct {
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	Timeout *uint32 `json:"timeout,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	OverloadTimeout *uint16 `json:"overload-timeout,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=0
	WarningThreshold *uint8 `json:"warning-threshold,omitempty"`
	LogOnly          *bool  `json:"log-only,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	MaxLsaCount *uint32 `json:"max-lsa-count,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverload struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverload struct {
	// +kubebuilder:default:=false
	Active *bool `json:"active,omitempty"`
	// +kubebuilder:default:=false
	OverloadIncludeExt1 *bool `json:"overload-include-ext-1,omitempty"`
	// +kubebuilder:default:=false
	OverloadIncludeExt2 *bool `json:"overload-include-ext-2,omitempty"`
	// +kubebuilder:default:=false
	OverloadIncludeExtStub *bool                                                                              `json:"overload-include-ext-stub,omitempty"`
	OverloadOnBoot         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadOverloadOnBoot `json:"overload-on-boot,omitempty"`
	RtrAdvLsaLimit         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverloadRtrAdvLsaLimit `json:"rtr-adv-lsa-limit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstall struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstall struct {
	BgpLs *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersLsaGenerate struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersLsaGenerate struct {
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

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersSpfWait struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersSpfWait struct {
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

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=600000
	// +kubebuilder:default:=1000
	LsaArrival  *uint32                                                                       `json:"lsa-arrival,omitempty"`
	LsaGenerate *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersLsaGenerate `json:"lsa-generate,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	RedistributeDelay *uint32                                                                   `json:"redistribute-delay,omitempty"`
	SpfWait           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimersSpfWait `json:"spf-wait,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	IncrementalSpfWait *uint32 `json:"incremental-spf-wait,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// +kubebuilder:default:=1000
	LsaAccumulate *uint32 `json:"lsa-accumulate,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceLsdb struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	LinkStateId *string `json:"link-state-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId *string `json:"router-id,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`all`;`external`;`inter-area-prefix`;`inter-area-router`;`intra-area-prefix`;`network`;`nssa`;`opaque`;`router`;`summary`
	Type *string `json:"type"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTracePacket struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTracePacket struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`all`;`dbdescr`;`hello`;`ls-ack`;`ls-request`;`ls-update`
	Type   *string `json:"type"`
	Detail *string `json:"detail,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`drop`;`egress`;`in-and-egress`;`ingress`
	Modifier *string `json:"modifier"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceRoutes struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceSpf struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	DestAddress *string `json:"dest-address,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTrace struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTrace struct {
	GracefulRestart *string                                                                             `json:"graceful-restart,omitempty"`
	Interfaces      *string                                                                             `json:"interfaces,omitempty"`
	Lsdb            *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceLsdb   `json:"lsdb,omitempty"`
	Misc            *string                                                                             `json:"misc,omitempty"`
	Packet          *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTracePacket `json:"packet,omitempty"`
	Routes          *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceRoutes `json:"routes,omitempty"`
	Spf             *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTraceSpf    `json:"spf,omitempty"`
	Adjacencies     *string                                                                             `json:"adjacencies,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptions struct {
	Trace *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptionsTrace `json:"trace,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTrafficEngineering struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTrafficEngineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstance struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstance struct {
	ExportLimit        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExportLimit        `json:"export-limit,omitempty"`
	ExternalDbOverflow *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceExternalDbOverflow `json:"external-db-overflow,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=10
	Preference        *uint8                                                                        `json:"preference,omitempty"`
	TeDatabaseInstall *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTeDatabaseInstall `json:"te-database-install,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`area`;`as`;`false`;`link`
	AdvertiseRouterCapability *string                                                            `json:"advertise-router-capability"`
	Area                      []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceArea `json:"area,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=150
	ExternalPreference *uint8 `json:"external-preference,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	InstanceId *uint32 `json:"instance-id,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxEcmpPaths *uint8                                                               `json:"max-ecmp-paths,omitempty"`
	Overload     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceOverload `json:"overload,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId           *string                                                                        `json:"router-id,omitempty"`
	TrafficEngineering *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name          *string                                                            `json:"name"`
	AddressFamily *string                                                            `json:"address-family,omitempty"`
	Asbr          *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceAsbr   `json:"asbr,omitempty"`
	ExportPolicy  *string                                                            `json:"export-policy,omitempty"`
	Timers        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTimers `json:"timers,omitempty"`
	Version       *string                                                            `json:"version"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState      *string                                                                     `json:"admin-state"`
	GracefulRestart *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceGracefulRestart `json:"graceful-restart,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8000000000
	// +kubebuilder:default:=400000000
	ReferenceBandwidth *uint64                                                                  `json:"reference-bandwidth,omitempty"`
	TraceOptions       *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstanceTraceOptions `json:"trace-options,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf struct {
	Instance []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfInstance `json:"instance,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfSpec struct {
	SrlNokiaNetworkInstanceName                         *string                                              `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf *SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf `json:"ospf"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfs API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfs
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfList{})
}