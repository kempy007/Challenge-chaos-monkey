# Challenge-chaos-monkey
Challenge is to build a chaos monkey service as per the [Challenge.md](./Challenge.md) I moved how I defined how I would go about this task to this readme below.

## Direction

Operator-sdk satisfies lang=golang, dockerfile, operator manifests and deployment options.

- [x] Use Operator-sdk to create CRD to configure a namespace else use all NS and to set a schedule (cron type). 
- [x] Operator watches for this resource and creates/destroys a chaos Pod(perhaps use go concurrency) as per CRUD Operation with CRD options. 
- [x] Filter out NS kube-system, monitoring
- [x] Create manifests to setup workloads NS and introduce deployments and static pods.
- [x] Create manifest for CRD for operator for workloads demo.
- [x] Provide instructions in readme of howto deploy this operator and setup the demo.

I did have starcoder enabled in my IDE, and I used starchat to ask a question to get unstuck as I have not been on this bicycle for a few years. I feel I need to be on the bicycle at least weekly to remain speedy,current and fresh.

I kept my dev notes in the [BuildNotes.md](./BuildNotes.md) file.

I have overun on my time but met the core requirements at the expense of writing the unittests thus functions maybe leaning towards monolith.

## setup_and_prerequisites

### MacOS

- brew install operator-sdk kubebuilder kustomize

## quick steps

This will run the operator in developer mode from your IDE, external to your cluster.

- git clone git@github.com:kempy007/Challenge-chaos-monkey.git
- cd Challenge-chaos-monkey
- kubectl apply -f demo-manifests/workloads.yaml
- kubectl apply -f demo-manifests/crd.ChaosMonkey.yaml
- make install run

Watch pods be deleted in workloads namespace. You write your own version of this CRD demo-manifests/crd.ChaosMonkey.yaml to config alternate schedule, this is how the service is will be configured in production.


## OLM Bundle install



## My Versions

operator-sdk version: "v1.29.0", commit: "78c564319585c0c348d1d7d9bbfeed1098fab006", kubernetes version: "v1.26.0", go version: "go1.20.4", GOOS: "darwin", GOARCH: "arm64"

WARNING: This version information is deprecated and will be replaced with the output from kubectl version --short.  Use --output=yaml|json to get the full version.
Client Version: version.Info{Major:"1", Minor:"27", GitVersion:"v1.27.1", GitCommit:"4c9411232e10168d7b050c49a1b59f6df9d7ea4b", GitTreeState:"clean", BuildDate:"2023-04-14T13:14:41Z", GoVersion:"go1.20.3", Compiler:"gc", Platform:"darwin/arm64"}
Kustomize Version: v5.0.1
Server Version: version.Info{Major:"1", Minor:"26", GitVersion:"v1.26.4+k3s1", GitCommit:"8d0255af07e95b841952563253d27b0d10bd72f0", GitTreeState:"clean", BuildDate:"2023-04-20T00:33:18Z", GoVersion:"go1.19.8", Compiler:"gc", Platform:"linux/amd64"}