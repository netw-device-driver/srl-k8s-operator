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
	// SrlnokiaInterfaceFinalizer is the name of the finalizer added to
	// SrlnokiaInterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaInterfaceFinalizer string = "Interface.srlinux.henderiw.be"
)

// InterfaceEthernetFlowControl struct
type InterfaceEthernetFlowControl struct {
	Transmit *bool `json:"transmit,omitempty"`
	Receive  *bool `json:"receive,omitempty"`
}

// InterfaceEthernet struct
type InterfaceEthernet struct {
	AggregateId   *string `json:"aggregate-id,omitempty"`
	AutoNegotiate *bool   `json:"auto-negotiate,omitempty"`
	// +kubebuilder:validation:Enum=`full`;`half`
	DuplexMode  *string                       `json:"duplex-mode,omitempty"`
	FlowControl *InterfaceEthernetFlowControl `json:"flow-control,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	LacpPortPriority *uint16 `json:"lacp-port-priority,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`1T`;`200G`;`25G`;`400G`;`40G`;`50G`
	PortSpeed *string `json:"port-speed,omitempty"`
}

// InterfaceLagLacp struct
type InterfaceLagLacp struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	AdminKey *uint16 `json:"admin-key,omitempty"`
	// +kubebuilder:validation:Enum=`FAST`;`SLOW`
	// +kubebuilder:default:=SLOW
	Interval *string `json:"interval,omitempty"`
	// +kubebuilder:validation:Enum=`ACTIVE`;`PASSIVE`
	// +kubebuilder:default:=ACTIVE
	LacpMode *string `json:"lacp-mode,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	SystemIdMac *string `json:"system-id-mac,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	SystemPriority *uint16 `json:"system-priority,omitempty"`
}

// InterfaceLag struct
type InterfaceLag struct {
	Lacp *InterfaceLagLacp `json:"lacp,omitempty"`
	// +kubebuilder:validation:Enum=`static`
	LacpFallbackMode *string `json:"lacp-fallback-mode,omitempty"`
	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=3600
	LacpFallbackTimeout *uint16 `json:"lacp-fallback-timeout,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`static`
	// +kubebuilder:default:=static
	LagType *string `json:"lag-type,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`25G`;`400G`;`40G`
	MemberSpeed *string `json:"member-speed,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	MinLinks *uint16 `json:"min-links,omitempty"`
}

// InterfaceQosOutputMulticastQueueScheduling struct
type InterfaceQosOutputMulticastQueueScheduling struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	PeakRatePercent *uint8 `json:"peak-rate-percent,omitempty"`
}

// InterfaceQosOutputMulticastQueue struct
type InterfaceQosOutputMulticastQueue struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	QueueId    *uint8                                      `json:"queue-id"`
	Scheduling *InterfaceQosOutputMulticastQueueScheduling `json:"scheduling,omitempty"`
	Template   *string                                     `json:"template,omitempty"`
}

// InterfaceQosOutputSchedulerTierNode struct
type InterfaceQosOutputSchedulerTierNode struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=11
	NodeNumber     *uint8 `json:"node-number"`
	StrictPriority *bool  `json:"strict-priority,omitempty"`
}

// InterfaceQosOutputSchedulerTier struct
type InterfaceQosOutputSchedulerTier struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4
	Level *uint8                                 `json:"level"`
	Node  []*InterfaceQosOutputSchedulerTierNode `json:"node,omitempty"`
}

// InterfaceQosOutputScheduler struct
type InterfaceQosOutputScheduler struct {
	Tier []*InterfaceQosOutputSchedulerTier `json:"tier,omitempty"`
}

// InterfaceQosOutputUnicastQueueScheduling struct
type InterfaceQosOutputUnicastQueueScheduling struct {
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

// InterfaceQosOutputUnicastQueue struct
type InterfaceQosOutputUnicastQueue struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	QueueId     *uint8                                    `json:"queue-id"`
	Scheduling  *InterfaceQosOutputUnicastQueueScheduling `json:"scheduling,omitempty"`
	Template    *string                                   `json:"template,omitempty"`
	VoqTemplate *string                                   `json:"voq-template,omitempty"`
}

// InterfaceQosOutput struct
type InterfaceQosOutput struct {
	MulticastQueue []*InterfaceQosOutputMulticastQueue `json:"multicast-queue,omitempty"`
	Scheduler      *InterfaceQosOutputScheduler        `json:"scheduler,omitempty"`
	UnicastQueue   []*InterfaceQosOutputUnicastQueue   `json:"unicast-queue,omitempty"`
}

// InterfaceQos struct
type InterfaceQos struct {
	Output *InterfaceQosOutput `json:"output,omitempty"`
}

// InterfaceSflow struct
type InterfaceSflow struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState *string `json:"admin-state,omitempty"`
}

// InterfaceTransceiver struct
type InterfaceTransceiver struct {
	DdmEvents *bool `json:"ddm-events,omitempty"`
	// +kubebuilder:validation:Enum=`base-r`;`disabled`;`rs-108`;`rs-528`;`rs-544`
	// +kubebuilder:default:=disabled
	ForwardErrorCorrection *string `json:"forward-error-correction,omitempty"`
	TxLaser                *bool   `json:"tx-laser,omitempty"`
}

// Interface struct
type Interface struct {
	Transceiver *InterfaceTransceiver `json:"transceiver,omitempty"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description  *string       `json:"description,omitempty"`
	Lag          *InterfaceLag `json:"lag,omitempty"`
	LoopbackMode *bool         `json:"loopback-mode,omitempty"`
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	Mtu *uint16 `json:"mtu,omitempty"`
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0|mgmt0-standby|system0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name        *string            `json:"name"`
	Ethernet    *InterfaceEthernet `json:"ethernet,omitempty"`
	Qos         *InterfaceQos      `json:"qos,omitempty"`
	Sflow       *InterfaceSflow    `json:"sflow,omitempty"`
	VlanTagging *bool              `json:"vlan-tagging,omitempty"`
}

// SrlnokiaInterfaceSpec struct
type SrlnokiaInterfaceSpec struct {
	SrlnokiaInterface *[]Interface `json:"interface"`
}

// SrlnokiaInterfaceStatus struct
type SrlnokiaInterfaceStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaInterfaceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaInterface is the Schema for the SrlnokiaInterfaces API
type SrlnokiaInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaInterfaceSpec   `json:"spec,omitempty"`
	Status SrlnokiaInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaInterfaceList contains a list of SrlnokiaInterfaces
type SrlnokiaInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaInterface{}, &SrlnokiaInterfaceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaInterface) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaInterface",
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

func (o *SrlnokiaInterface) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
