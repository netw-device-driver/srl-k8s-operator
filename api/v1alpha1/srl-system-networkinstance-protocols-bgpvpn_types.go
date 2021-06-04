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
	// SrlSystemNetworkinstanceProtocolsBgpvpnFinalizer is the name of the finalizer added to
	// SrlSystemNetworkinstanceProtocolsBgpvpn to block delete operations until the physical node can be
	// deprovisioned.
	SrlSystemNetworkinstanceProtocolsBgpvpnFinalizer string = "BgpVpn.srlinux.henderiw.be"
)

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget struct {
}

// SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct
type SystemNetworkinstanceProtocolsBgpvpnBgpInstance struct {
	RouteTarget *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteTarget `json:"route-target,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Id                 *uint8                                                             `json:"id"`
	RouteDistinguisher *SystemNetworkinstanceProtocolsBgpvpnBgpInstanceRouteDistinguisher `json:"route-distinguisher,omitempty"`
}

// SystemNetworkinstanceProtocolsBgpvpn struct
type SystemNetworkinstanceProtocolsBgpvpn struct {
	BgpInstance []*SystemNetworkinstanceProtocolsBgpvpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlSystemNetworkinstanceProtocolsBgpvpnSpec struct
type SrlSystemNetworkinstanceProtocolsBgpvpnSpec struct {
	SrlSystemNetworkinstanceProtocolsBgpvpn *SystemNetworkinstanceProtocolsBgpvpn `json:"bgp-vpn"`
}

// SrlSystemNetworkinstanceProtocolsBgpvpnStatus struct
type SrlSystemNetworkinstanceProtocolsBgpvpnStatus struct {
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
	UsedSpec *SrlSystemNetworkinstanceProtocolsBgpvpnSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlSystemNetworkinstanceProtocolsBgpvpn is the Schema for the SrlSystemNetworkinstanceProtocolsBgpvpns API
type SrlSystemNetworkinstanceProtocolsBgpvpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlSystemNetworkinstanceProtocolsBgpvpnSpec   `json:"spec,omitempty"`
	Status SrlSystemNetworkinstanceProtocolsBgpvpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsBgpvpnList contains a list of SrlSystemNetworkinstanceProtocolsBgpvpns
type SrlSystemNetworkinstanceProtocolsBgpvpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsBgpvpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsBgpvpn{}, &SrlSystemNetworkinstanceProtocolsBgpvpnList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlSystemNetworkinstanceProtocolsBgpvpn) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlSystemNetworkinstanceProtocolsBgpvpn",
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

func (o *SrlSystemNetworkinstanceProtocolsBgpvpn) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlSystemNetworkinstanceProtocolsBgpvpn) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
