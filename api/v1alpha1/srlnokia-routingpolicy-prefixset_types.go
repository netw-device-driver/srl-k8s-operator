
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
	// SrlnokiaRoutingpolicyPrefixsetFinalizer is the name of the finalizer added to
	// SrlnokiaRoutingpolicyPrefixset to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaRoutingpolicyPrefixsetFinalizer string = "RoutingpolicyPrefixset.srlinux.henderiw.be"
)
// RoutingpolicyPrefixsetPrefix struct
type RoutingpolicyPrefixsetPrefix struct {
  IpPrefixMaskLengthRange *string `json:"ip-prefix-mask-length-range,omitempty"`
  // +kubebuilder:validation:Optional
  // +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
  IpPrefix *string `json:"ip-prefix,omitempty"`
  // +kubebuilder:validation:Required
  // +kubebuilder:validation:Pattern=`([0-9]+\.\.[0-9]+)|exact`
  MaskLengthRange *string `json:"mask-length-range,omitempty"`
}
// RoutingpolicyPrefixset struct
type RoutingpolicyPrefixset struct {
  // +kubebuilder:validation:MinLength=1
  // +kubebuilder:validation:MaxLength=255
  // +kubebuilder:validation:Required
  // +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
  Name *string `json:"name"`
  Prefix []*RoutingpolicyPrefixsetPrefix `json:"prefix,omitempty"`
}

// SrlnokiaRoutingpolicyPrefixsetSpec struct
type SrlnokiaRoutingpolicyPrefixsetSpec struct{
  SrlnokiaRoutingpolicyPrefixset       *[]RoutingpolicyPrefixset `json:"prefix-set"`
}

// SrlnokiaRoutingpolicyPrefixsetStatus struct
type SrlnokiaRoutingpolicyPrefixsetStatus struct{
  // Target provides the status of the configuration on the device
  Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

  // UsedSpec provides the spec used for the configuration
  UsedSpec *SrlnokiaRoutingpolicyPrefixsetSpec `json:"usedSpec,omitempty"`

  // LastUpdated identifies when this status was last observed.
  // +optional
  LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaRoutingpolicyPrefixset is the Schema for the SrlnokiaRoutingpolicyPrefixsets API
type SrlnokiaRoutingpolicyPrefixset struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   SrlnokiaRoutingpolicyPrefixsetSpec   `json:"spec,omitempty"`
  Status SrlnokiaRoutingpolicyPrefixsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaRoutingpolicyPrefixsetList contains a list of SrlnokiaRoutingpolicyPrefixsets
type SrlnokiaRoutingpolicyPrefixsetList struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []SrlnokiaRoutingpolicyPrefixset `json:"items"`
}

func init() {
  SchemeBuilder.Register(&SrlnokiaRoutingpolicyPrefixset{}, &SrlnokiaRoutingpolicyPrefixsetList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaRoutingpolicyPrefixset) NewEvent(reason, message string) corev1.Event {
  t := metav1.Now()
  return corev1.Event{
    ObjectMeta: metav1.ObjectMeta{
      GenerateName: reason + "-",
      Namespace:    o.ObjectMeta.Namespace,
    },
    InvolvedObject: corev1.ObjectReference{
      Kind:       "SrlnokiaRoutingpolicyPrefixset",
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

func (o *SrlnokiaRoutingpolicyPrefixset) SetConfigStatus(t *string, c *ConfigStatus) {
  o.Status.Target[*t].ConfigStatus = c
}
