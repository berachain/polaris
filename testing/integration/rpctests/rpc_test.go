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

package rpctests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"pkg.berachain.dev/stargazer/eth/common"
	network "pkg.berachain.dev/stargazer/testing/utils/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	testKey = network.TestKey
	signer  = network.Signer

	txData = network.TxData
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "integration")
}

var _ = Describe("BlockAPIs", func() {
	var net *network.Network
	var client *ethclient.Client

	BeforeEach(func() {
		cfg := network.ConfigWithTestAccount()
		net = network.New(GinkgoT(), cfg)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).ToNot(HaveOccurred())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).ToNot(HaveOccurred())

	})

	It("eth_chainId - get chain id", func() {
		chainID, err := client.ChainID(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(chainID.String()).To(Equal("69420"))
	})

	It("eth_blockNumber - get latest block number", func() {
		// Dial GetBlock
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("eth_getBlockByNumber - get block header by number", func() {
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		header, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		Expect(err).ToNot(HaveOccurred())
		Expect(header.Number.Uint64()).To(Equal(blockNumber))
	})

	It("eth_getBlockByHash - get block header by hash for block 1", func() {

		// TODO: expected failure because offchain kv is not working yet, this reads from offchain kv
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).ToNot(HaveOccurred())
		hash := block.Hash()
		blockHeaderByHash, err := client.HeaderByHash(context.Background(), hash)
		Expect(err).ToNot(HaveOccurred())
		Expect(blockHeaderByHash.Hash()).To(Equal(block.Header().Hash()))
		Expect(blockHeaderByHash.Number).To(Equal(block.Number()))
	})

	It("eth_getBlockTransactionCountByHash - get txns in block by block hash. No txns submitted", func() {
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		Expect(err).ToNot(HaveOccurred())
		Expect(block.Number().Uint64()).To(Equal(blockNumber))
		count, err := client.TransactionCount(context.Background(), block.Hash())
		Expect(err).ToNot(HaveOccurred())
		Expect(count).To(Equal(uint(0)))

	})

	// TODO: get blockByNumber fails on repeated calls
	It("eth_getBlockTransactionCountByHash - get txns in block by block hash. 1 txn submitted", func() {

		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).ToNot(HaveOccurred())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).ToNot(HaveOccurred())

	})

})

var _ = Describe("TransactionAPIs", func() {
	var net *network.Network
	var client *ethclient.Client
	BeforeEach(func() {

		cfg := network.ConfigWithTestAccount()
		net = network.New(GinkgoT(), cfg)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).ToNot(HaveOccurred())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).ToNot(HaveOccurred())

	})

	// TODO: fails currently as unable to check for signed tx
	It("eth_sendTransaction", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).ToNot(HaveOccurred())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).ToNot(HaveOccurred())
	})

	It("eth_getTransactionByHash --- pending tx", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).ToNot(HaveOccurred())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).ToNot(HaveOccurred())

		txReceived, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		Expect(err).ToNot(HaveOccurred())
		Expect(isPending).To(BeTrue())
		Expect(txReceived.Hash()).To(Equal(signedTx.Hash()))
	})

	It("eth_getTransactionByHash --- finished tx", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).ToNot(HaveOccurred())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).ToNot(HaveOccurred())

		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())

		// TODO: update this to go over all blocks
		_, err = net.WaitForHeight(int64(blockNumber) + 5)
		Expect(err).ToNot(HaveOccurred())

		txReceived, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		Expect(err).ToNot(HaveOccurred())
		Expect(isPending).To(BeFalse()) // should work?
		Expect(txReceived.Hash()).To(Equal(signedTx.Hash()))
	})

	It("eth_getTransactionReceipt - get txn receipt for txn submitted by hash", func() {
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).ToNot(HaveOccurred())
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).ToNot(HaveOccurred())

		receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		Expect(err).ToNot(HaveOccurred())
		Expect(receipt.TxHash).To(Equal(signedTx.Hash()))

	})

	// TODO: verify if this works like this
	It("eth_getTransactionByBlockHashAndIndex - get txn by block hash and index for block 1 (no txns)", func() {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).ToNot(HaveOccurred())
		_, err = client.TransactionInBlock(context.Background(), block.Hash(), 1)
		Expect(err).To(HaveOccurred())
	})

})

var _ = Describe("GasPriceAPIs", func() {

	var net *network.Network
	var client *ethclient.Client

	BeforeEach(func() {
		cfg := network.DefaultConfig()
		net = network.New(GinkgoT(), cfg)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).ToNot(HaveOccurred())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).ToNot(HaveOccurred())
	})

	It("eth_gasPrice - check gas price", func() {
		gas, err := client.SuggestGasPrice(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(gas).To(Equal(big.NewInt(1000000001)))
	})

	It("eth_estimateGas -- checking if it's implemented", func() {
		gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     common.HexToAddress("0x1"),
			To:       nil,
			GasPrice: big.NewInt(1000000001),
			Value:    big.NewInt(1000000001),
			Data:     []byte("0x1"),
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(gas).To(BeNumerically(">", 0))
	})

})
