// SPDX-License-Identifier: Apache-2.0
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

package misc_test

import (
	"testing"

	tbindings "github.com/berachain/polaris/contracts/bindings/testing"
	network "github.com/berachain/polaris/e2e/localnet/network"
	utils "github.com/berachain/polaris/e2e/precompile"

	. "github.com/berachain/polaris/e2e/localnet/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMiscellaneousPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/precompile/misc")
}

var _ = Describe("Miscellaneous Precompile Tests", func() {
	var tf *network.TestFixture

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
	})

	AfterEach(func() {
		// Dump logs and stop the container here.
		if !CurrentSpecReport().Failure.IsZero() {
			logs, err := tf.DumpLogs()
			Expect(err).ToNot(HaveOccurred())
			GinkgoWriter.Println(logs)
		}
		Expect(tf.Teardown()).To(Succeed())
	})

	Describe("calling a precompile from the constructor", func() {
		It("should successfully deploy", func() {
			txr := tf.GenerateTransactOpts("alice")
			addr, tx, contract, err := tbindings.DeployPrecompileConstructor(txr, tf.EthClient())
			Expect(err).NotTo(HaveOccurred())

			err = tf.WaitForNextBlock()
			Expect(err).NotTo(HaveOccurred())

			ExpectSuccessReceipt(tf.EthClient(), tx)
			Expect(contract).ToNot(BeNil())
			Expect(addr).ToNot(BeEmpty())
		})
	})
})
