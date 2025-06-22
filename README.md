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
#### - step 1. Init cobra-cli application  [README1.md](README1.md)
#### - step 2. Add zerolog logging  [README1.md](README2.md)
#### - step 3. Add --log-level command line parameter [README3.md](README3.md)
#### - step 4. Add simple fast http server  [README4.md](README4.md)
#### - step 5. Add github CI [README5.md](README5.md)
