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
	"math/big"
	"sync"

	"cosmossdk.io/log"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	ethtxpool "github.com/ethereum/go-ethereum/core/txpool"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Mempool implements the mempool.Mempool & Lifecycle interfaces.
var (
	_ mempool.Mempool = (*Mempool)(nil)
	_ Lifecycle       = (*Mempool)(nil)
)

// Lifecycle represents a lifecycle object.
type Lifecycle interface {
	Start() error
	Stop() error
}

// GethTxPool is used for generating mocks.
type GethTxPool interface {
	eth.TxPool
}

// Mempool is a mempool that adheres to the cosmos mempool interface.
// It purposefully does not implement `Select` or `Remove` as the purpose of this mempool
// is to allow for transactions coming in from CometBFT's gossip to be added to the underlying
// geth txpool during `CheckTx`, that is the only purpose of `Mempool“.
type Mempool struct {
	eth.TxPool
	lifetime       int64
	chain          core.ChainReader
	handler        Lifecycle
	crc            CometRemoteCache
	blockBuilderMu *sync.RWMutex
	priceLimit     *big.Int
}

// New creates a new Mempool.
func New(
	chain core.ChainReader, txpool eth.TxPool, lifetime int64,
	blockBuilderMu *sync.RWMutex, priceLimit *big.Int,
) *Mempool {
	return &Mempool{
		TxPool:         txpool,
		chain:          chain,
		lifetime:       lifetime,
		crc:            newCometRemoteCache(),
		blockBuilderMu: blockBuilderMu,
		priceLimit:     priceLimit,
	}
}

// Init initializes the Mempool (notably the TxHandler).
func (m *Mempool) Init(
	logger log.Logger,
	txBroadcaster TxBroadcaster,
	txSerializer TxSerializer,
) {
	m.handler = newHandler(txBroadcaster, m.TxPool, txSerializer, m.crc, logger)
}

// Start starts the Mempool TxHandler.
func (m *Mempool) Start() error {
	return m.handler.Start()
}

// Stop stops the Mempool TxHandler.
func (m *Mempool) Stop() error {
	return m.handler.Stop()
}

// Insert attempts to insert a Tx into the app-side mempool returning an error upon failure.
func (m *Mempool) Insert(ctx context.Context, sdkTx sdk.Tx) error {
	sCtx := sdk.UnwrapSDKContext(ctx)
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return errors.New("only one message is supported")
	}

	wet, ok := utils.GetAs[*types.WrappedEthereumTransaction](msgs[0])
	if !ok {
		// We have to return nil for non-ethereum transactions as to not fail check-tx.
		return nil
	}

	// Add the eth tx to the Geth txpool.
	ethTx := wet.Unwrap()

	// Insert the tx into the txpool as a remote.
	m.blockBuilderMu.RLock()
	errs := m.TxPool.Add([]*ethtypes.Transaction{ethTx}, false, false)
	m.blockBuilderMu.RUnlock()

	// Handle case where a node broadcasts to itself, we don't want it to fail CheckTx.
	// Note: it's safe to check errs[0] because geth returns `errs` of length 1.
	if errors.Is(errs[0], ethtxpool.ErrAlreadyKnown) &&
		(sCtx.ExecMode() == sdk.ExecModeCheck || sCtx.ExecMode() == sdk.ExecModeReCheck) {
		telemetry.IncrCounter(float32(1), MetricKeyMempoolKnownTxs)
		sCtx.Logger().Info("mempool insert: tx already in mempool", "mode", sCtx.ExecMode())
		return nil
	} else if errs[0] != nil {
		return errs[0]
	}

	// Add the eth tx to the remote cache.
	_ = m.crc.MarkRemoteSeen(ethTx.Hash())

	return nil
}

// CountTx returns the number of transactions currently in the mempool.
func (m *Mempool) CountTx() int {
	runnable, blocked := m.TxPool.Stats()
	return runnable + blocked
}

// Select is an intentional no-op as we use a custom prepare proposal.
func (m *Mempool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

// Remove is an intentional no-op as the eth txpool handles removals.
func (m *Mempool) Remove(_ sdk.Tx) error {
	return nil
}
