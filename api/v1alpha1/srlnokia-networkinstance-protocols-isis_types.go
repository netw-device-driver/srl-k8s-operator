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
	// SrlnokiaNetworkinstanceProtocolsIsisFinalizer is the name of the finalizer added to
	// SrlnokiaNetworkinstanceProtocolsIsis to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaNetworkinstanceProtocolsIsisFinalizer string = "NetworkinstanceProtocolsIsis.srlinux.henderiw.be"
)

// NetworkinstanceProtocolsIsisInstanceAttachedBit struct
type NetworkinstanceProtocolsIsisInstanceAttachedBit struct {
	// +kubebuilder:default:=false
	Ignore *bool `json:"ignore,omitempty"`
	// +kubebuilder:default:=false
	Suppress *bool `json:"suppress,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceAuthentication struct {
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceAutoCost struct
type NetworkinstanceProtocolsIsisInstanceAutoCost struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8000000000
	ReferenceBandwidth *uint64 `json:"reference-bandwidth,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceGracefulRestart struct
type NetworkinstanceProtocolsIsisInstanceGracefulRestart struct {
	// +kubebuilder:default:=false
	HelperMode *bool `json:"helper-mode,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	RouteTag *uint32 `json:"route-tag,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 struct {
	SummaryAddress []*NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2SummaryAddress `json:"summary-address,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct
type NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies struct {
	Level1ToLevel2 *NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPoliciesLevel1ToLevel2 `json:"level1-to-level2,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication struct {
	Keychain            *string `json:"keychain,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast struct {
	// +kubebuilder:default:=false
	IncludeBfdTlv *bool `json:"include-bfd-tlv,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct
type NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	EnableBfd *bool `json:"enable-bfd,omitempty"`
	// +kubebuilder:default:=false
	IncludeBfdTlv *bool `json:"include-bfd-tlv,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization struct {
	Disable  *string `json:"disable,omitempty"`
	EndOfLib *bool   `json:"end-of-lib,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication struct {
	Keychain *string `json:"keychain,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20000
	// +kubebuilder:default:=9
	HelloInterval *uint32 `json:"hello-interval,omitempty"`
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=3
	HelloMultiplier *uint8 `json:"hello-multiplier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct
type NetworkinstanceProtocolsIsisInstanceInterfaceLevel struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	LevelNumber    *uint8                                                            `json:"level-number"`
	Authentication *NetworkinstanceProtocolsIsisInstanceInterfaceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	Disable *bool `json:"disable,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Ipv6UnicastMetric *uint32 `json:"ipv6-unicast-metric,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Metric *uint32 `json:"metric,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=64
	Priority *uint8                                                    `json:"priority,omitempty"`
	Timers   *NetworkinstanceProtocolsIsisInstanceInterfaceLevelTimers `json:"timers,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTimers struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=10
	CsnpInterval *uint16 `json:"csnp-interval,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=100
	LspPacingInterval *uint64 `json:"lsp-pacing-interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceInterface struct
type NetworkinstanceProtocolsIsisInstanceInterface struct {
	InterfaceName *string `json:"interface-name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Enum=`broadcast`;`point-to-point`
	CircuitType        *string                                                          `json:"circuit-type,omitempty"`
	Ipv4Unicast        *NetworkinstanceProtocolsIsisInstanceInterfaceIpv4Unicast        `json:"ipv4-unicast,omitempty"`
	LdpSynchronization *NetworkinstanceProtocolsIsisInstanceInterfaceLdpSynchronization `json:"ldp-synchronization,omitempty"`
	Level              []*NetworkinstanceProtocolsIsisInstanceInterfaceLevel            `json:"level,omitempty"`
	Timers             *NetworkinstanceProtocolsIsisInstanceInterfaceTimers             `json:"timers,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsIsisInstanceInterfaceTraceOptions       `json:"trace-options,omitempty"`
	Authentication     *NetworkinstanceProtocolsIsisInstanceInterfaceAuthentication     `json:"authentication,omitempty"`
	// +kubebuilder:validation:Enum=`adaptive`;`disable`;`loose`;`strict`
	// +kubebuilder:default:=disable
	HelloPadding *string                                                   `json:"hello-padding,omitempty"`
	Ipv6Unicast  *NetworkinstanceProtocolsIsisInstanceInterfaceIpv6Unicast `json:"ipv6-unicast,omitempty"`
	// +kubebuilder:default:=false
	Passive *bool `json:"passive,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv4Unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv4Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceIpv6Unicast struct
type NetworkinstanceProtocolsIsisInstanceIpv6Unicast struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:default:=false
	MultiTopology *bool `json:"multi-topology,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLdpSynchronization struct
type NetworkinstanceProtocolsIsisInstanceLdpSynchronization struct {
	// +kubebuilder:default:=false
	EndOfLib *bool `json:"end-of-lib,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:default:=60
	HoldDownTimer *uint16 `json:"hold-down-timer,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct
type NetworkinstanceProtocolsIsisInstanceLevelAuthentication struct {
	CsnpAuthentication  *bool   `json:"csnp-authentication,omitempty"`
	HelloAuthentication *bool   `json:"hello-authentication,omitempty"`
	Keychain            *string `json:"keychain,omitempty"`
	PsnpAuthentication  *bool   `json:"psnp-authentication,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelRoutePreference struct
type NetworkinstanceProtocolsIsisInstanceLevelRoutePreference struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	External *uint8 `json:"external,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	Internal *uint8 `json:"internal,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevelTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceLevelTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`lsdb`;`routes`;`spf`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceLevel struct
type NetworkinstanceProtocolsIsisInstanceLevel struct {
	TraceOptions *NetworkinstanceProtocolsIsisInstanceLevelTraceOptions `json:"trace-options,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	LevelNumber    *uint8                                                   `json:"level-number"`
	Authentication *NetworkinstanceProtocolsIsisInstanceLevelAuthentication `json:"authentication,omitempty"`
	// +kubebuilder:default:=false
	BgpLsExclude *bool `json:"bgp-ls-exclude,omitempty"`
	// +kubebuilder:validation:Enum=`narrow`;`wide`
	// +kubebuilder:default:=wide
	MetricStyle     *string                                                   `json:"metric-style,omitempty"`
	RoutePreference *NetworkinstanceProtocolsIsisInstanceLevelRoutePreference `json:"route-preference,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct
type NetworkinstanceProtocolsIsisInstanceOverloadImmediate struct {
	// +kubebuilder:default:=false
	MaxMetric *bool `json:"max-metric,omitempty"`
	// +kubebuilder:default:=false
	SetBit *bool `json:"set-bit,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverloadOnBoot struct
type NetworkinstanceProtocolsIsisInstanceOverloadOnBoot struct {
	MaxMetric *bool `json:"max-metric,omitempty"`
	SetBit    *bool `json:"set-bit,omitempty"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=1800
	Timeout *uint16 `json:"timeout,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceOverload struct
type NetworkinstanceProtocolsIsisInstanceOverload struct {
	// +kubebuilder:default:=false
	AdvertiseExternal *bool `json:"advertise-external,omitempty"`
	// +kubebuilder:default:=false
	AdvertiseInterlevel *bool                                                  `json:"advertise-interlevel,omitempty"`
	Immediate           *NetworkinstanceProtocolsIsisInstanceOverloadImmediate `json:"immediate,omitempty"`
	OnBoot              *NetworkinstanceProtocolsIsisInstanceOverloadOnBoot    `json:"on-boot,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct
type NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	BgpLsIdentifier *uint32 `json:"bgp-ls-identifier,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=-1
	IgpIdentifier *uint64 `json:"igp-identifier,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall struct
type NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall struct {
	BgpLs *NetworkinstanceProtocolsIsisInstanceTeDatabaseInstallBgpLs `json:"bgp-ls,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLspGeneration struct
type NetworkinstanceProtocolsIsisInstanceTimersLspGeneration struct {
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=10
	InitialWait *uint64 `json:"initial-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=120000
	// +kubebuilder:default:=5000
	MaxWait *uint64 `json:"max-wait,omitempty"`
	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// +kubebuilder:default:=1000
	SecondWait *uint64 `json:"second-wait,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersLspRefresh struct
type NetworkinstanceProtocolsIsisInstanceTimersLspRefresh struct {
	// +kubebuilder:default:=true
	HalfLifetime *bool `json:"half-lifetime,omitempty"`
	// +kubebuilder:validation:Minimum=150
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=600
	Interval *uint16 `json:"interval,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTimersSpf struct
type NetworkinstanceProtocolsIsisInstanceTimersSpf struct {
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

// NetworkinstanceProtocolsIsisInstanceTimers struct
type NetworkinstanceProtocolsIsisInstanceTimers struct {
	Spf           *NetworkinstanceProtocolsIsisInstanceTimersSpf           `json:"spf,omitempty"`
	LspGeneration *NetworkinstanceProtocolsIsisInstanceTimersLspGeneration `json:"lsp-generation,omitempty"`
	// +kubebuilder:validation:Minimum=350
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=1200
	LspLifetime *uint16                                               `json:"lsp-lifetime,omitempty"`
	LspRefresh  *NetworkinstanceProtocolsIsisInstanceTimersLspRefresh `json:"lsp-refresh,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTraceOptions struct
type NetworkinstanceProtocolsIsisInstanceTraceOptions struct {
	// +kubebuilder:validation:Enum=`adjacencies`;`graceful-restart`;`interfaces`;`packets-all`;`packets-l1-csnp`;`packets-l1-hello`;`packets-l1-lsp`;`packets-l1-psnp`;`packets-l2-csnp`;`packets-l2-hello`;`packets-l2-lsp`;`packets-l2-psnp`;`packets-p2p-hello`;`routes`;`summary-addresses`
	Trace *string `json:"trace,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTrafficEngineering struct
type NetworkinstanceProtocolsIsisInstanceTrafficEngineering struct {
	// +kubebuilder:default:=false
	Advertisement *bool `json:"advertisement,omitempty"`
	// +kubebuilder:default:=true
	LegacyLinkAttributeAdvertisement *bool `json:"legacy-link-attribute-advertisement,omitempty"`
}

// NetworkinstanceProtocolsIsisInstanceTransport struct
type NetworkinstanceProtocolsIsisInstanceTransport struct {
	// +kubebuilder:validation:Minimum=490
	// +kubebuilder:validation:Maximum=9490
	// +kubebuilder:default:=1492
	LspMtuSize *uint16 `json:"lsp-mtu-size,omitempty"`
}

// NetworkinstanceProtocolsIsisInstance struct
type NetworkinstanceProtocolsIsisInstance struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxEcmpPaths      *uint8                                                 `json:"max-ecmp-paths,omitempty"`
	TeDatabaseInstall *NetworkinstanceProtocolsIsisInstanceTeDatabaseInstall `json:"te-database-install,omitempty"`
	GracefulRestart   *NetworkinstanceProtocolsIsisInstanceGracefulRestart   `json:"graceful-restart,omitempty"`
	Ipv4Unicast       *NetworkinstanceProtocolsIsisInstanceIpv4Unicast       `json:"ipv4-unicast,omitempty"`
	Ipv6Unicast       *NetworkinstanceProtocolsIsisInstanceIpv6Unicast       `json:"ipv6-unicast,omitempty"`
	Level             []*NetworkinstanceProtocolsIsisInstanceLevel           `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`L1`;`L1L2`;`L2`
	// +kubebuilder:default:=L2
	LevelCapability               *string                                                            `json:"level-capability,omitempty"`
	AttachedBit                   *NetworkinstanceProtocolsIsisInstanceAttachedBit                   `json:"attached-bit,omitempty"`
	AutoCost                      *NetworkinstanceProtocolsIsisInstanceAutoCost                      `json:"auto-cost,omitempty"`
	InterLevelPropagationPolicies *NetworkinstanceProtocolsIsisInstanceInterLevelPropagationPolicies `json:"inter-level-propagation-policies,omitempty"`
	Interface                     []*NetworkinstanceProtocolsIsisInstanceInterface                   `json:"interface,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name     *string                                       `json:"name"`
	Overload *NetworkinstanceProtocolsIsisInstanceOverload `json:"overload,omitempty"`
	// +kubebuilder:default:=false
	PoiTlv *bool                                       `json:"poi-tlv,omitempty"`
	Timers *NetworkinstanceProtocolsIsisInstanceTimers `json:"timers,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState         *string                                                 `json:"admin-state,omitempty"`
	Authentication     *NetworkinstanceProtocolsIsisInstanceAuthentication     `json:"authentication,omitempty"`
	ExportPolicy       *string                                                 `json:"export-policy,omitempty"`
	LdpSynchronization *NetworkinstanceProtocolsIsisInstanceLdpSynchronization `json:"ldp-synchronization,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[0]{2}`
	Net                *string                                                 `json:"net,omitempty"`
	TraceOptions       *NetworkinstanceProtocolsIsisInstanceTraceOptions       `json:"trace-options,omitempty"`
	TrafficEngineering *NetworkinstanceProtocolsIsisInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	Transport          *NetworkinstanceProtocolsIsisInstanceTransport          `json:"transport,omitempty"`
}

// NetworkinstanceProtocolsIsis struct
type NetworkinstanceProtocolsIsis struct {
	Instance []*NetworkinstanceProtocolsIsisInstance `json:"instance,omitempty"`
}

// SrlnokiaNetworkinstanceProtocolsIsisSpec struct
type SrlnokiaNetworkinstanceProtocolsIsisSpec struct {
	SrlNokiaNetworkInstanceName          *string                       `json:"network-instance-name"`
	SrlnokiaNetworkinstanceProtocolsIsis *NetworkinstanceProtocolsIsis `json:"isis"`
}

// SrlnokiaNetworkinstanceProtocolsIsisStatus struct
type SrlnokiaNetworkinstanceProtocolsIsisStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaNetworkinstanceProtocolsIsisSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaNetworkinstanceProtocolsIsis is the Schema for the SrlnokiaNetworkinstanceProtocolsIsiss API
type SrlnokiaNetworkinstanceProtocolsIsis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaNetworkinstanceProtocolsIsisSpec   `json:"spec,omitempty"`
	Status SrlnokiaNetworkinstanceProtocolsIsisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaNetworkinstanceProtocolsIsisList contains a list of SrlnokiaNetworkinstanceProtocolsIsiss
type SrlnokiaNetworkinstanceProtocolsIsisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaNetworkinstanceProtocolsIsis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaNetworkinstanceProtocolsIsis{}, &SrlnokiaNetworkinstanceProtocolsIsisList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaNetworkinstanceProtocolsIsis) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaNetworkinstanceProtocolsIsis",
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

func (o *SrlnokiaNetworkinstanceProtocolsIsis) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
