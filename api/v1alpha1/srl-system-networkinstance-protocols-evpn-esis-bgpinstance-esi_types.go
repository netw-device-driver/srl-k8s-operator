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
	// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiFinalizer is the name of the finalizer added to
	// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi to block delete operations until the physical node can be
	// deprovisioned.
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiFinalizer string = "EthernetSegment.srlinux.henderiw.be"
)

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg struct {
	Capabilities *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlgCapabilities `json:"capabilities,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities struct {
	// +kubebuilder:default:=true
	AcDf *bool `json:"ac-df,omitempty"`
	// +kubebuilder:default:=false
	NonRevertive *bool `json:"non-revertive,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg struct {
	Capabilities *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlgCapabilities `json:"capabilities,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=32767
	PreferenceValue *uint32 `json:"preference-value,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm struct {
	DefaultAlg    *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmDefaultAlg    `json:"default-alg,omitempty"`
	PreferenceAlg *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithmPreferenceAlg `json:"preference-alg,omitempty"`
	// +kubebuilder:validation:Enum=`default`;`preference`
	// +kubebuilder:default:=default
	Type *string `json:"type,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ActivationTimer *uint32 `json:"activation-timer,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection struct {
	Algorithm                        *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionAlgorithm `json:"algorithm,omitempty"`
	InterfaceStandbySignalingOnNonDf *bool                                                                    `json:"interface-standby-signaling-on-non-df,omitempty"`
	Timers                           *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElectionTimers    `json:"timers,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEsi struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEsi struct {
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	OriginatingIp *string `json:"originating-ip,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct {
	Esi *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEsi `json:"ethernet-segment,omitempty"`
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop *string `json:"next-hop,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi       *string `json:"esi,omitempty"`
	Interface *string `json:"interface,omitempty"`
	// +kubebuilder:validation:Enum=`all-active`;`single-active`
	// +kubebuilder:default:=all-active
	MultiHomingMode *string                                                     `json:"multi-homing-mode,omitempty"`
	Routes          *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes `json:"routes,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=disable
	AdminState *string                                                         `json:"admin-state,omitempty"`
	DfElection *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfElection `json:"df-election,omitempty"`
}

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec struct
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec struct {
	SrlNokiaBgpInstanceId                                   *string                                                 `json:"bgp-instance-id"`
	SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi *[]SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"ethernet-segment"`
}

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus struct
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus struct {
	// ConfigurationDependencyTargetNotFound identifies if the target of the resource object is missing or not
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyTargetFound *TargetFoundStatus `json:"configurationDependencyTargetFound,omitempty"`

	// ConfigurationDependencyValidationStatus identifies the status of the LeafRef Validation of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ConfigurationDependencyValidationStatus *ValidationStatus `json:"configurationDependencyValidationStatus,omitempty"`

	// ConfigurationDependencyValidationDetails defines the validation details of the resource object
	ConfigurationDependencyValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi is the Schema for the SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis API
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec   `json:"spec,omitempty"`
	Status SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList contains a list of SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi",
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

func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
