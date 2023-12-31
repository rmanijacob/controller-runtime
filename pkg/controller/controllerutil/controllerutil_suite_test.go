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

package controllerutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rmanijacob/controller-runtime/pkg/client"
	"github.com/rmanijacob/controller-runtime/pkg/envtest"
	"github.com/rmanijacob/controller-runtime/pkg/envtest/printer"
	"k8s.io/client-go/rest"
)

func TestControllerutil(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteName := "Controllerutil Suite"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}

var testenv *envtest.Environment
var cfg *rest.Config
var c client.Client

var _ = BeforeSuite(func() {
	var err error

	testenv = &envtest.Environment{}

	cfg, err = testenv.Start()
	Expect(err).NotTo(HaveOccurred())

	c, err = client.New(cfg, client.Options{})
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(testenv.Stop()).To(Succeed())
})
