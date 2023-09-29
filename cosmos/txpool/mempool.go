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

package txpool

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Mempool implements the mempool.Mempool interface.
var _ mempool.Mempool = (*Mempool)(nil)

// GethTxPool represents the interface to interact with the geth txpool.
type GethTxPool interface {
	Add([]*coretypes.Transaction, bool, bool) []error
	Stats() (int, int)
}

// Mempool is a mempool that adheres to the cosmos mempool interface.
// It purposefully does not implement `Select` or `Remove` as the purpose of this mempool
// is to allow for transactions coming in from CometBFT's gossip to be added to the underlying
// geth txpool during `CheckTx`, that is the only purpose of `Mempool“.
type Mempool struct {
	txpool GethTxPool
}

// Insert attempts to insert a Tx into the app-side mempool returning
// an error upon failure.
func (m *Mempool) Insert(_ context.Context, sdkTx sdk.Tx) error {
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return errors.New("only one message is supported")
	}

	if wet, ok := (msgs[0]).(*types.WrappedEthereumTransaction); !ok {
		return errors.New("only WrappedEthereumTransactions are supported")
	} else if errs := m.txpool.Add([]*coretypes.Transaction{wet.AsTransaction()}, false, false); len(errs) != 0 {
		return errs[0]
	}

	return nil
}

// SetTxPool sets the underlying txpool.
func (m *Mempool) SetTxPool(txpool GethTxPool) {
	m.txpool = txpool
}

// CountTx returns the number of transactions currently in the mempool.
func (m *Mempool) CountTx() int {
	runnable, blocked := m.txpool.Stats()
	return runnable + blocked
}

// Select is an intention no-op as we use a custom prepare proposal.
func (m *Mempool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

// Remove is an intentional no-op as the eth txpool handles removals.
func (m *Mempool) Remove(sdk.Tx) error {
	return nil
}
