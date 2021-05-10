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
	// SrlNetworkinstanceNexthopgroupsFinalizer is the name of the finalizer added to
	// SrlNetworkinstanceNexthopgroups to block delete operations until the physical node can be
	// deprovisioned.
	SrlNetworkinstanceNexthopgroupsFinalizer string = "NextHopGroups.srlinux.henderiw.be"
)

// NetworkinstanceNexthopgroupsGroupBlackhole struct
type NetworkinstanceNexthopgroupsGroupBlackhole struct {
	// +kubebuilder:default:=false
	GenerateIcmp *bool `json:"generate-icmp,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd struct
type NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	LocalAddress *string `json:"local-address"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16384
	LocalDiscriminator *uint32 `json:"local-discriminator,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=16384
	RemoteDiscriminator *uint32 `json:"remote-discriminator,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthopFailureDetection struct
type NetworkinstanceNexthopgroupsGroupNexthopFailureDetection struct {
	EnableBfd *NetworkinstanceNexthopgroupsGroupNexthopFailureDetectionEnableBfd `json:"enable-bfd,omitempty"`
}

// NetworkinstanceNexthopgroupsGroupNexthop struct
type NetworkinstanceNexthopgroupsGroupNexthop struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Index *uint16 `json:"index"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState       *string                                                   `json:"admin-state,omitempty"`
	FailureDetection *NetworkinstanceNexthopgroupsGroupNexthopFailureDetection `json:"failure-detection,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	IpAddress            *string `json:"ip-address,omitempty"`
	PushedMplsLabelStack *string `json:"pushed-mpls-label-stack,omitempty"`
	// +kubebuilder:default:=true
	Resolve *bool `json:"resolve,omitempty"`
}

// NetworkinstanceNexthopgroupsGroup struct
type NetworkinstanceNexthopgroupsGroup struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string                                     `json:"admin-state,omitempty"`
	Blackhole  *NetworkinstanceNexthopgroupsGroupBlackhole `json:"blackhole,omitempty"`
	Nexthop    []*NetworkinstanceNexthopgroupsGroupNexthop `json:"nexthop,omitempty"`
}

// NetworkinstanceNexthopgroups struct
type NetworkinstanceNexthopgroups struct {
	Group []*NetworkinstanceNexthopgroupsGroup `json:"group,omitempty"`
}

// SrlNetworkinstanceNexthopgroupsSpec struct
type SrlNetworkinstanceNexthopgroupsSpec struct {
	SrlNokiaNetworkInstanceName     *string                       `json:"network-instance-name"`
	SrlNetworkinstanceNexthopgroups *NetworkinstanceNexthopgroups `json:"next-hop-groups"`
}

// SrlNetworkinstanceNexthopgroupsStatus struct
type SrlNetworkinstanceNexthopgroupsStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNetworkinstanceNexthopgroupsSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlNetworkinstanceNexthopgroups is the Schema for the SrlNetworkinstanceNexthopgroupss API
type SrlNetworkinstanceNexthopgroups struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNetworkinstanceNexthopgroupsSpec   `json:"spec,omitempty"`
	Status SrlNetworkinstanceNexthopgroupsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceNexthopgroupsList contains a list of SrlNetworkinstanceNexthopgroupss
type SrlNetworkinstanceNexthopgroupsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceNexthopgroups `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceNexthopgroups{}, &SrlNetworkinstanceNexthopgroupsList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlNetworkinstanceNexthopgroups) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNetworkinstanceNexthopgroups",
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

func (o *SrlNetworkinstanceNexthopgroups) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlNetworkinstanceNexthopgroups) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
