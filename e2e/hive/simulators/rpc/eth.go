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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/core/types"
	"gotest.tools/assert"
)

// func connectMultipleClients(t *TestEnv) {

// }

var (
	ctx = context.Background()
)

func chainIDSupport(t *TestEnv) {
	var (
		expectedChainID = big.NewInt(7) //nolint:gomnd // TODO: REFACTOR.
	)

	cID, err := t.Eth.ChainID(t.Ctx())
	assert.NilError(t, err, "could not get chain ID: %w", err)

	if expectedChainID.Cmp(cID) != 0 {
		t.Fatalf("expected chain ID %d, got %d", expectedChainID, cID)
	}
}

func gasPriceSupport(t *TestEnv) {
	gasPrice, err := t.Eth.SuggestGasPrice(ctx)
	assert.NilError(t, err, "could not get gasPrice: %w", err)
	if gasPrice == nil {
		t.Fatalf("gasPrice is nil")
	}
}

func blockNumberSupport(t *TestEnv) {
	blockNumber, err := t.Eth.BlockNumber(ctx)
	assert.NilError(t, err, "could not get blockNumber: %w", err)
	if blockNumber <= 0 {
		t.Fatalf("blockNumber <= 0, got: %d", blockNumber)
	}
}

func getBalanceSupport(t *TestEnv) {
	// balance, err := t.Eth.BalanceAt(ctx, , nil)

	var mySigner = types.NewLondonSigner(big.NewInt(5))
	key, err := crypto.GenerateKey()
	if err != nil {
		panic("sdfs")
	}

	types.SignNewTx(key, mySigner, &types.LegacyTx{})
	fmt.Printf("acount: %v\n", mySigner)
}

func estimateGasSupport(t *TestEnv) {

}

func getTransactionByHash(t *TestEnv) {

}
