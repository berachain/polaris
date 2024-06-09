// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
