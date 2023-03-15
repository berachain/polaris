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
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/gomega" //nolint:stylecheck,revive // Gomega makes sense in tests.
)

const (
	defaultTimeout = 10 * time.Second
)

// BuildTransactor builds a transaction opts object.
func BuildTransactor(
	client *ethclient.Client,
) *bind.TransactOpts {
	// Get the nonce from the RPC.
	// TODO: switch to pending once the txpool is finished. https://github.com/berachain/polaris/issues/385
	// Get the nonce from the RPC.
	blockNumber, err := client.BlockNumber(context.Background())
	Expect(err).ToNot(HaveOccurred())
	// nonce, err := client.PendingNonceAt(context.Background(), network.TestAddress)
	nonce, err := client.NonceAt(context.Background(), network.TestAddress, big.NewInt(int64(blockNumber)))

	Expect(err).ToNot(HaveOccurred())
	// Set up the auth object
	gasPrice, err := client.SuggestGasPrice(context.Background())
	Expect(err).ToNot(HaveOccurred())

	// Get the ChainID from the RPC.
	chainID, err := client.ChainID(context.Background())
	Expect(err).ToNot(HaveOccurred())

	// Build transaction opts object.
	auth, err := bind.NewKeyedTransactorWithChainID(network.ECDSATestKey, chainID)
	Expect(err).ToNot(HaveOccurred())
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 3_000_000  // in units
	auth.GasPrice = gasPrice
	return auth
}

// ExpectedMined waits for a transaction to be mined.
func ExpectMined(client *ethclient.Client, tx *coretypes.Transaction) {
	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
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
