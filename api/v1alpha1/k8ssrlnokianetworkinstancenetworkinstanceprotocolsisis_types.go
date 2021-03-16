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
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisFinalizer string = "Isis.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAttachedBit struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAttachedBit struct {
	// +kubebuilder:default:=false
	Ignore *bool `json:"ignore,omitempty"`
	// +kubebuilder:default:=false
	Suppress *bool `json:"suppress,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAuthentication struct {
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAutoCost struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAutoCost struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8000000000
	ReferenceBandwidth *uint64 `json:"reference-bandwidth,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceGracefulRestart struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceGracefulRestart struct {
	// +kubebuilder:default:=false
	HelperMode *bool `json:"helper-mode,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	RouteTag *uint32 `json:"route-tag,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct {
	SummaryAddress []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress `json:"summary-address,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct {
	Level1ToLevel2 *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 `json:"level1-to-level2,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceAuthentication struct {
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	IncludeBfdTlv *bool `json:"include-bfd-tlv,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	IncludeBfdTlv *bool `json:"include-bfd-tlv,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	EndOfLib *bool   `json:"end-of-lib,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelTimers struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20000
	// +kubebuilder:default:=9
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	HelloMultiplier *uint8 `json:"hello-multiplier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevel struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevel struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Ipv6UnicastMetric *uint32 `json:"ipv6-unicast-metric,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Metric *uint32 `json:"metric,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=64
	Priority *uint8                                                                           `json:"priority,omitempty"`
	Timers   *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelTimers `json:"timers,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	LevelNumber    *uint8                                                                                   `json:"level-number"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Disable *bool `json:"disable,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTimers struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	CsnpInterval *uint16 `json:"csnp-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=100
	LspPacingInterval *uint64 `json:"lsp-pacing-interval,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`adjacencies`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`
	Trace *string `json:"trace"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterface struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterface struct {
	InterfaceName *string `json:"interface-name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`adaptive`;`disable`;`loose`;`strict`
	// +kubebuilder:default:=disable
	HelloPadding *string                                                                          `json:"hello-padding"`
	Ipv6Unicast  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv6Unicast `json:"ipv6-unicast,omitempty"`
	Level        []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLevel     `json:"level,omitempty"`
	// +kubebuilder:default:=false
	Passive      *bool                                                                             `json:"passive,omitempty"`
	Timers       *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTimers       `json:"timers,omitempty"`
	TraceOptions *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceTraceOptions `json:"trace-options,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState     *string                                                                             `json:"admin-state"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	CircuitType        *string                                                                                 `json:"circuit-type"`
	Ipv4Unicast        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceIpv4Unicast        `json:"ipv4-unicast,omitempty"`
	LdpSynchronization *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterfaceLdpSynchronization `json:"ldp-synchronization,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv4Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv4Unicast struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv6Unicast struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv6Unicast struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:default:=false
	MultiTopology *bool `json:"multi-topology,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLdpSynchronization struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLdpSynchronization struct {
	// +kubebuilder:default:=false
	EndOfLib *bool `json:"end-of-lib,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelAuthentication struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelAuthentication struct {
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelRoutePreference struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelRoutePreference struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	Internal *uint8 `json:"internal,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	External *uint8 `json:"external,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`adjacencies`;`lsdb`;`routes`;`spf`
	Trace *string `json:"trace"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevel struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevel struct {
	TraceOptions *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelTraceOptions `json:"trace-options,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	LevelNumber    *uint8                                                                          `json:"level-number"`
	Authentication *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	BgpLsExclude *bool `json:"bgp-ls-exclude,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`narrow`;`wide`
	// +kubebuilder:default:=wide
	MetricStyle     *string                                                                          `json:"metric-style"`
	RoutePreference *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevelRoutePreference `json:"route-preference,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadImmediate struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadImmediate struct {
	// +kubebuilder:default:=false
	MaxMetric *bool `json:"max-metric,omitempty"`
	// +kubebuilder:default:=false
	SetBit *bool `json:"set-bit,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadOnBoot struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadOnBoot struct {
	MaxMetric *bool `json:"max-metric,omitempty"`
	SetBit    *bool `json:"set-bit,omitempty"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=1800
	Timeout *uint16 `json:"timeout,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverload struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverload struct {
	// +kubebuilder:default:=false
	AdvertiseExternal *bool `json:"advertise-external,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseInterlevel *bool                                                                         `json:"advertise-interlevel,omitempty"`
	Immediate           *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadImmediate `json:"immediate,omitempty"`
	OnBoot              *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverloadOnBoot    `json:"on-boot,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstall struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstall struct {
	BgpLs *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspGeneration struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspGeneration struct {
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=5000
	MaxWait *uint64 `json:"max-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	SecondWait *uint64 `json:"second-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=10
	InitialWait *uint64 `json:"initial-wait,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspRefresh struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspRefresh struct {
	// +kubebuilder:default:=true
	HalfLifetime *bool `json:"half-lifetime,omitempty"`
	// +kubebuilder:validation:Minimum=150
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=600
	Interval *uint16 `json:"interval,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersSpf struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersSpf struct {
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	InitialWait *uint64 `json:"initial-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=10000
	MaxWait *uint64 `json:"max-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	SecondWait *uint64 `json:"second-wait,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimers struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimers struct {
	LspGeneration *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspGeneration `json:"lsp-generation,omitempty"`
	// +kubebuilder:validation:Minimum=350
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1200
	LspLifetime *uint16                                                                      `json:"lsp-lifetime,omitempty"`
	LspRefresh  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersLspRefresh `json:"lsp-refresh,omitempty"`
	Spf         *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimersSpf        `json:"spf,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTraceOptions struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTraceOptions struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`adjacencies`;`graceful-restart`;`interfaces`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`;`routes`;`summary-addresses`
	Trace *string `json:"trace"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTrafficEngineering struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTrafficEngineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTransport struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTransport struct {
	// +kubebuilder:validation:Minimum=490
	// +kubebuilder:validation:Maximum=9490
	// +kubebuilder:default:=1492
	LspMtuSize *uint16 `json:"lsp-mtu-size,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstance struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstance struct {
	Timers      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTimers      `json:"timers,omitempty"`
	Transport   *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTransport   `json:"transport,omitempty"`
	Ipv4Unicast *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv4Unicast `json:"ipv4-unicast,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxEcmpPaths    *uint8                                                                      `json:"max-ecmp-paths,omitempty"`
	Authentication  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAuthentication  `json:"authentication,omitempty"`
	GracefulRestart *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceGracefulRestart `json:"graceful-restart,omitempty"`
	Interface       []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterface     `json:"interface,omitempty"`
	Ipv6Unicast     *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceIpv6Unicast     `json:"ipv6-unicast,omitempty"`
	Overload        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceOverload        `json:"overload,omitempty"`
	TraceOptions    *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTraceOptions    `json:"trace-options,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name               *string                                                                        `json:"name"`
	AttachedBit        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAttachedBit        `json:"attached-bit,omitempty"`
	LdpSynchronization *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLdpSynchronization `json:"ldp-synchronization,omitempty"`
	Level              []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceLevel            `json:"level,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`L1`;`L1L2`;`L2`
	// +kubebuilder:default:=L2
	LevelCapability *string `json:"level-capability"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[0]{2}`
	Net                *string                                                                        `json:"net,omitempty"`
	TeDatabaseInstall  *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTeDatabaseInstall  `json:"te-database-install,omitempty"`
	TrafficEngineering *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState   *string `json:"admin-state"`
	ExportPolicy *string `json:"export-policy,omitempty"`
	// +kubebuilder:default:=false
	PoiTlv                        *bool                                                                                     `json:"poi-tlv,omitempty"`
	AutoCost                      *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceAutoCost                      `json:"auto-cost,omitempty"`
	InterLevelPropagationPolicies *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstanceInterLevelPropagationPolicies `json:"inter-level-propagation-policies,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis struct {
	Instance []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisInstance `json:"instance,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisSpec struct {
	SrlNokiaNetworkInstanceName                         *string                                              `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis *SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis `json:"isis"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsiss API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsiss
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisList{})
}
