
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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	// SrlnokiaBfdFinalizer is the name of the finalizer added to
	// SrlnokiaBfd to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaBfdFinalizer string = "Bfd.srlinux.henderiw.be"
)
// BfdMicroBfdSessionsLagInterface struct
type BfdMicroBfdSessionsLagInterface struct {
  // +kubebuilder:validation:Minimum=10000
  // +kubebuilder:validation:Maximum=100000000
  // +kubebuilder:default:=1000000
  RequiredMinimumReceive *uint32 `json:"required-minimum-receive,omitempty"`
  Name *string `json:"name"`
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
  // +kubebuilder:validation:Optional
  // +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
  RemoteAddress *string `json:"remote-address,omitempty"`
}
// BfdMicroBfdSessions struct
type BfdMicroBfdSessions struct {
  LagInterface []*BfdMicroBfdSessionsLagInterface `json:"lag-interface,omitempty"`
}
// BfdSubinterface struct
type BfdSubinterface struct {
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
  // +kubebuilder:validation:MinLength=5
  // +kubebuilder:validation:MaxLength=25
  // +kubebuilder:validation:Required
  // +kubebuilder:validation:Pattern=`(system0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8]))\.(0|[1-9](\d){0,3}))`
  Id *string `json:"id"`
}
// Bfd struct
type Bfd struct {
  MicroBfdSessions *BfdMicroBfdSessions `json:"micro-bfd-sessions,omitempty"`
  Subinterface []*BfdSubinterface `json:"subinterface,omitempty"`
}

// SrlnokiaBfdSpec struct
type SrlnokiaBfdSpec struct{
  SrlnokiaBfd       *Bfd `json:"bfd"`
}

// SrlnokiaBfdStatus struct
type SrlnokiaBfdStatus struct{
  // Target provides the status of the configuration on the device
  Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

  // UsedSpec provides the spec used for the configuration
  UsedSpec *SrlnokiaBfdSpec `json:"usedSpec,omitempty"`

  // LastUpdated identifies when this status was last observed.
  // +optional
  LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaBfd is the Schema for the SrlnokiaBfds API
type SrlnokiaBfd struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   SrlnokiaBfdSpec   `json:"spec,omitempty"`
  Status SrlnokiaBfdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaBfdList contains a list of SrlnokiaBfds
type SrlnokiaBfdList struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []SrlnokiaBfd `json:"items"`
}

func init() {
  SchemeBuilder.Register(&SrlnokiaBfd{}, &SrlnokiaBfdList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaBfd) NewEvent(reason, message string) corev1.Event {
  t := metav1.Now()
  return corev1.Event{
    ObjectMeta: metav1.ObjectMeta{
      GenerateName: reason + "-",
      Namespace:    o.ObjectMeta.Namespace,
    },
    InvolvedObject: corev1.ObjectReference{
      Kind:       "SrlnokiaBfd",
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

func (o *SrlnokiaBfd) SetConfigStatus(t *string, c *ConfigStatus) {
  o.Status.Target[*t].ConfigStatus = c
}
