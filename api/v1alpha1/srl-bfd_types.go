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
	// SrlBfdFinalizer is the name of the finalizer added to
	// SrlBfd to block delete operations until the physical node can be
	// deprovisioned.
	SrlBfdFinalizer string = "Bfd.srlinux.henderiw.be"
)

// BfdMicroBfdSessionsLagInterface struct
type BfdMicroBfdSessionsLagInterface struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RemoteAddress *string `json:"remote-address,omitempty"`
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
	Name                   *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	DesiredMinimumTransmitInterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// +kubebuilder:validation:Minimum=3
	// +kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	DetectionMultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
}

// BfdMicroBfdSessions struct
type BfdMicroBfdSessions struct {
	LagInterface []*BfdMicroBfdSessionsLagInterface `json:"lag-interface,omitempty"`
}

// BfdSubinterface struct
type BfdSubinterface struct {
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:MaxLength=25
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
	Id *string `json:"id"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	DesiredMinimumTransmitInterval *uint32 `json:"desired-minimum-transmit-interval,omitempty"`
	// +kubebuilder:validation:Minimum=3
	// +kubebuilder:validation:Maximum=20
	// +kubebuilder:default:=3
	DetectionMultiplier *uint8 `json:"detection-multiplier,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=0
	MinimumEchoReceiveInterval *uint32 `json:"minimum-echo-receive-interval,omitempty"`
}

// Bfd struct
type Bfd struct {
	MicroBfdSessions *BfdMicroBfdSessions `json:"micro-bfd-sessions,omitempty"`
	Subinterface     []*BfdSubinterface   `json:"subinterface,omitempty"`
}

// SrlBfdSpec struct
type SrlBfdSpec struct {
	SrlBfd *Bfd `json:"bfd"`
}

// SrlBfdStatus struct
type SrlBfdStatus struct {
	// ConfigurationDependencyTargetNotFound identifies if the target of the resource object is missing or not
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyTargetFound *TargetFoundStatus `json:"configurationDependencyTargetFound,omitempty"`

	// ConfigurationDependencyLocalLeafrefValidationStatus identifies the status of the local LeafRef Validation of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyLocalLeafrefValidationStatus *ValidationStatus `json:"configurationDependencyLocalLeafrefValidationStatus,omitempty"`

	// ConfigurationDependencyLocalLeafrefValidationDetails defines the validation details of the resource object
	ConfigurationDependencyLocalLeafrefValidationDetails map[string]*ValidationDetails2 `json:"localLeafrefValidationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlBfdSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlBfd is the Schema for the SrlBfds API
type SrlBfd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlBfdSpec   `json:"spec,omitempty"`
	Status SrlBfdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlBfdList contains a list of SrlBfds
type SrlBfdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlBfd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlBfd{}, &SrlBfdList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlBfd) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlBfd",
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

func (o *SrlBfd) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlBfd) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
