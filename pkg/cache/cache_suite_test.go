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

package cache_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rmanijacob/controller-runtime/pkg/envtest"
	"github.com/rmanijacob/controller-runtime/pkg/envtest/printer"
	logf "github.com/rmanijacob/controller-runtime/pkg/log"
	"github.com/rmanijacob/controller-runtime/pkg/log/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestSource(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteName := "Cache Suite"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}

var testenv *envtest.Environment
var cfg *rest.Config
var clientset *kubernetes.Clientset

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	testenv = &envtest.Environment{}

	var err error
	cfg, err = testenv.Start()
	Expect(err).NotTo(HaveOccurred())

	clientset, err = kubernetes.NewForConfig(cfg)
	Expect(err).NotTo(HaveOccurred())
}, 60)

var _ = AfterSuite(func() {
	Expect(testenv.Stop()).To(Succeed())
})
