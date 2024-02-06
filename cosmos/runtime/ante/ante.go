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
	"errors"

	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/crypto/keys/ethsecp256k1"
	"github.com/berachain/polaris/cosmos/runtime/txpool"
	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Provider is a struct that holds the ante handlers for EVM and Cosmos.
type Provider struct {
	evmAnteHandler    sdk.AnteHandler // Ante handler for EVM transactions
	cosmosAnteHandler sdk.AnteHandler // Ante handler for Cosmos transactions
}

// NewAnteHandler creates a new Provider with a mempool and Cosmos ante handler.
// It sets up the EVM ante handler with the necessary decorators.
func NewAnteHandler(
	mempool *txpool.Mempool, cosmosAnteHandler sdk.AnteHandler,
) *Provider {
	evmAnteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // Set up the context decorator for the EVM ante handler
		mempool,                         // Add the mempool as a decorator for the EVM ante handler
	}

	// Return a new Provider with the set up EVM and provided Cosmos ante handlers
	return &Provider{
		evmAnteHandler:    sdk.ChainAnteDecorators(evmAnteDecorators...),
		cosmosAnteHandler: cosmosAnteHandler,
	}
}

// AnteHandler returns a function that handles ante for both EVM and Cosmos transactions.
// It checks the type of the transaction and calls the appropriate ante handler.
func (ah *Provider) AnteHandler() func(
	ctx sdk.Context, tx sdk.Tx, simulate bool,
) (sdk.Context, error) {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		// If the transaction contains a single EVM transaction, use the EVM ante handler
		if len(tx.GetMsgs()) == 1 {
			if _, ok := tx.GetMsgs()[0].(*evmtypes.WrappedEthereumTransaction); ok {
				return ah.evmAnteHandler(ctx, tx, simulate)
			} else if _, ok = tx.GetMsgs()[0].(*evmtypes.WrappedPayloadEnvelope); ok {
				if ctx.ExecMode() != sdk.ExecModeCheck {
					return ctx, nil
				}
				return ctx, errors.New("payload envelope is not supported in CheckTx")
			}
		}
		// Otherwise, use the Cosmos ante handler
		return ah.cosmosAnteHandler(ctx, tx, simulate)
	}
}

// EthSecp256k1SigVerificationGasConsumer is a function that consumes gas for the verification
// of an Ethereum Secp256k1 signature.
func EthSecp256k1SigVerificationGasConsumer(
	meter storetypes.GasMeter, sig signing.SignatureV2, params authtypes.Params,
) error {
	switch sig.PubKey.(type) {
	case *ethsecp256k1.PubKey:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: ethsecp256k1")
		return nil
	default:
		return ante.DefaultSigVerificationGasConsumer(meter, sig, params)
	}
}
