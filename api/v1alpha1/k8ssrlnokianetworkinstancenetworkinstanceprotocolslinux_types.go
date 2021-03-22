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
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxFinalizer string = "Linux.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct {
	// +kubebuilder:default:=true
	ExportNeighbors *bool `json:"export-neighbors,omitempty"`
	// +kubebuilder:default:=false
	ExportRoutes *bool `json:"export-routes,omitempty"`
	// +kubebuilder:default:=false
	ImportRoutes *bool `json:"import-routes,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec struct {
	SrlNokiaNetworkInstanceName                          *string                                               `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux *SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux `json:"linux"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxs API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxs
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux",
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

func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
