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

package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

// func connectMultipleClients(t *TestEnv) {

// }

var (
	ctx = context.Background()
)

// func chainIDSupport(t *TestEnv) {
// 	var (
// 		expectedChainID = big.NewInt(7) //nolint:gomnd // TODO: REFACTOR.
// 	)

// 	cID, err := t.Eth.ChainID(t.Ctx())
// 	assert.NilError(t, err, "could not get chain ID: %w", err)

// 	if expectedChainID.Cmp(cID) != 0 {
// 		t.Fatalf("expected chain ID %d, got %d", expectedChainID, cID)
// 	}
// }

// func gasPriceSupport(t *TestEnv) {
// 	gasPrice, err := t.Eth.SuggestGasPrice(ctx)
// 	assert.NilError(t, err, "could not get gasPrice: %w", err)
// 	if gasPrice == nil {
// 		t.Fatalf("gasPrice is nil")
// 	}
// }

// func blockNumberSupport(t *TestEnv) {
// 	blockNumber, err := t.Eth.BlockNumber(ctx)
// 	assert.NilError(t, err, "could not get blockNumber: %w", err)
// 	if blockNumber <= 0 {
// 		t.Fatalf("blockNumber <= 0, got: %d", blockNumber)
// 	}
// }

// func getBalanceSupport(t *TestEnv) {
// 	addr := t.Vault.createAccount(t, big.NewInt(5))
// 	blockNumber, _ := t.Eth.BlockNumber(ctx)
// 	balance, err := t.Eth.BalanceAt(ctx, addr, big.NewInt(int64(blockNumber)))
// 	assert.NilError(t, err, "could not get balance: %w", err)
// 	if balance.Cmp(new(big.Int)) != 1 {
// 		t.Fatalf("balance <= 0, got: %d", balance)
// 	}
// }

// func estimateGasSupport(t *TestEnv) {

// }

// func getTransactionByHash(t *TestEnv) {

// }

// TransactionReceiptTest sends a transaction and tests the receipt fields.
func TransactionReceiptTest(t *TestEnv) {
	var (
		_ = t.Vault.createAccount(t, big.NewInt(params.Ether))
	)

	// rawTx := types.NewTransaction(uint64(0), common.Address{}, big.NewInt(1), 100000, gasPrice, nil)
	// tx, err := t.Vault.signTransaction(key, rawTx)
	// if err != nil {
	// 	t.Fatalf("Unable to sign deploy tx: %v", err)
	// }

	// if err = t.Eth.SendTransaction(t.Ctx(), tx); err != nil {
	// 	t.Fatalf("Unable to send transaction: %v", err)
	// }

	// for i := 0; i < 60; i++ {
	// 	receipt, err := t.Eth.TransactionReceipt(t.Ctx(), tx.Hash())
	// 	if err == ethereum.NotFound {
	// 		time.Sleep(time.Second)
	// 		continue
	// 	}

	// 	if err != nil {
	// 		t.Errorf("Unable to fetch receipt: %v", err)
	// 	}
	// 	if receipt.TxHash != tx.Hash() {
	// 		t.Errorf("Receipt [tx=%x] contains invalid tx hash, want %x, got %x", tx.Hash(), receipt.TxHash)
	// 	}
	// 	if receipt.ContractAddress != (common.Address{}) {
	// 		t.Errorf("Receipt [tx=%x] contains invalid contract address, expected empty address but got %x", tx.Hash(), receipt.ContractAddress)
	// 	}
	// 	if receipt.Bloom.Big().Cmp(new(big.Int)) != 0 {
	// 		t.Errorf("Receipt [tx=%x] bloom not empty, %x", tx.Hash(), receipt.Bloom)
	// 	}
	// 	if receipt.GasUsed != params.TxGas {
	// 		t.Errorf("Receipt [tx=%x] has invalid gas used, want %d, got %d", tx.Hash(), params.TxGas, receipt.GasUsed)
	// 	}
	// 	if len(receipt.Logs) != 0 {
	// 		t.Errorf("Receipt [tx=%x] should not contain logs but got %d logs", tx.Hash(), len(receipt.Logs))
	// 	}
	// 	return
	// }
}
