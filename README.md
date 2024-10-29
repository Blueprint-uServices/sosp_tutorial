# Blueprint SOSP 2024 Tutorial

This tutorial provides a hands-on introduction on how systems researchers can use Blueprint to accelerate their microservice research.

The tutorial will be divided into 3 parts &mdash; (i) Running Blueprint applications out of the box; (ii) Modifying existing applications and taking measurements; (iii) Reproducing a metastability failure in an application.

## Pre-Requisites

The tutorial only requires a functional laptop. No additional resources are required. Prior to attending the tutorial, please install the following dependencies for a smooth experience.

### Dependencies

+ go 1.22+ : [Install](https://go.dev/doc/install)
+ docker : [Install](https://docs.docker.com/engine/install/)
+ (Optional) gRPC for Go: [Install](https://grpc.io/docs/languages/go/quickstart/)
+ (Optional) Thrift: [Download](https://thrift.apache.org/download), [Dependencies](https://thrift.apache.org/docs/install/debian.html)


## Detailed Instructions

+ [Part 1](./Part1.md): Running existing Blueprint apps
+ [Part 2](./Part2.md): Re-configuring Blueprint apps and collecting data
+ [Part 3](./Part3.md): Reproducing metastability failures
+ [Bonus Part](./PartBonus.md): Adding a new Blueprint plugin.