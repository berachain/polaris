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

package bank_test

import (
	"math/big"
	"testing"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	localnet "pkg.berachain.dev/polaris/e2e/localnet/network"
	"pkg.berachain.dev/polaris/e2e/localnet/utils"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/e2e/precompile/bank")
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
		tf = localnet.NewTestFixture(GinkgoT())
		bankPrecompile, _ = bindings.NewBankModule(
			common.HexToAddress("0x4381dC2aB14285160c808659aEe005D51255adD7"), tf.EthClient())
	})

	AfterEach(func() {
		err := tf.Teardown()
		Expect(err).ShouldNot(HaveOccurred())
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
		evmDenomMetadata := bindings.IBankModuleDenomMetadata{
			Name:        "Berachain bera",
			Symbol:      "BERA",
			Description: "The Bera.",
			DenomUnits: []bindings.IBankModuleDenomUnit{
				{Denom: "bera", Exponent: uint32(0), Aliases: []string{"bera"}},
				{Denom: "nbera", Exponent: uint32(9), Aliases: []string{"nanobera"}},
				{Denom: "abera", Exponent: uint32(18), Aliases: []string{"attobera"}},
			},
			Base:    "bera",
			Display: "bera",
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
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

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

		denomMetadata, err := bankPrecompile.GetDenomMetadata(nil, "bera")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(denomMetadata).To(Equal(evmDenomMetadata))

		sendEnabled, err := bankPrecompile.GetSendEnabled(nil, "abera")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(sendEnabled).To(BeTrue())
	})
})
