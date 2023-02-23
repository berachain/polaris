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
	"github.com/cosmos/cosmos-sdk/client/flags"

	crypto "pkg.berachain.dev/stargazer/crypto"
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
	mempool.EthTxPool
	rp rpc.Provider
}

func NewPlugin(rp rpc.Provider) Plugin {
	return &plugin{
		rp: rp,
	}
}

// `SendTx` sends a transaction to the transaction pool. It takes in a signed
// ethereum transaction from the rpc backend and wraps it in a Cosmos
// transaction. The Cosmos transaction is then broadcasted to the network.
func (p *plugin) SendTx(signedTx *coretypes.Transaction) error {
	clientCtx := p.rp.GetClientCtx()

	txBuilder, err := txpoolclient.NewEthTxBuilder(clientCtx)
	if err != nil {
		return err
	}

	// TODO: get evm denom from params.
	cosmosTx, err := txBuilder.BuildTx(signedTx, "abera")
	if err != nil {
		return err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		// b.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return err
	}

	syncCtx := clientCtx.WithBroadcastMode(flags.BroadcastSync)
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

func (p *plugin) GetAllTransactions() (coretypes.Transactions, error) {
	return coretypes.Transactions{}, nil
}

func (p *plugin) GetTransaction(hash common.Hash) *coretypes.Transaction {
	return nil
}

func (p *plugin) GetNonce(addr common.Address) uint64 {
	return 0
}

// `PubkeyFromTx` returns the public key of the signer of the transaction.
func PubkeyFromTx(signedTx *coretypes.Transaction, signer coretypes.Signer) (crypto.EthSecp256K1PubKey, error) {
	pk, err := signer.PubKey(signedTx)
	if err != nil {
		return crypto.EthSecp256K1PubKey{}, err
	}
	return crypto.EthSecp256K1PubKey{Key: pk}, nil
}
