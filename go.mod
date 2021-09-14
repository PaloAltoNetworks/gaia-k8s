module go.aporeto.io/gaia-k8s

go 1.16

require (
	github.com/golang/mock v1.5.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	go.aporeto.io/elemental v1.100.1-0.20210706184354-966eab3720af // indirect
	go.aporeto.io/gaia v1.102.0
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	k8s.io/code-generator v0.22.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
)

//replace go.aporeto.io/gaia v1.102.0 => ../gaia
