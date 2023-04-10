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

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bank", func() {
	BeforeEach(func() {

	})

	It("should call functions on the precompile directly", func() {

		//  function getBalance(address accountAddress, string calldata denom) external view returns (uint256);

		//  function getAllBalance(address accountAddress) external view returns (Coin[] memory);

		//  function getSpendableBalanceByDenom(address accountAddress, string calldata denom) external view returns (uint256);

		//  function getSpendableBalances(address accountAddress) external view returns (Coin[] memory);

		aberaSupply, err := bankPrecompile.GetSupplyOf(nil, "abera")
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("aberaSupply is: ", aberaSupply)

		totalSupply, err := bankPrecompile.GetTotalSupply(nil)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("totalSupply is: ", totalSupply)

		params, err := bankPrecompile.GetParams(nil)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("params is: ", params)

		// todo: set denom metadata then get 
		getTestMetadata()
		denomMetadata, err := bankPrecompile.GetDenomMetadata(nil, "abera")
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("denomMetadata is: ", denomMetadata)

		denomsMetadata, err := bankPrecompile.GetDenomsMetadata(nil)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("denomsMetadata is: ", denomsMetadata)

		//  function getDenomMetadata(string calldata denom) external view returns (DenomMetadata memory);

		//  function getSendEnabled(string[] calldata denoms) external view returns (SendEnabled memory);

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
