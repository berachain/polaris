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

package antelib

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/lib/utils"
)

// IgnoreDecorator is an AnteDecorator that wraps an existing AnteDecorator. It allows
// EthTransactions to skip said Decorator by checking to see if the transaction
// contains a message of the given type.
type IgnoreDecorator[D sdk.AnteDecorator, M sdk.Msg] struct {
	decorator D
}

// NewIgnoreDecorator returns a new IgnoreDecorator[D, M] instance.
func NewIgnoreDecorator[D sdk.AnteDecorator, M sdk.Msg](decorator D) *IgnoreDecorator[D, M] {
	return &IgnoreDecorator[D, M]{
		decorator: decorator,
	}
}

// AnteHandle implements the sdk.AnteDecorator interface, it is handle the
// type check for the message type.
func (sd IgnoreDecorator[D, M]) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if _, ok := utils.GetAs[M](msg); ok {
			return next(ctx, tx, simulate)
		}
	}

	return sd.decorator.AnteHandle(ctx, tx, simulate, next)
}
