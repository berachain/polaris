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
	"context"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"cosmossdk.io/core/appmodule"

	"github.com/berachain/polaris/cosmos/x/evm/keeper"
	"github.com/berachain/polaris/cosmos/x/evm/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// ConsensusVersion defines the current x/evm module consensus version.
const ConsensusVersion = 1

var (
	_ appmodule.HasServices          = AppModule{}
	_ appmodule.HasPrepareCheckState = AppModule{}
	_ appmodule.HasEndBlocker        = AppModule{}
	_ module.AppModule               = AppModule{}
	_ module.AppModuleBasic          = AppModuleBasic{}
)

// ==============================================================================
// AppModuleBasic
// ==============================================================================

// AppModuleBasic defines the basic application module used by the evm module.
type AppModuleBasic struct{}

// Name returns the evm module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the evm module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {
	// types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types.
func (b AppModuleBasic) RegisterInterfaces(r cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(r)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the evm module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *gwruntime.ServeMux) {}

// GetTxCmd returns no root tx command for the evm module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the root query command for the evm module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// ==============================================================================
// AppModule
// ==============================================================================

// AppModule implements an application module for the evm module.
type AppModule struct {
	AppModuleBasic
	keeper    *keeper.Keeper
	accKeeper AccountKeeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(
	keeper *keeper.Keeper,
	ak AccountKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		accKeeper:      ak,
	}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// RegisterInvariants registers the evm module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterMsgServiceServer(registrar, am.keeper)
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// PrepareCheckState prepares the application state for a check.
func (am AppModule) PrepareCheckState(ctx context.Context) error {
	return am.keeper.PrepareCheckState(ctx)
}

// Precommit performs precommit operations.
func (am AppModule) EndBlock(ctx context.Context) error {
	return am.keeper.EndBlock(ctx)
}
