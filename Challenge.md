# Scenario

Kubernetes clusters drive towards a declared state. Since Kubernetes environments are
typically very dynamic, Pods are intended to be ephemeral, and could be destroyed at any
point.

Since this can occur naturally, some companies test the resilience of workloads in advance by
doing this on purpose.

We would like you to write a program which runs inside the cluster, interacts with the kube apiserver, and deletes one pod at random in a particular namespace on a schedule. The program could be known as ‘PodChaosMonkey’. 

## Requirements
- Source code implementing the behaviour described, in one of the following languages:
  - Golang, Python, Node (typescript optional), C# or Java
- Dockerfile to build an image for the program
- Kubernetes manifests to run this in a local cluster
- A dummy workload running in the ‘workloads’ namespace as a target, public nginx image or similar is fine.


## Notes

Please don’t spend longer than 3 hours on this.

We mark these tests on areas like: 
 - functionality
 - understanding of Kubernetes
 - unit tests
 - code quality
 - readability 
 - the README

We don’t expect this to be production-ready code. Equally, if you have time to spare, consider whether all pods are suitable candidates for this testing, and how users of this program would configure it.

Be careful to target your implementation against the ‘workloads’ namespace or you may need to rebuild your cluster.

You can submit your code either by sharing a link to a Github repository or by emailing us a zip archive of your codebase.
