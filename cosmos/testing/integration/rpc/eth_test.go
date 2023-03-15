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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
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
	BeforeEach(func() {
		net, client = StartPolarisNetwork(GinkgoT())
		_ = net
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	It("eth_chainId, eth_gasPrice, eth_blockNumber, eth_getBalance", func() {
		chainID, err := client.ChainID(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(chainID.String()).To(Equal("69420"))
		gasPrice, err := client.SuggestGasPrice(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(gasPrice).ToNot(BeNil())
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).ToNot(HaveOccurred())
		Expect(blockNumber).To(BeNumerically(">", 0))
		balance, err := client.BalanceAt(context.Background(), network.TestAddress, nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(balance).To(Equal(big.NewInt(1000000000000000000)))
	})

	It("should deploy, mint tokens, and check balance", func() {
		// Deploy the contract
		erc20Contract := DeployERC20(BuildTransactor(client), client)

		// Mint tokens
		tx, err := erc20Contract.Mint(BuildTransactor(client),
			network.TestAddress, big.NewInt(100000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		// Check the erc20 balance
		erc20Balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, network.TestAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(erc20Balance).To(Equal(big.NewInt(100000000)))
	})
})
