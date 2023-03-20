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
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/precompile:integration")
}

var _ = Describe("Staking", func() {
	var net *network.Network
	var client *ethclient.Client
	var validator common.Address
	var stakingPrecompile *bindings.StakingModule

	BeforeEach(func() {
		net, client = StartPolarisNetwork(GinkgoT())
		validator = common.BytesToAddress(net.Validators[0].Address.Bytes())
		stakingPrecompile, _ = bindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), client)
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	It("should call functions on the precompile directly", func() {
		validators, err := stakingPrecompile.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(validators).To(ContainElement(validator))

		delegated, err := stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		txr := BuildTransactor(client)
		txr.Value = big.NewInt(1000000000000)
		tx, err := stakingPrecompile.Delegate(txr, validator, big.NewInt(100000000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		delegated, err = stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(100000000000))).To(Equal(0))
	})

	// It("should be able to call a precompile from a smart contract", func() {
	// 	_, tx, contract, err := tbindings.DeployLiquidStaking(
	// 		BuildTransactor(client),
	// 		client,
	// 		"myToken",
	// 		"MTK",
	// 		common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"),
	// 		validator,
	// 	)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	ExpectMined(client, tx)
	// 	ExpectSuccessReceipt(client, tx)

	// 	value, err := contract.TotalDelegated(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(value.Cmp(big.NewInt(0))).To(Equal(0))

	// 	addresses, err := contract.GetActiveValidators(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(addresses).To(HaveLen(1))
	// 	Expect(addresses[0]).To(Equal(validator))

	// 	// Send tokens to the contract
	// 	txr := BuildTransactor(client)
	// 	txr.GasLimit = 0
	// 	txr.Value = big.NewInt(100000000000)
	// 	tx, err = contract.Delegate(txr, big.NewInt(100000000000))
	// 	Expect(err).ToNot(HaveOccurred())
	// 	ExpectMined(client, tx)
	// 	ExpectSuccessReceipt(client, tx)

	// 	// Verify the delegation actually succeeded.
	// 	value, err = contract.TotalDelegated(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(value.Cmp(big.NewInt(100000000000))).To(Equal(0))
	// })
})
