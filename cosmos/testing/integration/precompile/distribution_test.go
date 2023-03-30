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
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"

	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
)

var _ = Describe("Distribution", func() {
	// It("should call functions on the precompile directly", func() {
	// 	ok, err := distributionPrecompile.GetWithdrawEnabled(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(ok).To(BeTrue())

	// 	// Set withdraw address.
	// 	txr := tf.GenerateTransactOpts("")
	// 	tx, err := distributionPrecompile.SetWithdrawAddress(txr, validator)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(tx).ToNot(BeNil())
	// 	ExpectMined(tf.EthClient, tx)
	// 	ExpectSuccessReceipt(tf.EthClient, tx)

	// 	// Set withdraw address bech32.
	// 	txr = tf.GenerateTransactOpts("")
	// 	bech32Addr := cosmlib.AddressToAccAddress(validator).String()
	// 	tx, err = distributionPrecompile.SetWithdrawAddress0(txr, bech32Addr)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(tx).ToNot(BeNil())
	// 	ExpectMined(tf.EthClient, tx)
	// 	ExpectSuccessReceipt(tf.EthClient, tx)

	// 	// Get withdraw enabled.
	// 	res, err := distributionPrecompile.GetWithdrawEnabled(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(res).To(BeTrue())
	// })

	It("should call functions on the precompile via a contract", func() {
		_, tx, contract, err := tbindings.DeployDistributionTestHelper(
			tf.GenerateTransactOpts(""),
			tf.EthClient,
			common.HexToAddress("0x93354845030274cD4bf1686Abd60AB28EC52e1a7"),
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		Expect(contract).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)

		// Set Withdraw Address.
		txr := tf.GenerateTransactOpts("")
		tx, err = contract.SetWithdrawAddress(txr, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(tx).ToNot(BeNil())
		ExpectMined(tf.EthClient, tx)
		ExpectSuccessReceipt(tf.EthClient, tx)
	})
})
