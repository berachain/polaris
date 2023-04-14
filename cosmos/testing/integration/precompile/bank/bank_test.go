// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package bank

import (
	"math/big"
	"os"
	"testing"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/precompile:integration")
}

var (
	tf             *integration.TestFixture
	bankPrecompile *bindings.BankModule
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	bankPrecompile, _ = bindings.NewBankModule(
		common.HexToAddress("0x4381dC2aB14285160c808659aEe005D51255adD7"), tf.EthClient)
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Bank", func() {
	denom := "abera"
	denom2 := "atoken"

	It("should call functions on the precompile directly", func() {
		coinsToBeSent := []bindings.IBankModuleCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(1000),
			},
		}
		expectedAllBalance := []bindings.IBankModuleCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(100),
			},
			{
				Denom:  denom2,
				Amount: big.NewInt(100),
			},
		}
		evmDenomMetadata := bindings.IBankModuleDenomMetadata{
			Name:        "Berachain bera",
			Symbol:      "BERA",
			Description: "The Bera.",
			DenomUnits: []bindings.IBankModuleDenomUnit{
				{Denom: "bera", Exponent: uint32(0), Aliases: []string{"bera"}},
				{Denom: "nbera", Exponent: uint32(9), Aliases: []string{"nanobera"}},
				{Denom: "abera", Exponent: uint32(18), Aliases: []string{"attobera"}},
			},
			Base:    "abera",
			Display: "bera",
		}

		// TestAddress3 initially has 1000000000 abera
		balance, err := bankPrecompile.GetBalance(nil, tf.Address("AccWithLessAbera"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000000000)))

		// Send 1000 bera from TestAddress to TestAddress3
		_, err = bankPrecompile.Send(
			tf.GenerateTransactOpts("MainAcc"),
			tf.Address("MainAcc"),
			tf.Address("AccWithLessAbera"),
			coinsToBeSent,
		)
		Expect(err).ShouldNot(HaveOccurred())

		// Wait one block.
		err = tf.Network.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// TestAddress3 now has 1000001000 abera
		balance, err = bankPrecompile.GetBalance(nil, tf.Address("AccWithLessAbera"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000001000)))

		// TestAddress2 has 100 abera and 100 atoken
		allBalance, err := bankPrecompile.GetAllBalances(nil, tf.Address("AccWith2Denoms"))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(allBalance).To(Equal(expectedAllBalance))

		spendableBalanceByDenom, err := bankPrecompile.GetSpendableBalance(nil, tf.Address("AccWith2Denoms"), denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalanceByDenom).To(Equal(big.NewInt(100)))

		spendableBalances, err := bankPrecompile.GetAllSpendableBalances(nil, tf.Address("AccWith2Denoms"))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalances).To(Equal(expectedAllBalance))

		atokenSupply, err := bankPrecompile.GetSupply(nil, denom2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(atokenSupply).To(Equal(big.NewInt(100)))

		totalSupply, err := bankPrecompile.GetAllSupply(nil)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(totalSupply).To(HaveLen(4))

		denomMetadata, err := bankPrecompile.GetDenomMetadata(nil, denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(denomMetadata).To(Equal(evmDenomMetadata))

		sendEnabled, err := bankPrecompile.GetSendEnabled(nil, "abera")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(sendEnabled).To(BeTrue())
	})

	It("should be able to call a precompile from a smart contract", func() {
		// deploy fundraiser contract with account 0
		contractAddr, tx, contract, err := tbindings.DeployFundraiser(
			tf.GenerateTransactOpts("MainAcc"),
			tf.EthClient,
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient, tx)

		coinsToDonate := []tbindings.IBankModuleCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(1000000),
			},
		}

		// donate 1000000 abera from account 0 to contractAddr
		_, err = contract.Donate(tf.GenerateTransactOpts("MainAcc"), coinsToDonate)
		Expect(err).ToNot(HaveOccurred())

		// Wait one block.
		err = tf.Network.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// contractAddr should have 1000000 abera
		balance, err := bankPrecompile.GetBalance(nil, contractAddr, denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000000)))

		// withdraw all 1000000 abera from contractAddr to account 0
		_, err = contract.WithdrawDonations(tf.GenerateTransactOpts("MainAcc"))
		Expect(err).ToNot(HaveOccurred())

		// Wait one block.
		err = tf.Network.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// contractAddr should have 0 abera
		balance, err = bankPrecompile.GetBalance(nil, contractAddr, denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance.Cmp(big.NewInt(0))).To(Equal(0))
	})
})
