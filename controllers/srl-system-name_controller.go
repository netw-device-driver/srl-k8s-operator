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

	"github.com/google/go-cmp/cmp"
	"github.com/metal3-io/baremetal-operator/pkg/utils"
	"github.com/stoewer/go-strcase"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	nddv1 "github.com/netw-device-driver/netw-device-controller/api/v1"
	"github.com/netw-device-driver/netwdevpb"

	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SystemNameIntraResourceleafRef = map[string]*ElementWithLeafRef{}

/*
	var SystemNameIntraResourceleafRef= map[string]*ElementWithLeafRef{
	}
*/

var SystemNameInterResourceleafRef = map[string][]ElementKeyValue{}

/*
	var SystemNameInterResourceleafRef = map[string]*ElementWithLeafRef{
	}
*/

// SrlSystemNameReconciler reconciles a SrlSystemName object
type SrlSystemNameReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ctx    context.Context
}

// Instead of passing a zillion arguments to the action of a phase,
// hold them in a context per device
type SrlSystemNameTargetReconcileInfo struct {
	Target map[string]*SrlSystemNameReconcileInfo
}

type SrlSystemNameReconcileInfo struct {
	target              *string
	resource            *string
	o                   *srlinuxv1alpha1.SrlSystemName
	level               *int32
	dependencies        *[]string
	leafRefDependencies *[]string
	deletepaths         *[]string
	request             ctrl.Request
	events              []corev1.Event
	errorMessage        *string
	postSaveCallbacks   []func()
	ctx                 context.Context
	log                 logr.Logger
}

// +kubebuilder:rbac:groups=ndd.henderiw.be,resources=networknodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlsystemnames,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlsystemnames/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlsystemnames/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update
// +kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch

func (r *SrlSystemNameReconciler) publishEvent(request ctrl.Request, event corev1.Event) {
	reqLogger := r.Log.WithValues("SrlSystemName", request.NamespacedName)
	reqLogger.Info("publishing event", "reason", event.Reason, "message", event.Message)
	err := r.Create(r.Ctx, &event)
	if err != nil {
		reqLogger.Info("failed to record event, ignoring",
			"reason", event.Reason, "message", event.Message, "error", err)
	}
	return
}

func (r *SrlSystemNameReconciler) updateEventHandler(e event.UpdateEvent) bool {
	_, oldOK := e.ObjectOld.(*srlinuxv1alpha1.SrlSystemName)
	_, newOK := e.ObjectNew.(*srlinuxv1alpha1.SrlSystemName)
	if !(oldOK && newOK) {
		// The thing that changed wasn't a host, so we
		// need to assume that we must update. This
		// happens when, for example, an owned Secret
		// changes.
		return true
	}

	//If the update increased the resource Generation then let's process it
	//if e.MetaNew.GetGeneration() != e.MetaOld.GetGeneration() {
	//	return true
	//}

	//Discard updates that did not increase the resource Generation (such as on Status.LastUpdated), except for the finalizers or annotations
	//if reflect.DeepEqual(e.MetaNew.GetFinalizers(), e.MetaOld.GetFinalizers()) && reflect.DeepEqual(e.MetaNew.GetAnnotations(), e.MetaOld.GetAnnotations()) {
	//	return false
	//}

	return true
}

// SetupWithManager sets up the controller with the Manager.
func (r *SrlSystemNameReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
	b := ctrl.NewControllerManagedBy(mgr).
		For(&srlinuxv1alpha1.SrlSystemName{}).
		WithEventFilter(
			predicate.Funcs{
				UpdateFunc: r.updateEventHandler,
			}).
		WithOptions(option).
		Watches(
			&source.Kind{Type: &nddv1.NetworkNode{}},
			handler.EnqueueRequestsFromMapFunc(r.NetworkNodeMapFunc),
		)

	_, err := b.Build(r)
	if err != nil {
		return errors.Wrap(err, "failed setting up with a controller manager")
	}
	return nil

}

// NetworkNodeMapFunc is a handler.ToRequestsFunc to be used to enqeue
// request for reconciliation of SrlSystemName.
func (r *SrlSystemNameReconciler) NetworkNodeMapFunc(o client.Object) []ctrl.Request {
	result := []ctrl.Request{}

	nn, ok := o.(*nddv1.NetworkNode)
	if !ok {
		panic(fmt.Sprintf("Expected a NodeTopology but got a %T", o))
	}
	r.Log.WithValues(nn.GetName(), nn.GetNamespace()).Info("NetworkNode MapFunction")

	selectors := []client.ListOption{
		client.InNamespace(nn.Namespace),
		client.MatchingLabels{},
	}
	os := &srlinuxv1alpha1.SrlSystemNameList{}
	if err := r.Client.List(context.TODO(), os, selectors...); err != nil {
		return result
	}

	for _, o := range os.Items {
		name := client.ObjectKey{
			Namespace: o.GetNamespace(),
			Name:      o.GetName(),
		}
		r.Log.WithValues(o.GetName(), o.GetNamespace()).Info("NetworkNode MapFunction ReQueue")
		result = append(result, ctrl.Request{NamespacedName: name})
	}

	// delay a bit to ensure the grpc server is started
	time.Sleep(2 * time.Second)

	return result
}

/*
	// last -> used to indicate in the validateObject the last element in the leafref. it is provided in the
	// validateObject function since you can have a list to walk through. As such we can supply the value direct
	// e -> element in the tree
	// x1 -> the data object
	// elementWithleafref -> the element with leafreaf object on which the values are supplied that match the leafref element
	func (r *SrlSystemNameReconciler) validateIfElementWithLeafRefExists(elements []string, i int, o interface{}, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
		//xType := reflect.TypeOf(o)
		//xValue := reflect.ValueOf(o)
		//r.Log.WithValues("xType", xType, "xValue", xValue).Info("validateObject")
		switch x := o.(type) {
		case map[string]interface{}:
			//r.Log.Info("validateIfElementWithLeafRefExists map[string]interface{}")
			if v, ok := x[elements[i]]; ok {
				//r.Log.WithValues("Element", elements[i], "Object", v, "Last", i == (len(elements)-1)).Info("validateIfElementWithLeafRefExists found")
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
				//r.Log.WithValues("Element", elements[i]).Info("validateIfElementWithLeafRefExists not found")
				return nil, false
			}
		case []interface{}:
			//r.Log.Info("validateIfElementWithLeafRefExists []interface{}")
			for _, v1 := range x {
				switch x := v1.(type) {
				case map[string]interface{}:
					if v, ok := x[elements[i]]; ok {
						//r.Log.WithValues("Element", elements[i], "Object", v, "Last", i == (len(elements)-1)).Info("validateIfElementWithLeafRefExists found")
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
						//r.Log.WithValues("Element", elements[i]).Info("validateIfElementWithLeafRefExists not found")
						return nil, false
					}
				}
			}
			if i == (len(elements) - 1) {
				return nil, true
			}
			return nil, false
		}
		//r.Log.Info("validateIfElementWithLeafRefExists not map[string]interface{} or []interface{}")
		return nil, true
	}
*/

/*
	func (r *SrlSystemNameReconciler) validateLeafRefExists(elements []string, i int, o interface{}, leafrefValue string, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
		//xType := reflect.TypeOf(o)
		//xValue := reflect.ValueOf(o)
		//r.Log.WithValues("xType", xType, "xValue", xValue).Info("validateObject")
		switch x := o.(type) {
		case map[string]interface{}:
			//r.Log.Info("validateLeafRefExists map[string]interface{}")
			if v, ok := x[elements[i]]; ok {
				//r.Log.WithValues("Element", elements[i], "LeafRefValue", leafrefValue, "Object", v, "Last", i == (len(elements)-1)).Info("validateLeafRefExists found")
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
				//r.Log.WithValues("Element", elements[i]).Info("validateLeafRefExists not found")
				return nil, false
			}
		case []interface{}:
			r.Log.Info("validateLeafRefExists []interface{}")
			f := false
			for _, v1 := range x {
				switch x := v1.(type) {
				case map[string]interface{}:
					if v, ok := x[elements[i]]; ok {
						//r.Log.WithValues("Element", elements[i], "LeafRefValue", leafrefValue, "Object", v, "Last", i == (len(elements)-1)).Info("validateLeafRefExists found")
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
						//r.Log.WithValues("Element", elements[i]).Info("validateLeafRefExists not found")
						return nil, false
					}
				}
			}
			if i == (len(elements) - 1) && f {
				return nil, true
			}
			return nil, false
		}
		//r.Log.Info("validateLeafRefExists not map[string]interface{} or []interface{}")
		return nil, true
	}
*/

func (r *SrlSystemNameReconciler) ValidateLocalLeafRefs(ctx context.Context, o *srlinuxv1alpha1.SrlSystemName) (err error) {
	r.Log.Info("Validate Local LeafRef Dependencies ...")

	// marshal data to json
	dd := struct {
		Name *srlinuxv1alpha1.SystemName `json:"name"`
	}{
		Name: o.Spec.SrlSystemName,
	}
	d := make([][]byte, 0)
	dj, err := json.Marshal(dd)
	if err != nil {
		return err
	}
	d = append(d, dj)

	for localLeafRef, leafRefInfo := range SystemNameIntraResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		lrd := r.FindLeafRef(localLeafRef, d, ekvl, leafRefInfo.REkvl)
		r.Log.WithValues("Local LeafRef Path ", localLeafRef, "leafref values", lrd).Info("Local LeafRef Values")
	}
	return nil
}

/*
	func (r *SrlSystemNameReconciler) validateLocalLeafRefs(o *srlinuxv1alpha1.SystemName) (err error) {
		// marshal data to json
		dd := struct {
			Name *srlinuxv1alpha1.SystemName `json:"name"`
		}{
			Name: o,
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

		for elementWithleafrefPath, elementWithleafref := range SystemNameIntraResourceleafRef {
			elementWithleafref.Values = make([]string, 0)
			elementWithleafref.LeafRefValues = make([]string, 0)
			// validate if the element with leafref exist
			elements := strings.Split(elementWithleafref.RelativePath2ObjectWithLeafRef, "/")
			x1 := x
			//r.Log.WithValues("X1", x1).Info("Data Input")

			// first element should be initialized with the first resource element
			elements[0] = "bgp"
			_, found := r.validateIfElementWithLeafRefExists(elements, 0, x1, elementWithleafref)
			if !found {
				elementWithleafref.Exists = false
			}

			r.Log.WithValues("elementWithleafrefPath", elementWithleafrefPath, "leafref values", elementWithleafref.Values).Info("LeafRef Values")
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
*/
// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SrlSystemNameReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("SrlSystemName", req.NamespacedName)

	r.Log.WithValues("ObjectName", req.NamespacedName).Info("reconciling SrlSystemName")

	o := &srlinuxv1alpha1.SrlSystemName{}
	if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get SrlSystemName config")
		return ctrl.Result{}, err
	}
	o.DeepCopy()

	//r.Log.WithValues("Object", o).Info("Object Info")

	// Add a finalizer to newly created objects.
	if o.DeletionTimestamp.IsZero() && !SrlSystemNamehasFinalizer(o) {
		r.Log.Info(
			"adding finalizer",
			"existingFinalizers", o.Finalizers,
			"newValue", srlinuxv1alpha1.SrlSystemNameFinalizer,
		)
		o.Finalizers = append(o.Finalizers,
			srlinuxv1alpha1.SrlSystemNameFinalizer)
		err := r.Update(context.TODO(), o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to add finalizer")
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// validate local leaf refs if resource is not in deleting state
	if o.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(o) {
		err := r.ValidateLocalLeafRefs(ctx, o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to validate local leafRef")
		}
		validationSuccess := true
		/*
			err := r.validateLocalLeafRefs(o)
			if err != nil {
				return ctrl.Result{}, errors.Wrap(err, "Marshal/Unmarshal errors")
			}
			validationSuccess := true
			o.Status.ConfigurationDependencyValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails, 0)
			for s, elementWithLeafRef := range SystemNameIntraResourceleafRef {
				if elementWithLeafRef.Exists {
					if !elementWithLeafRef.DependencyCheckSuccess {
						validationSuccess = false
					}
					o.Status.ConfigurationDependencyValidationDetails[s] = &srlinuxv1alpha1.ValidationDetails{
						Values:        &elementWithLeafRef.Values,
						LeafRefPath:   &elementWithLeafRef.RelativePath2LeafRef,
						LeafRefValues: &elementWithLeafRef.LeafRefValues,
					}
				} else {
					o.Status.ConfigurationDependencyValidationDetails[s] = &srlinuxv1alpha1.ValidationDetails{
						LeafRefPath: &elementWithLeafRef.RelativePath2LeafRef,
					}
				}
			}
		*/

		//if validationSuccess {
		//	o.Status.ValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
		//} else {
		//	o.Status.ValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
		//}

		if validationSuccess {
			// if the validation status was failed we want to update the event to indicate the success on the transition from failed -> success
			if o.Status.ConfigurationDependencyValidationStatus != nil && *o.Status.ConfigurationDependencyValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
				r.publishEvent(req, o.NewEvent("Validation success", ""))
			}
			o.Status.ConfigurationDependencyValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
		} else {
			// if the validation status did not change we dont have to publish a new event
			if o.Status.ConfigurationDependencyValidationStatus != nil && *o.Status.ConfigurationDependencyValidationStatus != srlinuxv1alpha1.ValidationStatusFailed {
				r.publishEvent(req, o.NewEvent("Validation failed", "Leaf Ref dependency missing"))
			}
			o.Status.ConfigurationDependencyValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
		}

		if err := r.saveSrlSystemNameStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}

		if !validationSuccess {
			return ctrl.Result{Requeue: true, RequeueAfter: validationErrorRetyrDelay}, nil
		}
	}

	t, dirty, err := r.FindTarget(ctx, o)
	if err != nil {
		switch err.(type) {
		case *TargetNotFoundError:
			// delete the resource
			// we can remove the finalizer w/o updating the device since no targets where found
			if !o.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(o) {
				// remove our finalizer from the list and update it.
				o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlSystemNameFinalizer)
				if err := r.Update(ctx, o); err != nil {
					return ctrl.Result{}, errors.Wrap(err,
						fmt.Sprintf("failed to remove finalizer"))
				}
				r.Log.Info("cleanup is complete, removed finalizer",
					"remaining", o.Finalizers)
				// Stop reconciliation as the resource is deleted
				return ctrl.Result{}, nil
			}
			// check status transitions in order to check if the status need to be updated and/or events need to be initialized
			if o.Status.ConfigurationDependencyTargetFound == nil {
				// publish event when the status was not yet initialized
				o.Status.ConfigurationDependencyTargetFound = srlinuxv1alpha1.TargetFoundStatusPtr(srlinuxv1alpha1.TargetFoundStatusFailed)
				r.publishEvent(req, o.NewEvent("Target not found", "No valid target defined to apply the resource upon"))
			} else {
				// only publish events on status transition
				if *o.Status.ConfigurationDependencyTargetFound != srlinuxv1alpha1.TargetFoundStatus(srlinuxv1alpha1.TargetFoundStatusFailed) {
					r.publishEvent(req, o.NewEvent("Target not found", "No valid target defined to apply the resource upon"))
				}
			}
			o.Status.ConfigurationDependencyTargetFound = srlinuxv1alpha1.TargetFoundStatusPtr(srlinuxv1alpha1.TargetFoundStatusFailed)
			// save resource status since last target got deleted
			if err = r.saveSrlSystemNameStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
			// when no target is available requeue to retry after requetimer
			return ctrl.Result{Requeue: true, RequeueAfter: targetNotFoundRetryDelay}, nil
		default:
			return ctrl.Result{}, err
		}
	}
	if o.Status.ConfigurationDependencyTargetFound == nil {
		// target status does not exist
		dirty = true
		o.Status.ConfigurationDependencyTargetFound = srlinuxv1alpha1.TargetFoundStatusPtr(srlinuxv1alpha1.TargetFoundStatusSuccess)
		targets := getTargets(t)
		r.publishEvent(req, o.NewEvent("Target found", targets))
	} else {
		if *o.Status.ConfigurationDependencyTargetFound != srlinuxv1alpha1.TargetFoundStatus(srlinuxv1alpha1.TargetFoundStatusSuccess) {
			dirty = true
			o.Status.ConfigurationDependencyTargetFound = srlinuxv1alpha1.TargetFoundStatusPtr(srlinuxv1alpha1.TargetFoundStatusSuccess)
			r.publishEvent(req, o.NewEvent("Target found", ""))
		}
	}
	// save resource status since items got deleted
	if dirty {
		if err = r.saveSrlSystemNameStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}
	}
	r.Log.WithValues("Targets", t).Info("Target Info")

	// find object spec difference and interleafref dependencies if resource is not in deleting state
	var diff bool
	var dp *[]string
	leafRefDependencies := make([]string, 0)
	if o.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(o) {
		diff, dp, err = r.FindSpecDiff(ctx, o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to find spec delta"))
		}
		if diff {
			r.publishEvent(req, o.NewEvent("Spec changed", "update the resource"))
		}
		r.Log.WithValues("Spec is different, update resource", diff, "Spec Delete Paths", *dp).Info("Spec Diff")
		// the diff handling is handled in the state machine later

		// find leafref dependencies
		leafRefDependencies, err = r.FindInterLeafRefDependencies(ctx, o)
		if err != nil {
			r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to get leafRef dependencies ")
		}
		r.Log.WithValues("Dependencies", leafRefDependencies).Info("LeafRef Dependencies")
	}

	// initialize the resource parameters
	level := int32(2)
	resource := "srlinux.henderiw.be" + "." + "SrlSystemName" + "." + strcase.UpperCamelCase(o.Name)

	dependencies := make([]string, 0)

	deletepaths := make([]string, 0)
	deletepaths = append(deletepaths, "/system/name")

	// path to be used for this object
	path := "/system"

	info := make(map[string]*SrlSystemNameReconcileInfo)
	result := make(map[string]reconcile.Result)
	actResult := make(map[string]actionResult)
	for _, target := range t {
		initialState := new(srlinuxv1alpha1.ConfigStatus)
		// the object was not processed on the target if len is 0
		if len(o.Status.Target) == 0 {
			o.Status.Target = make(map[string]*srlinuxv1alpha1.TargetStatus)
		}
		// initialize the target status if the status was not yet initialized
		if s, ok := o.Status.Target[target.TargetName]; !ok {
			o.Status.Target[target.TargetName] = &srlinuxv1alpha1.TargetStatus{
				ConfigStatus: srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone),
				ErrorCount:   intPtr(0),
			}
			initialState = o.Status.Target[target.TargetName].ConfigStatus
		} else {
			if diff {
				// if the resource was initalized and the object spec changed we should reinitialize the status on the device driver
				o.Status.Target[target.TargetName] = &srlinuxv1alpha1.TargetStatus{
					ConfigStatus: srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone),
					ErrorCount:   intPtr(0),
				}
				initialState = o.Status.Target[target.TargetName].ConfigStatus
			} else {
				initialState = s.ConfigStatus
			}
		}

		r.Log.Info("configuration status in reconcile",
			"target", target.TargetName,
			"status", initialState)
		info[target.TargetName] = &SrlSystemNameReconcileInfo{
			ctx:                 ctx,
			target:              &target.Target,
			log:                 r.Log.WithValues("ConfigState", initialState).WithValues("targetName", target.TargetName),
			o:                   o,
			request:             req,
			level:               &level,
			resource:            &resource,
			dependencies:        &dependencies,
			leafRefDependencies: &leafRefDependencies,
			deletepaths:         &deletepaths,
		}
		if *initialState == srlinuxv1alpha1.ConfigStatusNone {
			r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s, Configuration status old: None -> new: Configuring", target.TargetName), "New Resource or Resource Spec changed"))
			// update the cache through GRPC
			err := info[target.TargetName].UpdateCache(path)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("grpc update %q failed", *initialState))
				return ctrl.Result{}, err
			}
			o.Status.UsedSpec = &o.Spec
		}

		// activate the state machine

		r.Log.Info("object status",
			"target", target.TargetName,
			"status", o.Status.Target[target.TargetName])
		stateMachine := newSrlSystemNameStateMachine(o, r, &target.TargetName, info[target.TargetName])
		actResult[target.TargetName] = stateMachine.ReconcileState(info[target.TargetName])
		result[target.TargetName], err = actResult[target.TargetName].Result()
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("action %q failed", *initialState))
			return result[target.TargetName], err
		}
	}

	if !o.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(o) {
		deleted := true
		for _, target := range t {
			if result[target.TargetName].RequeueAfter != 0 {
				deleted = false
			}
		}
		if deleted {
			// delete complete
			// remove our finalizer from the list and update it.
			o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlSystemNameFinalizer)
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
			if err := r.saveSrlSystemNameStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
		}
		SrlSystemNamelogResult(info[target.TargetName], result[target.TargetName])

		// requeue for action update and action continue
		if result[target.TargetName].Requeue {
			return ctrl.Result{Requeue: true, RequeueAfter: result[target.TargetName].RequeueAfter}, nil
		}
	}

	for _, ri := range info {
		for _, e := range ri.events {
			r.publishEvent(req, e)
		}
	}

	return ctrl.Result{}, nil
}

func SrlSystemNamelogResult(info *SrlSystemNameReconcileInfo, result ctrl.Result) {
	if result.Requeue || result.RequeueAfter != 0 ||
		!utils.StringInList(info.o.Finalizers,
			srlinuxv1alpha1.SrlSystemNameFinalizer) {
		info.log.Info("done",
			"requeue", result.Requeue,
			"after", result.RequeueAfter)
	} else {
		info.log.Info("stopping on SrlSystemName",
			"message", info.o.Status)
	}
}

func (r *SrlSystemNameReconciler) saveSrlSystemNameStatus(ctx context.Context, o *srlinuxv1alpha1.SrlSystemName) error {
	t := metav1.Now()
	o.Status.DeepCopy()
	o.Status.LastUpdated = &t

	r.Log.Info("SrlSystemName",
		"status", o.Status)

	if err := r.Client.Status().Update(ctx, o); err != nil {
		r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to update SrlSystemName ")
		return err
	}
	return nil
}

func (r *SrlSystemNameReconciler) findLeafRefInTree(x1 interface{}, ekvl []ElementKeyValue, idx int, leafRefValues []string) []string {
	r.Log.WithValues("ekvl", ekvl, "idx", idx, "Data", x1, "leafRefValues", leafRefValues).Info("findLeafRefInTree")

	var tlrv []string
	switch x := x1.(type) {
	case map[string]interface{}:
		for k, x2 := range x {
			r.Log.WithValues("Key", k, "Value", x2, "leafRefValues", leafRefValues).Info("map[string]interface{}")
			if k == ekvl[idx].Element {
				if idx == len(ekvl)-1 {
					// last element/index in ekv
					if ekvl[idx].KeyName != "" {
						r.Log.WithValues("KeyName", ekvl[idx].KeyName).Info("map[string]interface{} Last Index")
						tlrv = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					} else {
						switch x3 := x2.(type) {
						case string:
							r.Log.WithValues("KeyName", "", "Value", x3, "Type", "string").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, x3)
							//return leafRefValuesPtr
						case int:
							x4 := strconv.Itoa(int(x3))
							r.Log.WithValues("KeyName", "", "Value", x4, "Type", "int").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, x4)
							//return leafRefValuesPtr
						default:
							r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} Last Index")
							//return leafRefValuesPtr
						}
					}
				} else {
					// not last element/index in ekv
					if ekvl[idx].KeyName != "" {
						r.Log.WithValues("KeyName", ekvl[idx].KeyName).Info("map[string]interface{} Not Last Index")
						tlrv = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					} else {
						r.Log.WithValues("KeyName", "").Info("map[string]interface{} Not Last Index")
						idx++
						tlrv = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					}
				}
			}
		}
	case []interface{}:
		for k, v := range x {
			r.Log.WithValues("Key", k, "Value", v, "leafRefValues", leafRefValues).Info("[]interface{}")
			switch x2 := v.(type) {
			case map[string]interface{}:
				for k3, x3 := range x2 {
					if k3 == ekvl[idx].KeyName {
						if idx == len(ekvl)-1 {
							// return the value
							switch x4 := x3.(type) {
							case string:
								r.Log.WithValues("KeyName", "", "Value", x4, "Type", "string").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x4)
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")

							case int:
								x5 := strconv.Itoa(int(x4))
								r.Log.WithValues("KeyName", "", "Value", x5, "Type", "int").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x5)
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
								//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
								//return leafRefValues
							default:
								r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} in []interface{} Last Index")
								//return leafRefValues
							}
						} else {
							i := idx
							i++
							tlrv = r.findLeafRefInTree(x2, ekvl, i, leafRefValues)
							//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
							if len(tlrv) > len(leafRefValues) {
								leafRefValues = tlrv
							}
							//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
						}
					}
				}
			}
		}
		//return leafRefValuesPtr
	case nil:
		r.Log.WithValues("x1", x1).Info("nil")
		//return leafRefValuesPtr
	}
	//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
	return leafRefValues
}

func (r *SrlSystemNameReconciler) FindLeafRef(localLeafRef string, d [][]byte, ekvl, rekvl []ElementKeyValue) []string {
	r.Log.WithValues("ekvl", ekvl, "rekvl", rekvl).Info("find LeafRef")
	leafRefDependencies := make([]string, 0)
	for _, b := range d {
		var x1 interface{}
		json.Unmarshal(b, &x1)

		leafRefValues := make([]string, 0)
		leafRefValues = r.findLeafRefInTree(x1, ekvl, 0, leafRefValues)
		r.Log.WithValues("LocalLeafRef", localLeafRef, "Values", leafRefValues).Info("find LeafRef Values")
		if len(leafRefValues) != 0 {
			//TODO parse value with rekvl
			//TODO append the result of the previous action with leafRefDependencies
		}
		for _, leafRefValue := range leafRefValues {
			split := strings.Split(leafRefValue, ".")
			leafRefDep := ""
			n := 0
			for _, rekv := range rekvl {
				if rekv.KeyName != "" {
					leafRefDep += "/" + rekv.Element + "[" + rekv.KeyName + "=" + split[n] + "]"
					n++
				} else {
					leafRefDep += "/" + rekv.Element
				}
			}
			leafRefDependencies = append(leafRefDependencies, leafRefDep)
		}
	}
	return leafRefDependencies
}

func (r *SrlSystemNameReconciler) FindInterLeafRefDependencies(ctx context.Context, o *srlinuxv1alpha1.SrlSystemName) ([]string, error) {
	r.Log.Info("Find LeafRef Dependencies ...")

	leafRefDependencies := make([]string, 0)

	// marshal data to json
	dd := struct {
		Name *srlinuxv1alpha1.SystemName `json:"name"`
	}{
		Name: o.Spec.SrlSystemName,
	}
	d := make([][]byte, 0)
	dj, err := json.Marshal(dd)
	if err != nil {
		return nil, err
	}
	d = append(d, dj)

	for localLeafRef, rekvl := range SystemNameInterResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		lrd := r.FindLeafRef(localLeafRef, d, ekvl, rekvl)
		if len(lrd) != 0 {
			leafRefDependencies = append(leafRefDependencies, lrd...)
		}
	}
	r.Log.WithValues("LeafRefDependencies", leafRefDependencies).Info("Final LeafRef Dependencies")
	return leafRefDependencies, nil
}

// FindSpecDiff tries to understand the difference from the latest spec to the newest spec
func (r *SrlSystemNameReconciler) FindSpecDiff(ctx context.Context, o *srlinuxv1alpha1.SrlSystemName) (bool, *[]string, error) {
	r.Log.Info("Find Spec Delta ...")

	deletepaths := make([]string, 0)

	if o.Status.UsedSpec != nil {
		r.Log.Info("Used spec",
			"used spec", *o.Status.UsedSpec)
		r.Log.Info("New spec",
			"nnew spec", o.Spec)
		if cmp.Equal(o.Spec, *o.Status.UsedSpec) {
			return false, &deletepaths, nil
		} else {
			return true, &deletepaths, nil
		}
	} else {
		// if no used spec was available it is the first time we go through the diff and hence we return diff == false
		return false, &deletepaths, nil
	}

	/*
		if o.Status.UsedSpec != nil {
			if *o.Spec.SrlNokiaNetworkInstanceName != *o.Status.UsedSpec.SrlNokiaNetworkInstanceName {
				deletepaths = append(deletepaths, fmt.Sprintf("/network-instance[name=%s]/protocols/bgp", *o.Status.UsedSpec.SrlNokiaNetworkInstanceName))
			}
		}
	*/
}

// FindTarget finds the SRL target for Object
func (r *SrlSystemNameReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.SrlSystemName) ([]*Target, bool, error) {
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
	nnl := &nddv1.NetworkNodeList{}
	if err := r.List(r.Ctx, nnl, selectors...); err != nil {
		r.Log.Error(err, "Failed to get NetworkNode List ")
		return nil, dirty, err
	}
	var targets []*Target

	for _, nn := range nnl.Items {
		// check if the network device has a target label and if it matches,
		// append the target to the target list
		//r.Log.WithValues("Network Device", nd).Info("Network Device info")
		if k, ok := nn.Labels["target"]; ok {
			if k == targetName {
				r.Log.WithValues("target", targetName).WithValues("DiscoveryStatus", nn.Status.DiscoveryStatus).Info("Target Label found")
				// the target matches and the network device driver is in ready state
				if nn.Status.DiscoveryStatus != nil && *nn.Status.DiscoveryStatus == nddv1.DiscoveryStatusReady {
					target := &Target{
						TargetName: nn.Name,
						Target:     "nddriver-service-" + nn.Name + ".nddriver-system.svc.cluster.local:" + strconv.Itoa(*nn.Status.GrpcServer.Port),
					}
					// check if the device was already provisioned
					if t, ok := o.Status.Target[nn.Name]; ok {
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
		if k, ok := nn.Labels["target-group"]; ok {
			if k == targetName {
				r.Log.WithValues("target", targetName).WithValues("DiscoveryStatus", nn.Status.DiscoveryStatus).Info("Target-group Label found")
				if nn.Status.DiscoveryStatus != nil && *nn.Status.DiscoveryStatus == nddv1.DiscoveryStatusReady {
					target := &Target{
						TargetName: nn.Name,
						Target:     "nddriver-service-" + nn.Name + ".nddriver-system.svc.cluster.local:" + strconv.Itoa(*nn.Status.GrpcServer.Port),
					}
					// check if the device was already provisioned
					if t, ok := o.Status.Target[nn.Name]; ok {
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

	// check for deleted items and remove the target from the status
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

// SrlSystemNamehasFinalizer checks if object has finalizer
func SrlSystemNamehasFinalizer(o *srlinuxv1alpha1.SrlSystemName) bool {
	return StringInList(o.Finalizers, srlinuxv1alpha1.SrlSystemNameFinalizer)
}

func (info *SrlSystemNameReconcileInfo) DeleteCache() error {
	if !info.o.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(info.o) {

		// prepare the grpc data
		req := &netwdevpb.CacheUpdateRequest{
			Resource:             *info.resource,
			Level:                *info.level,
			Action:               netwdevpb.CacheUpdateRequest_Delete,
			IndividualActionPath: *info.deletepaths,
			Dependencies:         *info.dependencies,
			LeafRefDependencies:  *info.leafRefDependencies,
		}

		updateCache(info.ctx, info.target, req)

		info.log.WithValues(
			"Resource", req.Resource).WithValues(
			"Data", req).Info("Published resource config delete data")

	}
	return nil
}

// Update Cache
func (info *SrlSystemNameReconcileInfo) UpdateCache(path string) error {

	// marshal data to json
	dd := struct {
		Name *srlinuxv1alpha1.SystemName `json:"name"`
	}{
		Name: info.o.Spec.SrlSystemName,
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
		IndividualActionPath: *info.deletepaths,
		ConfigData:           d,
		Dependencies:         *info.dependencies,
		LeafRefDependencies:  *info.leafRefDependencies,
	}
	updateCache(info.ctx, info.target, req)

	info.log.WithValues(
		"resource", req.Resource).WithValues(
		"Path", path).WithValues(
		"Data", req).Info("Published resource config update data")

	return nil
}

type SrlSystemNameStateMachine struct {
	Object     *srlinuxv1alpha1.SrlSystemName
	Reconciler *SrlSystemNameReconciler
	Target     *string
	TargetName *string
	NextState  *srlinuxv1alpha1.ConfigStatus
}

// appendEvent
func (info *SrlSystemNameReconcileInfo) appendEvent(reason, message string) {
	info.events = append(info.events, info.o.NewEvent(reason, message))
}

func newSrlSystemNameStateMachine(o *srlinuxv1alpha1.SrlSystemName,
	reconciler *SrlSystemNameReconciler, n *string,
	info *SrlSystemNameReconcileInfo) *SrlSystemNameStateMachine {
	currentState := o.Status.Target[*n].ConfigStatus
	r := SrlSystemNameStateMachine{
		Object:     o,
		NextState:  currentState, // Remain in current state by default
		Reconciler: reconciler,
		Target:     info.target,
		TargetName: n,
	}
	return &r
}

type SrlSystemNamestateHandler func(*SrlSystemNameReconcileInfo) actionResult

func (o *SrlSystemNameStateMachine) handlers() map[srlinuxv1alpha1.ConfigStatus]SrlSystemNamestateHandler {
	return map[srlinuxv1alpha1.ConfigStatus]SrlSystemNamestateHandler{
		srlinuxv1alpha1.ConfigStatusNone:             o.handleNone,
		srlinuxv1alpha1.ConfigStatusConfiguring:      o.handleConfiguring,
		srlinuxv1alpha1.ConfigStatusConfigureSuccess: o.handleConfigStatusConfigureSuccess,
		srlinuxv1alpha1.ConfigStatusConfigureFailed:  o.handleConfigStatusConfigureFailed,
		srlinuxv1alpha1.ConfigStatusDeleting:         o.handleDeleting,
	}
}

func (o *SrlSystemNameStateMachine) updateSrlSystemNameStateFrom(initialState *srlinuxv1alpha1.ConfigStatus,
	info *SrlSystemNameReconcileInfo) {
	if o.NextState != initialState {
		info.log.Info("changing configuration state",
			"old", initialState,
			"new", o.NextState)
		o.Object.Status.Target[*o.TargetName].ConfigStatus = o.NextState

		info.appendEvent(fmt.Sprintf("Target: %s, Configuration status old: %s -> new: %s", *o.TargetName, initialState.String(), o.NextState.String()), "")
	}
}

func (o *SrlSystemNameStateMachine) ReconcileState(info *SrlSystemNameReconcileInfo) actionResult {
	initialState := o.Object.Status.Target[*o.TargetName].ConfigStatus
	defer o.updateSrlSystemNameStateFrom(initialState, info)

	if o.checkInitiateDelete() {
		// initiate cache delete
		info.log.Info("Initiating SrlSystemNameStateMachine deletion")
		info.DeleteCache()
		// DONT LIKE THIS BELOW BUT REQUE SEEMS TO REQUE IMEEDIATELY
		//time.Sleep(15 * time.Second)
	}

	if stateHandler, found := o.handlers()[*initialState]; found {
		return stateHandler(info)
	}

	info.log.Info("No handler found for state", "state", initialState)
	return actionError{fmt.Errorf("No handler found for state \"%s\"", *initialState)}
}

func (o *SrlSystemNameStateMachine) checkInitiateDelete() bool {
	if !o.Object.DeletionTimestamp.IsZero() && SrlSystemNamehasFinalizer(o.Object) {
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

func (o *SrlSystemNameStateMachine) handleNone(info *SrlSystemNameReconcileInfo) actionResult {
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
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
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

func (o *SrlSystemNameStateMachine) handleConfiguring(info *SrlSystemNameReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
			return actionComplete{}
		}
		if cr.Data.Action == netwdevpb.CacheUpdateRequest_Delete {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
			return actionContinue{}
		}
	} else {
		info.log.Info("Object got removed by the device driver, most likely due to restart of device driver")
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone))
		return actionUpdate{delay: 1 * time.Second}
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

func (o *SrlSystemNameStateMachine) handleConfigStatusConfigureSuccess(info *SrlSystemNameReconcileInfo) actionResult {
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
	if !cr.Exists {
		info.log.Info("Object got removed by the device driver, most likely due to restart of device driver")
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone))
		return actionUpdate{delay: 1 * time.Second}
	}

	return actionComplete{}
}

func (o *SrlSystemNameStateMachine) handleConfigStatusConfigureFailed(info *SrlSystemNameReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess)
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
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

func (o *SrlSystemNameStateMachine) handleDeleting(info *SrlSystemNameReconcileInfo) actionResult {
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

func (o *SrlSystemNameStateMachine) DeleteFailed(info *SrlSystemNameReconcileInfo) actionResult {
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

func (o *SrlSystemNameStateMachine) DeleteSuccess(info *SrlSystemNameReconcileInfo) actionResult {
	return actionComplete{}
}
