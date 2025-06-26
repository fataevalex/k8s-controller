# Kubernetes Controller

It's a one of implementation of the Golang Kubernetes Controller course from FWDays.

## About

This project follows [the step-by-step tutorial](https://github.com/den-vasyliev/k8s-controller-tutorial-ref) for building production-grade Kubernetes controllers in Go. Each step is implemented as a separate commit/branch with detailed explanations.

**Course**: [Crash Course: Kubernetes controllers](https://fwdays.com/event/kubernetes-controllers-course)
**Instructors**: @den-vasyliev (Principal SRE), @Alex0M (Senior Platform Engineer)

## Prerequisites

- [Go](https://golang.org/dl/) 1.24 or newer
- [cobra-cli](https://github.com/spf13/cobra-cli) installed:
  ```sh
  go install github.com/spf13/cobra-cli@latest
  ```
  or for mac
  ```sh
  brew install cobra-cli
  ```

## Quick Start

### One-Command Setup

Get a complete Kubernetes development environment running in seconds:

```
# Clone the repository
git clone https://github.com/Searge/k8s-controller.git
cd k8s-controller
```

## Steps
- [x] [Cobra CLI](docs/01-cobra-cli/README.md)
- [x] [Zerolog logging](docs/02-zerolog-logging/README.md)
- [x] [Pflag, implementing POSIX/GNU-style --flags](docs/03-cobra-cli-pflag/README.md)
- [x] [Fast HTTP](docs/04-fast-http-server/README.md)
- [x] [Makefile, CI](docs/05-github-ci/README.md)
- [x] [List Kubernetes Deployments with client-go](docs/06-list-deployments/README.md)
- [x] [Deployment Informer with client-go](docs/07-Deployment-Informer/README.md)
- [x] [Deployment Informer with api-handler](docs/08-api-handler/README.md)
- [x] [Controller runtime](docs/09-controller-runtime/README.md)
- [x] [Leader Election and Metrics for Controller Manager](docs/10-leader-election/README.md)