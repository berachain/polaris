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
