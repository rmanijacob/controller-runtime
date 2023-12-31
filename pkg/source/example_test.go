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

package source_test

import (
	"github.com/rmanijacob/controller-runtime/pkg/controller"
	"github.com/rmanijacob/controller-runtime/pkg/event"
	"github.com/rmanijacob/controller-runtime/pkg/handler"
	"github.com/rmanijacob/controller-runtime/pkg/source"
	corev1 "k8s.io/api/core/v1"
)

var ctrl controller.Controller

// This example Watches for Pod Events (e.g. Create / Update / Delete) and enqueues a reconcile.Request
// with the Name and Namespace of the Pod.
func ExampleKind() {
	err := ctrl.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		// handle it
	}
}

// This example reads GenericEvents from a channel and enqueues a reconcile.Request containing the Name and Namespace
// provided by the event.
func ExampleChannel() {
	events := make(chan event.GenericEvent)

	err := ctrl.Watch(
		&source.Channel{Source: events},
		&handler.EnqueueRequestForObject{},
	)
	if err != nil {
		// handle it
	}
}
