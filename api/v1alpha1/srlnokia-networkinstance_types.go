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
	// SrlnokiaNetworkinstanceFinalizer is the name of the finalizer added to
	// SrlnokiaNetworkinstance to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaNetworkinstanceFinalizer string = "Networkinstance.srlinux.henderiw.be"
)

// NetworkinstanceBridgeTableMacDuplication struct
type NetworkinstanceBridgeTableMacDuplication struct {
	// +kubebuilder:default:=9
	HoldDownTime *uint32 `json:"hold-down-time,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	MonitoringWindow *uint32 `json:"monitoring-window,omitempty"`
	// +kubebuilder:validation:Minimum=3
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	NumMoves *uint32 `json:"num-moves,omitempty"`
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`
	// +kubebuilder:default:=stop-learning
	Action *string `json:"action,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// NetworkinstanceBridgeTableMacLearningAging struct
type NetworkinstanceBridgeTableMacLearningAging struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=300
	AgeTime *int32 `json:"age-time,omitempty"`
}

// NetworkinstanceBridgeTableMacLearning struct
type NetworkinstanceBridgeTableMacLearning struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                     `json:"admin-state,omitempty"`
	Aging      *NetworkinstanceBridgeTableMacLearningAging `json:"aging,omitempty"`
}

// NetworkinstanceBridgeTableMacLimit struct
type NetworkinstanceBridgeTableMacLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	MaximumEntries *int32 `json:"maximum-entries,omitempty"`
	// +kubebuilder:validation:Minimum=6
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	WarningThresholdPct *int32 `json:"warning-threshold-pct,omitempty"`
}

// NetworkinstanceBridgeTableStaticMacMac struct
type NetworkinstanceBridgeTableStaticMacMac struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Address     *string `json:"address"`
	Destination *string `json:"destination"`
}

// NetworkinstanceBridgeTableStaticMac struct
type NetworkinstanceBridgeTableStaticMac struct {
	Mac []*NetworkinstanceBridgeTableStaticMacMac `json:"mac,omitempty"`
}

// NetworkinstanceBridgeTable struct
type NetworkinstanceBridgeTable struct {
	MacLimit *NetworkinstanceBridgeTableMacLimit `json:"mac-limit,omitempty"`
	// +kubebuilder:default:=false
	ProtectAnycastGwMac *bool                                `json:"protect-anycast-gw-mac,omitempty"`
	StaticMac           *NetworkinstanceBridgeTableStaticMac `json:"static-mac,omitempty"`
	// +kubebuilder:default:=false
	DiscardUnknownDestMac *bool                                     `json:"discard-unknown-dest-mac,omitempty"`
	MacDuplication        *NetworkinstanceBridgeTableMacDuplication `json:"mac-duplication,omitempty"`
	MacLearning           *NetworkinstanceBridgeTableMacLearning    `json:"mac-learning,omitempty"`
}

// NetworkinstanceInterface struct
type NetworkinstanceInterface struct {
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// NetworkinstanceIpForwarding struct
type NetworkinstanceIpForwarding struct {
	ReceiveIpv4Check *bool `json:"receive-ipv4-check,omitempty"`
	ReceiveIpv6Check *bool `json:"receive-ipv6-check,omitempty"`
}

// NetworkinstanceIpLoadBalancingResilientHashPrefix struct
type NetworkinstanceIpLoadBalancingResilientHashPrefix struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:default:=1
	HashBucketsPerPath *uint8 `json:"hash-buckets-per-path,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MaxPaths *uint8 `json:"max-paths,omitempty"`
}

// NetworkinstanceIpLoadBalancing struct
type NetworkinstanceIpLoadBalancing struct {
	ResilientHashPrefix []*NetworkinstanceIpLoadBalancingResilientHashPrefix `json:"resilient-hash-prefix,omitempty"`
}

// NetworkinstanceMplsStaticMplsEntry struct
type NetworkinstanceMplsStaticMplsEntry struct {
	TopLabel *string `json:"top-label"`
	// +kubebuilder:default:=false
	CollectStats *bool   `json:"collect-stats,omitempty"`
	NextHopGroup *string `json:"next-hop-group,omitempty"`
	// +kubebuilder:validation:Enum=`pop`;`swap`
	// +kubebuilder:default:=swap
	Operation *string `json:"operation,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8 `json:"preference,omitempty"`
}

// NetworkinstanceMpls struct
type NetworkinstanceMpls struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState      *string                               `json:"admin-state,omitempty"`
	StaticMplsEntry []*NetworkinstanceMplsStaticMplsEntry `json:"static-mpls-entry,omitempty"`
	// +kubebuilder:default:=false
	TtlPropagation *bool `json:"ttl-propagation,omitempty"`
}

// NetworkinstanceMtu struct
type NetworkinstanceMtu struct {
	PathMtuDiscovery *bool `json:"path-mtu-discovery,omitempty"`
}

// NetworkinstanceTrafficEngineeringAdminGroupsGroup struct
type NetworkinstanceTrafficEngineeringAdminGroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	BitPosition *uint32 `json:"bit-position,omitempty"`
}

// NetworkinstanceTrafficEngineeringAdminGroups struct
type NetworkinstanceTrafficEngineeringAdminGroups struct {
	Group []*NetworkinstanceTrafficEngineeringAdminGroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficEngineeringInterfaceDelay struct
type NetworkinstanceTrafficEngineeringInterfaceDelay struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Static *uint32 `json:"static,omitempty"`
}

// NetworkinstanceTrafficEngineeringInterface struct
type NetworkinstanceTrafficEngineeringInterface struct {
	InterfaceName  *string                                          `json:"interface-name"`
	AdminGroup     *string                                          `json:"admin-group,omitempty"`
	Delay          *NetworkinstanceTrafficEngineeringInterfaceDelay `json:"delay,omitempty"`
	SrlgMembership *string                                          `json:"srlg-membership,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	TeMetric *uint32 `json:"te-metric,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	FromAddress *string `json:"from-address"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	ToAddress *string `json:"to-address,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct {
	StaticMember []*NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember `json:"static-member,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Value *uint32 `json:"value,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Cost *uint32 `json:"cost,omitempty"`
}

// NetworkinstanceTrafficEngineeringSharedRiskLinkGroups struct
type NetworkinstanceTrafficEngineeringSharedRiskLinkGroups struct {
	Group []*NetworkinstanceTrafficEngineeringSharedRiskLinkGroupsGroup `json:"group,omitempty"`
}

// NetworkinstanceTrafficEngineering struct
type NetworkinstanceTrafficEngineering struct {
	AdminGroups *NetworkinstanceTrafficEngineeringAdminGroups `json:"admin-groups,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AutonomousSystem *uint32                                       `json:"autonomous-system,omitempty"`
	Interface        []*NetworkinstanceTrafficEngineeringInterface `json:"interface,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4TeRouterId *string `json:"ipv4-te-router-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6TeRouterId       *string                                                `json:"ipv6-te-router-id,omitempty"`
	SharedRiskLinkGroups *NetworkinstanceTrafficEngineeringSharedRiskLinkGroups `json:"shared-risk-link-groups,omitempty"`
}

// NetworkinstanceVxlanInterface struct
type NetworkinstanceVxlanInterface struct {
	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=17
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,8}))`
	Name *string `json:"name"`
}

// Networkinstance struct
type Networkinstance struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string                     `json:"description,omitempty"`
	Interface   []*NetworkinstanceInterface `json:"interface,omitempty"`
	Mtu         *NetworkinstanceMtu         `json:"mtu,omitempty"`
	// +kubebuilder:default:=default
	Type        *string                     `json:"type,omitempty"`
	BridgeTable *NetworkinstanceBridgeTable `json:"bridge-table,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState      *string                         `json:"admin-state,omitempty"`
	IpForwarding    *NetworkinstanceIpForwarding    `json:"ip-forwarding,omitempty"`
	IpLoadBalancing *NetworkinstanceIpLoadBalancing `json:"ip-load-balancing,omitempty"`
	Mpls            *NetworkinstanceMpls            `json:"mpls,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId           *string                            `json:"router-id,omitempty"`
	TrafficEngineering *NetworkinstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
	VxlanInterface     []*NetworkinstanceVxlanInterface   `json:"vxlan-interface,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// SrlnokiaNetworkinstanceSpec struct
type SrlnokiaNetworkinstanceSpec struct {
	SrlnokiaNetworkinstance *[]Networkinstance `json:"network-instance"`
}

// SrlnokiaNetworkinstanceStatus struct
type SrlnokiaNetworkinstanceStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaNetworkinstanceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaNetworkinstance is the Schema for the SrlnokiaNetworkinstances API
type SrlnokiaNetworkinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaNetworkinstanceSpec   `json:"spec,omitempty"`
	Status SrlnokiaNetworkinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaNetworkinstanceList contains a list of SrlnokiaNetworkinstances
type SrlnokiaNetworkinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaNetworkinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaNetworkinstance{}, &SrlnokiaNetworkinstanceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaNetworkinstance) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaNetworkinstance",
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

func (o *SrlnokiaNetworkinstance) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
