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

package keeper

import (
	"fmt"
	"math/big"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/core"
	enclib "pkg.berachain.dev/polaris/lib/encoding"
	"pkg.berachain.dev/polaris/lib/utils"
)

// InitGenesis is called during the InitGenesis.
func (k *Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) error {
	// We configure the logger here because we want to get the logger off the context opposed to allocating a new one.
	k.ConfigureGethLogger(ctx)

	if err := k.validateEthGenesis(ctx, genState); err != nil {
		return err
	}

	// Initialize all the plugins.
	for _, plugin := range k.host.GetAllPlugins() {
		// checks whether plugin implements methods of HasGenesis and executes them if it does
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			plugin.InitGenesis(ctx, &genState)
		}
	}

	go func() {
		time.Sleep(2 * time.Second) //nolint:gomnd // we will fix this eventually.

		// Start the polaris "Node" in order to spin up things like the JSON-RPC server.
		if err := k.polaris.StartServices(); err != nil {
			// TODO: figure out how to propagate errors out of goroutine if that is a problem
			return
		}
	}()

	return nil
}

// ExportGenesis returns the exported genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesisState := new(types.GenesisState)
	for _, plugin := range k.host.GetAllPlugins() {
		if plugin, ok := utils.GetAs[plugins.HasGenesis](plugin); ok {
			plugin.ExportGenesis(ctx, genesisState)
		}
	}
	return genesisState
}

// validateEthGenesis is a InitGenesis helper which asserts that the ethereum genesis state is
// consistent with the cosmos genesis state.
func (k *Keeper) validateEthGenesis(ctx sdk.Context, genesisState types.GenesisState) error {
	// ethGenesis stores the unmarshalled ethereum genesis state
	ethGenesis := enclib.MustUnmarshalJSON[core.Genesis]([]byte(genesisState.Params.EthGenesis))

	// verify gas limit matches
	if ethGenesis.GasLimit != ctx.BlockGasMeter().Limit() {
		return fmt.Errorf("gas limit mismatch: expected %d, got %d", ethGenesis.GasLimit, ctx.GasMeter().Limit())
	}

	// verify chainID matches
	ctxChainId, _ := new(big.Int).SetString(ctx.ChainID(), 10)
	if result := ethGenesis.Config.ChainID.Cmp(ctxChainId); result != 0 {
		return fmt.Errorf("chainID mismatch: expected %d, got %d", ethGenesis.Config.ChainID, ctxChainId)
	}

	// verify coinbase and timestamp matches
	coinbase, timestamp := k.GetHost().GetBlockPlugin().GetNewBlockMetadata(ctx.BlockHeight())
	if ethGenesis.Coinbase != coinbase {
		return fmt.Errorf("coinbase mismatch: expected %s, got %s", ethGenesis.Coinbase, coinbase)
	}
	if ethGenesis.Timestamp != timestamp {
		return fmt.Errorf("timestamp mismatch: expected %d, got %d", ethGenesis.Timestamp, timestamp)
	}

	// verify balances match
	denom := genesisState.Params.EvmDenom
	for addr, account := range ethGenesis.Alloc {
		// no need to check for missing denom, since a missing denom will return 0 balance
		coin := k.bk.GetBalance(ctx, lib.AddressToAccAddress(addr), denom)
		if coin.Amount != math.NewIntFromBigInt(account.Balance) {
			return fmt.Errorf("account %s balance mismatch: expected %s, got %s", addr.Hex(), account.Balance, coin.Amount)
		}
	}
	return nil
}
