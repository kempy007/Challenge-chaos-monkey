# Notes

operator-sdk init --domain cypherpunk.io --repo github.com/kempy007/Challenge-chaos-monkey --plugins go.kubebuilder.io/v4-alpha --skip-go-version-check

operator-sdk create api --group=core --version=v1 --kind=ChaosMonkey --controller=true --resource=true


---
# prompt engineering - https://huggingface.co/spaces/HuggingFaceH4/starchat-playground

using golang and the operator-sdk give a code example of how I can access resources outside of my CRD.

Sure! Here's an example Go program that uses the Operator SDK to watch for changes to a CustomResourceDefinition (CRD) called "example.com/v1alpha1", which represents a simple resource with a string field named "message". The program then prints any updates to this field to standard output:

```
    package main
    
    import (
            "context"
            "flag"
            "fmt"
            operatorsdk "github.com/operator-framework/operator-sdk/pkg/sdk"
            metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
            "k8s.io/client-go/kubernetes"
            "k8s.io/client-go/rest"
    )
    
    type Example struct {
            metav1.TypeMeta   `json:",inline"`
            metav1.ObjectMeta `json:"metadata,omitempty"`
            Message          string   `json:"message,omitempty"`
    }
    
    func main() {
            namespace := flag.String("namespace", "", "Namespace where the CRD is defined")
            crdName := flag.String("crd-name", "", "Name of the CRD")
            flag.Parse()
    
            // Get the Kubernetes clientset
            config, err := rest.InClusterConfig()
            if err!= nil {
                    panic(err.Error())
            }
            clientset, err := kubernetes.NewForConfig(config)
            if err!= nil {
                    panic(err.Error())
            }
            
            // Watch for changes to the specified CRD
            watchlist := &source.Kind{Type: &Example{}}
            informerFactory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace(*namespace))
            controller := controller.NewController(&controller.ControllerOptions{
                    ClientSet:     clientset,
                    WatchList:     watchlist,
                    ResyncPeriod:   5 * time.Minute,
                    QueueSize:      100})
            informerFactory.Start(ctx.Done())
            defer informerFactory.Stop()
            controller.Run(ctx.Done())
    }
```

This program assumes you are running it inside a pod in the same cluster as your Kubernetes API server, so it will use the InClusterConfig function to get the necessary credentials. You'll need to replace "example.com/v1alpha1" with the actual group/version/kind of your own CRD. Note that this program only watches for changes to the entire CRD; if you want to watch specific fields within the CRD, you would need to modify the program accordingly.

Please Give another code example as above but also listing pods.

Here is an updated version of the previous code sample that lists all Pods in the current Namespace when the program starts up:

```
package main

import (
    "context"
    "flag"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/cache"

    operatorsdk "github.com/operator-framework/operator-sdk/pkg/sdk"
)

const (
    // Namespaces to target
    allNamespaces = ""
    // GroupVersion used for custom resource definition
    groupVersion = "example.com/v1alpha1"
    // Kind used for custom resource definition
    resourceKind = "Example"
)

var (
    kubeClient  *kubernetes.Clientset
    stopCh      chan struct{}
    crdsLister  cache.Indexer
    podsLister  corev1.PodInterface
    namespaces []string
)

func init() {
    flag.StringVar(&namespaces, "namespaces", namespaces, fmt.Sprintf("Comma separated list of namespaced to target (%q for all)", allNamespaces))
}

func main() {
    // Get the Kubernetes clientset
    config, err := rest.InClusterConfig()
    if err!= nil {
            panic(err.Error())
    }
    kubeClient, err = kubernetes.NewForConfig(config)
    if err!= nil {
            panic(err.Error())
    }

    // Create shared index informers for CRDs and Pods
    crdsInformerFactory := informers.NewSharedInformerFactory(kubeClient, 0)
    podInformer := crdsInformerFactory.Core().V1().Pods()

    // Start the informer factories to begin populating the caches
    stopCh = make(chan struct{})
    defer close(stopCh)
    crdsInformerFactory.Start(stopCh)
    podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
            AddFunc: func(obj interface{}) {
                    switch t := obj.(type) {
                    case *corev1.Pod:
                            processPod(t)
                    default:
                    }
            },
            UpdateFunc: func(_, newObj interface{}) {
                    switch t := newObj.(type) {
                    case *corev1.Pod:
                            processPod(t)
                    default:
                    }
            },
            DeleteFunc: func(obj interface{}) {
            },
    })

    // Wait for the caches to be synced before starting processing
    crdsInformerFactory.WaitForCacheSync(stopCh)

    // Watch for changes to the specified CRD
    watchlist := &source.Kind{Type: &Example{}}
    controller := controller.NewController(&controller.ControllerOptions{
            ClientSet: kubeClient,
            WatchList: watchlist,
            ResyncPeriod: 5 * time.Minute,
            QueueSize: 100,
    })
    controller.Run(stopCh)
}

func processPod(p *corev1.Pod) {
    for _, container := range p.Spec.Containers {
            fmt.Printf("%s:%d - %s\n", p.Status.PodIP, container.Ports[0].ContainerPort, p.Name)
    }
}
```

--- 
# Log

Started 16/06/23 ~18:00

- [x] Create a new project(repo)
- [x] Created strategy at end of challenge.md

Updated 16/06/23 ~18:30 

- [x] Create a new API
- [x] Create a new Controller
- [x] Create a new CRD
- [x] Create a new CR
- [x] Create a new RBAC
- [x] Create manifest for CRD
- [x] Create manifests for demo workloads

Updated 16/06/23 ~19:00 - ~22:00

- Getting interrupted alot broken flow
- [x] Research How to use Go concurrency
    - [] Research Go routines
    - [] Research Go channels
    - [] Research Go select
- [x] Re-reading my old projects for hints

Updated 17/06/23 ~11:00

- [x] Research How to access resources outside my CRD in operator-sdk
- [x] Prompt Engineered
- [x] kubeconfig working 11:15