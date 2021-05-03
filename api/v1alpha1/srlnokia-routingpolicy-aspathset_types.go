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
	// SrlnokiaRoutingpolicyAspathsetFinalizer is the name of the finalizer added to
	// SrlnokiaRoutingpolicyAspathset to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaRoutingpolicyAspathsetFinalizer string = "RoutingpolicyAspathset.srlinux.henderiw.be"
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

// SrlnokiaRoutingpolicyAspathsetSpec struct
type SrlnokiaRoutingpolicyAspathsetSpec struct {
	SrlnokiaRoutingpolicyAspathset *[]RoutingpolicyAspathset `json:"as-path-set"`
}

// SrlnokiaRoutingpolicyAspathsetStatus struct
type SrlnokiaRoutingpolicyAspathsetStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaRoutingpolicyAspathsetSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaRoutingpolicyAspathset is the Schema for the SrlnokiaRoutingpolicyAspathsets API
type SrlnokiaRoutingpolicyAspathset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaRoutingpolicyAspathsetSpec   `json:"spec,omitempty"`
	Status SrlnokiaRoutingpolicyAspathsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaRoutingpolicyAspathsetList contains a list of SrlnokiaRoutingpolicyAspathsets
type SrlnokiaRoutingpolicyAspathsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaRoutingpolicyAspathset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaRoutingpolicyAspathset{}, &SrlnokiaRoutingpolicyAspathsetList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaRoutingpolicyAspathset) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaRoutingpolicyAspathset",
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

func (o *SrlnokiaRoutingpolicyAspathset) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlnokiaRoutingpolicyAspathset) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
