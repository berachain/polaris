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

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	txpoolclient "pkg.berachain.dev/stargazer/x/evm/plugins/txpool/client"
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
	mempool   mempool.EthTxPool
	clientCtx client.Context
}

// `NewPlugin` returns a new transaction pool plugin.
func NewPlugin(rp rpc.Provider) Plugin {
	return &plugin{
		clientCtx: rp.GetClientCtx(),
	}
}

// `SendTx` sends a transaction to the transaction pool. It takes in a signed
// ethereum transaction from the rpc backend and wraps it in a Cosmos
// transaction. The Cosmos transaction is then broadcasted to the network.
func (p *plugin) SendTx(signedTx *coretypes.Transaction) error {
	// Serialize the transaction.
	txBytes, err := p.EthTransactionToTxBytes(signedTx)
	if err != nil {
		return err
	}

	// Send the transaction to the CometBFT mempool, which will
	// gossip it to peers via CometBFT's p2p layer.
	syncCtx := p.clientCtx.WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
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
func (p *plugin) SendPrivTx(signedEthTx *coretypes.Transaction) error {
	txBuilder, err := txpoolclient.NewEthTxBuilder(p.clientCtx)
	if err != nil {
		return err
	}

	// TODO: get evm denom from params.
	cosmosTx, err := txBuilder.BuildTx(signedEthTx, "abera")
	if err != nil {
		return err
	}

	// We insert into the local mempool, without gossiping to peers.
	// We use a blank sdk.Context{} as the context, as we don't need to
	// use it anyways. We set the priority as the gas price of the tx.
	return p.mempool.Insert(
		sdk.Context{}.WithPriority(signedEthTx.GasPrice().Int64()),
		cosmosTx,
	)
}

// `EthTransactionToTxBytes` converts an ethereum transaction to txBytes which allows for the to
// broadcast it to CometBFT.
func (p *plugin) EthTransactionToTxBytes(signedEthTx *coretypes.Transaction) ([]byte, error) {
	cosmosTx, err := p.EthTransactionToCosmosTx(signedEthTx)
	if err != nil {
		return nil, err
	}

	txBytes, err := p.clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		// b.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return nil, err
	}
	return txBytes, nil
}

// `EthTransactionToCosmosTx` converts an ethereum transaction to a Cosmos
// transaction.
func (p *plugin) EthTransactionToCosmosTx(signedEthTx *coretypes.Transaction) (sdk.Tx, error) {
	txBuilder, err := txpoolclient.NewEthTxBuilder(p.clientCtx)
	if err != nil {
		return nil, err
	}

	// TODO: get evm denom from params.
	cosmosTx, err := txBuilder.BuildTx(signedEthTx, "abera")
	if err != nil {
		return nil, err
	}
	return cosmosTx, nil
}

// `GetAllTransactions` returns all transactions in the transaction pool.
func (p *plugin) GetAllTransactions() (coretypes.Transactions, error) {
	return coretypes.Transactions{}, nil
}

// `GetTransactions` returns all transactions in the transaction pool.
func (p *plugin) GetTransaction(hash common.Hash) *coretypes.Transaction {
	return nil
}

// `GetTransactions` returns all transactions in the transaction pool.
func (p *plugin) GetNonce(addr common.Address) uint64 {
	return 0
}
