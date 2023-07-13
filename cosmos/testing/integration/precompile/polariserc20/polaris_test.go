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

package polariserc20_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	cbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos"
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	network "pkg.berachain.dev/polaris/e2e/localnet/network"
	utils "pkg.berachain.dev/polaris/e2e/localnet/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPolarisERC20(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/polariserc20:integration")
}

var _ = Describe("ERC20", func() {
	var tf *network.TestFixture

	BeforeEach(func() {
		// Setup the network and clients here.
		tf = network.NewTestFixture(GinkgoT())
	})

	AfterEach(func() {
		err := tf.Teardown()
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("deploy a polaris erc20 and call it from another contract", func() {
		It("should approve and use the transfer from method", func() {
			swapperAddress, tx, swapper, err := tbindings.DeploySwapper(tf.GenerateTransactOpts("alice"), tf.EthClient())
			Expect(err).ToNot(HaveOccurred())
			utils.ExpectSuccessReceipt(tf.EthClient(), tx)

			// check that the new ERC20 is minted to TestAddress
			tokenAddr, err := swapper.GetPolarisERC20(nil, "bAKT")
			Expect(err).ToNot(HaveOccurred())
			Expect(tokenAddr.Bytes()).To(Equal(common.Address{}.Bytes()))

			err = tf.WaitForNextBlock()
			Expect(err).ToNot(HaveOccurred())

			// Create a polaris erc20 contract from the address.
			tokenAddr, tx, token, err := cbindings.DeployPolarisERC20(
				tf.GenerateTransactOpts("alice"),
				tf.EthClient(),
				"bAKT",
			)
			Expect(err).ToNot(HaveOccurred())
			utils.ExpectSuccessReceipt(tf.EthClient(), tx)

			// Call the polaris erc20 contract to set the allowance of the swapper contract.
			tx, err = token.Approve(
				tf.GenerateTransactOpts("alice"),
				swapperAddress,
				big.NewInt(100),
			)
			Expect(err).ToNot(HaveOccurred())
			utils.ExpectSuccessReceipt(tf.EthClient(), tx)

			// Get the current allowance of the swapper contract.
			res, err := token.Allowance(
				nil,
				tf.Address("alice"),
				swapperAddress,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Cmp(big.NewInt(100))).To(Equal(0))

			// Call the swapper contract to swap the polaris erc20 token from the msg.sender.
			tx, err = swapper.Deposit(
				tf.GenerateTransactOpts("alice"),
				tokenAddr,
				big.NewInt(50),
			)
			Expect(err).ToNot(HaveOccurred())
			utils.ExpectSuccessReceipt(tf.EthClient(), tx)

			// Call the balance of the swapper contract to check the balance of the polaris erc20 token.
			res, err = token.BalanceOf(
				nil,
				swapperAddress,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Cmp(big.NewInt(50))).To(Equal(0))
		})

	})
})
