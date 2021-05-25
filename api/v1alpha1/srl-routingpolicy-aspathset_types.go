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
	// SrlRoutingpolicyAspathsetFinalizer is the name of the finalizer added to
	// SrlRoutingpolicyAspathset to block delete operations until the physical node can be
	// deprovisioned.
	SrlRoutingpolicyAspathsetFinalizer string = "AsPathSet.srlinux.henderiw.be"
)

// RoutingpolicyAspathset struct
type RoutingpolicyAspathset struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=65535
	Expression *string `json:"expression,omitempty"`
}

// SrlRoutingpolicyAspathsetSpec struct
type SrlRoutingpolicyAspathsetSpec struct {
	SrlRoutingpolicyAspathset *[]RoutingpolicyAspathset `json:"as-path-set"`
}

// SrlRoutingpolicyAspathsetStatus struct
type SrlRoutingpolicyAspathsetStatus struct {
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
	UsedSpec *SrlRoutingpolicyAspathsetSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlRoutingpolicyAspathset is the Schema for the SrlRoutingpolicyAspathsets API
type SrlRoutingpolicyAspathset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlRoutingpolicyAspathsetSpec   `json:"spec,omitempty"`
	Status SrlRoutingpolicyAspathsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyAspathsetList contains a list of SrlRoutingpolicyAspathsets
type SrlRoutingpolicyAspathsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyAspathset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyAspathset{}, &SrlRoutingpolicyAspathsetList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlRoutingpolicyAspathset) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlRoutingpolicyAspathset",
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

func (o *SrlRoutingpolicyAspathset) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlRoutingpolicyAspathset) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
