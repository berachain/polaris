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
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

func TestRpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/jsonrpc:integration")
}

var tf *integration.TestFixture

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Network", func() {
	It("should support eth_chainId", func() {
		chainID, err := tf.EthClient.ChainID(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(chainID.String()).To(Equal("69420"))
	})

	It("should support eth_gasPrice", func() {
		gasPrice, err := tf.EthClient.SuggestGasPrice(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(gasPrice).ToNot(BeNil())
	})

	It("should support eth_blockNumber", func() {
		blockNumber, err := tf.EthClient.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("should support eth_getBalance", func() {
		balance, err := tf.EthClient.BalanceAt(context.Background(), network.TestAddress, nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(balance.Cmp(big.NewInt(50000000000))).To(Equal(1))
	})

	It("should support eth_estimateGas", func() {
		// Estimate the gas required for a transaction
		from := network.TestAddress
		to := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
		value := big.NewInt(1000000000000)

		msg := ethereum.CallMsg{
			From:  from,
			To:    &to,
			Value: value,
		}

		gas, err := tf.EthClient.EstimateGas(context.Background(), msg)
		Expect(err).ToNot(HaveOccurred())
		Expect(gas).To(BeNumerically(">", 0))
	})

	It("should deploy an erc20, mint tokens, check balance and support eth_getTransactionByHash",
		func() {
			// Deploy a contract
			erc20Contract := DeployERC20(BuildTransactor(tf.EthClient), tf.EthClient)

			// Mint tokens
			tx, err := erc20Contract.Mint(BuildTransactor(tf.EthClient),
				network.TestAddress, big.NewInt(100000000))
			Expect(err).ToNot(HaveOccurred())

			// Get the transaction by its hash, it should be pending here.
			txHash := tx.Hash() // TODO: UNCOMMENT
			// fetchedTx, isPending, err :=
			// tf.EthClient.TransactionByHash(context.Background(), txHash)
			// Expect(err).ToNot(HaveOccurred())
			// Expect(isPending).To(BeTrue())
			// Expect(fetchedTx.Hash()).To(Equal(txHash))

			// Wait for it to be mined.
			// For this test, we aren't looking for the transaction in the mempool, so we wait
			// for it to be mined.
			// TODO: write a mempool searching test.
			ExpectMined(tf.EthClient, tx)
			ExpectSuccessReceipt(tf.EthClient, tx)

			// Get the transaction by its hash, it should be mined here.
			fetchedTx, isPending, err := tf.EthClient.TransactionByHash(context.Background(), txHash)
			Expect(err).ToNot(HaveOccurred())
			Expect(isPending).To(BeFalse())
			Expect(fetchedTx.Hash()).To(Equal(txHash))

			// Check the erc20 balance
			erc20Balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, network.TestAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(erc20Balance).To(Equal(big.NewInt(100000000)))
		})
})
