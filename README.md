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
```
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


### Fist stage build cobra-cli sample saaplication
 Create template of go application wich using cobra-cli framework
 ```
 #switch to feature barch 'step1-cobra-cli'
 git checkout  feature/step1-cobra-cli
 ```
**Initialize Go module (if not already):**
```
 go mod init github.com/fataevalex/k8s-controller
 ```
**Initialize Cobra:**
 ```sh
   cobra-cli init
   ```
**Build your CLI:**
   ```sh
   go build -o controller
   ```

**Run your CLI (shows help by default):**
   ```sh
   ./controller --help

## Project Structure

- `cmd/` — Contains your CLI commands.
- `main.go` — Entry point for your application.
- `cmd/go_basic.go`: Implements the command and struct logic
- `cmd/go_basic_test.go`: Unit tests for the struct methods 

This directory contains the `go_basic.go` file, which demonstrates basic usage of Go structs and methods within a Cobra CLI command.
