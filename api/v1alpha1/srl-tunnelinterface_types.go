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
	// SrlTunnelinterfaceFinalizer is the name of the finalizer added to
	// SrlTunnelinterface to block delete operations until the physical node can be
	// deprovisioned.
	SrlTunnelinterfaceFinalizer string = "TunnelInterface.srlinux.henderiw.be"
)

// Tunnelinterface struct
type Tunnelinterface struct {
	// +kubebuilder:validation:MinLength=6
	// +kubebuilder:validation:MaxLength=8
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(vxlan(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9]))`
	Name *string `json:"name"`
}

// SrlTunnelinterfaceSpec struct
type SrlTunnelinterfaceSpec struct {
	SrlTunnelinterface *[]Tunnelinterface `json:"tunnel-interface"`
}

// SrlTunnelinterfaceStatus struct
type SrlTunnelinterfaceStatus struct {
	// ConfigurationDependencyTargetNotFound identifies if the target of the resource object is missing or not
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyTargetFound *TargetFoundStatus `json:"configurationDependencyTargetFound,omitempty"`

	// ConfigurationDependencyValidationStatus identifies the status of the LeafRef Validation of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyValidationStatus *ValidationStatus `json:"configurationDependencyValidationStatus,omitempty"`

	// ConfigurationDependencyValidationDetails defines the validation details of the resource object
	ConfigurationDependencyValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlTunnelinterfaceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlTunnelinterface is the Schema for the SrlTunnelinterfaces API
type SrlTunnelinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlTunnelinterfaceSpec   `json:"spec,omitempty"`
	Status SrlTunnelinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlTunnelinterfaceList contains a list of SrlTunnelinterfaces
type SrlTunnelinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlTunnelinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlTunnelinterface{}, &SrlTunnelinterfaceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlTunnelinterface) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlTunnelinterface",
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

func (o *SrlTunnelinterface) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlTunnelinterface) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
