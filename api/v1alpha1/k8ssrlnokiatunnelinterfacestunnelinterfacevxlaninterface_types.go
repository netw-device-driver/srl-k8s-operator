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
	// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceFinalizer is the name of the finalizer added to
	// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceFinalizer string = "VxlanInterface.srlinux.henderiw.be"
)

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable struct {
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	DestinationMac *string `json:"destination-mac,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState          *string                                                                                                          `json:"admin-state,omitempty"`
	InnerEthernetHeader *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                                                                                         `json:"admin-state,omitempty"`
	Destination []*SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroupDestination `json:"destination,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups struct {
	Group []*SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroupsGroup `json:"group,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader struct {
	// +kubebuilder:default:=use-system-mac
	SourceMac *string `json:"source-mac,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress struct {
	// +kubebuilder:default:=use-system-ipv4-address
	SourceIp            *string                                                                         `json:"source-ip,omitempty"`
	DestinationGroups   *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressDestinationGroups   `json:"destination-groups,omitempty"`
	InnerEthernetHeader *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgressInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=99999999
	Index       *uint32                                                           `json:"index"`
	BridgeTable *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceBridgeTable `json:"bridge-table,omitempty"`
	Egress      *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceEgress      `json:"egress,omitempty"`
	Ingress     *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceIngress     `json:"ingress,omitempty"`
	Type        *string                                                           `json:"type"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceSpec struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceSpec struct {
	SrlNokiaTunnelInterfaceName                           *string                                                  `json:"tunnel-interface-name"`
	SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface *[]SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface `json:"vxlan-interface"`
}

// SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceStatus struct
type SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface is the Schema for the K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaces API
type K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceSpec   `json:"spec,omitempty"`
	Status SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceList contains a list of K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaces
type K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface{}, &K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterfaceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface",
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

func (o *K8sSrlNokiaTunnelInterfacesTunnelInterfaceVxlanInterface) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
