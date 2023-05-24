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

package chain

import (
	"context"

	"pkg.berachain.dev/polaris/eth/api"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// blockProducer is the block producer.
type blockProducer struct {
	polaris         api.Chain
	currentBlockNum uint64
}

// ProduceBlock produces a block from the mempool and returns it.
func (bp *blockProducer) ProduceBlock() error {
	bp.currentBlockNum++

	ctx := context.Background()
	// Prepare Polaris for a new block.
	bp.polaris.Prepare(ctx, bp.currentBlockNum)

	// TODO: get from mempool.
	txs := make(types.Transactions, 0)

	// Iterate through all the transactions in the mempool.
	for _, txn := range txs {
		_, err := bp.polaris.ProcessTransaction(context.Background(), txn)
		if err != nil {
			return err
		}
	}

	// Finalize the block.
	err := bp.polaris.Finalize(ctx)
	if err != nil {
		return err
	}

	return nil
}
