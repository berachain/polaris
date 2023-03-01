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

package network_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/params"

	"pkg.berachain.dev/stargazer/testutil/network"
)

var (
	dummyContract = common.HexToAddress("0x9fd0aA3B78277a1E717de9D3de434D4b812e5499")
	testKey, _    = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	signer = types.LatestSignerForChainID(params.DefaultChainConfig.ChainID)

	txData = &types.DynamicFeeTx{
		Nonce: 0,
		To:    &dummyContract,
		Gas:   100000,
		Data:  []byte("abcdef"),
	}
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testutil/network:integration")
}

var _ = Describe("BlockAPIs", func() {
	var net *network.Network
	var client *ethclient.Client

	BeforeEach(func() {
		cfg := network.DefaultConfig()
		genesis := make(map[string]json.RawMessage)
		fmt.Println("genesis", genesis)

		net = network.New(GinkgoT(), cfg)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).To(BeNil())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())

	})

	It("eth_chainId - get chain id", func() {
		chainID, err := client.ChainID(context.Background())
		Expect(err).To(BeNil())
		Expect(chainID.String()).To(Equal("42069"))
	})

	It("eth_blockNumber - get latest block number", func() {
		// Dial GetBlock
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("eth_getBlockByNumber - get block header by number", func() {
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		header, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		Expect(err).To(BeNil())
		Expect(header.Number).To(Equal(blockNumber))
	})

	It("eth_getBlockByHash - get block header by hash for block 1", func() {

		// TODO: expected failure because offchain kv is not working yet, this reads from offchain kv
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).To(BeNil())
		hash := block.Hash()
		blockHeaderByHash, err := client.HeaderByHash(context.Background(), hash)
		Expect(err).To(BeNil())
		Expect(*blockHeaderByHash).To(Equal(*block.Header()))
		Expect(blockHeaderByHash.Number).To(Equal(block.Number()))
	})

	It("eth_getBlockTransactionCountByHash - get txns in block by block hash. No txns submitted", func() {
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		Expect(err).To(BeNil())
		Expect(block.Number().Uint64()).To(Equal(blockNumber))
		count, err := client.TransactionCount(context.Background(), block.Hash())
		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))

	})

	// TODO: get blockByNumber fails on repeated calls
	It("eth_getBlockTransactionCountByHash - get txns in block by block hash. 1 txn submitted", func() {

		tx := types.NewTx(txData)
		signedTx, err := types.SignTx(tx, signer, testKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())

	})

})

var _ = Describe("TransactionAPIs", func() {
	var net *network.Network
	var client *ethclient.Client
	BeforeEach(func() {
		net = network.New(GinkgoT(), network.DefaultConfig())
		_, err := net.WaitForHeightWithTimeout(3, 15*time.Second)
		Expect(err).To(BeNil())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())

	})

	// TODO: fails currently as unable to check for signed tx
	It("eth_sendTransaction", func() {
		// create signed transaction
		tx := types.NewTx(txData)
		signedTx, err := types.SignTx(tx, signer, testKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())
	})

	// TODO: fails currently as unable to check for signed tx
	It("eth_getTransactionByHash", func() {
		// create signed transaction
		tx := types.NewTx(txData)
		signedTx, err := types.SignTx(tx, signer, testKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())

		txReceived, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		Expect(err).To(BeNil())
		Expect(isPending).To(BeFalse())
		Expect(txReceived.Hash()).To(Equal(signedTx.Hash()))
	})

	It("eth_getTransactionReceipt - get txn receipt for txn submitted by hash", func() {
		tx := types.NewTx(txData)
		signedTx, err := types.SignTx(tx, signer, testKey)
		Expect(err).To(BeNil())
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())
		receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		Expect(err).To(BeNil())
		Expect(receipt.TxHash).To(Equal(signedTx.Hash()))

	})

	// TODO: verify if this works like this
	It("eth_getTransactionByBlockHashAndIndex - get txn by block hash and index for block 1 (no txns)", func() {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).To(BeNil())
		_, err = client.TransactionInBlock(context.Background(), block.Hash(), 1)
		Expect(err).To(BeNil())
	})

	// // TODO: same as above so need to understand how it works exactly
	// It("eth_getTransactionByBlockNumberAndIndex - get txn by block num and index for block 1 (1 txn)", func() {
	// 	Expect(true)
	// })
})
