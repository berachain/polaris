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

package jsonrpc_test

import (
	"context"
	"math/big"

	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var _ = Describe("Tx Pool", func() {
	var contract *tbindings.ConsumeGas

	BeforeEach(func() {
		var err error
		var tx *coretypes.Transaction
		// Run some transactions for alice
		_, tx, contract, err = tbindings.DeployConsumeGas(
			tf.GenerateTransactOpts("alice"), client,
		)
		Expect(err).NotTo(HaveOccurred())
		ExpectSuccessReceipt(client, tx)
		tx, err = contract.ConsumeGas(tf.GenerateTransactOpts("alice"), big.NewInt(10000))
		Expect(err).NotTo(HaveOccurred())
		ExpectSuccessReceipt(client, tx)
		Expect(tf.Network.WaitForNextBlock()).To(Succeed())
	})

	It("should handle txpool requests: pending nonce", func() {
		aliceCurrNonce, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(aliceCurrNonce).To(BeNumerically(">=", 2))

		Expect(tf.Network.WaitForNextBlock()).To(Succeed())

		// send a transaction and make sure pending nonce is incremented
		_, err = contract.ConsumeGas(tf.GenerateTransactOpts("alice"), big.NewInt(10000))
		Expect(err).NotTo(HaveOccurred())
		alicePendingNonce, err := client.PendingNonceAt(context.Background(), tf.Address("alice"))
		Expect(err).NotTo(HaveOccurred())
		Expect(alicePendingNonce).To(Equal(aliceCurrNonce + 1))
		acn, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(acn).To(Equal(aliceCurrNonce))

		Expect(tf.Network.WaitForNextBlock()).To(Succeed())

		aliceCurrNonce, err = client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(aliceCurrNonce).To(Equal(alicePendingNonce))
	})

	It("should handle multiple transactions as queued", func() {
		beforeNonce, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())

		// send 10 transactions, each one with updated nonce
		for i := beforeNonce; i < beforeNonce+10; i++ {
			txr := tf.GenerateTransactOpts("alice")
			txr.Nonce = big.NewInt(int64(i))
			_, err = contract.ConsumeGas(txr, big.NewInt(500))
			Expect(err).NotTo(HaveOccurred())
		}

		Expect(tf.Network.WaitForNextBlock()).To(Succeed())

		// check that nonce is updated
		afterNonce, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(afterNonce).To(Equal(beforeNonce + 10))
	})
})
