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
	// SrlnokiaSystemNameFinalizer is the name of the finalizer added to
	// SrlnokiaSystemName to block delete operations until the physical node can be
	// deprovisioned.
	SrlnokiaSystemNameFinalizer string = "SystemName.srlinux.henderiw.be"
)

// SystemName struct
type SystemName struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.)*([a-zA-Z0-9_]([a-zA-Z0-9\-_]){0,61})?[a-zA-Z0-9]\.?)|\.`
	DomainName *string `json:"domain-name,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])`
	HostName *string `json:"host-name,omitempty"`
}

// SrlnokiaSystemNameSpec struct
type SrlnokiaSystemNameSpec struct {
	SrlnokiaSystemName *SystemName `json:"name"`
}

// SrlnokiaSystemNameStatus struct
type SrlnokiaSystemNameStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlnokiaSystemNameSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlnokiaSystemName is the Schema for the SrlnokiaSystemNames API
type SrlnokiaSystemName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlnokiaSystemNameSpec   `json:"spec,omitempty"`
	Status SrlnokiaSystemNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlnokiaSystemNameList contains a list of SrlnokiaSystemNames
type SrlnokiaSystemNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlnokiaSystemName `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlnokiaSystemName{}, &SrlnokiaSystemNameList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlnokiaSystemName) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlnokiaSystemName",
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

func (o *SrlnokiaSystemName) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlnokiaSystemName) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
