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

package miner

import (
	"context"

	evmkeeper "github.com/berachain/polaris/cosmos/x/evm/keeper"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

// EnvelopeSerializer is used to convert an envelope into a byte slice that represents
// a cosmos sdk.Tx.
type (
	EnvelopeSerializer interface {
		ToSdkTxBytes(*engine.ExecutionPayloadEnvelope, uint64) ([]byte, error)
	}

	App interface {
		BeginBlocker(sdk.Context) (sdk.BeginBlock, error)
		PreBlocker(sdk.Context, *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error)
		TxDecode(txBytes []byte) (sdk.Tx, error)
	}

	// EVMKeeper is an interface that defines the methods needed for the EVM setup.
	EVMKeeper interface {
		// Setup initializes the EVM keeper.
		Setup(evmkeeper.WrappedBlockchain) error
		SetLatestQueryContext(context.Context) error
	}
)
