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

package distribution_test

import (
	"os"
	"testing"

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

func TestDistributionPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/distribution:integration")
}

var (
	tf         *integration.TestFixture
	precompile *bindings.DistributionModule
	validator  common.Address
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	// Setup the governance precompile.
	precompile, _ = bindings.NewDistributionModule(
		common.HexToAddress("0x69"),
		tf.EthClient,
	)
	// Set the validator address.
	validator = common.Address(tf.Network.Validators[0].Address.Bytes())

	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Distribution Precompile", func() {
	It("Should be able to call the precompile directly", func() {
		// Wait one block.
		tf.Network.WaitForNextBlock()

		// Get withdraw address Enabled.
		_, err := precompile.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())

		// Set withdraw address Common.Address.
		txr := tf.GenerateTransactOpts("")
		tx, err := precompile.SetWithdrawAddress(txr, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Withdraw rewards.
		txr = tf.GenerateTransactOpts("")
		rv := network.GetDistrValidator()
		tx, err = precompile.WithdrawDelegatorReward(txr, network.TestAddress, rv)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)

		// Set withdraw address Bech32.
		txr = tf.GenerateTransactOpts("")
		bech32Addr := cosmlib.AddressToAccAddress(validator).String()
		tx, err = precompile.SetWithdrawAddress0(txr, bech32Addr)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})

	It("Should be able to call the precompile via a contract", func() {
		// Deploy the contract.
		_, tx, contract, err := tbindings.DeployDistributionTestHelper(
			tf.GenerateTransactOpts(""),
			tf.EthClient,
			common.HexToAddress("0x69"),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Get withdraw address Enabled.
		res, err := contract.GetWithdrawEnabled(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(BeTrue())

		// Withdraw the rewards.
		txr := tf.GenerateTransactOpts("")
		rv := network.GetDistrValidator()
		tx, err = contract.WithdrawRewards(txr, network.TestAddress, rv)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)

		// Set withdraw address Common.Address.
		txr = tf.GenerateTransactOpts("")
		tx, err = contract.SetWithdrawAddress(txr, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())

	})
})
