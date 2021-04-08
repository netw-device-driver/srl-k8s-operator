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
	// SrlNokiaBfdBfdFinalizer is the name of the finalizer added to
	// SrlNokiaBfdBfd to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaBfdBfdFinalizer string = "Bfd.srlinux.henderiw.be"
)

// SrlNokiaBfdBfdMicroBfdSessionsLagInterface struct
type SrlNokiaBfdBfdMicroBfdSessionsLagInterface struct {
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
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RemoteAddress *string `json:"remote-address,omitempty"`
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
	Name                   *string `json:"name"`
}

// SrlNokiaBfdBfdMicroBfdSessions struct
type SrlNokiaBfdBfdMicroBfdSessions struct {
	LagInterface []*SrlNokiaBfdBfdMicroBfdSessionsLagInterface `json:"lag-interface,omitempty"`
}

// SrlNokiaBfdBfdSubinterface struct
type SrlNokiaBfdBfdSubinterface struct {
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
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=100000000
	// +kubebuilder:default:=1000000
	RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
}

// SrlNokiaBfdBfd struct
type SrlNokiaBfdBfd struct {
	MicroBfdSessions *SrlNokiaBfdBfdMicroBfdSessions `json:"micro-bfd-sessions,omitempty"`
	Subinterface     []*SrlNokiaBfdBfdSubinterface   `json:"subinterface,omitempty"`
}

// SrlNokiaBfdBfdSpec struct
type SrlNokiaBfdBfdSpec struct {
	SrlNokiaBfdBfd *SrlNokiaBfdBfd `json:"bfd"`
}

// SrlNokiaBfdBfdStatus struct
type SrlNokiaBfdBfdStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaBfdBfdSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaBfdBfd is the Schema for the K8sSrlNokiaBfdBfds API
type K8sSrlNokiaBfdBfd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaBfdBfdSpec   `json:"spec,omitempty"`
	Status SrlNokiaBfdBfdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaBfdBfdList contains a list of K8sSrlNokiaBfdBfds
type K8sSrlNokiaBfdBfdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaBfdBfd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaBfdBfd{}, &K8sSrlNokiaBfdBfdList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaBfdBfd) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaBfdBfd",
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

func (o *K8sSrlNokiaBfdBfd) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}