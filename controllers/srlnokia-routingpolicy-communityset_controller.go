
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

	var RoutingpolicyCommunitysetIntraResourceleafRef= map[string]*LeafRef{
	}

	var RoutingpolicyCommunitysetInterResourceleafRef = map[string]*LeafRef{
	}
	
	// SrlnokiaRoutingpolicyCommunitysetReconciler reconciles a SrlnokiaRoutingpolicyCommunityset object
	type SrlnokiaRoutingpolicyCommunitysetReconciler struct {
		client.Client
		Log    logr.Logger
		Scheme *runtime.Scheme
		Ctx    context.Context
	}

	// Instead of passing a zillion arguments to the action of a phase,
	// hold them in a context per device
	type SrlnokiaRoutingpolicyCommunitysetTargetReconcileInfo struct {
		Target map[string]*SrlnokiaRoutingpolicyCommunitysetReconcileInfo
	}

	type SrlnokiaRoutingpolicyCommunitysetReconcileInfo struct {
		target            *string
		resource          *string
		o                 *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset
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
	// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokiaroutingpolicycommunitysets,verbs=get;list;watch;create;update;patch;delete
	// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokiaroutingpolicycommunitysets/status,verbs=get;update;patch
	// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=srlnokiaroutingpolicycommunitysets/finalizers,verbs=update
	
	// SetupWithManager sets up the controller with the Manager.
	func (r *SrlnokiaRoutingpolicyCommunitysetReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
		b := ctrl.NewControllerManagedBy(mgr).
			For(&srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset{}).
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
	// request for reconciliation of SrlnokiaRoutingpolicyCommunityset.
	func (r *SrlnokiaRoutingpolicyCommunitysetReconciler) NetworkDeviceMapFunc(o client.Object) []ctrl.Request {
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
		os := &srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetList{}
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

	// Reconcile is part of the main kubernetes reconciliation loop which aims to
	// move the current state of the cluster closer to the desired state.
	// TODO(user): Modify the Reconcile function to compare the state specified by
	// the SrlnokiaRoutingpolicyCommunityset object against the actual cluster state, and then
	// perform operations to make the cluster state reflect the state specified by
	// the user.
	//
	// For more details, check Reconcile and its Result here:
	// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
	func (r *SrlnokiaRoutingpolicyCommunitysetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
		_ = r.Log.WithValues("SrlnokiaRoutingpolicyCommunityset", req.NamespacedName)

		r.Log.WithValues("ObjectName", req.NamespacedName).Info("reconciling SrlnokiaRoutingpolicyCommunityset")

		o := &srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset{}
		if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get SrlnokiaRoutingpolicyCommunityset config")
			return ctrl.Result{}, err
		}
		o.DeepCopy()

		//r.Log.WithValues("Object", o).Info("Object Info")

		// Add a finalizer to newly created objects.
		if o.DeletionTimestamp.IsZero() && !SrlnokiaRoutingpolicyCommunitysethasFinalizer(o) {
			r.Log.Info(
				"adding finalizer",
				"existingFinalizers", o.Finalizers,
				"newValue", srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetFinalizer,
			)
			o.Finalizers = append(o.Finalizers,
				srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetFinalizer)
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
					if err = r.saveSrlnokiaRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
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
			if err = r.saveSrlnokiaRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
				return ctrl.Result{}, errors.Wrap(err,
					fmt.Sprintf("failed to save status"))
			}
		}
		r.Log.WithValues("Targets", t).Info("Target Info")

		// initialize the resource parameters
		level := int32(2)
		resource := "srlinux.henderiw.be" + "." + "SrlnokiaRoutingpolicyCommunityset" + "." + strcase.UpperCamelCase(o.Name)
		hkeys0 := make([]string, 0)
		for _, n := range *o.Spec.SrlnokiaRoutingpolicyCommunityset {
			hkeys0 = append(hkeys0, *n.Name)
		}
	
		dependencies := make([]string, 0)

		deletepaths := make([]string, 0)
		for _, hkey0 := range hkeys0 {
			deletepaths = append(deletepaths, fmt.Sprintf("/routing-policy/community-set[name=%s]", hkey0))
		}

		// path to be used for this object
		path := "/routing-policy"

		info := make(map[string]*SrlnokiaRoutingpolicyCommunitysetReconcileInfo)
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
			info[target.TargetName] = &SrlnokiaRoutingpolicyCommunitysetReconcileInfo{
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
			stateMachine := newSrlnokiaRoutingpolicyCommunitysetStateMachine(o, r, &target.TargetName, info[target.TargetName])
			actResult[target.TargetName] = stateMachine.ReconcileState(info[target.TargetName])
			result[target.TargetName], err = actResult[target.TargetName].Result()
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("action %q failed", *initialState))
				return result[target.TargetName], err
			}
		}

		if !o.DeletionTimestamp.IsZero() && SrlnokiaRoutingpolicyCommunitysethasFinalizer(o) {
			deleted := true
			for _, target := range t {
				if result[target.TargetName].RequeueAfter != 0 {
					deleted = false
				}
			}
			if deleted {
				// delete complete
				// remove our finalizer from the list and update it.
				o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetFinalizer)
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
				if err := r.saveSrlnokiaRoutingpolicyCommunitysetStatus(ctx, o); err != nil {
					return ctrl.Result{}, errors.Wrap(err,
						fmt.Sprintf("failed to save status"))
				}
			}
			SrlnokiaRoutingpolicyCommunitysetlogResult(info[target.TargetName], result[target.TargetName])
	
			// requeue for action update and action continue
			if result[target.TargetName].Requeue {
				return ctrl.Result{Requeue: true, RequeueAfter: result[target.TargetName].RequeueAfter}, nil
			}
		}
	
		return ctrl.Result{}, nil
	}
	
	func SrlnokiaRoutingpolicyCommunitysetlogResult(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo, result ctrl.Result) {
		if result.Requeue || result.RequeueAfter != 0 ||
			!utils.StringInList(info.o.Finalizers,
				srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetFinalizer) {
			info.log.Info("done",
				"requeue", result.Requeue,
				"after", result.RequeueAfter)
		} else {
			info.log.Info("stopping on SrlnokiaRoutingpolicyCommunityset",
				"message", info.o.Status)
		}
	}

	func (r *SrlnokiaRoutingpolicyCommunitysetReconciler) saveSrlnokiaRoutingpolicyCommunitysetStatus(ctx context.Context, o *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset) error {
		t := metav1.Now()
		o.Status.DeepCopy()
		o.Status.LastUpdated = &t
	
		r.Log.Info("SrlnokiaRoutingpolicyCommunityset",
			"status", o.Status)
	
		if err := r.Client.Status().Update(ctx, o); err != nil {
			r.Log.WithValues(o.Name, o.Namespace).Error(err, "Failed to update SrlnokiaRoutingpolicyCommunityset ")
			return err
		}
		return nil
	}

	// FindTarget finds the SRL target for Object
	func (r *SrlnokiaRoutingpolicyCommunitysetReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset) ([]*Target, bool, error) {
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

	// SrlnokiaRoutingpolicyCommunitysethasFinalizer checks if object has finalizer
	func SrlnokiaRoutingpolicyCommunitysethasFinalizer(o *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset) bool {
		return StringInList(o.Finalizers, srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunitysetFinalizer)
	}

	func (info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) DeleteCache(deletepaths, dependencies *[]string) error {
		if !info.o.DeletionTimestamp.IsZero() && SrlnokiaRoutingpolicyCommunitysethasFinalizer(info.o) {

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
	func (info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) UpdateCache(path string, deletepaths, dependencies []string) error {
		
		// marshal data to json
		
		
		
		
		
		
		
		
		d := make([][]byte, 0)
		for _, obj := range *info.o.Spec.SrlnokiaRoutingpolicyCommunityset {
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

	type SrlnokiaRoutingpolicyCommunitysetStateMachine struct {
		Object     *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset
		Reconciler *SrlnokiaRoutingpolicyCommunitysetReconciler
		Target     *string
		TargetName *string
		NextState  *srlinuxv1alpha1.ConfigStatus
	}
	
	// appendEvent
	func (info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) appendEvent(reason, message string) {
		info.events = append(info.events, info.o.NewEvent(reason, message))
	}
	
	func newSrlnokiaRoutingpolicyCommunitysetStateMachine(o *srlinuxv1alpha1.SrlnokiaRoutingpolicyCommunityset,
		reconciler *SrlnokiaRoutingpolicyCommunitysetReconciler, n *string,
		info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) *SrlnokiaRoutingpolicyCommunitysetStateMachine {
		currentState := o.Status.Target[*n].ConfigStatus
		r := SrlnokiaRoutingpolicyCommunitysetStateMachine{
			Object:     o,
			NextState:  currentState, // Remain in current state by default
			Reconciler: reconciler,
			Target:     info.target,
			TargetName: n,
		}
		return &r
	}
	
	type SrlnokiaRoutingpolicyCommunitysetstateHandler func(*SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handlers() map[srlinuxv1alpha1.ConfigStatus]SrlnokiaRoutingpolicyCommunitysetstateHandler {
		return map[srlinuxv1alpha1.ConfigStatus]SrlnokiaRoutingpolicyCommunitysetstateHandler{
			srlinuxv1alpha1.ConfigStatusNone:             o.handleNone,
			srlinuxv1alpha1.ConfigStatusConfiguring:      o.handleConfiguring,
			srlinuxv1alpha1.ConfigStatusConfigureSuccess: o.handleConfigStatusConfigureSuccess,
			srlinuxv1alpha1.ConfigStatusConfigureFailed:  o.handleConfigStatusConfigureFailed,
			srlinuxv1alpha1.ConfigStatusDeleting:         o.handleDeleting,
		}
	}
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) updateSrlnokiaRoutingpolicyCommunitysetStateFrom(initialState *srlinuxv1alpha1.ConfigStatus,
		info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) {
		if o.NextState != initialState {
			info.log.Info("changing configuration state",
				"old", initialState,
				"new", o.NextState)
			o.Object.Status.Target[*o.TargetName].ConfigStatus = o.NextState
		}
	}
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) ReconcileState(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
		initialState := o.Object.Status.Target[*o.TargetName].ConfigStatus
		defer o.updateSrlnokiaRoutingpolicyCommunitysetStateFrom(initialState, info)
	
		if o.checkInitiateDelete() {
			// initiate cache delete
			info.log.Info("Initiating SrlnokiaRoutingpolicyCommunitysetStateMachine deletion")
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) checkInitiateDelete() bool {
		if !o.Object.DeletionTimestamp.IsZero() && SrlnokiaRoutingpolicyCommunitysethasFinalizer(o.Object) {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handleNone(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
		}
		return actionUpdate{delay: 10 * time.Second}
	}
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handleConfiguring(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handleConfigStatusConfigureSuccess(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handleConfigStatusConfigureFailed(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) handleDeleting(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) DeleteFailed(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
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
	
	func (o *SrlnokiaRoutingpolicyCommunitysetStateMachine) DeleteSuccess(info *SrlnokiaRoutingpolicyCommunitysetReconcileInfo) actionResult {
		return actionComplete{}
	}
