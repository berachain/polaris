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

package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/cosmos/config"
	"github.com/berachain/polaris/cosmos/runtime/txpool"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth/core"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// host represents the host chain
	*Host

	// provider is the struct that houses the Polaris EVM.
	chain  core.Blockchain
	txpool *txpool.Mempool
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	ak state.AccountKeeper,
	storeKey storetypes.StoreKey,
	pcs func() *ethprecompile.Injector,
	qc func() func(height int64, prove bool) (sdk.Context, error),
	polarisCfg *config.Config,
) *Keeper {
	host := NewHost(
		*polarisCfg,
		storeKey,
		ak,
		pcs,
		qc,
	)
	return &Keeper{
		Host: host,
	}
}

func (k *Keeper) Setup(chain core.Blockchain, txPool *txpool.Mempool) error {
	k.chain = chain
	k.txpool = txPool
	return k.SetupPrecompiles()
}

func (k *Keeper) GetHost() core.PolarisHostChain {
	return k.Host
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(types.ModuleName)
}
