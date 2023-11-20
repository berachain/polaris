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

package localnet_test

import (
	"context"
	"os"

	"github.com/berachain/polaris/contracts/bindings/testing"
	localnet "github.com/berachain/polaris/e2e/localnet/network"
	coretypes "github.com/berachain/polaris/eth/core/types"

	. "github.com/berachain/polaris/e2e/localnet/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("evm load testing", func() {
	var (
		tf       *localnet.TestFixture
		contract *testing.LoadTestOps
	)

	BeforeEach(func() {
		if os.Getenv("REMOTE_NODE_URL") != "" {
			tf = localnet.NewRemoteTestFixture(GinkgoT(), os.Getenv("REMOTE_NODE_URL"))
		} else {
			tf = localnet.NewTestFixture(GinkgoT(), localnet.NewFixtureConfig(
				"../../../e2e/precompile/polard/config/",
				"polard/base:v0.0.0",
				"polard/localnet:v0.0.0",
				"goodcontainer",
				"8545/tcp",
				"8546/tcp",
				"1.21.3",
			))
			Expect(tf).ToNot(BeNil())
		}
		// Wait for the next block.
		Expect(tf.WaitForNextBlock()).ToNot(HaveOccurred())

		// Check to make sure "alice" has balance to complete the test.
		balance, err := tf.EthClient().BalanceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(*balance).To(BeNumerically(">", 0))

		var tx *coretypes.Transaction
		_, tx, contract, err = testing.DeployLoadTestOps(
			tf.GenerateTransactOpts("alice"), tf.EthClient(),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	AfterEach(func() {
		// Dump logs and stop the containter here.
		if !CurrentSpecReport().Failure.IsZero() {
			logs, err := tf.DumpLogs()
			Expect(err).ToNot(HaveOccurred())
			GinkgoWriter.Println(logs)
		}
		Expect(tf.Teardown()).To(Succeed())
	})

	Context("evm load testing", func() {
		It("should run a load test", func() {
			tx, err := contract.LoadData(tf.GenerateTransactOpts("alice"))
			Expect(err).ToNot(HaveOccurred())
			ExpectSuccessReceipt(tf.EthClient(), tx)
		})
	})
})
