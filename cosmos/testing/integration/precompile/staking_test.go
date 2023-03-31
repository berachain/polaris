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

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
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

var (
	tf                   *integration.TestFixture
	stakingPrecompile    *bindings.StakingModule
	governancePrecompile *bindings.GovernanceModule // TODO: merge with feat/distr-precompile where setup is refactored.
	validator            common.Address
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	validator = common.Address(tf.Network.Validators[0].Address.Bytes())
	stakingPrecompile, _ = bindings.NewStakingModule(
		common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient)

	// Setup the governance precompile.
	governancePrecompile, _ = bindings.NewGovernanceModule(
		common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"), tf.EthClient,
	)
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Staking", func() {
	It("should call functions on the precompile directly", func() {
		validators, err := stakingPrecompile.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(validators).To(ContainElement(validator))

		delegated, err := stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		txr := tf.GenerateTransactOpts("")
		txr.Value = big.NewInt(1000000000000)
		tx, err := stakingPrecompile.Delegate(txr, validator, big.NewInt(100000000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		delegated, err = stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(100000000000))).To(Equal(0))
	})

	It("should be able to call a precompile from a smart contract", func() {
		_, tx, contract, err := tbindings.DeployLiquidStaking(
			tf.GenerateTransactOpts(""),
			tf.EthClient,
			"myToken",
			"MTK",
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"),
			validator,
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		value, err := contract.TotalDelegated(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Cmp(big.NewInt(0))).To(Equal(0))

		addresses, err := contract.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(addresses).To(HaveLen(1))
		Expect(addresses[0]).To(Equal(validator))

		// Send tokens to the contract
		txr := tf.GenerateTransactOpts("")
		txr.GasLimit = 0
		txr.Value = big.NewInt(100000000000)
		tx, err = contract.Delegate(txr, big.NewInt(100000000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Verify the delegation actually succeeded.
		value, err = contract.TotalDelegated(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Cmp(big.NewInt(100000000000))).To(Equal(0))
	})
	// TODO: Move to own file once setup.go is merged in here.
	When("Testing Governance Precompile", func() {
		It("Should be able to call precompile directly", func() {
			// Prepare the Message.
			govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
			message := &banktypes.MsgSend{
				FromAddress: cosmlib.AddressToAccAddress(govAcc).String(),
				ToAddress:   cosmlib.AddressToAccAddress(network.TestAddress).String(),
				Amount:      initDeposit,
			}
			messageBz, err := message.Marshal()
			Expect(err).ToNot(HaveOccurred())
			// Prepare the Proposal.
			proposal := v1.MsgSubmitProposal{
				InitialDeposit: initDeposit,
				Proposer:       cosmlib.AddressToAccAddress(network.TestAddress).String(),
				Metadata:       "metadata",
				Title:          "title",
				Summary:        "summary",
				Expedited:      false,
			}
			proposalBz, err := proposal.Marshal()
			Expect(err).ToNot(HaveOccurred())

			// Submit the transaction.
			txr := tf.GenerateTransactOpts("")
			txr.Value = big.NewInt(100)
			tx, err := governancePrecompile.SubmitProposal(txr, proposalBz, messageBz)
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)

			// Get a proposal from the chain.
			res, err := governancePrecompile.GetProposal(nil, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Id).To(Equal(uint64(2)))

			// Get the proposals from the chain.
			proposals, err := governancePrecompile.GetProposals(nil, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(proposals).To(HaveLen(2)) // one in genesis, one we just submitted.

			// Vote on the proposal.
			txr = tf.GenerateTransactOpts("")
			tx, err = governancePrecompile.Vote(txr, 2, 1, "metadata")
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)

			// TODO: Fix this test.
			// Vote Weighted.
			// txr = tf.GenerateTransactOpts("")
			// weight, err := math.LegacyNewDecFromStr("0.4")
			// Expect(err).ToNot(HaveOccurred())
			// tx, err = governancePrecompile.VoteWeighted(
			// 	txr,
			// 	2,
			// 	[]bindings.IGovernanceModuleWeightedVoteOption{
			// 		{
			// 			VoteOption: int32(1),
			// 			Weight:     weight.String(),
			// 		},
			// 	},
			// 	"metadata",
			// )
			// Expect(err).ToNot(HaveOccurred())
			// ExpectMined(tf.EthClient, tx)
		})

		It("should be able to call from a contract", func() {
			// Deploy the governance wrapper contract.
			_, tx, contract, err := tbindings.DeployGovernanceWrapper(
				tf.GenerateTransactOpts(""),
				tf.EthClient,
				common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
			)
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
			ExpectSuccessReceipt(tf.EthClient, tx)

			// Prepare the Message.
			govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
			message := &banktypes.MsgSend{
				FromAddress: cosmlib.AddressToAccAddress(govAcc).String(),
				ToAddress:   cosmlib.AddressToAccAddress(network.TestAddress).String(),
				Amount:      initDeposit,
			}
			messageBz, err := message.Marshal()
			Expect(err).ToNot(HaveOccurred())
			// Prepare the Proposal.
			proposal := v1.MsgSubmitProposal{
				InitialDeposit: initDeposit,
				Proposer:       cosmlib.AddressToAccAddress(network.TestAddress).String(),
				Metadata:       "metadata",
				Title:          "title",
				Summary:        "summary",
				Expedited:      false,
			}
			proposalBz, err := proposal.Marshal()
			Expect(err).ToNot(HaveOccurred())

			// Send funds to the contract.
			txr := tf.GenerateTransactOpts("")
			tx, err = contract.SubmitProposalWrapepr(txr, proposalBz, messageBz)
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
		})
	})

})
