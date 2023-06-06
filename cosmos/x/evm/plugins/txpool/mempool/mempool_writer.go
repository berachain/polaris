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

package mempool

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"

	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// Insert is called when a transaction is added to the mempool.
func (gtp *WrappedGethTxPool) Insert(_ context.Context, tx sdk.Tx) error {
	// TODO: @calbera we might want to call with sync here? idk if this could cause problems.
	// TODO: trying sync on CI rn
	return gtp.AddRemotes(coretypes.Transactions{evmtypes.GetAsEthTx(tx)})[0]
}

// InsertSync is called when a transaction is added to the mempool. (For testing purposes).
func (gtp *WrappedGethTxPool) InsertSync(_ context.Context, tx sdk.Tx) error {
	return gtp.AddRemotesSync(coretypes.Transactions{evmtypes.GetAsEthTx(tx)})[0]
}

// Remove is called when a transaction is removed from the mempool.
func (gtp *WrappedGethTxPool) Remove(tx sdk.Tx) error {
	if ethTx := evmtypes.GetAsEthTx(tx); ethTx != nil {
		removed := gtp.RemoveTx(ethTx.Hash(), true)
		if removed < 1 {
			// TODO: RemoveTx will return 0 if the tx was removed from future queue. We should
			// handle this specific case better.
			return sdkmempool.ErrTxNotFound
		}
	}
	return nil
}
