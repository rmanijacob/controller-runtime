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

package manager

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rmanijacob/controller-runtime/pkg/envtest"
	"github.com/rmanijacob/controller-runtime/pkg/envtest/printer"
	logf "github.com/rmanijacob/controller-runtime/pkg/log"
	"github.com/rmanijacob/controller-runtime/pkg/log/zap"
	"github.com/rmanijacob/controller-runtime/pkg/metrics"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestSource(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteName := "Manager Suite"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}

var testenv *envtest.Environment
var cfg *rest.Config
var clientset *kubernetes.Clientset

// clientTransport is used to force-close keep-alives in tests that check for leaks.
var clientTransport *http.Transport

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	testenv = &envtest.Environment{}

	var err error
	cfg, err = testenv.Start()
	Expect(err).NotTo(HaveOccurred())

	cfg.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		// NB(directxman12): we can't set Transport *and* use TLS options,
		// so we grab the transport right after it gets created so that we can
		// type-assert on it (hopefully)?
		// hopefully this doesn't break 🤞
		transport, isTransport := rt.(*http.Transport)
		if !isTransport {
			panic(fmt.Sprintf("wasn't able to grab underlying transport from REST client's RoundTripper, can't figure out how to close keep-alives: expected an *http.Transport, got %#v", rt))
		}
		clientTransport = transport
		return rt
	}

	clientset, err = kubernetes.NewForConfig(cfg)
	Expect(err).NotTo(HaveOccurred())

	// Prevent the metrics listener being created
	metrics.DefaultBindAddress = "0"
}, 60)

var _ = AfterSuite(func() {
	Expect(testenv.Stop()).To(Succeed())

	// Put the DefaultBindAddress back
	metrics.DefaultBindAddress = ":8080"
})
