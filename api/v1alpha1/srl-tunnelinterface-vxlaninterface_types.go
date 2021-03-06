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
	// SrlTunnelinterfaceVxlaninterfaceFinalizer is the name of the finalizer added to
	// SrlTunnelinterfaceVxlaninterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlTunnelinterfaceVxlaninterfaceFinalizer string = "VxlanInterface.srlinux.henderiw.be"
)

// TunnelinterfaceVxlaninterfaceBridgeTable struct
type TunnelinterfaceVxlaninterfaceBridgeTable struct {
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	DestinationMac *string `json:"destination-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination struct {
	InnerEthernetHeader *TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestinationInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState  *string                                                                 `json:"admin-state,omitempty"`
	Destination []*TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroupDestination `json:"destination,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi *string `json:"esi,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressDestinationGroups struct
type TunnelinterfaceVxlaninterfaceEgressDestinationGroups struct {
	Group []*TunnelinterfaceVxlaninterfaceEgressDestinationGroupsGroup `json:"group,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader struct
type TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader struct {
	// +kubebuilder:default:=use-system-mac
	SourceMac *string `json:"source-mac,omitempty"`
}

// TunnelinterfaceVxlaninterfaceEgress struct
type TunnelinterfaceVxlaninterfaceEgress struct {
	DestinationGroups   *TunnelinterfaceVxlaninterfaceEgressDestinationGroups   `json:"destination-groups,omitempty"`
	InnerEthernetHeader *TunnelinterfaceVxlaninterfaceEgressInnerEthernetHeader `json:"inner-ethernet-header,omitempty"`
	// +kubebuilder:default:=use-system-ipv4-address
	SourceIp *string `json:"source-ip,omitempty"`
}

// TunnelinterfaceVxlaninterfaceIngress struct
type TunnelinterfaceVxlaninterfaceIngress struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16777215
	Vni *uint32 `json:"vni"`
}

// TunnelinterfaceVxlaninterface struct
type TunnelinterfaceVxlaninterface struct {
	Egress  *TunnelinterfaceVxlaninterfaceEgress  `json:"egress,omitempty"`
	Ingress *TunnelinterfaceVxlaninterfaceIngress `json:"ingress,omitempty"`
	Type    *string                               `json:"type"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=99999999
	Index       *uint32                                   `json:"index"`
	BridgeTable *TunnelinterfaceVxlaninterfaceBridgeTable `json:"bridge-table,omitempty"`
}

// SrlTunnelinterfaceVxlaninterfaceSpec struct
type SrlTunnelinterfaceVxlaninterfaceSpec struct {
	SrlNokiaTunnelInterfaceName      *string                          `json:"tunnel-interface-name"`
	SrlTunnelinterfaceVxlaninterface *[]TunnelinterfaceVxlaninterface `json:"vxlan-interface"`
}

// SrlTunnelinterfaceVxlaninterfaceStatus struct
type SrlTunnelinterfaceVxlaninterfaceStatus struct {
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
	UsedSpec *SrlTunnelinterfaceVxlaninterfaceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlTunnelinterfaceVxlaninterface is the Schema for the SrlTunnelinterfaceVxlaninterfaces API
type SrlTunnelinterfaceVxlaninterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlTunnelinterfaceVxlaninterfaceSpec   `json:"spec,omitempty"`
	Status SrlTunnelinterfaceVxlaninterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceVxlaninterfaceList contains a list of SrlTunnelinterfaceVxlaninterfaces
type SrlTunnelinterfaceVxlaninterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterfaceVxlaninterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterfaceVxlaninterface{}, &SrlTunnelinterfaceVxlaninterfaceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlTunnelinterfaceVxlaninterface) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlTunnelinterfaceVxlaninterface",
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

func (o *SrlTunnelinterfaceVxlaninterface) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlTunnelinterfaceVxlaninterface) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
