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

package jsonrpc

import (
	"context"
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var _ = Describe("Tx Pool", func() {
	// BeforeEach(func() {

	// })

	It("should handle txpool requests", func() {
		// Run some transactions for bob
		_, tx, contract, err := tbindings.DeployConsumeGas(
			tf.GenerateTransactOpts("alice"), client,
		)
		Expect(err).NotTo(HaveOccurred())
		ExpectSuccessReceipt(client, tx)
		tx, err = contract.ConsumeGas(tf.GenerateTransactOpts("alice"), big.NewInt(10000))
		Expect(err).NotTo(HaveOccurred())
		tf.Network.WaitForNextBlock()
		bobCurrNonce, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(bobCurrNonce).To(BeNumerically(">=", 2))
	})
})
