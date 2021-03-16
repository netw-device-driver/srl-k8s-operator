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
	// SrlNokiaNetworkInstanceNetworkInstanceFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstance to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceFinalizer string = "NetworkInstance.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacDuplication struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacDuplication struct {
	// +kubebuilder:validation:Minimum=3
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default:=5
	NumMoves *uint32 `json:"num-moves,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`blackhole`;`oper-down`;`stop-learning`
	// +kubebuilder:default:=stop-learning
	Action *string `json:"action"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:default:=9
	HoldDownTime *uint32 `json:"hold-down-time,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=15
	// +kubebuilder:default:=3
	MonitoringWindow *uint32 `json:"monitoring-window,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearningAging struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearningAging struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=86400
	// +kubebuilder:default:=300
	AgeTime *int32 `json:"age-time,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearning struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearning struct {
	Aging *SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearningAging `json:"aging,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLimit struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLimit struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8192
	// +kubebuilder:default:=250
	MaximumEntries *int32 `json:"maximum-entries,omitempty"`
	// +kubebuilder:validation:Minimum=6
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=95
	WarningThresholdPct *int32 `json:"warning-threshold-pct,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMacMac struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMacMac struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Address     *string `json:"address"`
	Destination *string `json:"destination"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMac struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMac struct {
	Mac []*SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMacMac `json:"mac,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceBridgeTable struct
type SrlNokiaNetworkInstanceNetworkInstanceBridgeTable struct {
	MacDuplication *SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacDuplication `json:"mac-duplication,omitempty"`
	MacLearning    *SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLearning    `json:"mac-learning,omitempty"`
	MacLimit       *SrlNokiaNetworkInstanceNetworkInstanceBridgeTableMacLimit       `json:"mac-limit,omitempty"`
	// +kubebuilder:default:=false
	ProtectAnycastGwMac *bool                                                       `json:"protect-anycast-gw-mac,omitempty"`
	StaticMac           *SrlNokiaNetworkInstanceNetworkInstanceBridgeTableStaticMac `json:"static-mac,omitempty"`
	// +kubebuilder:default:=false
	DiscardUnknownDestMac *bool `json:"discard-unknown-dest-mac,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceInterface struct
type SrlNokiaNetworkInstanceNetworkInstanceInterface struct {
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0\.0|system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Name *string `json:"name"`
}

// SrlNokiaNetworkInstanceNetworkInstanceIpForwarding struct
type SrlNokiaNetworkInstanceNetworkInstanceIpForwarding struct {
	ReceiveIpv4Check *bool `json:"receive-ipv4-check,omitempty"`
	ReceiveIpv6Check *bool `json:"receive-ipv6-check,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancingResilientHashPrefix struct
type SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancingResilientHashPrefix struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
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

// SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancing struct
type SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancing struct {
	ResilientHashPrefix []*SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancingResilientHashPrefix `json:"resilient-hash-prefix,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceMplsStaticMplsEntry struct
type SrlNokiaNetworkInstanceNetworkInstanceMplsStaticMplsEntry struct {
	NextHopGroup *string `json:"next-hop-group,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`pop`;`swap`
	// +kubebuilder:default:=swap
	Operation *string `json:"operation"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=5
	Preference *uint8  `json:"preference,omitempty"`
	TopLabel   *string `json:"top-label"`
	// +kubebuilder:default:=false
	CollectStats *bool `json:"collect-stats,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceMpls struct
type SrlNokiaNetworkInstanceNetworkInstanceMpls struct {
	StaticMplsEntry []*SrlNokiaNetworkInstanceNetworkInstanceMplsStaticMplsEntry `json:"static-mpls-entry,omitempty"`
	// +kubebuilder:default:=false
	TtlPropagation *bool `json:"ttl-propagation,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state"`
}

// SrlNokiaNetworkInstanceNetworkInstanceMtu struct
type SrlNokiaNetworkInstanceNetworkInstanceMtu struct {
	PathMtuDiscovery *bool `json:"path-mtu-discovery,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroupsGroup struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	BitPosition *uint32 `json:"bit-position,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroups struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroups struct {
	Group []*SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroupsGroup `json:"group,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterfaceDelay struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterfaceDelay struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Static *uint32 `json:"static,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterface struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterface struct {
	Delay          *SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterfaceDelay `json:"delay,omitempty"`
	SrlgMembership *string                                                                 `json:"srlg-membership,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	TeMetric      *uint32 `json:"te-metric,omitempty"`
	InterfaceName *string `json:"interface-name"`
	AdminGroup    *string `json:"admin-group,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	ToAddress *string `json:"to-address,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	FromAddress *string `json:"from-address"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroup struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Value *uint32 `json:"value,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Cost         *uint32                                                                                          `json:"cost,omitempty"`
	StaticMember []*SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroupStaticMember `json:"static-member,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroups struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroups struct {
	Group []*SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroupsGroup `json:"group,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineering struct
type SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineering struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Ipv4TeRouterId *string `json:"ipv4-te-router-id,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Ipv6TeRouterId       *string                                                                       `json:"ipv6-te-router-id,omitempty"`
	SharedRiskLinkGroups *SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringSharedRiskLinkGroups `json:"shared-risk-link-groups,omitempty"`
	AdminGroups          *SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringAdminGroups          `json:"admin-groups,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AutonomousSystem *uint32                                                              `json:"autonomous-system,omitempty"`
	Interface        []*SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineeringInterface `json:"interface,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceVxlanInterface struct
type SrlNokiaNetworkInstanceNetworkInstanceVxlanInterface struct {
	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=17
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,8}))`
	Name *string `json:"name"`
}

// SrlNokiaNetworkInstanceNetworkInstance struct
type SrlNokiaNetworkInstanceNetworkInstance struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	RouterId       *string                                                 `json:"router-id,omitempty"`
	VxlanInterface []*SrlNokiaNetworkInstanceNetworkInstanceVxlanInterface `json:"vxlan-interface,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState   *string                                             `json:"admin-state"`
	BridgeTable  *SrlNokiaNetworkInstanceNetworkInstanceBridgeTable  `json:"bridge-table,omitempty"`
	Interface    []*SrlNokiaNetworkInstanceNetworkInstanceInterface  `json:"interface,omitempty"`
	IpForwarding *SrlNokiaNetworkInstanceNetworkInstanceIpForwarding `json:"ip-forwarding,omitempty"`
	// +kubebuilder:default:=default
	Type *string `json:"type,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Description        *string                                                   `json:"description,omitempty"`
	IpLoadBalancing    *SrlNokiaNetworkInstanceNetworkInstanceIpLoadBalancing    `json:"ip-load-balancing,omitempty"`
	Mpls               *SrlNokiaNetworkInstanceNetworkInstanceMpls               `json:"mpls,omitempty"`
	Mtu                *SrlNokiaNetworkInstanceNetworkInstanceMtu                `json:"mtu,omitempty"`
	TrafficEngineering *SrlNokiaNetworkInstanceNetworkInstanceTrafficEngineering `json:"traffic-engineering,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceSpec struct {
	SrlNokiaNetworkInstanceNetworkInstance *[]SrlNokiaNetworkInstanceNetworkInstance `json:"network-instance"`
}

// SrlNokiaNetworkInstanceNetworkInstanceStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstance is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstances API
type K8sSrlNokiaNetworkInstanceNetworkInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstances
type K8sSrlNokiaNetworkInstanceNetworkInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstance{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceList{})
}
