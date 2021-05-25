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
	// SrlRoutingpolicyPrefixsetFinalizer is the name of the finalizer added to
	// SrlRoutingpolicyPrefixset to block delete operations until the physical node can be
	// deprovisioned.
	SrlRoutingpolicyPrefixsetFinalizer string = "PrefixSet.srlinux.henderiw.be"
)

// RoutingpolicyPrefixsetPrefix struct
type RoutingpolicyPrefixsetPrefix struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9]+\.\.[0-9]+)|exact`
	MaskLengthRange         *string `json:"mask-length-range,omitempty"`
	IpPrefixMaskLengthRange *string `json:"ip-prefix-mask-length-range,omitempty"`
}

// RoutingpolicyPrefixset struct
type RoutingpolicyPrefixset struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name   *string                         `json:"name"`
	Prefix []*RoutingpolicyPrefixsetPrefix `json:"prefix,omitempty"`
}

// SrlRoutingpolicyPrefixsetSpec struct
type SrlRoutingpolicyPrefixsetSpec struct {
	SrlRoutingpolicyPrefixset *[]RoutingpolicyPrefixset `json:"prefix-set"`
}

// SrlRoutingpolicyPrefixsetStatus struct
type SrlRoutingpolicyPrefixsetStatus struct {
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
	UsedSpec *SrlRoutingpolicyPrefixsetSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlRoutingpolicyPrefixset is the Schema for the SrlRoutingpolicyPrefixsets API
type SrlRoutingpolicyPrefixset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlRoutingpolicyPrefixsetSpec   `json:"spec,omitempty"`
	Status SrlRoutingpolicyPrefixsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPrefixsetList contains a list of SrlRoutingpolicyPrefixsets
type SrlRoutingpolicyPrefixsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyPrefixset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyPrefixset{}, &SrlRoutingpolicyPrefixsetList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlRoutingpolicyPrefixset) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlRoutingpolicyPrefixset",
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

func (o *SrlRoutingpolicyPrefixset) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlRoutingpolicyPrefixset) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
