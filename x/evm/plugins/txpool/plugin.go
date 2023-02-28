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
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	errorslib "pkg.berachain.dev/stargazer/lib/errors"
	mempool "pkg.berachain.dev/stargazer/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/stargazer/x/evm/rpc"
)

// `Plugin` represents the transaction pool plugin.
var _ Plugin = (*plugin)(nil)

// `Plugin` represents the transaction pool plugin.
type Plugin interface {
	core.TxPoolPlugin
}

// `plugin` represents the transaction pool plugin.
type plugin struct {
	mempool     *mempool.EthTxPool
	rpcProvider rpc.Provider
}

// `NewPlugin` returns a new transaction pool plugin.
func NewPlugin(rpcProvider rpc.Provider, ethTxMempool *mempool.EthTxPool) Plugin {
	return &plugin{
		mempool:     ethTxMempool,
		rpcProvider: rpcProvider,
	}
}

// `SendTx` sends a transaction to the transaction pool. It takes in a signed
// ethereum transaction from the rpc backend and wraps it in a Cosmos
// transaction. The Cosmos transaction is then broadcasted to the network.
func (p *plugin) SendTx(signedEthTx *coretypes.Transaction) error {
	// Serialize the transaction to Bytes
	txBytes, err := NewSerializer(p.rpcProvider.GetClientCtx()).Serialize(signedEthTx)
	if err != nil {
		return errorslib.Wrap(err, "failed to serialize transaction")
	}

	// Send the transaction to the CometBFT mempool, which will
	// gossip it to peers via CometBFT's p2p layer.
	syncCtx := p.rpcProvider.GetClientCtx().WithBroadcastMode(flags.BroadcastSync)

	rsp, err := syncCtx.BroadcastTx(txBytes)
	fmt.Println("ABCI RESP", rsp)
	if rsp != nil && rsp.Code != 0 {
		err = errorsmod.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog)
	}
	if err != nil {
		// b.logger.Error("failed to broadcast tx", "error", err.Errsor())
		return err
	}
	return nil
}

// `SendPrivTx` sends a private transaction to the transaction pool. It takes in
// a signed ethereum transaction from the rpc backend and wraps it in a Cosmos
// transaction. The Cosmos transaction is injected into the local mempool, but is
// NOT gossiped to peers.
func (p *plugin) SendPrivTx(signedTx *coretypes.Transaction) error {
	cosmosTx, err := NewSerializer(p.rpcProvider.GetClientCtx()).SerializeToSdkTx(signedTx)
	if err != nil {
		return err
	}

	// We insert into the local mempool, without gossiping to peers.
	// We use a blank sdk.Context{} as the context, as we don't need to
	// use it anyways. We set the priority as the gas price of the tx.
	return p.mempool.Insert(sdk.Context{}.WithPriority(signedTx.GasPrice().Int64()), cosmosTx)
}

// `GetAllTransactions` returns all transactions in the transaction pool.
func (p *plugin) GetAllTransactions() (coretypes.Transactions, error) {
	return p.mempool.GetPoolTransactions(), nil
}

// `GetTransactions` returns the transaction by hash in the transaction pool.
func (p *plugin) GetTransaction(hash common.Hash) *coretypes.Transaction {
	return p.mempool.GetTransaction(hash)
}

func (p *plugin) GetNonce(addr common.Address) uint64 {
	// TODO: implement this
	return 0
}
