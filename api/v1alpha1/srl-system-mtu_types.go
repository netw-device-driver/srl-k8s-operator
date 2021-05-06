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
	// SrlSystemMtuFinalizer is the name of the finalizer added to
	// SrlSystemMtu to block delete operations until the physical node can be
	// deprovisioned.
	SrlSystemMtuFinalizer string = "SystemMtu.srlinux.henderiw.be"
)

// SystemMtu struct
type SystemMtu struct {
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultPortMtu *uint16 `json:"default-port-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=552
	// +kubebuilder:validation:Maximum=9232
	// +kubebuilder:default:=552
	MinPathMtu *uint16 `json:"min-path-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=1280
	// +kubebuilder:validation:Maximum=9486
	// +kubebuilder:default:=1500
	DefaultIpMtu *uint16 `json:"default-ip-mtu,omitempty"`
	// +kubebuilder:validation:Minimum=1500
	// +kubebuilder:validation:Maximum=9500
	// +kubebuilder:default:=9232
	DefaultL2Mtu *uint16 `json:"default-l2-mtu,omitempty"`
}

// SrlSystemMtuSpec struct
type SrlSystemMtuSpec struct {
	SrlSystemMtu *SystemMtu `json:"mtu"`
}

// SrlSystemMtuStatus struct
type SrlSystemMtuStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlSystemMtuSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlSystemMtu is the Schema for the SrlSystemMtus API
type SrlSystemMtu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlSystemMtuSpec   `json:"spec,omitempty"`
	Status SrlSystemMtuStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemMtuList contains a list of SrlSystemMtus
type SrlSystemMtuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemMtu `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemMtu{}, &SrlSystemMtuList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlSystemMtu) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlSystemMtu",
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

func (o *SrlSystemMtu) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlSystemMtu) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
