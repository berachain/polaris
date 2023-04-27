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

package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/txpool"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/params"
)

// PolarisTxPool defines the required functions for the txpool.
type PolarisTxPool interface {
	AddLocal(tx *types.Transaction) error
	AddLocals(txs []*types.Transaction) []error
	AddRemote(tx *types.Transaction) error
	AddRemotes(txs []*types.Transaction) []error
	AddRemotesSync(txs []*types.Transaction) []error
	Content() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	ContentFrom(addr common.Address) (types.Transactions, types.Transactions)
	GasPrice() *big.Int
	Get(hash common.Hash) *types.Transaction
	Has(hash common.Hash) bool
	Locals() []common.Address
	Nonce(addr common.Address) uint64
	Pending(enforceTips bool) map[common.Address]types.Transactions
	RemoveTx(hash common.Hash, outofbound bool) int
	SetGasPrice(price *big.Int)
	Stats() (int, int)
	Status(hashes []common.Hash) []txpool.TxStatus
	Stop()
	SubscribeNewTxsEvent(ch chan<- NewTxsEvent) event.Subscription
}

// NewPolarisTxPool returns a new PolarisTxPool with the default txpool config.
func NewPolarisTxPool(chainConfig *params.ChainConfig, bc txpool.BlockChain) PolarisTxPool {
	return txpool.NewTxPool(txpool.DefaultConfig, chainConfig, bc)
}
