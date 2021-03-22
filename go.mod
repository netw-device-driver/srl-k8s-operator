module github.com/srl-wim/srl-k8s-operator

go 1.15

require (
	github.com/go-logr/logr v0.3.0
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/metal3-io/baremetal-operator v0.0.0-20210310165305-c74264a6ddcb
	github.com/netw-device-driver/netw-device-controller v0.1.4
	github.com/netw-device-driver/netwdevpb v0.1.15
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/stoewer/go-strcase v1.2.0
	golang.org/x/net v0.0.0-20201216054612-986b41b23924 // indirect
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d // indirect
	google.golang.org/grpc v1.34.0
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
)
