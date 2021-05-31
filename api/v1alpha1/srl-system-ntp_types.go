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
	// SrlSystemNtpFinalizer is the name of the finalizer added to
	// SrlSystemNtp to block delete operations until the physical node can be
	// deprovisioned.
	SrlSystemNtpFinalizer string = "Ntp.srlinux.henderiw.be"
)

// SystemNtpServer struct
type SystemNtpServer struct {
	// +kubebuilder:default:=false
	Iburst *bool `json:"iburst,omitempty"`
	// +kubebuilder:default:=false
	Prefer *bool `json:"prefer,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address"`
}

// SystemNtp struct
type SystemNtp struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState      *string            `json:"admin-state,omitempty"`
	NetworkInstance *string            `json:"network-instance"`
	Server          []*SystemNtpServer `json:"server,omitempty"`
}

// SrlSystemNtpSpec struct
type SrlSystemNtpSpec struct {
	SrlSystemNtp *SystemNtp `json:"ntp"`
}

// SrlSystemNtpStatus struct
type SrlSystemNtpStatus struct {
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
	UsedSpec *SrlSystemNtpSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlSystemNtp is the Schema for the SrlSystemNtps API
type SrlSystemNtp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlSystemNtpSpec   `json:"spec,omitempty"`
	Status SrlSystemNtpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNtpList contains a list of SrlSystemNtps
type SrlSystemNtpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNtp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNtp{}, &SrlSystemNtpList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlSystemNtp) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlSystemNtp",
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

func (o *SrlSystemNtp) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlSystemNtp) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
