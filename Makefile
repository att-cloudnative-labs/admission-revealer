# Image URL to use all building/pushing image targets
IMG ?= admission-revealer:v3

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: webhook

# Build webhook binary
webhook: fmt vet
	go build -o webhook main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: fmt vet
	go run ./main.go

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy:
	kubectl apply -f config/

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}
