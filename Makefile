APP = k8s-controller
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_FLAGS = -v -o $(APP) -ldflags "-X=github.com/den-vasyliev/$(APP)/cmd.appVersion=$(VERSION)"
ENVTEST ?= $(LOCALBIN)/setup-envtest
ENVTEST_VERSION ?= latest
LOCALBIN ?= $(shell pwd)/bin

.PHONY: all build test test-coverage run docker-build clean envtest

all: build

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
ENVTEST ?= $(LOCALBIN)/setup-envtest

## Tool Versions
ENVTEST_VERSION ?= release-0.19

format:
	gofmt -s -w ./

lint:
	golint

envtest: $(ENVTEST) ## Download setup-envtest locally if necessary.
$(ENVTEST): $(LOCALBIN)
	$(call go-install-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest,$(ENVTEST_VERSION))


build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_FLAGS) main.go

test: envtest
	go install gotest.tools/gotestsum@latest
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use --bin-dir $(LOCALBIN) -p path)" gotestsum --junitfile report.xml --format testname ./... ${TEST_ARGS}

test-coverage: envtest
	go install github.com/boumenot/gocover-cobertura@latest
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use --bin-dir $(LOCALBIN) -p path)" go test -coverprofile=coverage.out -covermode=count ./...
	go tool cover -func=coverage.out
	gocover-cobertura < coverage.out > coverage.xml

test-integration:
	 go test -v ./cmd -run TestDeploymentIntegration

run:
	go run main.go

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t $(APP):latest .

clean:
	rm -f $(APP)