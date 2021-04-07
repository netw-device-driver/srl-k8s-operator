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
	// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesFinalizer string = "AggregateRoutes.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteAggregator struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteAggregator struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	Address *string `json:"address,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	AsNumber *uint32 `json:"as-number,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteCommunities struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteCommunities struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`.*:.*`
	// +kubebuilder:validation:Pattern=`([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})`
	// +kubebuilder:validation:Pattern=`.*:.*:.*`
	Add *string `json:"add,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRoute struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRoute struct {
	Communities  *SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteCommunities `json:"communities,omitempty"`
	GenerateIcmp *bool                                                                  `json:"generate-icmp,omitempty"`
	// +kubebuilder:default:=false
	SummaryOnly *bool `json:"summary-only,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix *string `json:"prefix"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                                               `json:"admin-state,omitempty"`
	Aggregator *SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRouteAggregator `json:"aggregator,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes struct {
	Route []*SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesRoute `json:"route,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesSpec struct {
	SrlNokiaNetworkInstanceName                           *string                                                `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes *SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes `json:"aggregate-routes"`
}

// SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutess API
type K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutess
type K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes",
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

func (o *K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
