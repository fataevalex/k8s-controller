APP = k8s-controller
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_FLAGS = -v -o $(APP) -ldflags "-X=github.com/den-vasyliev/$(APP)/cmd.appVersion=$(VERSION)"

.PHONY: all build test run docker-build clean

all: build

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_FLAGS) main.go

test:
	go test ./...

test-integration:
	 go test -v ./cmd -run TestDeploymentIntegration

run:
	go run main.go

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t $(APP):latest .

clean:
	rm -f $(APP)