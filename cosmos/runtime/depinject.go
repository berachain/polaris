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

package runtime

import (
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	"pkg.berachain.dev/polaris/cosmos/config"
	evmkeeper "pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	ethlog "pkg.berachain.dev/polaris/eth/log"
	"pkg.berachain.dev/polaris/eth/polar"
)

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	Logger    log.Logger
	EVMKeeper *evmkeeper.Keeper
	Config    func() *config.Config
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out

	Polaris *Polaris
}

// New creates a new Polaris runtime from the provided
// dependencies.
func New(input DepInjectInput) DepInjectOutput {
	cfg := input.Config()
	node, err := polar.NewGethNetworkingStack(&cfg.Node)
	if err != nil {
		panic(err)
	}

	polaris := polar.NewWithNetworkingStack(
		&cfg.Polar, input.EVMKeeper.Host, node, ethlog.FuncHandler(
			func(r *ethlog.Record) error {
				polarisGethLogger := input.Logger.With("module", "polaris-geth")
				switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
				case ethlog.LvlTrace:
				case ethlog.LvlDebug:
					polarisGethLogger.Debug(r.Msg, r.Ctx...)
				case ethlog.LvlInfo:
					polarisGethLogger.Info(r.Msg, r.Ctx...)
				case ethlog.LvlWarn:
				case ethlog.LvlCrit:
				case ethlog.LvlError:
					polarisGethLogger.Error(r.Msg, r.Ctx...)
				}
				return nil
			}),
	)

	return DepInjectOutput{
		Polaris: &Polaris{
			Polaris:   polaris,
			EVMKeeper: input.EVMKeeper,
		},
	}
}
