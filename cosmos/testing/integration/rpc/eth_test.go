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
	"github.com/ethereum/go-ethereum/ethclient"

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

var (
	net      *network.Network
	client   *ethclient.Client
	wsClient *ethclient.Client
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	net = StartPolarisTestNetwork(GinkgoT())
	client, wsClient = BuildEthClients(GinkgoT(), net)
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

var _ = Describe("Network", func() {
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

	It("eth_estimateGas", func() {
		// Estimate the gas required for a transaction
		from := network.TestAddress
		to := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
		value := big.NewInt(1000000000000)

		msg := ethereum.CallMsg{
			From:  from,
			To:    &to,
			Value: value,
		}

		gas, err := client.EstimateGas(context.Background(), msg)
		Expect(err).ToNot(HaveOccurred())
		Expect(gas).To(BeNumerically(">", 0))
	})

	It("should deploy, mint tokens, and check balance, eth_getTransactionByHash", func() {
		// Deploy the contract
		erc20Contract := DeployERC20(BuildTransactor(client), client)

		// Mint tokens
		tx, err := erc20Contract.Mint(BuildTransactor(client),
			network.TestAddress, big.NewInt(100000000))
		Expect(err).ToNot(HaveOccurred())

		// Get the transaction by its hash, it should be pending here.
		txHash := tx.Hash() // TODO: UNCOMMENT
		// fetchedTx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		// Expect(err).ToNot(HaveOccurred())
		// Expect(isPending).To(BeTrue())
		// Expect(fetchedTx.Hash()).To(Equal(txHash))

		// Wait for it to be mined.
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		// Get the transaction by its hash, it should be mined here.
		fetchedTx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		Expect(err).ToNot(HaveOccurred())
		Expect(isPending).To(BeFalse())
		Expect(fetchedTx.Hash()).To(Equal(txHash))

		// Check the erc20 balance
		erc20Balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, network.TestAddress)
		Expect(err).ToNot(HaveOccurred())
		Expect(erc20Balance).To(Equal(big.NewInt(100000000)))
	})
})
