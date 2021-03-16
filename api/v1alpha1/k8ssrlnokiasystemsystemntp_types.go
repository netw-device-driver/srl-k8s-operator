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
)

const (
	// SrlNokiaSystemSystemNtpFinalizer is the name of the finalizer added to
	// SrlNokiaSystemSystemNtp to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaSystemSystemNtpFinalizer string = "Ntp.srlinux.henderiw.be"
)

// SrlNokiaSystemSystemNtpServer struct
type SrlNokiaSystemSystemNtpServer struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	Address *string `json:"address"`
	// +kubebuilder:default:=false
	Iburst *bool `json:"iburst,omitempty"`
	// +kubebuilder:default:=false
	Prefer *bool `json:"prefer,omitempty"`
}

// SrlNokiaSystemSystemNtp struct
type SrlNokiaSystemSystemNtp struct {
	Server []*SrlNokiaSystemSystemNtpServer `json:"server,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=`disable`;`enable`
	AdminState      *string `json:"admin-state"`
	NetworkInstance *string `json:"network-instance"`
}

// SrlNokiaSystemSystemNtpSpec struct
type SrlNokiaSystemSystemNtpSpec struct {
	SrlNokiaSystemSystemNtp *SrlNokiaSystemSystemNtp `json:"ntp"`
}

// SrlNokiaSystemSystemNtpStatus struct
type SrlNokiaSystemSystemNtpStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaSystemSystemNtp is the Schema for the K8sSrlNokiaSystemSystemNtps API
type K8sSrlNokiaSystemSystemNtp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaSystemSystemNtpSpec   `json:"spec,omitempty"`
	Status SrlNokiaSystemSystemNtpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaSystemSystemNtpList contains a list of K8sSrlNokiaSystemSystemNtps
type K8sSrlNokiaSystemSystemNtpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaSystemSystemNtp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaSystemSystemNtp{}, &K8sSrlNokiaSystemSystemNtpList{})
}
