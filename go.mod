module go.aporeto.io/gaia-k8s

go 1.16

require (
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	go.aporeto.io/gaia v1.102.0
	k8s.io/apimachinery v0.22.0
	k8s.io/code-generator v0.22.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
)

replace go.aporeto.io/gaia v1.102.0 => go.aporeto.io/gaia v1.94.1-0.20210915121255-371799927b54
