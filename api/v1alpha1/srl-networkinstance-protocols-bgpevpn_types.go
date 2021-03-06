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
	// SrlNetworkinstanceProtocolsBgpevpnFinalizer is the name of the finalizer added to
	// SrlNetworkinstanceProtocolsBgpevpn to block delete operations until the physical node can be
	// deprovisioned.
	SrlNetworkinstanceProtocolsBgpevpnFinalizer string = "BgpEvpn.srlinux.henderiw.be"
)

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	OriginatingIp *string `json:"originating-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp struct {
	// +kubebuilder:default:=true
	Advertise *bool `json:"advertise,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable struct {
	InclusiveMcast *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableInclusiveMcast `json:"inclusive-mcast,omitempty"`
	MacIp          *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTableMacIp          `json:"mac-ip,omitempty"`
	// +kubebuilder:default:=use-system-ipv4-address
	NextHop *string `json:"next-hop,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp struct {
	// +kubebuilder:default:=false
	AdvertiseGatewayMac *bool `json:"advertise-gateway-mac,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable struct {
	MacIp *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTableMacIp `json:"mac-ip,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes struct
type NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes struct {
	RouteTable  *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesRouteTable  `json:"route-table,omitempty"`
	BridgeTable *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutesBridgeTable `json:"bridge-table,omitempty"`
}

// NetworkinstanceProtocolsBgpevpnBgpInstance struct
type NetworkinstanceProtocolsBgpevpnBgpInstance struct {
	VxlanInterface *string `json:"vxlan-interface,omitempty"`
	Id             *string `json:"id"`
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:=enable
	AdminState *string `json:"admin-state,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	// +kubebuilder:default:=0
	DefaultAdminTag *uint32 `json:"default-admin-tag,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:default:=1
	Ecmp *uint8 `json:"ecmp,omitempty"`
	// +kubebuilder:validation:Enum=`vxlan`
	// +kubebuilder:default:=vxlan
	EncapsulationType *string `json:"encapsulation-type,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Evi    *uint32                                           `json:"evi"`
	Routes *NetworkinstanceProtocolsBgpevpnBgpInstanceRoutes `json:"routes,omitempty"`
}

// NetworkinstanceProtocolsBgpevpn struct
type NetworkinstanceProtocolsBgpevpn struct {
	BgpInstance []*NetworkinstanceProtocolsBgpevpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlNetworkinstanceProtocolsBgpevpnSpec struct
type SrlNetworkinstanceProtocolsBgpevpnSpec struct {
	SrlNokiaNetworkInstanceName        *string                          `json:"network-instance-name"`
	SrlNetworkinstanceProtocolsBgpevpn *NetworkinstanceProtocolsBgpevpn `json:"bgp-evpn"`
}

// SrlNetworkinstanceProtocolsBgpevpnStatus struct
type SrlNetworkinstanceProtocolsBgpevpnStatus struct {
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
	UsedSpec *SrlNetworkinstanceProtocolsBgpevpnSpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlNetworkinstanceProtocolsBgpevpn is the Schema for the SrlNetworkinstanceProtocolsBgpevpns API
type SrlNetworkinstanceProtocolsBgpevpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNetworkinstanceProtocolsBgpevpnSpec   `json:"spec,omitempty"`
	Status SrlNetworkinstanceProtocolsBgpevpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsBgpevpnList contains a list of SrlNetworkinstanceProtocolsBgpevpns
type SrlNetworkinstanceProtocolsBgpevpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsBgpevpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsBgpevpn{}, &SrlNetworkinstanceProtocolsBgpevpnList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlNetworkinstanceProtocolsBgpevpn) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNetworkinstanceProtocolsBgpevpn",
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

func (o *SrlNetworkinstanceProtocolsBgpevpn) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlNetworkinstanceProtocolsBgpevpn) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}
