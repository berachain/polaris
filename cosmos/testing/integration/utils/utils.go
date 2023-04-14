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

package utils

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/gomega" //nolint:stylecheck,revive,gostaticcheck  // Gomega makes sense in tests.
)

const (
	DefaultTimeout = 10 * time.Second
	TxTimeout      = 30 * time.Second
)

// StartPolarisNetwork starts a new in-memory Polaris chain.
func StartPolarisNetwork(t network.TestingT) (*network.Network, *ethclient.Client) {
	var err error
	net := network.New(t, network.DefaultConfig())
	time.Sleep(1 * time.Second)
	_, err = net.WaitForHeightWithTimeout(1, DefaultTimeout)
	Expect(err).ToNot(HaveOccurred())

	// Dial an Ethereum RPC Endpoint
	rpcClient, err := gethrpc.DialContext(context.Background(), net.Validators[0].APIAddress+"/eth/rpc")
	Expect(err).ToNot(HaveOccurred())
	client := ethclient.NewClient(rpcClient)
	Expect(err).ToNot(HaveOccurred())

	return net, client
}

// ExpectedMined waits for a transaction to be mined.
func ExpectMined(client *ethclient.Client, tx *coretypes.Transaction) {
	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	_, err := bind.WaitMined(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())
}

// ExpectSuccessReceipt waits for the transaction to be mined and returns the receipt.
// It also checks that the transaction was successful.
func ExpectSuccessReceipt(
	client *ethclient.Client,
	tx *coretypes.Transaction,
) *coretypes.Receipt {
	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	_, err := bind.WaitMined(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())

	// Verify the receipt is good.
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	Expect(receipt.Status).To(Equal(uint64(0x1))) //nolint:gomnd // success.
	return receipt
}

// DeployERC20 deploys a new ERC20 contract and waits for the transaction to be mined.
// Upon success, it returns a binding to the contract.
func DeployERC20(
	auth *bind.TransactOpts,
	client *ethclient.Client,
) *bindings.SolmateERC20 {
	// Deploy the contract
	expectedAddr, tx, contract, err := bindings.DeploySolmateERC20(auth, client)
	Expect(err).ToNot(HaveOccurred())

	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	_, err = bind.WaitDeployed(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())

	// Just to be safe we also check the receipt.
	receipt, err := bind.WaitMined(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())

	// Ensure that the receipt was successful.
	Expect(err).ToNot(HaveOccurred())
	Expect(receipt.Status).To(Equal(uint64(0x1))) //nolint:gomnd // success.
	// Ensure that the contract address is correct.
	Expect(expectedAddr).To(Equal(receipt.ContractAddress))

	return contract
}
