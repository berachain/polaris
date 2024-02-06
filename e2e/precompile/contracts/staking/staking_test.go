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

package staking_test

import (
	"math/big"
	"testing"
	"time"

	bbindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/bank"
	bindings "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/staking"
	tbindings "github.com/berachain/polaris/contracts/bindings/testing"
	network "github.com/berachain/polaris/e2e/localnet/network"
	utils "github.com/berachain/polaris/e2e/precompile"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/berachain/polaris/e2e/localnet/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/precompile/staking")
}

var _ = Describe("Staking", func() {
	var (
		tf                *network.TestFixture
		stakingPrecompile *bindings.StakingModule
		bankPrecompile    *bbindings.BankModule
		delegateAmt       = big.NewInt(123450000000)
		validator         common.Address
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
		stakingPrecompile, _ = bindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient())
		bankPrecompile, _ = bbindings.NewBankModule(
			common.BytesToAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
			tf.EthClient(),
		)
		validators, _, err := stakingPrecompile.GetBondedValidators(nil, bindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(validators).To(HaveLen(1))
		validator = common.BytesToAddress(validators[0].OperatorAddr[:])
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
		delegated, err := stakingPrecompile.GetDelegation(nil, tf.Address("alice"), validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		txr := tf.GenerateTransactOpts("alice")
		txr.Value = delegateAmt
		tx, err := stakingPrecompile.Delegate(txr, validator, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		delegated, err = stakingPrecompile.GetDelegation(nil, tf.Address("alice"), validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(delegateAmt)).To(Equal(0))

		delVals, _, err := stakingPrecompile.GetDelegatorValidators(
			nil, tf.Address("alice"), bindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(delVals).To(HaveLen(1))
		Expect(delVals[0].OperatorAddr).To(Equal(validator))

		undelegateAmt := new(big.Int).Div(delegateAmt, big.NewInt(2))
		tx, err = stakingPrecompile.Undelegate(
			tf.GenerateTransactOpts("alice"),
			validator,
			undelegateAmt,
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
		Expect(tf.WaitForBlock(1)).To(Succeed())

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

		vals, _, err := stakingPrecompile.GetValidators(nil, bindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(vals).To(HaveLen(1))
		Expect(vals[0].OperatorAddr).To(Equal(validator))

		val, err := stakingPrecompile.GetValidator(nil, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(val.OperatorAddr).To(Equal(vals[0].OperatorAddr))
	})

	It("should be able to call a precompile from a smart contract", func() {
		contractAddr, tx, contract, err := tbindings.DeployLiquidStaking(
			tf.GenerateTransactOpts("alice"),
			tf.EthClient(),
			"myToken",
			"MTK",
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

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
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Send tokens to the contract to delegate and mint LSD.
		txr = tf.GenerateTransactOpts("alice")
		txr.GasLimit = 0
		txr.Value = delegateAmt
		tx, err = contract.Delegate(txr, delegateAmt)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

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
