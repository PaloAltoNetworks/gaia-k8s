module go.aporeto.io/gaia-k8s

go 1.16

require (
	github.com/go-openapi/spec v0.20.4
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	go.aporeto.io/gaia v1.94.1-0.20210917202132-43da66fb7de7
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/code-generator v0.22.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
)

replace go.aporeto.io/gaia => ../gaia
