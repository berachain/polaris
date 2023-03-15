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
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/jsonrpc:integration")
}

var _ = Describe("Network", func() {
	var net *network.Network
	var client *ethclient.Client
	var rpcClient *gethrpc.Client
	var ctx context.Context
	BeforeEach(func() {
		net = network.New(GinkgoT(), network.DefaultConfig())
		time.Sleep(1 * time.Second)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).ToNot(HaveOccurred())

		// Dial an Ethereum RPC Endpoint
		rpcClient, err = gethrpc.DialContext(ctx, net.Validators[0].APIAddress+"/eth/rpc")
		Expect(err).ToNot(HaveOccurred())
		client = ethclient.NewClient(rpcClient)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	It("eth_chainId", func() {
		chainID, err := client.ChainID(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(chainID.String()).To(Equal("69420"))
	})

	It("eth_gasPrice", func() {
		gasPrice, err := client.SuggestGasPrice(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(gasPrice).ToNot(BeNil())
	})

	It("eth_blockNumber", func() {
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("should deploy, mint tokens, and check balance", func() {
		nonce, err := client.PendingNonceAt(context.Background(), network.TestAddress)
		Expect(err).ToNot(HaveOccurred())

		// Set up the auth object
		gasPrice, err := client.SuggestGasPrice(context.Background())
		Expect(err).ToNot(HaveOccurred())

		chainID, err := client.ChainID(context.Background())

		Expect(err).ToNot(HaveOccurred())

		auth, err := bind.NewKeyedTransactorWithChainID(network.ECDSATestKey, chainID)
		Expect(err).ToNot(HaveOccurred())

		// Build transaction opts object.
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)        // in wei
		auth.GasLimit = uint64(3_000_000) // in units
		auth.GasPrice = gasPrice

		balance, err := client.BalanceAt(context.Background(), network.TestAddress, nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000000000000000000)))

		// // Deploy the contract
		// _, _, _, err = bindings.DeploySolmateERC20(auth, client)
		// Expect(err).ToNot(HaveOccurred())
		// _, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		// _, err = bind.WaitMined(ctx, client, tx)
		// Expect(err).ToNot(HaveOccurred())

		// // Mint tokens
		// auth.Nonce = big.NewInt(int64(nonce + 1))
		// tx, err = contract.Mint(auth, network.TestAddress, big.NewInt(100000000))
		// Expect(err).ToNot(HaveOccurred())
		// ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		// _, err = bind.WaitMined(ctx, client, tx)
		// Expect(err).ToNot(HaveOccurred())

		// // Check the balance
		// balance, err := contract.BalanceOf(&bind.CallOpts{}, network.TestAddress)
		// Expect(err).ToNot(HaveOccurred())
		// Expect(balance.String()).To(Equal("100000000"))
	})
})

// func expectSuccessReceipt(
// 	client *ethclient.Client,
// 	hash common.Hash,
// ) *coretypes.Receipt {
// 	receipt, err := client.TransactionReceipt(context.Background(), hash)
// 	Expect(err).ToNot(HaveOccurred())
// 	Expect(receipt.Status).To(Equal(uint64(0x1)))
// 	return receipt
// }
