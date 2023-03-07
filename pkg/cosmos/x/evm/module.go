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
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/pkg/cosmos/x/evm/types"
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
	types.RegisterInterfaces(r)
}

// `RegisterGRPCGatewayRoutes` registers the gRPC Gateway routes for the evm module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	// if err := types.RegisterQueryServiceHandlerClient(context.Background(), mux,
	// types.NewQueryClient(clientCtx)); err != nil {
	// 	panic(err)
	// }
	// evmrpc.RegisterJSONRPCServer(clientCtx, mux, app.EVMKeeper.GetRPCProvider()) maybe here?

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
	types.RegisterMsgServiceServer(cfg.MsgServer(), am.keeper)
	// types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// `ConsensusVersion` implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// `BeginBlock` returns the begin blocker for the evm module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	am.keeper.BeginBlocker(ctx)
}

// `EndBlock` returns the end blocker for the evm module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	am.keeper.EndBlocker(ctx)
	return []abci.ValidatorUpdate{}
}
