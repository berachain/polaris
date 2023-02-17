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

	"github.com/berachain/stargazer/x/evm/keeper"
	"github.com/berachain/stargazer/x/evm/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

// `ConsensusVersion` defines the current x/evm module consensus version.
const ConsensusVersion = 1

var (
	_ module.HasServices         = AppModule{}
	_ module.BeginBlockAppModule = AppModule{}
	_ module.EndBlockAppModule   = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
)

// ==============================================================================
// AppModuleBasic
// ==============================================================================

// `AppModuleBasic` defines the basic application module used by the evm module.
type AppModuleBasic struct{}

// `Name` returns the evm module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// `RegisterLegacyAminoCodec` registers the evm module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// types.RegisterLegacyAminoCodec(cdc)
}

// `RegisterInterfaces` registers the module's interface types.
func (b AppModuleBasic) RegisterInterfaces(r cdctypes.InterfaceRegistry) {
	// types.RegisterInterfaces(r)
}

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

// `RegisterGRPCGatewayRoutes` registers the gRPC Gateway routes for the evm module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	// if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
	// 	panic(err)
	// }
}

// `GetTxCmd` returns no root tx command for the evm module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// `GetQueryCmd` returns the root query command for the evm module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// ==============================================================================
// AppModule
// ==============================================================================

// `AppModule` implements an application module for the evm module.
type AppModule struct {
	AppModuleBasic
	keeper     *keeper.Keeper
	accKeeper  AccountKeeper
	bankKeeper BankKeeper
}

// `NewAppModule` creates a new AppModule object.
func NewAppModule(
	keeper *keeper.Keeper,
	ak AccountKeeper,
	bk BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		accKeeper:      ak,
		bankKeeper:     bk,
	}
}

// `IsOnePerModuleType` implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// `IsAppModule` implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// `RegisterInvariants` registers the evm module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// `RegisterServices` registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	// types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// `InitGenesis` performs genesis initialization for the evm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	// var genesisState types.GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)
	return []abci.ValidatorUpdate{}
}

// `ExportGenesis` returns the exported genesis state as raw bytes for the evm
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	// gs := am.keeper.ExportGenesis(ctx)
	// return cdc.MustMarshalJSON(gs)
	return json.RawMessage{}
}

// `ConsensusVersion` implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// `BeginBlock` returns the begin blocker for the evm module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	// BeginBlocker(ctx, am.keeper, am.inflationCalculator)
}

// `EndBlock` returns the end blocker for the evm module. It returns no validator
// updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
