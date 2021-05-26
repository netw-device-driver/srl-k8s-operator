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
	// SrlNetworkinstanceAggregateroutesFinalizer is the name of the finalizer added to
	// SrlNetworkinstanceAggregateroutes to block delete operations until the physical node can be
	// deprovisioned.
	SrlNetworkinstanceAggregateroutesFinalizer string = "AggregateRoutes.srlinux.henderiw.be"
)

// NetworkinstanceAggregateroutesRouteAggregator struct
type NetworkinstanceAggregateroutesRouteAggregator struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber *uint32 `json:"as-number,omitempty"`
}

// NetworkinstanceAggregateroutesRouteCommunities struct
type NetworkinstanceAggregateroutesRouteCommunities struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])|.*:.*|([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})|.*:.*:.*`
	Add *string `json:"add,omitempty"`
}

// NetworkinstanceAggregateroutesRoute struct
type NetworkinstanceAggregateroutesRoute struct {
	Communities  *NetworkinstanceAggregateroutesRouteCommunities `json:"communities,omitempty"`
	GenerateIcmp *bool                                           `json:"generate-icmp,omitempty"`
	// +kubebuilder:default:=false
	SummaryOnly *bool `json:"summary-only,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                        `json:"admin-state,omitempty"`
	Aggregator *NetworkinstanceAggregateroutesRouteAggregator `json:"aggregator,omitempty"`
}

// NetworkinstanceAggregateroutes struct
type NetworkinstanceAggregateroutes struct {
	Route []*NetworkinstanceAggregateroutesRoute `json:"route,omitempty"`
}

// SrlNetworkinstanceAggregateroutesSpec struct
type SrlNetworkinstanceAggregateroutesSpec struct {
	SrlNokiaNetworkInstanceName       *string                         `json:"network-instance-name"`
	SrlNetworkinstanceAggregateroutes *NetworkinstanceAggregateroutes `json:"aggregate-routes"`
}

// SrlNetworkinstanceAggregateroutesStatus struct
type SrlNetworkinstanceAggregateroutesStatus struct {
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
	UsedSpec *SrlNetworkinstanceAggregateroutesSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlNetworkinstanceAggregateroutes is the Schema for the SrlNetworkinstanceAggregateroutess API
type SrlNetworkinstanceAggregateroutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNetworkinstanceAggregateroutesSpec   `json:"spec,omitempty"`
	Status SrlNetworkinstanceAggregateroutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceAggregateroutesList contains a list of SrlNetworkinstanceAggregateroutess
type SrlNetworkinstanceAggregateroutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceAggregateroutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceAggregateroutes{}, &SrlNetworkinstanceAggregateroutesList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlNetworkinstanceAggregateroutes) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNetworkinstanceAggregateroutes",
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

func (o *SrlNetworkinstanceAggregateroutes) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlNetworkinstanceAggregateroutes) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
