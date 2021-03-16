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
	natsServer     string
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
	flag.StringVar(&natsServer, "nats-server", "",
		"The address the natsServer to subscribe to")
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

	if err := (&controllers.K8sSrlNokiaInterfacesInterfaceReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaInterfacesInterface}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaInterfacesInterface")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaInterfacesInterfaceSubinterfaceReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaInterfacesInterfaceSubinterface}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaInterfacesInterfaceSubinterface")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstance}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstance")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutesReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceAggregateRoutes")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroupsReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceNextHopGroups")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgp")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpnReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpevpn")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpnReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsBgpVpn")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsisReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsIsis")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinuxReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsLinux")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspfReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceProtocolsOspf")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutesReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaNetworkInstanceNetworkInstanceStaticRoutes")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaSystemSystemMtuReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaSystemSystemMtu}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaSystemSystemMtu")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaSystemSystemNameReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaSystemSystemName}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaSystemSystemName")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaSystemSystemNetworkInstanceReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaSystemSystemNetworkInstance}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaSystemSystemNetworkInstance")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaSystemSystemNtpReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaSystemSystemNtp}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaSystemSystemNtp")
		os.Exit(1)
	}

	if err := (&controllers.K8sSrlNokiaTunnelInterfacesTunnelInterfaceReconciler{
		Client: mgr.GetClient(),
		Server: &natsServer,
		Log:    ctrl.Log.WithName("Controller").WithName("{K8sSrlNokiaTunnelInterfacesTunnelInterface}"),
		Scheme: mgr.GetScheme(),
		Ctx:    ctx,
	}).SetupWithManager(ctx, mgr, concurrency(srlConcurrency)); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "K8sSrlNokiaTunnelInterfacesTunnelInterface")
		os.Exit(1)
	}

}

func concurrency(c int) controller.Options {
	return controller.Options{MaxConcurrentReconciles: c}
}
