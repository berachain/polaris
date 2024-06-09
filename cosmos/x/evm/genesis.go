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

package evm

import (
	"encoding/json"

	"github.com/berachain/polaris/eth/core"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns default genesis state as raw bytes for the evm
// module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	ethGen := core.DefaultGenesis
	rawGenesis, err := ethGen.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return rawGenesis
}

// ValidateGenesis performs genesis state validation for the evm module.
func (AppModuleBasic) ValidateGenesis(
	_ codec.JSONCodec,
	_ client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	ethGen := new(core.Genesis)
	if err := ethGen.UnmarshalJSON(bz); err != nil {
		return err
	}

	// TODO: figure out what in geth we need to call in order to validate the genesis.

	return nil
}

// InitGenesis performs genesis initialization for the evm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(
	ctx sdk.Context,
	_ codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var ethGen core.Genesis
	if err := ethGen.UnmarshalJSON(data); err != nil {
		panic(err)
	}

	if err := am.keeper.InitGenesis(ctx, &ethGen); err != nil {
		panic(err)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the evm
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) json.RawMessage {
	ethGen := am.keeper.ExportGenesis(ctx)
	ethGenBz, err := ethGen.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return ethGenBz
}
