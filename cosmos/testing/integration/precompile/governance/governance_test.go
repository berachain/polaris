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

package governance_test

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

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/governance:integration")
}

var (
	tf         *integration.TestFixture
	precompile *bindings.GovernanceModule
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	// Setup the governance precompile.
	precompile, _ = bindings.NewGovernanceModule(
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

var _ = Describe("Governance Precompile Directly", func() {
	It("should be able to call the precompile methods directly", func() {
		// Prepare the message.
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

		// Call the precompile create proposal method.
		txr := tf.GenerateTransactOpts("")
		txr.Value = big.NewInt(100)
		tx, err := precompile.SubmitProposal(txr, proposalBz, messageBz)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)

		// Check that the proposal was created.
		proposals, err := precompile.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(proposals).To(HaveLen(2)) // genesis proposal + new proposal

		// Check get proposal.
		res, err := precompile.GetProposal(nil, 2)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.Id).To(Equal(uint64(2)))

		// Should be able to vote on the proposal.
		txr = tf.GenerateTransactOpts("")
		tx, err = precompile.Vote(txr, 2, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Should be able to cancel a proposal.
		txr = tf.GenerateTransactOpts("")
		tx, err = precompile.CancelProposal(txr, 1)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})

	It("should be able to call the precompile methods via a contract", func() {
		// Deploy the contract.
		_, tx, contract, err := tbindings.DeployGovernanceWrapper(
			tf.GenerateTransactOpts(""),
			tf.EthClient,
			common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Should be able to get proposal.
		res, err := contract.GetProposal(nil, 2)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.Id).To(Equal(uint64(2)))

		// Should be able to get proposals.
		res1, err := contract.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(res1).To(HaveLen(1)) // just the genesis proposal.

		// Should be able to create a proposal.
		govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
		initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
		message := &banktypes.MsgSend{
			FromAddress: cosmlib.AddressToAccAddress(govAcc).String(),
			ToAddress:   cosmlib.AddressToAccAddress(network.TestAddress).String(),
			Amount:      initDeposit,
		}
		messageBz, err := message.Marshal()
		Expect(err).ToNot(HaveOccurred())
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
		tx, err = contract.SubmitProposalWrapepr(
			tf.GenerateTransactOpts(""),
			proposalBz,
			messageBz,
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})
})
