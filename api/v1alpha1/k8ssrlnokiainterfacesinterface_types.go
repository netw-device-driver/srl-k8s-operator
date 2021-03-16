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
	// SrlNokiaInterfacesInterfaceFinalizer is the name of the finalizer added to
	// SrlNokiaInterfacesInterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaInterfacesInterfaceFinalizer string = "Interface.srlinux.henderiw.be"
)

// SrlNokiaInterfacesInterfaceEthernetFlowControl struct
type SrlNokiaInterfacesInterfaceEthernetFlowControl struct {
	Receive  *bool `json:"receive,omitempty"`
	Transmit *bool `json:"transmit,omitempty"`
}

// SrlNokiaInterfacesInterfaceEthernet struct
type SrlNokiaInterfacesInterfaceEthernet struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	LacpPortPriority *uint16 `json:"lacp-port-priority,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`1T`;`200G`;`25G`;`400G`;`40G`;`50G`
	PortSpeed     *string `json:"port-speed"`
	AggregateId   *string `json:"aggregate-id,omitempty"`
	AutoNegotiate *bool   `json:"auto-negotiate,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`full`;`half`
	DuplexMode  *string                                         `json:"duplex-mode"`
	FlowControl *SrlNokiaInterfacesInterfaceEthernetFlowControl `json:"flow-control,omitempty"`
}

// SrlNokiaInterfacesInterfaceLagLacp struct
type SrlNokiaInterfacesInterfaceLagLacp struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	SystemPriority *uint16 `json:"system-priority,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	AdminKey *uint16 `json:"admin-key,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`FAST`;`SLOW`
	// +kubebuilder:default:=SLOW
	Interval *string `json:"interval"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`ACTIVE`;`PASSIVE`
	// +kubebuilder:default:=ACTIVE
	LacpMode *string `json:"lacp-mode"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	SystemIdMac *string `json:"system-id-mac,omitempty"`
}

// SrlNokiaInterfacesInterfaceLag struct
type SrlNokiaInterfacesInterfaceLag struct {
	Lacp *SrlNokiaInterfacesInterfaceLagLacp `json:"lacp,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`static`
	LacpFallbackMode *string `json:"lacp-fallback-mode"`
	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=3600
	LacpFallbackTimeout *uint16 `json:"lacp-fallback-timeout,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`lacp`;`static`
	// +kubebuilder:default:=static
	LagType *string `json:"lag-type"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`25G`;`400G`;`40G`
	MemberSpeed *string `json:"member-speed"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MinLinks *uint16 `json:"min-links,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputMulticastQueueScheduling struct
type SrlNokiaInterfacesInterfaceQosOutputMulticastQueueScheduling struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	PeakRatePercent *uint8 `json:"peak-rate-percent,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputMulticastQueue struct
type SrlNokiaInterfacesInterfaceQosOutputMulticastQueue struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	QueueId    *uint8                                                        `json:"queue-id"`
	Scheduling *SrlNokiaInterfacesInterfaceQosOutputMulticastQueueScheduling `json:"scheduling,omitempty"`
	Template   *string                                                       `json:"template,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputSchedulerTierNode struct
type SrlNokiaInterfacesInterfaceQosOutputSchedulerTierNode struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=11
	NodeNumber     *uint8 `json:"node-number"`
	StrictPriority *bool  `json:"strict-priority,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputSchedulerTier struct
type SrlNokiaInterfacesInterfaceQosOutputSchedulerTier struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4
	Level *uint8                                                   `json:"level"`
	Node  []*SrlNokiaInterfacesInterfaceQosOutputSchedulerTierNode `json:"node,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputScheduler struct
type SrlNokiaInterfacesInterfaceQosOutputScheduler struct {
	Tier []*SrlNokiaInterfacesInterfaceQosOutputSchedulerTier `json:"tier,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputUnicastQueueScheduling struct
type SrlNokiaInterfacesInterfaceQosOutputUnicastQueueScheduling struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	PeakRatePercent *uint8 `json:"peak-rate-percent,omitempty"`
	// +kubebuilder:default:=true
	StrictPriority *bool `json:"strict-priority,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutputUnicastQueue struct
type SrlNokiaInterfacesInterfaceQosOutputUnicastQueue struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	QueueId     *uint8                                                      `json:"queue-id"`
	Scheduling  *SrlNokiaInterfacesInterfaceQosOutputUnicastQueueScheduling `json:"scheduling,omitempty"`
	Template    *string                                                     `json:"template,omitempty"`
	VoqTemplate *string                                                     `json:"voq-template,omitempty"`
}

// SrlNokiaInterfacesInterfaceQosOutput struct
type SrlNokiaInterfacesInterfaceQosOutput struct {
	Scheduler      *SrlNokiaInterfacesInterfaceQosOutputScheduler        `json:"scheduler,omitempty"`
	UnicastQueue   []*SrlNokiaInterfacesInterfaceQosOutputUnicastQueue   `json:"unicast-queue,omitempty"`
	MulticastQueue []*SrlNokiaInterfacesInterfaceQosOutputMulticastQueue `json:"multicast-queue,omitempty"`
}

// SrlNokiaInterfacesInterfaceQos struct
type SrlNokiaInterfacesInterfaceQos struct {
	Output *SrlNokiaInterfacesInterfaceQosOutput `json:"output,omitempty"`
}

// SrlNokiaInterfacesInterfaceSflow struct
type SrlNokiaInterfacesInterfaceSflow struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state"`
}

// SrlNokiaInterfacesInterfaceTransceiver struct
type SrlNokiaInterfacesInterfaceTransceiver struct {
	DdmEvents *bool `json:"ddm-events,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`base-r`;`disabled`;`rs-108`;`rs-528`;`rs-544`
	ForwardErrorCorrection *string `json:"forward-error-correction"`
	TxLaser                *bool   `json:"tx-laser,omitempty"`
}

// SrlNokiaInterfacesInterface struct
type SrlNokiaInterfacesInterface struct {
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	Mtu         *uint16                           `json:"mtu,omitempty"`
	Sflow       *SrlNokiaInterfacesInterfaceSflow `json:"sflow,omitempty"`
	VlanTagging *bool                             `json:"vlan-tagging,omitempty"`
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0|mgmt0-standby|system0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name *string `json:"name"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Description  *string                         `json:"description,omitempty"`
	Lag          *SrlNokiaInterfacesInterfaceLag `json:"lag,omitempty"`
	LoopbackMode *bool                           `json:"loopback-mode,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                                 `json:"admin-state"`
	Ethernet    *SrlNokiaInterfacesInterfaceEthernet    `json:"ethernet,omitempty"`
	Qos         *SrlNokiaInterfacesInterfaceQos         `json:"qos,omitempty"`
	Transceiver *SrlNokiaInterfacesInterfaceTransceiver `json:"transceiver,omitempty"`
}

// SrlNokiaInterfacesInterfaceSpec struct
type SrlNokiaInterfacesInterfaceSpec struct {
	SrlNokiaInterfacesInterface *[]SrlNokiaInterfacesInterface `json:"interface"`
}

// SrlNokiaInterfacesInterfaceStatus struct
type SrlNokiaInterfacesInterfaceStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaInterfacesInterface is the Schema for the K8sSrlNokiaInterfacesInterfaces API
type K8sSrlNokiaInterfacesInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaInterfacesInterfaceSpec   `json:"spec,omitempty"`
	Status SrlNokiaInterfacesInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaInterfacesInterfaceList contains a list of K8sSrlNokiaInterfacesInterfaces
type K8sSrlNokiaInterfacesInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaInterfacesInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaInterfacesInterface{}, &K8sSrlNokiaInterfacesInterfaceList{})
}
