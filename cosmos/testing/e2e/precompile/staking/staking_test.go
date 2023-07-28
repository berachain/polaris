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

package staking_test

import (
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	bbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	network "pkg.berachain.dev/polaris/e2e/localnet/network"
	utils "pkg.berachain.dev/polaris/e2e/localnet/utils"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/e2e/precompile/staking")
}

var _ = Describe("Staking", func() {
	var (
		tf                *network.TestFixture
		stakingPrecompile *bindings.StakingModule
		bankPrecompile    *bbindings.BankModule
		validator         common.Address
		delegateAmt       = big.NewInt(123450000000)
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT())

		validator = tf.ValAddr()
		stakingPrecompile, _ = bindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient())
		bankPrecompile, _ = bbindings.NewBankModule(
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
			tf.EthClient(),
		)
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

	It("should call functions on the precompile directly", func() {
		validators, err := stakingPrecompile.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(validators).To(ContainElement(validator))

		delegated, err := stakingPrecompile.GetDelegation(nil, tf.Address("alice"), validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		txr := tf.GenerateTransactOpts("alice")
		txr.Value = delegateAmt
		tx, err := stakingPrecompile.Delegate(txr, validator, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

		delegated, err = stakingPrecompile.GetDelegation(nil, tf.Address("alice"), validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(delegateAmt)).To(Equal(0))

		delVals, err := stakingPrecompile.GetDelegatorValidators(nil, tf.Address("alice"))
		Expect(err).ToNot(HaveOccurred())
		Expect(delVals).To(HaveLen(1))
		delValAddr, err := sdk.ValAddressFromBech32(delVals[0].OperatorAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(cosmlib.ValAddressToEthAddress(delValAddr)).To(Equal(validator))

		undelegateAmt := new(big.Int).Div(delegateAmt, big.NewInt(2))
		tx, err = stakingPrecompile.Undelegate(
			tf.GenerateTransactOpts("alice"),
			validator,
			undelegateAmt,
		)
		Expect(err).ToNot(HaveOccurred())
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

		ude, err := stakingPrecompile.GetUnbondingDelegation(
			nil,
			tf.Address("alice"),
			validator,
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(ude).To(HaveLen(1))
		Expect(ude[0].CreationHeight).To(BeNumerically(">=", 1))
		Expect(ude[0].CompletionTime).ToNot(BeEmpty())
		Expect(ude[0].Balance.Cmp(undelegateAmt)).To(Equal(0))

		vals, err := stakingPrecompile.GetValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(vals).To(HaveLen(1))
		valAddr, err := sdk.ValAddressFromBech32(vals[0].OperatorAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(cosmlib.ValAddressToEthAddress(valAddr)).To(Equal(validator))

		val, err := stakingPrecompile.GetValidator(nil, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(val.OperatorAddress).To(Equal(vals[0].OperatorAddress))
	})

	It("should be able to call a precompile from a smart contract", func() {
		contractAddr, tx, contract, err := tbindings.DeployLiquidStaking(
			tf.GenerateTransactOpts("alice"),
			tf.EthClient(),
			"myToken",
			"MTK",
		)
		Expect(err).ToNot(HaveOccurred())
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

		delegated, err := stakingPrecompile.GetDelegation(nil, contractAddr, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		addresses, err := contract.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(addresses).To(HaveLen(1))
		Expect(addresses[0]).To(Equal(validator))

		txr := tf.GenerateTransactOpts("alice")
		amt := big.NewInt(123450000000)
		tx, err = bankPrecompile.Send(txr, contractAddr, []bbindings.CosmosCoin{
			{
				Denom:  "abera",
				Amount: amt,
			},
		})
		Expect(err).ToNot(HaveOccurred())
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

		// Send tokens to the contract to delegate and mint LSD.
		txr = tf.GenerateTransactOpts("alice")
		txr.GasLimit = 0
		txr.Value = delegateAmt
		tx, err = contract.Delegate(txr, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		utils.ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for a couple blocks to query.
		time.Sleep(4 * time.Second)

		// Verify the delegation actually succeeded.
		delegated, err = stakingPrecompile.GetDelegation(nil, contractAddr, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(delegateAmt)).To(Equal(0))

		// Check the balance of LSD ERC20 is minted to sender.
		balance, err := contract.BalanceOf(nil, tf.Address("alice"))
		Expect(err).ToNot(HaveOccurred())
		Expect(balance.Cmp(delegateAmt)).To(Equal(0))
	})
})
