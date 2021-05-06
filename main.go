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

package main

import (
	"context"
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	nddv1 "github.com/netw-device-driver/netw-device-controller/api/v1"
	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	"github.com/srl-wim/srl-k8s-operator/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme         = runtime.NewScheme()
	setupLog       = ctrl.Log.WithName("setup")
	srlConcurrency int
	//natsServer string
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(srlinuxv1alpha1.AddToScheme(scheme))

	utilruntime.Must(nddv1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	//flag.StringVar(&natsServer, "nats-server", "",
	//	"The address the natsServer to subscribe to")
	flag.IntVar(&srlConcurrency, "srl-oncurrency", 1,
		"Number of items to process simultaneously")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "4047ae29.henderiw.be",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Setup the context that's going to be used in controllers and for the manager.
	ctx := ctrl.SetupSignalHandler()

	setupReconcilers(ctx, mgr)

	// +kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("health", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("check", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func setupReconcilers(ctx context.Context, mgr ctrl.Manager) {
	if err := (&controllers.SrlBfdReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlBfd"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlBfd")
		os.Exit(1)
	}
	if err := (&controllers.SrlInterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlInterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlInterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlInterfaceSubinterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlInterfaceSubinterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlInterfaceSubinterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstance"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstance")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceAggregateroutesReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceAggregateroutes"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceAggregateroutes")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceNexthopgroupsReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceNexthopgroups"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceNexthopgroups")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsBgpReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsBgp"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsBgp")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsBgpevpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsBgpevpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsBgpevpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsBgpvpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsBgpvpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsBgpvpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsIsisReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsIsis"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsIsis")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsLinuxReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsLinux"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsLinux")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceProtocolsOspfReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceProtocolsOspf"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceProtocolsOspf")
		os.Exit(1)
	}
	if err := (&controllers.SrlNetworkinstanceStaticroutesReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlNetworkinstanceStaticroutes"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlNetworkinstanceStaticroutes")
		os.Exit(1)
	}
	if err := (&controllers.SrlRoutingpolicyAspathsetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlRoutingpolicyAspathset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlRoutingpolicyAspathset")
		os.Exit(1)
	}
	if err := (&controllers.SrlRoutingpolicyCommunitysetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlRoutingpolicyCommunityset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlRoutingpolicyCommunityset")
		os.Exit(1)
	}
	if err := (&controllers.SrlRoutingpolicyPolicyReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlRoutingpolicyPolicy"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlRoutingpolicyPolicy")
		os.Exit(1)
	}
	if err := (&controllers.SrlRoutingpolicyPrefixsetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlRoutingpolicyPrefixset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlRoutingpolicyPrefixset")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemMtuReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemMtu"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemMtu")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemNameReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemName"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemName")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemNetworkinstanceProtocolsBgpvpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemNetworkinstanceProtocolsBgpvpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemNetworkinstanceProtocolsBgpvpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi")
		os.Exit(1)
	}
	if err := (&controllers.SrlSystemNtpReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlSystemNtp"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlSystemNtp")
		os.Exit(1)
	}
	if err := (&controllers.SrlTunnelinterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlTunnelinterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlTunnelinterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlTunnelinterfaceVxlaninterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlTunnelinterfaceVxlaninterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlTunnelinterfaceVxlaninterface")
		os.Exit(1)
	}
}

func concurrency(c int) controller.Options {
	return controller.Options{MaxConcurrentReconciles: c}
}
