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

package misc

import (
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var _ = Describe("Miscellaneous Precompile Tests", func() {
	Describe("calling a precompile from the constructor", func() {
		It("should successfully deploy", func() {
			addr, tx, contract, err := tbindings.DeployPrecompileConstructor(
				tf.GenerateTransactOpts(""), tf.EthClient,
			)
			Expect(err).NotTo(HaveOccurred())
			ExpectSuccessReceipt(tf.EthClient, tx)
			Expect(contract).ToNot(BeNil())
			Expect(addr).ToNot(BeEmpty())

			err = tf.Network.WaitForNextBlock()
			Expect(err).NotTo(HaveOccurred())

			aberaAddr, err := contract.Abera(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(aberaAddr).ToNot(BeEmpty())
			aberaStr, err := contract.Denom(nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(aberaStr).To(Equal("abera"))
		})
	})
})
