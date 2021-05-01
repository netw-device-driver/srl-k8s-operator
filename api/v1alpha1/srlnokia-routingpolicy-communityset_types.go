
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
	// SrlnokiaRoutingpolicyCommunitysetFinalizer is the name of the finalizer added to
	// SrlnokiaRoutingpolicyCommunityset to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaRoutingpolicyCommunitysetFinalizer string = "RoutingpolicyCommunityset.srlinux.henderiw.be"
)
// RoutingpolicyCommunityset struct
type RoutingpolicyCommunityset struct {
  // +kubebuilder:validation:MinLength=1
  // +kubebuilder:validation:MaxLength=255
  // +kubebuilder:validation:Required
  // +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
  Name *string `json:"name"`
  // +kubebuilder:validation:Optional
  // +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
  Member *string `json:"member,omitempty"`
}

// SrlnokiaRoutingpolicyCommunitysetSpec struct
type SrlnokiaRoutingpolicyCommunitysetSpec struct{
  SrlnokiaRoutingpolicyCommunityset       *[]RoutingpolicyCommunityset `json:"community-set"`
}

// SrlnokiaRoutingpolicyCommunitysetStatus struct
type SrlnokiaRoutingpolicyCommunitysetStatus struct{
  // Target provides the status of the configuration on the device
  Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

  // UsedSpec provides the spec used for the configuration
  UsedSpec *SrlnokiaRoutingpolicyCommunitysetSpec `json:"usedSpec,omitempty"`

  // LastUpdated identifies when this status was last observed.
  // +optional
  LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaRoutingpolicyCommunityset is the Schema for the SrlnokiaRoutingpolicyCommunitysets API
type SrlnokiaRoutingpolicyCommunityset struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   SrlnokiaRoutingpolicyCommunitysetSpec   `json:"spec,omitempty"`
  Status SrlnokiaRoutingpolicyCommunitysetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaRoutingpolicyCommunitysetList contains a list of SrlnokiaRoutingpolicyCommunitysets
type SrlnokiaRoutingpolicyCommunitysetList struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []SrlnokiaRoutingpolicyCommunityset `json:"items"`
}

func init() {
  SchemeBuilder.Register(&SrlnokiaRoutingpolicyCommunityset{}, &SrlnokiaRoutingpolicyCommunitysetList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaRoutingpolicyCommunityset) NewEvent(reason, message string) corev1.Event {
  t := metav1.Now()
  return corev1.Event{
    ObjectMeta: metav1.ObjectMeta{
      GenerateName: reason + "-",
      Namespace:    o.ObjectMeta.Namespace,
    },
    InvolvedObject: corev1.ObjectReference{
      Kind:       "SrlnokiaRoutingpolicyCommunityset",
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

func (o *SrlnokiaRoutingpolicyCommunityset) SetConfigStatus(t *string, c *ConfigStatus) {
  o.Status.Target[*t].ConfigStatus = c
}
