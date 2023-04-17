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
	"os"
	"testing"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var (
	tf         *integration.TestFixture
	precompile *bindings.GovernanceModule
	wrapper    *tbindings.GovernanceWrapper
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

	// Deploy the contract.
	_, tx, contract, err := tbindings.DeployGovernanceWrapper(
		tf.GenerateTransactOpts("alice"),
		tf.EthClient,
		common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2"),
	)
	Expect(err).ToNot(HaveOccurred())
	ExpectMined(tf.EthClient, tx)
	ExpectSuccessReceipt(tf.EthClient, tx)
	wrapper = contract

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
		res, err := precompile.GetProposal(nil, 2)
		Expect(err).ToNot(HaveOccurred())
		Expect(res.Id).To(Equal(uint64(2)))

		// Call via wrapper.
		res2, err := wrapper.GetProposal(nil, 3)
		Expect(err).ToNot(HaveOccurred())
		Expect(res2.Id).To(Equal(uint64(3)))
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
		tx, err := precompile.Vote(txr, 2, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Call via wrapper.
		txr = tf.GenerateTransactOpts("alice")
		tx, err = wrapper.Vote(txr, 2, 1, "metadata")
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})
	It("Should be able to cancel a proposal", func() {
		// Call directly.
		txr := tf.GenerateTransactOpts("alice")
		tx, err := precompile.CancelProposal(txr, 2)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		// Call via wrapper
		//TODO: Need https://github.com/berachain/polaris/issues/550, bc msg.sender != proposer
	})
})
