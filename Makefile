export MAKEFLAGS += --warn-undefined-variables
export SHELL := /bin/bash -o pipefail

export GO111MODULE = on
export GOPRIVATE = go.aporeto.io,github.com/aporeto-inc

default: checkfmt lint test

golangci-lint:
ifeq (, $(shell which golangci-lint))
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
GOLANGCI_LINT=$(shell which golangci-lint)
else
GOLANGCI_LINT=$(shell which golangci-lint)
endif

lint: golangci-lint
	$(GOLANGCI_LINT) run \
		--disable-all \
		--deadline=5m \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=golint \
		--enable=unused \
		--enable=structcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=typecheck \
		--enable=unparam \
		./...

checkfmt:
ifneq ($(shell { goimports -l ./pkg/apis/networkprismacloudio/ ; goimports -l ./pkg/generated/ ; }),)
	@echo Please run \`make fmt\`
	exit 1
endif
	@echo Success

test:
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.cov

##########################
### non-public targets ###
##########################

fmt:
	goimports -w ./pkg/apis/networkprismacloudio/
	goimports -w ./pkg/generated/

.PHONY:codegen
codegen: get-codegen 
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind NetworkRuleSetPolicy
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind ExternalNetwork 
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind NetworkRuleSetPolicy --namespaced true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind ExternalNetwork --namespaced true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind ProcessingUnit --namespaced true --readonly true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind PUTrafficAction --namespaced true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind ProcessingUnit --readonly true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind PUTrafficAction
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind Enforcer --readonly true
	k8s-api-codegen create api --group networkprismacloudio --version v1 --kind EnforcerProfile 
	make fmt

regenerate: 
	rm -rf ./pkg/apis/networkprismacloudio/
	rm -rf ./pkg/generated/
	mkdir ./pkg/generated/
	make codegen
	make generate

get-codegen: ## Download get-codegen.
	go get go.aporeto.io/k8s-api-codegen@latest

generate: 
	./hack/update_codegen.sh

#generate-conversion: conversion-gen
#	$(CONVERSION_GEN) -v 20 -O zz_generated.conversion -i ./pkg/apis/networkprismacloudio/v2  -h ./hack/boilerplate.go.txt

download_dependencies: 
	make deepcopy-gen
	make conversion-gen
	make defaulter-gen
	make client-gen
	make lister-gen
	make informer-gen
	make openapi-gen

deepcopy-gen:
ifeq (, $(shell which deepcopy-gen))
	go get k8s.io/code-generator/cmd/deepcopy-gen@latest
DEEPCOPY_GEN=$(shell which deepcopy-gen)
else
DEEPCOPY_GEN=$(shell which deepcopy-gen)
endif

conversion-gen:
ifeq (, $(shell which conversion-gen))
	go get k8s.io/code-generator/cmd/conversion-gen@latest
CONVERSION_GEN=$(shell which conversion-gen)
else
CONVERSION_GEN=$(shell which conversion-gen)
endif

defaulter-gen:
ifeq (, $(shell which defaulter-gen))
	go get k8s.io/code-generator/cmd/defaulter-gen@latest
DEFAULTER_GEN=$(shell which defaulter-gen)
else
DEFAULTER_GEN=$(shell which defaulter-gen)
endif

client-gen:
ifeq (, $(shell which client-gen))
	go get k8s.io/code-generator/cmd/client-gen@latest
CLIENT_GEN=$(shell which client-gen)
else
CLIENT_GEN=$(shell which client-gen)
endif

lister-gen:
ifeq (, $(shell which lister-gen))
	go get k8s.io/code-generator/cmd/lister-gen@latest
LISTER_GEN=$(shell which lister-gen)
else
LISTER_GEN=$(shell which lister-gen)
endif

informer-gen:
ifeq (, $(shell which informer-gen))
	go get k8s.io/code-generator/cmd/informer-gen@latest
INFORMER_GEN=$(shell which informer-gen)
else
INFORMER_GEN=$(shell which informer-gen)
endif

openapi-gen:
ifeq (, $(shell which openapi-gen))
	go get k8s.io/code-generator/cmd/openapi-gen@latest
OPENAPI_GEN=$(shell which openapi-gen)
else
OPENAPI_GEN=$(shell which openapi-gen)
endif
