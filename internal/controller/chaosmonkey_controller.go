/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "github.com/kempy007/Challenge-chaos-monkey/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ChaosMonkeyReconciler reconciles a ChaosMonkey object
type ChaosMonkeyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core.cypherpunk.io,resources=chaosmonkeys,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.cypherpunk.io,resources=chaosmonkeys/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.cypherpunk.io,resources=chaosmonkeys/finalizers,verbs=update

// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ChaosMonkey object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ChaosMonkeyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	//crd := corev1.ChaosMonkeySpec{} 	//ERROR: does not implement client.Object (missing method DeepCopyObject)
	var crd corev1.ChaosMonkey

	err := r.Client.Get(ctx, req.NamespacedName, &crd)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	NS := crd.Spec.Namespace
	SCHED := crd.Spec.Schedule
	TIMER := crd.Spec.Timer
	STRTIM := fmt.Sprintf("%d", TIMER)

	if NS == "" {
		log.Log.Info("CRD has no namespace set")
	}
	if SCHED == "" {
		log.Log.Info("CRD has no schedule set")
	}
	if TIMER == 0 {
		log.Log.Info("CRD has no timer set")
	}

	log.Log.Info("ChaosMonkey named: " + crd.Name + " has namespace: " + NS + ", schedule: " + SCHED + ", timer: " + STRTIM)

	//TODO: Add Channel Routine to execute scheduled actions
	runTimer(NS, TIMER)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ChaosMonkeyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.ChaosMonkey{}).
		Complete(r)
}

// Timer Routine from Namespace, kill random pod in namespace
func runTimer(NS string, T int) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Log.Info(err.Error())
		config, err = ctrl.GetConfig()
		if err != nil {
			log.Log.Info(err.Error())
		}
	}
	if config == nil {
		log.Log.Info("No config")
		return
	}
	// Get all pods in namespace
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Log.Info(err.Error())
	}
	if clientSet == nil {
		log.Log.Info("No clientSet")
		return
	}
	podClient := clientSet.CoreV1().Pods(NS)
	pods, err := podClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Log.Info(err.Error())
	}
	if pods == nil {
		log.Log.Info("No pods")
		return
	}

	// log each pods.Items
	for _, pod := range pods.Items {
		log.Log.Info(pod.Name)
	}

	// // Kill random pod
	// randPod := pods.Items[0].Name
	// log.Log.Info("Killing pod: " + randPod)
	// err = podClient.Delete(context.TODO(), randPod, metav1.DeleteOptions{})
	// if err != nil {
	// 	//log.Log.Info(err.Error())
	// }
}
