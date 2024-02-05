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

package distribution_test

import (
	"math/big"
	"testing"

	bbindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/bank"
	bindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/distribution"
	sbindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/staking"
	tbindings "github.com/berachain/polaris/contracts/bindings/testing"
	network "github.com/berachain/polaris/e2e/localnet/network"
	utils "github.com/berachain/polaris/e2e/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/berachain/polaris/e2e/localnet/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDistributionPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/precompile/distribution")
}

var _ = Describe("Distribution Precompile", func() {
	var (
		tf                *network.TestFixture
		precompile        *bindings.DistributionModule
		stakingPrecompile *sbindings.StakingModule
		bankPrecompile    *bbindings.BankModule
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
		// Setup the governance precompile.
		precompile, _ = bindings.NewDistributionModule(
			common.HexToAddress("0x69"),
			tf.EthClient(),
		)
		// Setup the staking precompile.
		stakingPrecompile, _ = sbindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient())
		bankPrecompile, _ = bbindings.NewBankModule(
			common.BytesToAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
			tf.EthClient(),
		)
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

	It("should be able to get if withdraw address is enabled", func() {
		res, err := precompile.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(BeTrue())
	})

	It("should be able to set withdraw address with cosmos address", func() {
		addr := sdk.AccAddress("addr")
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.SetWithdrawAddress(txr, common.BytesToAddress(addr))
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("should be able to set withdraw address with ethereum address", func() {
		addr := sdk.AccAddress("addr")
		ethAddr := common.BytesToAddress(addr)
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.SetWithdrawAddress(txr, ethAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("should be able to get delegator reward", func() {
		// Delegate some tokens to an active validator.
		validators, _, err := stakingPrecompile.GetBondedValidators(nil, sbindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		val := validators[0]
		delegateAmt := big.NewInt(123450000000)
		txr := tf.GenerateTransactOpts("alice")
		txr.Value = delegateAmt
		tx, err := stakingPrecompile.Delegate(txr, val.OperatorAddr, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for 5 blocks to be produced, to make sure there are rewards.
		for i := 0; i < 5; i++ {
			Expect(tf.WaitForNextBlock()).To(Succeed())
		}

		// Preview the withdraw rewards.
		rewards, err := precompile.GetTotalDelegatorReward(nil, tf.Address("alice"))
		Expect(err).ToNot(HaveOccurred())
		Expect(rewards).ToNot(BeNil())
		for _, reward := range rewards {
			Expect(reward.Amount.Cmp(new(big.Int))).To(Equal(1))
		}

		// Withdraw the rewards.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = precompile.WithdrawDelegatorReward(txr, tf.Address("alice"), val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})

	It("Should be able to call the precompile via the contract", func() {
		// Deploy the contract.
		contractAddress, tx, contract, err := tbindings.DeployDistributionWrapper(
			tf.GenerateTransactOpts("alice"),
			tf.EthClient(),
			common.HexToAddress("0x69"),
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		txr := tf.GenerateTransactOpts("alice")
		amt := big.NewInt(123450000000)
		tx, err = bankPrecompile.Send(txr, contractAddress, []bbindings.CosmosCoin{
			{
				Denom:  "abera",
				Amount: amt,
			},
		})
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Delegate some tokens to a validator.
		validators, _, err := stakingPrecompile.GetBondedValidators(nil, sbindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		val := validators[0]
		amt = big.NewInt(123450000000)
		txr = tf.GenerateTransactOpts("alice")
		txr.Value = amt
		tx, err = contract.Delegate(txr, val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for the 2 block to be produced, to make sure there are rewards.
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// Withdraw the rewards.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = contract.WithdrawRewards(txr, contractAddress, val.OperatorAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Get withdraw address enabled.
		res, err := contract.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(BeTrue())

		// Set the withdraw address.
		addr := sdk.AccAddress("addr")
		ethAddr := common.BytesToAddress(addr)
		txr = tf.GenerateTransactOpts("alice")
		tx, err = contract.SetWithdrawAddress(txr, ethAddr)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})
})
