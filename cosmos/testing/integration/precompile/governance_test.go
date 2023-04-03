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

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ = Describe("Governance Precompile", func() {
	When("Calling the governance precompile directly", func() {
		It("should be able to create a proposal", func() {
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

			// Call the precompile.
			txr := tf.GenerateTransactOpts("")
			txr.Value = big.NewInt(100)
			tx, err := governancePrecompile.SubmitProposal(txr, proposalBz, messageBz)
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
		})
		It("should be able to get proposal", func() {
			res, err := governancePrecompile.GetProposal(nil, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Id).To(Equal(uint64(2)))
		})
		It("should be able to get proposals", func() {
			proposals, err := governancePrecompile.GetProposals(nil, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(proposals).To(HaveLen(2)) // one in genesis, one we just submitted.
		})
		It("should be able to vote on a proposal", func() {
			txr := tf.GenerateTransactOpts("")
			tx, err := governancePrecompile.Vote(txr, 2, 1, "metadata")
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
		})
	})
	When("Calling the precompile via a contract", func() {
		var (
			wrapperContract *tbindings.GovernanceWrapper
		)
		BeforeEach(func() {
			// Deploy the governance wrapper contract.
			_, tx, contract, err := tbindings.DeployGovernanceWrapper(
				tf.GenerateTransactOpts(""),
				tf.EthClient,
				common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
			)
			wrapperContract = contract
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
			ExpectSuccessReceipt(tf.EthClient, tx)
		})
		It("should be able to get proposal", func() {
			res, err := wrapperContract.GetProposal(nil, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Id).To(Equal(uint64(2)))
		})
		It("should be able to get proposals", func() {
			res2, err := wrapperContract.GetProposals(nil, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(res2).To(HaveLen(2))
		})
		It("should be able to create a proposal", func() {
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
			// Submit a proposal.
			txr := tf.GenerateTransactOpts("")
			tx, err := wrapperContract.SubmitProposalWrapepr(txr, proposalBz, messageBz)
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(tf.EthClient, tx)
		})
		// TODO: Test Times Out but works in isolation.
		// !!DEV
		// It("should be able to vote on a proposal", func() {
		// 	// Vote on the proposal.
		// 	txr := tf.GenerateTransactOpts("")
		// 	tx, err := wrapperContract.Vote(txr, 2, 1, "metadata")
		// 	Expect(err).ToNot(HaveOccurred())
		// 	ExpectMined(tf.EthClient, tx)
		// })
	})
})
