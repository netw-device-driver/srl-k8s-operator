/*
	Copyright 2021 Wim Henderickx.

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

package controllers

import (
	log "github.com/sirupsen/logrus"
	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func addFinalizer2Resource(lrr *LeafRefResource) error {
	key := types.NamespacedName{
		Namespace: lrr.nameSpace,
		Name:      lrr.leafRefResourceObjectName,
	}

	var o client.Object
	switch lrr.leafRefResourceName {
	case "SrlBfd":
		o = &srlinuxv1alpha1.SrlBfd{}
	case "SrlInterface":
		o = &srlinuxv1alpha1.SrlInterface{}
	case "SrlInterfaceSubinterface":
		o = &srlinuxv1alpha1.SrlInterfaceSubinterface{}
	case "SrlNetworkinstance":
		o = &srlinuxv1alpha1.SrlNetworkinstance{}
	case "SrlNetworkinstanceAggregateroutes":
		o = &srlinuxv1alpha1.SrlNetworkinstanceAggregateroutes{}
	case "SrlNetworkinstanceNexthopgroups":
		o = &srlinuxv1alpha1.SrlNetworkinstanceNexthopgroups{}
	case "SrlNetworkinstanceProtocolsBgp":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgp{}
	case "SrlNetworkinstanceProtocolsBgpevpn":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{}
	case "SrlNetworkinstanceProtocolsBgpvpn":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgpvpn{}
	case "SrlNetworkinstanceProtocolsIsis":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsIsis{}
	case "SrlNetworkinstanceProtocolsLinux":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsLinux{}
	case "SrlNetworkinstanceProtocolsOspf":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsOspf{}
	case "SrlNetworkinstanceStaticroutes":
		o = &srlinuxv1alpha1.SrlNetworkinstanceStaticroutes{}
	case "SrlRoutingpolicyAspathset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyAspathset{}
	case "SrlRoutingpolicyCommunityset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyCommunityset{}
	case "SrlRoutingpolicyPolicy":
		o = &srlinuxv1alpha1.SrlRoutingpolicyPolicy{}
	case "SrlRoutingpolicyPrefixset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyPrefixset{}
	case "SrlSystemMtu":
		o = &srlinuxv1alpha1.SrlSystemMtu{}
	case "SrlSystemName":
		o = &srlinuxv1alpha1.SrlSystemName{}
	case "SrlSystemNetworkinstanceProtocolsBgpvpn":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpn":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}
	case "SrlSystemNtp":
		o = &srlinuxv1alpha1.SrlSystemNtp{}
	case "SrlTunnelinterface":
		o = &srlinuxv1alpha1.SrlTunnelinterface{}
	case "SrlTunnelinterfaceVxlaninterface":
		o = &srlinuxv1alpha1.SrlTunnelinterfaceVxlaninterface{}
	}

	if err := lrr.client.Get(lrr.ctx, key, o); err != nil {
		return err
	}

	f := o.GetFinalizers()
	found := false
	for _, ff := range f {
		if ff == lrr.resourceName+"."+lrr.resourceObjectName+"."+lrr.target {
			found = true
		}
	}
	if !found {
		f = append(f, lrr.resourceName+"."+lrr.resourceObjectName+"."+lrr.target)
		o.SetFinalizers(f)
		if err := lrr.client.Update(lrr.ctx, o); err != nil {
			return err
		}
	}

	log.Infof("Resource Dependent finalizer: %v", o.GetFinalizers())
	return nil
}

func deleteFinalizer2Resource(lrr *LeafRefResource) error {
	key := types.NamespacedName{
		Namespace: lrr.nameSpace,
		Name:      lrr.leafRefResourceObjectName,
	}

	var o client.Object
	switch lrr.leafRefResourceName {
	case "SrlBfd":
		o = &srlinuxv1alpha1.SrlBfd{}
	case "SrlInterface":
		o = &srlinuxv1alpha1.SrlInterface{}
	case "SrlInterfaceSubinterface":
		o = &srlinuxv1alpha1.SrlInterfaceSubinterface{}
	case "SrlNetworkinstance":
		o = &srlinuxv1alpha1.SrlNetworkinstance{}
	case "SrlNetworkinstanceAggregateroutes":
		o = &srlinuxv1alpha1.SrlNetworkinstanceAggregateroutes{}
	case "SrlNetworkinstanceNexthopgroups":
		o = &srlinuxv1alpha1.SrlNetworkinstanceNexthopgroups{}
	case "SrlNetworkinstanceProtocolsBgp":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgp{}
	case "SrlNetworkinstanceProtocolsBgpevpn":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgpevpn{}
	case "SrlNetworkinstanceProtocolsBgpvpn":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsBgpvpn{}
	case "SrlNetworkinstanceProtocolsIsis":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsIsis{}
	case "SrlNetworkinstanceProtocolsLinux":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsLinux{}
	case "SrlNetworkinstanceProtocolsOspf":
		o = &srlinuxv1alpha1.SrlNetworkinstanceProtocolsOspf{}
	case "SrlNetworkinstanceStaticroutes":
		o = &srlinuxv1alpha1.SrlNetworkinstanceStaticroutes{}
	case "SrlRoutingpolicyAspathset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyAspathset{}
	case "SrlRoutingpolicyCommunityset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyCommunityset{}
	case "SrlRoutingpolicyPolicy":
		o = &srlinuxv1alpha1.SrlRoutingpolicyPolicy{}
	case "SrlRoutingpolicyPrefixset":
		o = &srlinuxv1alpha1.SrlRoutingpolicyPrefixset{}
	case "SrlSystemMtu":
		o = &srlinuxv1alpha1.SrlSystemMtu{}
	case "SrlSystemName":
		o = &srlinuxv1alpha1.SrlSystemName{}
	case "SrlSystemNetworkinstanceProtocolsBgpvpn":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsBgpvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpn":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi":
		o = &srlinuxv1alpha1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}
	case "SrlSystemNtp":
		o = &srlinuxv1alpha1.SrlSystemNtp{}
	case "SrlTunnelinterface":
		o = &srlinuxv1alpha1.SrlTunnelinterface{}
	case "SrlTunnelinterfaceVxlaninterface":
		o = &srlinuxv1alpha1.SrlTunnelinterfaceVxlaninterface{}
	}

	if err := lrr.client.Get(lrr.ctx, key, o); err != nil {
		return err
	}

	f := o.GetFinalizers()
	log.Infof("deleteFinalizer2Resource total finalizer: %v", f)
	for _, ff := range f {
		log.Infof("deleteFinalizer2Resource finalizer: %s", ff)
		log.Infof("deleteFinalizer2Resource finalizer resource: %s", lrr.resourceName+"."+lrr.resourceObjectName+"."+lrr.target)
		if ff == lrr.resourceName+"."+lrr.resourceObjectName+"."+lrr.target {
			f = removeString(f, ff)
			o.SetFinalizers(f)
			log.Infof("Resource Dependent finalizer after delete: %v", f)
			if err := lrr.client.Update(lrr.ctx, o); err != nil {
				return err
			}
		}
	}
	return nil
}
