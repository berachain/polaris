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

package core

import (
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
)

// `blockchain` is the canonical, persistent object that operates the Stargazer EVM.
type blockchain struct {
	// `StateProcessor` is the canonical, persistent state processor that runs the EVM.
	*StateProcessor
	// `host` is the host chain that the Stargazer EVM is running on.
	host StargazerHostChain
}

// `NewChain` creates and returns a `api.Chain` with the given EVM chain configuration and host.
func NewChain(host StargazerHostChain) *blockchain { //nolint:revive // temp.
	bc := &blockchain{
		host: host,
	}
	bc.StateProcessor = bc.buildStateProcessor(vm.Config{}, true)
	return bc
}

// `Host` returns the host chain that the Stargazer EVM is running on.
func (bc *blockchain) Host() StargazerHostChain {
	return bc.host
}

// `buildStateProcessor` builds and returns a `StateProcessor` with the given EVM configuration and
// commit flag.
func (bc *blockchain) buildStateProcessor(vmConfig vm.Config, commit bool) *StateProcessor {
	return NewStateProcessor(bc.host, state.NewStateDB(bc.host.GetStatePlugin()), vmConfig, commit)
}

func (bc *blockchain) CurrentHeader() *types.StargazerHeader {
	return bc.StateProcessor.block.StargazerHeader
}
func (bc *blockchain) CurrentBlock() *types.StargazerBlock {
	return bc.StateProcessor.block
}
