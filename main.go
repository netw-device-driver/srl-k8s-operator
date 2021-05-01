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
	if err := (&controllers.SrlnokiaBfdReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaBfd"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaBfd")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaInterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaInterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaInterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaInterfaceSubinterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaInterfaceSubinterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaInterfaceSubinterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstance"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstance")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceAggregateroutesReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceAggregateroutes"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceAggregateroutes")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceNexthopgroupsReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceNexthopgroups"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceNexthopgroups")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsBgpReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsBgp"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsBgp")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsBgpevpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsBgpevpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsBgpevpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsBgpvpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsBgpvpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsBgpvpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsIsisReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsIsis"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsIsis")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsLinuxReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsLinux"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsLinux")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceProtocolsOspfReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceProtocolsOspf"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceProtocolsOspf")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaNetworkinstanceStaticroutesReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaNetworkinstanceStaticroutes"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaNetworkinstanceStaticroutes")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaRoutingpolicyAspathsetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaRoutingpolicyAspathset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaRoutingpolicyAspathset")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaRoutingpolicyCommunitysetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaRoutingpolicyCommunityset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaRoutingpolicyCommunityset")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaRoutingpolicyPolicyReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaRoutingpolicyPolicy"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaRoutingpolicyPolicy")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaRoutingpolicyPrefixsetReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaRoutingpolicyPrefixset"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaRoutingpolicyPrefixset")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaSystemMtuReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaSystemMtu"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaSystemMtu")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaSystemNameReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaSystemName"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaSystemName")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaSystemNetworkinstanceProtocolsBgpvpnReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaSystemNetworkinstanceProtocolsBgpvpn"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaSystemNetworkinstanceProtocolsBgpvpn")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaSystemNtpReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaSystemNtp"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaSystemNtp")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaTunnelinterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaTunnelinterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaTunnelinterface")
		os.Exit(1)
	}
	if err := (&controllers.SrlnokiaTunnelinterfaceVxlaninterfaceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("Controller").WithName("SrlnokiaTunnelinterfaceVxlaninterface"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SrlnokiaTunnelinterfaceVxlaninterface")
		os.Exit(1)
	}
}

func concurrency(c int) controller.Options {
	return controller.Options{MaxConcurrentReconciles: c}
}
