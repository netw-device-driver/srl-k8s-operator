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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnFinalizer is the name of the finalizer added to
	// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnFinalizer string = "BgpVpn.srlinux.henderiw.be"
)

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]).(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	Rd *string `json:"rd,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`color:[0-1]{2}:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	ExportRt *string `json:"export-rt,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`color:[0-1]{2}:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])`
	ImportRt *string `json:"import-rt,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstance struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstance struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Id                 *uint8                                                                              `json:"id"`
	ExportPolicy       *string                                                                             `json:"export-policy,omitempty"`
	ImportPolicy       *string                                                                             `json:"import-policy,omitempty"`
	RouteDistinguisher *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteDistinguisher `json:"route-distinguisher,omitempty"`
	RouteTarget        *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstanceRouteTarget        `json:"route-target,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn struct {
	BgpInstance []*SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnBgpInstance `json:"bgp-instance,omitempty"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnSpec struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnSpec struct {
	SrlNokiaNetworkInstanceName                           *string                                                `json:"network-instance-name"`
	SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn *SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn `json:"bgp-vpn"`
}

// SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnStatus struct
type SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn is the Schema for the K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpns API
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnSpec   `json:"spec,omitempty"`
	Status SrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnList contains a list of K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpns
type K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn{}, &K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnList{})
}
