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

package vm

import (
	"github.com/berachain/stargazer/eth/params"
)

// `EVMFactory` is used to build new Stargazer `EVM`s.
type EVMFactory struct {
	// `precompileManager` is responsible for keeping track of the stateful precompile
	// containers that are available to the EVM and executing them.
	precompileManager PrecompileManager
}

// `NewEVMFactory` creates and returns a new `EVMFactory` with the given `precompileManager`.
func NewEVMFactory(precompileManager PrecompileManager) *EVMFactory {
	return &EVMFactory{
		precompileManager: precompileManager,
	}
}

// `Build` creates and returns a new `vm.StargazerEVM`.
func (ef *EVMFactory) Build(
	ssdb StargazerStateDB,
	blockCtx BlockContext,
	chainConfig *params.EthChainConfig,
	noBaseFee bool,
) StargazerEVM {
	return NewStargazerEVM(
		blockCtx, TxContext{}, ssdb, chainConfig, Config{}, ef.precompileManager)
}
