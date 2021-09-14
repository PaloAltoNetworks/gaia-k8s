module go.aporeto.io/gaia-k8s

go 1.16

require (
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cobra v1.1.3
	go.aporeto.io/elemental v1.100.1-0.20210706184354-966eab3720af
	go.aporeto.io/gaia v1.102.0
	go.aporeto.io/manipulate v1.124.0
	go.aporeto.io/midgard-lib v1.72.0
	go.aporeto.io/tg v1.34.1-0.20210528201128-159c302ba155
	k8s.io/apimachinery v0.22.0
	k8s.io/apiserver v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/code-generator v0.22.0
	k8s.io/component-base v0.22.0
	k8s.io/klog/v2 v2.9.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
)

//replace go.aporeto.io/gaia v1.102.0 => ../gaia
