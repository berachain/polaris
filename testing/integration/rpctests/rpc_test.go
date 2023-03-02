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

package rpc_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"

	"pkg.berachain.dev/stargazer/testutil/network"
)

var (
	// dummyContract  = network.DummyContract
	testKey        = network.TestKey
	addressFromKey = network.AddressFromKey
	signer         = network.Signer

	txData = network.TxData
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testutil/rpc:integration")
}

var _ = Describe("BlockAPIs", func() {
	var net *network.Network
	var client *ethclient.Client

	BeforeEach(func() {
		cfg := network.DefaultConfig()

		var authState authtypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authState)
		newAccount := authtypes.NewBaseAccount(addressFromKey.Bytes(), testKey.PubKey(), 99, 0)
		accounts, err := authtypes.PackAccounts([]authtypes.GenesisAccount{newAccount})
		Expect(err).To(BeNil())
		authState.Accounts = append(authState.Accounts, accounts[0])
		cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authState)

		var bankState banktypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankState)
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: sdk.MustBech32ifyAddressBytes("cosmos", addressFromKey.Bytes()),
			Coins:   sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000000000000))),
		})
		cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankState)

		net = network.New(GinkgoT(), cfg)
		_, err = net.WaitForHeightWithTimeout(1, 15*time.Second)
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
		Expect(header.Number.Uint64()).To(Equal(blockNumber))
	})

	It("eth_getBlockByHash - get block header by hash for block 1", func() {

		// TODO: expected failure because offchain kv is not working yet, this reads from offchain kv
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).To(BeNil())
		hash := block.Hash()
		blockHeaderByHash, err := client.HeaderByHash(context.Background(), hash)
		Expect(err).To(BeNil())
		Expect(blockHeaderByHash.Hash()).To(Equal(block.Header().Hash()))
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
		Expect(count).To(Equal(uint(0)))

	})

	// TODO: get blockByNumber fails on repeated calls
	It("eth_getBlockTransactionCountByHash - get txns in block by block hash. 1 txn submitted", func() {

		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).To(BeNil())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
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
		cfg := network.DefaultConfig()

		var authState authtypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authState)
		newAccount := authtypes.NewBaseAccount(addressFromKey.Bytes(), testKey.PubKey(), 99, 0)
		accounts, err := authtypes.PackAccounts([]authtypes.GenesisAccount{newAccount})
		Expect(err).To(BeNil())
		authState.Accounts = append(authState.Accounts, accounts[0])
		cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authState)

		var bankState banktypes.GenesisState
		cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankState)
		bankState.Balances = append(bankState.Balances, banktypes.Balance{
			Address: sdk.MustBech32ifyAddressBytes("cosmos", addressFromKey.Bytes()),
			Coins:   sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000000000000))),
		})
		cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankState)

		net = network.New(GinkgoT(), cfg)
		_, err = net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).To(BeNil())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())

	})

	// TODO: fails currently as unable to check for signed tx
	It("eth_sendTransaction", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).To(BeNil())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())
	})

	It("eth_getTransactionByHash --- pending tx", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).To(BeNil())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())

		txReceived, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		Expect(err).To(BeNil())
		Expect(isPending).To(BeTrue())
		Expect(txReceived.Hash()).To(Equal(signedTx.Hash()))
	})

	It("eth_getTransactionByHash --- finished tx", func() {
		// create signed transaction
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).To(BeNil())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
		Expect(err).To(BeNil())
		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		Expect(err).To(BeNil())

		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		net.WaitForHeight(int64(blockNumber) + 5)

		txReceived, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
		Expect(err).To(BeNil())
		Expect(isPending).To(BeFalse()) // should work?
		Expect(txReceived.Hash()).To(Equal(signedTx.Hash()))
	})

	It("eth_getTransactionReceipt - get txn receipt for txn submitted by hash", func() {
		tx := ethtypes.NewTx(txData)
		ethKey, err := testKey.ToECDSA()
		Expect(err).To(BeNil())
		signedTx, err := ethtypes.SignTx(tx, signer, ethKey)
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
		Expect(err).To(BeNil())
		client, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())
	})

	It("eth_gasPrice - check gas price", func() {
		gas, err := client.SuggestGasPrice(context.Background())
		Expect(err).To(BeNil())
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
		Expect(err).To(BeNil())
		Expect(gas).To(BeNumerically(">", 0))
	})

})

var _ = Describe("BlockAPIs", func() {

})
