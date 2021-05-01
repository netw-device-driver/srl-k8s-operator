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
	// SrlnokiaSystemNetworkinstanceProtocolsBgpvpnFinalizer is the name of the finalizer added to
	// SrlnokiaSystemNetworkinstanceProtocolsBgpvpn to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaSystemNetworkinstanceProtocolsBgpvpnFinalizer string = "SystemNetworkinstanceProtocolsBgpvpn.srlinux.henderiw.be"
)

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct {
	RouteDistinguisher *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher `json:"route-distinguisher,omitempty"`
	RouteTarget        *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget        `json:"route-target,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Id *uint8 `json:"id"`
}

// SystemNetworkinstanceProtocolsBgpvpn struct
type SystemNetworkinstanceProtocolsBgpvpn struct {
	BgpInstance []*SystemNetworkinstanceProtocolsBgpvpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlnokiaSystemNetworkinstanceProtocolsBgpvpnSpec struct
type SrlnokiaSystemNetworkinstanceProtocolsBgpvpnSpec struct {
	SrlnokiaSystemNetworkinstanceProtocolsBgpvpn *SystemNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn"`
}

// SrlnokiaSystemNetworkinstanceProtocolsBgpvpnStatus struct
type SrlnokiaSystemNetworkinstanceProtocolsBgpvpnStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaSystemNetworkinstanceProtocolsBgpvpnSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaSystemNetworkinstanceProtocolsBgpvpn is the Schema for the SrlnokiaSystemNetworkinstanceProtocolsBgpvpns API
type SrlnokiaSystemNetworkinstanceProtocolsBgpvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaSystemNetworkinstanceProtocolsBgpvpnSpec   `json:"spec,omitempty"`
	Status SrlnokiaSystemNetworkinstanceProtocolsBgpvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaSystemNetworkinstanceProtocolsBgpvpnList contains a list of SrlnokiaSystemNetworkinstanceProtocolsBgpvpns
type SrlnokiaSystemNetworkinstanceProtocolsBgpvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaSystemNetworkinstanceProtocolsBgpvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaSystemNetworkinstanceProtocolsBgpvpn{}, &SrlnokiaSystemNetworkinstanceProtocolsBgpvpnList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaSystemNetworkinstanceProtocolsBgpvpn) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaSystemNetworkinstanceProtocolsBgpvpn",
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

func (o *SrlnokiaSystemNetworkinstanceProtocolsBgpvpn) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
