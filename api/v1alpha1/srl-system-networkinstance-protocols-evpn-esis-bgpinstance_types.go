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
	// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceFinalizer is the name of the finalizer added to
	// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance to block delete operations until the physical node can be
	// deprovisioned.
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceFinalizer string = "BgpInstance.srlinux.henderiw.be"
)

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	Id *string `json:"id"`
}

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec struct
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec struct {
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance *[]SystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"bgp-instance"`
}

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus struct
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance is the Schema for the SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstances API
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceSpec   `json:"spec,omitempty"`
	Status SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList contains a list of SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstances
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance",
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

func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
