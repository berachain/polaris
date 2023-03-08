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

package evm

import (
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/rpc/api"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
)

// `DefaultGenesis` returns default genesis state as raw bytes for the evm
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// `ValidateGenesis` performs genesis state validation for the evm module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(data)
}

// `InitGenesis` performs genesis initialization for the evm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	// We configure the logger here because we want to get the logger off the context opposed to allocating a new one.
	am.keeper.ConfigureGethLogger(ctx)

	// Initialize all the plugins.
	for _, plugin := range am.keeper.GetAllPlugins() {
		plugin.InitGenesis(ctx, &genesisState)
	}

	// TODO: Clean this up its really jank, we should move this, feels like a bad spot for it, but works.
	// Currently since we are registering TransactionAPI using the native ethereum backend, it needs to be able to
	// read the chainID from the ConfigurationPlugin (which is on disk). If we are enabling the APIs before
	// InitGenesis is called, then we get a nil pointer error since the ConfigurationPlugin is not yet initialized.
	if err := am.keeper.GetRPCProvider().RegisterAPIs(api.GetExtraFn); err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{}
}

// `ExportGenesis` returns the exported genesis state as raw bytes for the evm
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesisState := new(types.GenesisState)
	for _, plugin := range am.keeper.GetAllPlugins() {
		plugin.ExportGenesis(ctx, genesisState)
	}
	return cdc.MustMarshalJSON(genesisState)
}
