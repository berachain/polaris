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

package governance

import (
	"math/big"
	"os"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var (
	tf         *integration.TestFixture
	precompile *bindings.GovernanceModule
)

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/governance:integration")
}

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

var _ = Describe("Governance Precompile", func() {
	It("Should be able to submit a proposal", func() {
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
	})

	It("Should be able to get proposals", func() {
		// Call the precompile get proposals method.
		proposals, err := precompile.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(proposals).To(HaveLen(1))
	})

	It("Should be able to get a proposal", func() {
		// Call the precompile get proposal method.
		proposal, err := precompile.GetProposal(nil, 1)
		Expect(err).ToNot(HaveOccurred())
		Expect(proposal.Id).To(Equal(uint64(1)))
	})

	It("Should be able to call the vote precompile", func() {
		propBz, msgBz := createProposalAndMsg()
		txr := tf.GenerateTransactOpts("")
		txr.Value = big.NewInt(100)
		tx, err := precompile.SubmitProposal(txr, propBz, msgBz)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)

		// Call the precompile get proposal method.
		proposal, err := precompile.GetProposal(nil, 2)
		Expect(err).ToNot(HaveOccurred())

		// Wait a 2 blocks to make sure that the vote period of 1 second is in effect.
		err = tf.Network.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())
		err = tf.Network.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// Call the precompile vote method.
		txr = tf.GenerateTransactOpts("")
		txr.Value = big.NewInt(100)
		tx, err = precompile.Vote(txr, proposal.Id, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
	})
})

func createProposalAndMsg() ([]byte, []byte) {
	govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
	initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
	message := &banktypes.MsgSend{
		FromAddress: cosmlib.AddressToAccAddress(govAcc).String(),
		ToAddress:   cosmlib.AddressToAccAddress(network.TestAddress).String(),
		Amount:      initDeposit,
	}
	messageBz, err := message.Marshal()
	if err != nil {
		panic(err)
	}

	// Prepare the Proposal.
	proposal := v1.MsgSubmitProposal{
		InitialDeposit: initDeposit,
		Proposer:       cosmlib.AddressToAccAddress(network.TestAddress).String(),
		Metadata:       "metadata",
		Title:          "title",
		Summary:        "summary",
		Expedited:      true, // So can be voted on.
	}
	proposalBz, err := proposal.Marshal()
	if err != nil {
		panic(err)
	}

	return proposalBz, messageBz
}
