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

package precompile

import (
	"fmt"
	"math/big"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"

	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bank", func() {
	denom := "abera"
	denom2 := "atoken"

	BeforeEach(func() {

	})

	It("should call functions on the precompile directly", func() {

		coinsToBeSent := []precompile.IBankModuleCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(1000),
			},
		}
		expectedAllBalance := []precompile.IBankModuleCoin{
			{
				Denom:  denom,
				Amount: big.NewInt(100),
			},
			{
				Denom:  denom2,
				Amount: big.NewInt(100),
			},
		}

		fmt.Println("bytes for sending", []byte(fmt.Sprintf("%v", coinsToBeSent)))

		fmt.Println("TestAddress1 is: ", network.TestAddress)
		fmt.Println("TestAddress2 is: ", network.TestAddress2)
		fmt.Println("coinsToBeSent is: ", coinsToBeSent)

		// tx, err := bankPrecompile.Send(tf.GenerateTransactOpts(""), network.TestAddress, network.TestAddress3, coinsToBeSent)
		// Expect(err).ShouldNot(HaveOccurred())
		// fmt.Printf("tx is: %v", *tx)

		// time.Sleep(4)

		balance, err := bankPrecompile.GetBalance(nil, network.TestAddress, denom)
		Expect(err).ShouldNot(HaveOccurred())
		fmt.Println("bera balance of TestAddress is: ", balance)

		//  function getBalance(address accountAddress, string calldata denom) external view returns (uint256);
		balance, err = bankPrecompile.GetBalance(nil, network.TestAddress2, denom)
		Expect(err).ShouldNot(HaveOccurred())
		fmt.Println("bera balance of TestAddress2 is: ", balance)
		// Expect(balance).To(Equal(big.NewInt(1000000000000000000)))

		balance, err = bankPrecompile.GetBalance(nil, network.TestAddress3, denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(100)))

		allBalance, err := bankPrecompile.GetAllBalance(nil, network.TestAddress2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(allBalance).To(Equal(expectedAllBalance))

		spendableBalanceByDenom, err := bankPrecompile.GetSpendableBalanceByDenom(nil, network.TestAddress2, denom)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalanceByDenom).To(Equal(big.NewInt(100)))

		spendableBalances, err := bankPrecompile.GetSpendableBalances(nil, network.TestAddress2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(spendableBalances).To(Equal(expectedAllBalance))

		atokenSupply, err := bankPrecompile.GetSupplyOf(nil, denom2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(atokenSupply).To(Equal(big.NewInt(100)))

		totalSupply, err := bankPrecompile.GetTotalSupply(nil)
		Expect(err).ShouldNot(HaveOccurred())
		fmt.Println("totalSupply is: ", totalSupply)

		params, err := bankPrecompile.GetParams(nil)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(params.DefaultSendEnabled).To(BeTrue())

		// denomMetadata, err := bankPrecompile.GetDenomMetadata(nil, denom)
		// Expect(err).ShouldNot(HaveOccurred())
		// fmt.Println("denomMetadata is: ", denomMetadata)

		// denomsMetadata, err := bankPrecompile.GetDenomsMetadata(nil)
		// Expect(err).ShouldNot(HaveOccurred())
		// fmt.Println("denomsMetadata is: ", denomsMetadata)

		//  function getSendEnabled(string[] calldata denoms) external view returns (SendEnabled memory);
		fmt.Println("BEFORE!!!!!!!!!!!!!!!")
		sendEnabled, err := bankPrecompile.GetSendEnabled(nil, []string{ "abera", "atoken"})
		Expect(err).ShouldNot(HaveOccurred())
		// Expect(sendEnabled).To(BeTrue())
		fmt.Println("sendEnabled is: ", sendEnabled)
		fmt.Println("AFTER!!!!!!!!!!!!!!!")
		//  function send(address fromAddress, address toAddress, Coin calldata amount) external payable returns (bool);

		//  function multiSend(Input calldata input, Output[] memory outputs) external payable returns (bool);

	})

	// It("should be able to call a precompile from a smart contract", func() {
	// 	contractAddr, tx, contract, err := tbindings.DeployLiquidBank(
	// 		tf.GenerateTransactOpts(""),
	// 		tf.EthClient,
	// 		"myToken",
	// 		"MTK",
	// 	)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	ExpectMined(tf.EthClient, tx)
	// 	ExpectSuccessReceipt(tf.EthClient, tx)

	// 	delegated, err := bankPrecompile.GetDelegation(nil, contractAddr, validator)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

	// 	addresses, err := contract.GetActiveValidators(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(addresses).To(HaveLen(1))
	// 	Expect(addresses[0]).To(Equal(validator))

	// 	// Send tokens to the contract
	// 	txr := tf.GenerateTransactOpts("")
	// 	txr.GasLimit = 0
	// 	txr.Value = big.NewInt(100000000000)
	// 	tx, err = contract.Delegate(txr, big.NewInt(100000000000))
	// 	Expect(err).ToNot(HaveOccurred())
	// 	ExpectMined(tf.EthClient, tx)
	// 	ExpectSuccessReceipt(tf.EthClient, tx)

	// 	// Verify the delegation actually succeeded.
	// 	delegated, err = bankPrecompile.GetDelegation(nil, contractAddr, validator)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(delegated.Cmp(big.NewInt(100000000000))).To(Equal(0))
	// })
})

func getTestMetadata() []banktypes.Metadata {
	return []banktypes.Metadata{
		{
			Name:        "Berachain bera",
			Symbol:      "BERA",
			Description: "The Bera.",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "bera", Exponent: uint32(0), Aliases: []string{"bera"}},
				{Denom: "nbera", Exponent: uint32(9), Aliases: []string{"nanobera"}},
				{Denom: "abera", Exponent: uint32(18), Aliases: []string{"attobera"}},
			},
			Base:    "abera",
			Display: "bera",
		},
		{
			Name:        "Token",
			Symbol:      "TOKEN",
			Description: "The native staking token of the Token Hub.",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "1token", Exponent: uint32(5), Aliases: []string{"decitoken"}},
				{Denom: "2token", Exponent: uint32(4), Aliases: []string{"centitoken"}},
				{Denom: "3token", Exponent: uint32(7), Aliases: []string{"dekatoken"}},
			},
			Base:    "utoken",
			Display: "token",
		},
	}
}
