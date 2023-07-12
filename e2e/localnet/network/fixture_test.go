// SPDX-License-Identifier: MIT
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

package localnet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	geth "github.com/ethereum/go-ethereum"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ = Describe("JSON RPC tests", func() {
	var (
		tf     *TestFixture
		client *ethclient.Client
	)

	BeforeEach(func() {
		tf = NewTestFixture(GinkgoT())
		Expect(tf).ToNot(BeNil())
		client = tf.c.EthClient()
	})

	AfterEach(func() {
		Expect(tf.Teardown()).To(Succeed())
	})

	Context("eth namespace", func() {
		It("should connect -- multiple clients", func() {
			// Dial an Ethereum RPC Endpoint
			rpcClient, err := gethrpc.DialContext(context.Background(), tf.c.GetHTTPEndpoint())
			Expect(err).ToNot(HaveOccurred())
			c := ethclient.NewClient(rpcClient)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).ToNot(BeNil())
		})

		It("should support eth_chainId", func() {
			chainID, err := client.ChainID(context.Background())
			Expect(chainID.String()).To(Equal("2061"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should support eth_gasPrice", func() {
			gasPrice, err := client.SuggestGasPrice(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(gasPrice).ToNot(BeNil())
		})

		It("should support eth_blockNumber", func() {
			// Get the latest block
			blockNumber, err := client.BlockNumber(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(blockNumber).To(BeNumerically(">", 0))
		})

		It("should support eth_getBalance", func() {
			// Get the balance of an account
			balance, err := client.BalanceAt(context.Background(), tf.Address("alice"), nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(balance.Uint64()).To(BeNumerically(">", 0))
		})

		It("should support eth_estimateGas", func() {
			// Estimate the gas required for a transaction
			from := tf.Address("alice")
			to := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
			value := big.NewInt(1000000000000)

			msg := geth.CallMsg{
				From:  from,
				To:    &to,
				Value: value,
			}

			gas, err := client.EstimateGas(context.Background(), msg)
			Expect(err).ToNot(HaveOccurred())
			Expect(gas).To(BeNumerically(">", 0))
		})

		It("should deploy, mint tokens and check balance, eth_getTransactionByHash", func() {
			// Deploy the contract
			erc20Contract, deployedAddress := DeployERC20(tf.GenerateTransactOpts("alice"), client)

			// Mint tokens
			tx, err := erc20Contract.Mint(tf.GenerateTransactOpts("alice"),
				tf.Address("alice"), big.NewInt(100000000))
			Expect(err).ToNot(HaveOccurred())

			// Get the transaction by its hash, it should be pending here.
			txHash := tx.Hash()

			// Wait for the receipt.
			receipt := ExpectSuccessReceipt(client, tx)
			Expect(receipt.Logs).To(HaveLen(2))
			for i, log := range receipt.Logs {
				Expect(log.Address).To(Equal(deployedAddress))
				Expect(log.BlockHash).To(Equal(receipt.BlockHash))
				Expect(log.TxHash).To(Equal(txHash))
				Expect(log.TxIndex).To(Equal(uint(0)))
				Expect(log.BlockNumber).To(Equal(receipt.BlockNumber.Uint64()))
				Expect(log.Index).To(Equal(uint(i)))
			}

			// Get the transaction by its hash, it should be mined here.
			fetchedTx, isPending, err := client.TransactionByHash(context.Background(), txHash)
			Expect(err).ToNot(HaveOccurred())
			Expect(isPending).To(BeFalse())
			Expect(fetchedTx.Hash()).To(Equal(txHash))

			// Check the erc20 balance
			erc20Balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, tf.Address("alice"))
			Expect(err).ToNot(HaveOccurred())
			Expect(erc20Balance).To(Equal(big.NewInt(100000000)))
		})
	})
	Context("txpool namespace", func() {
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
			Expect(tf.c.WaitForNextBlock()).To(Succeed())
		})

		It("should handle txpool requests: pending nonce", func() {
			aliceCurrNonce, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(aliceCurrNonce).To(BeNumerically(">=", 2))
			Expect(tf.c.WaitForNextBlock()).To(Succeed())

			// send a transaction and make sure pending nonce is incremented
			_, err = contract.ConsumeGas(tf.GenerateTransactOpts("alice"), big.NewInt(10000))
			Expect(err).NotTo(HaveOccurred())
			alicePendingNonce, err := client.PendingNonceAt(context.Background(), tf.Address("alice"))
			Expect(err).NotTo(HaveOccurred())
			Expect(alicePendingNonce).To(Equal(aliceCurrNonce + 1))
			acn, err := client.NonceAt(context.Background(), tf.Address("alice"), nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(acn).To(Equal(aliceCurrNonce))

			Expect(tf.c.WaitForNextBlock()).To(Succeed())

			aliceCurrNonce, err = client.NonceAt(context.Background(), tf.Address("alice"), nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(aliceCurrNonce).To(Equal(alicePendingNonce))
		})

		It("should handle multiple transactions as queued", func() {
			// Get the starting nonce.
			beforeNonce, err := client.PendingNonceAt(context.Background(), tf.Address("charlie"))
			Expect(err).NotTo(HaveOccurred())

			// send 10 transactions, each one with updated nonce
			var txs []*coretypes.Transaction
			for i := beforeNonce; i < beforeNonce+10; i++ {
				txr := tf.GenerateTransactOpts("charlie")
				txr.Nonce = big.NewInt(int64(i))
				var tx *coretypes.Transaction
				tx, err = contract.ConsumeGas(txr, big.NewInt(50))
				txs = append(txs, tx)
				Expect(err).ToNot(HaveOccurred())
			}

			// check that nonce is updated in memory.
			afterNonce, err := client.PendingNonceAt(context.Background(), tf.Address("charlie"))
			Expect(err).NotTo(HaveOccurred())
			Expect(afterNonce).To(Equal(beforeNonce + uint64(len(txs))))

			// check to make sure all the txs went thru.
			for _, tx := range txs {
				ExpectSuccessReceipt(client, tx)
			}

			// verify the nonce has increased on disk.
			afterNonce, err = client.NonceAt(context.Background(), tf.Address("charlie"), nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(afterNonce).To(Equal(beforeNonce + 10))
		})
	})
	Context("ws namespace", func() {
		var (
			ctx      context.Context
			wsclient *ethclient.Client
		)

		BeforeEach(func() {
			ctx = context.Background()
			wsclient = tf.c.EthWsClient()
		})

		It("should connect -- multiple clients", func() {
			// Dial an Ethereum websocket Endpoint
			ws, err := gethrpc.DialWebsocket(ctx, tf.c.GetWSEndpoint(), "*")
			Expect(err).ToNot(HaveOccurred())
			wsClient := ethclient.NewClient(ws)
			Expect(err).ToNot(HaveOccurred())
			Expect(wsClient).ToNot(BeNil())
		})

		It("should subscribe to new heads", func() {
			// Subscribe to new heads
			sub, err := wsclient.SubscribeNewHead(ctx, make(chan *gethtypes.Header))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})

		It("should subscribe to logs", func() {
			// Subscribe to logs
			sub, err := wsclient.SubscribeFilterLogs(ctx, geth.FilterQuery{}, make(chan gethtypes.Log))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})
	})
})
