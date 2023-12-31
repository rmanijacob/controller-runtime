/*
Copyright 2020 The Kubernetes Authors.

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

package config_test

import (
	"fmt"
	"os"

	"github.com/rmanijacob/controller-runtime/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/rmanijacob/controller-runtime/examples/configfile/custom/v1alpha1"
)

var scheme = runtime.NewScheme()

func init() {
	_ = v1alpha1.AddToScheme(scheme)
}

// This example will load a file using Complete with only
// defaults set.
func ExampleFile() {
	// This will load a config file from ./config.yaml
	loader := config.File()
	if _, err := loader.Complete(); err != nil {
		fmt.Println("failed to load config")
		os.Exit(1)
	}
}

// This example will load the file from a custom path.
func ExampleDeferredFileLoader_atPath() {
	loader := config.File().AtPath("/var/run/controller-runtime/config.yaml")
	if _, err := loader.Complete(); err != nil {
		fmt.Println("failed to load config")
		os.Exit(1)
	}
}

// This example sets up loader with a custom scheme.
func ExampleDeferredFileLoader_injectScheme() {
	loader := config.File()
	err := loader.InjectScheme(scheme)
	if err != nil {
		fmt.Println("failed to inject scheme")
		os.Exit(1)
	}

	_, err = loader.Complete()
	if err != nil {
		fmt.Println("failed to load config")
		os.Exit(1)
	}
}

// This example sets up the loader with a custom scheme and custom type.
func ExampleDeferredFileLoader_ofKind() {
	loader := config.File().OfKind(&v1alpha1.CustomControllerManagerConfiguration{})
	err := loader.InjectScheme(scheme)
	if err != nil {
		fmt.Println("failed to inject scheme")
		os.Exit(1)
	}
	_, err = loader.Complete()
	if err != nil {
		fmt.Println("failed to load config")
		os.Exit(1)
	}
}
