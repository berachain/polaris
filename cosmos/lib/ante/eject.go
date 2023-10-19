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

package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewAnteHandler creates a new instance of AnteHandler with EjectOnRecheckTxDecorator.
func NewAnteHandler() sdk.AnteHandler {
	anteDecorators := []sdk.AnteDecorator{
		&EjectOnRecheckTxDecorator{},
	}

	return sdk.ChainAnteDecorators(anteDecorators...)
}

// EjectOnRecheckTxDecorator will return an error if the context is a recheck tx.
// This is used to forcibly eject transactions from the CometBFT mempool after they
// have been passed down to the application, as we want to prevent the comet mempool
// from growing in size.
type EjectOnRecheckTxDecorator struct{}

// Antehandle implements sdk.AnteHandler.
func (EjectOnRecheckTxDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	var newCtx sdk.Context
	if ctx.IsReCheckTx() {
		return ctx, fmt.Errorf("recheck tx")
	}

	return next(newCtx, tx, simulate)
}
