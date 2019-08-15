Look ma no YAML

# Tl;dr?

This is a very simple tool to generate kubernetes YML via `go generate`.

It piggy-backs on top of [ko](http://github.com/google/ko) to generate Docker images
from Go code, and [knative serving](http://github.com/knative/serving) to easily
run them on Kubernetes.

Basically, no more YAML wrangling, just `go generate` and `kubectl apply`.

# How to Use?

## Pre-requisites

- **Ko:** from [github.com/google/ko](http://github.com/google/ko).

## Install

~~~~
go install github.com/julz/kgen/cmd/kgen
~~~~

## Quick-Start

1. Add a `//go:generate kgen ksvc` to your main package(s)
   * See [examples/simple/main.go](examples/simple/main.go#L10) for an example
1. Run `go generate ./...`
1. Your YML is ready! Just `ko apply -f build/yml`. (See the [ko
   docs](http://github.com/google/ko) for more
   about this step).

No more manual YML wrangling!.
