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

var NetworkinstanceStaticroutesIntraResourceleafRef = map[string]*ElementWithLeafRef{}

/*
	var NetworkinstanceStaticroutesIntraResourceleafRef= map[string]*ElementWithLeafRef{
	}
*/

var NetworkinstanceStaticroutesInterResourceleafRef = map[string][]ElementKeyValue{
	"/static-routes/route[prefix=]/next-hop-group": []ElementKeyValue{
		{"network-instance", "", ""},
		{"next-hop-groups", "", ""},
		{"group", "name", "string"},
	},
}

/*
	var NetworkinstanceStaticroutesInterResourceleafRef = map[string]*ElementWithLeafRef{
		"/static-routes/route[prefix=]/next-hop-group" : {
			AbsolutePath2LeafRef: "/network-instance/next-hop-groups",
			RelativePath2LeafRef: "",
		    RelativePath2ObjectWithLeafRef: "",
		    ElementName:         "group",
		    KeyName:             "name",
		},
	}
*/

// SrlNetworkinstanceStaticroutesReconciler reconciles a SrlNetworkinstanceStaticroutes object
type SrlNetworkinstanceStaticroutesReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ctx    context.Context
}

// Instead of passing a zillion arguments to the action of a phase,
// hold them in a context per device
type SrlNetworkinstanceStaticroutesTargetReconcileInfo struct {
	Target map[string]*SrlNetworkinstanceStaticroutesReconcileInfo
}

type SrlNetworkinstanceStaticroutesReconcileInfo struct {
	target              *string
	resource            *string
	o                   *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes
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
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnetworkinstancestaticroutes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnetworkinstancestaticroutes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnetworkinstancestaticroutes/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update
// +kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch

func (r *SrlNetworkinstanceStaticroutesReconciler) publishEvent(request ctrl.Request, event corev1.Event) {
	reqLogger := r.Log.WithValues("SrlNetworkinstanceStaticroutes", request.NamespacedName)
	reqLogger.Info("publishing event", "reason", event.Reason, "message", event.Message)
	err := r.Create(r.Ctx, &event)
	if err != nil {
		reqLogger.Info("failed to record event, ignoring",
			"reason", event.Reason, "message", event.Message, "error", err)
	}
	return
}

func (r *SrlNetworkinstanceStaticroutesReconciler) updateEventHandler(e event.UpdateEvent) bool {
	_, oldOK := e.ObjectOld.(*srlinuxv1alpha1.SrlNetworkinstanceStaticroutes)
	_, newOK := e.ObjectNew.(*srlinuxv1alpha1.SrlNetworkinstanceStaticroutes)
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
func (r *SrlNetworkinstanceStaticroutesReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
	b := ctrl.NewControllerManagedBy(mgr).
		For(&srlinuxv1alpha1.SrlNetworkinstanceStaticroutes{}).
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
// request for reconciliation of SrlNetworkinstanceStaticroutes.
func (r *SrlNetworkinstanceStaticroutesReconciler) NetworkNodeMapFunc(o client.Object) []ctrl.Request {
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
	os := &srlinuxv1alpha1.SrlNetworkinstanceStaticroutesList{}
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
	func (r *SrlNetworkinstanceStaticroutesReconciler) validateIfElementWithLeafRefExists(elements []string, i int, o interface{}, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
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
	func (r *SrlNetworkinstanceStaticroutesReconciler) validateLeafRefExists(elements []string, i int, o interface{}, leafrefValue string, elementWithleafref *ElementWithLeafRef) (interface{}, bool) {
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

func (r *SrlNetworkinstanceStaticroutesReconciler) ValidateLocalLeafRefs(ctx context.Context, o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) (err error) {
	r.Log.Info("Validate Local LeafRef Dependencies ...")

	// marshal data to json
	dd := struct {
		StaticRoutes *srlinuxv1alpha1.NetworkinstanceStaticroutes `json:"static-routes"`
	}{
		StaticRoutes: o.Spec.SrlNetworkinstanceStaticroutes,
	}
	d := make([][]byte, 0)
	dj, err := json.Marshal(dd)
	if err != nil {
		return err
	}
	d = append(d, dj)

	for localLeafRef, leafRefInfo := range NetworkinstanceStaticroutesIntraResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		remoteLeafRefPaths, localLeafRefPaths := r.FindLocalLeafRef(localLeafRef, d, ekvl, leafRefInfo.REkvl)
		r.Log.WithValues("Local LeafRef Path ", localLeafRef, "remoteLeafRefPaths", remoteLeafRefPaths, "localLeafRefPaths", localLeafRefPaths).Info("Local/Remote LeafRef Paths")

		//leafRefInfo.Exists = false
		//leafRefInfo.RemoteLeafRefs = make([]string, 0)
		//leafRefInfo.LocalLeafRefValues = localLeafRefPaths
		leafRefInfo.LocalResolvedLeafRefInfo = make(map[string]*srlinuxv1alpha1.RemoteLeafRefInfo)

		for i, remoteLeafRefPath := range remoteLeafRefPaths {
			//leafRefInfo.Exists = true
			//leafRefInfo.LocalLeafRefValues = append(leafRefInfo.LocalLeafRefValues, localLeafRefPaths[i])
			//leafRefInfo.RemoteLeafRefs = append(leafRefInfo.RemoteLeafRefs, remoteLeafRef)

			rekvl := getHierarchicalElements(remoteLeafRefPath)
			rlvs := r.FindRemoteLeafRef(remoteLeafRefPath, d, rekvl)
			r.Log.WithValues("Remote LeafRef Path ", remoteLeafRefPath, "remote leafref values", rlvs).Info("Remote LeafRef Values")
			found := false
			//leafRefInfo.DependencyCheckSuccess = false

			for _, values := range rlvs {
				if values == rekvl[len(rekvl)-1].KeyValue {
					found = true
					//leafRefInfo.DependencyCheckSuccess = true
					leafRefInfo.LocalResolvedLeafRefInfo[localLeafRefPaths[i]] = &srlinuxv1alpha1.RemoteLeafRefInfo{
						RemoteLeafRef:   stringPtr(remoteLeafRefPath),
						DependencyCheck: srlinuxv1alpha1.DependencyCheckPtr(srlinuxv1alpha1.DependencyCheckSuccess),
					}
					r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("remote Leafref FOUND, all good")
				}
			}
			if !found {
				leafRefInfo.LocalResolvedLeafRefInfo[localLeafRefPaths[i]] = &srlinuxv1alpha1.RemoteLeafRefInfo{
					RemoteLeafRef:   stringPtr(remoteLeafRefPath),
					DependencyCheck: srlinuxv1alpha1.DependencyCheckPtr(srlinuxv1alpha1.DependencyCheckFailed),
				}
				r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("remote Leafref NOT FOUND, missing leaf reference")
			}
		}
		r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("leafref STATUS")
	}
	r.Log.WithValues("NetworkinstanceStaticroutesIntraResourceleafRef", NetworkinstanceStaticroutesIntraResourceleafRef).Info("leafref STATUS All")
	return nil
}

/*
	func (r *SrlNetworkinstanceStaticroutesReconciler) validateLocalLeafRefs(o *srlinuxv1alpha1.NetworkinstanceStaticroutes) (err error) {
		// marshal data to json
		dd := struct {
			StaticRoutes *srlinuxv1alpha1.NetworkinstanceStaticroutes `json:"static-routes"`
		}{
			StaticRoutes: o,
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

		for elementWithleafrefPath, elementWithleafref := range NetworkinstanceStaticroutesIntraResourceleafRef {
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
func (r *SrlNetworkinstanceStaticroutesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("SrlNetworkinstanceStaticroutes", req.NamespacedName)

	r.Log.WithValues("ObjectName", req.NamespacedName).Info("reconciling SrlNetworkinstanceStaticroutes")

	o := &srlinuxv1alpha1.SrlNetworkinstanceStaticroutes{}
	if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get SrlNetworkinstanceStaticroutes config")
		return ctrl.Result{}, err
	}
	o.DeepCopy()

	//r.Log.WithValues("Object", o).Info("Object Info")

	// Add a finalizer to newly created objects.
	if o.DeletionTimestamp.IsZero() && !SrlNetworkinstanceStaticrouteshasFinalizer(o) {
		r.Log.Info(
			"adding finalizer",
			"existingFinalizers", o.Finalizers,
			"newValue", srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer,
		)
		o.Finalizers = append(o.Finalizers,
			srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer)
		err := r.Update(context.TODO(), o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to add finalizer")
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// validate local leaf refs if resource is not in deleting state
	if o.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(o) {
		err := r.ValidateLocalLeafRefs(ctx, o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to validate local leafRef")
		}
		validationSuccess := true
		o.Status.ConfigurationDependencyLocalLeafrefValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails2, 0)
		for localLeafRef, leafRefInfo := range NetworkinstanceStaticroutesIntraResourceleafRef {
			if len(leafRefInfo.LocalResolvedLeafRefInfo) > 0 {
				o.Status.ConfigurationDependencyLocalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails2{
					LocalResolvedLeafRefInfo: make(map[string]*srlinuxv1alpha1.RemoteLeafRefInfo),
				}
				for localLeafRefPath, RemoteLeafRefInfo := range leafRefInfo.LocalResolvedLeafRefInfo {
					if *RemoteLeafRefInfo.DependencyCheck != srlinuxv1alpha1.DependencyCheckSuccess {
						validationSuccess = false
					}
					o.Status.ConfigurationDependencyLocalLeafrefValidationDetails[localLeafRef].LocalResolvedLeafRefInfo[localLeafRefPath] = &srlinuxv1alpha1.RemoteLeafRefInfo{
						RemoteLeafRef:   RemoteLeafRefInfo.RemoteLeafRef,
						DependencyCheck: RemoteLeafRefInfo.DependencyCheck,
					}
				}
			} else {
				o.Status.ConfigurationDependencyLocalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails2{}
			}
		}
		/*
			err := r.validateLocalLeafRefs(o)
			if err != nil {
				return ctrl.Result{}, errors.Wrap(err, "Marshal/Unmarshal errors")
			}
			validationSuccess := true
			o.Status.ConfigurationDependencyValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails, 0)
			for s, elementWithLeafRef := range NetworkinstanceStaticroutesIntraResourceleafRef {
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
			if o.Status.ConfigurationDependencyLocalLeafrefValidationStatus != nil && *o.Status.ConfigurationDependencyLocalLeafrefValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
				r.publishEvent(req, o.NewEvent("Validation success", ""))
			}
			o.Status.ConfigurationDependencyLocalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
		} else {
			// if the validation status did not change we dont have to publish a new event
			if o.Status.ConfigurationDependencyLocalLeafrefValidationStatus != nil && *o.Status.ConfigurationDependencyLocalLeafrefValidationStatus != srlinuxv1alpha1.ValidationStatusFailed {
				r.publishEvent(req, o.NewEvent("Validation failed", "Leaf Ref dependency missing"))
			}
			o.Status.ConfigurationDependencyLocalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
		}

		if err := r.saveSrlNetworkinstanceStaticroutesStatus(ctx, o); err != nil {
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
			if !o.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(o) {
				// remove our finalizer from the list and update it.
				o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer)
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
			if err = r.saveSrlNetworkinstanceStaticroutesStatus(ctx, o); err != nil {
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
		if err = r.saveSrlNetworkinstanceStaticroutesStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}
	}
	r.Log.WithValues("Targets", t).Info("Target Info")

	// find object spec difference and interleafref dependencies if resource is not in deleting state
	var diff bool
	var dp *[]string
	leafRefDependencies := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	if o.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(o) {
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
		leafRefDependencies, localLeafRefPaths, err = r.FindInterLeafRefDependencies(ctx, o)
		if err != nil {
			r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to get leafRef dependencies ")
		}
		r.Log.WithValues("leafRefDependencies", leafRefDependencies, "localLeafRefPaths", localLeafRefPaths).Info("LeafRef Dependencies")
	}

	// initialize the resource parameters
	level := int32(2)
	resource := "srlinux.henderiw.be" + "." + "SrlNetworkinstanceStaticroutes" + "." + strcase.UpperCamelCase(o.Name)
	hkey0 := *o.Spec.SrlNokiaNetworkInstanceName

	dependencies := make([]string, 0)
	dependencies = append(dependencies, fmt.Sprintf("/network-instance[name=%s]", hkey0))
	//dependencies = append(dependencies, fmt.Sprintf("/network-instance[name=%s]", hkey0))

	deletepaths := make([]string, 0)
	deletepaths = append(deletepaths, fmt.Sprintf("/network-instance[name=%s]/static-routes", hkey0))

	// path to be used for this object
	path := fmt.Sprintf("/network-instance[name=%s]", hkey0)

	info := make(map[string]*SrlNetworkinstanceStaticroutesReconcileInfo)
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
		info[target.TargetName] = &SrlNetworkinstanceStaticroutesReconcileInfo{
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
		stateMachine := newSrlNetworkinstanceStaticroutesStateMachine(o, r, &target.TargetName, info[target.TargetName])
		actResult[target.TargetName] = stateMachine.ReconcileState(info[target.TargetName])
		result[target.TargetName], err = actResult[target.TargetName].Result()
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("action %q failed", *initialState))
			return result[target.TargetName], err
		}
	}

	if !o.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(o) {
		deleted := true
		for _, target := range t {
			if result[target.TargetName].RequeueAfter != 0 {
				deleted = false
			}
		}
		if deleted {
			// delete complete
			// remove our finalizer from the list and update it.
			o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer)
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
			if err := r.saveSrlNetworkinstanceStaticroutesStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
		}
		SrlNetworkinstanceStaticrouteslogResult(info[target.TargetName], result[target.TargetName])

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

func SrlNetworkinstanceStaticrouteslogResult(info *SrlNetworkinstanceStaticroutesReconcileInfo, result ctrl.Result) {
	if result.Requeue || result.RequeueAfter != 0 ||
		!utils.StringInList(info.o.Finalizers,
			srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer) {
		info.log.Info("done",
			"requeue", result.Requeue,
			"after", result.RequeueAfter)
	} else {
		info.log.Info("stopping on SrlNetworkinstanceStaticroutes",
			"message", info.o.Status)
	}
}

func (r *SrlNetworkinstanceStaticroutesReconciler) saveSrlNetworkinstanceStaticroutesStatus(ctx context.Context, o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) error {
	t := metav1.Now()
	o.Status.DeepCopy()
	o.Status.LastUpdated = &t

	r.Log.Info("SrlNetworkinstanceStaticroutes",
		"status", o.Status)

	if err := r.Client.Status().Update(ctx, o); err != nil {
		r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to update SrlNetworkinstanceStaticroutes ")
		return err
	}
	return nil
}

func (r *SrlNetworkinstanceStaticroutesReconciler) findLeafRefInTree(x1 interface{}, ekvl []ElementKeyValue, idx int, leafRefValues, localLeafRefPaths []string, lridx int) ([]string, []string) {
	r.Log.WithValues("ekvl", ekvl, "idx", idx, "Data", x1, "leafRefValues", leafRefValues, "localLeafRefPath", localLeafRefPaths).Info("findLeafRefInTree")

	var tlrv []string
	switch x := x1.(type) {
	case map[string]interface{}:
		for k, x2 := range x {
			//r.Log.WithValues("Key", k, "Value", x2, "leafRefValues", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("map[string]interface{}")
			if k == ekvl[idx].Element {
				if idx == len(ekvl)-1 {
					// last element/index in ekv
					if ekvl[idx].KeyName != "" {
						r.Log.WithValues("KeyName", ekvl[idx].KeyName).Info("map[string]interface{} Last Index")
						tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues, localLeafRefPaths, lridx)
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
							localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "=" + x3
							//return leafRefValuesPtr
						case int:
							x4 := strconv.Itoa(int(x3))
							r.Log.WithValues("KeyName", "", "Value", x4, "Type", "int").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, x4)
							localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "=" + x4
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
						tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues, localLeafRefPaths, lridx)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					} else {
						r.Log.WithValues("KeyName", "").Info("map[string]interface{} Not Last Index")
						localLeafRefPaths[lridx] += "/" + ekvl[idx].Element
						idx++
						tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues, localLeafRefPaths, lridx)
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
		leafreforig := localLeafRefPaths[lridx]
		for n, v := range x {
			//r.Log.WithValues("Key", i, "Value", v, "leafRefValues", leafRefValues, "localLeafRefPath", localLeafRefPaths).Info("[]interface{}")
			switch x2 := v.(type) {
			case map[string]interface{}:
				for k3, x3 := range x2 {
					if k3 == ekvl[idx].KeyName {
						if n > 0 {
							localLeafRefPaths = append(localLeafRefPaths, leafreforig)
							lridx++
						}
						if idx == len(ekvl)-1 {
							// return the value
							switch x4 := x3.(type) {
							case string:
								r.Log.WithValues("KeyName", "", "Value", x4, "Type", "string").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x4)
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + x4 + "]"
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")

							case int:
								x5 := strconv.Itoa(int(x4))
								r.Log.WithValues("KeyName", "", "Value", x5, "Type", "int").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x5)
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + x5 + "]"
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
								//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
								//return leafRefValues
							default:
								r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} in []interface{} Last Index")
								//return leafRefValues
							}
						} else {
							r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} in []interface{} Not Last Index")
							switch x4 := x3.(type) {
							case string:
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + x4 + "]"
							}
							i := idx
							i++
							tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, i, leafRefValues, localLeafRefPaths, lridx)
							//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
							if len(tlrv) > len(leafRefValues) {
								leafRefValues = tlrv
							}
							//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
						}
					}
					//i++
				}
			}
		}
		//return leafRefValuesPtr
	case nil:
		r.Log.WithValues("x1", x1).Info("nil")
		//return leafRefValuesPtr
	}
	//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
	return leafRefValues, localLeafRefPaths
}

func (r *SrlNetworkinstanceStaticroutesReconciler) FindRemoteLeafRef(remoteLeafRef string, d [][]byte, rekvl []ElementKeyValue) []string {
	r.Log.WithValues("remoteLeafRef", remoteLeafRef, "rekvl", rekvl).Info("Find Remote LeafRef")
	leafRefValues := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	for _, b := range d {
		var x1 interface{}
		json.Unmarshal(b, &x1)

		localLeafRefPaths = append(localLeafRefPaths, "")
		leafRefValues, localLeafRefPaths = r.findLeafRefInTree(x1, rekvl, 0, leafRefValues, localLeafRefPaths, 0)
		r.Log.WithValues("remoteLeafRef", remoteLeafRef, "Values", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("Find remote LeafRef Values")
	}
	return leafRefValues
}

func (r *SrlNetworkinstanceStaticroutesReconciler) FindLocalLeafRef(localLeafRef string, d [][]byte, ekvl, rekvl []ElementKeyValue) ([]string, []string) {
	r.Log.WithValues("ekvl", ekvl, "rekvl", rekvl).Info("find LeafRef")
	leafRefDependencies := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	for _, b := range d {
		var x1 interface{}
		json.Unmarshal(b, &x1)

		leafRefValues := make([]string, 0)
		localLeafRefPaths = append(localLeafRefPaths, "")
		leafRefValues, localLeafRefPaths = r.findLeafRefInTree(x1, ekvl, 0, leafRefValues, localLeafRefPaths, 0)
		r.Log.WithValues("LocalLeafRef", localLeafRef, "Values", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("find LeafRef Values")
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
	return leafRefDependencies, localLeafRefPaths
}

func (r *SrlNetworkinstanceStaticroutesReconciler) FindInterLeafRefDependencies(ctx context.Context, o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) ([]string, []string, error) {
	r.Log.Info("Find LeafRef Dependencies ...")

	// marshal data to json
	dd := struct {
		StaticRoutes *srlinuxv1alpha1.NetworkinstanceStaticroutes `json:"static-routes"`
	}{
		StaticRoutes: o.Spec.SrlNetworkinstanceStaticroutes,
	}
	d := make([][]byte, 0)
	dj, err := json.Marshal(dd)
	if err != nil {
		return nil, nil, err
	}
	d = append(d, dj)

	leafRefDependencies := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	for localLeafRef, rekvl := range NetworkinstanceProtocolsBgpInterResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		lrd, lrp := r.FindLocalLeafRef(localLeafRef, d, ekvl, rekvl)
		if len(lrd) != 0 {
			leafRefDependencies = append(leafRefDependencies, lrd...)
			localLeafRefPaths = append(localLeafRefPaths, lrp...)
		}
	}
	r.Log.WithValues("LeafRefDependencies", leafRefDependencies).Info("Final LeafRef Dependencies")
	return leafRefDependencies, localLeafRefPaths, nil
}

// FindSpecDiff tries to understand the difference from the latest spec to the newest spec
func (r *SrlNetworkinstanceStaticroutesReconciler) FindSpecDiff(ctx context.Context, o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) (bool, *[]string, error) {
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
func (r *SrlNetworkinstanceStaticroutesReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) ([]*Target, bool, error) {
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

// SrlNetworkinstanceStaticrouteshasFinalizer checks if object has finalizer
func SrlNetworkinstanceStaticrouteshasFinalizer(o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes) bool {
	return StringInList(o.Finalizers, srlinuxv1alpha1.SrlNetworkinstanceStaticroutesFinalizer)
}

func (info *SrlNetworkinstanceStaticroutesReconcileInfo) DeleteCache() error {
	if !info.o.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(info.o) {

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
func (info *SrlNetworkinstanceStaticroutesReconcileInfo) UpdateCache(path string) error {

	// marshal data to json
	dd := struct {
		StaticRoutes *srlinuxv1alpha1.NetworkinstanceStaticroutes `json:"static-routes"`
	}{
		StaticRoutes: info.o.Spec.SrlNetworkinstanceStaticroutes,
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

type SrlNetworkinstanceStaticroutesStateMachine struct {
	Object     *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes
	Reconciler *SrlNetworkinstanceStaticroutesReconciler
	Target     *string
	TargetName *string
	NextState  *srlinuxv1alpha1.ConfigStatus
}

// appendEvent
func (info *SrlNetworkinstanceStaticroutesReconcileInfo) appendEvent(reason, message string) {
	info.events = append(info.events, info.o.NewEvent(reason, message))
}

func newSrlNetworkinstanceStaticroutesStateMachine(o *srlinuxv1alpha1.SrlNetworkinstanceStaticroutes,
	reconciler *SrlNetworkinstanceStaticroutesReconciler, n *string,
	info *SrlNetworkinstanceStaticroutesReconcileInfo) *SrlNetworkinstanceStaticroutesStateMachine {
	currentState := o.Status.Target[*n].ConfigStatus
	r := SrlNetworkinstanceStaticroutesStateMachine{
		Object:     o,
		NextState:  currentState, // Remain in current state by default
		Reconciler: reconciler,
		Target:     info.target,
		TargetName: n,
	}
	return &r
}

type SrlNetworkinstanceStaticroutesstateHandler func(*SrlNetworkinstanceStaticroutesReconcileInfo) actionResult

func (o *SrlNetworkinstanceStaticroutesStateMachine) handlers() map[srlinuxv1alpha1.ConfigStatus]SrlNetworkinstanceStaticroutesstateHandler {
	return map[srlinuxv1alpha1.ConfigStatus]SrlNetworkinstanceStaticroutesstateHandler{
		srlinuxv1alpha1.ConfigStatusNone:             o.handleNone,
		srlinuxv1alpha1.ConfigStatusConfiguring:      o.handleConfiguring,
		srlinuxv1alpha1.ConfigStatusConfigureSuccess: o.handleConfigStatusConfigureSuccess,
		srlinuxv1alpha1.ConfigStatusConfigureFailed:  o.handleConfigStatusConfigureFailed,
		srlinuxv1alpha1.ConfigStatusDeleting:         o.handleDeleting,
	}
}

func (o *SrlNetworkinstanceStaticroutesStateMachine) updateSrlNetworkinstanceStaticroutesStateFrom(initialState *srlinuxv1alpha1.ConfigStatus,
	info *SrlNetworkinstanceStaticroutesReconcileInfo) {
	if o.NextState != initialState {
		info.log.Info("changing configuration state",
			"old", initialState,
			"new", o.NextState)
		o.Object.Status.Target[*o.TargetName].ConfigStatus = o.NextState

		info.appendEvent(fmt.Sprintf("Target: %s, Configuration status old: %s -> new: %s", *o.TargetName, initialState.String(), o.NextState.String()), "")
	}
}

func (o *SrlNetworkinstanceStaticroutesStateMachine) ReconcileState(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
	initialState := o.Object.Status.Target[*o.TargetName].ConfigStatus
	defer o.updateSrlNetworkinstanceStaticroutesStateFrom(initialState, info)

	if o.checkInitiateDelete() {
		// initiate cache delete
		info.log.Info("Initiating SrlNetworkinstanceStaticroutesStateMachine deletion")
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) checkInitiateDelete() bool {
	if !o.Object.DeletionTimestamp.IsZero() && SrlNetworkinstanceStaticrouteshasFinalizer(o.Object) {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) handleNone(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) handleConfiguring(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) handleConfigStatusConfigureSuccess(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) handleConfigStatusConfigureFailed(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) handleDeleting(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) DeleteFailed(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
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

func (o *SrlNetworkinstanceStaticroutesStateMachine) DeleteSuccess(info *SrlNetworkinstanceStaticroutesReconcileInfo) actionResult {
	return actionComplete{}
}
