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

package governance_test

import (
	"math/big"
	"testing"

	"cosmossdk.io/core/address"
	sdkmath "cosmossdk.io/math"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	bbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/bank"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	cosmtypes "pkg.berachain.dev/polaris/cosmos/types"
	network "pkg.berachain.dev/polaris/e2e/localnet/network"
	utils "pkg.berachain.dev/polaris/e2e/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/e2e/localnet/utils"
)

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/precompile/governance")
}

var tf *network.TestFixture

var _ = Describe("Call the Precompile Directly", func() {
	var (
		precompile     *bindings.GovernanceModule
		wrapper        *tbindings.GovernanceWrapper
		bankPrecompile *bbindings.BankModule
		wrapperAddr    common.Address
		accCodec       = addresscodec.NewBech32Codec(cosmtypes.Bech32PrefixAccAddr)
	)

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT(), utils.NewPolarisFixtureConfig())
		err := tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())

		// Setup the governance precompile.
		precompile, _ = bindings.NewGovernanceModule(
			common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"), tf.EthClient(),
		)
		// Setup the bank precompile.
		bankPrecompile, _ = bbindings.NewBankModule(
			common.HexToAddress("0x4381dC2aB14285160c808659aEe005D51255adD7"), tf.EthClient())

		// Deploy the contract.
		var tx *types.Transaction
		wrapperAddr, tx, wrapper, err = tbindings.DeployGovernanceWrapper(
			tf.GenerateTransactOpts("alice"),
			tf.EthClient(),
			common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient(), tx)
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Alice Submits a proposal.
		amt := sdkmath.NewInt(100000000)
		prop, _ := propAndMsgBz(
			accCodec, cosmlib.MustStringFromEthAddress(accCodec, tf.Address("alice")), amt)
		txr := tf.GenerateTransactOpts("alice")
		tx, err = precompile.SubmitProposal(txr, prop)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Send coins to the wrapper.
		coins := []bbindings.CosmosCoin{
			{
				Denom:  "abgt",
				Amount: big.NewInt(amt.Int64()),
			},
		}
		txr = tf.GenerateTransactOpts("alice")
		tx, err = bankPrecompile.Send(txr, wrapperAddr, coins)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wrapper submits a proposal.
		prop, _ = propAndMsgBz(accCodec, cosmlib.MustStringFromEthAddress(accCodec, wrapperAddr), amt)
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.Submit(txr, prop, "abgt", big.NewInt(amt.Int64()))
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Wait for next block.
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())
		err = tf.WaitForNextBlock()
		Expect(err).ToNot(HaveOccurred())
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

	It("Should be able to get a proposal", func() {
		// Call directly.
		res, err := precompile.GetProposal(nil, 1)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.Id).To(Equal(uint64(1)))

		// Call via wrapper.
		res2, err := wrapper.GetProposal(nil, 1)
		Expect(err).ToNot(HaveOccurred())
		Expect(res2.Id).To(Equal(uint64(1)))

		// Call directly.
		getProposalsRes, pageRes, err := precompile.GetProposals(
			nil, 0,
			bindings.CosmosPageRequest{
				Key:        "test",
				Offset:     0,
				Limit:      10,
				CountTotal: true,
				Reverse:    false,
			},
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(getProposalsRes).To(HaveLen(2))
		Expect(pageRes).ToNot(BeNil())

		// Call via wrapper.
		wrapperRes, err := wrapper.GetProposals(nil, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(wrapperRes).To(HaveLen(2))

		// Call directly.
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.Vote(txr, 1, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Call via wrapper.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.Vote(txr, 1, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Call directly.
		votes, _, err := precompile.GetProposalVotes(nil, 1, bindings.CosmosPageRequest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(votes).To(HaveLen(2))

		// Call directly.
		aliceVote, err := precompile.GetProposalVotesByVoter(nil, 1, tf.Address("alice"))
		Expect(err).ToNot(HaveOccurred())
		Expect(aliceVote.Voter).To(Equal(tf.Address("alice")))

		// Call directly.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = precompile.CancelProposal(txr, 1)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)

		// Call via wrapper.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.CancelProposal(txr, 2)
		Expect(err).ToNot(HaveOccurred())
		ExpectSuccessReceipt(tf.EthClient(), tx)
	})
})

func propAndMsgBz(accCodec address.Codec, proposer string, amount sdkmath.Int) ([]byte, []byte) {
	// Prepare the message.
	govAcc := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
	initDeposit := sdk.NewCoins(sdk.NewCoin("abgt", amount))
	message := &banktypes.MsgSend{
		FromAddress: cosmlib.MustStringFromEthAddress(accCodec, govAcc),
		ToAddress:   cosmlib.MustStringFromEthAddress(accCodec, tf.Address("alice")),
		Amount:      initDeposit,
	}
	messageBz, err := message.Marshal()
	Expect(err).ToNot(HaveOccurred())

	// Prepare the Proposal.
	proposal := v1.MsgSubmitProposal{
		InitialDeposit: initDeposit,
		Proposer:       proposer,
		Metadata:       "metadata",
		Title:          "title",
		Summary:        "summary",
		Expedited:      true,
	}
	proposalBz, err := proposal.Marshal()
	Expect(err).ToNot(HaveOccurred())

	return proposalBz, messageBz
}
