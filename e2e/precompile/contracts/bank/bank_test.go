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

package bank_test

import (
	"math/big"
	"testing"

	bindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/bank"
	localnet "github.com/berachain/polaris/e2e/localnet/network"
	utils "github.com/berachain/polaris/e2e/precompile"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/berachain/polaris/e2e/localnet/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/precompile/bank")
}

var _ = Describe("Bank", func() {
	var (
		tf             *localnet.TestFixture
		bankPrecompile *bindings.BankModule

		denom  = "abera"
		denom2 = "atoken"
		denom3 = "stake"
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = localnet.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
		bankPrecompile, _ = bindings.NewBankModule(
			common.HexToAddress("0x4381dC2aB14285160c808659aEe005D51255adD7"), tf.EthClient())
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

	It("should call functions on the precompile directly", func() {
		numberOfDenoms := 7
		coinsToBeSent := []bindings.CosmosCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(1000),
			},
		}
		expectedAllBalance := []bindings.CosmosCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(100),
			},
			{
				Denom:  denom2,
				Amount: big.NewInt(100),
			},
			{
				Denom:  denom3,
				Amount: big.NewInt(1000000000000000000),
			},
		}

		// charlie initially has 1000000000000000000 abera
		balance, err := bankPrecompile.GetBalance(nil, tf.Address("charlie"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance.Cmp(big.NewInt(1000000000000000000))).To(Equal(0))

		// Send 1000 bera from alice to charlie
		tx, err := bankPrecompile.Send(
			tf.GenerateTransactOpts("alice"),
			tf.Address("charlie"),
			coinsToBeSent,
		)
		Expect(err).ShouldNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait one block.
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// charlie now has 1000000000000001000 abera
		balance, err = bankPrecompile.GetBalance(nil, tf.Address("charlie"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000000000000001000)))

		// bob has 100 abera and 100 atoken
		allBalance, err := bankPrecompile.GetAllBalances(nil, tf.Address("bob"))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(allBalance).To(Equal(expectedAllBalance))

		spendableBalanceByDenom, err := bankPrecompile.GetSpendableBalance(nil, tf.Address("bob"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalanceByDenom).To(Equal(big.NewInt(100)))

		spendableBalances, err := bankPrecompile.GetAllSpendableBalances(nil, tf.Address("bob"))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalances).To(Equal(expectedAllBalance))

		atokenSupply, err := bankPrecompile.GetSupply(nil, "asupply")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(atokenSupply).To(Equal(big.NewInt(1000000000000000000)))

		totalSupply, err := bankPrecompile.GetAllSupply(nil)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(totalSupply).To(HaveLen(numberOfDenoms))
	})
})
