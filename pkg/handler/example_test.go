/*
Copyright 2018 The Kubernetes Authors.

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

package handler_test

import (
	"github.com/rmanijacob/controller-runtime/pkg/client"
	"github.com/rmanijacob/controller-runtime/pkg/controller"
	"github.com/rmanijacob/controller-runtime/pkg/event"
	"github.com/rmanijacob/controller-runtime/pkg/handler"
	"github.com/rmanijacob/controller-runtime/pkg/reconcile"
	"github.com/rmanijacob/controller-runtime/pkg/source"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

var c controller.Controller

// This example watches Pods and enqueues Requests with the Name and Namespace of the Pod from
// the Event (i.e. change caused by a Create, Update, Delete).
func ExampleEnqueueRequestForObject() {
	// controller is a controller.controller
	err := c.Watch(
		&source.Kind{Type: &corev1.Pod{}},
		&handler.EnqueueRequestForObject{},
	)
	if err != nil {
		// handle it
	}
}

// This example watches ReplicaSets and enqueues a Request containing the Name and Namespace of the
// owning (direct) Deployment responsible for the creation of the ReplicaSet.
func ExampleEnqueueRequestForOwner() {
	// controller is a controller.controller
	err := c.Watch(
		&source.Kind{Type: &appsv1.ReplicaSet{}},
		&handler.EnqueueRequestForOwner{
			OwnerType:    &appsv1.Deployment{},
			IsController: true,
		},
	)
	if err != nil {
		// handle it
	}
}

// This example watches Deployments and enqueues a Request contain the Name and Namespace of different
// objects (of Type: MyKind) using a mapping function defined by the user.
func ExampleEnqueueRequestsFromMapFunc() {
	// controller is a controller.controller
	err := c.Watch(
		&source.Kind{Type: &appsv1.Deployment{}},
		handler.EnqueueRequestsFromMapFunc(func(a client.Object) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      a.GetName() + "-1",
					Namespace: a.GetNamespace(),
				}},
				{NamespacedName: types.NamespacedName{
					Name:      a.GetName() + "-2",
					Namespace: a.GetNamespace(),
				}},
			}
		}),
	)
	if err != nil {
		// handle it
	}
}

// This example implements handler.EnqueueRequestForObject.
func ExampleFuncs() {
	// controller is a controller.controller
	err := c.Watch(
		&source.Kind{Type: &corev1.Pod{}},
		handler.Funcs{
			CreateFunc: func(e event.CreateEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Object.GetName(),
					Namespace: e.Object.GetNamespace(),
				}})
			},
			UpdateFunc: func(e event.UpdateEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.ObjectNew.GetName(),
					Namespace: e.ObjectNew.GetNamespace(),
				}})
			},
			DeleteFunc: func(e event.DeleteEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Object.GetName(),
					Namespace: e.Object.GetNamespace(),
				}})
			},
			GenericFunc: func(e event.GenericEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Object.GetName(),
					Namespace: e.Object.GetNamespace(),
				}})
			},
		},
	)
	if err != nil {
		// handle it
	}
}
