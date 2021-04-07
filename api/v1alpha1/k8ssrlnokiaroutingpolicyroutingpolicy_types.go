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
	// SrlNokiaRoutingPolicyRoutingPolicyFinalizer is the name of the finalizer added to
	// SrlNokiaRoutingPolicyRoutingPolicy to block delete operations until the physical node can be
	// deprovisioned.
	SrlNokiaRoutingPolicyRoutingPolicyFinalizer string = "RoutingPolicy.srlinux.henderiw.be"
)

// SrlNokiaRoutingPolicyRoutingPolicyAsPathSet struct
type SrlNokiaRoutingPolicyRoutingPolicyAsPathSet struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=65535
	Expression *string `json:"expression,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyCommunitySet struct
type SrlNokiaRoutingPolicyRoutingPolicyCommunitySet struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])`
	// +kubebuilder:validation:Pattern=`.*:.*`
	// +kubebuilder:validation:Pattern=`([1-9][0-9]{0,9}):([1-9][0-9]{0,9}):([1-9][0-9]{0,9})`
	// +kubebuilder:validation:Pattern=`.*:.*:.*`
	Member *string `json:"member,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPath struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPath struct {
	Prepend *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                                        `json:"remove,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpCommunities struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpLocalPreference struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpLocalPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpOrigin struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgp struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgp struct {
	LocalPreference *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpOrigin          `json:"origin,omitempty"`
	AsPath          *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpAsPath          `json:"as-path,omitempty"`
	Communities     *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgpCommunities     `json:"communities,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAccept struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAccept struct {
	Bgp *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAcceptBgp `json:"bgp,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextEntry struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextEntry struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextPolicy struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextPolicy struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionReject struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionReject struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultAction struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultAction struct {
	NextPolicy *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionReject     `json:"reject,omitempty"`
	Accept     *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionAccept     `json:"accept,omitempty"`
	NextEntry  *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultActionNextEntry  `json:"next-entry,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPathPrepend struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPathPrepend struct {
	AsNumber *string `json:"as-number,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=50
	// +kubebuilder:default:=1
	RepeatN *uint8 `json:"repeat-n,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPath struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPath struct {
	Prepend *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPathPrepend `json:"prepend,omitempty"`
	Remove  *bool                                                                          `json:"remove,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	Replace *uint32 `json:"replace,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpCommunities struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpCommunities struct {
	Add     *string `json:"add,omitempty"`
	Remove  *string `json:"remove,omitempty"`
	Replace *string `json:"replace,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpLocalPreference struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpLocalPreference struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	Set *uint32 `json:"set,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpOrigin struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpOrigin struct {
	// +kubebuilder:validation:Enum=`egp`;`igp`;`incomplete`
	Set *string `json:"set,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgp struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgp struct {
	Communities     *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpCommunities     `json:"communities,omitempty"`
	LocalPreference *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpLocalPreference `json:"local-preference,omitempty"`
	Origin          *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpOrigin          `json:"origin,omitempty"`
	AsPath          *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgpAsPath          `json:"as-path,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAccept struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAccept struct {
	Bgp *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAcceptBgp `json:"bgp,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextEntry struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextEntry struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextPolicy struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextPolicy struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionReject struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionReject struct {
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementAction struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementAction struct {
	Accept     *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionAccept     `json:"accept,omitempty"`
	NextEntry  *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextEntry  `json:"next-entry,omitempty"`
	NextPolicy *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionNextPolicy `json:"next-policy,omitempty"`
	Reject     *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementActionReject     `json:"reject,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpAsPathLength struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpAsPathLength struct {
	// +kubebuilder:validation:Enum=`eq`;`ge`;`le`
	// +kubebuilder:default:=eq
	Operator *string `json:"operator,omitempty"`
	// +kubebuilder:default:=false
	Unique *bool `json:"unique,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	Value *uint8 `json:"value"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpEvpn struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpEvpn struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	RouteType *uint8 `json:"route-type,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgp struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgp struct {
	AsPathLength *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpAsPathLength `json:"as-path-length,omitempty"`
	AsPathSet    *string                                                                `json:"as-path-set,omitempty"`
	CommunitySet *string                                                                `json:"community-set,omitempty"`
	Evpn         *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgpEvpn         `json:"evpn,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchIsis struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchIsis struct {
	// +kubebuilder:validation:Enum=`external`;`internal`
	RouteType *string `json:"route-type,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	Level *uint8 `json:"level,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchOspf struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchOspf struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	// +kubebuilder:validation:Pattern=`[0-9\.]*`
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%!+(BADINDEX))?`
	AreaId *string `json:"area-id,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	InstanceId *uint32 `json:"instance-id,omitempty"`
	RouteType  *string `json:"route-type,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatch struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatch struct {
	Family    *string                                                     `json:"family,omitempty"`
	Isis      *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchIsis `json:"isis,omitempty"`
	Ospf      *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchOspf `json:"ospf,omitempty"`
	PrefixSet *string                                                     `json:"prefix-set,omitempty"`
	Protocol  *string                                                     `json:"protocol,omitempty"`
	Bgp       *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatchBgp  `json:"bgp,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicyStatement struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicyStatement struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4294967295
	SequenceId *uint32                                                  `json:"sequence-id"`
	Action     *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementAction `json:"action,omitempty"`
	Match      *SrlNokiaRoutingPolicyRoutingPolicyPolicyStatementMatch  `json:"match,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPolicy struct
type SrlNokiaRoutingPolicyRoutingPolicyPolicy struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name          *string                                                `json:"name"`
	DefaultAction *SrlNokiaRoutingPolicyRoutingPolicyPolicyDefaultAction `json:"default-action,omitempty"`
	Statement     []*SrlNokiaRoutingPolicyRoutingPolicyPolicyStatement   `json:"statement,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPrefixSetPrefix struct
type SrlNokiaRoutingPolicyRoutingPolicyPrefixSetPrefix struct {
	IpPrefixMaskLengthRange *string `json:"ip-prefix-mask-length-range,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`([0-9]+\.\.[0-9]+)|exact`
	MaskLengthRange *string `json:"mask-length-range,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicyPrefixSet struct
type SrlNokiaRoutingPolicyRoutingPolicyPrefixSet struct {
	Prefix []*SrlNokiaRoutingPolicyRoutingPolicyPrefixSetPrefix `json:"prefix,omitempty"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
}

// SrlNokiaRoutingPolicyRoutingPolicy struct
type SrlNokiaRoutingPolicyRoutingPolicy struct {
	AsPathSet    []*SrlNokiaRoutingPolicyRoutingPolicyAsPathSet    `json:"as-path-set,omitempty"`
	CommunitySet []*SrlNokiaRoutingPolicyRoutingPolicyCommunitySet `json:"community-set,omitempty"`
	Policy       []*SrlNokiaRoutingPolicyRoutingPolicyPolicy       `json:"policy,omitempty"`
	PrefixSet    []*SrlNokiaRoutingPolicyRoutingPolicyPrefixSet    `json:"prefix-set,omitempty"`
}

// SrlNokiaRoutingPolicyRoutingPolicySpec struct
type SrlNokiaRoutingPolicyRoutingPolicySpec struct {
	SrlNokiaRoutingPolicyRoutingPolicy *SrlNokiaRoutingPolicyRoutingPolicy `json:"routing-policy"`
}

// SrlNokiaRoutingPolicyRoutingPolicyStatus struct
type SrlNokiaRoutingPolicyRoutingPolicyStatus struct {
	// Target provides the status of the configuration on the device
	Target map[string]*TargetStatus `json:"targetStatus,omitempty"`

	// UsedSpec provides the spec used for the configuration
	UsedSpec *SrlNokiaRoutingPolicyRoutingPolicySpec `json:"usedSpec,omitempty"`

	// LastUpdated identifies when this status was last observed.
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sSrlNokiaRoutingPolicyRoutingPolicy is the Schema for the K8sSrlNokiaRoutingPolicyRoutingPolicys API
type K8sSrlNokiaRoutingPolicyRoutingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SrlNokiaRoutingPolicyRoutingPolicySpec   `json:"spec,omitempty"`
	Status SrlNokiaRoutingPolicyRoutingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sSrlNokiaRoutingPolicyRoutingPolicyList contains a list of K8sSrlNokiaRoutingPolicyRoutingPolicys
type K8sSrlNokiaRoutingPolicyRoutingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sSrlNokiaRoutingPolicyRoutingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sSrlNokiaRoutingPolicyRoutingPolicy{}, &K8sSrlNokiaRoutingPolicyRoutingPolicyList{})
}

// NewEvent creates a new event associated with the object and ready
// to be published to the kubernetes API.
func (o *K8sSrlNokiaRoutingPolicyRoutingPolicy) NewEvent(reason, message string) corev1.Event {
	t := metav1.Now()
	return corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: reason + "-",
			Namespace:    o.ObjectMeta.Namespace,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:       "SrlNokiaRoutingPolicyRoutingPolicy",
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

func (o *K8sSrlNokiaRoutingPolicyRoutingPolicy) SetConfigStatus(t *string, c *ConfigStatus) {
	o.Status.Target[*t].ConfigStatus = c
}
