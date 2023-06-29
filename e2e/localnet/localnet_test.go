// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package localnet

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLocalnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var _ = Describe("Fixture", func() {
	var (
		c   Localnet
		err error
	)

	BeforeEach(func() {
		_, err = NewDockerizedNetwork(
			"base",
			baseImageName,
			baseContext,
			baseDockerfile,
			"8545/tcp",
			"8546/tcp",
			map[string]string{
				"GO_VERSION":               "1.20.4",
				"PRECOMPILE_CONTRACTS_DIR": "./contracts",
				"GOOS":                     "linux",
				"GOARCH":                   "arm64",
			},
		)
		Expect(err).ToNot(HaveOccurred())

		c, err = NewDockerizedNetwork(
			"localnet1",
			localnetImageName,
			localnetContext,
			localnetDockerfile,
			"8545/tcp",
			"8546/tcp",
			map[string]string{
				"GO_VERSION":   "1.20.4",
				"BASE_IMAGE":   baseImageName,
				"GENESIS_PATH": "config",
			},
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())
	})

	AfterEach(func() {
		Expect(c.Stop()).To(Succeed())

	})

	It("should create a container", func() {
		err = c.Start()
		Expect(err).ToNot(HaveOccurred())
	})
})
