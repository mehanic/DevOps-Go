# Golang Lessons (Kubernetes Operators)

## Introduction

1. History of why Golang emerged, what problems was it created to solve
   - Resources/Links:
     - https://medium.com/geekculture/learn-go-part-1-the-beginning-723746f2e8b0
2. What projects currently use Golang? (emphasis on CNCF projects with Kubernetes being the main project)
3. What other languages compare to Golang? (when to choose one over the other)
   - Subpoints:
     - Go is simple to understand (not a lot of language features)
     - Go has built-in concurrency with the runtime
     - Lots of CloudNative projects are written in Go

## Let's Begin!

1. Installing Golang
   - Subpoints:
     - Explanation of the various environment variables (`GOPATH`, `GO111MODULE`, `GOBIN`)
     - Basic usage of the `go` command with some subcommands
2. "Hello, World!"
   - Subpoints:
     - Explanation of `package`s and the structure of them
     - Explanation of what is actually happening once you run the program with `go run` or `go build`
3. LearnByExample Go (Basic)
   - Subpoints:
     - Types
     - Constants/Variables
     - Functions
     - Arrays/Slices
     - Loops
     - Maps
     - Conditionals (`if/else`, `switch`)
     - Pointers
     - Structs
     - Methods (Pointer vs. Value)
     - Interfaces

## Keep (Go)ing!

1. Go projects
   - Subpoints:
     - Go modules
     - Constructing a Go project
     - Referencing of [this](https://github.com/golang-standards/project-layout) to show how Go projects are organized
     - HTTP Clients and Server
2. Write a basic HTTP server
   - Subpoints:
     - HTTP server with two endpoints (`GET`, `POST`). JSON decoding
3. Use Golang Kubernetes client library
   - Subpoints:
     - Explanation of what `kubectl` does underneath the hood
       - Makes authenticated? API requests to the kube api server
     - Install `k8s` Golang library
     - Use client to do emulate basic "kubectl" like requests
       - `get` pods
       - `create` deployment
4. Running Golang project elsewhere
   - Subpoints:
     - Dockerizing Golang project
     - Considerations for running k8s aware Golang project in k8s cluster
       - Have to use `InClusterConfig`

## Kubernetes Controllers/Operators

1. Dynamic Admission Controller
   - Subpoints:
     - Introducing basics [here]()
     - Utility of controllers
       - Validating Resources
       - Mutating Resources based on whatever logic you want (side effects)
     - Built in controllers (deployment, cronjob, etc) [here](https://github.com/kubernetes/kubernetes/tree/master/pkg/controller)
     - Third Party Controllers:
       - [aws-lb-controller](https://github.com/kubernetes-sigs/aws-load-balancer-controller)
       - [ingress-nginx](https://github.com/kubernetes/ingress-nginx)
     - Create custom admission controller
       - Validating Webhook (k8s deployment check for presense of environment variable)
       - Mutating Webhook (inject environment variable for every k8s deployment)
2. What is a Kubernetes operator?
   - Subpoints:
     - Introducing basics [here](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
     - Why create operators?
       - Help to extend functionality to interesting use cases
       - Usually used to interact with a 3rd party service and control resource lifecycle on that third party
     - CRDs:
       - [ArgoCD](https://github.com/argoproj/argo-cd)
       - [Crossplane](https://github.com/crossplane/crossplane)
       - [Tekton](https://tekton.dev/)
3. Operator SDK
   - Subpoints:
     - K8s operators are complicated to create from scratch
     - [Operator SDK](https://sdk.operatorframework.io/) makes it easier
     - Explanation of generated code
     - [Tutorial](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/) on how to use
4. Creation of custom operator(s)
   - Subpoints:
     - GitHub repository operator
     - S3 bucket operator
