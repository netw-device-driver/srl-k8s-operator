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
	//"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stoewer/go-strcase"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	//"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	nddv1 "github.com/netw-device-driver/netw-device-controller/api/v1"
	"github.com/netw-device-driver/netwdevpb"

	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RoutingpolicyCommunitysetInternalResourceleafRef = map[string]*ElementWithLeafRef{}

var RoutingpolicyCommunitysetExternalResourceleafRef = map[string]*ElementWithLeafRef{}

// SrlRoutingpolicyCommunitysetReconciler reconciles a SrlRoutingpolicyCommunityset object
type SrlRoutingpolicyCommunitysetReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ctx    context.Context
}

// Instead of passing a zillion arguments to the action of a phase,
// hold them in a context per device
type SrlRoutingpolicyCommunitysetTargetReconcileInfo struct {
	Target map[string]*SrlRoutingpolicyCommunitysetReconcileInfo
}

type SrlRoutingpolicyCommunitysetReconcileInfo struct {
	target              *string
	resource            *string
	o                   *srlinuxv1alpha1.SrlRoutingpolicyCommunityset
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
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlroutingpolicycommunitysets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlroutingpolicycommunitysets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlroutingpolicycommunitysets/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update
// +kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;update

func (r *SrlRoutingpolicyCommunitysetReconciler) publishEvent(request ctrl.Request, event corev1.Event) {
	reqLogger := r.Log.WithValues("SrlRoutingpolicyCommunityset", request.NamespacedName)
	reqLogger.Info("publishing event", "reason", event.Reason, "message", event.Message)
	err := r.Create(r.Ctx, &event)
	if err != nil {
		reqLogger.Info("failed to record event, ignoring",
			"reason", event.Reason, "message", event.Message, "error", err)
	}
	return
}

/*
	func (r *SrlRoutingpolicyCommunitysetReconciler) updateEventHandler(e event.UpdateEvent) bool {
		_, oldOK := e.ObjectOld.(*srlinuxv1alpha1.SrlRoutingpolicyCommunityset)
		_, newOK := e.ObjectNew.(*srlinuxv1alpha1.SrlRoutingpolicyCommunityset)
		if !(oldOK && newOK) {
			// The thing that changed wasn't a host, so we
			// need to assume that we must update. This
			// happens when, for example, an owned Secret
			// changes.
			return true
		}
		return true
	}
*/

// SetupWithManager sets up the controller with the Manager.
func (r *SrlRoutingpolicyCommunitysetReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
	b := ctrl.NewControllerManagedBy(mgr).
		For(&srlinuxv1alpha1.SrlRoutingpolicyCommunityset{}).
		WithEventFilter(IgnoreUpdateWithoutGenerationChangePredicate()).
		//WithEventFilter(
		//	predicate.Funcs{
		//		UpdateFunc: r.updateEventHandler,
		//	}).
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
// request for reconciliation of SrlRoutingpolicyCommunityset.
func (r *SrlRoutingpolicyCommunitysetReconciler) NetworkNodeMapFunc(o client.Object) []ctrl.Request {
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
	os := &srlinuxv1alpha1.SrlRoutingpolicyCommunitysetList{}
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

func (r *SrlRoutingpolicyCommunitysetReconciler) ValidateParentDependency(ctx context.Context, cm *string, dependencies *[]string) (bool, error) {
	var x1 interface{}
	json.Unmarshal([]byte(*cm), &x1)

	parentDependencyFound := true
	for _, dep := range *dependencies {
		r.Log.WithValues("Dependency", dep).Info("ValidateParentDependency")
		ekvl := getHierarchicalElements(dep)
		parentDependencyFound = r.findPathInTree(x1, ekvl, 0)
		if !parentDependencyFound {
			return parentDependencyFound, nil
		}
	}
	return parentDependencyFound, nil
}

func (r *SrlRoutingpolicyCommunitysetReconciler) ValidateExternalLeafRefs(ctx context.Context, o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset, cm *string) (err error) {
	r.Log.Info("Validate External LeafRef Dependencies ...")

	// marshal data to json
	d := make([][]byte, 0)
	for _, obj := range *o.Spec.SrlRoutingpolicyCommunityset {
		o := make([]srlinuxv1alpha1.RoutingpolicyCommunityset, 0)
		o = append(o, obj)
		dd := struct {
			CommunitySet *[]srlinuxv1alpha1.RoutingpolicyCommunityset `json:"community-set"`
		}{
			CommunitySet: &o,
		}
		dj, err := json.Marshal(dd)
		if err != nil {
			return err
		}
		d = append(d, dj)
	}

	c := make([][]byte, 0)
	c = append(d, []byte(*cm))

	for localLeafRef, leafRefInfo := range RoutingpolicyCommunitysetExternalResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		remoteLeafRefPaths, localLeafRefPaths := r.FindLocalLeafRef(localLeafRef, d, ekvl, leafRefInfo.REkvl)
		//r.Log.WithValues("Local LeafRef Path ", localLeafRef, "remoteLeafRefPaths", remoteLeafRefPaths, "localLeafRefPaths", localLeafRefPaths).Info("External Local/Remote LeafRef Paths")

		leafRefInfo.LocalResolvedLeafRefInfo = make(map[string]*srlinuxv1alpha1.RemoteLeafRefInfo)
		for i, remoteLeafRefPath := range remoteLeafRefPaths {
			rekvl := getHierarchicalElements(remoteLeafRefPath)
			rlvs := r.FindRemoteLeafRef(remoteLeafRefPath, c, rekvl)
			//r.Log.WithValues("Remote LeafRef Path ", remoteLeafRefPath, "remote leafref values", rlvs).Info("External Remote LeafRef Values")
			found := false

			for _, values := range rlvs {
				if values == rekvl[len(rekvl)-1].KeyValue {
					found = true
					//leafRefInfo.DependencyCheckSuccess = true
					leafRefInfo.LocalResolvedLeafRefInfo[localLeafRefPaths[i]] = &srlinuxv1alpha1.RemoteLeafRefInfo{
						RemoteLeafRef:   stringPtr(remoteLeafRefPath),
						DependencyCheck: srlinuxv1alpha1.DependencyCheckPtr(srlinuxv1alpha1.DependencyCheckSuccess),
					}
					//r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("External Remote Leafref FOUND, all good")
				}
			}
			if !found {
				leafRefInfo.LocalResolvedLeafRefInfo[localLeafRefPaths[i]] = &srlinuxv1alpha1.RemoteLeafRefInfo{
					RemoteLeafRef:   stringPtr(remoteLeafRefPath),
					DependencyCheck: srlinuxv1alpha1.DependencyCheckPtr(srlinuxv1alpha1.DependencyCheckFailed),
				}
				r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("External Remote Leafref NOT FOUND, missing leaf reference")
			}
		}
		//r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("External leafref STATUS")
	}
	r.Log.WithValues("RoutingpolicyCommunitysetExternalResourceleafRef", RoutingpolicyCommunitysetExternalResourceleafRef).Info("External leafref STATUS All")

	return nil
}

func (r *SrlRoutingpolicyCommunitysetReconciler) ValidateLocalLeafRefs(ctx context.Context, o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) (err error) {
	r.Log.Info("Validate Local LeafRef Dependencies ...")

	// marshal data to json
	d := make([][]byte, 0)
	for _, obj := range *o.Spec.SrlRoutingpolicyCommunityset {
		o := make([]srlinuxv1alpha1.RoutingpolicyCommunityset, 0)
		o = append(o, obj)
		dd := struct {
			CommunitySet *[]srlinuxv1alpha1.RoutingpolicyCommunityset `json:"community-set"`
		}{
			CommunitySet: &o,
		}
		dj, err := json.Marshal(dd)
		if err != nil {
			return err
		}
		d = append(d, dj)
	}

	for localLeafRef, leafRefInfo := range RoutingpolicyCommunitysetInternalResourceleafRef {
		// get the ekvl for the local leafref
		ekvl := getHierarchicalElements(localLeafRef)

		// check if the leafref is configured in the resource
		// if not we dont have a leafref dependency in this resource
		remoteLeafRefPaths, localLeafRefPaths := r.FindLocalLeafRef(localLeafRef, d, ekvl, leafRefInfo.REkvl)
		//r.Log.WithValues("Local LeafRef Path ", localLeafRef, "remoteLeafRefPaths", remoteLeafRefPaths, "localLeafRefPaths", localLeafRefPaths).Info("Local/Remote LeafRef Paths")

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
					//r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("remote Leafref FOUND, all good")
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
		//r.Log.WithValues("localLeafRef", localLeafRef, "leafRefInfo", leafRefInfo).Info("leafref STATUS")
	}
	r.Log.WithValues("RoutingpolicyCommunitysetInternalResourceleafRef", RoutingpolicyCommunitysetInternalResourceleafRef).Info("leafref STATUS All")
	return nil
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SrlRoutingpolicyCommunitysetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("SrlRoutingpolicyCommunityset", req.NamespacedName)

	r.Log.WithValues("ObjectName", req.NamespacedName).Info("reconciling SrlRoutingpolicyCommunityset")

	o := &srlinuxv1alpha1.SrlRoutingpolicyCommunityset{}
	if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get SrlRoutingpolicyCommunityset config")
		return ctrl.Result{}, err
	}
	o.DeepCopy()

	//r.Log.WithValues("Object", o).Info("Object Info")

	// Add a finalizer to newly created objects.
	if o.DeletionTimestamp.IsZero() && !SrlRoutingpolicyCommunitysethasFinalizer(o) {
		r.Log.Info(
			"adding finalizer",
			"existingFinalizers", o.Finalizers,
			"newValue", srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer,
		)
		o.Finalizers = append(o.Finalizers,
			srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer)
		err := r.Update(ctx, o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to add finalizer")
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// check other dependencies, if so we cannot delete the object
	if !o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
		if SrlRoutingpolicyCommunitysethasOtherFinalizer(o) {
			// other dependencies block the deletion, requeue the request
			return ctrl.Result{Requeue: true, RequeueAfter: deleteDependencyRetryDelay}, nil
		}

	}

	// validate local leaf refs if resource is not in deleting state
	if o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
		err := r.ValidateLocalLeafRefs(ctx, o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to validate local leafRef")
		}
		validationSuccess := true
		o.Status.ConfigurationDependencyInternalLeafrefValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails, 0)
		for localLeafRef, leafRefInfo := range RoutingpolicyCommunitysetInternalResourceleafRef {
			if len(leafRefInfo.LocalResolvedLeafRefInfo) > 0 {
				o.Status.ConfigurationDependencyInternalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails{
					LocalResolvedLeafRefInfo: make(map[string]*srlinuxv1alpha1.RemoteLeafRefInfo),
				}
				for localLeafRefPath, RemoteLeafRefInfo := range leafRefInfo.LocalResolvedLeafRefInfo {
					if *RemoteLeafRefInfo.DependencyCheck != srlinuxv1alpha1.DependencyCheckSuccess {
						validationSuccess = false
					}
					o.Status.ConfigurationDependencyInternalLeafrefValidationDetails[localLeafRef].LocalResolvedLeafRefInfo[localLeafRefPath] = &srlinuxv1alpha1.RemoteLeafRefInfo{
						RemoteLeafRef:   RemoteLeafRefInfo.RemoteLeafRef,
						DependencyCheck: RemoteLeafRefInfo.DependencyCheck,
					}
				}
			} else {
				o.Status.ConfigurationDependencyInternalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails{}
			}
		}

		if o.Status.ConfigurationDependencyInternalLeafrefValidationStatus == nil {
			if validationSuccess {
				r.publishEvent(req, o.NewEvent("Internal Leafref Validation success", ""))
				o.Status.ConfigurationDependencyInternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
			} else {
				r.publishEvent(req, o.NewEvent("Internal Leafref Validation failed", "Leaf Ref dependency missing"))
				o.Status.ConfigurationDependencyInternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
			}
		} else {
			if validationSuccess {
				// if the validation status was failed we want to update the event to indicate the success on the transition from failed -> success
				if *o.Status.ConfigurationDependencyInternalLeafrefValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
					r.publishEvent(req, o.NewEvent("Internal Leafref Validation success", ""))
				}
				o.Status.ConfigurationDependencyInternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
			} else {
				// if the validation status did not change we dont have to publish a new event
				if *o.Status.ConfigurationDependencyInternalLeafrefValidationStatus != srlinuxv1alpha1.ValidationStatusFailed {
					r.publishEvent(req, o.NewEvent("Internal Leafref Validation failed", "Leaf Ref dependency missing"))
				}
				o.Status.ConfigurationDependencyInternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
			}
		}

		if err := r.saveSrlRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}

		if !validationSuccess {
			return ctrl.Result{Requeue: true, RequeueAfter: internalLeafRefvalidationErrorretryDelay}, nil
		}
	}

	t, dirty, err := r.FindTarget(ctx, o)
	if err != nil {
		switch err.(type) {
		case *TargetNotFoundError:
			// delete the resource
			// we can remove the finalizer w/o updating the device since no targets where found
			if !o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
				// remove the leafref dependency finalizers from remote leafref objects
				for targetName, ts := range o.Status.Target {
					// localLeafRef, leafRefInfo
					for _, leafRefInfo := range ts.ConfigurationDependencyExternalLeafrefValidationDetails {
						// localLeafRefPath, RemoteLeafRefInfo
						for _, RemoteLeafRefInfo := range leafRefInfo.LocalResolvedLeafRefInfo {
							if RemoteLeafRefInfo.RemoteResourceObject != nil {
								r.Log.WithValues("RemoteResourceObject", *RemoteLeafRefInfo.RemoteResourceObject).Info("Remote LeafRef Object")
								if *RemoteLeafRefInfo.RemoteResourceObject != "" {
									// first part of split, split[0] is the resource, 2nd part is the resourceName, split[1]
									split := strings.Split(*RemoteLeafRefInfo.RemoteResourceObject, ".")
									lrr := &LeafRefResource{
										ctx:                       ctx,
										client:                    r.Client,
										nameSpace:                 o.GetNamespace(),
										resourceName:              "CommunitySet",
										resourceObjectName:        strcase.UpperCamelCase(o.GetName()),
										leafRefResourceName:       split[0],
										leafRefResourceObjectName: split[1],
										target:                    targetName,
									}
									deleteFinalizer2Resource(lrr)
								} else {
									r.Log.Info("Remote LeafRef dependency is empty, somethign went wrong")
								}
							}
						}
					}
				}
				// remove our finalizer from the list and update it.
				o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer)
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
			if err = r.saveSrlRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
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
		if err = r.saveSrlRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}
	}
	r.Log.WithValues("Targets", t).Info("Target Info")

	// find object spec difference if resource is not in deleting state
	var diff bool
	var dp *[]string
	leafRefDependencies := make([]string, 0)
	//localLeafRefPaths := make([]string, 0)
	if o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
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
	}

	// initialize the resource parameters
	level := int32(2)
	resource := "srlinux.henderiw.be" + "." + "SrlRoutingpolicyCommunityset" + "." + strcase.UpperCamelCase(o.Name)
	hkeys0 := make([]string, 0)
	for _, n := range *o.Spec.SrlRoutingpolicyCommunityset {
		hkeys0 = append(hkeys0, *n.Name)
	}

	dependencies := make([]string, 0)

	deletepaths := make([]string, 0)
	for _, hkey0 := range hkeys0 {
		deletepaths = append(deletepaths, fmt.Sprintf("/routing-policy/community-set[name=%s]", hkey0))
	}

	// path to be used for this object
	path := "/routing-policy"

	// validate parent dependency and external leafref dependencies if not in deleteing status
	if o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
		for _, target := range t {
			// initialize the status if not yet done
			// the object was not processed on the target if len is 0
			if len(o.Status.Target) == 0 {
				o.Status.Target = make(map[string]*srlinuxv1alpha1.TargetStatus)
			}
			if _, ok := o.Status.Target[target.TargetName]; !ok {
				o.Status.Target[target.TargetName] = &srlinuxv1alpha1.TargetStatus{
					ConfigStatus: srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone),
					ErrorCount:   intPtr(0),
					ConfigurationDependencyParentValidationDetails:          make(map[string]*srlinuxv1alpha1.ValidationDetails, 0),
					ConfigurationDependencyExternalLeafrefValidationDetails: make(map[string]*srlinuxv1alpha1.ValidationDetails, 0),
				}
			}

			// check if shared memory was initialized
			if _, ok := SrlSharedInfo[target.TargetName]; !ok {
				SrlSharedInfo[target.TargetName] = make(map[string]string)
			}
			// add entry in shared memory
			if err := addSharedMemoryEntry(stringPtr(target.TargetName), stringPtr("SrlRoutingpolicyCommunityset"), stringPtr(req.Name), stringSlicePtr(deletepaths)); err != nil {
				return ctrl.Result{}, err
			}
			r.Log.WithValues("target", target.TargetName, "shared memory data", SrlSharedInfo[target.TargetName]).Info("shared memory info")

			// get configmap
			cm, err := r.getConfigMap(ctx, stringPtr(target.TargetName))
			if err != nil {
				return ctrl.Result{}, err
			}
			// update configmap with deletepaths
			/*
				if err := r.addEntryConfigMap(ctx, stringPtr(target.TargetName), stringPtr(req.Name), stringSlicePtr(deletepaths)); err != nil {
					return ctrl.Result{}, err
				}
			*/
			var x1 interface{}
			json.Unmarshal([]byte(*cm), &x1)

			// validate Parent Dependency
			parentDependencyFound, err := r.ValidateParentDependency(ctx, cm, stringSlicePtr(dependencies))
			r.Log.WithValues("Target", target.TargetName, "ParentDependencyFound", parentDependencyFound).Info("Parent Dependency")

			if o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus == nil {
				if parentDependencyFound {
					r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s Parent Dependency Found", target.TargetName), fmt.Sprintf("Dependency %v", dependencies)))
					o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
				} else {
					r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s Parent Dependency Not Found", target.TargetName), fmt.Sprintf("Dependency %v", dependencies)))
					o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
				}
			} else {
				if parentDependencyFound {
					// if the parentDependencyFound status was found we want to update the event to indicate the success on the transition from failed -> success
					if *o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
						r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s Parent Dependency Found", target.TargetName), fmt.Sprintf("Dependency %v", dependencies)))
					}
					o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
				} else {
					// if the validation status did not change we dont have to publish a new event
					if *o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus != srlinuxv1alpha1.ValidationStatusFailed {
						r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s Parent Dependency Not Found", target.TargetName), fmt.Sprintf("Dependency %v", dependencies)))
					}
					o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
				}
			}

			err = r.ValidateExternalLeafRefs(ctx, o, cm)
			if err != nil {
				return ctrl.Result{}, errors.Wrap(err, "failed to validate external leafRef")
			}

			validationSuccess := true
			o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationDetails = make(map[string]*srlinuxv1alpha1.ValidationDetails, 0)
			for localLeafRef, leafRefInfo := range RoutingpolicyCommunitysetExternalResourceleafRef {
				if len(leafRefInfo.LocalResolvedLeafRefInfo) > 0 {
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails{
						LocalResolvedLeafRefInfo: make(map[string]*srlinuxv1alpha1.RemoteLeafRefInfo),
					}
					for localLeafRefPath, RemoteLeafRefInfo := range leafRefInfo.LocalResolvedLeafRefInfo {
						if *RemoteLeafRefInfo.DependencyCheck != srlinuxv1alpha1.DependencyCheckSuccess {
							validationSuccess = false

							o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationDetails[localLeafRef].LocalResolvedLeafRefInfo[localLeafRefPath] = &srlinuxv1alpha1.RemoteLeafRefInfo{
								RemoteLeafRef:   RemoteLeafRefInfo.RemoteLeafRef,
								DependencyCheck: RemoteLeafRefInfo.DependencyCheck,
							}
						} else {
							res, err := getRemoteleafRefResource(stringPtr(target.TargetName), RemoteLeafRefInfo)
							//res, err := r.GetRemoteleafRefResource(ctx, stringPtr(target.TargetName), RemoteLeafRefInfo)
							if err != nil {
								return ctrl.Result{}, errors.Wrap(err,
									fmt.Sprintf("failed to get remote leaf ref resource"))
							}
							r.Log.WithValues("Resource", *res).Info("Remote LeafRef resource")

							if *res != "" {
								// first part of split, split[0] is the resource, 2nd part is the resourceName, split[1]
								split := strings.Split(*res, ".")
								lrr := &LeafRefResource{
									ctx:                       ctx,
									client:                    r.Client,
									nameSpace:                 o.GetNamespace(),
									resourceName:              "CommunitySet",
									resourceObjectName:        strcase.UpperCamelCase(o.GetName()),
									leafRefResourceName:       split[0],
									leafRefResourceObjectName: split[1],
									target:                    target.TargetName,
								}
								addFinalizer2Resource(lrr)
							} else {
								r.Log.Info("Remote LeafRef dependency is empty, somethign went wrong")
							}
							o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationDetails[localLeafRef].LocalResolvedLeafRefInfo[localLeafRefPath] = &srlinuxv1alpha1.RemoteLeafRefInfo{
								RemoteLeafRef:        RemoteLeafRefInfo.RemoteLeafRef,
								DependencyCheck:      RemoteLeafRefInfo.DependencyCheck,
								RemoteResourceObject: res,
							}
						}
					}
				} else {
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationDetails[localLeafRef] = &srlinuxv1alpha1.ValidationDetails{}
				}
			}

			if o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus == nil {
				if validationSuccess {
					r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s External Leafref Validation success", target.TargetName), ""))
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
				} else {
					r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s External Leafref Validation failed", target.TargetName), "Leaf Ref dependency missing"))
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
				}
			} else {
				if validationSuccess {
					// if the validation status was failed we want to update the event to indicate the success on the transition from failed -> success
					if *o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
						r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s External Leafref Validation success", target.TargetName), ""))
					}
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusSuccess)
				} else {
					// if the validation status did not change we dont have to publish a new event
					if *o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus != srlinuxv1alpha1.ValidationStatusFailed {
						r.publishEvent(req, o.NewEvent(fmt.Sprintf("Target: %s External Leafref Validation failed", target.TargetName), "Leaf Ref dependency missing"))
					}
					o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus = srlinuxv1alpha1.ValidationStatusPtr(srlinuxv1alpha1.ValidationStatusFailed)
				}
			}
		}
		if err := r.saveSrlRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to save status"))
		}
		// check validation status and requeue if a validation error is reported
		for _, target := range t {
			if *o.Status.Target[target.TargetName].ConfigurationDependencyParentValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
				return ctrl.Result{Requeue: true, RequeueAfter: parentDependencyRetyrDelay}, nil
			}
			if *o.Status.Target[target.TargetName].ConfigurationDependencyExternalLeafrefValidationStatus == srlinuxv1alpha1.ValidationStatusFailed {
				return ctrl.Result{Requeue: true, RequeueAfter: externalLeafRefvalidationErrorretryDelay}, nil
			}
		}
	}

	info := make(map[string]*SrlRoutingpolicyCommunitysetReconcileInfo)
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
		info[target.TargetName] = &SrlRoutingpolicyCommunitysetReconcileInfo{
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
		stateMachine := newSrlRoutingpolicyCommunitysetStateMachine(o, r, &target.TargetName, info[target.TargetName])
		actResult[target.TargetName] = stateMachine.ReconcileState(info[target.TargetName])
		result[target.TargetName], err = actResult[target.TargetName].Result()
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("action %q failed", *initialState))
			return result[target.TargetName], err
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
			if err := r.saveSrlRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
		}
		SrlRoutingpolicyCommunitysetlogResult(info[target.TargetName], result[target.TargetName])

		// requeue for action update and action continue
		if result[target.TargetName].Requeue {
			return ctrl.Result{Requeue: true, RequeueAfter: result[target.TargetName].RequeueAfter}, nil
		}
	}

	if !o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o) {
		deleted := true
		for _, target := range t {
			// delete entry in shared memory
			if err := deleteSharedMemoryEntry(stringPtr(target.TargetName), stringPtr("SrlInterface"), stringPtr(req.Name)); err != nil {
				return ctrl.Result{}, err
			}
			r.Log.WithValues("target", target.TargetName, "shared memory data", SrlSharedInfo[target.TargetName]).Info("shared memory info")

			//Remove configmap entry
			/*
				if err := r.deleteEntryConfigMap(ctx, stringPtr(target.TargetName), stringPtr(req.Name)); err != nil {
					return ctrl.Result{}, err
				}
			*/
			if result[target.TargetName].RequeueAfter != 0 {
				deleted = false
			}
		}
		if deleted {
			// delete complete
			r.Log.WithValues("Finalizers", o.Finalizers).Info("Finalizers")
			// remove the leafref dependency finalizers from remote leafref objects
			for targetName, ts := range o.Status.Target {
				// localLeafRef, leafRefInfo
				for _, leafRefInfo := range ts.ConfigurationDependencyExternalLeafrefValidationDetails {
					// localLeafRefPath, RemoteLeafRefInfo
					for _, RemoteLeafRefInfo := range leafRefInfo.LocalResolvedLeafRefInfo {
						if RemoteLeafRefInfo.RemoteResourceObject != nil {
							r.Log.WithValues("RemoteResourceObject", *RemoteLeafRefInfo.RemoteResourceObject).Info("Remote LeafRef Object")
							if *RemoteLeafRefInfo.RemoteResourceObject != "" {
								// first part of split, split[0] is the resource, 2nd part is the resourceName, split[1]
								split := strings.Split(*RemoteLeafRefInfo.RemoteResourceObject, ".")
								lrr := &LeafRefResource{
									ctx:                       ctx,
									client:                    r.Client,
									nameSpace:                 o.GetNamespace(),
									resourceName:              "CommunitySet",
									resourceObjectName:        strcase.UpperCamelCase(o.GetName()),
									leafRefResourceName:       split[0],
									leafRefResourceObjectName: split[1],
									target:                    targetName,
								}
								deleteFinalizer2Resource(lrr)
							} else {
								r.Log.Info("Remote LeafRef dependency is empty, somethign went wrong")
							}
						}
					}
				}
			}
			// remove our finalizer from the list and update it.
			r.Log.Info("removing finalizer")
			o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer)
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

	for _, ri := range info {
		for _, e := range ri.events {
			r.publishEvent(req, e)
		}
	}

	return ctrl.Result{}, nil
}

func SrlRoutingpolicyCommunitysetlogResult(info *SrlRoutingpolicyCommunitysetReconcileInfo, result ctrl.Result) {
	if result.Requeue || result.RequeueAfter != 0 ||
		!StringInList(info.o.Finalizers,
			srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer) {
		info.log.Info("done",
			"requeue", result.Requeue,
			"after", result.RequeueAfter)
	} else {
		info.log.Info("stopping reconcile SrlRoutingpolicyCommunityset")
		//info.log.Info("stopping reconcile SrlRoutingpolicyCommunityset",
		//	"message", info.o.Status)
	}
}

func (r *SrlRoutingpolicyCommunitysetReconciler) saveSrlRoutingpolicyCommunitysetStatus(ctx context.Context, o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) error {
	t := metav1.Now()
	o.Status.DeepCopy()
	o.Status.LastUpdated = &t

	r.Log.Info("SrlRoutingpolicyCommunityset",
		"status", o.Status)

	if err := r.Client.Status().Update(ctx, o); err != nil {
		r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to update SrlRoutingpolicyCommunityset ")
		return err
	}
	return nil
}

func (r *SrlRoutingpolicyCommunitysetReconciler) getConfigMap(ctx context.Context, targetName *string) (*string, error) {
	cmKey := types.NamespacedName{
		Namespace: "nddriver-system",
		Name:      "nddriver-cm-" + *targetName,
	}
	cm := &corev1.ConfigMap{}
	if err := r.Get(ctx, cmKey, cm); err != nil {
		r.Log.Error(err, "Failed to get configmap")
		return nil, err
	}

	if _, ok := cm.Data["config.json"]; !ok {
		r.Log.WithValues("targetName", targetName).Info("ConfigMap is empty")
	}
	//r.Log.WithValues("targetName", targetName).Info("ConfigMap content")
	return stringPtr(cm.Data["config.json"]), nil
}

/*
	func (r *SrlRoutingpolicyCommunitysetReconciler) addEntryConfigMap(ctx context.Context, targetName, oName *string, deletepaths *[]string) error {
		dp := &DeletePaths{
			DeletePaths: deletepaths,
		}

		d, err := json.Marshal(dp)
		if err != nil {
			r.Log.Error(err, "Failed to marshal data")
			return err
		}

		cmKey := types.NamespacedName{
			Namespace: "nddriver-system",
			Name:      "nddriver-cm-" + *targetName,
		}
		cm := &corev1.ConfigMap{}
		if err := r.Get(ctx, cmKey, cm); err != nil {
			r.Log.Error(err, "Failed to get configmap")
			return err
		}

		//cm.Data["SrlRoutingpolicyCommunityset"] = string(d)
		cm.Data[fmt.Sprintf("SrlRoutingpolicyCommunityset.%s",*oName)] = string(d)

		if err := r.Update(ctx, cm); err != nil {
			r.Log.Error(err, "Failed to update configmap")
			return err
		}

		return nil
	}
*/

/*
	func (r *SrlRoutingpolicyCommunitysetReconciler) deleteEntryConfigMap(ctx context.Context, targetName, oName *string) error {
		cmKey := types.NamespacedName{
			Namespace: "nddriver-system",
			Name:      "nddriver-cm-" + *targetName,
		}
		cm := &corev1.ConfigMap{}
		if err := r.Get(ctx, cmKey, cm); err != nil {
			r.Log.Error(err, "Failed to get configmap")
			return err
		}

		if _, ok := cm.Data[fmt.Sprintf("SrlRoutingpolicyCommunityset.%s",*oName)]; ok {
			delete(cm.Data, fmt.Sprintf("SrlRoutingpolicyCommunityset.%s",*oName))
			if err := r.Update(ctx, cm); err != nil {
				r.Log.Error(err, "Failed to update configmap")
				return err
			}
		}
		return nil
	}
*/

/*
	func (r *SrlRoutingpolicyCommunitysetReconciler) GetRemoteleafRefResource(ctx context.Context, targetName *string, remoteleafRef *srlinuxv1alpha1.RemoteLeafRefInfo) (*string, error) {
		// get configmap
		cmKey := types.NamespacedName{
			Namespace: "nddriver-system",
			Name:      "nddriver-cm-" + *targetName,
		}
		cm := &corev1.ConfigMap{}
		if err := r.Get(ctx, cmKey, cm); err != nil {
			r.Log.Error(err, "Failed to get configmap")
			return nil, err
		}
		resource := new(string)
		p := new(string)
		for res, dps := range cm.Data {
			if strings.HasPrefix(res, "Srl") {
				//r.Log.WithValues("DeletePaths", dps, "Remote Leafref", *remoteleafRef.RemoteLeafRef).Info("GetRemoteleafRefResource info")
				var x1 interface{}
				json.Unmarshal([]byte(dps), &x1)
				f, dp := matchDeletePath(x1, remoteleafRef.RemoteLeafRef)
				if f {
					//r.Log.WithValues("DeletePath", *dp).Info("Path Found")
					if len(*dp) > len(*p) {
						resource = stringPtr(res)
						*p = *dp
					}
				}
			}
		}
		return resource, nil
	}
*/

func (r *SrlRoutingpolicyCommunitysetReconciler) findPathInTree(x1 interface{}, ekvl []ElementKeyValue, idx int) bool {
	//r.Log.WithValues("ekvl", ekvl, "idx", idx, "Data", x1).Info("findPathInTree")

	switch x := x1.(type) {
	case map[string]interface{}:
		for k, x2 := range x {
			if k == ekvl[idx].Element {
				if idx == len(ekvl)-1 {
					// last element/index in ekv
					if ekvl[idx].KeyName != "" {
						//r.Log.WithValues("ElementName", k, "KeyName", ekvl[idx].KeyName).Info("findPathInTree map[string]interface{} Last Index")
						return r.findPathInTree(x2, ekvl, idx)
					} else {
						//r.Log.WithValues("ElementName", k, "KeyName", "").Info("findPathInTree map[string]interface{} Last Index")
						return true
					}
				} else {
					// not last element/index in ekv
					if ekvl[idx].KeyName != "" {
						//r.Log.WithValues("ElementName", k, "KeyName", ekvl[idx].KeyName).Info("findPathInTree map[string]interface{} Not Last Index")
						return r.findPathInTree(x2, ekvl, idx)
					} else {
						//r.Log.WithValues("ElementName", k, "KeyName", "").Info("findPathInTree map[string]interface{} Not Last Index")
						idx++
						return r.findPathInTree(x2, ekvl, idx)
					}
				}
			}
		}
	case []interface{}:
		for _, v := range x {
			switch x2 := v.(type) {
			case map[string]interface{}:
				for k3, x3 := range x2 {
					if k3 == ekvl[idx].KeyName {
						if idx == len(ekvl)-1 {
							switch x3.(type) {
							case string, uint32:
								if x3 == ekvl[idx].KeyValue {
									//r.Log.WithValues("ElementName", k3, "KeyName", "").Info("findPathInTree map[string]interface{} in []interface{} Last Index found")
									return true
								} else {
									//r.Log.WithValues("ElementName", k3, "KeyName", "").Info("findPathInTree map[string]interface{} in []interface{} Last Index not found")
								}
							}
						} else {
							//r.Log.WithValues("ElementName", k3, "KeyName", "").Info("findPathInTree map[string]interface{} in []interface{} Not Last Index")
							idx++
							return r.findPathInTree(x3, ekvl, idx)
						}
					}
				}
			}
		}
	case nil:
		//r.Log.WithValues("x1", x1).Info("findPathInTree nil")
		return false
	}
	r.Log.Info("findPathInTree end")
	return false
}

func (r *SrlRoutingpolicyCommunitysetReconciler) findLeafRefInTree(x1 interface{}, ekvl []ElementKeyValue, idx int, leafRefValues, localLeafRefPaths []string, lridx int) ([]string, []string) {
	//r.Log.WithValues("ekvl", ekvl, "idx", idx, "Data", x1, "leafRefValues", leafRefValues, "localLeafRefPath", localLeafRefPaths).Info("findLeafRefInTree")

	var tlrv []string
	switch x := x1.(type) {
	case map[string]interface{}:
		for k, x2 := range x {
			//r.Log.WithValues("Key", k, "Value", x2, "leafRefValues", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("map[string]interface{}")
			if k == ekvl[idx].Element {
				if idx == len(ekvl)-1 {
					// last element/index in ekv
					if ekvl[idx].KeyName != "" {
						//r.Log.WithValues("KeyName", ekvl[idx].KeyName).Info("map[string]interface{} Last Index")
						tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues, localLeafRefPaths, lridx)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					} else {
						switch x3 := x2.(type) {
						case string:
							//r.Log.WithValues("KeyName", "", "Value", x3, "Type", "string").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, x3)
							localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "=" + x3
							//return leafRefValuesPtr
						case int:
							x4 := strconv.Itoa(int(x3))
							//r.Log.WithValues("KeyName", "", "Value", x4, "Type", "int").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, x4)
							localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "=" + x4
							//return leafRefValuesPtr
						case float64:
							//r.Log.WithValues("KeyName", "", "Value", x3, "Type", "float64").Info("map[string]interface{} Last Index")
							leafRefValues = append(leafRefValues, fmt.Sprintf("%.0f", x3))
							localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "=" + fmt.Sprintf("%f", x3)

						default:
							//r.Log.WithValues("Type", reflect.TypeOf(x3)).Info("Default type")
							//r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} Last Index")
							//return leafRefValuesPtr
						}
					}
				} else {
					// not last element/index in ekv
					if ekvl[idx].KeyName != "" {
						//r.Log.WithValues("KeyName", ekvl[idx].KeyName).Info("map[string]interface{} Not Last Index")
						tlrv, localLeafRefPaths = r.findLeafRefInTree(x2, ekvl, idx, leafRefValues, localLeafRefPaths, lridx)
						//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
						if len(tlrv) > len(leafRefValues) {
							leafRefValues = tlrv
						}
						//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
					} else {
						//r.Log.WithValues("KeyName", "").Info("map[string]interface{} Not Last Index")
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
								//r.Log.WithValues("KeyName", "", "Value", x4, "Type", "string").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x4)
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + x4 + "]"
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")

							case int:
								x5 := strconv.Itoa(int(x4))
								//r.Log.WithValues("KeyName", "", "Value", x5, "Type", "int").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, x5)
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + x5 + "]"
								//r.Log.WithValues("leafRefValues", tlrv).Info("findLeafRefInTree return")
								//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
								//return leafRefValues
							case float64:
								//r.Log.WithValues("KeyName", "", "Value", x4, "Type", "float64").Info("map[string]interface{} in []interface{} Last Index")
								leafRefValues = append(leafRefValues, fmt.Sprintf("%.0f", x4))
								localLeafRefPaths[lridx] += "/" + ekvl[idx].Element + "[" + ekvl[idx].KeyName + "=" + fmt.Sprintf("%f", x4) + "]"
							default:
								//r.Log.WithValues("Type", reflect.TypeOf(x4)).Info("Default type")
								//r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} in []interface{} Last Index")
								//return leafRefValues
							}
						} else {
							//r.Log.WithValues("KeyName", "", "Value", nil, "Type", "Default").Info("map[string]interface{} in []interface{} Not Last Index")
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
		//r.Log.WithValues("x1", x1).Info("nil")
		//return leafRefValuesPtr
	}
	//r.Log.WithValues("leafRefValues", leafRefValues).Info("findLeafRefInTree return")
	return leafRefValues, localLeafRefPaths
}

func (r *SrlRoutingpolicyCommunitysetReconciler) FindRemoteLeafRef(remoteLeafRef string, d [][]byte, rekvl []ElementKeyValue) []string {
	//r.Log.WithValues("remoteLeafRef", remoteLeafRef, "rekvl", rekvl).Info("Find Remote LeafRef")
	leafRefValues := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	for _, b := range d {
		var x1 interface{}
		json.Unmarshal(b, &x1)

		localLeafRefPaths = append(localLeafRefPaths, "")
		leafRefValues, localLeafRefPaths = r.findLeafRefInTree(x1, rekvl, 0, leafRefValues, localLeafRefPaths, 0)
		//r.Log.WithValues("remoteLeafRef", remoteLeafRef, "Values", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("Find remote LeafRef Values")
	}
	return leafRefValues
}

func (r *SrlRoutingpolicyCommunitysetReconciler) FindLocalLeafRef(localLeafRef string, d [][]byte, ekvl, rekvl []ElementKeyValue) ([]string, []string) {
	//r.Log.WithValues("ekvl", ekvl, "rekvl", rekvl).Info("find LeafRef")
	leafRefDependencies := make([]string, 0)
	localLeafRefPaths := make([]string, 0)
	for _, b := range d {
		var x1 interface{}
		json.Unmarshal(b, &x1)

		leafRefValues := make([]string, 0)
		localLeafRefPaths = append(localLeafRefPaths, "")
		leafRefValues, localLeafRefPaths = r.findLeafRefInTree(x1, ekvl, 0, leafRefValues, localLeafRefPaths, 0)
		//r.Log.WithValues("LocalLeafRef", localLeafRef, "Values", leafRefValues, "localLeafRefPaths", localLeafRefPaths).Info("find LeafRef Values")
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

// FindSpecDiff tries to understand the difference from the latest spec to the newest spec
func (r *SrlRoutingpolicyCommunitysetReconciler) FindSpecDiff(ctx context.Context, o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) (bool, *[]string, error) {
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
func (r *SrlRoutingpolicyCommunitysetReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) ([]*Target, bool, error) {
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

// SrlRoutingpolicyCommunitysethasFinalizer checks if object has finalizer
func SrlRoutingpolicyCommunitysethasFinalizer(o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) bool {
	return StringInList(o.Finalizers, srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer)
}

// SSrlRoutingpolicyCommunitysethasOtherFinalizer checks if object has other finalizers
func SrlRoutingpolicyCommunitysethasOtherFinalizer(o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset) bool {
	for _, f := range o.Finalizers {
		if f != srlinuxv1alpha1.SrlRoutingpolicyCommunitysetFinalizer {
			return true
		}
	}
	return false
}

func (info *SrlRoutingpolicyCommunitysetReconcileInfo) DeleteCache() error {
	if !info.o.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(info.o) {

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
func (info *SrlRoutingpolicyCommunitysetReconcileInfo) UpdateCache(path string) error {

	// marshal data to json
	d := make([][]byte, 0)
	for _, obj := range *info.o.Spec.SrlRoutingpolicyCommunityset {
		o := make([]srlinuxv1alpha1.RoutingpolicyCommunityset, 0)
		o = append(o, obj)
		dd := struct {
			CommunitySet *[]srlinuxv1alpha1.RoutingpolicyCommunityset `json:"community-set"`
		}{
			CommunitySet: &o,
		}
		dj, err := json.Marshal(dd)
		if err != nil {
			return err
		}
		d = append(d, dj)
	}
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

type SrlRoutingpolicyCommunitysetStateMachine struct {
	Object     *srlinuxv1alpha1.SrlRoutingpolicyCommunityset
	Reconciler *SrlRoutingpolicyCommunitysetReconciler
	Target     *string
	TargetName *string
	NextState  *srlinuxv1alpha1.ConfigStatus
}

// appendEvent
func (info *SrlRoutingpolicyCommunitysetReconcileInfo) appendEvent(reason, message string) {
	info.events = append(info.events, info.o.NewEvent(reason, message))
}

func newSrlRoutingpolicyCommunitysetStateMachine(o *srlinuxv1alpha1.SrlRoutingpolicyCommunityset,
	reconciler *SrlRoutingpolicyCommunitysetReconciler, n *string,
	info *SrlRoutingpolicyCommunitysetReconcileInfo) *SrlRoutingpolicyCommunitysetStateMachine {
	currentState := o.Status.Target[*n].ConfigStatus
	r := SrlRoutingpolicyCommunitysetStateMachine{
		Object:     o,
		NextState:  currentState, // Remain in current state by default
		Reconciler: reconciler,
		Target:     info.target,
		TargetName: n,
	}
	return &r
}

type SrlRoutingpolicyCommunitysetstateHandler func(*SrlRoutingpolicyCommunitysetReconcileInfo) actionResult

func (o *SrlRoutingpolicyCommunitysetStateMachine) handlers() map[srlinuxv1alpha1.ConfigStatus]SrlRoutingpolicyCommunitysetstateHandler {
	return map[srlinuxv1alpha1.ConfigStatus]SrlRoutingpolicyCommunitysetstateHandler{
		srlinuxv1alpha1.ConfigStatusNone:             o.handleNone,
		srlinuxv1alpha1.ConfigStatusConfiguring:      o.handleConfiguring,
		srlinuxv1alpha1.ConfigStatusConfigureSuccess: o.handleConfigureSuccess,
		srlinuxv1alpha1.ConfigStatusConfigureFailed:  o.handleConfigureFailed,
		srlinuxv1alpha1.ConfigStatusDeleting:         o.handleDeleting,
		srlinuxv1alpha1.ConfigStatusDeleteSuccess:    o.handleDeleteSuccess,
		srlinuxv1alpha1.ConfigStatusDeleteFailed:     o.handleDeleteFailed,
	}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) updateSrlRoutingpolicyCommunitysetStateFrom(initialState *srlinuxv1alpha1.ConfigStatus,
	info *SrlRoutingpolicyCommunitysetReconcileInfo) {
	if o.NextState != initialState {
		info.log.Info("changing configuration state",
			"old", initialState,
			"new", o.NextState)
		o.Object.Status.Target[*o.TargetName].ConfigStatus = o.NextState

		info.appendEvent(fmt.Sprintf("Target: %s, Configuration status old: %s -> new: %s", *o.TargetName, initialState.String(), o.NextState.String()), "")
	}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) ReconcileState(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	initialState := o.Object.Status.Target[*o.TargetName].ConfigStatus
	defer o.updateSrlRoutingpolicyCommunitysetStateFrom(initialState, info)

	if o.checkInitiateDelete() {
		// initiate cache delete
		info.log.Info("Initiating SrlRoutingpolicyCommunitysetStateMachine deletion")
		info.DeleteCache()
	}

	if stateHandler, found := o.handlers()[*initialState]; found {
		return stateHandler(info)
	}

	info.log.Info("No handler found for state", "state", initialState)
	return actionError{fmt.Errorf("No handler found for state \"%s\"", *initialState)}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) checkInitiateDelete() bool {
	if !o.Object.DeletionTimestamp.IsZero() && SrlRoutingpolicyCommunitysethasFinalizer(o.Object) {
		// Delete requested
		switch *o.NextState {
		default:
			// new state deleting
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
		case srlinuxv1alpha1.ConfigStatusDeleting,
			srlinuxv1alpha1.ConfigStatusDeleteFailed,
			srlinuxv1alpha1.ConfigStatusDeleteSuccess:
			// Already in deleting state. Allow state machine to run.
			return false
		}
		return true
	}
	// delete not requested
	return false
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleNone(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	//info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			info.log.Info("object status",
				"target", o.Target,
				"status", o.Object.Status)
			*o.NextState = srlinuxv1alpha1.ConfigStatusConfigureSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
			return actionComplete{}
		}
	}
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleteSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
			return actionUpdate{delay: 10 * time.Second}
		}
	} else {
		// update action
		*o.NextState = srlinuxv1alpha1.ConfigStatusConfiguring
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
		o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(cr.Status.String()))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleConfiguring(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	//info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			*o.NextState = srlinuxv1alpha1.ConfigStatusConfigureSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
			return actionComplete{}
		}
		if cr.Data.Action == netwdevpb.CacheUpdateRequest_Delete {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
			return actionContinue{}
		}
	} else {
		info.log.Info("Object got removed by the device driver, most likely due to restart of device driver")
		*o.NextState = srlinuxv1alpha1.ConfigStatusNone
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone))
		return actionUpdate{delay: 1 * time.Second}
	}
	if o.NextState == srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting) {
		// delete action
		if !cr.Exists {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleteSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
			return actionUpdate{delay: 10 * time.Second}
		}
	} else {
		// update action
		o.NextState = srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring)
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleConfigureSuccess(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	info.log.Info("handleConfigStatusConfigureSuccess CacheStatusResponse", "Response", cr)
	info.log.Info("handleConfigStatusConfigureSuccess NextState", "NextState", *o.NextState)
	if *o.NextState == srlinuxv1alpha1.ConfigStatusDeleting {
		info.log.Info("handleConfigStatusConfigureSuccess -> Next State ConfigStatusDeleting")
		// delete action
		*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
		// this was done to handle a recirculation and allow time for the leafref dependency finalizers to be removed
		return actionUpdate{delay: 5 * time.Second}
	}
	if !cr.Exists {
		info.log.Info("Object got removed by the device driver, most likely due to restart of device driver")
		*o.NextState = srlinuxv1alpha1.ConfigStatusNone
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusNone))
		return actionUpdate{delay: 1 * time.Second}
	}

	return actionComplete{}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleConfigureFailed(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	//info.log.Info("CacheStatusResponse", "Response", cr)
	if cr.Exists {
		if cr.Status == netwdevpb.CacheStatusReply_UpdateProcessedSuccess {
			*o.NextState = srlinuxv1alpha1.ConfigStatusConfigureSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfigureSuccess))
			o.Object.SetConfigStatusDetails(o.TargetName, stringPtr(""))
			return actionComplete{}
		}
	}
	if *o.NextState == srlinuxv1alpha1.ConfigStatusDeleting {
		// delete action
		if !cr.Exists {
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleteSuccess
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
			return actionComplete{}
		} else {
			// delete action
			*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
			o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
			// this was done to handle a recirculation and allow time for the leafref dependency finalizers to be removed
			return actionUpdate{delay: 5 * time.Second}
		}
	} else {
		// update action
		*o.NextState = srlinuxv1alpha1.ConfigStatusConfiguring
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusConfiguring))
	}
	return actionUpdate{delay: 10 * time.Second}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleDeleting(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
		// delete action
		*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
		// this was done to handle a recirculation and allow time for the leafref dependency finalizers to be removed
		return actionUpdate{delay: 10 * time.Second}
	}

}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleDeleteSuccess(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	return actionFinished{}
}

func (o *SrlRoutingpolicyCommunitysetStateMachine) handleDeleteFailed(info *SrlRoutingpolicyCommunitysetReconcileInfo) actionResult {
	cr, err := getCachStatus(o.Reconciler.Ctx, o.Target, info.resource, *info.level)
	if err != nil {
		return actionFailed{dirty: true, errorCount: *info.o.Status.Target[*o.TargetName].ErrorCount}
	}
	//info.log.Info("CacheStatusResponse", "Response", cr)

	if !cr.Exists {
		*o.NextState = srlinuxv1alpha1.ConfigStatusDeleteSuccess
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleteSuccess))
		return actionFinished{}
	} else {
		// delete action
		*o.NextState = srlinuxv1alpha1.ConfigStatusDeleting
		o.Object.SetConfigStatus(o.TargetName, srlinuxv1alpha1.ConfigStatusPtr(srlinuxv1alpha1.ConfigStatusDeleting))
		// this was done to handle a recirculation and allow time for the leafref dependency finalizers to be removed
		return actionUpdate{delay: 10 * time.Second}
	}
}
