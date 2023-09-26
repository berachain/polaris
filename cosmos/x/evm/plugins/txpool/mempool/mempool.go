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
	"github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/core/txpool"
)

var _ mempool.Mempool = (*WrappedGethTxPool)(nil)

type WrappedGethTxPool struct {
	*txpool.TxPool
}

// func (wgtp *WrappedGethTxPool) commitTransactions() error {

// }

func (wgtp *WrappedGethTxPool) CountTx() int {
	return 0
}

func (wgtp *WrappedGethTxPool) Insert(_ context.Context, _ sdk.Tx) error {
	return nil
}

func (wgtp *WrappedGethTxPool) Select(context.Context, [][]byte) mempool.Iterator {
	return nil
}

func (wgtp *WrappedGethTxPool) Remove(sdk.Tx) error {
	// intentionally a no-op
	return nil
}
