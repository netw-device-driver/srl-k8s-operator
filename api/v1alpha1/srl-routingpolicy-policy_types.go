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
	// SrlRoutingpolicyPolicyFinalizer is the name of the finalizer added to
	// SrlRoutingpolicyPolicy to block delete operations until the physical node can be
	// deprovisioned.
	SrlRoutingpolicyPolicyFinalizer string = "Policy.srlinux.henderiw.be"
)

// RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct
type RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpAsPath struct
type RoutingpolicyPolicyDefaultActionAcceptBgpAsPath struct {
	Prepend *RoutingpolicyPolicyDefaultActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                   `json:"remove,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpCommunities struct
type RoutingpolicyPolicyDefaultActionAcceptBgpCommunities struct {
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
	Add     *string `json:"add,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference struct
type RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgpOrigin struct
type RoutingpolicyPolicyDefaultActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAcceptBgp struct
type RoutingpolicyPolicyDefaultActionAcceptBgp struct {
	AsPath          *RoutingpolicyPolicyDefaultActionAcceptBgpAsPath          `json:"as-path,omitempty"`
	Communities     *RoutingpolicyPolicyDefaultActionAcceptBgpCommunities     `json:"communities,omitempty"`
	LocalPreference *RoutingpolicyPolicyDefaultActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyDefaultActionAcceptBgpOrigin          `json:"origin,omitempty"`
}

// RoutingpolicyPolicyDefaultActionAccept struct
type RoutingpolicyPolicyDefaultActionAccept struct {
	Bgp *RoutingpolicyPolicyDefaultActionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyDefaultActionNextEntry struct
type RoutingpolicyPolicyDefaultActionNextEntry struct {
}

// RoutingpolicyPolicyDefaultActionNextPolicy struct
type RoutingpolicyPolicyDefaultActionNextPolicy struct {
}

// RoutingpolicyPolicyDefaultActionReject struct
type RoutingpolicyPolicyDefaultActionReject struct {
}

// RoutingpolicyPolicyDefaultAction struct
type RoutingpolicyPolicyDefaultAction struct {
	Accept     *RoutingpolicyPolicyDefaultActionAccept     `json:"accept,omitempty"`
	NextEntry  *RoutingpolicyPolicyDefaultActionNextEntry  `json:"next-entry,omitempty"`
	NextPolicy *RoutingpolicyPolicyDefaultActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyDefaultActionReject     `json:"reject,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend struct
type RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpAsPath struct
type RoutingpolicyPolicyStatementActionAcceptBgpAsPath struct {
	Prepend *RoutingpolicyPolicyStatementActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                     `json:"remove,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct
type RoutingpolicyPolicyStatementActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference struct
type RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct
type RoutingpolicyPolicyStatementActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// RoutingpolicyPolicyStatementActionAcceptBgp struct
type RoutingpolicyPolicyStatementActionAcceptBgp struct {
	Communities     *RoutingpolicyPolicyStatementActionAcceptBgpCommunities     `json:"communities,omitempty"`
	LocalPreference *RoutingpolicyPolicyStatementActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *RoutingpolicyPolicyStatementActionAcceptBgpOrigin          `json:"origin,omitempty"`
	AsPath          *RoutingpolicyPolicyStatementActionAcceptBgpAsPath          `json:"as-path,omitempty"`
}

// RoutingpolicyPolicyStatementActionAccept struct
type RoutingpolicyPolicyStatementActionAccept struct {
	Bgp *RoutingpolicyPolicyStatementActionAcceptBgp `json:"bgp,omitempty"`
}

// RoutingpolicyPolicyStatementActionNextEntry struct
type RoutingpolicyPolicyStatementActionNextEntry struct {
}

// RoutingpolicyPolicyStatementActionNextPolicy struct
type RoutingpolicyPolicyStatementActionNextPolicy struct {
}

// RoutingpolicyPolicyStatementActionReject struct
type RoutingpolicyPolicyStatementActionReject struct {
}

// RoutingpolicyPolicyStatementAction struct
type RoutingpolicyPolicyStatementAction struct {
	NextEntry  *RoutingpolicyPolicyStatementActionNextEntry  `json:"next-entry,omitempty"`
	NextPolicy *RoutingpolicyPolicyStatementActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *RoutingpolicyPolicyStatementActionReject     `json:"reject,omitempty"`
	Accept     *RoutingpolicyPolicyStatementActionAccept     `json:"accept,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgpAsPathLength struct
type RoutingpolicyPolicyStatementMatchBgpAsPathLength struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	// +kubebuilder:default:=eq
	Operator *string `json:"operator,omitempty"`
	// +kubebuilder:default:=false
	Unique *bool `json:"unique,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	Value *uint8 `json:"value"`
}

// RoutingpolicyPolicyStatementMatchBgpEvpn struct
type RoutingpolicyPolicyStatementMatchBgpEvpn struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	RouteType *uint8 `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchBgp struct
type RoutingpolicyPolicyStatementMatchBgp struct {
	AsPathLength *RoutingpolicyPolicyStatementMatchBgpAsPathLength `json:"as-path-length,omitempty"`
	AsPathSet    *string                                           `json:"as-path-set,omitempty"`
	CommunitySet *string                                           `json:"community-set,omitempty"`
	Evpn         *RoutingpolicyPolicyStatementMatchBgpEvpn         `json:"evpn,omitempty"`
}

// RoutingpolicyPolicyStatementMatchIsis struct
type RoutingpolicyPolicyStatementMatchIsis struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Level *uint8 `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=`external`;`internal`
	RouteType *string `json:"route-type,omitempty"`
}

// RoutingpolicyPolicyStatementMatchOspf struct
type RoutingpolicyPolicyStatementMatchOspf struct {
	RouteType *string `json:"route-type,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|[0-9\.]*|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])([\p{N}\p{L}]+)?`
	AreaId *string `json:"area-id,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	InstanceId *uint32 `json:"instance-id,omitempty"`
}

// RoutingpolicyPolicyStatementMatch struct
type RoutingpolicyPolicyStatementMatch struct {
	Protocol  *string                                `json:"protocol,omitempty"`
	Bgp       *RoutingpolicyPolicyStatementMatchBgp  `json:"bgp,omitempty"`
	Family    *string                                `json:"family,omitempty"`
	Isis      *RoutingpolicyPolicyStatementMatchIsis `json:"isis,omitempty"`
	Ospf      *RoutingpolicyPolicyStatementMatchOspf `json:"ospf,omitempty"`
	PrefixSet *string                                `json:"prefix-set,omitempty"`
}

// RoutingpolicyPolicyStatement struct
type RoutingpolicyPolicyStatement struct {
	Match *RoutingpolicyPolicyStatementMatch `json:"match,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	SequenceId *uint32                             `json:"sequence-id"`
	Action     *RoutingpolicyPolicyStatementAction `json:"action,omitempty"`
}

// RoutingpolicyPolicy struct
type RoutingpolicyPolicy struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name          *string                           `json:"name"`
	DefaultAction *RoutingpolicyPolicyDefaultAction `json:"default-action,omitempty"`
	Statement     []*RoutingpolicyPolicyStatement   `json:"statement,omitempty"`
}

// SrlRoutingpolicyPolicySpec struct
type SrlRoutingpolicyPolicySpec struct {
	SrlRoutingpolicyPolicy *[]RoutingpolicyPolicy `json:"policy"`
}

// SrlRoutingpolicyPolicyStatus struct
type SrlRoutingpolicyPolicyStatus struct {
	// ValidationStatus defines the validation status of the resource object
	// +kubebuilder:validation:Enum=Success;Failed
	ValidationStatus *ValidationStatus `json:"validationStatus,omitempty"`

	// ValidationDetails defines the validation details of the resource object
	ValidationDetails map[string]*ValidationDetails `json:"validationDetails,omitempty"`

	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlRoutingpolicyPolicySpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SrlRoutingpolicyPolicy is the Schema for the SrlRoutingpolicyPolicys API
type SrlRoutingpolicyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlRoutingpolicyPolicySpec   `json:"spec,omitempty"`
	Status SrlRoutingpolicyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlRoutingpolicyPolicyList contains a list of SrlRoutingpolicyPolicys
type SrlRoutingpolicyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlRoutingpolicyPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlRoutingpolicyPolicy{}, &SrlRoutingpolicyPolicyList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *SrlRoutingpolicyPolicy) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlRoutingpolicyPolicy",
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

func (o *SrlRoutingpolicyPolicy) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
func (o *SrlRoutingpolicyPolicy) SetConfigStatusDetails(t *string, c *string) {
	o.Status.Target[*t].ConfigStatusDetails = c
}