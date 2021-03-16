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

	"github.com/stoewer/go-strcase"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	nddv1 "github.com/netw-device-driver/netw-device-controller/api/v1"

	"github.com/netw-device-driver/netwdevpb"
	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	"github.com/srl-wim/srl-k8s-operator/pkg/natssrl"
)

// K8sSrlNokiaInterfacesInterfaceReconciler reconciles a SrlNokiaInterfacesInterface object
type K8sSrlNokiaInterfacesInterfaceReconciler struct {
	client.Client
	Server *string
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ctx    context.Context
}

// +kubebuilder:rbac:groups=ndd.henderiw.be,resources=networknodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=k8ssrlnokiainterfacesinterfaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=k8ssrlnokiainterfacesinterfaces/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=srlinux.henderiw.be,resources=k8ssrlnokiainterfacesinterfaces/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SrlNokiaInterfacesInterface object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *K8sSrlNokiaInterfacesInterfaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("K8sSrlNokiaInterfacesInterface", req.NamespacedName)

	r.Log.Info("reconciling K8sSrlNokiaInterfacesInterface")

	o := &srlinuxv1alpha1.K8sSrlNokiaInterfacesInterface{}
	if err := r.Client.Get(ctx, req.NamespacedName, o); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		r.Log.WithValues(req.Name, req.Namespace).Error(err, "Failed to get K8sSrlNokiaInterfacesInterface config")
		return ctrl.Result{}, err
	}
	o.DeepCopy()

	r.Log.WithValues("Object", o).Info("Object Info")

	// Add a finalizer to newly created objects.
	if o.DeletionTimestamp.IsZero() && !SrlNokiaInterfacesInterfacehasFinalizer(o) {
		r.Log.Info(
			"adding finalizer",
			"existingFinalizers", o.Finalizers,
			"newValue", srlinuxv1alpha1.SrlNokiaInterfacesInterfaceFinalizer,
		)
		o.Finalizers = append(o.Finalizers,
			srlinuxv1alpha1.SrlNokiaInterfacesInterfaceFinalizer)
		err := r.Update(context.TODO(), o)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "failed to add finalizer")
		}
		return ctrl.Result{Requeue: true}, nil
	}

	t, err := r.FindTarget(ctx, o)
	if err != nil {
		switch err.(type) {
		case *TargetNotFoundError:
			// when no target is available requeue to retry after requetimer
			return ctrl.Result{Requeue: true, RequeueAfter: targetNotFoundRetryDelay}, nil
		default:
			return ctrl.Result{}, err
		}
	}
	r.Log.WithValues("Target", *t).Info("Target Info")

	level := int32(1)
	topic := "ndd" + "." + *t + "." + "K8sSrlNokiaInterfacesInterface" + strcase.UpperCamelCase(o.Name)

	hkeys0 := make([]string, 0)
	for _, n := range *o.Spec.SrlNokiaInterfacesInterface {
		hkeys0 = append(hkeys0, *n.Name)
	}

	dependencies := make([]string, 0)

	deletepaths := make([]string, 0)

	for _, hkey0 := range hkeys0 {
		deletepaths = append(deletepaths, fmt.Sprintf("/interface[name=%s]", hkey0))
	}

	// The object is being deleted
	if !o.DeletionTimestamp.IsZero() && SrlNokiaInterfacesInterfacehasFinalizer(o) {

		// prepare the nats data
		data := &netwdevpb.ConfigMessage{
			Level:                level,
			Action:               netwdevpb.ConfigMessage_Delete,
			IndividualActionPath: deletepaths,
			Dependencies:         dependencies,
		}

		n := &natssrl.Client{
			Server: *r.Server,
			Topic:  topic,
		}
		n.Publish(data)

		r.Log.WithValues(
			"Topic", n.Topic).WithValues(
			"Data", data).Info("Published data")

		// remove our finalizer from the list and update it.
		o.Finalizers = removeString(o.Finalizers, srlinuxv1alpha1.SrlNokiaInterfacesInterfaceFinalizer)
		if err := r.Update(context.Background(), o); err != nil {
			return ctrl.Result{}, errors.Wrap(err,
				fmt.Sprintf("failed to remove finalizer"))
		}
		r.Log.Info("cleanup is complete, removed finalizer",
			"remaining", o.Finalizers)
		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	// path to be used for this object

	path := "/"

	// marshal data to json

	d := make([][]byte, 0)
	for _, obj := range *o.Spec.SrlNokiaInterfacesInterface {
		o := make([]srlinuxv1alpha1.SrlNokiaInterfacesInterface, 0)
		o = append(o, obj)
		dd := struct {
			Interface *[]srlinuxv1alpha1.SrlNokiaInterfacesInterface `json:"interface"`
		}{
			Interface: &o,
		}
		dj, err := json.Marshal(dd)
		if err != nil {
			return ctrl.Result{}, err
		}
		d = append(d, dj)
	}

	// prepare the
	data := &netwdevpb.ConfigMessage{
		Level:                level,
		Action:               netwdevpb.ConfigMessage_Update,
		AggregateActionPath:  path,
		IndividualActionPath: deletepaths,
		Data:                 d,
		Dependencies:         dependencies,
	}

	n := &natssrl.Client{
		Server: *r.Server,
		Topic:  topic,
	}
	n.Publish(data)

	r.Log.WithValues(
		"Topic", n.Topic).WithValues(
		"Path", path).WithValues(
		"Data", data).Info("Published data")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *K8sSrlNokiaInterfacesInterfaceReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, option controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&srlinuxv1alpha1.K8sSrlNokiaInterfacesInterface{}).
		WithOptions(option).
		Complete(r)
}

// FindTarget finds the SRL target for Object
func (r *K8sSrlNokiaInterfacesInterfaceReconciler) FindTarget(ctx context.Context, o *srlinuxv1alpha1.K8sSrlNokiaInterfacesInterface) (target *string, err error) {
	r.Log.Info("Find target ...")
	var targetName string
	for k, v := range o.Labels {
		if k == "target" {
			targetName = v
		}
	}
	// get network node list
	selectors := []client.ListOption{
		client.MatchingLabels{},
	}

	nn := &nddv1.NetworkNodeList{}
	if err := r.List(r.Ctx, nn, selectors...); err != nil {
		r.Log.Error(err, "Failed to get NetworkNode List ")
		return nil, err
	}
	for _, n := range nn.Items {
		if n.Name == targetName {
			target = &targetName
		}
	}
	// Target not found
	if target == nil {
		return target, &TargetNotFoundError{message: "The Target cannot be found, update label or discovery object"}
	}
	return target, nil
}

// SrlNokiaInterfacesInterfacehasFinalizer checks if object has finalizer
func SrlNokiaInterfacesInterfacehasFinalizer(o *srlinuxv1alpha1.K8sSrlNokiaInterfacesInterface) bool {
	return StringInList(o.Finalizers, srlinuxv1alpha1.SrlNokiaInterfacesInterfaceFinalizer)
}
