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
	"math/big"
	"time"

	bindings "github.com/berachain/polaris/contracts/bindings/testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	//lint:ignore ST1001 Gomega makes sense in tests
	. "github.com/onsi/gomega" //nolint:stylecheck,revive,gostaticcheck // Gomega makes sense in tests
)

const (
	DefaultTimeout = 15 * time.Second
	TxTimeout      = 30 * time.Second
)

// ExpectedMined waits for a transaction to be mined.
func ExpectMined(client *ethclient.Client, tx *ethtypes.Transaction) {
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
	tx *ethtypes.Transaction,
) *ethtypes.Receipt {
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

// ExpectFailedReceipt waits for the transaction to be mined and returns the receipt.
// It also checks that the transaction was failed.
func ExpectFailedReceipt(
	client *ethclient.Client,
	tx *ethtypes.Transaction,
) *ethtypes.Receipt {
	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	_, err := bind.WaitMined(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())

	// Verify the receipt is good but status failed.
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	Expect(receipt.Status).To(Equal(uint64(0x0))) //nolint:gomnd // fail.
	return receipt
}

// DeployERC20 deploys a new ERC20 contract and waits for the transaction to be mined.
// Upon success, it returns a binding to the contract and the address of the contract.
func DeployERC20(
	auth *bind.TransactOpts,
	client *ethclient.Client,
) (*bindings.SolmateERC20, common.Address) {
	// Deploy the contract
	expectedAddr, tx, contract, err := bindings.DeploySolmateERC20(auth, client)
	Expect(err).ToNot(HaveOccurred())

	// Wait for the transaction to be mined.
	ctx, cancel := context.WithTimeout(context.Background(), TxTimeout)
	defer cancel()

	_, err = bind.WaitMined(ctx, client, tx)
	Expect(err).ToNot(HaveOccurred())

	time.Sleep(500 * time.Millisecond) //nolint:gomnd // temporary.
	code, err := client.CodeAt(ctx, expectedAddr, big.NewInt(-1))
	Expect(err).ToNot(HaveOccurred())
	Expect(code).ToNot(BeEmpty())

	return contract, expectedAddr
}
