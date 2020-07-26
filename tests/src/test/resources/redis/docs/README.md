## Prerequisites

You need a Kubernetes cluster up and running and Persistent Storage available with a default `Storage Class` defined.

## Getting Started

Deploy the `Operator` using the following command:

`kubectl kudo install redis`

It deploys a Redis Cluster composed of 6 instances. There are 3 masters and 1 replica per master.
