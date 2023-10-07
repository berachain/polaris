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

	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Mempool implements the mempool.Mempool & Lifecycle interfaces.
var (
	_ mempool.Mempool = (*Mempool)(nil)
	_ Lifecycle       = (*Mempool)(nil)
)

// GethTxPool represents the interface to interact with the geth txpool.
type GethTxPool interface {
	Add([]*coretypes.Transaction, bool, bool) []error
	Stats() (int, int)
	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
}

// Lifecycle represents a lifecycle object.
type Lifecycle interface {
	Start() error
	Stop() error
}

// Mempool is a mempool that adheres to the cosmos mempool interface.
// It purposefully does not implement `Select` or `Remove` as the purpose of this mempool
// is to allow for transactions coming in from CometBFT's gossip to be added to the underlying
// geth txpool during `CheckTx`, that is the only purpose of `Mempool“.
type Mempool struct {
	txpool  GethTxPool
	handler Lifecycle
}

// NewMempool creates a new Mempool.
func New(txpool GethTxPool) *Mempool {
	return &Mempool{
		txpool: txpool,
	}
}

// Init intializes the Mempool (notably the TxHandler).
func (m *Mempool) Init(
	logger log.Logger,
	txBroadcaster TxBroadcaster,
	txSerializer TxSerializer,
) {
	m.handler = newHandler(txBroadcaster, m.txpool, txSerializer, logger)
}

// Start starts the Mempool TxHandler.
func (m *Mempool) Start() error {
	return m.handler.Start()
}

// Stop stops the Mempool TxHandler.
func (m *Mempool) Stop() error {
	return m.handler.Stop()
}

// Insert attempts to insert a Tx into the app-side mempool returning
// an error upon failure.
func (m *Mempool) Insert(ctx context.Context, sdkTx sdk.Tx) error {
	sCtx := sdk.UnwrapSDKContext(ctx)
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return errors.New("only one message is supported")
	}

	if wet, ok := utils.GetAs[*types.WrappedEthereumTransaction](msgs[0]); !ok {
		return errors.New("only WrappedEthereumTransactions are supported")
	} else if errs := m.txpool.Add(
		[]*coretypes.Transaction{wet.Unwrap()}, false, false,
	); len(errs) != 0 {
		// Handle case where a node broadcasts to itself, we don't want it to fail CheckTx.
		if errors.Is(errs[0], legacypool.ErrAlreadyKnown) && sCtx.IsCheckTx() {
			return nil
		}
		return errs[0]
	}

	return nil
}

// CountTx returns the number of transactions currently in the mempool.
func (m *Mempool) CountTx() int {
	runnable, blocked := m.txpool.Stats()
	return runnable + blocked
}

// Select is an intentional no-op as we use a custom prepare proposal.
func (m *Mempool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

// Remove is an intentional no-op as the eth txpool handles removals.
func (m *Mempool) Remove(sdk.Tx) error {
	return nil
}
