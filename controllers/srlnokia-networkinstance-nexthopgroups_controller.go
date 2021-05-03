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
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/metal3-io/baremetal-operator/pkg/utils"
	"github.com/stoewer/go-strcase"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	nddv1 "github.com/netw-device-driver/netw-device-controller/api/v1"
	"github.com/netw-device-driver/netwdevpb"

	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkinstanceNexthopgroupsIntraResourceleafRef = map[string]*ElementWithLeafRef{}

var NetworkinstanceNexthopgroupsInterResourceleafRef = map[string]*ElementWithLeafRef{}

// SrlnokiaNetworkinstanceNexthopgroupsReconciler reconciles a SrlnokiaNetworkinstanceNexthopgroups object
type SrlnokiaNetworkinstanceNexthopgroupsReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ctx    context.Context
}

// Instead of passing a zillion arguments to the action of a phase,
// hold them in a context per device
type SrlnokiaNetworkinstanceNexthopgroupsTargetReconcileInfo struct {
	Target map[string]*SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo
}

type SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo struct {
	target            *string
	resource          *string
	o                 *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups
	level             *int32
	dependencies      *[]string
	deletepaths       *[]string
	request           ctrl.Request
	events            []corev1.Event
	errorMessage      *string
	postSaveCallbacks []func()
	ctx               context.Context
	log               logr.Logger
}

// +kubebuilder:rbac:groups=ndd.henderiw.be,resources=networkdevices,verbs=get;list;watch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokianetworkinstancenexthopgroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokianetworkinstancenexthopgroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokianetworkinstancenexthopgroups/finalizers,verbs=update

// SetupWithManager sets up the controller with the Manager.
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
	b := ctrl.NewControllerManagedBy(mgr).
		For(&srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups{}).
		WithOptions(option).
		Watches(
			&source.Kind{Type: &nddv1.NetworkDevice{}},
			handler.EnqueueRequestsFromMapFunc(r.NetworkDeviceMapFunc),
		)

	_, err := b.Build(r)
	if err != nil {
		return errors.Wrap(err, "failed setting up with a controller manager")
	}
	return nil

}

// NetworkDeviceMapFunc is a handler.ToRequestsFunc to be used to enqeue
// request for reconciliation of SrlnokiaNetworkinstanceNexthopgroups.
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) NetworkDeviceMapFunc(o client.Object) []ctrl.Request {
	result := []ctrl.Request{}

	nd, ok := o.(*nddv1.NetworkDevice)
	if !ok {
		panic(fmt.Sprintf("Expected a NodeTopology but got a %T", o))
	}
	r.Log.WithValues(nd.GetName(), nd.GetNamespace()).Info("NetworkDevice MapFunction")

	selectors := []client.ListOption{
		client.InNamespace(nd.Namespace),
		client.MatchingLabels{},
	}
	os := &srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsList{}
	if err := r.Client.List(context.TODO(), os, selectors...); err != nil {
		return result
	}

	for _, o := range os.Items {
		name := client.ObjectKey{
			Namespace: o.GetNamespace(),
			Name:      o.GetName(),
		}
		r.Log.WithValues(o.GetName(), o.GetNamespace()).Info("NetworkDevice MapFunction ReQueue")
		result = append(result, ctrl.Request{NamespacedName: name})
	}

	// delay a bit to ensure the grpc server is started
	time.Sleep(2 * time.Second)

	return result
}

// last -> used to indicate in the validateObject the last element in the leafref. it is provided in the
// validateObject function since you can have a list to walk through. As such we can supply the value direct
// e -> element in the tree
// x1 -> the data object
// elementWithleafref -> the element with leafreaf object on which the values are supplied that match the leafref element
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) validateIfElementWithLeafRefExists(elements []string, i int, o interface{}, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
	//xType := reflect.TypeOf(o)
	//xValue := reflect.ValueOf(o)
	//r.Log.WithValues("xType", xType, "xValue", xValue).Info("validateObject")
	switch x := o.(type) {
	case map[string]interface{}:
		r.Log.Info("validateIfElementWithLeafRefExists map[string]interface{}")
		if v, ok := x[elements[i]]; ok {
			r.Log.WithValues("Element", elements[i], "Object", v, "Last", i == (len(elements)-1)).Info("validateObject found")
			// if last
			if i == (len(elements) - 1) {
				switch val := v.(type) {
				case string:
					found := false
					for _, leafrefValues := range elementWithleafref.Values {
						if string(val) == leafrefValues {
							found = true
						}
					}
					if !found {
						elementWithleafref.Exists = true
						elementWithleafref.Values = append(elementWithleafref.Values, string(val))
					}
				}
				return v, true
			}
			i++
			_, found := r.validateIfElementWithLeafRefExists(elements, i, v, elementWithleafref)
			if !found {
				return nil, false
			}
		} else {
			r.Log.WithValues("Element", elements[i]).Info("validateObject not found")
			return nil, false
		}
	case []interface{}:
		r.Log.Info("validateIfElementWithLeafRefExists []interface{}")
		for _, v1 := range x {
			switch x := v1.(type) {
			case map[string]interface{}:
				if v, ok := x[elements[i]]; ok {
					r.Log.WithValues("Element", elements[i], "Object", v, "Last", i == (len(elements)-1)).Info("validateIfElementWithLeafRefExists found")
					// if last
					if i == (len(elements) - 1) {
						switch val := v.(type) {
						case string:
							found := false
							for _, leafrefValues := range elementWithleafref.Values {
								if string(val) == leafrefValues {
									found = true
								}
							}
							if !found {
								elementWithleafref.Exists = true
								elementWithleafref.Values = append(elementWithleafref.Values, string(val))
							}
						}
						//return v, true
					} else {
						_, found := r.validateIfElementWithLeafRefExists(elements, i, v, elementWithleafref)
						if !found {
							return nil, false
						}
					}
				} else {
					r.Log.WithValues("Element", elements[i]).Info("validateObject not found")
					return nil, false
				}
			}
		}
		if i == (len(elements) - 1) {
			return nil, true
		}
		return nil, false
	}
	r.Log.Info("validateObject not map[string]interface{} or []interface{}")
	return nil, true
}

func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) validateLeafRefExists(elements []string, i int, o interface{}, leafrefValue string, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
	//xType := reflect.TypeOf(o)
	//xValue := reflect.ValueOf(o)
	//r.Log.WithValues("xType", xType, "xValue", xValue).Info("validateObject")
	switch x := o.(type) {
	case map[string]interface{}:
		r.Log.Info("validateLeafRefExists map[string]interface{}")
		if v, ok := x[elements[i]]; ok {
			r.Log.WithValues("Element", elements[i], "LeafRefValue", leafrefValue, "Object", v, "Last", i == (len(elements)-1)).Info("validateLeafRefExists found")
			if i == (len(elements) - 1) {
				switch val := v.(type) {
				case string:
					if leafrefValue == val {
						found := false
						for _, leafrefValues := range elementWithleafref.LeafRefValues {
							if string(val) == leafrefValues {
								found = true
							}
						}
						if !found {
							elementWithleafref.LeafRefValues = append(elementWithleafref.LeafRefValues, val)
						}
						return val, true
					} else {
						found := false
						for _, leafrefValues := range elementWithleafref.LeafRefValues {
							if string(val) == leafrefValues {
								found = true
							}
						}
						if !found {
							elementWithleafref.LeafRefValues = append(elementWithleafref.LeafRefValues, val)
						}
						return val, false
					}
				}
			}
			i++
			_, found := r.validateLeafRefExists(elements, i, v, leafrefValue, elementWithleafref)
			if !found {
				return nil, false
			}
		} else {
			r.Log.WithValues("Element", elements[i]).Info("validateLeafRefExists not found")
			return nil, false
		}
	case []interface{}:
		r.Log.Info("validateLeafRefExists []interface{}")
		f := false
		for _, v1 := range x {
			switch x := v1.(type) {
			case map[string]interface{}:
				if v, ok := x[elements[i]]; ok {
					r.Log.WithValues("Element", elements[i], "LeafRefValue", leafrefValue, "Object", v, "Last", i == (len(elements)-1)).Info("validateLeafRefExists found")
					if i == (len(elements) - 1) {
						switch val := v.(type) {
						case string:
							if leafrefValue == val {
								found := false
								for _, leafrefValues := range elementWithleafref.LeafRefValues {
									if string(val) == leafrefValues {
										found = true
									}
								}
								if !found {
									elementWithleafref.LeafRefValues = append(elementWithleafref.LeafRefValues, val)
								}
								f = true
								//return val, true
							} else {
								found := false
								for _, leafrefValues := range elementWithleafref.LeafRefValues {
									if string(val) == leafrefValues {
										found = true
									}
								}
								if !found {
									elementWithleafref.LeafRefValues = append(elementWithleafref.LeafRefValues, val)
								}
								//return val, false
							}
						}
					} else {
						_, found := r.validateLeafRefExists(elements, i, v, leafrefValue, elementWithleafref)
						if !found {
							return nil, false
						}
					}
					//return v, true
				} else {
					r.Log.WithValues("Element", elements[i]).Info("validateLeafRefExists not found")
					return nil, false
				}
			}
		}
		if i == (len(elements)-1) && f {
			return nil, true
		}
		return nil, false
	}
	r.Log.Info("validateLeafRefExists not map[string]interface{} or []interface{}")
	return nil, true
}

func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) validateLocalLeafRefs(o *srlinuxv1alpha1.NetworkinstanceNexthopgroups) (err error) {
	// marshal data to json
	dd := struct {
		NextHopGroups *srlinuxv1alpha1.NetworkinstanceNexthopgroups `json:"next-hop-groups"`
	}{
		NextHopGroups: o,
	}
	d, err := json.Marshal(dd)
	if err != nil {
		return err
	}
	// unmarshal data to json
	var x interface{}
	err = json.Unmarshal(d, &x)
	if err != nil {
		return err
	}

	for _, elementWithleafref := range NetworkinstanceNexthopgroupsIntraResourceleafRef {
		elementWithleafref.Values = make([]string, 0)
		elementWithleafref.LeafRefValues = make([]string, 0)
		// validate if the element with leafref exist
		elements := strings.Split(elementWithleafref.RelativePath2ObjectWithLeafRef, "/")
		x1 := x
		r.Log.WithValues("X1", x1).Info("Data Input")

		// first element should be initialized with the first resource element
		elements[0] = "bgp"
		_, found := r.validateIfElementWithLeafRefExists(elements, 0, x1, elementWithleafref)
		if !found {
			elementWithleafref.Exists = false
		}

		r.Log.WithValues("leafref values", elementWithleafref.Values).Info("LeafRef Values")
		elementWithleafref.DependencyCheckSuccess = true
		for _, leafReafValue := range elementWithleafref.Values {
			elements := strings.Split(elementWithleafref.RelativePath2LeafRef, "/")
			x1 := x

			// first element should be initialized with the first resource element
			elements[0] = "bgp"

			_, found = r.validateLeafRefExists(elements, 0, x1, leafReafValue, elementWithleafref)
			if !found {
				elementWithleafref.DependencyCheckSuccess = false
				r.Log.WithValues("ElementWithLeafref", elementWithleafref).Info("Leafref NOT FOUND, Object has missing leafs")
			} else {
				r.Log.WithValues("ElementWithLeafref", elementWithleafref).Info("Leafref FOUND, all good")
			}
		}
	}
	return nil
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("SrlnokiaNetworkinstanceNexthopgroups", req.NamespacedName)

	r.Log.WithValues("ObjectName", req.NamespacedName).Info("reconciling SrlnokiaNetworkinstanceNexthopgroups")

	o := &srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups{}
	if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get SrlnokiaNetworkinstanceNexthopgroups config")
		return ctrl.Result{}, err
	}
	o.DeepCopy()

	//r.Log.WithValues("Object", o).Info("Object Info")

	// validate local leaf refs
	err := r.validateLocalLeafRefs(o.Spec.SrlnokiaNetworkinstanceNexthopgroups)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "Marshal/Unmarshal errors")
	}
	validationSuccess := true
	o.Status.ValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails, 0)
	for s, elementWithLeafRef := range NetworkinstanceProtocolsBgpIntraResourceleafRef {
		if elementWithLeafRef.Exists {
			if !elementWithLeafRef.DependencyCheckSuccess {
				validationSuccess = false
			}
			o.Status.ValidationDetails[s] = &srlinuxv1alpha1.ValidationDetails{
				Values:        &elementWithLeafRef.Values,
				LeafRefPath:   &elementWithLeafRef.RelativePath2LeafRef,
				LeafRefValues: &elementWithLeafRef.LeafRefValues,
			}
		} else {
			o.Status.ValidationDetails[s] = &srlinuxv1alpha1.ValidationDetails{
				LeafRefPath: &elementWithLeafRef.RelativePath2LeafRef,
			}
		}
	}

	if validationSuccess {
		o.Status.ValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
	} else {
		o.Status.ValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
	}

	if err := r.saveSrlnokiaNetworkinstanceNexthopgroupsStatus(ctx, o); err != nil {
		return ctrl.Result{}, errors.Wrap(err,
			fmt.Sprintf("failed to save status"))
	}

	if !validationSuccess {
		return ctrl.Result{Requeue: true, RequeueAfter: validationErrorRetyrDelay}, nil
	}

	// find object delta
	LastUsedSpec := o.Status.UsedSpec
	if LastUsedSpec != nil {
		r.Log.WithValues("LastUsedSpec", LastUsedSpec).Info("Last used Spec Info")
	}
	delta, err := r.FindSpecDelta(ctx, o)
	r.Log.WithValues("Spec Detla", *delta).Info("Find Spec Delta")
	o.Status.UsedSpec = &o.Spec

	// Add a finalizer to newly created objects.
	if o.DeletionTimestamp.IsZero() && !SrlnokiaNetworkinstanceNexthopgroupshasFinalizer(o) {
		r.Log.Info(
			"adding finalizer",
			"existingFinalizers", o.Finalizers,
			"newValue", srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsFinalizer,
		)
		o.Finalizers = append(o.Finalizers,
			srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsFinalizer)
		err := r.Update(context.TODO(), o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to add finalizer")
		}
		return ctrl.Result{Requeue: true}, nil
	}

	t, dirty, err := r.FindTarget(ctx, o)
	if err != nil {
		switch err.(type) {
		case *TargetNotFoundError:
			// save resource status since last target got deleted
			if dirty {
				if err = r.saveSrlnokiaNetworkinstanceNexthopgroupsStatus(ctx, o); err != nil {
					return ctrl.Result{}, errors.Wrap(err,
						fmt.Sprintf("failed to save status"))
				}
			}
			// when no target is available requeue to retry after requetimer
			return ctrl.Result{Requeue: true, RequeueAfter: targetNotFoundRetryDelay}, nil
		default:
			return ctrl.Result{}, err
		}
	}
	// save resource status since items got deleted
	if dirty {
		if err = r.saveSrlnokiaNetworkinstanceNexthopgroupsStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}
	}
	r.Log.WithValues("Targets", t).Info("Target Info")

	// initialize the resource parameters
	level := int32(2)
	resource := "srlinux.henderiw.be" + "." + "SrlnokiaNetworkinstanceNexthopgroups" + "." + strcase.UpperCamelCase(o.Name)
	hkey0 := *o.Spec.SrlNokiaNetworkInstanceName

	dependencies := make([]string, 0)
	dependencies = append(dependencies, fmt.Sprintf("/networkinstance[name=%s]", hkey0))

	deletepaths := make([]string, 0)
	deletepaths = append(deletepaths, fmt.Sprintf("/network-instance[name=%s]/next-hop-groups", hkey0))

	// path to be used for this object
	path := fmt.Sprintf("/network-instance[name=%s]", hkey0)

	info := make(map[string]*SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo)
	result := make(map[string]reconcile.Result)
	actResult := make(map[string]actionResult)
	for _, target := range t {
		initialState := new(srlinuxv1alpha1.ConfigStatus)
		if len(o.Status.Target) == 0 {
			o.Status.Target = make(map[string]*srlinuxv1alpha1.TargetStatus)
		}
		if s, ok := o.Status.Target[target.TargetName]; !ok {
			o.Status.Target[target.TargetName] = &srlinuxv1alpha1.TargetStatus{
				ConfigStatus: srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone),
				ErrorCount:   intPtr(0),
			}
			initialState = o.Status.Target[target.TargetName].ConfigStatus
		} else {
			initialState = s.ConfigStatus
		}

		r.Log.Info("configuration status in reconcile",
			"target", target.TargetName,
			"status", initialState)
		info[target.TargetName] = &SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo{
			ctx:          ctx,
			target:       &target.Target,
			log:          r.Log.WithValues("ConfigState", initialState).WithValues("targetName", target.TargetName),
			o:            o,
			request:      req,
			level:        &level,
			resource:     &resource,
			dependencies: &dependencies,
			deletepaths:  &deletepaths,
		}
		if *initialState == srlinuxv1alpha1.ConfigStatusNone {
			// update the cache through GRPC
			err := info[target.TargetName].UpdateCache(path, deletepaths, dependencies)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("grpc update %q failed", *initialState))
				return ctrl.Result{}, err
			}
			// DONT LIKE THIS BELOW BUT REQUE SEEMS TO REQUE IMEDIATELY, NOT SURE WHY
			//time.Sleep(15 * time.Second)
		}

		// activate the state machine

		r.Log.Info("object status",
			"target", target.TargetName,
			"status", o.Status.Target[target.TargetName])
		stateMachine := newSrlnokiaNetworkinstanceNexthopgroupsStateMachine(o, r, &target.TargetName, info[target.TargetName])
		actResult[target.TargetName] = stateMachine.ReconcileState(info[target.TargetName])
		result[target.TargetName], err = actResult[target.TargetName].Result()
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("action %q failed", *initialState))
			return result[target.TargetName], err
		}
	}

	if !o.DeletionTimestamp.IsZero() && SrlnokiaNetworkinstanceNexthopgroupshasFinalizer(o) {
		deleted := true
		for _, target := range t {
			if result[target.TargetName].RequeueAfter != 0 {
				deleted = false
			}
		}
		if deleted {
			// delete complete
			// remove our finalizer from the list and update it.
			o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsFinalizer)
			if err := r.Update(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to remove finalizer"))
			}
			r.Log.Info("cleanup is complete, removed finalizer",
				"remaining", o.Finalizers)
			// Stop reconciliation as the item is deleted
			return ctrl.Result{}, nil
		}
	}
	// Only save status when we're told to, otherwise we
	// introduce an infinite loop reconciling the same object over and
	// over when there is an unrecoverable error (tracked through the
	// error state).

	for _, target := range t {
		dirty := false
		if actResult[target.TargetName].Dirty() {
			dirty = true
		}
		if dirty {
			if err := r.saveSrlnokiaNetworkinstanceNexthopgroupsStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
		}
		SrlnokiaNetworkinstanceNexthopgroupslogResult(info[target.TargetName], result[target.TargetName])

		// requeue for action update and action continue
		if result[target.TargetName].Requeue {
			return ctrl.Result{Requeue: true, RequeueAfter: result[target.TargetName].RequeueAfter}, nil
		}
	}

	return ctrl.Result{}, nil
}

func SrlnokiaNetworkinstanceNexthopgroupslogResult(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo, result ctrl.Result) {
	if result.Requeue || result.RequeueAfter != 0 ||
		!utils.StringInList(info.o.Finalizers,
			srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsFinalizer) {
		info.log.Info("done",
			"requeue", result.Requeue,
			"after", result.RequeueAfter)
	} else {
		info.log.Info("stopping on SrlnokiaNetworkinstanceNexthopgroups",
			"message", info.o.Status)
	}
}

func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) saveSrlnokiaNetworkinstanceNexthopgroupsStatus(ctx context.Context, o *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups) error {
	t := metav1.Now()
	o.Status.DeepCopy()
	o.Status.LastUpdated = &t

	r.Log.Info("SrlnokiaNetworkinstanceNexthopgroups",
		"status", o.Status)

	if err := r.Client.Status().Update(ctx, o); err != nil {
		r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to update SrlnokiaNetworkinstanceNexthopgroups ")
		return err
	}
	return nil
}

// FindTarget finds the SRL target for Object
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) FindSpecDelta(ctx context.Context, o *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups) (*[]string, error) {
	r.Log.Info("Find Spec Delta ...")

	deletepaths := make([]string, 0)

	/*
		if o.Status.UsedSpec != nil {
			if *o.Spec.SrlNokiaNetworkInstanceName != *o.Status.UsedSpec.SrlNokiaNetworkInstanceName {
				deletepaths = append(deletepaths, fmt.Sprintf("/network-instance[name=%s]/protocols/bgp", *o.Status.UsedSpec.SrlNokiaNetworkInstanceName))
			}
		}
	*/

	return &deletepaths, nil
}

// FindTarget finds the SRL target for Object
func (r *SrlnokiaNetworkinstanceNexthopgroupsReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups) ([]*Target, bool, error) {
	r.Log.Info("Find target ...")

	// get the target name the resource should get applied to
	var targetName string
	var dirty bool
	if v, ok := o.Labels["target"]; ok {
		targetName = v
	}
	r.Log.WithValues("target", targetName).Info("Trying to find target name device driver")
	// get network device list, to find the target, based on attached labels
	selectors := []client.ListOption{
		client.MatchingLabels{},
	}
	ndl := &nddv1.NetworkDeviceList{}
	if err := r.List(r.Ctx, ndl, selectors...); err != nil {
		r.Log.Error(err, "Failed to get NetworkDevice List ")
		return nil, dirty, err
	}
	var targets []*Target

	for _, nd := range ndl.Items {
		// check if the network device has a target label and if it matches,
		// append the target to the target list
		//r.Log.WithValues("Network Device", nd).Info("Network Device info")
		if k, ok := nd.Labels["target"]; ok {
			if k == targetName {
				r.Log.WithValues("target", targetName).WithValues("DiscoveryStatus", nd.Status.DiscoveryStatus).Info("Target Label found")
				// the target matches and the network device driver is in ready state
				if nd.Status.DiscoveryStatus != nil && *nd.Status.DiscoveryStatus == nddv1.DiscoveryStatusReady {
					target := &Target{
						TargetName: nd.Name,
						Target:     "nddriver-service-" + nd.Name + ".nddriver-system.svc.cluster.local:" + strconv.Itoa(*nd.Status.GrpcServer.Port),
					}
					// check if the device was already provisioned
					if t, ok := o.Status.Target[nd.Name]; ok {
						// target was already known to the resource and configured, so we exclude
						if t.ConfigStatus != srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess) {
							targets = append(targets, target)
						}
					} else {
						// device was not provisionded so far
						targets = append(targets, target)
					}
				}
			}
		}
		// check if the network device has a target-group label and if it matches,
		// append the target to the target list
		if k, ok := nd.Labels["target-group"]; ok {
			if k == targetName {
				r.Log.WithValues("target", targetName).WithValues("DiscoveryStatus", nd.Status.DiscoveryStatus).Info("Target-group Label found")
				if nd.Status.DiscoveryStatus != nil && *nd.Status.DiscoveryStatus == nddv1.DiscoveryStatusReady {
					target := &Target{
						TargetName: nd.Name,
						Target:     "nddriver-service-" + nd.Name + ".nddriver-system.svc.cluster.local:" + strconv.Itoa(*nd.Status.GrpcServer.Port),
					}
					// check if the device was already provisioned
					if t, ok := o.Status.Target[nd.Name]; ok {
						// target was already known to the resource and configured, so we exclude
						if t.ConfigStatus != srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess) {
							targets = append(targets, target)
						}
					} else {
						// device was not provisionded so far
						targets = append(targets, target)
					}
				}
			}
		}
	}

	// check for deleted items
	for activeTargetName := range o.Status.Target {
		activeTargetDeleted := true
		for _, target := range targets {
			if target.TargetName == activeTargetName {
				activeTargetDeleted = false
			}
		}
		if activeTargetDeleted {
			// delete the status
			delete(o.Status.Target, activeTargetName)
			r.Log.WithValues("target", activeTargetName).Info("Target network device driver got deleted")
			dirty = true
		}
	}

	if len(targets) == 0 {
		// Target not found, return target not found error
		return nil, dirty, &TargetNotFoundError{message: "The Target cannot be found, update label or discovery object"}
	}
	return targets, dirty, nil
}

// SrlnokiaNetworkinstanceNexthopgroupshasFinalizer checks if object has finalizer
func SrlnokiaNetworkinstanceNexthopgroupshasFinalizer(o *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups) bool {
	return StringInList(o.Finalizers, srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroupsFinalizer)
}

func (info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) DeleteCache(deletepaths, dependencies *[]string) error {
	if !info.o.DeletionTimestamp.IsZero() && SrlnokiaNetworkinstanceNexthopgroupshasFinalizer(info.o) {

		// prepare the grpc data
		req := &netwdevpb.CacheUpdateRequest{
			Resource:             *info.resource,
			Level:                *info.level,
			Action:               netwdevpb.CacheUpdateRequest_Delete,
			IndividualActionPath: *deletepaths,
			Dependencies:         *dependencies,
		}

		updateCache(info.ctx, info.target, req)

		info.log.WithValues(
			"Resource", req.Resource).WithValues(
			"Data", req).Info("Published resource config delete data")

	}
	return nil
}

// Update Cache
func (info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) UpdateCache(path string, deletepaths, dependencies []string) error {

	// marshal data to json

	dd := struct {
		NextHopGroups *srlinuxv1alpha1.NetworkinstanceNexthopgroups `json:"next-hop-groups"`
	}{
		NextHopGroups: info.o.Spec.SrlnokiaNetworkinstanceNexthopgroups,
	}
	d := make([][]byte, 0)
	dj, err := json.Marshal(dd)
	if err != nil {
		return err
	}
	d = append(d, dj)

	// update the cache
	req := &netwdevpb.CacheUpdateRequest{
		Resource:             *info.resource,
		Level:                *info.level,
		Action:               netwdevpb.CacheUpdateRequest_Update,
		AggregateActionPath:  path,
		IndividualActionPath: deletepaths,
		ConfigData:           d,
		Dependencies:         dependencies,
	}
	updateCache(info.ctx, info.target, req)

	info.log.WithValues(
		"resource", req.Resource).WithValues(
		"Path", path).WithValues(
		"Data", req).Info("Published resource config update data")

	return nil
}

type SrlnokiaNetworkinstanceNexthopgroupsStateMachine struct {
	Object     *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups
	Reconciler *SrlnokiaNetworkinstanceNexthopgroupsReconciler
	Target     *string
	TargetName *string
	NextState  *srlinuxv1alpha1.ConfigStatus
}

// appendEvent
func (info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) appendEvent(reason, message string) {
	info.events = append(info.events, info.o.NewEvent(reason, message))
}

func newSrlnokiaNetworkinstanceNexthopgroupsStateMachine(o *srlinuxv1alpha1.SrlnokiaNetworkinstanceNexthopgroups,
	reconciler *SrlnokiaNetworkinstanceNexthopgroupsReconciler, n *string,
	info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) *SrlnokiaNetworkinstanceNexthopgroupsStateMachine {
	currentState := o.Status.Target[*n].ConfigStatus
	r := SrlnokiaNetworkinstanceNexthopgroupsStateMachine{
		Object:     o,
		NextState:  currentState, // Remain in current state by default
		Reconciler: reconciler,
		Target:     info.target,
		TargetName: n,
	}
	return &r
}

type SrlnokiaNetworkinstanceNexthopgroupsstateHandler func(*SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handlers() map[srlinuxv1alpha1.ConfigStatus]SrlnokiaNetworkinstanceNexthopgroupsstateHandler {
	return map[srlinuxv1alpha1.ConfigStatus]SrlnokiaNetworkinstanceNexthopgroupsstateHandler{
		srlinuxv1alpha1.ConfigStatusNone:             o.handleNone,
		srlinuxv1alpha1.ConfigStatusConfiguring:      o.handleConfiguring,
		srlinuxv1alpha1.ConfigStatusConfigureSuccess: o.handleConfigStatusConfigureSuccess,
		srlinuxv1alpha1.ConfigStatusConfigureFailed:  o.handleConfigStatusConfigureFailed,
		srlinuxv1alpha1.ConfigStatusDeleting:         o.handleDeleting,
	}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) updateSrlnokiaNetworkinstanceNexthopgroupsStateFrom(initialState *srlinuxv1alpha1.ConfigStatus,
	info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) {
	if o.NextState != initialState {
		info.log.Info("changing configuration state",
			"old", initialState,
			"new", o.NextState)
		o.Object.Status.Target[*o.TargetName].ConfigStatus = o.NextState
	}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) ReconcileState(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	initialState := o.Object.Status.Target[*o.TargetName].ConfigStatus
	defer o.updateSrlnokiaNetworkinstanceNexthopgroupsStateFrom(initialState, info)

	if o.checkInitiateDelete() {
		// initiate cache delete
		info.log.Info("Initiating SrlnokiaNetworkinstanceNexthopgroupsStateMachine deletion")
		info.DeleteCache(info.deletepaths, info.dependencies)
		// DONT LIKE THIS BELOW BUT REQUE SEEMS TO REQUE IMEEDIATELY
		//time.Sleep(15 * time.Second)
	}

	if stateHandler, found := o.handlers()[*initialState]; found {
		return stateHandler(info)
	}

	info.log.Info("No handler found for state", "state", initialState)
	return actionError{fmt.Errorf("No handler found for state \"%s\"", *initialState)}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) checkInitiateDelete() bool {
	if !o.Object.DeletionTimestamp.IsZero() && SrlnokiaNetworkinstanceNexthopgroupshasFinalizer(o.Object) {
		// Delete requested
		switch o.NextState {
		default:
			// new state deleting
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting)
		case srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting),
			srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed),
			srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess):
			// Already in deleting state. Allow state machine to run.
			return false
		}
		return true
	}
	// delete not requested
	return false
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handleNone(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			info.log.Info("object status",
				"target", o.Target,
				"status", o.Object.Status)
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			return actionComplete{}
		}
	}
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
			return actionComplete{}
		}
	} else {
		// update action
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
		o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(cr.Status.String()))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handleConfiguring(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			return actionComplete{}
		}
	}
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
			return actionComplete{}
		}
	} else {
		// update action
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handleConfigStatusConfigureSuccess(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
		}
		return actionUpdate{delay: 10 * time.Second}
	}

	return actionComplete{}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handleConfigStatusConfigureFailed(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			return actionComplete{}
		}
	}
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
			return actionComplete{}
		}
	} else {
		// update action
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) handleDeleting(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)

	if !cr.Exists {
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
		return actionComplete{}
	} else {
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
	}
	return actionUpdate{delay: 10 * time.Second}

}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) DeleteFailed(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)

	if !cr.Exists {
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
		return actionComplete{}
	} else {
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteFailed))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlnokiaNetworkinstanceNexthopgroupsStateMachine) DeleteSuccess(info *SrlnokiaNetworkinstanceNexthopgroupsReconcileInfo) actionResult {
	return actionComplete{}
}
