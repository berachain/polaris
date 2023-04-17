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

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var (
	tf             *integration.TestFixture
	precompile     *bindings.GovernanceModule
	wrapper        *tbindings.GovernanceWrapper
	bankPrecompile *bindings.BankModule
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
	// Setup the bank precompile.
	bankPrecompile, _ = bindings.NewBankModule(
		common.HexToAddress("0x4381dC2aB14285160c808659aEe005D51255adD7"), tf.EthClient)

	// Deploy the contract.
	contractAddr, tx, contract, err := tbindings.DeployGovernanceWrapper(
		tf.GenerateTransactOpts("alice"),
		tf.EthClient,
		common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
	)
	Expect(err).ToNot(HaveOccurred())
	ExpectMined(tf.EthClient, tx)
	ExpectSuccessReceipt(tf.EthClient, tx)
	wrapper = contract

	// Alice Submits a proposal.
	amt := sdk.NewInt(100000000)
	prop, msg := propAndMsgBz(cosmlib.AddressToAccAddress(tf.Address("alice")).String(), amt)
	txr := tf.GenerateTransactOpts("alice")
	tx, err = precompile.SubmitProposal(txr, prop, msg)
	Expect(err).ToNot(HaveOccurred())
	ExpectMined(tf.EthClient, tx)
	ExpectSuccessReceipt(tf.EthClient, tx)

	// Send coins to the wrapper.
	coins := []bindings.IBankModuleCoin{
		{
			Denom:  "stake",
			Amount: big.NewInt(amt.Int64()),
		},
	}
	txr = tf.GenerateTransactOpts("alice")
	tx, err = bankPrecompile.Send(txr, tf.Address("alice"), contractAddr, coins)
	Expect(err).ToNot(HaveOccurred())
	ExpectMined(tf.EthClient, tx)
	ExpectSuccessReceipt(tf.EthClient, tx)

	// Wrapper submits a proposal.
	prop, msg = propAndMsgBz(cosmlib.AddressToAccAddress(contractAddr).String(), amt)
	txr = tf.GenerateTransactOpts("alice")
	tx, err = wrapper.Submit(txr, prop, msg, "stake", big.NewInt(amt.Int64()))
	Expect(err).ToNot(HaveOccurred())
	ExpectMined(tf.EthClient, tx)
	ExpectSuccessReceipt(tf.EthClient, tx)

	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Call the Precompile Directly", func() {

	It("Should be able to get a proposal", func() {
		// Call directly.
		res, err := precompile.GetProposal(nil, 1)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.Id).To(Equal(uint64(1)))

		// Call via wrapper.
		res2, err := wrapper.GetProposal(nil, 1)
		Expect(err).ToNot(HaveOccurred())
		Expect(res2.Id).To(Equal(uint64(1)))
	})

	It("Should be able to get proposals", func() {
		// Call directly.
		res, err := precompile.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(HaveLen(2))

		// Call via wrapper.
		res2, err := wrapper.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(res2).To(HaveLen(2))
	})

	It("Should be able to vote on a proposal", func() {
		// Call directly.
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.Vote(txr, 1, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Call via wrapper.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.Vote(txr, 1, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})

	It("Should be able to cancel a proposal", func() {
		// Call directly.
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.CancelProposal(txr, 1)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Call via wrapper.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.CancelProposal(txr, 2)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})
})

func propAndMsgBz(proposer string, amount math.Int) ([]byte, []byte) {
	// Prepare the message.
	govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
	initDeposit := sdk.NewCoins(sdk.NewCoin("stake", amount))
	message := &banktypes.MsgSend{
		FromAddress: cosmlib.AddressToAccAddress(govAcc).String(),
		ToAddress:   cosmlib.AddressToAccAddress(tf.Address("alice")).String(),
		Amount:      initDeposit,
	}
	messageBz, err := message.Marshal()
	Expect(err).ToNot(HaveOccurred())

	// Prepare the Proposal.
	proposal := v1.MsgSubmitProposal{
		InitialDeposit: initDeposit,
		// Proposer:       cosmlib.AddressToAccAddress(tf.Address("alice")).String(),
		Proposer:  proposer,
		Metadata:  "metadata",
		Title:     "title",
		Summary:   "summary",
		Expedited: true,
	}
	proposalBz, err := proposal.Marshal()
	Expect(err).ToNot(HaveOccurred())

	return proposalBz, messageBz
}
