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
	// SrlNokiaSystemSystemMtuFinalizer is the name of the finalizer added to
	// SrlNokiaSystemSystemMtu to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaSystemSystemMtuFinalizer string = "Mtu.srlinux.henderiw.be"
)

// SrlNokiaSystemSystemMtu struct
type SrlNokiaSystemSystemMtu struct {
	// +kubebuilder:validation:Minimum=1280
	// +kubebuilder:validation:Maximum=9486
	// +kubebuilder:default:=1500
	DefaultIpMtu *uint16 `json:"default-ip-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultL2Mtu *uint16 `json:"default-l2-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultPortMtu *uint16 `json:"default-port-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=552
	// +kubebuilder:validation:Maximum=9232
	// +kubebuilder:default:=552
	MinPathMtu *uint16 `json:"min-path-mtu,omitempty"`
}

// SrlNokiaSystemSystemMtuSpec struct
type SrlNokiaSystemSystemMtuSpec struct {
	SrlNokiaSystemSystemMtu *SrlNokiaSystemSystemMtu `json:"mtu"`
}

// SrlNokiaSystemSystemMtuStatus struct
type SrlNokiaSystemSystemMtuStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaSystemSystemMtuSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaSystemSystemMtu is the Schema for the K8sSrlNokiaSystemSystemMtus API
type K8sSrlNokiaSystemSystemMtu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaSystemSystemMtuSpec   `json:"spec,omitempty"`
	Status SrlNokiaSystemSystemMtuStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaSystemSystemMtuList contains a list of K8sSrlNokiaSystemSystemMtus
type K8sSrlNokiaSystemSystemMtuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaSystemSystemMtu `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaSystemSystemMtu{}, &K8sSrlNokiaSystemSystemMtuList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaSystemSystemMtu) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaSystemSystemMtu",
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

func (o *K8sSrlNokiaSystemSystemMtu) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
