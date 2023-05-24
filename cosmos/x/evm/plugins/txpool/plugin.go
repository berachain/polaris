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
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gofrs/uuid"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	mempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// Compile-time type assertion.
var _ Plugin = (*plugin)(nil)

// Plugin defines the required functions of the transaction pool plugin.
type Plugin interface {
	plugins.Base
	core.TxPoolPlugin
	SetNonceRetriever(mempool.NonceRetriever)
	SetClientContext(client.Context)
}

// plugin represents the transaction pool plugin.
type plugin struct {
	*mempool.EthTxPool

	clientContext client.Context
	cp            ConfigurationPlugin

	// txFeed and scope is used to send new batch transactions to new txs subscribers when the
	// batch is added to the mempool.
	txFeed event.Feed
	scope  event.SubscriptionScope
}

// NewPlugin returns a new transaction pool plugin.
func NewPlugin(cp ConfigurationPlugin, ethTxMempool *mempool.EthTxPool) Plugin {
	return &plugin{
		EthTxPool: ethTxMempool,
		cp:        cp,
	}
}

// SetClientContext implements the Plugin interface.
func (p *plugin) SetClientContext(ctx client.Context) {
	p.clientContext = ctx
}

// SubscribeNewTxsEvent returns a new event subscription for the new txs feed.
func (p *plugin) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return p.scope.Track(p.txFeed.Subscribe(ch))
}

// SendTx sends a transaction to the transaction pool. It takes in a signed Ethereum transaction
// from the rpc backend and wraps it in a Cosmos transaction. The Cosmos transaction is then
// broadcasted to the network.
func (p *plugin) SendTx(signedEthTx *coretypes.Transaction) error {
	// Serialize the transaction to Bytes
	txBytes, err := SerializeToBytes(p.cp.GetEvmDenom(), p.clientContext, signedEthTx)
	if err != nil {
		return errorslib.Wrap(err, "failed to serialize transaction")
	}

	// Send the transaction to the CometBFT mempool, which will gossip it to peers via CometBFT's
	// p2p layer.
	syncCtx := p.clientContext.WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
	if rsp != nil && rsp.Code != 0 {
		err = errorsmod.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog)
	}
	if err != nil {
		// b.logger.Error("failed to broadcast tx", "error", err.Errsor())
		return err
	}

	// Currently sending an individual new txs event for every new tx added to the mempool via
	// broadcast.
	// TODO: support sending batch new txs events when adding queued txs to the pending txs.
	p.txFeed.Send(core.NewTxsEvent{Txs: coretypes.Transactions{signedEthTx}})

	return nil
}

// SendPrivTx sends a private transaction to the transaction pool. It takes in a signed ethereum
// transaction from the rpc backend and wraps it in a Cosmos transaction. The Cosmos transaction is
// injected into the local mempool, but is NOT gossiped to peers.
func (p *plugin) SendPrivTx(signedTx *coretypes.Transaction) error {
	cosmosTx, err := SerializeToSdkTx(p.cp.GetEvmDenom(), p.clientContext, signedTx)
	if err != nil {
		return err
	}

	// We insert into the local mempool, without gossiping to peers. We use a blank sdk.Context{}
	// as the context, as we don't need to use it anyways. We set the priority as the gas price of
	// the tx.
	return p.EthTxPool.Insert(sdk.Context{}.WithPriority(signedTx.GasPrice().Int64()), cosmosTx)
}

func (p *plugin) SendBundle(txs types.Transactions, blockNumber int64, uuid uuid.UUID, signingAddress common.Address, minTimestamp uint64, maxTimestamp uint64, revertingTxHashes []common.Hash) error {
	// TODO
	return nil
}

func (p *plugin) IsPlugin() {}
